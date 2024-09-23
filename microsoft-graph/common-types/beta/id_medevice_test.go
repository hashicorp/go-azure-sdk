package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceId{}

func TestNewMeDeviceID(t *testing.T) {
	id := NewMeDeviceID("deviceId")

	if id.DeviceId != "deviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceId'", id.DeviceId, "deviceId")
	}
}

func TestFormatMeDeviceID(t *testing.T) {
	actual := NewMeDeviceID("deviceId").ID()
	expected := "/me/devices/deviceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDeviceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceId
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
			// Valid URI
			Input: "/me/devices/deviceId",
			Expected: &MeDeviceId{
				DeviceId: "deviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/devices/deviceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceID(v.Input)
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

	}
}

func TestParseMeDeviceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceId
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
			// Valid URI
			Input: "/me/devices/deviceId",
			Expected: &MeDeviceId{
				DeviceId: "deviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/devices/deviceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiD",
			Expected: &MeDeviceId{
				DeviceId: "dEvIcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeDeviceId(t *testing.T) {
	segments := MeDeviceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDeviceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
