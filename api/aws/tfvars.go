package aws

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/elastisys/ck8s/api"
)

// AWSTFVars TODO
type AWSTFVars struct {
	Region string `json:"region" mapstructure:"region"`

	PrefixSC string `json:"prefix_sc" mapstructure:"prefix_sc"`
	PrefixWC string `json:"prefix_wc" mapstructure:"prefix_wc"`

	MachinesSC map[string]api.Machine `json:"machines_sc" mapstructure:"machines_sc" validate:"required,min=1"`
	MachinesWC map[string]api.Machine `json:"machines_wc" mapstructure:"machines_wc" validate:"required,min=1"`

	PublicIngressCIDRWhitelist []string `json:"public_ingress_cidr_whitelist" mapstructure:"public_ingress_cidr_whitelist" validate:"required"`
	APIServerWhitelist         []string `json:"api_server_whitelist" mapstructure:"api_server_whitelist" validate:"required"`
	NodeportWhitelist          []string `json:"nodeport_whitelist" mapstructure:"nodeport_whitelist" validate:"required"`
}

func (e *Cluster) CloneMachine(name string) (string, error) {
	machines := e.Machines()

	cloneName := uuid.New().String()

	machine, ok := machines[name]
	if !ok {
		return "", fmt.Errorf("machine not found: %s", name)
	}

	machines[cloneName] = machine

	return cloneName, nil
}

func (e *Cluster) Machines() (machines map[string]api.Machine) {
	switch e.config.ClusterType {
	case api.ServiceCluster:
		return e.tfvars.MachinesSC
	case api.WorkloadCluster:
		return e.tfvars.MachinesWC
	}

	panic("invalid cluster type")
}

func (e *Cluster) RemoveMachine(name string) error {
	machines := e.Machines()
	delete(machines, name)
	return nil
}
