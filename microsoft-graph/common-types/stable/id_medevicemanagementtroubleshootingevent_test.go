package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceManagementTroubleshootingEventId{}

func TestNewMeDeviceManagementTroubleshootingEventID(t *testing.T) {
	id := NewMeDeviceManagementTroubleshootingEventID("deviceManagementTroubleshootingEventId")

	if id.DeviceManagementTroubleshootingEventId != "deviceManagementTroubleshootingEventId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTroubleshootingEventId'", id.DeviceManagementTroubleshootingEventId, "deviceManagementTroubleshootingEventId")
	}
}

func TestFormatMeDeviceManagementTroubleshootingEventID(t *testing.T) {
	actual := NewMeDeviceManagementTroubleshootingEventID("deviceManagementTroubleshootingEventId").ID()
	expected := "/me/deviceManagementTroubleshootingEvents/deviceManagementTroubleshootingEventId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDeviceManagementTroubleshootingEventID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceManagementTroubleshootingEventId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/deviceManagementTroubleshootingEvents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/deviceManagementTroubleshootingEvents/deviceManagementTroubleshootingEventId",
			Expected: &MeDeviceManagementTroubleshootingEventId{
				DeviceManagementTroubleshootingEventId: "deviceManagementTroubleshootingEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/deviceManagementTroubleshootingEvents/deviceManagementTroubleshootingEventId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceManagementTroubleshootingEventID(v.Input)
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

func TestParseMeDeviceManagementTroubleshootingEventIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceManagementTroubleshootingEventId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/deviceManagementTroubleshootingEvents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEmAnAgEmEnTtRoUbLeShOoTiNgEvEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/deviceManagementTroubleshootingEvents/deviceManagementTroubleshootingEventId",
			Expected: &MeDeviceManagementTroubleshootingEventId{
				DeviceManagementTroubleshootingEventId: "deviceManagementTroubleshootingEventId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/deviceManagementTroubleshootingEvents/deviceManagementTroubleshootingEventId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEmAnAgEmEnTtRoUbLeShOoTiNgEvEnTs/dEvIcEmAnAgEmEnTtRoUbLeShOoTiNgEvEnTiD",
			Expected: &MeDeviceManagementTroubleshootingEventId{
				DeviceManagementTroubleshootingEventId: "dEvIcEmAnAgEmEnTtRoUbLeShOoTiNgEvEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEmAnAgEmEnTtRoUbLeShOoTiNgEvEnTs/dEvIcEmAnAgEmEnTtRoUbLeShOoTiNgEvEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceManagementTroubleshootingEventIDInsensitively(v.Input)
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

func TestSegmentsForMeDeviceManagementTroubleshootingEventId(t *testing.T) {
	segments := MeDeviceManagementTroubleshootingEventId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDeviceManagementTroubleshootingEventId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
