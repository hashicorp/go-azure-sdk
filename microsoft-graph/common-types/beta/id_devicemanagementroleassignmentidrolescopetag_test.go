package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRoleAssignmentIdRoleScopeTagId{}

func TestNewDeviceManagementRoleAssignmentIdRoleScopeTagID(t *testing.T) {
	id := NewDeviceManagementRoleAssignmentIdRoleScopeTagID("deviceAndAppManagementRoleAssignmentIdValue", "roleScopeTagIdValue")

	if id.DeviceAndAppManagementRoleAssignmentId != "deviceAndAppManagementRoleAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceAndAppManagementRoleAssignmentId'", id.DeviceAndAppManagementRoleAssignmentId, "deviceAndAppManagementRoleAssignmentIdValue")
	}

	if id.RoleScopeTagId != "roleScopeTagIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleScopeTagId'", id.RoleScopeTagId, "roleScopeTagIdValue")
	}
}

func TestFormatDeviceManagementRoleAssignmentIdRoleScopeTagID(t *testing.T) {
	actual := NewDeviceManagementRoleAssignmentIdRoleScopeTagID("deviceAndAppManagementRoleAssignmentIdValue", "roleScopeTagIdValue").ID()
	expected := "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentIdValue/roleScopeTags/roleScopeTagIdValue"
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
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentIdValue/roleScopeTags",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentIdValue/roleScopeTags/roleScopeTagIdValue",
			Expected: &DeviceManagementRoleAssignmentIdRoleScopeTagId{
				DeviceAndAppManagementRoleAssignmentId: "deviceAndAppManagementRoleAssignmentIdValue",
				RoleScopeTagId:                         "roleScopeTagIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentIdValue/roleScopeTags/roleScopeTagIdValue/extra",
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
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/dEvIcEaNdApPmAnAgEmEnTrOlEaSsIgNmEnTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentIdValue/roleScopeTags",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/dEvIcEaNdApPmAnAgEmEnTrOlEaSsIgNmEnTiDvAlUe/rOlEsCoPeTaGs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentIdValue/roleScopeTags/roleScopeTagIdValue",
			Expected: &DeviceManagementRoleAssignmentIdRoleScopeTagId{
				DeviceAndAppManagementRoleAssignmentId: "deviceAndAppManagementRoleAssignmentIdValue",
				RoleScopeTagId:                         "roleScopeTagIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentIdValue/roleScopeTags/roleScopeTagIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/dEvIcEaNdApPmAnAgEmEnTrOlEaSsIgNmEnTiDvAlUe/rOlEsCoPeTaGs/rOlEsCoPeTaGiDvAlUe",
			Expected: &DeviceManagementRoleAssignmentIdRoleScopeTagId{
				DeviceAndAppManagementRoleAssignmentId: "dEvIcEaNdApPmAnAgEmEnTrOlEaSsIgNmEnTiDvAlUe",
				RoleScopeTagId:                         "rOlEsCoPeTaGiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/dEvIcEaNdApPmAnAgEmEnTrOlEaSsIgNmEnTiDvAlUe/rOlEsCoPeTaGs/rOlEsCoPeTaGiDvAlUe/extra",
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
