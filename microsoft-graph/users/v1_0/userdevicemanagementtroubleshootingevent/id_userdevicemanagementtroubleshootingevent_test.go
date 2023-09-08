package userdevicemanagementtroubleshootingevent

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserDeviceManagementTroubleshootingEventId{}

func TestNewUserDeviceManagementTroubleshootingEventID(t *testing.T) {
	id := NewUserDeviceManagementTroubleshootingEventID("userIdValue", "deviceManagementTroubleshootingEventIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.DeviceManagementTroubleshootingEventId != "deviceManagementTroubleshootingEventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTroubleshootingEventId'", id.DeviceManagementTroubleshootingEventId, "deviceManagementTroubleshootingEventIdValue")
	}
}

func TestFormatUserDeviceManagementTroubleshootingEventID(t *testing.T) {
	actual := NewUserDeviceManagementTroubleshootingEventID("userIdValue", "deviceManagementTroubleshootingEventIdValue").ID()
	expected := "/users/userIdValue/deviceManagementTroubleshootingEvents/deviceManagementTroubleshootingEventIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserDeviceManagementTroubleshootingEventID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserDeviceManagementTroubleshootingEventId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/deviceManagementTroubleshootingEvents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/deviceManagementTroubleshootingEvents/deviceManagementTroubleshootingEventIdValue",
			Expected: &UserDeviceManagementTroubleshootingEventId{
				UserId:                                 "userIdValue",
				DeviceManagementTroubleshootingEventId: "deviceManagementTroubleshootingEventIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/deviceManagementTroubleshootingEvents/deviceManagementTroubleshootingEventIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserDeviceManagementTroubleshootingEventID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.DeviceManagementTroubleshootingEventId != v.Expected.DeviceManagementTroubleshootingEventId {
			t.Fatalf("Expected %q but got %q for DeviceManagementTroubleshootingEventId", v.Expected.DeviceManagementTroubleshootingEventId, actual.DeviceManagementTroubleshootingEventId)
		}

	}
}

func TestParseUserDeviceManagementTroubleshootingEventIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserDeviceManagementTroubleshootingEventId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/deviceManagementTroubleshootingEvents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dEvIcEmAnAgEmEnTtRoUbLeShOoTiNgEvEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/deviceManagementTroubleshootingEvents/deviceManagementTroubleshootingEventIdValue",
			Expected: &UserDeviceManagementTroubleshootingEventId{
				UserId:                                 "userIdValue",
				DeviceManagementTroubleshootingEventId: "deviceManagementTroubleshootingEventIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/deviceManagementTroubleshootingEvents/deviceManagementTroubleshootingEventIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dEvIcEmAnAgEmEnTtRoUbLeShOoTiNgEvEnTs/dEvIcEmAnAgEmEnTtRoUbLeShOoTiNgEvEnTiDvAlUe",
			Expected: &UserDeviceManagementTroubleshootingEventId{
				UserId:                                 "uSeRiDvAlUe",
				DeviceManagementTroubleshootingEventId: "dEvIcEmAnAgEmEnTtRoUbLeShOoTiNgEvEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dEvIcEmAnAgEmEnTtRoUbLeShOoTiNgEvEnTs/dEvIcEmAnAgEmEnTtRoUbLeShOoTiNgEvEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserDeviceManagementTroubleshootingEventIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.DeviceManagementTroubleshootingEventId != v.Expected.DeviceManagementTroubleshootingEventId {
			t.Fatalf("Expected %q but got %q for DeviceManagementTroubleshootingEventId", v.Expected.DeviceManagementTroubleshootingEventId, actual.DeviceManagementTroubleshootingEventId)
		}

	}
}

func TestSegmentsForUserDeviceManagementTroubleshootingEventId(t *testing.T) {
	segments := UserDeviceManagementTroubleshootingEventId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserDeviceManagementTroubleshootingEventId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
