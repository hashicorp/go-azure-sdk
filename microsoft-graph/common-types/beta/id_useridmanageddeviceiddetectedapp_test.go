package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedDeviceIdDetectedAppId{}

func TestNewUserIdManagedDeviceIdDetectedAppID(t *testing.T) {
	id := NewUserIdManagedDeviceIdDetectedAppID("userIdValue", "managedDeviceIdValue", "detectedAppIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ManagedDeviceId != "managedDeviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceIdValue")
	}

	if id.DetectedAppId != "detectedAppIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DetectedAppId'", id.DetectedAppId, "detectedAppIdValue")
	}
}

func TestFormatUserIdManagedDeviceIdDetectedAppID(t *testing.T) {
	actual := NewUserIdManagedDeviceIdDetectedAppID("userIdValue", "managedDeviceIdValue", "detectedAppIdValue").ID()
	expected := "/users/userIdValue/managedDevices/managedDeviceIdValue/detectedApps/detectedAppIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdManagedDeviceIdDetectedAppID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdManagedDeviceIdDetectedAppId
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
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/detectedApps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/detectedApps/detectedAppIdValue",
			Expected: &UserIdManagedDeviceIdDetectedAppId{
				UserId:          "userIdValue",
				ManagedDeviceId: "managedDeviceIdValue",
				DetectedAppId:   "detectedAppIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/detectedApps/detectedAppIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdManagedDeviceIdDetectedAppID(v.Input)
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

		if actual.DetectedAppId != v.Expected.DetectedAppId {
			t.Fatalf("Expected %q but got %q for DetectedAppId", v.Expected.DetectedAppId, actual.DetectedAppId)
		}

	}
}

func TestParseUserIdManagedDeviceIdDetectedAppIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdManagedDeviceIdDetectedAppId
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
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/detectedApps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/dEtEcTeDaPpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/detectedApps/detectedAppIdValue",
			Expected: &UserIdManagedDeviceIdDetectedAppId{
				UserId:          "userIdValue",
				ManagedDeviceId: "managedDeviceIdValue",
				DetectedAppId:   "detectedAppIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/detectedApps/detectedAppIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/dEtEcTeDaPpS/dEtEcTeDaPpIdVaLuE",
			Expected: &UserIdManagedDeviceIdDetectedAppId{
				UserId:          "uSeRiDvAlUe",
				ManagedDeviceId: "mAnAgEdDeViCeIdVaLuE",
				DetectedAppId:   "dEtEcTeDaPpIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/dEtEcTeDaPpS/dEtEcTeDaPpIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdManagedDeviceIdDetectedAppIDInsensitively(v.Input)
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

		if actual.DetectedAppId != v.Expected.DetectedAppId {
			t.Fatalf("Expected %q but got %q for DetectedAppId", v.Expected.DetectedAppId, actual.DetectedAppId)
		}

	}
}

func TestSegmentsForUserIdManagedDeviceIdDetectedAppId(t *testing.T) {
	segments := UserIdManagedDeviceIdDetectedAppId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdManagedDeviceIdDetectedAppId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
