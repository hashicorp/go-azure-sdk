package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnlineMeetingIdRegistrationCustomQuestionId{}

func TestNewMeOnlineMeetingIdRegistrationCustomQuestionID(t *testing.T) {
	id := NewMeOnlineMeetingIdRegistrationCustomQuestionID("onlineMeetingId", "meetingRegistrationQuestionId")

	if id.OnlineMeetingId != "onlineMeetingId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingId")
	}

	if id.MeetingRegistrationQuestionId != "meetingRegistrationQuestionId" {
		t.Fatalf("Expected %q but got %q for Segment 'MeetingRegistrationQuestionId'", id.MeetingRegistrationQuestionId, "meetingRegistrationQuestionId")
	}
}

func TestFormatMeOnlineMeetingIdRegistrationCustomQuestionID(t *testing.T) {
	actual := NewMeOnlineMeetingIdRegistrationCustomQuestionID("onlineMeetingId", "meetingRegistrationQuestionId").ID()
	expected := "/me/onlineMeetings/onlineMeetingId/registration/customQuestions/meetingRegistrationQuestionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnlineMeetingIdRegistrationCustomQuestionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingIdRegistrationCustomQuestionId
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
			Input: "/me/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingId/registration",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingId/registration/customQuestions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingId/registration/customQuestions/meetingRegistrationQuestionId",
			Expected: &MeOnlineMeetingIdRegistrationCustomQuestionId{
				OnlineMeetingId:               "onlineMeetingId",
				MeetingRegistrationQuestionId: "meetingRegistrationQuestionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingId/registration/customQuestions/meetingRegistrationQuestionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingIdRegistrationCustomQuestionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OnlineMeetingId != v.Expected.OnlineMeetingId {
			t.Fatalf("Expected %q but got %q for OnlineMeetingId", v.Expected.OnlineMeetingId, actual.OnlineMeetingId)
		}

		if actual.MeetingRegistrationQuestionId != v.Expected.MeetingRegistrationQuestionId {
			t.Fatalf("Expected %q but got %q for MeetingRegistrationQuestionId", v.Expected.MeetingRegistrationQuestionId, actual.MeetingRegistrationQuestionId)
		}

	}
}

func TestParseMeOnlineMeetingIdRegistrationCustomQuestionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingIdRegistrationCustomQuestionId
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
			Input: "/me/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingId/registration",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEgIsTrAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingId/registration/customQuestions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEgIsTrAtIoN/cUsToMqUeStIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingId/registration/customQuestions/meetingRegistrationQuestionId",
			Expected: &MeOnlineMeetingIdRegistrationCustomQuestionId{
				OnlineMeetingId:               "onlineMeetingId",
				MeetingRegistrationQuestionId: "meetingRegistrationQuestionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingId/registration/customQuestions/meetingRegistrationQuestionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEgIsTrAtIoN/cUsToMqUeStIoNs/mEeTiNgReGiStRaTiOnQuEsTiOnId",
			Expected: &MeOnlineMeetingIdRegistrationCustomQuestionId{
				OnlineMeetingId:               "oNlInEmEeTiNgId",
				MeetingRegistrationQuestionId: "mEeTiNgReGiStRaTiOnQuEsTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEgIsTrAtIoN/cUsToMqUeStIoNs/mEeTiNgReGiStRaTiOnQuEsTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingIdRegistrationCustomQuestionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OnlineMeetingId != v.Expected.OnlineMeetingId {
			t.Fatalf("Expected %q but got %q for OnlineMeetingId", v.Expected.OnlineMeetingId, actual.OnlineMeetingId)
		}

		if actual.MeetingRegistrationQuestionId != v.Expected.MeetingRegistrationQuestionId {
			t.Fatalf("Expected %q but got %q for MeetingRegistrationQuestionId", v.Expected.MeetingRegistrationQuestionId, actual.MeetingRegistrationQuestionId)
		}

	}
}

func TestSegmentsForMeOnlineMeetingIdRegistrationCustomQuestionId(t *testing.T) {
	segments := MeOnlineMeetingIdRegistrationCustomQuestionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnlineMeetingIdRegistrationCustomQuestionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
