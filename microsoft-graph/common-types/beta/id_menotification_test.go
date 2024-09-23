package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeNotificationId{}

func TestNewMeNotificationID(t *testing.T) {
	id := NewMeNotificationID("notificationId")

	if id.NotificationId != "notificationId" {
		t.Fatalf("Expected %q but got %q for Segment 'NotificationId'", id.NotificationId, "notificationId")
	}
}

func TestFormatMeNotificationID(t *testing.T) {
	actual := NewMeNotificationID("notificationId").ID()
	expected := "/me/notifications/notificationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeNotificationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeNotificationId
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
			Input: "/me/notifications",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/notifications/notificationId",
			Expected: &MeNotificationId{
				NotificationId: "notificationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/notifications/notificationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeNotificationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.NotificationId != v.Expected.NotificationId {
			t.Fatalf("Expected %q but got %q for NotificationId", v.Expected.NotificationId, actual.NotificationId)
		}

	}
}

func TestParseMeNotificationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeNotificationId
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
			Input: "/me/notifications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/nOtIfIcAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/notifications/notificationId",
			Expected: &MeNotificationId{
				NotificationId: "notificationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/notifications/notificationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/nOtIfIcAtIoNs/nOtIfIcAtIoNiD",
			Expected: &MeNotificationId{
				NotificationId: "nOtIfIcAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/nOtIfIcAtIoNs/nOtIfIcAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeNotificationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.NotificationId != v.Expected.NotificationId {
			t.Fatalf("Expected %q but got %q for NotificationId", v.Expected.NotificationId, actual.NotificationId)
		}

	}
}

func TestSegmentsForMeNotificationId(t *testing.T) {
	segments := MeNotificationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeNotificationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
