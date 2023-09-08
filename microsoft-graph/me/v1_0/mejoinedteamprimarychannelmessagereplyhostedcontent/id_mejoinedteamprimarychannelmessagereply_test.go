package mejoinedteamprimarychannelmessagereplyhostedcontent

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamPrimaryChannelMessageReplyId{}

func TestNewMeJoinedTeamPrimaryChannelMessageReplyID(t *testing.T) {
	id := NewMeJoinedTeamPrimaryChannelMessageReplyID("teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.ChatMessageId != "chatMessageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId'", id.ChatMessageId, "chatMessageIdValue")
	}

	if id.ChatMessageId1 != "chatMessageId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId1'", id.ChatMessageId1, "chatMessageId1Value")
	}
}

func TestFormatMeJoinedTeamPrimaryChannelMessageReplyID(t *testing.T) {
	actual := NewMeJoinedTeamPrimaryChannelMessageReplyID("teamIdValue", "chatMessageIdValue", "chatMessageId1Value").ID()
	expected := "/me/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamPrimaryChannelMessageReplyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamPrimaryChannelMessageReplyId
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
			Input: "/me/joinedTeams/teamIdValue/primaryChannel",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/primaryChannel/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value",
			Expected: &MeJoinedTeamPrimaryChannelMessageReplyId{
				TeamId:         "teamIdValue",
				ChatMessageId:  "chatMessageIdValue",
				ChatMessageId1: "chatMessageId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamPrimaryChannelMessageReplyID(v.Input)
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

func TestParseMeJoinedTeamPrimaryChannelMessageReplyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamPrimaryChannelMessageReplyId
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
			Input: "/me/joinedTeams/teamIdValue/primaryChannel",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/primaryChannel/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value",
			Expected: &MeJoinedTeamPrimaryChannelMessageReplyId{
				TeamId:         "teamIdValue",
				ChatMessageId:  "chatMessageIdValue",
				ChatMessageId1: "chatMessageId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe",
			Expected: &MeJoinedTeamPrimaryChannelMessageReplyId{
				TeamId:         "tEaMiDvAlUe",
				ChatMessageId:  "cHaTmEsSaGeIdVaLuE",
				ChatMessageId1: "cHaTmEsSaGeId1vAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamPrimaryChannelMessageReplyIDInsensitively(v.Input)
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

func TestSegmentsForMeJoinedTeamPrimaryChannelMessageReplyId(t *testing.T) {
	segments := MeJoinedTeamPrimaryChannelMessageReplyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamPrimaryChannelMessageReplyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
