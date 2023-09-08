package meonlinemeetingregistrationregistrant

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnlineMeetingRegistrationRegistrantId{}

func TestNewMeOnlineMeetingRegistrationRegistrantID(t *testing.T) {
	id := NewMeOnlineMeetingRegistrationRegistrantID("onlineMeetingIdValue", "meetingRegistrantBaseIdValue")

	if id.OnlineMeetingId != "onlineMeetingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingIdValue")
	}

	if id.MeetingRegistrantBaseId != "meetingRegistrantBaseIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MeetingRegistrantBaseId'", id.MeetingRegistrantBaseId, "meetingRegistrantBaseIdValue")
	}
}

func TestFormatMeOnlineMeetingRegistrationRegistrantID(t *testing.T) {
	actual := NewMeOnlineMeetingRegistrationRegistrantID("onlineMeetingIdValue", "meetingRegistrantBaseIdValue").ID()
	expected := "/me/onlineMeetings/onlineMeetingIdValue/registration/registrants/meetingRegistrantBaseIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnlineMeetingRegistrationRegistrantID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingRegistrationRegistrantId
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
			Input: "/me/onlineMeetings/onlineMeetingIdValue/registration/registrants",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/registration/registrants/meetingRegistrantBaseIdValue",
			Expected: &MeOnlineMeetingRegistrationRegistrantId{
				OnlineMeetingId:         "onlineMeetingIdValue",
				MeetingRegistrantBaseId: "meetingRegistrantBaseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingIdValue/registration/registrants/meetingRegistrantBaseIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingRegistrationRegistrantID(v.Input)
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

		if actual.MeetingRegistrantBaseId != v.Expected.MeetingRegistrantBaseId {
			t.Fatalf("Expected %q but got %q for MeetingRegistrantBaseId", v.Expected.MeetingRegistrantBaseId, actual.MeetingRegistrantBaseId)
		}

	}
}

func TestParseMeOnlineMeetingRegistrationRegistrantIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingRegistrationRegistrantId
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
			Input: "/me/onlineMeetings/onlineMeetingIdValue/registration/registrants",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEgIsTrAtIoN/rEgIsTrAnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/registration/registrants/meetingRegistrantBaseIdValue",
			Expected: &MeOnlineMeetingRegistrationRegistrantId{
				OnlineMeetingId:         "onlineMeetingIdValue",
				MeetingRegistrantBaseId: "meetingRegistrantBaseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingIdValue/registration/registrants/meetingRegistrantBaseIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEgIsTrAtIoN/rEgIsTrAnTs/mEeTiNgReGiStRaNtBaSeIdVaLuE",
			Expected: &MeOnlineMeetingRegistrationRegistrantId{
				OnlineMeetingId:         "oNlInEmEeTiNgIdVaLuE",
				MeetingRegistrantBaseId: "mEeTiNgReGiStRaNtBaSeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEgIsTrAtIoN/rEgIsTrAnTs/mEeTiNgReGiStRaNtBaSeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingRegistrationRegistrantIDInsensitively(v.Input)
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

		if actual.MeetingRegistrantBaseId != v.Expected.MeetingRegistrantBaseId {
			t.Fatalf("Expected %q but got %q for MeetingRegistrantBaseId", v.Expected.MeetingRegistrantBaseId, actual.MeetingRegistrantBaseId)
		}

	}
}

func TestSegmentsForMeOnlineMeetingRegistrationRegistrantId(t *testing.T) {
	segments := MeOnlineMeetingRegistrationRegistrantId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnlineMeetingRegistrationRegistrantId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
