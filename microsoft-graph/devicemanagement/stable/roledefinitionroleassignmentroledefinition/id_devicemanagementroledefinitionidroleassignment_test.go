package roledefinitionroleassignmentroledefinition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRoleDefinitionIdRoleAssignmentId{}

func TestNewDeviceManagementRoleDefinitionIdRoleAssignmentID(t *testing.T) {
	id := NewDeviceManagementRoleDefinitionIdRoleAssignmentID("roleDefinitionIdValue", "roleAssignmentIdValue")

	if id.RoleDefinitionId != "roleDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleDefinitionId'", id.RoleDefinitionId, "roleDefinitionIdValue")
	}

	if id.RoleAssignmentId != "roleAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleAssignmentId'", id.RoleAssignmentId, "roleAssignmentIdValue")
	}
}

func TestFormatDeviceManagementRoleDefinitionIdRoleAssignmentID(t *testing.T) {
	actual := NewDeviceManagementRoleDefinitionIdRoleAssignmentID("roleDefinitionIdValue", "roleAssignmentIdValue").ID()
	expected := "/deviceManagement/roleDefinitions/roleDefinitionIdValue/roleAssignments/roleAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementRoleDefinitionIdRoleAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementRoleDefinitionIdRoleAssignmentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleDefinitions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleDefinitions/roleDefinitionIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleDefinitions/roleDefinitionIdValue/roleAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/roleDefinitions/roleDefinitionIdValue/roleAssignments/roleAssignmentIdValue",
			Expected: &DeviceManagementRoleDefinitionIdRoleAssignmentId{
				RoleDefinitionId: "roleDefinitionIdValue",
				RoleAssignmentId: "roleAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleDefinitions/roleDefinitionIdValue/roleAssignments/roleAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementRoleDefinitionIdRoleAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RoleDefinitionId != v.Expected.RoleDefinitionId {
			t.Fatalf("Expected %q but got %q for RoleDefinitionId", v.Expected.RoleDefinitionId, actual.RoleDefinitionId)
		}

		if actual.RoleAssignmentId != v.Expected.RoleAssignmentId {
			t.Fatalf("Expected %q but got %q for RoleAssignmentId", v.Expected.RoleAssignmentId, actual.RoleAssignmentId)
		}

	}
}

func TestParseDeviceManagementRoleDefinitionIdRoleAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementRoleDefinitionIdRoleAssignmentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEdEfInItIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleDefinitions/roleDefinitionIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEdEfInItIoNs/rOlEdEfInItIoNiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleDefinitions/roleDefinitionIdValue/roleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEdEfInItIoNs/rOlEdEfInItIoNiDvAlUe/rOlEaSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/roleDefinitions/roleDefinitionIdValue/roleAssignments/roleAssignmentIdValue",
			Expected: &DeviceManagementRoleDefinitionIdRoleAssignmentId{
				RoleDefinitionId: "roleDefinitionIdValue",
				RoleAssignmentId: "roleAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleDefinitions/roleDefinitionIdValue/roleAssignments/roleAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEdEfInItIoNs/rOlEdEfInItIoNiDvAlUe/rOlEaSsIgNmEnTs/rOlEaSsIgNmEnTiDvAlUe",
			Expected: &DeviceManagementRoleDefinitionIdRoleAssignmentId{
				RoleDefinitionId: "rOlEdEfInItIoNiDvAlUe",
				RoleAssignmentId: "rOlEaSsIgNmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEdEfInItIoNs/rOlEdEfInItIoNiDvAlUe/rOlEaSsIgNmEnTs/rOlEaSsIgNmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementRoleDefinitionIdRoleAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RoleDefinitionId != v.Expected.RoleDefinitionId {
			t.Fatalf("Expected %q but got %q for RoleDefinitionId", v.Expected.RoleDefinitionId, actual.RoleDefinitionId)
		}

		if actual.RoleAssignmentId != v.Expected.RoleAssignmentId {
			t.Fatalf("Expected %q but got %q for RoleAssignmentId", v.Expected.RoleAssignmentId, actual.RoleAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementRoleDefinitionIdRoleAssignmentId(t *testing.T) {
	segments := DeviceManagementRoleDefinitionIdRoleAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementRoleDefinitionIdRoleAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
