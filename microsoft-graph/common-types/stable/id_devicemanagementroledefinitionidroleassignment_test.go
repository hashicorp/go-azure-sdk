package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRoleDefinitionIdRoleAssignmentId{}

func TestNewDeviceManagementRoleDefinitionIdRoleAssignmentID(t *testing.T) {
	id := NewDeviceManagementRoleDefinitionIdRoleAssignmentID("roleDefinitionId", "roleAssignmentId")

	if id.RoleDefinitionId != "roleDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleDefinitionId'", id.RoleDefinitionId, "roleDefinitionId")
	}

	if id.RoleAssignmentId != "roleAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleAssignmentId'", id.RoleAssignmentId, "roleAssignmentId")
	}
}

func TestFormatDeviceManagementRoleDefinitionIdRoleAssignmentID(t *testing.T) {
	actual := NewDeviceManagementRoleDefinitionIdRoleAssignmentID("roleDefinitionId", "roleAssignmentId").ID()
	expected := "/deviceManagement/roleDefinitions/roleDefinitionId/roleAssignments/roleAssignmentId"
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
			Input: "/deviceManagement/roleDefinitions/roleDefinitionId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleDefinitions/roleDefinitionId/roleAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/roleDefinitions/roleDefinitionId/roleAssignments/roleAssignmentId",
			Expected: &DeviceManagementRoleDefinitionIdRoleAssignmentId{
				RoleDefinitionId: "roleDefinitionId",
				RoleAssignmentId: "roleAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleDefinitions/roleDefinitionId/roleAssignments/roleAssignmentId/extra",
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
			Input: "/deviceManagement/roleDefinitions/roleDefinitionId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEdEfInItIoNs/rOlEdEfInItIoNiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleDefinitions/roleDefinitionId/roleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEdEfInItIoNs/rOlEdEfInItIoNiD/rOlEaSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/roleDefinitions/roleDefinitionId/roleAssignments/roleAssignmentId",
			Expected: &DeviceManagementRoleDefinitionIdRoleAssignmentId{
				RoleDefinitionId: "roleDefinitionId",
				RoleAssignmentId: "roleAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleDefinitions/roleDefinitionId/roleAssignments/roleAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEdEfInItIoNs/rOlEdEfInItIoNiD/rOlEaSsIgNmEnTs/rOlEaSsIgNmEnTiD",
			Expected: &DeviceManagementRoleDefinitionIdRoleAssignmentId{
				RoleDefinitionId: "rOlEdEfInItIoNiD",
				RoleAssignmentId: "rOlEaSsIgNmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEdEfInItIoNs/rOlEdEfInItIoNiD/rOlEaSsIgNmEnTs/rOlEaSsIgNmEnTiD/extra",
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
