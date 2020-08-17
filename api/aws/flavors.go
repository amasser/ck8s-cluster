package aws

import (
	"github.com/elastisys/ck8s/api"
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
		config: AWSConfig{
			BaseConfig: *api.DefaultBaseConfig(
				clusterType,
				api.AWS,
				clusterName,
			),
			S3Region: "us-west-1",
		},
		secret: AWSSecret{
			BaseSecret: api.BaseSecret{
				S3AccessKey: "changeme",
				S3SecretKey: "changeme",
			},
			AWSAccessKeyID:     "changeme",
			AWSSecretAccessKey: "changeme",
			DNSAccessKeyID:     "changeme",
			DNSSecretAccessKey: "changeme",
		},
		tfvars: AWSTFVars{
			PublicIngressCIDRWhitelist: []string{},
			APIServerWhitelist:         []string{},
			NodeportWhitelist:          []string{},
		},
	}
}

// Development TODO
func Development(clusterType api.ClusterType, clusterName string) api.Cluster {
	cluster := Default(clusterType, clusterName)

	cluster.tfvars.Region = "us-west-1"

	cluster.tfvars.MachinesSC = map[string]api.Machine{
		"master-0": {
			NodeType: api.Master,
			Size:     "t3.small",
		},
		"worker-0": {
			NodeType: api.Worker,
			Size:     "t3.xlarge",
		},
		"worker-1": {
			NodeType: api.Worker,
			Size:     "t3.large",
		},
	}

	cluster.tfvars.MachinesWC = map[string]api.Machine{
		"master-0": {
			NodeType: api.Master,
			Size:     "t3.small",
		},
		"worker-0": {
			NodeType: api.Worker,
			Size:     "t3.large",
		},
		// TODO Should we use two nodes here?
		"worker-1": {
			NodeType: api.Worker,
			Size:     "t3.large",
		},
	}

	return cluster
}

// Production TODO
func Production(clusterType api.ClusterType, clusterName string) api.Cluster {
	cluster := Default(clusterType, clusterName)

	// TODO

	return cluster
}
