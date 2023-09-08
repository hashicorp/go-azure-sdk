package userjoinedteamscheduletimeoffreason

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamScheduleTimeOffReasonId{}

func TestNewUserJoinedTeamScheduleTimeOffReasonID(t *testing.T) {
	id := NewUserJoinedTeamScheduleTimeOffReasonID("userIdValue", "teamIdValue", "timeOffReasonIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.TimeOffReasonId != "timeOffReasonIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TimeOffReasonId'", id.TimeOffReasonId, "timeOffReasonIdValue")
	}
}

func TestFormatUserJoinedTeamScheduleTimeOffReasonID(t *testing.T) {
	actual := NewUserJoinedTeamScheduleTimeOffReasonID("userIdValue", "teamIdValue", "timeOffReasonIdValue").ID()
	expected := "/users/userIdValue/joinedTeams/teamIdValue/schedule/timeOffReasons/timeOffReasonIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserJoinedTeamScheduleTimeOffReasonID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserJoinedTeamScheduleTimeOffReasonId
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/timeOffReasons",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/timeOffReasons/timeOffReasonIdValue",
			Expected: &UserJoinedTeamScheduleTimeOffReasonId{
				UserId:          "userIdValue",
				TeamId:          "teamIdValue",
				TimeOffReasonId: "timeOffReasonIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/timeOffReasons/timeOffReasonIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserJoinedTeamScheduleTimeOffReasonID(v.Input)
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

		if actual.TeamId != v.Expected.TeamId {
			t.Fatalf("Expected %q but got %q for TeamId", v.Expected.TeamId, actual.TeamId)
		}

		if actual.TimeOffReasonId != v.Expected.TimeOffReasonId {
			t.Fatalf("Expected %q but got %q for TimeOffReasonId", v.Expected.TimeOffReasonId, actual.TimeOffReasonId)
		}

	}
}

func TestParseUserJoinedTeamScheduleTimeOffReasonIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserJoinedTeamScheduleTimeOffReasonId
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/timeOffReasons",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/tImEoFfReAsOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/timeOffReasons/timeOffReasonIdValue",
			Expected: &UserJoinedTeamScheduleTimeOffReasonId{
				UserId:          "userIdValue",
				TeamId:          "teamIdValue",
				TimeOffReasonId: "timeOffReasonIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/timeOffReasons/timeOffReasonIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/tImEoFfReAsOnS/tImEoFfReAsOnIdVaLuE",
			Expected: &UserJoinedTeamScheduleTimeOffReasonId{
				UserId:          "uSeRiDvAlUe",
				TeamId:          "tEaMiDvAlUe",
				TimeOffReasonId: "tImEoFfReAsOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/tImEoFfReAsOnS/tImEoFfReAsOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserJoinedTeamScheduleTimeOffReasonIDInsensitively(v.Input)
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

		if actual.TeamId != v.Expected.TeamId {
			t.Fatalf("Expected %q but got %q for TeamId", v.Expected.TeamId, actual.TeamId)
		}

		if actual.TimeOffReasonId != v.Expected.TimeOffReasonId {
			t.Fatalf("Expected %q but got %q for TimeOffReasonId", v.Expected.TimeOffReasonId, actual.TimeOffReasonId)
		}

	}
}

func TestSegmentsForUserJoinedTeamScheduleTimeOffReasonId(t *testing.T) {
	segments := UserJoinedTeamScheduleTimeOffReasonId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserJoinedTeamScheduleTimeOffReasonId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
