package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationProfileId{}

func TestNewDeviceManagementDeviceConfigurationProfileID(t *testing.T) {
	id := NewDeviceManagementDeviceConfigurationProfileID("deviceConfigurationProfileIdValue")

	if id.DeviceConfigurationProfileId != "deviceConfigurationProfileIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationProfileId'", id.DeviceConfigurationProfileId, "deviceConfigurationProfileIdValue")
	}
}

func TestFormatDeviceManagementDeviceConfigurationProfileID(t *testing.T) {
	actual := NewDeviceManagementDeviceConfigurationProfileID("deviceConfigurationProfileIdValue").ID()
	expected := "/deviceManagement/deviceConfigurationProfiles/deviceConfigurationProfileIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceConfigurationProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationProfileId
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
			Input: "/deviceManagement/deviceConfigurationProfiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurationProfiles/deviceConfigurationProfileIdValue",
			Expected: &DeviceManagementDeviceConfigurationProfileId{
				DeviceConfigurationProfileId: "deviceConfigurationProfileIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurationProfiles/deviceConfigurationProfileIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationProfileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceConfigurationProfileId != v.Expected.DeviceConfigurationProfileId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationProfileId", v.Expected.DeviceConfigurationProfileId, actual.DeviceConfigurationProfileId)
		}

	}
}

func TestParseDeviceManagementDeviceConfigurationProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationProfileId
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
			Input: "/deviceManagement/deviceConfigurationProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnPrOfIlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurationProfiles/deviceConfigurationProfileIdValue",
			Expected: &DeviceManagementDeviceConfigurationProfileId{
				DeviceConfigurationProfileId: "deviceConfigurationProfileIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurationProfiles/deviceConfigurationProfileIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnPrOfIlEs/dEvIcEcOnFiGuRaTiOnPrOfIlEiDvAlUe",
			Expected: &DeviceManagementDeviceConfigurationProfileId{
				DeviceConfigurationProfileId: "dEvIcEcOnFiGuRaTiOnPrOfIlEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnPrOfIlEs/dEvIcEcOnFiGuRaTiOnPrOfIlEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationProfileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceConfigurationProfileId != v.Expected.DeviceConfigurationProfileId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationProfileId", v.Expected.DeviceConfigurationProfileId, actual.DeviceConfigurationProfileId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceConfigurationProfileId(t *testing.T) {
	segments := DeviceManagementDeviceConfigurationProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceConfigurationProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
