package mejoinedteamchannelmessagereply

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamChannelMessageId{}

func TestNewMeJoinedTeamChannelMessageID(t *testing.T) {
	id := NewMeJoinedTeamChannelMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.ChannelId != "channelIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChannelId'", id.ChannelId, "channelIdValue")
	}

	if id.ChatMessageId != "chatMessageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId'", id.ChatMessageId, "chatMessageIdValue")
	}
}

func TestFormatMeJoinedTeamChannelMessageID(t *testing.T) {
	actual := NewMeJoinedTeamChannelMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue").ID()
	expected := "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamChannelMessageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamChannelMessageId
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
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue",
			Expected: &MeJoinedTeamChannelMessageId{
				TeamId:        "teamIdValue",
				ChannelId:     "channelIdValue",
				ChatMessageId: "chatMessageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamChannelMessageID(v.Input)
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

	}
}

func TestParseMeJoinedTeamChannelMessageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamChannelMessageId
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
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue",
			Expected: &MeJoinedTeamChannelMessageId{
				TeamId:        "teamIdValue",
				ChannelId:     "channelIdValue",
				ChatMessageId: "chatMessageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/messages/chatMessageIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/cHaNnElS/cHaNnElIdVaLuE/mEsSaGeS/cHaTmEsSaGeIdVaLuE",
			Expected: &MeJoinedTeamChannelMessageId{
				TeamId:        "tEaMiDvAlUe",
				ChannelId:     "cHaNnElIdVaLuE",
				ChatMessageId: "cHaTmEsSaGeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/cHaNnElS/cHaNnElIdVaLuE/mEsSaGeS/cHaTmEsSaGeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamChannelMessageIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeJoinedTeamChannelMessageId(t *testing.T) {
	segments := MeJoinedTeamChannelMessageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamChannelMessageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
