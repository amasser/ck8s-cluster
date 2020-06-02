package runner

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"go.uber.org/zap"
)

var NodeNotFoundErr = errors.New("kubernetes node not found")

type Kubectl struct {
	runner Runner

	kubeconfigPath string
	nodePrefix     string

	logger *zap.Logger
}

func NewKubectl(
	logger *zap.Logger,
	runner Runner,
	kubeconfigPath string,
	nodePrefix string,
) *Kubectl {
	return &Kubectl{
		runner: runner,

		kubeconfigPath: kubeconfigPath,
		nodePrefix:     nodePrefix,

		logger: logger.With(
			zap.String("kubeconfig", kubeconfigPath),
			zap.String("node_prefix", nodePrefix),
		),
	}
}

func (k *Kubectl) command(args ...string) *Command {
	return NewCommand(
		"sops",
		"exec-file", k.kubeconfigPath,
		fmt.Sprintf("KUBECONFIG={} kubectl %s", strings.Join(args, " ")),
	)
}

func (k *Kubectl) fullNodeName(name string) string {
	return fmt.Sprintf("%s%s", k.nodePrefix, name)
}

// NodeExists runs `sops exec-file KUBECONFIG 'kubectl get node NAME'` in the
// background. If the node is not found a NodeNotFoundErr error is returned.
func (k *Kubectl) NodeExists(name string) error {
	k.logger.Debug("kubectl_node_exists")

	cmd := k.command("get", "node", k.fullNodeName(name))

	// TODO: This assumes a lot about the command output. We should replace
	// 		 this with a proper Kubernetes client lib implementation ASAP.
	cmd.OutputHandler = func(stdoutPipe, stderrPipe io.Reader) error {
		stderr, err := ioutil.ReadAll(stderrPipe)
		if err != nil {
			return err
		}
		if strings.Contains(string(stderr), "not found") {
			return NodeNotFoundErr
		}
		return nil
	}

	return k.runner.Background(cmd)
}

// Drain runs `sops exec-file KUBECONFIG 'kubectl drain --ignore-daemonsets
// --delete-local-data NAME'`
func (k *Kubectl) Drain(name string) error {
	k.logger.Debug("kubectl_drain")
	return k.runner.Run(k.command(
		"drain",
		"--ignore-daemonsets",
		"--delete-local-data",
		k.fullNodeName(name),
	))
}

// DeleteNode runs `sops exec-file KUBECONFIG 'kubectl delete node NAME'`
func (k *Kubectl) DeleteNode(name string) error {
	k.logger.Debug("kubectl_delete_node")
	return k.runner.Run(k.command("delete", "node", k.fullNodeName(name)))
}