package exoscale

import (
	"testing"

	"github.com/elastisys/ck8s/api"
	"github.com/google/go-cmp/cmp"
)

func TestFlavors(t *testing.T) {
	clusterType := api.ServiceCluster
	clusterName := "foo"

	type testCase struct {
		want, got api.Cluster
	}

	testCases := []testCase{{
		want: &Cluster{
			config: ExoscaleConfig{
				BaseConfig: api.BaseConfig{
					ClusterType:               clusterType,
					CloudProviderType:         api.Exoscale,
					EnvironmentName:           clusterName,
					DNSPrefix:                 clusterName,
					S3BucketNameHarbor:        clusterName + "-harbor",
					S3BucketNameVelero:        clusterName + "-velero",
					S3BucketNameElasticsearch: clusterName + "-es-backup",
					S3BucketNameInfluxDB:      clusterName + "-influxdb",
					S3BucketNameFluentd:       clusterName + "-sc-logs",
				},
				S3RegionAddress: "sos-ch-gva-2.exo.io",
			},
			secret: ExoscaleSecret{
				BaseSecret: api.BaseSecret{
					S3AccessKey: "changeme",
					S3SecretKey: "changeme",
				},
				APIKey:    "changeme",
				SecretKey: "changeme",
			},
			tfvars: ExoscaleTFVars{
				PublicIngressCIDRWhitelist: []string{},
				APIServerWhitelist:         []string{},
				NodeportWhitelist:          []string{},
			},
		},
		got: Default(clusterType, clusterName),
	}, {
		want: &Cluster{
			config: ExoscaleConfig{
				BaseConfig: api.BaseConfig{
					ClusterType:               clusterType,
					CloudProviderType:         api.Exoscale,
					EnvironmentName:           clusterName,
					DNSPrefix:                 clusterName,
					S3BucketNameHarbor:        clusterName + "-harbor",
					S3BucketNameVelero:        clusterName + "-velero",
					S3BucketNameElasticsearch: clusterName + "-es-backup",
					S3BucketNameInfluxDB:      clusterName + "-influxdb",
					S3BucketNameFluentd:       clusterName + "-sc-logs",
				},
				S3RegionAddress: "sos-ch-gva-2.exo.io",
			},
			secret: ExoscaleSecret{
				BaseSecret: api.BaseSecret{
					S3AccessKey: "changeme",
					S3SecretKey: "changeme",
				},
				APIKey:    "changeme",
				SecretKey: "changeme",
			},
			tfvars: ExoscaleTFVars{
				PublicIngressCIDRWhitelist: []string{},
				APIServerWhitelist:         []string{},
				NodeportWhitelist:          []string{},
				MachinesSC: map[string]ExoscaleMachine{
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
						ESLocalStorageCapacity: 12,
					},
					"worker-1": {
						Machine: api.Machine{
							NodeType: api.Worker,
							Size:     "Large",
						},
						ESLocalStorageCapacity: 12,
					},
				},
				MachinesWC: map[string]ExoscaleMachine{
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
				},
				NFSSize: "Small",
			},
		},
		got: Development(clusterType, clusterName),
	}, {
		want: &Cluster{
			config: ExoscaleConfig{
				BaseConfig: api.BaseConfig{
					ClusterType:               clusterType,
					CloudProviderType:         api.Exoscale,
					EnvironmentName:           clusterName,
					DNSPrefix:                 clusterName,
					S3BucketNameHarbor:        clusterName + "-harbor",
					S3BucketNameVelero:        clusterName + "-velero",
					S3BucketNameElasticsearch: clusterName + "-es-backup",
					S3BucketNameInfluxDB:      clusterName + "-influxdb",
					S3BucketNameFluentd:       clusterName + "-sc-logs",
				},
				S3RegionAddress: "sos-ch-gva-2.exo.io",
			},
			secret: ExoscaleSecret{
				BaseSecret: api.BaseSecret{
					S3AccessKey: "changeme",
					S3SecretKey: "changeme",
				},
				APIKey:    "changeme",
				SecretKey: "changeme",
			},
			tfvars: ExoscaleTFVars{
				PublicIngressCIDRWhitelist: []string{},
				APIServerWhitelist:         []string{},
				NodeportWhitelist:          []string{},

				MachinesSC: map[string]ExoscaleMachine{
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
						ESLocalStorageCapacity: 140,
					},
					"worker-2": {
						Machine: api.Machine{
							NodeType: api.Worker,
							Size:     "Large",
						},
						ESLocalStorageCapacity: 140,
					},
					"worker-3": {
						Machine: api.Machine{
							NodeType: api.Worker,
							Size:     "Large",
						},
						ESLocalStorageCapacity: 0,
					},
				},
				MachinesWC: map[string]ExoscaleMachine{
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
				},
				NFSSize: "Small",
			},
		},
		got: Production(clusterType, clusterName),
	}}

	for _, tc := range testCases {
		if diff := cmp.Diff(tc.want, tc.got, cmp.AllowUnexported(Cluster{})); diff != "" {
			t.Errorf("flavor mismatch (-want +got):\n%s", diff)
		}
	}
}
