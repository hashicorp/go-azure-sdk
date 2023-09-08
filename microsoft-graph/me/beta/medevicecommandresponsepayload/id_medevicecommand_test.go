package medevicecommandresponsepayload

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeDeviceCommandId{}

func TestNewMeDeviceCommandID(t *testing.T) {
	id := NewMeDeviceCommandID("deviceIdValue", "commandIdValue")

	if id.DeviceId != "deviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceId'", id.DeviceId, "deviceIdValue")
	}

	if id.CommandId != "commandIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CommandId'", id.CommandId, "commandIdValue")
	}
}

func TestFormatMeDeviceCommandID(t *testing.T) {
	actual := NewMeDeviceCommandID("deviceIdValue", "commandIdValue").ID()
	expected := "/me/devices/deviceIdValue/commands/commandIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDeviceCommandID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceCommandId
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
			Input: "/me/devices/deviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/devices/deviceIdValue/commands",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/devices/deviceIdValue/commands/commandIdValue",
			Expected: &MeDeviceCommandId{
				DeviceId:  "deviceIdValue",
				CommandId: "commandIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/devices/deviceIdValue/commands/commandIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceCommandID(v.Input)
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

func TestParseMeDeviceCommandIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceCommandId
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
			Input: "/me/devices/deviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/devices/deviceIdValue/commands",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiDvAlUe/cOmMaNdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/devices/deviceIdValue/commands/commandIdValue",
			Expected: &MeDeviceCommandId{
				DeviceId:  "deviceIdValue",
				CommandId: "commandIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/devices/deviceIdValue/commands/commandIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiDvAlUe/cOmMaNdS/cOmMaNdIdVaLuE",
			Expected: &MeDeviceCommandId{
				DeviceId:  "dEvIcEiDvAlUe",
				CommandId: "cOmMaNdIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiDvAlUe/cOmMaNdS/cOmMaNdIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceCommandIDInsensitively(v.Input)
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

func TestSegmentsForMeDeviceCommandId(t *testing.T) {
	segments := MeDeviceCommandId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDeviceCommandId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
