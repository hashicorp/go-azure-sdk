package meonlinemeetingregistrationcustomquestion

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnlineMeetingRegistrationCustomQuestionId{}

func TestNewMeOnlineMeetingRegistrationCustomQuestionID(t *testing.T) {
	id := NewMeOnlineMeetingRegistrationCustomQuestionID("onlineMeetingIdValue", "meetingRegistrationQuestionIdValue")

	if id.OnlineMeetingId != "onlineMeetingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingIdValue")
	}

	if id.MeetingRegistrationQuestionId != "meetingRegistrationQuestionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MeetingRegistrationQuestionId'", id.MeetingRegistrationQuestionId, "meetingRegistrationQuestionIdValue")
	}
}

func TestFormatMeOnlineMeetingRegistrationCustomQuestionID(t *testing.T) {
	actual := NewMeOnlineMeetingRegistrationCustomQuestionID("onlineMeetingIdValue", "meetingRegistrationQuestionIdValue").ID()
	expected := "/me/onlineMeetings/onlineMeetingIdValue/registration/customQuestions/meetingRegistrationQuestionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnlineMeetingRegistrationCustomQuestionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingRegistrationCustomQuestionId
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
			Input: "/me/onlineMeetings/onlineMeetingIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/registration",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/registration/customQuestions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/registration/customQuestions/meetingRegistrationQuestionIdValue",
			Expected: &MeOnlineMeetingRegistrationCustomQuestionId{
				OnlineMeetingId:               "onlineMeetingIdValue",
				MeetingRegistrationQuestionId: "meetingRegistrationQuestionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingIdValue/registration/customQuestions/meetingRegistrationQuestionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingRegistrationCustomQuestionID(v.Input)
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

func TestParseMeOnlineMeetingRegistrationCustomQuestionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingRegistrationCustomQuestionId
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
			Input: "/me/onlineMeetings/onlineMeetingIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/registration",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEgIsTrAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/registration/customQuestions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEgIsTrAtIoN/cUsToMqUeStIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/registration/customQuestions/meetingRegistrationQuestionIdValue",
			Expected: &MeOnlineMeetingRegistrationCustomQuestionId{
				OnlineMeetingId:               "onlineMeetingIdValue",
				MeetingRegistrationQuestionId: "meetingRegistrationQuestionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingIdValue/registration/customQuestions/meetingRegistrationQuestionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEgIsTrAtIoN/cUsToMqUeStIoNs/mEeTiNgReGiStRaTiOnQuEsTiOnIdVaLuE",
			Expected: &MeOnlineMeetingRegistrationCustomQuestionId{
				OnlineMeetingId:               "oNlInEmEeTiNgIdVaLuE",
				MeetingRegistrationQuestionId: "mEeTiNgReGiStRaTiOnQuEsTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEgIsTrAtIoN/cUsToMqUeStIoNs/mEeTiNgReGiStRaTiOnQuEsTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingRegistrationCustomQuestionIDInsensitively(v.Input)
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

func TestSegmentsForMeOnlineMeetingRegistrationCustomQuestionId(t *testing.T) {
	segments := MeOnlineMeetingRegistrationCustomQuestionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnlineMeetingRegistrationCustomQuestionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
