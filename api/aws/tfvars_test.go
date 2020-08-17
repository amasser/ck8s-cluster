package aws

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/elastisys/ck8s/api"
)

func TestCloneMachine(t *testing.T) {
	testName := "foo"

	type tfvarsPart struct {
		nameSizeMap map[string]string
	}

	cluster := Default(-1, "testName")

	cluster.tfvars.MachinesSC = map[string]api.Machine{
		testName: {
			NodeType: api.Master,
			Size:     "t3.small",
		},
	}

	cluster.tfvars.MachinesWC = map[string]api.Machine{
		testName: {
			NodeType: api.Worker,
			Size:     "t3.large",
		},
	}

	for _, clusterType := range []api.ClusterType{
		api.ServiceCluster,
		api.WorkloadCluster,
	} {
		cluster.config.ClusterType = clusterType

		if _, err := cluster.CloneMachine(testName); err != nil {
			t.Fatalf(
				"error while cloning %s machine: %s",
				clusterType.String(), err,
			)
		}
	}

	for _, clusterType := range []api.ClusterType{
		api.ServiceCluster,
		api.WorkloadCluster,
	} {
		cluster.config.ClusterType = clusterType

		cloneName, err := cluster.CloneMachine(testName)
		if err != nil {
			t.Fatalf(
				"error while cloning %s machine: %s",
				clusterType.String(), err,
			)
		}

		machines := cluster.Machines()

		clonedMachine, ok := machines[cloneName]
		if !ok {
			t.Errorf(
				"cloned machine missing: %s", cloneName,
			)
		}

		if diff := cmp.Diff(machines[testName], clonedMachine); diff != "" {
			t.Errorf("clone mismatch (-want +got):\n%s", diff)
		}
	}
}

func TestRemoveMachine(t *testing.T) {
	testName := "bar"

	got, want := Default(-1, "testName"), Default(-1, "testName")

	got.tfvars = AWSTFVars{
		MachinesSC: map[string]api.Machine{
			testName: {
				NodeType: api.Master,
				Size:     "t3.small",
			},
		},
		MachinesWC: map[string]api.Machine{
			testName: {
				NodeType: api.Worker,
				Size:     "t3.large",
			},
		},
	}

	want.tfvars = AWSTFVars{
		MachinesSC: map[string]api.Machine{},
		MachinesWC: map[string]api.Machine{},
	}

	for _, clusterType := range []api.ClusterType{
		api.ServiceCluster,
		api.WorkloadCluster,
	} {
		got.config.ClusterType = clusterType

		if err := got.RemoveMachine(testName); err != nil {
			t.Fatalf(
				"error while removing %s machine: %s",
				clusterType.String(), err,
			)
		}
	}

	if diff := cmp.Diff(want.tfvars, got.tfvars); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
