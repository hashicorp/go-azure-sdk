package mejoinedteamscheduleoffershiftrequest

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamScheduleOfferShiftRequestId{}

func TestNewMeJoinedTeamScheduleOfferShiftRequestID(t *testing.T) {
	id := NewMeJoinedTeamScheduleOfferShiftRequestID("teamIdValue", "offerShiftRequestIdValue")

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.OfferShiftRequestId != "offerShiftRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OfferShiftRequestId'", id.OfferShiftRequestId, "offerShiftRequestIdValue")
	}
}

func TestFormatMeJoinedTeamScheduleOfferShiftRequestID(t *testing.T) {
	actual := NewMeJoinedTeamScheduleOfferShiftRequestID("teamIdValue", "offerShiftRequestIdValue").ID()
	expected := "/me/joinedTeams/teamIdValue/schedule/offerShiftRequests/offerShiftRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamScheduleOfferShiftRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamScheduleOfferShiftRequestId
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
			Input: "/me/joinedTeams",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/schedule",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/schedule/offerShiftRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/schedule/offerShiftRequests/offerShiftRequestIdValue",
			Expected: &MeJoinedTeamScheduleOfferShiftRequestId{
				TeamId:              "teamIdValue",
				OfferShiftRequestId: "offerShiftRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/schedule/offerShiftRequests/offerShiftRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamScheduleOfferShiftRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TeamId != v.Expected.TeamId {
			t.Fatalf("Expected %q but got %q for TeamId", v.Expected.TeamId, actual.TeamId)
		}

		if actual.OfferShiftRequestId != v.Expected.OfferShiftRequestId {
			t.Fatalf("Expected %q but got %q for OfferShiftRequestId", v.Expected.OfferShiftRequestId, actual.OfferShiftRequestId)
		}

	}
}

func TestParseMeJoinedTeamScheduleOfferShiftRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamScheduleOfferShiftRequestId
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
			Input: "/me/joinedTeams",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/schedule",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/schedule/offerShiftRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/oFfErShIfTrEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/schedule/offerShiftRequests/offerShiftRequestIdValue",
			Expected: &MeJoinedTeamScheduleOfferShiftRequestId{
				TeamId:              "teamIdValue",
				OfferShiftRequestId: "offerShiftRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/schedule/offerShiftRequests/offerShiftRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/oFfErShIfTrEqUeStS/oFfErShIfTrEqUeStIdVaLuE",
			Expected: &MeJoinedTeamScheduleOfferShiftRequestId{
				TeamId:              "tEaMiDvAlUe",
				OfferShiftRequestId: "oFfErShIfTrEqUeStIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/oFfErShIfTrEqUeStS/oFfErShIfTrEqUeStIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamScheduleOfferShiftRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TeamId != v.Expected.TeamId {
			t.Fatalf("Expected %q but got %q for TeamId", v.Expected.TeamId, actual.TeamId)
		}

		if actual.OfferShiftRequestId != v.Expected.OfferShiftRequestId {
			t.Fatalf("Expected %q but got %q for OfferShiftRequestId", v.Expected.OfferShiftRequestId, actual.OfferShiftRequestId)
		}

	}
}

func TestSegmentsForMeJoinedTeamScheduleOfferShiftRequestId(t *testing.T) {
	segments := MeJoinedTeamScheduleOfferShiftRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamScheduleOfferShiftRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
