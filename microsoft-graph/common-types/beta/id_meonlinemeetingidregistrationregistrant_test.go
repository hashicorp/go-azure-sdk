package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnlineMeetingIdRegistrationRegistrantId{}

func TestNewMeOnlineMeetingIdRegistrationRegistrantID(t *testing.T) {
	id := NewMeOnlineMeetingIdRegistrationRegistrantID("onlineMeetingId", "meetingRegistrantBaseId")

	if id.OnlineMeetingId != "onlineMeetingId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingId")
	}

	if id.MeetingRegistrantBaseId != "meetingRegistrantBaseId" {
		t.Fatalf("Expected %q but got %q for Segment 'MeetingRegistrantBaseId'", id.MeetingRegistrantBaseId, "meetingRegistrantBaseId")
	}
}

func TestFormatMeOnlineMeetingIdRegistrationRegistrantID(t *testing.T) {
	actual := NewMeOnlineMeetingIdRegistrationRegistrantID("onlineMeetingId", "meetingRegistrantBaseId").ID()
	expected := "/me/onlineMeetings/onlineMeetingId/registration/registrants/meetingRegistrantBaseId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnlineMeetingIdRegistrationRegistrantID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingIdRegistrationRegistrantId
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
			Input: "/me/onlineMeetings/onlineMeetingId/registration/registrants",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingId/registration/registrants/meetingRegistrantBaseId",
			Expected: &MeOnlineMeetingIdRegistrationRegistrantId{
				OnlineMeetingId:         "onlineMeetingId",
				MeetingRegistrantBaseId: "meetingRegistrantBaseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingId/registration/registrants/meetingRegistrantBaseId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingIdRegistrationRegistrantID(v.Input)
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

func TestParseMeOnlineMeetingIdRegistrationRegistrantIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingIdRegistrationRegistrantId
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
			Input: "/me/onlineMeetings/onlineMeetingId/registration/registrants",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEgIsTrAtIoN/rEgIsTrAnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingId/registration/registrants/meetingRegistrantBaseId",
			Expected: &MeOnlineMeetingIdRegistrationRegistrantId{
				OnlineMeetingId:         "onlineMeetingId",
				MeetingRegistrantBaseId: "meetingRegistrantBaseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingId/registration/registrants/meetingRegistrantBaseId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEgIsTrAtIoN/rEgIsTrAnTs/mEeTiNgReGiStRaNtBaSeId",
			Expected: &MeOnlineMeetingIdRegistrationRegistrantId{
				OnlineMeetingId:         "oNlInEmEeTiNgId",
				MeetingRegistrantBaseId: "mEeTiNgReGiStRaNtBaSeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEgIsTrAtIoN/rEgIsTrAnTs/mEeTiNgReGiStRaNtBaSeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingIdRegistrationRegistrantIDInsensitively(v.Input)
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

func TestSegmentsForMeOnlineMeetingIdRegistrationRegistrantId(t *testing.T) {
	segments := MeOnlineMeetingIdRegistrationRegistrantId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnlineMeetingIdRegistrationRegistrantId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
