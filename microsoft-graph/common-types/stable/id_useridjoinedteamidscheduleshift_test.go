package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdScheduleShiftId{}

func TestNewUserIdJoinedTeamIdScheduleShiftID(t *testing.T) {
	id := NewUserIdJoinedTeamIdScheduleShiftID("userId", "teamId", "shiftId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.ShiftId != "shiftId" {
		t.Fatalf("Expected %q but got %q for Segment 'ShiftId'", id.ShiftId, "shiftId")
	}
}

func TestFormatUserIdJoinedTeamIdScheduleShiftID(t *testing.T) {
	actual := NewUserIdJoinedTeamIdScheduleShiftID("userId", "teamId", "shiftId").ID()
	expected := "/users/userId/joinedTeams/teamId/schedule/shifts/shiftId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdJoinedTeamIdScheduleShiftID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdScheduleShiftId
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
			Input: "/users/userId/joinedTeams/teamId/schedule/shifts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/joinedTeams/teamId/schedule/shifts/shiftId",
			Expected: &UserIdJoinedTeamIdScheduleShiftId{
				UserId:  "userId",
				TeamId:  "teamId",
				ShiftId: "shiftId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/joinedTeams/teamId/schedule/shifts/shiftId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdScheduleShiftID(v.Input)
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

		if actual.ShiftId != v.Expected.ShiftId {
			t.Fatalf("Expected %q but got %q for ShiftId", v.Expected.ShiftId, actual.ShiftId)
		}

	}
}

func TestParseUserIdJoinedTeamIdScheduleShiftIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdScheduleShiftId
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
			Input: "/users/userId/joinedTeams/teamId/schedule/shifts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/sChEdUlE/sHiFtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/joinedTeams/teamId/schedule/shifts/shiftId",
			Expected: &UserIdJoinedTeamIdScheduleShiftId{
				UserId:  "userId",
				TeamId:  "teamId",
				ShiftId: "shiftId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/joinedTeams/teamId/schedule/shifts/shiftId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/sChEdUlE/sHiFtS/sHiFtId",
			Expected: &UserIdJoinedTeamIdScheduleShiftId{
				UserId:  "uSeRiD",
				TeamId:  "tEaMiD",
				ShiftId: "sHiFtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/sChEdUlE/sHiFtS/sHiFtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdScheduleShiftIDInsensitively(v.Input)
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

		if actual.ShiftId != v.Expected.ShiftId {
			t.Fatalf("Expected %q but got %q for ShiftId", v.Expected.ShiftId, actual.ShiftId)
		}

	}
}

func TestSegmentsForUserIdJoinedTeamIdScheduleShiftId(t *testing.T) {
	segments := UserIdJoinedTeamIdScheduleShiftId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdJoinedTeamIdScheduleShiftId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
