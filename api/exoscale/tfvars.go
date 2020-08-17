package exoscale

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/elastisys/ck8s/api"
)

type ExoscaleMachine struct {
	api.Machine `mapstructure:",squash"`

	ESLocalStorageCapacity int `json:"es_local_storage_capacity" mapstructure:"es_local_storage_capacity"`
}

type ExoscaleTFVars struct {
	PrefixSC string `json:"prefix_sc" mapstructure:"prefix_sc"`
	PrefixWC string `json:"prefix_wc" mapstructure:"prefix_wc"`

	MachinesSC map[string]ExoscaleMachine `json:"machines_sc" mapstructure:"machines_sc" validate:"required,min=1"`
	MachinesWC map[string]ExoscaleMachine `json:"machines_wc" mapstructure:"machines_wc" validate:"required,min=1"`

	NFSSize string `json:"nfs_size" mapstructure:"nfs_size" validate:"required"`

	PublicIngressCIDRWhitelist []string `json:"public_ingress_cidr_whitelist" mapstructure:"public_ingress_cidr_whitelist" validate:"required"`

	APIServerWhitelist []string `json:"api_server_whitelist" mapstructure:"api_server_whitelist" validate:"required"`
	NodeportWhitelist  []string `json:"nodeport_whitelist" mapstructure:"nodeport_whitelist" validate:"required"`
}

func (e *Cluster) Machines() map[string]api.Machine {
	machines := map[string]api.Machine{}

	for name, machine := range e.machinesByClusterType() {
		machines[name] = machine.Machine
	}

	return machines
}

func (e *Cluster) CloneMachine(name string) (string, error) {
	machines := e.machinesByClusterType()

	cloneName := uuid.New().String()

	machine, ok := machines[name]
	if !ok {
		return "", fmt.Errorf("machine not found: %s", name)
	}

	machines[cloneName] = machine

	return cloneName, nil
}

func (e *Cluster) RemoveMachine(name string) error {
	delete(e.machinesByClusterType(), name)
	return nil
}

func (e *Cluster) machinesByClusterType() map[string]ExoscaleMachine {
	switch e.config.ClusterType {
	case api.ServiceCluster:
		return e.tfvars.MachinesSC
	case api.WorkloadCluster:
		return e.tfvars.MachinesWC
	}

	panic("invalid cluster type")
}
