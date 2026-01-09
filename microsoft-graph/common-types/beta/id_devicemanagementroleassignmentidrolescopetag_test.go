package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRoleAssignmentIdRoleScopeTagId{}

func TestNewDeviceManagementRoleAssignmentIdRoleScopeTagID(t *testing.T) {
	id := NewDeviceManagementRoleAssignmentIdRoleScopeTagID("deviceAndAppManagementRoleAssignmentId", "roleScopeTagId")

	if id.DeviceAndAppManagementRoleAssignmentId != "deviceAndAppManagementRoleAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceAndAppManagementRoleAssignmentId'", id.DeviceAndAppManagementRoleAssignmentId, "deviceAndAppManagementRoleAssignmentId")
	}

	if id.RoleScopeTagId != "roleScopeTagId" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleScopeTagId'", id.RoleScopeTagId, "roleScopeTagId")
	}
}

func TestFormatDeviceManagementRoleAssignmentIdRoleScopeTagID(t *testing.T) {
	actual := NewDeviceManagementRoleAssignmentIdRoleScopeTagID("deviceAndAppManagementRoleAssignmentId", "roleScopeTagId").ID()
	expected := "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentId/roleScopeTags/roleScopeTagId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementRoleAssignmentIdRoleScopeTagID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementRoleAssignmentIdRoleScopeTagId
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
			Input: "/deviceManagement/roleAssignments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentId/roleScopeTags",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentId/roleScopeTags/roleScopeTagId",
			Expected: &DeviceManagementRoleAssignmentIdRoleScopeTagId{
				DeviceAndAppManagementRoleAssignmentId: "deviceAndAppManagementRoleAssignmentId",
				RoleScopeTagId:                         "roleScopeTagId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentId/roleScopeTags/roleScopeTagId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementRoleAssignmentIdRoleScopeTagID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceAndAppManagementRoleAssignmentId != v.Expected.DeviceAndAppManagementRoleAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceAndAppManagementRoleAssignmentId", v.Expected.DeviceAndAppManagementRoleAssignmentId, actual.DeviceAndAppManagementRoleAssignmentId)
		}

		if actual.RoleScopeTagId != v.Expected.RoleScopeTagId {
			t.Fatalf("Expected %q but got %q for RoleScopeTagId", v.Expected.RoleScopeTagId, actual.RoleScopeTagId)
		}

	}
}

func TestParseDeviceManagementRoleAssignmentIdRoleScopeTagIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementRoleAssignmentIdRoleScopeTagId
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
			Input: "/deviceManagement/roleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/dEvIcEaNdApPmAnAgEmEnTrOlEaSsIgNmEnTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentId/roleScopeTags",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/dEvIcEaNdApPmAnAgEmEnTrOlEaSsIgNmEnTiD/rOlEsCoPeTaGs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentId/roleScopeTags/roleScopeTagId",
			Expected: &DeviceManagementRoleAssignmentIdRoleScopeTagId{
				DeviceAndAppManagementRoleAssignmentId: "deviceAndAppManagementRoleAssignmentId",
				RoleScopeTagId:                         "roleScopeTagId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentId/roleScopeTags/roleScopeTagId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/dEvIcEaNdApPmAnAgEmEnTrOlEaSsIgNmEnTiD/rOlEsCoPeTaGs/rOlEsCoPeTaGiD",
			Expected: &DeviceManagementRoleAssignmentIdRoleScopeTagId{
				DeviceAndAppManagementRoleAssignmentId: "dEvIcEaNdApPmAnAgEmEnTrOlEaSsIgNmEnTiD",
				RoleScopeTagId:                         "rOlEsCoPeTaGiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/dEvIcEaNdApPmAnAgEmEnTrOlEaSsIgNmEnTiD/rOlEsCoPeTaGs/rOlEsCoPeTaGiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementRoleAssignmentIdRoleScopeTagIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceAndAppManagementRoleAssignmentId != v.Expected.DeviceAndAppManagementRoleAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceAndAppManagementRoleAssignmentId", v.Expected.DeviceAndAppManagementRoleAssignmentId, actual.DeviceAndAppManagementRoleAssignmentId)
		}

		if actual.RoleScopeTagId != v.Expected.RoleScopeTagId {
			t.Fatalf("Expected %q but got %q for RoleScopeTagId", v.Expected.RoleScopeTagId, actual.RoleScopeTagId)
		}

	}
}

func TestSegmentsForDeviceManagementRoleAssignmentIdRoleScopeTagId(t *testing.T) {
	segments := DeviceManagementRoleAssignmentIdRoleScopeTagId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementRoleAssignmentIdRoleScopeTagId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
