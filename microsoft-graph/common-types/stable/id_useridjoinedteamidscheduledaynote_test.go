package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdScheduleDayNoteId{}

func TestNewUserIdJoinedTeamIdScheduleDayNoteID(t *testing.T) {
	id := NewUserIdJoinedTeamIdScheduleDayNoteID("userId", "teamId", "dayNoteId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.DayNoteId != "dayNoteId" {
		t.Fatalf("Expected %q but got %q for Segment 'DayNoteId'", id.DayNoteId, "dayNoteId")
	}
}

func TestFormatUserIdJoinedTeamIdScheduleDayNoteID(t *testing.T) {
	actual := NewUserIdJoinedTeamIdScheduleDayNoteID("userId", "teamId", "dayNoteId").ID()
	expected := "/users/userId/joinedTeams/teamId/schedule/dayNotes/dayNoteId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdJoinedTeamIdScheduleDayNoteID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdScheduleDayNoteId
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
			Input: "/users/userId/joinedTeams",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/schedule",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/schedule/dayNotes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/joinedTeams/teamId/schedule/dayNotes/dayNoteId",
			Expected: &UserIdJoinedTeamIdScheduleDayNoteId{
				UserId:    "userId",
				TeamId:    "teamId",
				DayNoteId: "dayNoteId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/joinedTeams/teamId/schedule/dayNotes/dayNoteId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdScheduleDayNoteID(v.Input)
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

		if actual.DayNoteId != v.Expected.DayNoteId {
			t.Fatalf("Expected %q but got %q for DayNoteId", v.Expected.DayNoteId, actual.DayNoteId)
		}

	}
}

func TestParseUserIdJoinedTeamIdScheduleDayNoteIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdScheduleDayNoteId
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
			Input: "/users/userId/joinedTeams",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/schedule",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/sChEdUlE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/schedule/dayNotes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/sChEdUlE/dAyNoTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/joinedTeams/teamId/schedule/dayNotes/dayNoteId",
			Expected: &UserIdJoinedTeamIdScheduleDayNoteId{
				UserId:    "userId",
				TeamId:    "teamId",
				DayNoteId: "dayNoteId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/joinedTeams/teamId/schedule/dayNotes/dayNoteId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/sChEdUlE/dAyNoTeS/dAyNoTeId",
			Expected: &UserIdJoinedTeamIdScheduleDayNoteId{
				UserId:    "uSeRiD",
				TeamId:    "tEaMiD",
				DayNoteId: "dAyNoTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/sChEdUlE/dAyNoTeS/dAyNoTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdScheduleDayNoteIDInsensitively(v.Input)
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

		if actual.DayNoteId != v.Expected.DayNoteId {
			t.Fatalf("Expected %q but got %q for DayNoteId", v.Expected.DayNoteId, actual.DayNoteId)
		}

	}
}

func TestSegmentsForUserIdJoinedTeamIdScheduleDayNoteId(t *testing.T) {
	segments := UserIdJoinedTeamIdScheduleDayNoteId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdJoinedTeamIdScheduleDayNoteId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
