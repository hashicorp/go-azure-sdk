package usermobileapptroubleshootingevent

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMobileAppTroubleshootingEventId{}

func TestNewUserMobileAppTroubleshootingEventID(t *testing.T) {
	id := NewUserMobileAppTroubleshootingEventID("userIdValue", "mobileAppTroubleshootingEventIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.MobileAppTroubleshootingEventId != "mobileAppTroubleshootingEventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MobileAppTroubleshootingEventId'", id.MobileAppTroubleshootingEventId, "mobileAppTroubleshootingEventIdValue")
	}
}

func TestFormatUserMobileAppTroubleshootingEventID(t *testing.T) {
	actual := NewUserMobileAppTroubleshootingEventID("userIdValue", "mobileAppTroubleshootingEventIdValue").ID()
	expected := "/users/userIdValue/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserMobileAppTroubleshootingEventID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserMobileAppTroubleshootingEventId
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
			Input: "/users/userIdValue/mobileAppTroubleshootingEvents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue",
			Expected: &UserMobileAppTroubleshootingEventId{
				UserId:                          "userIdValue",
				MobileAppTroubleshootingEventId: "mobileAppTroubleshootingEventIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserMobileAppTroubleshootingEventID(v.Input)
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

		if actual.MobileAppTroubleshootingEventId != v.Expected.MobileAppTroubleshootingEventId {
			t.Fatalf("Expected %q but got %q for MobileAppTroubleshootingEventId", v.Expected.MobileAppTroubleshootingEventId, actual.MobileAppTroubleshootingEventId)
		}

	}
}

func TestParseUserMobileAppTroubleshootingEventIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserMobileAppTroubleshootingEventId
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
			Input: "/users/userIdValue/mobileAppTroubleshootingEvents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mObIlEaPpTrOuBlEsHoOtInGeVeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue",
			Expected: &UserMobileAppTroubleshootingEventId{
				UserId:                          "userIdValue",
				MobileAppTroubleshootingEventId: "mobileAppTroubleshootingEventIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/mobileAppTroubleshootingEvents/mobileAppTroubleshootingEventIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtIdVaLuE",
			Expected: &UserMobileAppTroubleshootingEventId{
				UserId:                          "uSeRiDvAlUe",
				MobileAppTroubleshootingEventId: "mObIlEaPpTrOuBlEsHoOtInGeVeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mObIlEaPpTrOuBlEsHoOtInGeVeNtS/mObIlEaPpTrOuBlEsHoOtInGeVeNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserMobileAppTroubleshootingEventIDInsensitively(v.Input)
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

		if actual.MobileAppTroubleshootingEventId != v.Expected.MobileAppTroubleshootingEventId {
			t.Fatalf("Expected %q but got %q for MobileAppTroubleshootingEventId", v.Expected.MobileAppTroubleshootingEventId, actual.MobileAppTroubleshootingEventId)
		}

	}
}

func TestSegmentsForUserMobileAppTroubleshootingEventId(t *testing.T) {
	segments := UserMobileAppTroubleshootingEventId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserMobileAppTroubleshootingEventId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
