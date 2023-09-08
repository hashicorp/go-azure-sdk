package usermanageddevicelogcollectionrequest

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserManagedDeviceLogCollectionRequestId{}

func TestNewUserManagedDeviceLogCollectionRequestID(t *testing.T) {
	id := NewUserManagedDeviceLogCollectionRequestID("userIdValue", "managedDeviceIdValue", "deviceLogCollectionResponseIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ManagedDeviceId != "managedDeviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceIdValue")
	}

	if id.DeviceLogCollectionResponseId != "deviceLogCollectionResponseIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceLogCollectionResponseId'", id.DeviceLogCollectionResponseId, "deviceLogCollectionResponseIdValue")
	}
}

func TestFormatUserManagedDeviceLogCollectionRequestID(t *testing.T) {
	actual := NewUserManagedDeviceLogCollectionRequestID("userIdValue", "managedDeviceIdValue", "deviceLogCollectionResponseIdValue").ID()
	expected := "/users/userIdValue/managedDevices/managedDeviceIdValue/logCollectionRequests/deviceLogCollectionResponseIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserManagedDeviceLogCollectionRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserManagedDeviceLogCollectionRequestId
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
			Input: "/users/userIdValue/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/logCollectionRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/logCollectionRequests/deviceLogCollectionResponseIdValue",
			Expected: &UserManagedDeviceLogCollectionRequestId{
				UserId:                        "userIdValue",
				ManagedDeviceId:               "managedDeviceIdValue",
				DeviceLogCollectionResponseId: "deviceLogCollectionResponseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/logCollectionRequests/deviceLogCollectionResponseIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserManagedDeviceLogCollectionRequestID(v.Input)
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

func TestParseUserManagedDeviceLogCollectionRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserManagedDeviceLogCollectionRequestId
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
			Input: "/users/userIdValue/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/logCollectionRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/lOgCoLlEcTiOnReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/logCollectionRequests/deviceLogCollectionResponseIdValue",
			Expected: &UserManagedDeviceLogCollectionRequestId{
				UserId:                        "userIdValue",
				ManagedDeviceId:               "managedDeviceIdValue",
				DeviceLogCollectionResponseId: "deviceLogCollectionResponseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/logCollectionRequests/deviceLogCollectionResponseIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/lOgCoLlEcTiOnReQuEsTs/dEvIcElOgCoLlEcTiOnReSpOnSeIdVaLuE",
			Expected: &UserManagedDeviceLogCollectionRequestId{
				UserId:                        "uSeRiDvAlUe",
				ManagedDeviceId:               "mAnAgEdDeViCeIdVaLuE",
				DeviceLogCollectionResponseId: "dEvIcElOgCoLlEcTiOnReSpOnSeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/lOgCoLlEcTiOnReQuEsTs/dEvIcElOgCoLlEcTiOnReSpOnSeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserManagedDeviceLogCollectionRequestIDInsensitively(v.Input)
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

func TestSegmentsForUserManagedDeviceLogCollectionRequestId(t *testing.T) {
	segments := UserManagedDeviceLogCollectionRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserManagedDeviceLogCollectionRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
