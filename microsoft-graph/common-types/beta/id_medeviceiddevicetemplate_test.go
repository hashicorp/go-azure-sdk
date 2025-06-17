package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceIdDeviceTemplateId{}

func TestNewMeDeviceIdDeviceTemplateID(t *testing.T) {
	id := NewMeDeviceIdDeviceTemplateID("deviceId", "deviceTemplateId")

	if id.DeviceId != "deviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceId'", id.DeviceId, "deviceId")
	}

	if id.DeviceTemplateId != "deviceTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceTemplateId'", id.DeviceTemplateId, "deviceTemplateId")
	}
}

func TestFormatMeDeviceIdDeviceTemplateID(t *testing.T) {
	actual := NewMeDeviceIdDeviceTemplateID("deviceId", "deviceTemplateId").ID()
	expected := "/me/devices/deviceId/deviceTemplate/deviceTemplateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDeviceIdDeviceTemplateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceIdDeviceTemplateId
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
			Input: "/me/devices/deviceId/deviceTemplate",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/devices/deviceId/deviceTemplate/deviceTemplateId",
			Expected: &MeDeviceIdDeviceTemplateId{
				DeviceId:         "deviceId",
				DeviceTemplateId: "deviceTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/devices/deviceId/deviceTemplate/deviceTemplateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceIdDeviceTemplateID(v.Input)
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

		if actual.DeviceTemplateId != v.Expected.DeviceTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceTemplateId", v.Expected.DeviceTemplateId, actual.DeviceTemplateId)
		}

	}
}

func TestParseMeDeviceIdDeviceTemplateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceIdDeviceTemplateId
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
			Input: "/me/devices/deviceId/deviceTemplate",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiD/dEvIcEtEmPlAtE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/devices/deviceId/deviceTemplate/deviceTemplateId",
			Expected: &MeDeviceIdDeviceTemplateId{
				DeviceId:         "deviceId",
				DeviceTemplateId: "deviceTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/devices/deviceId/deviceTemplate/deviceTemplateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiD/dEvIcEtEmPlAtE/dEvIcEtEmPlAtEiD",
			Expected: &MeDeviceIdDeviceTemplateId{
				DeviceId:         "dEvIcEiD",
				DeviceTemplateId: "dEvIcEtEmPlAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiD/dEvIcEtEmPlAtE/dEvIcEtEmPlAtEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceIdDeviceTemplateIDInsensitively(v.Input)
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

		if actual.DeviceTemplateId != v.Expected.DeviceTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceTemplateId", v.Expected.DeviceTemplateId, actual.DeviceTemplateId)
		}

	}
}

func TestSegmentsForMeDeviceIdDeviceTemplateId(t *testing.T) {
	segments := MeDeviceIdDeviceTemplateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDeviceIdDeviceTemplateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
