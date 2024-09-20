package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdChannelIdMemberId{}

func TestNewUserIdJoinedTeamIdChannelIdMemberID(t *testing.T) {
	id := NewUserIdJoinedTeamIdChannelIdMemberID("userId", "teamId", "channelId", "conversationMemberId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.ChannelId != "channelId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChannelId'", id.ChannelId, "channelId")
	}

	if id.ConversationMemberId != "conversationMemberId" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationMemberId'", id.ConversationMemberId, "conversationMemberId")
	}
}

func TestFormatUserIdJoinedTeamIdChannelIdMemberID(t *testing.T) {
	actual := NewUserIdJoinedTeamIdChannelIdMemberID("userId", "teamId", "channelId", "conversationMemberId").ID()
	expected := "/users/userId/joinedTeams/teamId/channels/channelId/members/conversationMemberId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdJoinedTeamIdChannelIdMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdChannelIdMemberId
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
			Input: "/users/userId/joinedTeams/teamId/channels",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/channels/channelId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/channels/channelId/members",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/joinedTeams/teamId/channels/channelId/members/conversationMemberId",
			Expected: &UserIdJoinedTeamIdChannelIdMemberId{
				UserId:               "userId",
				TeamId:               "teamId",
				ChannelId:            "channelId",
				ConversationMemberId: "conversationMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/joinedTeams/teamId/channels/channelId/members/conversationMemberId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdChannelIdMemberID(v.Input)
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

		if actual.ChannelId != v.Expected.ChannelId {
			t.Fatalf("Expected %q but got %q for ChannelId", v.Expected.ChannelId, actual.ChannelId)
		}

		if actual.ConversationMemberId != v.Expected.ConversationMemberId {
			t.Fatalf("Expected %q but got %q for ConversationMemberId", v.Expected.ConversationMemberId, actual.ConversationMemberId)
		}

	}
}

func TestParseUserIdJoinedTeamIdChannelIdMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdChannelIdMemberId
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
			Input: "/users/userId/joinedTeams/teamId/channels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/cHaNnElS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/channels/channelId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/cHaNnElS/cHaNnElId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/channels/channelId/members",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/cHaNnElS/cHaNnElId/mEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/joinedTeams/teamId/channels/channelId/members/conversationMemberId",
			Expected: &UserIdJoinedTeamIdChannelIdMemberId{
				UserId:               "userId",
				TeamId:               "teamId",
				ChannelId:            "channelId",
				ConversationMemberId: "conversationMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/joinedTeams/teamId/channels/channelId/members/conversationMemberId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/cHaNnElS/cHaNnElId/mEmBeRs/cOnVeRsAtIoNmEmBeRiD",
			Expected: &UserIdJoinedTeamIdChannelIdMemberId{
				UserId:               "uSeRiD",
				TeamId:               "tEaMiD",
				ChannelId:            "cHaNnElId",
				ConversationMemberId: "cOnVeRsAtIoNmEmBeRiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/cHaNnElS/cHaNnElId/mEmBeRs/cOnVeRsAtIoNmEmBeRiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdChannelIdMemberIDInsensitively(v.Input)
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

		if actual.ChannelId != v.Expected.ChannelId {
			t.Fatalf("Expected %q but got %q for ChannelId", v.Expected.ChannelId, actual.ChannelId)
		}

		if actual.ConversationMemberId != v.Expected.ConversationMemberId {
			t.Fatalf("Expected %q but got %q for ConversationMemberId", v.Expected.ConversationMemberId, actual.ConversationMemberId)
		}

	}
}

func TestSegmentsForUserIdJoinedTeamIdChannelIdMemberId(t *testing.T) {
	segments := UserIdJoinedTeamIdChannelIdMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdJoinedTeamIdChannelIdMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
