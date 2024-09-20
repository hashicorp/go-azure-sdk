package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedDeviceIdLogCollectionRequestId{}

func TestNewUserIdManagedDeviceIdLogCollectionRequestID(t *testing.T) {
	id := NewUserIdManagedDeviceIdLogCollectionRequestID("userId", "managedDeviceId", "deviceLogCollectionResponseId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ManagedDeviceId != "managedDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceId")
	}

	if id.DeviceLogCollectionResponseId != "deviceLogCollectionResponseId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceLogCollectionResponseId'", id.DeviceLogCollectionResponseId, "deviceLogCollectionResponseId")
	}
}

func TestFormatUserIdManagedDeviceIdLogCollectionRequestID(t *testing.T) {
	actual := NewUserIdManagedDeviceIdLogCollectionRequestID("userId", "managedDeviceId", "deviceLogCollectionResponseId").ID()
	expected := "/users/userId/managedDevices/managedDeviceId/logCollectionRequests/deviceLogCollectionResponseId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdManagedDeviceIdLogCollectionRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdManagedDeviceIdLogCollectionRequestId
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
			Input: "/users/userId/managedDevices/managedDeviceId/logCollectionRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/managedDevices/managedDeviceId/logCollectionRequests/deviceLogCollectionResponseId",
			Expected: &UserIdManagedDeviceIdLogCollectionRequestId{
				UserId:                        "userId",
				ManagedDeviceId:               "managedDeviceId",
				DeviceLogCollectionResponseId: "deviceLogCollectionResponseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/managedDevices/managedDeviceId/logCollectionRequests/deviceLogCollectionResponseId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdManagedDeviceIdLogCollectionRequestID(v.Input)
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

		if actual.DeviceLogCollectionResponseId != v.Expected.DeviceLogCollectionResponseId {
			t.Fatalf("Expected %q but got %q for DeviceLogCollectionResponseId", v.Expected.DeviceLogCollectionResponseId, actual.DeviceLogCollectionResponseId)
		}

	}
}

func TestParseUserIdManagedDeviceIdLogCollectionRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdManagedDeviceIdLogCollectionRequestId
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
			Input: "/users/userId/managedDevices/managedDeviceId/logCollectionRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/lOgCoLlEcTiOnReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/managedDevices/managedDeviceId/logCollectionRequests/deviceLogCollectionResponseId",
			Expected: &UserIdManagedDeviceIdLogCollectionRequestId{
				UserId:                        "userId",
				ManagedDeviceId:               "managedDeviceId",
				DeviceLogCollectionResponseId: "deviceLogCollectionResponseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/managedDevices/managedDeviceId/logCollectionRequests/deviceLogCollectionResponseId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/lOgCoLlEcTiOnReQuEsTs/dEvIcElOgCoLlEcTiOnReSpOnSeId",
			Expected: &UserIdManagedDeviceIdLogCollectionRequestId{
				UserId:                        "uSeRiD",
				ManagedDeviceId:               "mAnAgEdDeViCeId",
				DeviceLogCollectionResponseId: "dEvIcElOgCoLlEcTiOnReSpOnSeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/lOgCoLlEcTiOnReQuEsTs/dEvIcElOgCoLlEcTiOnReSpOnSeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdManagedDeviceIdLogCollectionRequestIDInsensitively(v.Input)
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

		if actual.DeviceLogCollectionResponseId != v.Expected.DeviceLogCollectionResponseId {
			t.Fatalf("Expected %q but got %q for DeviceLogCollectionResponseId", v.Expected.DeviceLogCollectionResponseId, actual.DeviceLogCollectionResponseId)
		}

	}
}

func TestSegmentsForUserIdManagedDeviceIdLogCollectionRequestId(t *testing.T) {
	segments := UserIdManagedDeviceIdLogCollectionRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdManagedDeviceIdLogCollectionRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
