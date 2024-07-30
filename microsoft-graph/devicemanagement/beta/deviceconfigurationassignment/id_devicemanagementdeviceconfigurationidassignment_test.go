package deviceconfigurationassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationIdAssignmentId{}

func TestNewDeviceManagementDeviceConfigurationIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementDeviceConfigurationIdAssignmentID("deviceConfigurationIdValue", "deviceConfigurationAssignmentIdValue")

	if id.DeviceConfigurationId != "deviceConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationId'", id.DeviceConfigurationId, "deviceConfigurationIdValue")
	}

	if id.DeviceConfigurationAssignmentId != "deviceConfigurationAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationAssignmentId'", id.DeviceConfigurationAssignmentId, "deviceConfigurationAssignmentIdValue")
	}
}

func TestFormatDeviceManagementDeviceConfigurationIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementDeviceConfigurationIdAssignmentID("deviceConfigurationIdValue", "deviceConfigurationAssignmentIdValue").ID()
	expected := "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/assignments/deviceConfigurationAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceConfigurationIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationIdAssignmentId
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
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/assignments/deviceConfigurationAssignmentIdValue",
			Expected: &DeviceManagementDeviceConfigurationIdAssignmentId{
				DeviceConfigurationId:           "deviceConfigurationIdValue",
				DeviceConfigurationAssignmentId: "deviceConfigurationAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/assignments/deviceConfigurationAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationIdAssignmentID(v.Input)
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

		if actual.DeviceConfigurationAssignmentId != v.Expected.DeviceConfigurationAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationAssignmentId", v.Expected.DeviceConfigurationAssignmentId, actual.DeviceConfigurationAssignmentId)
		}

	}
}

func TestParseDeviceManagementDeviceConfigurationIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationIdAssignmentId
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
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnIdVaLuE/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/assignments/deviceConfigurationAssignmentIdValue",
			Expected: &DeviceManagementDeviceConfigurationIdAssignmentId{
				DeviceConfigurationId:           "deviceConfigurationIdValue",
				DeviceConfigurationAssignmentId: "deviceConfigurationAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/assignments/deviceConfigurationAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnIdVaLuE/aSsIgNmEnTs/dEvIcEcOnFiGuRaTiOnAsSiGnMeNtIdVaLuE",
			Expected: &DeviceManagementDeviceConfigurationIdAssignmentId{
				DeviceConfigurationId:           "dEvIcEcOnFiGuRaTiOnIdVaLuE",
				DeviceConfigurationAssignmentId: "dEvIcEcOnFiGuRaTiOnAsSiGnMeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnIdVaLuE/aSsIgNmEnTs/dEvIcEcOnFiGuRaTiOnAsSiGnMeNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationIdAssignmentIDInsensitively(v.Input)
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

		if actual.DeviceConfigurationAssignmentId != v.Expected.DeviceConfigurationAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationAssignmentId", v.Expected.DeviceConfigurationAssignmentId, actual.DeviceConfigurationAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceConfigurationIdAssignmentId(t *testing.T) {
	segments := DeviceManagementDeviceConfigurationIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceConfigurationIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
