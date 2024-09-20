package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdWindowsInformationProtectionDeviceRegistrationId{}

func TestNewUserIdWindowsInformationProtectionDeviceRegistrationID(t *testing.T) {
	id := NewUserIdWindowsInformationProtectionDeviceRegistrationID("userId", "windowsInformationProtectionDeviceRegistrationId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.WindowsInformationProtectionDeviceRegistrationId != "windowsInformationProtectionDeviceRegistrationId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsInformationProtectionDeviceRegistrationId'", id.WindowsInformationProtectionDeviceRegistrationId, "windowsInformationProtectionDeviceRegistrationId")
	}
}

func TestFormatUserIdWindowsInformationProtectionDeviceRegistrationID(t *testing.T) {
	actual := NewUserIdWindowsInformationProtectionDeviceRegistrationID("userId", "windowsInformationProtectionDeviceRegistrationId").ID()
	expected := "/users/userId/windowsInformationProtectionDeviceRegistrations/windowsInformationProtectionDeviceRegistrationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdWindowsInformationProtectionDeviceRegistrationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdWindowsInformationProtectionDeviceRegistrationId
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
			Input: "/users/userId/windowsInformationProtectionDeviceRegistrations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/windowsInformationProtectionDeviceRegistrations/windowsInformationProtectionDeviceRegistrationId",
			Expected: &UserIdWindowsInformationProtectionDeviceRegistrationId{
				UserId: "userId",
				WindowsInformationProtectionDeviceRegistrationId: "windowsInformationProtectionDeviceRegistrationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/windowsInformationProtectionDeviceRegistrations/windowsInformationProtectionDeviceRegistrationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdWindowsInformationProtectionDeviceRegistrationID(v.Input)
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

		if actual.WindowsInformationProtectionDeviceRegistrationId != v.Expected.WindowsInformationProtectionDeviceRegistrationId {
			t.Fatalf("Expected %q but got %q for WindowsInformationProtectionDeviceRegistrationId", v.Expected.WindowsInformationProtectionDeviceRegistrationId, actual.WindowsInformationProtectionDeviceRegistrationId)
		}

	}
}

func TestParseUserIdWindowsInformationProtectionDeviceRegistrationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdWindowsInformationProtectionDeviceRegistrationId
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
			Input: "/users/userId/windowsInformationProtectionDeviceRegistrations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/wInDoWsInFoRmAtIoNpRoTeCtIoNdEvIcErEgIsTrAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/windowsInformationProtectionDeviceRegistrations/windowsInformationProtectionDeviceRegistrationId",
			Expected: &UserIdWindowsInformationProtectionDeviceRegistrationId{
				UserId: "userId",
				WindowsInformationProtectionDeviceRegistrationId: "windowsInformationProtectionDeviceRegistrationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/windowsInformationProtectionDeviceRegistrations/windowsInformationProtectionDeviceRegistrationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/wInDoWsInFoRmAtIoNpRoTeCtIoNdEvIcErEgIsTrAtIoNs/wInDoWsInFoRmAtIoNpRoTeCtIoNdEvIcErEgIsTrAtIoNiD",
			Expected: &UserIdWindowsInformationProtectionDeviceRegistrationId{
				UserId: "uSeRiD",
				WindowsInformationProtectionDeviceRegistrationId: "wInDoWsInFoRmAtIoNpRoTeCtIoNdEvIcErEgIsTrAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/wInDoWsInFoRmAtIoNpRoTeCtIoNdEvIcErEgIsTrAtIoNs/wInDoWsInFoRmAtIoNpRoTeCtIoNdEvIcErEgIsTrAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdWindowsInformationProtectionDeviceRegistrationIDInsensitively(v.Input)
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

		if actual.WindowsInformationProtectionDeviceRegistrationId != v.Expected.WindowsInformationProtectionDeviceRegistrationId {
			t.Fatalf("Expected %q but got %q for WindowsInformationProtectionDeviceRegistrationId", v.Expected.WindowsInformationProtectionDeviceRegistrationId, actual.WindowsInformationProtectionDeviceRegistrationId)
		}

	}
}

func TestSegmentsForUserIdWindowsInformationProtectionDeviceRegistrationId(t *testing.T) {
	segments := UserIdWindowsInformationProtectionDeviceRegistrationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdWindowsInformationProtectionDeviceRegistrationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
