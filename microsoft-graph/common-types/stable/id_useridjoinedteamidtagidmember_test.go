package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdTagIdMemberId{}

func TestNewUserIdJoinedTeamIdTagIdMemberID(t *testing.T) {
	id := NewUserIdJoinedTeamIdTagIdMemberID("userId", "teamId", "teamworkTagId", "teamworkTagMemberId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.TeamworkTagId != "teamworkTagId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamworkTagId'", id.TeamworkTagId, "teamworkTagId")
	}

	if id.TeamworkTagMemberId != "teamworkTagMemberId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamworkTagMemberId'", id.TeamworkTagMemberId, "teamworkTagMemberId")
	}
}

func TestFormatUserIdJoinedTeamIdTagIdMemberID(t *testing.T) {
	actual := NewUserIdJoinedTeamIdTagIdMemberID("userId", "teamId", "teamworkTagId", "teamworkTagMemberId").ID()
	expected := "/users/userId/joinedTeams/teamId/tags/teamworkTagId/members/teamworkTagMemberId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdJoinedTeamIdTagIdMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdTagIdMemberId
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
			Input: "/users/userId/joinedTeams/teamId/tags",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/tags/teamworkTagId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/tags/teamworkTagId/members",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/joinedTeams/teamId/tags/teamworkTagId/members/teamworkTagMemberId",
			Expected: &UserIdJoinedTeamIdTagIdMemberId{
				UserId:              "userId",
				TeamId:              "teamId",
				TeamworkTagId:       "teamworkTagId",
				TeamworkTagMemberId: "teamworkTagMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/joinedTeams/teamId/tags/teamworkTagId/members/teamworkTagMemberId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdTagIdMemberID(v.Input)
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

		if actual.TeamworkTagId != v.Expected.TeamworkTagId {
			t.Fatalf("Expected %q but got %q for TeamworkTagId", v.Expected.TeamworkTagId, actual.TeamworkTagId)
		}

		if actual.TeamworkTagMemberId != v.Expected.TeamworkTagMemberId {
			t.Fatalf("Expected %q but got %q for TeamworkTagMemberId", v.Expected.TeamworkTagMemberId, actual.TeamworkTagMemberId)
		}

	}
}

func TestParseUserIdJoinedTeamIdTagIdMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdTagIdMemberId
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
			Input: "/users/userId/joinedTeams/teamId/tags",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/tAgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/tags/teamworkTagId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/tAgS/tEaMwOrKtAgId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/tags/teamworkTagId/members",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/tAgS/tEaMwOrKtAgId/mEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/joinedTeams/teamId/tags/teamworkTagId/members/teamworkTagMemberId",
			Expected: &UserIdJoinedTeamIdTagIdMemberId{
				UserId:              "userId",
				TeamId:              "teamId",
				TeamworkTagId:       "teamworkTagId",
				TeamworkTagMemberId: "teamworkTagMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/joinedTeams/teamId/tags/teamworkTagId/members/teamworkTagMemberId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/tAgS/tEaMwOrKtAgId/mEmBeRs/tEaMwOrKtAgMeMbErId",
			Expected: &UserIdJoinedTeamIdTagIdMemberId{
				UserId:              "uSeRiD",
				TeamId:              "tEaMiD",
				TeamworkTagId:       "tEaMwOrKtAgId",
				TeamworkTagMemberId: "tEaMwOrKtAgMeMbErId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/tAgS/tEaMwOrKtAgId/mEmBeRs/tEaMwOrKtAgMeMbErId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdTagIdMemberIDInsensitively(v.Input)
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

		if actual.TeamworkTagId != v.Expected.TeamworkTagId {
			t.Fatalf("Expected %q but got %q for TeamworkTagId", v.Expected.TeamworkTagId, actual.TeamworkTagId)
		}

		if actual.TeamworkTagMemberId != v.Expected.TeamworkTagMemberId {
			t.Fatalf("Expected %q but got %q for TeamworkTagMemberId", v.Expected.TeamworkTagMemberId, actual.TeamworkTagMemberId)
		}

	}
}

func TestSegmentsForUserIdJoinedTeamIdTagIdMemberId(t *testing.T) {
	segments := UserIdJoinedTeamIdTagIdMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdJoinedTeamIdTagIdMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
