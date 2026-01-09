package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDeviceIdCommandId{}

func TestNewUserIdDeviceIdCommandID(t *testing.T) {
	id := NewUserIdDeviceIdCommandID("userId", "deviceId", "commandId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DeviceId != "deviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceId'", id.DeviceId, "deviceId")
	}

	if id.CommandId != "commandId" {
		t.Fatalf("Expected %q but got %q for Segment 'CommandId'", id.CommandId, "commandId")
	}
}

func TestFormatUserIdDeviceIdCommandID(t *testing.T) {
	actual := NewUserIdDeviceIdCommandID("userId", "deviceId", "commandId").ID()
	expected := "/users/userId/devices/deviceId/commands/commandId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDeviceIdCommandID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDeviceIdCommandId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/devices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/devices/deviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/devices/deviceId/commands",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/devices/deviceId/commands/commandId",
			Expected: &UserIdDeviceIdCommandId{
				UserId:    "userId",
				DeviceId:  "deviceId",
				CommandId: "commandId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/devices/deviceId/commands/commandId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDeviceIdCommandID(v.Input)
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

func TestParseUserIdDeviceIdCommandIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDeviceIdCommandId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/devices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dEvIcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/devices/deviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dEvIcEs/dEvIcEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/devices/deviceId/commands",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dEvIcEs/dEvIcEiD/cOmMaNdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/devices/deviceId/commands/commandId",
			Expected: &UserIdDeviceIdCommandId{
				UserId:    "userId",
				DeviceId:  "deviceId",
				CommandId: "commandId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/devices/deviceId/commands/commandId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dEvIcEs/dEvIcEiD/cOmMaNdS/cOmMaNdId",
			Expected: &UserIdDeviceIdCommandId{
				UserId:    "uSeRiD",
				DeviceId:  "dEvIcEiD",
				CommandId: "cOmMaNdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dEvIcEs/dEvIcEiD/cOmMaNdS/cOmMaNdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDeviceIdCommandIDInsensitively(v.Input)
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

func TestSegmentsForUserIdDeviceIdCommandId(t *testing.T) {
	segments := UserIdDeviceIdCommandId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDeviceIdCommandId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
