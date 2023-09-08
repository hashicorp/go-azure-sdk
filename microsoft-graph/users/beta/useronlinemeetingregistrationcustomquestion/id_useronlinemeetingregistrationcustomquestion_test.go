package useronlinemeetingregistrationcustomquestion

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnlineMeetingRegistrationCustomQuestionId{}

func TestNewUserOnlineMeetingRegistrationCustomQuestionID(t *testing.T) {
	id := NewUserOnlineMeetingRegistrationCustomQuestionID("userIdValue", "onlineMeetingIdValue", "meetingRegistrationQuestionIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.OnlineMeetingId != "onlineMeetingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingIdValue")
	}

	if id.MeetingRegistrationQuestionId != "meetingRegistrationQuestionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MeetingRegistrationQuestionId'", id.MeetingRegistrationQuestionId, "meetingRegistrationQuestionIdValue")
	}
}

func TestFormatUserOnlineMeetingRegistrationCustomQuestionID(t *testing.T) {
	actual := NewUserOnlineMeetingRegistrationCustomQuestionID("userIdValue", "onlineMeetingIdValue", "meetingRegistrationQuestionIdValue").ID()
	expected := "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/registration/customQuestions/meetingRegistrationQuestionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserOnlineMeetingRegistrationCustomQuestionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnlineMeetingRegistrationCustomQuestionId
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
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/registration/customQuestions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/registration/customQuestions/meetingRegistrationQuestionIdValue",
			Expected: &UserOnlineMeetingRegistrationCustomQuestionId{
				UserId:                        "userIdValue",
				OnlineMeetingId:               "onlineMeetingIdValue",
				MeetingRegistrationQuestionId: "meetingRegistrationQuestionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/registration/customQuestions/meetingRegistrationQuestionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnlineMeetingRegistrationCustomQuestionID(v.Input)
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

		if actual.MeetingRegistrationQuestionId != v.Expected.MeetingRegistrationQuestionId {
			t.Fatalf("Expected %q but got %q for MeetingRegistrationQuestionId", v.Expected.MeetingRegistrationQuestionId, actual.MeetingRegistrationQuestionId)
		}

	}
}

func TestParseUserOnlineMeetingRegistrationCustomQuestionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnlineMeetingRegistrationCustomQuestionId
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
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/registration/customQuestions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEgIsTrAtIoN/cUsToMqUeStIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/registration/customQuestions/meetingRegistrationQuestionIdValue",
			Expected: &UserOnlineMeetingRegistrationCustomQuestionId{
				UserId:                        "userIdValue",
				OnlineMeetingId:               "onlineMeetingIdValue",
				MeetingRegistrationQuestionId: "meetingRegistrationQuestionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/registration/customQuestions/meetingRegistrationQuestionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEgIsTrAtIoN/cUsToMqUeStIoNs/mEeTiNgReGiStRaTiOnQuEsTiOnIdVaLuE",
			Expected: &UserOnlineMeetingRegistrationCustomQuestionId{
				UserId:                        "uSeRiDvAlUe",
				OnlineMeetingId:               "oNlInEmEeTiNgIdVaLuE",
				MeetingRegistrationQuestionId: "mEeTiNgReGiStRaTiOnQuEsTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEgIsTrAtIoN/cUsToMqUeStIoNs/mEeTiNgReGiStRaTiOnQuEsTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnlineMeetingRegistrationCustomQuestionIDInsensitively(v.Input)
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

		if actual.MeetingRegistrationQuestionId != v.Expected.MeetingRegistrationQuestionId {
			t.Fatalf("Expected %q but got %q for MeetingRegistrationQuestionId", v.Expected.MeetingRegistrationQuestionId, actual.MeetingRegistrationQuestionId)
		}

	}
}

func TestSegmentsForUserOnlineMeetingRegistrationCustomQuestionId(t *testing.T) {
	segments := UserOnlineMeetingRegistrationCustomQuestionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserOnlineMeetingRegistrationCustomQuestionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
