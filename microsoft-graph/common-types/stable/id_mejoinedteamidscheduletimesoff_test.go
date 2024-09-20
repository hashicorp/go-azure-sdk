package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdScheduleTimesOffId{}

func TestNewMeJoinedTeamIdScheduleTimesOffID(t *testing.T) {
	id := NewMeJoinedTeamIdScheduleTimesOffID("teamId", "timeOffId")

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.TimeOffId != "timeOffId" {
		t.Fatalf("Expected %q but got %q for Segment 'TimeOffId'", id.TimeOffId, "timeOffId")
	}
}

func TestFormatMeJoinedTeamIdScheduleTimesOffID(t *testing.T) {
	actual := NewMeJoinedTeamIdScheduleTimesOffID("teamId", "timeOffId").ID()
	expected := "/me/joinedTeams/teamId/schedule/timesOff/timeOffId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamIdScheduleTimesOffID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdScheduleTimesOffId
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
			Input: "/me/joinedTeams/teamId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/schedule",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/schedule/timesOff",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/schedule/timesOff/timeOffId",
			Expected: &MeJoinedTeamIdScheduleTimesOffId{
				TeamId:    "teamId",
				TimeOffId: "timeOffId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/schedule/timesOff/timeOffId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdScheduleTimesOffID(v.Input)
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

		if actual.TimeOffId != v.Expected.TimeOffId {
			t.Fatalf("Expected %q but got %q for TimeOffId", v.Expected.TimeOffId, actual.TimeOffId)
		}

	}
}

func TestParseMeJoinedTeamIdScheduleTimesOffIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdScheduleTimesOffId
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
			Input: "/me/joinedTeams/teamId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/schedule",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/schedule/timesOff",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE/tImEsOfF",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/schedule/timesOff/timeOffId",
			Expected: &MeJoinedTeamIdScheduleTimesOffId{
				TeamId:    "teamId",
				TimeOffId: "timeOffId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/schedule/timesOff/timeOffId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE/tImEsOfF/tImEoFfId",
			Expected: &MeJoinedTeamIdScheduleTimesOffId{
				TeamId:    "tEaMiD",
				TimeOffId: "tImEoFfId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE/tImEsOfF/tImEoFfId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdScheduleTimesOffIDInsensitively(v.Input)
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

		if actual.TimeOffId != v.Expected.TimeOffId {
			t.Fatalf("Expected %q but got %q for TimeOffId", v.Expected.TimeOffId, actual.TimeOffId)
		}

	}
}

func TestSegmentsForMeJoinedTeamIdScheduleTimesOffId(t *testing.T) {
	segments := MeJoinedTeamIdScheduleTimesOffId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamIdScheduleTimesOffId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
