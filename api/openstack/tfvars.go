package openstack

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/elastisys/ck8s/api"
)

type TFVars struct {
	PrefixSC string `json:"prefix_sc" mapstructure:"prefix_sc"`
	PrefixWC string `json:"prefix_wc" mapstructure:"prefix_wc"`

	MachinesSC map[string]api.Machine `json:"machines_sc" mapstructure:"machines_sc" validate:"required,min=1"`
	MachinesWC map[string]api.Machine `json:"machines_wc" mapstructure:"machines_wc" validate:"required,min=1"`

	MasterAntiAffinityPolicySC string `json:"master_anti_affinity_policy_sc" mapstructure:"master_anti_affinity_policy_sc"`
	WorkerAntiAffinityPolicySC string `json:"worker_anti_affinity_policy_sc" mapstructure:"worker_anti_affinity_policy_sc"`
	MasterAntiAffinityPolicyWC string `json:"master_anti_affinity_policy_wc" mapstructure:"master_anti_affinity_policy_wc"`
	WorkerAntiAffinityPolicyWC string `json:"worker_anti_affinity_policy_wc" mapstructure:"worker_anti_affinity_policy_wc"`

	PublicIngressCIDRWhitelist []string `json:"public_ingress_cidr_whitelist" mapstructure:"public_ingress_cidr_whitelist" validate:"required"`

	APIServerWhitelist []string `json:"api_server_whitelist" mapstructure:"api_server_whitelist" validate:"required"`
	NodeportWhitelist  []string `json:"nodeport_whitelist" mapstructure:"nodeport_whitelist" validate:"required"`

	AWSDNSZoneID  string `json:"aws_dns_zone_id" mapstructure:"aws_dns_zone_id" validate:"required"`
	AWSDNSRoleARN string `json:"aws_dns_role_arn" mapstructure:"aws_dns_role_arn" validate:"required"`

	ExternalNetworkID   string `json:"external_network_id" mapstructure:"external_network_id" validate:"required"`
	ExternalNetworkName string `json:"external_network_name" mapstructure:"external_network_name" validate:"required"`
}

func (e *Cluster) CloneMachine(name string) (string, error) {
	machines := e.Machines()

	// TODO Find the root cause for this issue
	cloneName := strings.Replace(uuid.New().String(), "-", "", -1)

	machine, ok := machines[name]
	if !ok {
		return "", fmt.Errorf("machine not found: %s", name)
	}

	machines[cloneName] = machine

	return cloneName, nil
}

func (e *Cluster) Machines() (machines map[string]api.Machine) {
	switch e.Config.ClusterType {
	case api.ServiceCluster:
		return e.TFVars.MachinesSC
	case api.WorkloadCluster:
		return e.TFVars.MachinesWC
	}

	panic("invalid cluster type")
}

func (e *Cluster) RemoveMachine(name string) error {
	machines := e.Machines()
	delete(machines, name)
	return nil
}
