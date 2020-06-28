package aws

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/elastisys/ck8s/api"
)

func testState(
	t *testing.T,
	clusterType api.ClusterType,
	clusterName string,
) api.ClusterState {
	switch clusterType {
	case api.ServiceCluster:
		clusterName += "-service-cluster"
	case api.WorkloadCluster:
		clusterName += "-workload-cluster"
	}

	tfOutput := &terraformOutput{
		ClusterType: clusterType,
		ClusterName: clusterName,
	}

	data, err := ioutil.ReadFile("testdata/terraform-output.json")
	if err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(data, &tfOutput); err != nil {
		t.Fatal(err)
	}

	return tfOutput
}

func TestTerraformOutputBaseDomain(t *testing.T) {
	testCases := map[api.ClusterType]string{
		api.ServiceCluster:  "", // TODO This isn't part of aws terraform atm
		api.WorkloadCluster: "",
	}

	for clusterType, want := range testCases {
		tfOutput := testState(t, clusterType, "ck8stest")

		got := tfOutput.BaseDomain()
		if got != want {
			t.Errorf(
				"Base domain mismatch, want: %s, got: %s",
				want, got,
			)
		}
	}
}

func TestTerraformOutputControlPlanePublicIP(t *testing.T) {
	testCases := map[api.ClusterType]string{
		api.ServiceCluster:  "tf-lb-20200624083157273500000008-1544241826.us-west-1.elb.amazonaws.com",
		api.WorkloadCluster: "tf-lb-20200624083144986100000006-1183550514.us-west-1.elb.amazonaws.com",
	}

	for clusterType, want := range testCases {
		tfOutput := testState(t, clusterType, "ck8stest")

		got := tfOutput.ControlPlanePublicIP()
		if got != want {
			t.Errorf(
				"control plane public IP mismatch, want: %s, got: %s",
				want, got,
			)
		}
	}
}

func TestTerraformOutputMachines(t *testing.T) {
	testCases := map[api.ClusterType][]api.MachineState{
		api.ServiceCluster: {{
			Machine: api.Machine{
				NodeType: api.Master,
				Name:     "master-0",
			},
			PublicIP:  "54.183.132.152",
			PrivateIP: "172.16.1.23",
		}, {
			Machine: api.Machine{
				NodeType: api.Worker,
				Name:     "worker-0",
			},
			PublicIP:  "54.241.200.247",
			PrivateIP: "172.16.1.199",
		}, {
			Machine: api.Machine{
				NodeType: api.Worker,
				Name:     "worker-1",
			},
			PublicIP:  "54.241.61.156",
			PrivateIP: "172.16.1.24",
		}},
		api.WorkloadCluster: {{
			Machine: api.Machine{
				NodeType: api.Master,
				Name:     "master-0",
			},
			PublicIP:  "54.241.106.114",
			PrivateIP: "172.16.1.128",
		}, {
			Machine: api.Machine{
				NodeType: api.Worker,
				Name:     "worker-0",
			},
			PublicIP:  "54.153.41.224",
			PrivateIP: "172.16.1.12",
		}, {
			Machine: api.Machine{
				NodeType: api.Worker,
				Name:     "worker-1",
			},
			PublicIP:  "18.144.147.194",
			PrivateIP: "172.16.1.234",
		}},
	}

	for clusterType, wantMachines := range testCases {
		tfOutput := testState(t, clusterType, "ck8stest")

		gotMachines := tfOutput.Machines()

		if diff := cmp.Diff(wantMachines, gotMachines); diff != "" {
			t.Errorf("machines mismatch (-want +got):\n%s", diff)
		}

		for _, wantMachine := range wantMachines {
			gotMachine, err := tfOutput.Machine(
				wantMachine.NodeType,
				wantMachine.Name,
			)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(wantMachine, gotMachine); diff != "" {
				t.Errorf("machine mismatch (-want +got):\n%s", diff)
			}
		}
	}
}

func TestTerraformOutputMachinesNotFound(t *testing.T) {
	tfOutput := testState(t, api.ServiceCluster, "ck8stest")
	_, err := tfOutput.Machine(api.Master, "test")
	notFoundErr := &api.MachineStateNotFoundError{}
	if !errors.As(err, &notFoundErr) {
		t.Error("expected MachineStateNotFoundError")
	}
}
