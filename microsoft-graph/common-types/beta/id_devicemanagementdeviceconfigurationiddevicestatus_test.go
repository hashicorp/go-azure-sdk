package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationIdDeviceStatusId{}

func TestNewDeviceManagementDeviceConfigurationIdDeviceStatusID(t *testing.T) {
	id := NewDeviceManagementDeviceConfigurationIdDeviceStatusID("deviceConfigurationId", "deviceConfigurationDeviceStatusId")

	if id.DeviceConfigurationId != "deviceConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationId'", id.DeviceConfigurationId, "deviceConfigurationId")
	}

	if id.DeviceConfigurationDeviceStatusId != "deviceConfigurationDeviceStatusId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationDeviceStatusId'", id.DeviceConfigurationDeviceStatusId, "deviceConfigurationDeviceStatusId")
	}
}

func TestFormatDeviceManagementDeviceConfigurationIdDeviceStatusID(t *testing.T) {
	actual := NewDeviceManagementDeviceConfigurationIdDeviceStatusID("deviceConfigurationId", "deviceConfigurationDeviceStatusId").ID()
	expected := "/deviceManagement/deviceConfigurations/deviceConfigurationId/deviceStatuses/deviceConfigurationDeviceStatusId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceConfigurationIdDeviceStatusID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationIdDeviceStatusId
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
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/deviceStatuses",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/deviceStatuses/deviceConfigurationDeviceStatusId",
			Expected: &DeviceManagementDeviceConfigurationIdDeviceStatusId{
				DeviceConfigurationId:             "deviceConfigurationId",
				DeviceConfigurationDeviceStatusId: "deviceConfigurationDeviceStatusId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/deviceStatuses/deviceConfigurationDeviceStatusId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationIdDeviceStatusID(v.Input)
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

		if actual.DeviceConfigurationDeviceStatusId != v.Expected.DeviceConfigurationDeviceStatusId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationDeviceStatusId", v.Expected.DeviceConfigurationDeviceStatusId, actual.DeviceConfigurationDeviceStatusId)
		}

	}
}

func TestParseDeviceManagementDeviceConfigurationIdDeviceStatusIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationIdDeviceStatusId
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
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/deviceStatuses",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnId/dEvIcEsTaTuSeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/deviceStatuses/deviceConfigurationDeviceStatusId",
			Expected: &DeviceManagementDeviceConfigurationIdDeviceStatusId{
				DeviceConfigurationId:             "deviceConfigurationId",
				DeviceConfigurationDeviceStatusId: "deviceConfigurationDeviceStatusId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationId/deviceStatuses/deviceConfigurationDeviceStatusId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnId/dEvIcEsTaTuSeS/dEvIcEcOnFiGuRaTiOnDeViCeStAtUsId",
			Expected: &DeviceManagementDeviceConfigurationIdDeviceStatusId{
				DeviceConfigurationId:             "dEvIcEcOnFiGuRaTiOnId",
				DeviceConfigurationDeviceStatusId: "dEvIcEcOnFiGuRaTiOnDeViCeStAtUsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnId/dEvIcEsTaTuSeS/dEvIcEcOnFiGuRaTiOnDeViCeStAtUsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationIdDeviceStatusIDInsensitively(v.Input)
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

		if actual.DeviceConfigurationDeviceStatusId != v.Expected.DeviceConfigurationDeviceStatusId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationDeviceStatusId", v.Expected.DeviceConfigurationDeviceStatusId, actual.DeviceConfigurationDeviceStatusId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceConfigurationIdDeviceStatusId(t *testing.T) {
	segments := DeviceManagementDeviceConfigurationIdDeviceStatusId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceConfigurationIdDeviceStatusId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
