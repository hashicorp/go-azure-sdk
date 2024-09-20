package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAutopilotEventId{}

func TestNewDeviceManagementAutopilotEventID(t *testing.T) {
	id := NewDeviceManagementAutopilotEventID("deviceManagementAutopilotEventId")

	if id.DeviceManagementAutopilotEventId != "deviceManagementAutopilotEventId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementAutopilotEventId'", id.DeviceManagementAutopilotEventId, "deviceManagementAutopilotEventId")
	}
}

func TestFormatDeviceManagementAutopilotEventID(t *testing.T) {
	actual := NewDeviceManagementAutopilotEventID("deviceManagementAutopilotEventId").ID()
	expected := "/deviceManagement/autopilotEvents/deviceManagementAutopilotEventId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementAutopilotEventID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAutopilotEventId
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
			Input: "/deviceManagement/autopilotEvents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/autopilotEvents/deviceManagementAutopilotEventId",
			Expected: &DeviceManagementAutopilotEventId{
				DeviceManagementAutopilotEventId: "deviceManagementAutopilotEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/autopilotEvents/deviceManagementAutopilotEventId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAutopilotEventID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementAutopilotEventId != v.Expected.DeviceManagementAutopilotEventId {
			t.Fatalf("Expected %q but got %q for DeviceManagementAutopilotEventId", v.Expected.DeviceManagementAutopilotEventId, actual.DeviceManagementAutopilotEventId)
		}

	}
}

func TestParseDeviceManagementAutopilotEventIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAutopilotEventId
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
			Input: "/deviceManagement/autopilotEvents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aUtOpIlOtEvEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/autopilotEvents/deviceManagementAutopilotEventId",
			Expected: &DeviceManagementAutopilotEventId{
				DeviceManagementAutopilotEventId: "deviceManagementAutopilotEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/autopilotEvents/deviceManagementAutopilotEventId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aUtOpIlOtEvEnTs/dEvIcEmAnAgEmEnTaUtOpIlOtEvEnTiD",
			Expected: &DeviceManagementAutopilotEventId{
				DeviceManagementAutopilotEventId: "dEvIcEmAnAgEmEnTaUtOpIlOtEvEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aUtOpIlOtEvEnTs/dEvIcEmAnAgEmEnTaUtOpIlOtEvEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAutopilotEventIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementAutopilotEventId != v.Expected.DeviceManagementAutopilotEventId {
			t.Fatalf("Expected %q but got %q for DeviceManagementAutopilotEventId", v.Expected.DeviceManagementAutopilotEventId, actual.DeviceManagementAutopilotEventId)
		}

	}
}

func TestSegmentsForDeviceManagementAutopilotEventId(t *testing.T) {
	segments := DeviceManagementAutopilotEventId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementAutopilotEventId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
