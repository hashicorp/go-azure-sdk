package mejoinedteamprimarychannelmessagehostedcontent

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamPrimaryChannelMessageId{}

func TestNewMeJoinedTeamPrimaryChannelMessageID(t *testing.T) {
	id := NewMeJoinedTeamPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.ChatMessageId != "chatMessageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId'", id.ChatMessageId, "chatMessageIdValue")
	}
}

func TestFormatMeJoinedTeamPrimaryChannelMessageID(t *testing.T) {
	actual := NewMeJoinedTeamPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue").ID()
	expected := "/me/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamPrimaryChannelMessageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamPrimaryChannelMessageId
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
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue",
			Expected: &MeJoinedTeamPrimaryChannelMessageId{
				TeamId:        "teamIdValue",
				ChatMessageId: "chatMessageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamPrimaryChannelMessageID(v.Input)
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

	}
}

func TestParseMeJoinedTeamPrimaryChannelMessageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamPrimaryChannelMessageId
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
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue",
			Expected: &MeJoinedTeamPrimaryChannelMessageId{
				TeamId:        "teamIdValue",
				ChatMessageId: "chatMessageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE",
			Expected: &MeJoinedTeamPrimaryChannelMessageId{
				TeamId:        "tEaMiDvAlUe",
				ChatMessageId: "cHaTmEsSaGeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamPrimaryChannelMessageIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeJoinedTeamPrimaryChannelMessageId(t *testing.T) {
	segments := MeJoinedTeamPrimaryChannelMessageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamPrimaryChannelMessageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
