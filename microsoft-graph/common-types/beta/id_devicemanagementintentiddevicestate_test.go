package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdDeviceStateId{}

func TestNewDeviceManagementIntentIdDeviceStateID(t *testing.T) {
	id := NewDeviceManagementIntentIdDeviceStateID("deviceManagementIntentIdValue", "deviceManagementIntentDeviceStateIdValue")

	if id.DeviceManagementIntentId != "deviceManagementIntentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentId'", id.DeviceManagementIntentId, "deviceManagementIntentIdValue")
	}

	if id.DeviceManagementIntentDeviceStateId != "deviceManagementIntentDeviceStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentDeviceStateId'", id.DeviceManagementIntentDeviceStateId, "deviceManagementIntentDeviceStateIdValue")
	}
}

func TestFormatDeviceManagementIntentIdDeviceStateID(t *testing.T) {
	actual := NewDeviceManagementIntentIdDeviceStateID("deviceManagementIntentIdValue", "deviceManagementIntentDeviceStateIdValue").ID()
	expected := "/deviceManagement/intents/deviceManagementIntentIdValue/deviceStates/deviceManagementIntentDeviceStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementIntentIdDeviceStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntentIdDeviceStateId
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
			Input: "/deviceManagement/intents",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/deviceStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/deviceStates/deviceManagementIntentDeviceStateIdValue",
			Expected: &DeviceManagementIntentIdDeviceStateId{
				DeviceManagementIntentId:            "deviceManagementIntentIdValue",
				DeviceManagementIntentDeviceStateId: "deviceManagementIntentDeviceStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/deviceStates/deviceManagementIntentDeviceStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntentIdDeviceStateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementIntentId != v.Expected.DeviceManagementIntentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentId", v.Expected.DeviceManagementIntentId, actual.DeviceManagementIntentId)
		}

		if actual.DeviceManagementIntentDeviceStateId != v.Expected.DeviceManagementIntentDeviceStateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentDeviceStateId", v.Expected.DeviceManagementIntentDeviceStateId, actual.DeviceManagementIntentDeviceStateId)
		}

	}
}

func TestParseDeviceManagementIntentIdDeviceStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntentIdDeviceStateId
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
			Input: "/deviceManagement/intents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/deviceStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/dEvIcEsTaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/deviceStates/deviceManagementIntentDeviceStateIdValue",
			Expected: &DeviceManagementIntentIdDeviceStateId{
				DeviceManagementIntentId:            "deviceManagementIntentIdValue",
				DeviceManagementIntentDeviceStateId: "deviceManagementIntentDeviceStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/deviceStates/deviceManagementIntentDeviceStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/dEvIcEsTaTeS/dEvIcEmAnAgEmEnTiNtEnTdEvIcEsTaTeIdVaLuE",
			Expected: &DeviceManagementIntentIdDeviceStateId{
				DeviceManagementIntentId:            "dEvIcEmAnAgEmEnTiNtEnTiDvAlUe",
				DeviceManagementIntentDeviceStateId: "dEvIcEmAnAgEmEnTiNtEnTdEvIcEsTaTeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/dEvIcEsTaTeS/dEvIcEmAnAgEmEnTiNtEnTdEvIcEsTaTeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntentIdDeviceStateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementIntentId != v.Expected.DeviceManagementIntentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentId", v.Expected.DeviceManagementIntentId, actual.DeviceManagementIntentId)
		}

		if actual.DeviceManagementIntentDeviceStateId != v.Expected.DeviceManagementIntentDeviceStateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentDeviceStateId", v.Expected.DeviceManagementIntentDeviceStateId, actual.DeviceManagementIntentDeviceStateId)
		}

	}
}

func TestSegmentsForDeviceManagementIntentIdDeviceStateId(t *testing.T) {
	segments := DeviceManagementIntentIdDeviceStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementIntentIdDeviceStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
