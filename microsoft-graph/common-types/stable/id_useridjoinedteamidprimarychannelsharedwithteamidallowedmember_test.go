package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{}

func TestNewUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID(t *testing.T) {
	id := NewUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID("userId", "teamId", "sharedWithChannelTeamInfoId", "conversationMemberId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.SharedWithChannelTeamInfoId != "sharedWithChannelTeamInfoId" {
		t.Fatalf("Expected %q but got %q for Segment 'SharedWithChannelTeamInfoId'", id.SharedWithChannelTeamInfoId, "sharedWithChannelTeamInfoId")
	}

	if id.ConversationMemberId != "conversationMemberId" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationMemberId'", id.ConversationMemberId, "conversationMemberId")
	}
}

func TestFormatUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID(t *testing.T) {
	actual := NewUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID("userId", "teamId", "sharedWithChannelTeamInfoId", "conversationMemberId").ID()
	expected := "/users/userId/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId/allowedMembers/conversationMemberId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId
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
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId/allowedMembers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId/allowedMembers/conversationMemberId",
			Expected: &UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{
				UserId:                      "userId",
				TeamId:                      "teamId",
				SharedWithChannelTeamInfoId: "sharedWithChannelTeamInfoId",
				ConversationMemberId:        "conversationMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId/allowedMembers/conversationMemberId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberID(v.Input)
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

		if actual.ConversationMemberId != v.Expected.ConversationMemberId {
			t.Fatalf("Expected %q but got %q for ConversationMemberId", v.Expected.ConversationMemberId, actual.ConversationMemberId)
		}

	}
}

func TestParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId
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
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId/allowedMembers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoId/aLlOwEdMeMbErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId/allowedMembers/conversationMemberId",
			Expected: &UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{
				UserId:                      "userId",
				TeamId:                      "teamId",
				SharedWithChannelTeamInfoId: "sharedWithChannelTeamInfoId",
				ConversationMemberId:        "conversationMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId/allowedMembers/conversationMemberId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoId/aLlOwEdMeMbErS/cOnVeRsAtIoNmEmBeRiD",
			Expected: &UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{
				UserId:                      "uSeRiD",
				TeamId:                      "tEaMiD",
				SharedWithChannelTeamInfoId: "sHaReDwItHcHaNnElTeAmInFoId",
				ConversationMemberId:        "cOnVeRsAtIoNmEmBeRiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoId/aLlOwEdMeMbErS/cOnVeRsAtIoNmEmBeRiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberIDInsensitively(v.Input)
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

		if actual.ConversationMemberId != v.Expected.ConversationMemberId {
			t.Fatalf("Expected %q but got %q for ConversationMemberId", v.Expected.ConversationMemberId, actual.ConversationMemberId)
		}

	}
}

func TestSegmentsForUserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId(t *testing.T) {
	segments := UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdJoinedTeamIdPrimaryChannelSharedWithTeamIdAllowedMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
