package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdNotificationId{}

func TestNewUserIdNotificationID(t *testing.T) {
	id := NewUserIdNotificationID("userId", "notificationId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.NotificationId != "notificationId" {
		t.Fatalf("Expected %q but got %q for Segment 'NotificationId'", id.NotificationId, "notificationId")
	}
}

func TestFormatUserIdNotificationID(t *testing.T) {
	actual := NewUserIdNotificationID("userId", "notificationId").ID()
	expected := "/users/userId/notifications/notificationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdNotificationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdNotificationId
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
			Input: "/users/userId/notifications",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/notifications/notificationId",
			Expected: &UserIdNotificationId{
				UserId:         "userId",
				NotificationId: "notificationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/notifications/notificationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdNotificationID(v.Input)
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

		if actual.NotificationId != v.Expected.NotificationId {
			t.Fatalf("Expected %q but got %q for NotificationId", v.Expected.NotificationId, actual.NotificationId)
		}

	}
}

func TestParseUserIdNotificationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdNotificationId
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
			Input: "/users/userId/notifications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/nOtIfIcAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/notifications/notificationId",
			Expected: &UserIdNotificationId{
				UserId:         "userId",
				NotificationId: "notificationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/notifications/notificationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/nOtIfIcAtIoNs/nOtIfIcAtIoNiD",
			Expected: &UserIdNotificationId{
				UserId:         "uSeRiD",
				NotificationId: "nOtIfIcAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/nOtIfIcAtIoNs/nOtIfIcAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdNotificationIDInsensitively(v.Input)
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

		if actual.NotificationId != v.Expected.NotificationId {
			t.Fatalf("Expected %q but got %q for NotificationId", v.Expected.NotificationId, actual.NotificationId)
		}

	}
}

func TestSegmentsForUserIdNotificationId(t *testing.T) {
	segments := UserIdNotificationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdNotificationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
