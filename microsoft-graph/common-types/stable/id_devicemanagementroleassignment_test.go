package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRoleAssignmentId{}

func TestNewDeviceManagementRoleAssignmentID(t *testing.T) {
	id := NewDeviceManagementRoleAssignmentID("deviceAndAppManagementRoleAssignmentId")

	if id.DeviceAndAppManagementRoleAssignmentId != "deviceAndAppManagementRoleAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceAndAppManagementRoleAssignmentId'", id.DeviceAndAppManagementRoleAssignmentId, "deviceAndAppManagementRoleAssignmentId")
	}
}

func TestFormatDeviceManagementRoleAssignmentID(t *testing.T) {
	actual := NewDeviceManagementRoleAssignmentID("deviceAndAppManagementRoleAssignmentId").ID()
	expected := "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementRoleAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementRoleAssignmentId
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
			// Valid URI
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentId",
			Expected: &DeviceManagementRoleAssignmentId{
				DeviceAndAppManagementRoleAssignmentId: "deviceAndAppManagementRoleAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementRoleAssignmentID(v.Input)
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

	}
}

func TestParseDeviceManagementRoleAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementRoleAssignmentId
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
			// Valid URI
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentId",
			Expected: &DeviceManagementRoleAssignmentId{
				DeviceAndAppManagementRoleAssignmentId: "deviceAndAppManagementRoleAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleAssignments/deviceAndAppManagementRoleAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/dEvIcEaNdApPmAnAgEmEnTrOlEaSsIgNmEnTiD",
			Expected: &DeviceManagementRoleAssignmentId{
				DeviceAndAppManagementRoleAssignmentId: "dEvIcEaNdApPmAnAgEmEnTrOlEaSsIgNmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/dEvIcEaNdApPmAnAgEmEnTrOlEaSsIgNmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementRoleAssignmentIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementRoleAssignmentId(t *testing.T) {
	segments := DeviceManagementRoleAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementRoleAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
