package mejoinedteamscheduletimesoff

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamScheduleTimesOffId{}

func TestNewMeJoinedTeamScheduleTimesOffID(t *testing.T) {
	id := NewMeJoinedTeamScheduleTimesOffID("teamIdValue", "timeOffIdValue")

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.TimeOffId != "timeOffIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TimeOffId'", id.TimeOffId, "timeOffIdValue")
	}
}

func TestFormatMeJoinedTeamScheduleTimesOffID(t *testing.T) {
	actual := NewMeJoinedTeamScheduleTimesOffID("teamIdValue", "timeOffIdValue").ID()
	expected := "/me/joinedTeams/teamIdValue/schedule/timesOff/timeOffIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamScheduleTimesOffID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamScheduleTimesOffId
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
			Input: "/me/joinedTeams/teamIdValue/schedule/timesOff",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/schedule/timesOff/timeOffIdValue",
			Expected: &MeJoinedTeamScheduleTimesOffId{
				TeamId:    "teamIdValue",
				TimeOffId: "timeOffIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/schedule/timesOff/timeOffIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamScheduleTimesOffID(v.Input)
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

func TestParseMeJoinedTeamScheduleTimesOffIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamScheduleTimesOffId
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
			Input: "/me/joinedTeams/teamIdValue/schedule/timesOff",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/tImEsOfF",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/schedule/timesOff/timeOffIdValue",
			Expected: &MeJoinedTeamScheduleTimesOffId{
				TeamId:    "teamIdValue",
				TimeOffId: "timeOffIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/schedule/timesOff/timeOffIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/tImEsOfF/tImEoFfIdVaLuE",
			Expected: &MeJoinedTeamScheduleTimesOffId{
				TeamId:    "tEaMiDvAlUe",
				TimeOffId: "tImEoFfIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/tImEsOfF/tImEoFfIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamScheduleTimesOffIDInsensitively(v.Input)
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

func TestSegmentsForMeJoinedTeamScheduleTimesOffId(t *testing.T) {
	segments := MeJoinedTeamScheduleTimesOffId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamScheduleTimesOffId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
