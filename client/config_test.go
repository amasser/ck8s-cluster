package client

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/elastisys/ck8s/api"
	"github.com/elastisys/ck8s/api/aws"
	"github.com/elastisys/ck8s/api/citycloud"
	"github.com/elastisys/ck8s/api/exoscale"
	"github.com/elastisys/ck8s/api/openstack"
	"github.com/elastisys/ck8s/api/safespring"
	"github.com/elastisys/ck8s/testutil"
)

func TestTFVarsRead(t *testing.T) {
	clusterType := api.ServiceCluster

	type testCase struct {
		path string
		want interface{}
		got  api.Cluster
	}

	for _, tc := range []testCase{{
		path: "testdata/exoscale-tfvars.json",
		want: &exoscale.ExoscaleTFVars{
			MachinesSC: map[string]exoscale.ExoscaleMachine{
				"master-0": {
					Machine: api.Machine{
						NodeType: api.Master,
						Size:     "Small",
					}},
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
			},
			MachinesWC: map[string]exoscale.ExoscaleMachine{
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
			NFSSize:                    "Small",
			PublicIngressCIDRWhitelist: []string{"1.2.3.4/32", "4.3.2.1/32"},
			APIServerWhitelist:         []string{"1.2.3.4/32", "4.3.2.1/32"},
			NodeportWhitelist:          []string{"1.2.3.4/32", "4.3.2.1/32"},
		},
		got: exoscale.Default(clusterType, ""),
	}, {
		path: "testdata/citycloud-tfvars.json",
		want: &openstack.TFVars{
			MachinesSC: map[string]api.Machine{
				"master-0": {
					NodeType: api.Master,
					Size:     "96c7903e-32f0-421d-b6a2-a45c97b15665",
				},
				"worker-0": {
					NodeType: api.Worker,
					Size:     "d430b3cd-0216-43ff-878c-c08689c0001b",
				},
				"worker-1": {
					NodeType: api.Worker,
					Size:     "572a3b2e-6329-4053-b872-aecb1e70d8a6",
				},
			},
			MachinesWC: map[string]api.Machine{
				"master-0": {
					NodeType: api.Master,
					Size:     "96c7903e-32f0-421d-b6a2-a45c97b15665",
				},
				"worker-0": {
					NodeType: api.Worker,
					Size:     "572a3b2e-6329-4053-b872-aecb1e70d8a6",
				},
			},
			MasterAntiAffinityPolicySC: "anti-affinity",
			MasterAntiAffinityPolicyWC: "anti-affinity",
			PublicIngressCIDRWhitelist: []string{"1.2.3.4/32", "4.3.2.1/32"},
			APIServerWhitelist:         []string{"1.2.3.4/32", "4.3.2.1/32"},
			NodeportWhitelist:          []string{"1.2.3.4/32", "4.3.2.1/32"},
			ExternalNetworkID:          "2aec7a99-3783-4e2a-bd2b-bbe4fef97d1c",
			ExternalNetworkName:        "ext-net",
			AWSDNSZoneID:               "testAWSDNSZoneID",
			AWSDNSRoleARN:              "testAWSDNSRoleARN",
		},
		got: citycloud.Default(clusterType, ""),
	}, {
		path: "testdata/safespring-tfvars.json",
		want: &openstack.TFVars{
			MachinesSC: map[string]api.Machine{
				"master-0": {
					NodeType: api.Master,
					Size:     "dc67a9eb-0685-4bb6-9383-a01c717e02e8",
				},
				"worker-0": {
					NodeType: api.Worker,
					Size:     "ea0dbe3b-f93a-47e0-84e4-b09ec5873bdf",
				},
				"worker-1": {
					NodeType: api.Worker,
					Size:     "dc67a9eb-0685-4bb6-9383-a01c717e02e8",
				},
				"loadbalancer-0": {
					NodeType: api.LoadBalancer,
					Size:     "51d480b8-2517-4ba8-bfe0-c649ac93eb61",
				},
			},
			MachinesWC: map[string]api.Machine{
				"master-0": {
					NodeType: api.Master,
					Size:     "dc67a9eb-0685-4bb6-9383-a01c717e02e8",
				},
				"worker-0": {
					NodeType: api.Worker,
					Size:     "dc67a9eb-0685-4bb6-9383-a01c717e02e8",
				},
				"loadbalancer-0": {
					NodeType: api.LoadBalancer,
					Size:     "51d480b8-2517-4ba8-bfe0-c649ac93eb61",
				},
			},
			MasterAntiAffinityPolicySC: "anti-affinity",
			MasterAntiAffinityPolicyWC: "anti-affinity",
			WorkerAntiAffinityPolicySC: "anti-affinity",
			WorkerAntiAffinityPolicyWC: "anti-affinity",
			PublicIngressCIDRWhitelist: []string{"1.2.3.4/32", "4.3.2.1/32"},
			APIServerWhitelist:         []string{"1.2.3.4/32", "4.3.2.1/32"},
			NodeportWhitelist:          []string{"1.2.3.4/32", "4.3.2.1/32"},
			ExternalNetworkID:          "2aec7a99-3783-4e2a-bd2b-bbe4fef97d1c",
			ExternalNetworkName:        "ext-net",
			AWSDNSZoneID:               "testAWSDNSZoneID",
			AWSDNSRoleARN:              "testAWSDNSRoleARN",
		},
		got: safespring.Default(clusterType, ""),
	}, {
		path: "testdata/aws-tfvars.json",
		want: &aws.AWSTFVars{
			Region: "us-west-1",
			MachinesSC: map[string]api.Machine{
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
			},
			MachinesWC: map[string]api.Machine{
				"master-0": {
					NodeType: api.Master,
					Size:     "t3.small",
				},
				"worker-0": {
					NodeType: api.Worker,
					Size:     "t3.large",
				},
				"worker-1": {
					NodeType: api.Worker,
					Size:     "t3.large",
				},
			},
			PublicIngressCIDRWhitelist: []string{"1.2.3.4/32", "4.3.2.1/32"},
			APIServerWhitelist:         []string{"1.2.3.4/32", "4.3.2.1/32"},
			NodeportWhitelist:          []string{"1.2.3.4/32", "4.3.2.1/32"},
		},
		got: aws.Default(clusterType, ""),
	}} {
		logTest, logger := testutil.NewTestLogger([]string{
			"config_handler_tfvars_read",
		})

		configHandler := NewConfigHandler(
			logger,
			clusterType,
			api.ConfigPath{
				api.TFVarsFile: {
					Path:   tc.path,
					Format: "json",
				},
			},
			api.CodePath{},
		)

		if err := configHandler.readTFVars(tc.got); err != nil {
			t.Fatalf("error reading tfvars (%s): %s", tc.path, err)
		}

		if diff := cmp.Diff(tc.want, tc.got.TFVars()); diff != "" {
			t.Errorf("%s mismatch (-want +got):\n%s", tc.path, diff)
		}

		logTest.Diff(t)
	}
}
