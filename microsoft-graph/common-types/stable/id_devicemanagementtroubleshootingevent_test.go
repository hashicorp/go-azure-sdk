package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTroubleshootingEventId{}

func TestNewDeviceManagementTroubleshootingEventID(t *testing.T) {
	id := NewDeviceManagementTroubleshootingEventID("deviceManagementTroubleshootingEventId")

	if id.DeviceManagementTroubleshootingEventId != "deviceManagementTroubleshootingEventId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTroubleshootingEventId'", id.DeviceManagementTroubleshootingEventId, "deviceManagementTroubleshootingEventId")
	}
}

func TestFormatDeviceManagementTroubleshootingEventID(t *testing.T) {
	actual := NewDeviceManagementTroubleshootingEventID("deviceManagementTroubleshootingEventId").ID()
	expected := "/deviceManagement/troubleshootingEvents/deviceManagementTroubleshootingEventId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementTroubleshootingEventID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTroubleshootingEventId
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
			Input: "/deviceManagement/troubleshootingEvents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/troubleshootingEvents/deviceManagementTroubleshootingEventId",
			Expected: &DeviceManagementTroubleshootingEventId{
				DeviceManagementTroubleshootingEventId: "deviceManagementTroubleshootingEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/troubleshootingEvents/deviceManagementTroubleshootingEventId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTroubleshootingEventID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementTroubleshootingEventId != v.Expected.DeviceManagementTroubleshootingEventId {
			t.Fatalf("Expected %q but got %q for DeviceManagementTroubleshootingEventId", v.Expected.DeviceManagementTroubleshootingEventId, actual.DeviceManagementTroubleshootingEventId)
		}

	}
}

func TestParseDeviceManagementTroubleshootingEventIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTroubleshootingEventId
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
			Input: "/deviceManagement/troubleshootingEvents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tRoUbLeShOoTiNgEvEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/troubleshootingEvents/deviceManagementTroubleshootingEventId",
			Expected: &DeviceManagementTroubleshootingEventId{
				DeviceManagementTroubleshootingEventId: "deviceManagementTroubleshootingEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/troubleshootingEvents/deviceManagementTroubleshootingEventId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tRoUbLeShOoTiNgEvEnTs/dEvIcEmAnAgEmEnTtRoUbLeShOoTiNgEvEnTiD",
			Expected: &DeviceManagementTroubleshootingEventId{
				DeviceManagementTroubleshootingEventId: "dEvIcEmAnAgEmEnTtRoUbLeShOoTiNgEvEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tRoUbLeShOoTiNgEvEnTs/dEvIcEmAnAgEmEnTtRoUbLeShOoTiNgEvEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTroubleshootingEventIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementTroubleshootingEventId != v.Expected.DeviceManagementTroubleshootingEventId {
			t.Fatalf("Expected %q but got %q for DeviceManagementTroubleshootingEventId", v.Expected.DeviceManagementTroubleshootingEventId, actual.DeviceManagementTroubleshootingEventId)
		}

	}
}

func TestSegmentsForDeviceManagementTroubleshootingEventId(t *testing.T) {
	segments := DeviceManagementTroubleshootingEventId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementTroubleshootingEventId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
