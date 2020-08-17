package exoscale

import (
	"github.com/elastisys/ck8s/api"
)

const (
	FlavorDevelopment api.ClusterFlavor = "dev"
	FlavorProduction  api.ClusterFlavor = "prod"
)

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
			ESLocalStorageCapacity: 26,
		},
		"worker-1": {
			Machine: api.Machine{
				NodeType: api.Worker,
				Size:     "Large",
			},
			ESLocalStorageCapacity: 26,
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

	// TODO

	return cluster
}
