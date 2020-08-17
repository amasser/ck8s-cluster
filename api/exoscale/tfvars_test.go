package exoscale

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/elastisys/ck8s/api"
)

func TestCloneMachine(t *testing.T) {
	testName := "foo"

	type tfvarsPart struct {
		nameSlice []string
		sizeMap   map[string]string
		esCapMap  map[string]int
	}

	cluster := Default(-1, "testName")

	cluster.tfvars.MachinesSC = map[string]ExoscaleMachine{
		testName: {
			Machine: api.Machine{
				NodeType: api.Master,
				Size:     "Small",
			},
			ESLocalStorageCapacity: 10,
		},
	}
	cluster.tfvars.MachinesWC = map[string]ExoscaleMachine{
		testName: {
			Machine: api.Machine{
				NodeType: api.Worker,
				Size:     "Large",
			},
		},
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

	got.tfvars = ExoscaleTFVars{
		MachinesSC: map[string]ExoscaleMachine{
			testName: {
				Machine: api.Machine{
					NodeType: api.Master,
					Size:     "Small",
				},
				ESLocalStorageCapacity: 10,
			},
		},
		MachinesWC: map[string]ExoscaleMachine{
			testName: {
				Machine: api.Machine{
					NodeType: api.Worker,
					Size:     "Large",
				},
			},
		},
	}

	want.tfvars = ExoscaleTFVars{
		MachinesSC: map[string]ExoscaleMachine{},
		MachinesWC: map[string]ExoscaleMachine{},
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
