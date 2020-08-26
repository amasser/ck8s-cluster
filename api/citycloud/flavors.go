package citycloud

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
		openstack.Default(clusterType, api.CityCloud, clusterName),
	}

	cluster.Cluster.Config.IdentityAPIVersion = "3"
	cluster.Cluster.Config.AuthURL = "https://kna1.citycloud.com:5000"
	cluster.Cluster.Config.RegionName = "Kna1"
	cluster.Cluster.Config.S3RegionAddress = "s3-kna1.citycloud.com:8080"

	cluster.Cluster.TFVars.ExternalNetworkID = "fba95253-5543-4078-b793-e2de58c31378"
	cluster.Cluster.TFVars.ExternalNetworkName = "ext-net"

	return cluster
}

// Development TODO
func Development(clusterType api.ClusterType, clusterName string) api.Cluster {
	cluster := Default(clusterType, clusterName)

	cluster.Cluster.TFVars.MachinesSC = map[string]api.Machine{
		"master-0": {
			NodeType: api.Master,
			// 2 cores 4GB mem 50GB storage
			Size: "96c7903e-32f0-421d-b6a2-a45c97b15665",
		},
		"worker-0": {
			NodeType: api.Worker,
			// 4 core 16GB mem 50GB storage
			Size: "d430b3cd-0216-43ff-878c-c08689c0001b",
		},
		"worker-1": {
			NodeType: api.Worker,
			// 4 core 8GB mem 50GB storage
			Size: "572a3b2e-6329-4053-b872-aecb1e70d8a6",
		},
	}

	cluster.Cluster.TFVars.MachinesWC = map[string]api.Machine{
		"master-0": {
			NodeType: api.Master,
			// 2 cores 4GB mem 50GB storage
			Size: "96c7903e-32f0-421d-b6a2-a45c97b15665",
		},
		"worker-0": {
			NodeType: api.Worker,
			// 4 core 8GB mem 50GB storage
			Size: "572a3b2e-6329-4053-b872-aecb1e70d8a6",
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
