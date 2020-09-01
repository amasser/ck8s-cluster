package aws

import (
	"github.com/elastisys/ck8s/api"
)

const (
	FlavorDevelopment api.ClusterFlavor = "dev"
	FlavorProduction  api.ClusterFlavor = "prod"
)

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

func Production(clusterType api.ClusterType, clusterName string) api.Cluster {
	cluster := Default(clusterType, clusterName)

	cluster.tfvars.Region = "us-west-1"

	cluster.tfvars.MachinesSC = map[string]api.Machine{
		// Masters ------------------------------------
		"master-0": {
			NodeType: api.Master,
			Size:     "t3.small",
		},
		"master-1": {
			NodeType: api.Master,
			Size:     "t3.small",
		},
		"master-2": {
			NodeType: api.Master,
			Size:     "t3.small",
		},
		// Workers ------------------------------------
		// TODO:
		// - Safespring has 8 cores for the "extra-large" and 4 for "large"
		//   but here we have only 4 and 2 respectivly.
		// - Maybe we should switch to non-burstable instances?
		"worker-0": {
			NodeType: api.Worker,
			Size:     "t3.xlarge",
		},
		"worker-1": {
			NodeType: api.Worker,
			Size:     "t3.large",
		},
		"worker-2": {
			NodeType: api.Worker,
			Size:     "t3.large",
		},
		"worker-3": {
			NodeType: api.Worker,
			Size:     "t3.large",
		},
	}

	cluster.tfvars.MachinesWC = map[string]api.Machine{
		// Masters ------------------------------------
		"master-0": {
			NodeType: api.Master,
			Size:     "t3.small",
		},
		"master-1": {
			NodeType: api.Master,
			Size:     "t3.small",
		},
		"master-2": {
			NodeType: api.Master,
			Size:     "t3.small",
		},
		// Workers ------------------------------------
		"worker-ck8s-0": {
			NodeType: api.Worker,
			Size:     "t3.large",
		},
		"worker-0": {
			NodeType: api.Worker,
			Size:     "t3.large",
		},
		"worker-1": {
			NodeType: api.Worker,
			Size:     "t3.large",
		},
		"worker-2": {
			NodeType: api.Worker,
			Size:     "t3.large",
		},
	}

	return cluster
}
