package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId{}

func TestNewMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentID(t *testing.T) {
	id := NewMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value", "chatMessageHostedContentIdValue")

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.ChannelId != "channelIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChannelId'", id.ChannelId, "channelIdValue")
	}

	if id.ChatMessageId != "chatMessageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId'", id.ChatMessageId, "chatMessageIdValue")
	}

	if id.ChatMessageId1 != "chatMessageId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId1'", id.ChatMessageId1, "chatMessageId1Value")
	}

	if id.ChatMessageHostedContentId != "chatMessageHostedContentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageHostedContentId'", id.ChatMessageHostedContentId, "chatMessageHostedContentIdValue")
	}
}

func TestFormatMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentID(t *testing.T) {
	actual := NewMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value", "chatMessageHostedContentIdValue").ID()
	expected := "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents/chatMessageHostedContentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId
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
			Input: "/me/joinedTeams/teamIdValue/channels",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue/replies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents/chatMessageHostedContentIdValue",
			Expected: &MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId{
				TeamId:                     "teamIdValue",
				ChannelId:                  "channelIdValue",
				ChatMessageId:              "chatMessageIdValue",
				ChatMessageId1:             "chatMessageId1Value",
				ChatMessageHostedContentId: "chatMessageHostedContentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents/chatMessageHostedContentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentID(v.Input)
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

		if actual.ChatMessageId1 != v.Expected.ChatMessageId1 {
			t.Fatalf("Expected %q but got %q for ChatMessageId1", v.Expected.ChatMessageId1, actual.ChatMessageId1)
		}

		if actual.ChatMessageHostedContentId != v.Expected.ChatMessageHostedContentId {
			t.Fatalf("Expected %q but got %q for ChatMessageHostedContentId", v.Expected.ChatMessageHostedContentId, actual.ChatMessageHostedContentId)
		}

	}
}

func TestParseMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId
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
			Input: "/me/joinedTeams/teamIdValue/channels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/cHaNnElS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/cHaNnElS/cHaNnElIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/cHaNnElS/cHaNnElIdVaLuE/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/cHaNnElS/cHaNnElIdVaLuE/mEsSaGeS/cHaTmEsSaGeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue/replies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/cHaNnElS/cHaNnElIdVaLuE/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/cHaNnElS/cHaNnElIdVaLuE/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/cHaNnElS/cHaNnElIdVaLuE/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe/hOsTeDcOnTeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents/chatMessageHostedContentIdValue",
			Expected: &MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId{
				TeamId:                     "teamIdValue",
				ChannelId:                  "channelIdValue",
				ChatMessageId:              "chatMessageIdValue",
				ChatMessageId1:             "chatMessageId1Value",
				ChatMessageHostedContentId: "chatMessageHostedContentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents/chatMessageHostedContentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/cHaNnElS/cHaNnElIdVaLuE/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe/hOsTeDcOnTeNtS/cHaTmEsSaGeHoStEdCoNtEnTiDvAlUe",
			Expected: &MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId{
				TeamId:                     "tEaMiDvAlUe",
				ChannelId:                  "cHaNnElIdVaLuE",
				ChatMessageId:              "cHaTmEsSaGeIdVaLuE",
				ChatMessageId1:             "cHaTmEsSaGeId1vAlUe",
				ChatMessageHostedContentId: "cHaTmEsSaGeHoStEdCoNtEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/cHaNnElS/cHaNnElIdVaLuE/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe/hOsTeDcOnTeNtS/cHaTmEsSaGeHoStEdCoNtEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentIDInsensitively(v.Input)
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

		if actual.ChatMessageId1 != v.Expected.ChatMessageId1 {
			t.Fatalf("Expected %q but got %q for ChatMessageId1", v.Expected.ChatMessageId1, actual.ChatMessageId1)
		}

		if actual.ChatMessageHostedContentId != v.Expected.ChatMessageHostedContentId {
			t.Fatalf("Expected %q but got %q for ChatMessageHostedContentId", v.Expected.ChatMessageHostedContentId, actual.ChatMessageHostedContentId)
		}

	}
}

func TestSegmentsForMeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId(t *testing.T) {
	segments := MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamIdChannelIdMessageIdReplyIdHostedContentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
