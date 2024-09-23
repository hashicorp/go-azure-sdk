package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDeviceIdUsageRightId{}

func TestNewMeDeviceIdUsageRightID(t *testing.T) {
	id := NewMeDeviceIdUsageRightID("deviceId", "usageRightId")

	if id.DeviceId != "deviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceId'", id.DeviceId, "deviceId")
	}

	if id.UsageRightId != "usageRightId" {
		t.Fatalf("Expected %q but got %q for Segment 'UsageRightId'", id.UsageRightId, "usageRightId")
	}
}

func TestFormatMeDeviceIdUsageRightID(t *testing.T) {
	actual := NewMeDeviceIdUsageRightID("deviceId", "usageRightId").ID()
	expected := "/me/devices/deviceId/usageRights/usageRightId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDeviceIdUsageRightID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceIdUsageRightId
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
			Input: "/me/devices/deviceId/usageRights",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/devices/deviceId/usageRights/usageRightId",
			Expected: &MeDeviceIdUsageRightId{
				DeviceId:     "deviceId",
				UsageRightId: "usageRightId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/devices/deviceId/usageRights/usageRightId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceIdUsageRightID(v.Input)
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

		if actual.UsageRightId != v.Expected.UsageRightId {
			t.Fatalf("Expected %q but got %q for UsageRightId", v.Expected.UsageRightId, actual.UsageRightId)
		}

	}
}

func TestParseMeDeviceIdUsageRightIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceIdUsageRightId
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
			Input: "/me/devices/deviceId/usageRights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiD/uSaGeRiGhTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/devices/deviceId/usageRights/usageRightId",
			Expected: &MeDeviceIdUsageRightId{
				DeviceId:     "deviceId",
				UsageRightId: "usageRightId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/devices/deviceId/usageRights/usageRightId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiD/uSaGeRiGhTs/uSaGeRiGhTiD",
			Expected: &MeDeviceIdUsageRightId{
				DeviceId:     "dEvIcEiD",
				UsageRightId: "uSaGeRiGhTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiD/uSaGeRiGhTs/uSaGeRiGhTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceIdUsageRightIDInsensitively(v.Input)
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

		if actual.UsageRightId != v.Expected.UsageRightId {
			t.Fatalf("Expected %q but got %q for UsageRightId", v.Expected.UsageRightId, actual.UsageRightId)
		}

	}
}

func TestSegmentsForMeDeviceIdUsageRightId(t *testing.T) {
	segments := MeDeviceIdUsageRightId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDeviceIdUsageRightId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
