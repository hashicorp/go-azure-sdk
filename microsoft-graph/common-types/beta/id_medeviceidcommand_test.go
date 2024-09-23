package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceIdCommandId{}

func TestNewMeDeviceIdCommandID(t *testing.T) {
	id := NewMeDeviceIdCommandID("deviceId", "commandId")

	if id.DeviceId != "deviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceId'", id.DeviceId, "deviceId")
	}

	if id.CommandId != "commandId" {
		t.Fatalf("Expected %q but got %q for Segment 'CommandId'", id.CommandId, "commandId")
	}
}

func TestFormatMeDeviceIdCommandID(t *testing.T) {
	actual := NewMeDeviceIdCommandID("deviceId", "commandId").ID()
	expected := "/me/devices/deviceId/commands/commandId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDeviceIdCommandID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceIdCommandId
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
			Input: "/me/devices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/devices/deviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/devices/deviceId/commands",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/devices/deviceId/commands/commandId",
			Expected: &MeDeviceIdCommandId{
				DeviceId:  "deviceId",
				CommandId: "commandId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/devices/deviceId/commands/commandId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceIdCommandID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceId != v.Expected.DeviceId {
			t.Fatalf("Expected %q but got %q for DeviceId", v.Expected.DeviceId, actual.DeviceId)
		}

		if actual.CommandId != v.Expected.CommandId {
			t.Fatalf("Expected %q but got %q for CommandId", v.Expected.CommandId, actual.CommandId)
		}

	}
}

func TestParseMeDeviceIdCommandIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceIdCommandId
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
			Input: "/me/devices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/devices/deviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/devices/deviceId/commands",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiD/cOmMaNdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/devices/deviceId/commands/commandId",
			Expected: &MeDeviceIdCommandId{
				DeviceId:  "deviceId",
				CommandId: "commandId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/devices/deviceId/commands/commandId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiD/cOmMaNdS/cOmMaNdId",
			Expected: &MeDeviceIdCommandId{
				DeviceId:  "dEvIcEiD",
				CommandId: "cOmMaNdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiD/cOmMaNdS/cOmMaNdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceIdCommandIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceId != v.Expected.DeviceId {
			t.Fatalf("Expected %q but got %q for DeviceId", v.Expected.DeviceId, actual.DeviceId)
		}

		if actual.CommandId != v.Expected.CommandId {
			t.Fatalf("Expected %q but got %q for CommandId", v.Expected.CommandId, actual.CommandId)
		}

	}
}

func TestSegmentsForMeDeviceIdCommandId(t *testing.T) {
	segments := MeDeviceIdCommandId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDeviceIdCommandId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
