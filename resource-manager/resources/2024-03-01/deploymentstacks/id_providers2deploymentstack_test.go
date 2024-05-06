package deploymentstacks

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &Providers2DeploymentStackId{}

func TestNewProviders2DeploymentStackID(t *testing.T) {
	id := NewProviders2DeploymentStackID("managementGroupIdValue", "deploymentStackValue")

	if id.ManagementGroupId != "managementGroupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagementGroupId'", id.ManagementGroupId, "managementGroupIdValue")
	}

	if id.DeploymentStackName != "deploymentStackValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeploymentStackName'", id.DeploymentStackName, "deploymentStackValue")
	}
}

func TestFormatProviders2DeploymentStackID(t *testing.T) {
	actual := NewProviders2DeploymentStackID("managementGroupIdValue", "deploymentStackValue").ID()
	expected := "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Resources/deploymentStacks/deploymentStackValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviders2DeploymentStackID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2DeploymentStackId
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
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Resources",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Resources/deploymentStacks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Resources/deploymentStacks/deploymentStackValue",
			Expected: &Providers2DeploymentStackId{
				ManagementGroupId:   "managementGroupIdValue",
				DeploymentStackName: "deploymentStackValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Resources/deploymentStacks/deploymentStackValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2DeploymentStackID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagementGroupId != v.Expected.ManagementGroupId {
			t.Fatalf("Expected %q but got %q for ManagementGroupId", v.Expected.ManagementGroupId, actual.ManagementGroupId)
		}

		if actual.DeploymentStackName != v.Expected.DeploymentStackName {
			t.Fatalf("Expected %q but got %q for DeploymentStackName", v.Expected.DeploymentStackName, actual.DeploymentStackName)
		}

	}
}

func TestParseProviders2DeploymentStackIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2DeploymentStackId
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
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Resources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.rEsOuRcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Resources/deploymentStacks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/dEpLoYmEnTsTaCkS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Resources/deploymentStacks/deploymentStackValue",
			Expected: &Providers2DeploymentStackId{
				ManagementGroupId:   "managementGroupIdValue",
				DeploymentStackName: "deploymentStackValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Resources/deploymentStacks/deploymentStackValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/dEpLoYmEnTsTaCkS/dEpLoYmEnTsTaCkVaLuE",
			Expected: &Providers2DeploymentStackId{
				ManagementGroupId:   "mAnAgEmEnTgRoUpIdVaLuE",
				DeploymentStackName: "dEpLoYmEnTsTaCkVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/dEpLoYmEnTsTaCkS/dEpLoYmEnTsTaCkVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2DeploymentStackIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagementGroupId != v.Expected.ManagementGroupId {
			t.Fatalf("Expected %q but got %q for ManagementGroupId", v.Expected.ManagementGroupId, actual.ManagementGroupId)
		}

		if actual.DeploymentStackName != v.Expected.DeploymentStackName {
			t.Fatalf("Expected %q but got %q for DeploymentStackName", v.Expected.DeploymentStackName, actual.DeploymentStackName)
		}

	}
}

func TestSegmentsForProviders2DeploymentStackId(t *testing.T) {
	segments := Providers2DeploymentStackId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("Providers2DeploymentStackId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
