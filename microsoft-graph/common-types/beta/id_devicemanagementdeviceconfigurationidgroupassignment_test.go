package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationIdGroupAssignmentId{}

func TestNewDeviceManagementDeviceConfigurationIdGroupAssignmentID(t *testing.T) {
	id := NewDeviceManagementDeviceConfigurationIdGroupAssignmentID("deviceConfigurationId", "deviceConfigurationGroupAssignmentId")

	if id.DeviceConfigurationId != "deviceConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationId'", id.DeviceConfigurationId, "deviceConfigurationId")
	}

	if id.DeviceConfigurationGroupAssignmentId != "deviceConfigurationGroupAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationGroupAssignmentId'", id.DeviceConfigurationGroupAssignmentId, "deviceConfigurationGroupAssignmentId")
	}
}

func TestFormatDeviceManagementDeviceConfigurationIdGroupAssignmentID(t *testing.T) {
	actual := NewDeviceManagementDeviceConfigurationIdGroupAssignmentID("deviceConfigurationId", "deviceConfigurationGroupAssignmentId").ID()
	expected := "/deviceManagement/deviceConfigurations/deviceConfigurationId/groupAssignments/deviceConfigurationGroupAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceConfigurationIdGroupAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationIdGroupAssignmentId
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
			Input: "/deviceManagement/deviceConfigurations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/groupAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/groupAssignments/deviceConfigurationGroupAssignmentId",
			Expected: &DeviceManagementDeviceConfigurationIdGroupAssignmentId{
				DeviceConfigurationId:                "deviceConfigurationId",
				DeviceConfigurationGroupAssignmentId: "deviceConfigurationGroupAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/groupAssignments/deviceConfigurationGroupAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationIdGroupAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceConfigurationId != v.Expected.DeviceConfigurationId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationId", v.Expected.DeviceConfigurationId, actual.DeviceConfigurationId)
		}

		if actual.DeviceConfigurationGroupAssignmentId != v.Expected.DeviceConfigurationGroupAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationGroupAssignmentId", v.Expected.DeviceConfigurationGroupAssignmentId, actual.DeviceConfigurationGroupAssignmentId)
		}

	}
}

func TestParseDeviceManagementDeviceConfigurationIdGroupAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationIdGroupAssignmentId
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
			Input: "/deviceManagement/deviceConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/groupAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnId/gRoUpAsSiGnMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/groupAssignments/deviceConfigurationGroupAssignmentId",
			Expected: &DeviceManagementDeviceConfigurationIdGroupAssignmentId{
				DeviceConfigurationId:                "deviceConfigurationId",
				DeviceConfigurationGroupAssignmentId: "deviceConfigurationGroupAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/groupAssignments/deviceConfigurationGroupAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnId/gRoUpAsSiGnMeNtS/dEvIcEcOnFiGuRaTiOnGrOuPaSsIgNmEnTiD",
			Expected: &DeviceManagementDeviceConfigurationIdGroupAssignmentId{
				DeviceConfigurationId:                "dEvIcEcOnFiGuRaTiOnId",
				DeviceConfigurationGroupAssignmentId: "dEvIcEcOnFiGuRaTiOnGrOuPaSsIgNmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnId/gRoUpAsSiGnMeNtS/dEvIcEcOnFiGuRaTiOnGrOuPaSsIgNmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationIdGroupAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceConfigurationId != v.Expected.DeviceConfigurationId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationId", v.Expected.DeviceConfigurationId, actual.DeviceConfigurationId)
		}

		if actual.DeviceConfigurationGroupAssignmentId != v.Expected.DeviceConfigurationGroupAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationGroupAssignmentId", v.Expected.DeviceConfigurationGroupAssignmentId, actual.DeviceConfigurationGroupAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceConfigurationIdGroupAssignmentId(t *testing.T) {
	segments := DeviceManagementDeviceConfigurationIdGroupAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceConfigurationIdGroupAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
