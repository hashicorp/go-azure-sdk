package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId{}

func TestNewUserIdJoinedTeamIdPrimaryChannelSharedWithTeamID(t *testing.T) {
	id := NewUserIdJoinedTeamIdPrimaryChannelSharedWithTeamID("userId", "teamId", "sharedWithChannelTeamInfoId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.SharedWithChannelTeamInfoId != "sharedWithChannelTeamInfoId" {
		t.Fatalf("Expected %q but got %q for Segment 'SharedWithChannelTeamInfoId'", id.SharedWithChannelTeamInfoId, "sharedWithChannelTeamInfoId")
	}
}

func TestFormatUserIdJoinedTeamIdPrimaryChannelSharedWithTeamID(t *testing.T) {
	actual := NewUserIdJoinedTeamIdPrimaryChannelSharedWithTeamID("userId", "teamId", "sharedWithChannelTeamInfoId").ID()
	expected := "/users/userId/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId
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
			Input: "/users/userId/joinedTeams/teamId/primaryChannel",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/primaryChannel/sharedWithTeams",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId",
			Expected: &UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId{
				UserId:                      "userId",
				TeamId:                      "teamId",
				SharedWithChannelTeamInfoId: "sharedWithChannelTeamInfoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamID(v.Input)
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

		if actual.SharedWithChannelTeamInfoId != v.Expected.SharedWithChannelTeamInfoId {
			t.Fatalf("Expected %q but got %q for SharedWithChannelTeamInfoId", v.Expected.SharedWithChannelTeamInfoId, actual.SharedWithChannelTeamInfoId)
		}

	}
}

func TestParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId
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
			Input: "/users/userId/joinedTeams/teamId/primaryChannel",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/primaryChannel/sharedWithTeams",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/sHaReDwItHtEaMs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId",
			Expected: &UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId{
				UserId:                      "userId",
				TeamId:                      "teamId",
				SharedWithChannelTeamInfoId: "sharedWithChannelTeamInfoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoId",
			Expected: &UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId{
				UserId:                      "uSeRiD",
				TeamId:                      "tEaMiD",
				SharedWithChannelTeamInfoId: "sHaReDwItHcHaNnElTeAmInFoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIDInsensitively(v.Input)
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

		if actual.SharedWithChannelTeamInfoId != v.Expected.SharedWithChannelTeamInfoId {
			t.Fatalf("Expected %q but got %q for SharedWithChannelTeamInfoId", v.Expected.SharedWithChannelTeamInfoId, actual.SharedWithChannelTeamInfoId)
		}

	}
}

func TestSegmentsForUserIdJoinedTeamIdPrimaryChannelSharedWithTeamId(t *testing.T) {
	segments := UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
