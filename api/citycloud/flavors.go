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
	return &Cluster{
		config: openstack.OpenstackConfig{
			BaseConfig: *api.DefaultBaseConfig(
				clusterType,
				api.CityCloud,
				clusterName,
			),

			IdentityAPIVersion: "3",
			AuthURL:            "https://kna1.citycloud.com:5000",
			RegionName:         "Kna1",

			ProjectID:         "changeme",
			ProjectDomainName: "changeme",
			UserDomainName:    "changeme",

			S3RegionAddress: "s3-kna1.citycloud.com:8080",
		},
		secret: openstack.OpenstackSecret{
			BaseSecret: api.BaseSecret{
				S3AccessKey: "changeme",
				S3SecretKey: "changeme",
			},
			AWSAccessKeyID:     "changeme",
			AWSSecretAccessKey: "changeme",

			Username: "changeme",
			Password: "changeme",
		},
		tfvars: openstack.OpenstackTFVars{
			PublicIngressCIDRWhitelist: []string{},
			APIServerWhitelist:         []string{},
			NodeportWhitelist:          []string{},

			ExternalNetworkID:   "fba95253-5543-4078-b793-e2de58c31378",
			ExternalNetworkName: "ext-net",

			AWSDNSZoneID:  "changeme",
			AWSDNSRoleARN: "changeme",
		},
	}
}

// Development TODO
func Development(clusterType api.ClusterType, clusterName string) api.Cluster {
	cluster := Default(clusterType, clusterName)

	cluster.tfvars.MasterNamesSC = []string{"master-0"}
	cluster.tfvars.MasterNameSizeMapSC = map[string]string{
		"master-0": "96c7903e-32f0-421d-b6a2-a45c97b15665", // 2 Core 4gb mem 50gb storage
	}

	cluster.tfvars.WorkerNamesSC = []string{"worker-0", "worker-1"}
	cluster.tfvars.WorkerNameSizeMapSC = map[string]string{
		"worker-0": "d430b3cd-0216-43ff-878c-c08689c0001b", // 4 core 16gb mem 50gb storage
		"worker-1": "572a3b2e-6329-4053-b872-aecb1e70d8a6", // 4 core 8gb mem 50gb storage
	}
	cluster.tfvars.MasterNamesWC = []string{"master-0"}
	cluster.tfvars.MasterNameSizeMapWC = map[string]string{
		"master-0": "96c7903e-32f0-421d-b6a2-a45c97b15665", // 2 core 4gb mem 50gb storage
	}

	cluster.tfvars.WorkerNamesWC = []string{"worker-0"}
	cluster.tfvars.WorkerNameSizeMapWC = map[string]string{
		"worker-0": "572a3b2e-6329-4053-b872-aecb1e70d8a6", // 4 core 8gb mem 50gb storage
	}

	cluster.tfvars.MasterAntiAffinityPolicySC = "anti-affinity"
	cluster.tfvars.MasterAntiAffinityPolicyWC = "anti-affinity"

	return cluster
}

// Production TODO
func Production(clusterType api.ClusterType, clusterName string) api.Cluster {
	cluster := Default(clusterType, clusterName)

	// TODO

	return cluster
}
