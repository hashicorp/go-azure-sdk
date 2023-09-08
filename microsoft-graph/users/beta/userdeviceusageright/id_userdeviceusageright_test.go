package userdeviceusageright

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserDeviceUsageRightId{}

func TestNewUserDeviceUsageRightID(t *testing.T) {
	id := NewUserDeviceUsageRightID("userIdValue", "deviceIdValue", "usageRightIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.DeviceId != "deviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceId'", id.DeviceId, "deviceIdValue")
	}

	if id.UsageRightId != "usageRightIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UsageRightId'", id.UsageRightId, "usageRightIdValue")
	}
}

func TestFormatUserDeviceUsageRightID(t *testing.T) {
	actual := NewUserDeviceUsageRightID("userIdValue", "deviceIdValue", "usageRightIdValue").ID()
	expected := "/users/userIdValue/devices/deviceIdValue/usageRights/usageRightIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserDeviceUsageRightID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserDeviceUsageRightId
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
			Input: "/users/userIdValue/devices/deviceIdValue/usageRights",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/devices/deviceIdValue/usageRights/usageRightIdValue",
			Expected: &UserDeviceUsageRightId{
				UserId:       "userIdValue",
				DeviceId:     "deviceIdValue",
				UsageRightId: "usageRightIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/devices/deviceIdValue/usageRights/usageRightIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserDeviceUsageRightID(v.Input)
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

		if actual.UsageRightId != v.Expected.UsageRightId {
			t.Fatalf("Expected %q but got %q for UsageRightId", v.Expected.UsageRightId, actual.UsageRightId)
		}

	}
}

func TestParseUserDeviceUsageRightIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserDeviceUsageRightId
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
			Input: "/users/userIdValue/devices/deviceIdValue/usageRights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dEvIcEs/dEvIcEiDvAlUe/uSaGeRiGhTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/devices/deviceIdValue/usageRights/usageRightIdValue",
			Expected: &UserDeviceUsageRightId{
				UserId:       "userIdValue",
				DeviceId:     "deviceIdValue",
				UsageRightId: "usageRightIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/devices/deviceIdValue/usageRights/usageRightIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dEvIcEs/dEvIcEiDvAlUe/uSaGeRiGhTs/uSaGeRiGhTiDvAlUe",
			Expected: &UserDeviceUsageRightId{
				UserId:       "uSeRiDvAlUe",
				DeviceId:     "dEvIcEiDvAlUe",
				UsageRightId: "uSaGeRiGhTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dEvIcEs/dEvIcEiDvAlUe/uSaGeRiGhTs/uSaGeRiGhTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserDeviceUsageRightIDInsensitively(v.Input)
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

		if actual.UsageRightId != v.Expected.UsageRightId {
			t.Fatalf("Expected %q but got %q for UsageRightId", v.Expected.UsageRightId, actual.UsageRightId)
		}

	}
}

func TestSegmentsForUserDeviceUsageRightId(t *testing.T) {
	segments := UserDeviceUsageRightId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserDeviceUsageRightId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
