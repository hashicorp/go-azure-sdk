package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdPrimaryChannelMessageIdReplyId{}

func TestNewMeJoinedTeamIdPrimaryChannelMessageIdReplyID(t *testing.T) {
	id := NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID("teamId", "chatMessageId", "chatMessageId1")

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.ChatMessageId != "chatMessageId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId'", id.ChatMessageId, "chatMessageId")
	}

	if id.ChatMessageId1 != "chatMessageId1" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId1'", id.ChatMessageId1, "chatMessageId1")
	}
}

func TestFormatMeJoinedTeamIdPrimaryChannelMessageIdReplyID(t *testing.T) {
	actual := NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID("teamId", "chatMessageId", "chatMessageId1").ID()
	expected := "/me/joinedTeams/teamId/primaryChannel/messages/chatMessageId/replies/chatMessageId1"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamIdPrimaryChannelMessageIdReplyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdPrimaryChannelMessageIdReplyId
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
			Input: "/me/joinedTeams/teamId/primaryChannel",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/primaryChannel/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/primaryChannel/messages/chatMessageId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/primaryChannel/messages/chatMessageId/replies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/primaryChannel/messages/chatMessageId/replies/chatMessageId1",
			Expected: &MeJoinedTeamIdPrimaryChannelMessageIdReplyId{
				TeamId:         "teamId",
				ChatMessageId:  "chatMessageId",
				ChatMessageId1: "chatMessageId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/primaryChannel/messages/chatMessageId/replies/chatMessageId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdPrimaryChannelMessageIdReplyID(v.Input)
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

		if actual.ChatMessageId != v.Expected.ChatMessageId {
			t.Fatalf("Expected %q but got %q for ChatMessageId", v.Expected.ChatMessageId, actual.ChatMessageId)
		}

		if actual.ChatMessageId1 != v.Expected.ChatMessageId1 {
			t.Fatalf("Expected %q but got %q for ChatMessageId1", v.Expected.ChatMessageId1, actual.ChatMessageId1)
		}

	}
}

func TestParseMeJoinedTeamIdPrimaryChannelMessageIdReplyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdPrimaryChannelMessageIdReplyId
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
			Input: "/me/joinedTeams/teamId/primaryChannel",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/primaryChannel/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/primaryChannel/messages/chatMessageId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/primaryChannel/messages/chatMessageId/replies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeId/rEpLiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/primaryChannel/messages/chatMessageId/replies/chatMessageId1",
			Expected: &MeJoinedTeamIdPrimaryChannelMessageIdReplyId{
				TeamId:         "teamId",
				ChatMessageId:  "chatMessageId",
				ChatMessageId1: "chatMessageId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/primaryChannel/messages/chatMessageId/replies/chatMessageId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeId/rEpLiEs/cHaTmEsSaGeId1",
			Expected: &MeJoinedTeamIdPrimaryChannelMessageIdReplyId{
				TeamId:         "tEaMiD",
				ChatMessageId:  "cHaTmEsSaGeId",
				ChatMessageId1: "cHaTmEsSaGeId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeId/rEpLiEs/cHaTmEsSaGeId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdPrimaryChannelMessageIdReplyIDInsensitively(v.Input)
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

		if actual.ChatMessageId != v.Expected.ChatMessageId {
			t.Fatalf("Expected %q but got %q for ChatMessageId", v.Expected.ChatMessageId, actual.ChatMessageId)
		}

		if actual.ChatMessageId1 != v.Expected.ChatMessageId1 {
			t.Fatalf("Expected %q but got %q for ChatMessageId1", v.Expected.ChatMessageId1, actual.ChatMessageId1)
		}

	}
}

func TestSegmentsForMeJoinedTeamIdPrimaryChannelMessageIdReplyId(t *testing.T) {
	segments := MeJoinedTeamIdPrimaryChannelMessageIdReplyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamIdPrimaryChannelMessageIdReplyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
