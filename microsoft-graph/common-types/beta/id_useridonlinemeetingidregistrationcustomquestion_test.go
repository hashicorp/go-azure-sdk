package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingIdRegistrationCustomQuestionId{}

func TestNewUserIdOnlineMeetingIdRegistrationCustomQuestionID(t *testing.T) {
	id := NewUserIdOnlineMeetingIdRegistrationCustomQuestionID("userId", "onlineMeetingId", "meetingRegistrationQuestionId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.OnlineMeetingId != "onlineMeetingId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingId")
	}

	if id.MeetingRegistrationQuestionId != "meetingRegistrationQuestionId" {
		t.Fatalf("Expected %q but got %q for Segment 'MeetingRegistrationQuestionId'", id.MeetingRegistrationQuestionId, "meetingRegistrationQuestionId")
	}
}

func TestFormatUserIdOnlineMeetingIdRegistrationCustomQuestionID(t *testing.T) {
	actual := NewUserIdOnlineMeetingIdRegistrationCustomQuestionID("userId", "onlineMeetingId", "meetingRegistrationQuestionId").ID()
	expected := "/users/userId/onlineMeetings/onlineMeetingId/registration/customQuestions/meetingRegistrationQuestionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdOnlineMeetingIdRegistrationCustomQuestionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnlineMeetingIdRegistrationCustomQuestionId
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
			Input: "/users/userId/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/registration",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/registration/customQuestions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/registration/customQuestions/meetingRegistrationQuestionId",
			Expected: &UserIdOnlineMeetingIdRegistrationCustomQuestionId{
				UserId:                        "userId",
				OnlineMeetingId:               "onlineMeetingId",
				MeetingRegistrationQuestionId: "meetingRegistrationQuestionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onlineMeetings/onlineMeetingId/registration/customQuestions/meetingRegistrationQuestionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnlineMeetingIdRegistrationCustomQuestionID(v.Input)
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

func TestParseUserIdOnlineMeetingIdRegistrationCustomQuestionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnlineMeetingIdRegistrationCustomQuestionId
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
			Input: "/users/userId/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/registration",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEgIsTrAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/registration/customQuestions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEgIsTrAtIoN/cUsToMqUeStIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/registration/customQuestions/meetingRegistrationQuestionId",
			Expected: &UserIdOnlineMeetingIdRegistrationCustomQuestionId{
				UserId:                        "userId",
				OnlineMeetingId:               "onlineMeetingId",
				MeetingRegistrationQuestionId: "meetingRegistrationQuestionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onlineMeetings/onlineMeetingId/registration/customQuestions/meetingRegistrationQuestionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEgIsTrAtIoN/cUsToMqUeStIoNs/mEeTiNgReGiStRaTiOnQuEsTiOnId",
			Expected: &UserIdOnlineMeetingIdRegistrationCustomQuestionId{
				UserId:                        "uSeRiD",
				OnlineMeetingId:               "oNlInEmEeTiNgId",
				MeetingRegistrationQuestionId: "mEeTiNgReGiStRaTiOnQuEsTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEgIsTrAtIoN/cUsToMqUeStIoNs/mEeTiNgReGiStRaTiOnQuEsTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnlineMeetingIdRegistrationCustomQuestionIDInsensitively(v.Input)
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

func TestSegmentsForUserIdOnlineMeetingIdRegistrationCustomQuestionId(t *testing.T) {
	segments := UserIdOnlineMeetingIdRegistrationCustomQuestionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdOnlineMeetingIdRegistrationCustomQuestionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
