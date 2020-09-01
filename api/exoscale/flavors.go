package exoscale

import (
	"github.com/elastisys/ck8s/api"
)

const (
	FlavorDevelopment api.ClusterFlavor = "dev"
	FlavorProduction  api.ClusterFlavor = "prod"
)

// Sizes
// Name            RAM      vCPUs
// ------------------------------
// SMALL          2 GB    2 Cores
// MEDIUM         4 GB    2 Cores
// LARGE          8 GB    4 Cores
// EXTRA-LARGE   16 GB    4 Cores
// HUGE          32 GB    8 Cores
// MEGA          64 GB   12 Cores
// TITAN        128 GB   16 Cores
// JUMBO        225 GB   24 Cores

func Default(clusterType api.ClusterType, clusterName string) *Cluster {
	return &Cluster{
		config: ExoscaleConfig{
			BaseConfig: *api.DefaultBaseConfig(
				clusterType,
				api.Exoscale,
				clusterName,
			),
			S3RegionAddress: "sos-ch-gva-2.exo.io",
		},
		secret: ExoscaleSecret{
			BaseSecret: *api.DefaultBaseSecret(),
			APIKey:     "changeme",
			SecretKey:  "changeme",
		},
		tfvars: ExoscaleTFVars{
			PublicIngressCIDRWhitelist: []string{},
			APIServerWhitelist:         []string{},
			NodeportWhitelist:          []string{},
		},
	}
}

func Development(clusterType api.ClusterType, clusterName string) api.Cluster {
	cluster := Default(clusterType, clusterName)

	cluster.tfvars.MachinesSC = map[string]ExoscaleMachine{
		"master-0": {
			Machine: api.Machine{
				NodeType: api.Master,
				Size:     "Small",
			},
		},
		"worker-0": {
			Machine: api.Machine{
				NodeType: api.Worker,
				Size:     "Extra-large",
			},
			// Match ES_DATA_STORAGE_SIZE in config.sh
			// Note that this value is in GB while config.sh uses Gi
			ESLocalStorageCapacity: 12,
		},
		"worker-1": {
			Machine: api.Machine{
				NodeType: api.Worker,
				Size:     "Large",
			},
			// Match ES_DATA_STORAGE_SIZE in config.sh
			// Note that this value is in GB while config.sh uses Gi
			ESLocalStorageCapacity: 12,
		},
	}

	cluster.tfvars.MachinesWC = map[string]ExoscaleMachine{
		"master-0": {
			Machine: api.Machine{
				NodeType: api.Master,
				Size:     "Small",
			},
		},
		"worker-0": {
			Machine: api.Machine{
				NodeType: api.Worker,
				Size:     "Large",
			},
		},
	}

	cluster.tfvars.NFSSize = "Small"

	return cluster
}

func Production(clusterType api.ClusterType, clusterName string) api.Cluster {
	cluster := Default(clusterType, clusterName)

	cluster.tfvars.MachinesSC = map[string]ExoscaleMachine{
		// Masters ------------------------------------
		"master-0": {
			Machine: api.Machine{
				NodeType: api.Master,
				Size:     "Medium",
			},
		},
		"master-1": {
			Machine: api.Machine{
				NodeType: api.Master,
				Size:     "Medium",
			},
		},
		"master-2": {
			Machine: api.Machine{
				NodeType: api.Master,
				Size:     "Medium",
			},
		},
		// Workers ------------------------------------
		// TODO:
		// - Safespring has 8 cores for the "extra-large" but here we have only 4
		// - How many nodes with local storage do we need?
		"worker-0": {
			Machine: api.Machine{
				NodeType: api.Worker,
				Size:     "Extra-large",
			},
			ESLocalStorageCapacity: 0,
		},
		"worker-1": {
			Machine: api.Machine{
				NodeType: api.Worker,
				Size:     "Large",
			},
			// Match ES_DATA_STORAGE_SIZE in config.sh
			// Note that this value is in GB while config.sh uses Gi
			ESLocalStorageCapacity: 140,
		},
		"worker-2": {
			Machine: api.Machine{
				NodeType: api.Worker,
				Size:     "Large",
			},
			// Match ES_DATA_STORAGE_SIZE in config.sh
			// Note that this value is in GB while config.sh uses Gi
			ESLocalStorageCapacity: 140,
		},
		"worker-3": {
			Machine: api.Machine{
				NodeType: api.Worker,
				Size:     "Large",
			},
			ESLocalStorageCapacity: 0,
		},
	}

	cluster.tfvars.MachinesWC = map[string]ExoscaleMachine{
		// Masters ------------------------------------
		"master-0": {
			Machine: api.Machine{
				NodeType: api.Master,
				Size:     "Medium",
			},
		},
		"master-1": {
			Machine: api.Machine{
				NodeType: api.Master,
				Size:     "Medium",
			},
		},
		"master-2": {
			Machine: api.Machine{
				NodeType: api.Master,
				Size:     "Medium",
			},
		},
		// Workers ------------------------------------
		"worker-ck8s-0": {
			Machine: api.Machine{
				NodeType: api.Worker,
				Size:     "Large",
			},
		},
		"worker-0": {
			Machine: api.Machine{
				NodeType: api.Worker,
				Size:     "Large",
			},
		},
		"worker-1": {
			Machine: api.Machine{
				NodeType: api.Worker,
				Size:     "Large",
			},
		},
		"worker-2": {
			Machine: api.Machine{
				NodeType: api.Worker,
				Size:     "Large",
			},
		},
	}

	cluster.tfvars.NFSSize = "Small"

	return cluster
}
