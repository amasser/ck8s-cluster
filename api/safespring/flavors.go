package safespring

import (
	"github.com/elastisys/ck8s/api"
	"github.com/elastisys/ck8s/api/openstack"
)

const (
	// FlavorDevelopment TODO
	FlavorDevelopment api.ClusterFlavor = "dev"

	// FlavorProduction TODO
	FlavorProduction api.ClusterFlavor = "prod"
)

// Default TODO
func Default(clusterType api.ClusterType, clusterName string) *Cluster {
	cluster := &Cluster{
		openstack.Default(clusterType, api.Safespring, clusterName),
	}

	cluster.Cluster.Config.IdentityAPIVersion = "3"
	cluster.Cluster.Config.AuthURL = "https://keystone.api.cloud.ipnett.se/v3"
	cluster.Cluster.Config.RegionName = "se-east-1"
	cluster.Cluster.Config.S3RegionAddress = "s3.sto1.safedc.net"

	cluster.Cluster.TFVars.ExternalNetworkID = "71b10496-2617-47ae-abbc-36239f0863bb"
	cluster.Cluster.TFVars.ExternalNetworkName = "public-v4"

	return cluster
}

// Development TODO
func Development(clusterType api.ClusterType, clusterName string) api.Cluster {
	cluster := Default(clusterType, clusterName)

	cluster.Cluster.TFVars.MachinesSC = map[string]api.Machine{
		"master-0": {
			NodeType: api.Master,
			// TODO: could go with smaller flavor here if made available
			// lb.large.1d
			Size: "dc67a9eb-0685-4bb6-9383-a01c717e02e8",
		},
		"worker-0": {
			NodeType: api.Worker,
			// lb.xlarge.1d
			Size: "ea0dbe3b-f93a-47e0-84e4-b09ec5873bdf",
		},
		"worker-1": {
			NodeType: api.Worker,
			// lb.large.1d
			Size: "dc67a9eb-0685-4bb6-9383-a01c717e02e8",
		},
		"loadbalancer-0": {
			NodeType: api.LoadBalancer,
			// lb.tiny
			Size: "51d480b8-2517-4ba8-bfe0-c649ac93eb61",
		},
	}

	cluster.Cluster.TFVars.MachinesWC = map[string]api.Machine{
		"master-0": {
			NodeType: api.Master,
			// TODO: could go with smaller flavor here if made available
			// lb.large.1d
			Size: "dc67a9eb-0685-4bb6-9383-a01c717e02e8",
		},
		"worker-0": {
			NodeType: api.Worker,
			// lb.large.1d
			Size: "dc67a9eb-0685-4bb6-9383-a01c717e02e8",
		},
		"loadbalancer-0": {
			NodeType: api.LoadBalancer,
			// lb.tiny
			Size: "51d480b8-2517-4ba8-bfe0-c649ac93eb61",
		},
	}

	cluster.Cluster.TFVars.MasterAntiAffinityPolicySC = "anti-affinity"
	cluster.Cluster.TFVars.MasterAntiAffinityPolicyWC = "anti-affinity"

	return cluster
}

// Production TODO
func Production(clusterType api.ClusterType, clusterName string) api.Cluster {
	cluster := Default(clusterType, clusterName)

	// TODO

	return cluster
}
