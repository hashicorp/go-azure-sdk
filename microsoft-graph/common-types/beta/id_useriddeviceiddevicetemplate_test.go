package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDeviceIdDeviceTemplateId{}

func TestNewUserIdDeviceIdDeviceTemplateID(t *testing.T) {
	id := NewUserIdDeviceIdDeviceTemplateID("userId", "deviceId", "deviceTemplateId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DeviceId != "deviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceId'", id.DeviceId, "deviceId")
	}

	if id.DeviceTemplateId != "deviceTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceTemplateId'", id.DeviceTemplateId, "deviceTemplateId")
	}
}

func TestFormatUserIdDeviceIdDeviceTemplateID(t *testing.T) {
	actual := NewUserIdDeviceIdDeviceTemplateID("userId", "deviceId", "deviceTemplateId").ID()
	expected := "/users/userId/devices/deviceId/deviceTemplate/deviceTemplateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDeviceIdDeviceTemplateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDeviceIdDeviceTemplateId
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
			Input: "/users/userId/devices/deviceId/deviceTemplate",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/devices/deviceId/deviceTemplate/deviceTemplateId",
			Expected: &UserIdDeviceIdDeviceTemplateId{
				UserId:           "userId",
				DeviceId:         "deviceId",
				DeviceTemplateId: "deviceTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/devices/deviceId/deviceTemplate/deviceTemplateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDeviceIdDeviceTemplateID(v.Input)
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

		if actual.DeviceTemplateId != v.Expected.DeviceTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceTemplateId", v.Expected.DeviceTemplateId, actual.DeviceTemplateId)
		}

	}
}

func TestParseUserIdDeviceIdDeviceTemplateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDeviceIdDeviceTemplateId
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
			Input: "/users/userId/devices/deviceId/deviceTemplate",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dEvIcEs/dEvIcEiD/dEvIcEtEmPlAtE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/devices/deviceId/deviceTemplate/deviceTemplateId",
			Expected: &UserIdDeviceIdDeviceTemplateId{
				UserId:           "userId",
				DeviceId:         "deviceId",
				DeviceTemplateId: "deviceTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/devices/deviceId/deviceTemplate/deviceTemplateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dEvIcEs/dEvIcEiD/dEvIcEtEmPlAtE/dEvIcEtEmPlAtEiD",
			Expected: &UserIdDeviceIdDeviceTemplateId{
				UserId:           "uSeRiD",
				DeviceId:         "dEvIcEiD",
				DeviceTemplateId: "dEvIcEtEmPlAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dEvIcEs/dEvIcEiD/dEvIcEtEmPlAtE/dEvIcEtEmPlAtEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDeviceIdDeviceTemplateIDInsensitively(v.Input)
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

		if actual.DeviceTemplateId != v.Expected.DeviceTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceTemplateId", v.Expected.DeviceTemplateId, actual.DeviceTemplateId)
		}

	}
}

func TestSegmentsForUserIdDeviceIdDeviceTemplateId(t *testing.T) {
	segments := UserIdDeviceIdDeviceTemplateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDeviceIdDeviceTemplateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
