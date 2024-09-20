package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedDeviceIdDeviceConfigurationStateId{}

func TestNewUserIdManagedDeviceIdDeviceConfigurationStateID(t *testing.T) {
	id := NewUserIdManagedDeviceIdDeviceConfigurationStateID("userId", "managedDeviceId", "deviceConfigurationStateId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ManagedDeviceId != "managedDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceId")
	}

	if id.DeviceConfigurationStateId != "deviceConfigurationStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationStateId'", id.DeviceConfigurationStateId, "deviceConfigurationStateId")
	}
}

func TestFormatUserIdManagedDeviceIdDeviceConfigurationStateID(t *testing.T) {
	actual := NewUserIdManagedDeviceIdDeviceConfigurationStateID("userId", "managedDeviceId", "deviceConfigurationStateId").ID()
	expected := "/users/userId/managedDevices/managedDeviceId/deviceConfigurationStates/deviceConfigurationStateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdManagedDeviceIdDeviceConfigurationStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdManagedDeviceIdDeviceConfigurationStateId
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
			Input: "/users/userId/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId/deviceConfigurationStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/managedDevices/managedDeviceId/deviceConfigurationStates/deviceConfigurationStateId",
			Expected: &UserIdManagedDeviceIdDeviceConfigurationStateId{
				UserId:                     "userId",
				ManagedDeviceId:            "managedDeviceId",
				DeviceConfigurationStateId: "deviceConfigurationStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/managedDevices/managedDeviceId/deviceConfigurationStates/deviceConfigurationStateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdManagedDeviceIdDeviceConfigurationStateID(v.Input)
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

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.DeviceConfigurationStateId != v.Expected.DeviceConfigurationStateId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationStateId", v.Expected.DeviceConfigurationStateId, actual.DeviceConfigurationStateId)
		}

	}
}

func TestParseUserIdManagedDeviceIdDeviceConfigurationStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdManagedDeviceIdDeviceConfigurationStateId
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
			Input: "/users/userId/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId/deviceConfigurationStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/dEvIcEcOnFiGuRaTiOnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/managedDevices/managedDeviceId/deviceConfigurationStates/deviceConfigurationStateId",
			Expected: &UserIdManagedDeviceIdDeviceConfigurationStateId{
				UserId:                     "userId",
				ManagedDeviceId:            "managedDeviceId",
				DeviceConfigurationStateId: "deviceConfigurationStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/managedDevices/managedDeviceId/deviceConfigurationStates/deviceConfigurationStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/dEvIcEcOnFiGuRaTiOnStAtEs/dEvIcEcOnFiGuRaTiOnStAtEiD",
			Expected: &UserIdManagedDeviceIdDeviceConfigurationStateId{
				UserId:                     "uSeRiD",
				ManagedDeviceId:            "mAnAgEdDeViCeId",
				DeviceConfigurationStateId: "dEvIcEcOnFiGuRaTiOnStAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/dEvIcEcOnFiGuRaTiOnStAtEs/dEvIcEcOnFiGuRaTiOnStAtEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdManagedDeviceIdDeviceConfigurationStateIDInsensitively(v.Input)
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

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.DeviceConfigurationStateId != v.Expected.DeviceConfigurationStateId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationStateId", v.Expected.DeviceConfigurationStateId, actual.DeviceConfigurationStateId)
		}

	}
}

func TestSegmentsForUserIdManagedDeviceIdDeviceConfigurationStateId(t *testing.T) {
	segments := UserIdManagedDeviceIdDeviceConfigurationStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdManagedDeviceIdDeviceConfigurationStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
