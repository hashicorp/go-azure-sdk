package deployments

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = Providers2DeploymentId{}

func TestNewProviders2DeploymentID(t *testing.T) {
	id := NewProviders2DeploymentID("groupIdValue", "deploymentValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.DeploymentName != "deploymentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeploymentName'", id.DeploymentName, "deploymentValue")
	}
}

func TestFormatProviders2DeploymentID(t *testing.T) {
	actual := NewProviders2DeploymentID("groupIdValue", "deploymentValue").ID()
	expected := "/providers/Microsoft.Management/managementGroups/groupIdValue/providers/Microsoft.Resources/deployments/deploymentValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviders2DeploymentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2DeploymentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue/providers/Microsoft.Resources",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue/providers/Microsoft.Resources/deployments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue/providers/Microsoft.Resources/deployments/deploymentValue",
			Expected: &Providers2DeploymentId{
				GroupId:        "groupIdValue",
				DeploymentName: "deploymentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue/providers/Microsoft.Resources/deployments/deploymentValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2DeploymentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.DeploymentName != v.Expected.DeploymentName {
			t.Fatalf("Expected %q but got %q for DeploymentName", v.Expected.DeploymentName, actual.DeploymentName)
		}

	}
}

func TestParseProviders2DeploymentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2DeploymentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/gRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/gRoUpIdVaLuE/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue/providers/Microsoft.Resources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/gRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.rEsOuRcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue/providers/Microsoft.Resources/deployments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/gRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/dEpLoYmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue/providers/Microsoft.Resources/deployments/deploymentValue",
			Expected: &Providers2DeploymentId{
				GroupId:        "groupIdValue",
				DeploymentName: "deploymentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/groupIdValue/providers/Microsoft.Resources/deployments/deploymentValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/gRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/dEpLoYmEnTs/dEpLoYmEnTvAlUe",
			Expected: &Providers2DeploymentId{
				GroupId:        "gRoUpIdVaLuE",
				DeploymentName: "dEpLoYmEnTvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/gRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/dEpLoYmEnTs/dEpLoYmEnTvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2DeploymentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.DeploymentName != v.Expected.DeploymentName {
			t.Fatalf("Expected %q but got %q for DeploymentName", v.Expected.DeploymentName, actual.DeploymentName)
		}

	}
}

func TestSegmentsForProviders2DeploymentId(t *testing.T) {
	segments := Providers2DeploymentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("Providers2DeploymentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
