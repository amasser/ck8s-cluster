package openstack

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/elastisys/ck8s/api"
)

func TestCloneMachine(t *testing.T) {
	testName := "foo"

	type tfvarsPart struct {
		NameSlice []string
		SizeMap   map[string]string
	}

	cluster := Default(-1, "", "testName")

	cluster.TFVars.MachinesSC = map[string]api.Machine{
		testName: {
			NodeType: api.Master,
			Size:     "a1093fde-0772-474b-aced-42a5a2d36814",
		},
	}

	cluster.TFVars.MachinesWC = map[string]api.Machine{
		testName: {
			NodeType: api.Worker,
			Size:     "3232fa6c-3af1-4608-b0f9-acce2415a7a8",
		},
	}

	for _, clusterType := range []api.ClusterType{
		api.ServiceCluster,
		api.WorkloadCluster,
	} {
		cluster.Config.ClusterType = clusterType

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

	got := Default(-1, api.Safespring, "testName")
	want := Default(-1, api.CityCloud, "testName")

	got.TFVars = TFVars{
		MachinesSC: map[string]api.Machine{
			testName: {
				NodeType: api.Master,
				Size:     "a1093fde-0772-474b-aced-42a5a2d36814",
			},
		},
		MachinesWC: map[string]api.Machine{
			testName: {
				NodeType: api.Worker,
				Size:     "3232fa6c-3af1-4608-b0f9-acce2415a7a8",
			},
		},
	}

	want.TFVars = TFVars{
		MachinesSC: map[string]api.Machine{},
		MachinesWC: map[string]api.Machine{},
	}

	for _, clusterType := range []api.ClusterType{
		api.ServiceCluster,
		api.WorkloadCluster,
	} {
		got.Config.ClusterType = clusterType

		if err := got.RemoveMachine(testName); err != nil {
			t.Fatalf(
				"error while removing %s machine: %s",
				clusterType.String(), err,
			)
		}
	}

	if diff := cmp.Diff(want.TFVars, got.TFVars); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
