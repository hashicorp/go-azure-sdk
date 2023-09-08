package medeviceusageright

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeDeviceUsageRightId{}

func TestNewMeDeviceUsageRightID(t *testing.T) {
	id := NewMeDeviceUsageRightID("deviceIdValue", "usageRightIdValue")

	if id.DeviceId != "deviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceId'", id.DeviceId, "deviceIdValue")
	}

	if id.UsageRightId != "usageRightIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UsageRightId'", id.UsageRightId, "usageRightIdValue")
	}
}

func TestFormatMeDeviceUsageRightID(t *testing.T) {
	actual := NewMeDeviceUsageRightID("deviceIdValue", "usageRightIdValue").ID()
	expected := "/me/devices/deviceIdValue/usageRights/usageRightIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDeviceUsageRightID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceUsageRightId
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
			Input: "/me/devices/deviceIdValue/usageRights",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/devices/deviceIdValue/usageRights/usageRightIdValue",
			Expected: &MeDeviceUsageRightId{
				DeviceId:     "deviceIdValue",
				UsageRightId: "usageRightIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/devices/deviceIdValue/usageRights/usageRightIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceUsageRightID(v.Input)
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

func TestParseMeDeviceUsageRightIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDeviceUsageRightId
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
			Input: "/me/devices/deviceIdValue/usageRights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiDvAlUe/uSaGeRiGhTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/devices/deviceIdValue/usageRights/usageRightIdValue",
			Expected: &MeDeviceUsageRightId{
				DeviceId:     "deviceIdValue",
				UsageRightId: "usageRightIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/devices/deviceIdValue/usageRights/usageRightIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiDvAlUe/uSaGeRiGhTs/uSaGeRiGhTiDvAlUe",
			Expected: &MeDeviceUsageRightId{
				DeviceId:     "dEvIcEiDvAlUe",
				UsageRightId: "uSaGeRiGhTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dEvIcEs/dEvIcEiDvAlUe/uSaGeRiGhTs/uSaGeRiGhTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDeviceUsageRightIDInsensitively(v.Input)
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

func TestSegmentsForMeDeviceUsageRightId(t *testing.T) {
	segments := MeDeviceUsageRightId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDeviceUsageRightId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
