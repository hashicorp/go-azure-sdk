package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingIdRegistrationRegistrantId{}

func TestNewUserIdOnlineMeetingIdRegistrationRegistrantID(t *testing.T) {
	id := NewUserIdOnlineMeetingIdRegistrationRegistrantID("userId", "onlineMeetingId", "meetingRegistrantBaseId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.OnlineMeetingId != "onlineMeetingId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingId")
	}

	if id.MeetingRegistrantBaseId != "meetingRegistrantBaseId" {
		t.Fatalf("Expected %q but got %q for Segment 'MeetingRegistrantBaseId'", id.MeetingRegistrantBaseId, "meetingRegistrantBaseId")
	}
}

func TestFormatUserIdOnlineMeetingIdRegistrationRegistrantID(t *testing.T) {
	actual := NewUserIdOnlineMeetingIdRegistrationRegistrantID("userId", "onlineMeetingId", "meetingRegistrantBaseId").ID()
	expected := "/users/userId/onlineMeetings/onlineMeetingId/registration/registrants/meetingRegistrantBaseId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdOnlineMeetingIdRegistrationRegistrantID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnlineMeetingIdRegistrationRegistrantId
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
			Input: "/users/userId/onlineMeetings/onlineMeetingId/registration/registrants",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/registration/registrants/meetingRegistrantBaseId",
			Expected: &UserIdOnlineMeetingIdRegistrationRegistrantId{
				UserId:                  "userId",
				OnlineMeetingId:         "onlineMeetingId",
				MeetingRegistrantBaseId: "meetingRegistrantBaseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onlineMeetings/onlineMeetingId/registration/registrants/meetingRegistrantBaseId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnlineMeetingIdRegistrationRegistrantID(v.Input)
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

func TestParseUserIdOnlineMeetingIdRegistrationRegistrantIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnlineMeetingIdRegistrationRegistrantId
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
			Input: "/users/userId/onlineMeetings/onlineMeetingId/registration/registrants",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEgIsTrAtIoN/rEgIsTrAnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/registration/registrants/meetingRegistrantBaseId",
			Expected: &UserIdOnlineMeetingIdRegistrationRegistrantId{
				UserId:                  "userId",
				OnlineMeetingId:         "onlineMeetingId",
				MeetingRegistrantBaseId: "meetingRegistrantBaseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onlineMeetings/onlineMeetingId/registration/registrants/meetingRegistrantBaseId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEgIsTrAtIoN/rEgIsTrAnTs/mEeTiNgReGiStRaNtBaSeId",
			Expected: &UserIdOnlineMeetingIdRegistrationRegistrantId{
				UserId:                  "uSeRiD",
				OnlineMeetingId:         "oNlInEmEeTiNgId",
				MeetingRegistrantBaseId: "mEeTiNgReGiStRaNtBaSeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEgIsTrAtIoN/rEgIsTrAnTs/mEeTiNgReGiStRaNtBaSeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnlineMeetingIdRegistrationRegistrantIDInsensitively(v.Input)
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

func TestSegmentsForUserIdOnlineMeetingIdRegistrationRegistrantId(t *testing.T) {
	segments := UserIdOnlineMeetingIdRegistrationRegistrantId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdOnlineMeetingIdRegistrationRegistrantId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
