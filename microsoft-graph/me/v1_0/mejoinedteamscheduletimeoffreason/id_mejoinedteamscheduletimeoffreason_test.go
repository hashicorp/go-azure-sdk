package mejoinedteamscheduletimeoffreason

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamScheduleTimeOffReasonId{}

func TestNewMeJoinedTeamScheduleTimeOffReasonID(t *testing.T) {
	id := NewMeJoinedTeamScheduleTimeOffReasonID("teamIdValue", "timeOffReasonIdValue")

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.TimeOffReasonId != "timeOffReasonIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TimeOffReasonId'", id.TimeOffReasonId, "timeOffReasonIdValue")
	}
}

func TestFormatMeJoinedTeamScheduleTimeOffReasonID(t *testing.T) {
	actual := NewMeJoinedTeamScheduleTimeOffReasonID("teamIdValue", "timeOffReasonIdValue").ID()
	expected := "/me/joinedTeams/teamIdValue/schedule/timeOffReasons/timeOffReasonIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamScheduleTimeOffReasonID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamScheduleTimeOffReasonId
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
			Input: "/me/joinedTeams/teamIdValue/schedule/timeOffReasons",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/schedule/timeOffReasons/timeOffReasonIdValue",
			Expected: &MeJoinedTeamScheduleTimeOffReasonId{
				TeamId:          "teamIdValue",
				TimeOffReasonId: "timeOffReasonIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/schedule/timeOffReasons/timeOffReasonIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamScheduleTimeOffReasonID(v.Input)
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

		if actual.TimeOffReasonId != v.Expected.TimeOffReasonId {
			t.Fatalf("Expected %q but got %q for TimeOffReasonId", v.Expected.TimeOffReasonId, actual.TimeOffReasonId)
		}

	}
}

func TestParseMeJoinedTeamScheduleTimeOffReasonIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamScheduleTimeOffReasonId
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
			Input: "/me/joinedTeams/teamIdValue/schedule/timeOffReasons",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/tImEoFfReAsOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/schedule/timeOffReasons/timeOffReasonIdValue",
			Expected: &MeJoinedTeamScheduleTimeOffReasonId{
				TeamId:          "teamIdValue",
				TimeOffReasonId: "timeOffReasonIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/schedule/timeOffReasons/timeOffReasonIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/tImEoFfReAsOnS/tImEoFfReAsOnIdVaLuE",
			Expected: &MeJoinedTeamScheduleTimeOffReasonId{
				TeamId:          "tEaMiDvAlUe",
				TimeOffReasonId: "tImEoFfReAsOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/tImEoFfReAsOnS/tImEoFfReAsOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamScheduleTimeOffReasonIDInsensitively(v.Input)
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

		if actual.TimeOffReasonId != v.Expected.TimeOffReasonId {
			t.Fatalf("Expected %q but got %q for TimeOffReasonId", v.Expected.TimeOffReasonId, actual.TimeOffReasonId)
		}

	}
}

func TestSegmentsForMeJoinedTeamScheduleTimeOffReasonId(t *testing.T) {
	segments := MeJoinedTeamScheduleTimeOffReasonId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamScheduleTimeOffReasonId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
