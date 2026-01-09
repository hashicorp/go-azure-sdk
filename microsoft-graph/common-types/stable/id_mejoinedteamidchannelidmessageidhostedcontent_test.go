package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdChannelIdMessageIdHostedContentId{}

func TestNewMeJoinedTeamIdChannelIdMessageIdHostedContentID(t *testing.T) {
	id := NewMeJoinedTeamIdChannelIdMessageIdHostedContentID("teamId", "channelId", "chatMessageId", "chatMessageHostedContentId")

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.ChannelId != "channelId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChannelId'", id.ChannelId, "channelId")
	}

	if id.ChatMessageId != "chatMessageId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId'", id.ChatMessageId, "chatMessageId")
	}

	if id.ChatMessageHostedContentId != "chatMessageHostedContentId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageHostedContentId'", id.ChatMessageHostedContentId, "chatMessageHostedContentId")
	}
}

func TestFormatMeJoinedTeamIdChannelIdMessageIdHostedContentID(t *testing.T) {
	actual := NewMeJoinedTeamIdChannelIdMessageIdHostedContentID("teamId", "channelId", "chatMessageId", "chatMessageHostedContentId").ID()
	expected := "/me/joinedTeams/teamId/channels/channelId/messages/chatMessageId/hostedContents/chatMessageHostedContentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamIdChannelIdMessageIdHostedContentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdChannelIdMessageIdHostedContentId
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
			Input: "/me/joinedTeams/teamId/channels",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/channels/channelId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/channels/channelId/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/channels/channelId/messages/chatMessageId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/channels/channelId/messages/chatMessageId/hostedContents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/channels/channelId/messages/chatMessageId/hostedContents/chatMessageHostedContentId",
			Expected: &MeJoinedTeamIdChannelIdMessageIdHostedContentId{
				TeamId:                     "teamId",
				ChannelId:                  "channelId",
				ChatMessageId:              "chatMessageId",
				ChatMessageHostedContentId: "chatMessageHostedContentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/channels/channelId/messages/chatMessageId/hostedContents/chatMessageHostedContentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdChannelIdMessageIdHostedContentID(v.Input)
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

		if actual.ChannelId != v.Expected.ChannelId {
			t.Fatalf("Expected %q but got %q for ChannelId", v.Expected.ChannelId, actual.ChannelId)
		}

		if actual.ChatMessageId != v.Expected.ChatMessageId {
			t.Fatalf("Expected %q but got %q for ChatMessageId", v.Expected.ChatMessageId, actual.ChatMessageId)
		}

		if actual.ChatMessageHostedContentId != v.Expected.ChatMessageHostedContentId {
			t.Fatalf("Expected %q but got %q for ChatMessageHostedContentId", v.Expected.ChatMessageHostedContentId, actual.ChatMessageHostedContentId)
		}

	}
}

func TestParseMeJoinedTeamIdChannelIdMessageIdHostedContentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdChannelIdMessageIdHostedContentId
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
			Input: "/me/joinedTeams/teamId/channels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/cHaNnElS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/channels/channelId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/cHaNnElS/cHaNnElId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/channels/channelId/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/cHaNnElS/cHaNnElId/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/channels/channelId/messages/chatMessageId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/cHaNnElS/cHaNnElId/mEsSaGeS/cHaTmEsSaGeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/channels/channelId/messages/chatMessageId/hostedContents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/cHaNnElS/cHaNnElId/mEsSaGeS/cHaTmEsSaGeId/hOsTeDcOnTeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/channels/channelId/messages/chatMessageId/hostedContents/chatMessageHostedContentId",
			Expected: &MeJoinedTeamIdChannelIdMessageIdHostedContentId{
				TeamId:                     "teamId",
				ChannelId:                  "channelId",
				ChatMessageId:              "chatMessageId",
				ChatMessageHostedContentId: "chatMessageHostedContentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/channels/channelId/messages/chatMessageId/hostedContents/chatMessageHostedContentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/cHaNnElS/cHaNnElId/mEsSaGeS/cHaTmEsSaGeId/hOsTeDcOnTeNtS/cHaTmEsSaGeHoStEdCoNtEnTiD",
			Expected: &MeJoinedTeamIdChannelIdMessageIdHostedContentId{
				TeamId:                     "tEaMiD",
				ChannelId:                  "cHaNnElId",
				ChatMessageId:              "cHaTmEsSaGeId",
				ChatMessageHostedContentId: "cHaTmEsSaGeHoStEdCoNtEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/cHaNnElS/cHaNnElId/mEsSaGeS/cHaTmEsSaGeId/hOsTeDcOnTeNtS/cHaTmEsSaGeHoStEdCoNtEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdChannelIdMessageIdHostedContentIDInsensitively(v.Input)
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

		if actual.ChannelId != v.Expected.ChannelId {
			t.Fatalf("Expected %q but got %q for ChannelId", v.Expected.ChannelId, actual.ChannelId)
		}

		if actual.ChatMessageId != v.Expected.ChatMessageId {
			t.Fatalf("Expected %q but got %q for ChatMessageId", v.Expected.ChatMessageId, actual.ChatMessageId)
		}

		if actual.ChatMessageHostedContentId != v.Expected.ChatMessageHostedContentId {
			t.Fatalf("Expected %q but got %q for ChatMessageHostedContentId", v.Expected.ChatMessageHostedContentId, actual.ChatMessageHostedContentId)
		}

	}
}

func TestSegmentsForMeJoinedTeamIdChannelIdMessageIdHostedContentId(t *testing.T) {
	segments := MeJoinedTeamIdChannelIdMessageIdHostedContentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamIdChannelIdMessageIdHostedContentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
