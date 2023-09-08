package userdevicecommand

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserDeviceCommandId{}

func TestNewUserDeviceCommandID(t *testing.T) {
	id := NewUserDeviceCommandID("userIdValue", "deviceIdValue", "commandIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.DeviceId != "deviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceId'", id.DeviceId, "deviceIdValue")
	}

	if id.CommandId != "commandIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CommandId'", id.CommandId, "commandIdValue")
	}
}

func TestFormatUserDeviceCommandID(t *testing.T) {
	actual := NewUserDeviceCommandID("userIdValue", "deviceIdValue", "commandIdValue").ID()
	expected := "/users/userIdValue/devices/deviceIdValue/commands/commandIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserDeviceCommandID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserDeviceCommandId
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
			Input: "/users/userIdValue/devices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/devices/deviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/devices/deviceIdValue/commands",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/devices/deviceIdValue/commands/commandIdValue",
			Expected: &UserDeviceCommandId{
				UserId:    "userIdValue",
				DeviceId:  "deviceIdValue",
				CommandId: "commandIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/devices/deviceIdValue/commands/commandIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserDeviceCommandID(v.Input)
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

		if actual.DeviceId != v.Expected.DeviceId {
			t.Fatalf("Expected %q but got %q for DeviceId", v.Expected.DeviceId, actual.DeviceId)
		}

		if actual.CommandId != v.Expected.CommandId {
			t.Fatalf("Expected %q but got %q for CommandId", v.Expected.CommandId, actual.CommandId)
		}

	}
}

func TestParseUserDeviceCommandIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserDeviceCommandId
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
			Input: "/users/userIdValue/devices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dEvIcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/devices/deviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dEvIcEs/dEvIcEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/devices/deviceIdValue/commands",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dEvIcEs/dEvIcEiDvAlUe/cOmMaNdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/devices/deviceIdValue/commands/commandIdValue",
			Expected: &UserDeviceCommandId{
				UserId:    "userIdValue",
				DeviceId:  "deviceIdValue",
				CommandId: "commandIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/devices/deviceIdValue/commands/commandIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dEvIcEs/dEvIcEiDvAlUe/cOmMaNdS/cOmMaNdIdVaLuE",
			Expected: &UserDeviceCommandId{
				UserId:    "uSeRiDvAlUe",
				DeviceId:  "dEvIcEiDvAlUe",
				CommandId: "cOmMaNdIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dEvIcEs/dEvIcEiDvAlUe/cOmMaNdS/cOmMaNdIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserDeviceCommandIDInsensitively(v.Input)
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

		if actual.DeviceId != v.Expected.DeviceId {
			t.Fatalf("Expected %q but got %q for DeviceId", v.Expected.DeviceId, actual.DeviceId)
		}

		if actual.CommandId != v.Expected.CommandId {
			t.Fatalf("Expected %q but got %q for CommandId", v.Expected.CommandId, actual.CommandId)
		}

	}
}

func TestSegmentsForUserDeviceCommandId(t *testing.T) {
	segments := UserDeviceCommandId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserDeviceCommandId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
