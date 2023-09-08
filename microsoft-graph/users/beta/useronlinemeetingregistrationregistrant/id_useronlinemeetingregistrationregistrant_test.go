package useronlinemeetingregistrationregistrant

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnlineMeetingRegistrationRegistrantId{}

func TestNewUserOnlineMeetingRegistrationRegistrantID(t *testing.T) {
	id := NewUserOnlineMeetingRegistrationRegistrantID("userIdValue", "onlineMeetingIdValue", "meetingRegistrantBaseIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.OnlineMeetingId != "onlineMeetingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingIdValue")
	}

	if id.MeetingRegistrantBaseId != "meetingRegistrantBaseIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MeetingRegistrantBaseId'", id.MeetingRegistrantBaseId, "meetingRegistrantBaseIdValue")
	}
}

func TestFormatUserOnlineMeetingRegistrationRegistrantID(t *testing.T) {
	actual := NewUserOnlineMeetingRegistrationRegistrantID("userIdValue", "onlineMeetingIdValue", "meetingRegistrantBaseIdValue").ID()
	expected := "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/registration/registrants/meetingRegistrantBaseIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserOnlineMeetingRegistrationRegistrantID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnlineMeetingRegistrationRegistrantId
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
			Input: "/users/userIdValue/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/registration",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/registration/registrants",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/registration/registrants/meetingRegistrantBaseIdValue",
			Expected: &UserOnlineMeetingRegistrationRegistrantId{
				UserId:                  "userIdValue",
				OnlineMeetingId:         "onlineMeetingIdValue",
				MeetingRegistrantBaseId: "meetingRegistrantBaseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/registration/registrants/meetingRegistrantBaseIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnlineMeetingRegistrationRegistrantID(v.Input)
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

		if actual.OnlineMeetingId != v.Expected.OnlineMeetingId {
			t.Fatalf("Expected %q but got %q for OnlineMeetingId", v.Expected.OnlineMeetingId, actual.OnlineMeetingId)
		}

		if actual.MeetingRegistrantBaseId != v.Expected.MeetingRegistrantBaseId {
			t.Fatalf("Expected %q but got %q for MeetingRegistrantBaseId", v.Expected.MeetingRegistrantBaseId, actual.MeetingRegistrantBaseId)
		}

	}
}

func TestParseUserOnlineMeetingRegistrationRegistrantIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnlineMeetingRegistrationRegistrantId
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
			Input: "/users/userIdValue/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/registration",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEgIsTrAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/registration/registrants",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEgIsTrAtIoN/rEgIsTrAnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/registration/registrants/meetingRegistrantBaseIdValue",
			Expected: &UserOnlineMeetingRegistrationRegistrantId{
				UserId:                  "userIdValue",
				OnlineMeetingId:         "onlineMeetingIdValue",
				MeetingRegistrantBaseId: "meetingRegistrantBaseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/registration/registrants/meetingRegistrantBaseIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEgIsTrAtIoN/rEgIsTrAnTs/mEeTiNgReGiStRaNtBaSeIdVaLuE",
			Expected: &UserOnlineMeetingRegistrationRegistrantId{
				UserId:                  "uSeRiDvAlUe",
				OnlineMeetingId:         "oNlInEmEeTiNgIdVaLuE",
				MeetingRegistrantBaseId: "mEeTiNgReGiStRaNtBaSeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEgIsTrAtIoN/rEgIsTrAnTs/mEeTiNgReGiStRaNtBaSeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnlineMeetingRegistrationRegistrantIDInsensitively(v.Input)
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

		if actual.OnlineMeetingId != v.Expected.OnlineMeetingId {
			t.Fatalf("Expected %q but got %q for OnlineMeetingId", v.Expected.OnlineMeetingId, actual.OnlineMeetingId)
		}

		if actual.MeetingRegistrantBaseId != v.Expected.MeetingRegistrantBaseId {
			t.Fatalf("Expected %q but got %q for MeetingRegistrantBaseId", v.Expected.MeetingRegistrantBaseId, actual.MeetingRegistrantBaseId)
		}

	}
}

func TestSegmentsForUserOnlineMeetingRegistrationRegistrantId(t *testing.T) {
	segments := UserOnlineMeetingRegistrationRegistrantId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserOnlineMeetingRegistrationRegistrantId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
