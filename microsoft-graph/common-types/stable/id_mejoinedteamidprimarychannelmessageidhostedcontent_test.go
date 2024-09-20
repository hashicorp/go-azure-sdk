package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId{}

func TestNewMeJoinedTeamIdPrimaryChannelMessageIdHostedContentID(t *testing.T) {
	id := NewMeJoinedTeamIdPrimaryChannelMessageIdHostedContentID("teamId", "chatMessageId", "chatMessageHostedContentId")

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.ChatMessageId != "chatMessageId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId'", id.ChatMessageId, "chatMessageId")
	}

	if id.ChatMessageHostedContentId != "chatMessageHostedContentId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageHostedContentId'", id.ChatMessageHostedContentId, "chatMessageHostedContentId")
	}
}

func TestFormatMeJoinedTeamIdPrimaryChannelMessageIdHostedContentID(t *testing.T) {
	actual := NewMeJoinedTeamIdPrimaryChannelMessageIdHostedContentID("teamId", "chatMessageId", "chatMessageHostedContentId").ID()
	expected := "/me/joinedTeams/teamId/primaryChannel/messages/chatMessageId/hostedContents/chatMessageHostedContentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamIdPrimaryChannelMessageIdHostedContentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId
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
			Input: "/me/joinedTeams/teamId/primaryChannel/messages/chatMessageId/hostedContents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/primaryChannel/messages/chatMessageId/hostedContents/chatMessageHostedContentId",
			Expected: &MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId{
				TeamId:                     "teamId",
				ChatMessageId:              "chatMessageId",
				ChatMessageHostedContentId: "chatMessageHostedContentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/primaryChannel/messages/chatMessageId/hostedContents/chatMessageHostedContentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdPrimaryChannelMessageIdHostedContentID(v.Input)
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

		if actual.ChatMessageHostedContentId != v.Expected.ChatMessageHostedContentId {
			t.Fatalf("Expected %q but got %q for ChatMessageHostedContentId", v.Expected.ChatMessageHostedContentId, actual.ChatMessageHostedContentId)
		}

	}
}

func TestParseMeJoinedTeamIdPrimaryChannelMessageIdHostedContentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId
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
			Input: "/me/joinedTeams/teamId/primaryChannel/messages/chatMessageId/hostedContents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeId/hOsTeDcOnTeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/primaryChannel/messages/chatMessageId/hostedContents/chatMessageHostedContentId",
			Expected: &MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId{
				TeamId:                     "teamId",
				ChatMessageId:              "chatMessageId",
				ChatMessageHostedContentId: "chatMessageHostedContentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/primaryChannel/messages/chatMessageId/hostedContents/chatMessageHostedContentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeId/hOsTeDcOnTeNtS/cHaTmEsSaGeHoStEdCoNtEnTiD",
			Expected: &MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId{
				TeamId:                     "tEaMiD",
				ChatMessageId:              "cHaTmEsSaGeId",
				ChatMessageHostedContentId: "cHaTmEsSaGeHoStEdCoNtEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeId/hOsTeDcOnTeNtS/cHaTmEsSaGeHoStEdCoNtEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdPrimaryChannelMessageIdHostedContentIDInsensitively(v.Input)
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

		if actual.ChatMessageHostedContentId != v.Expected.ChatMessageHostedContentId {
			t.Fatalf("Expected %q but got %q for ChatMessageHostedContentId", v.Expected.ChatMessageHostedContentId, actual.ChatMessageHostedContentId)
		}

	}
}

func TestSegmentsForMeJoinedTeamIdPrimaryChannelMessageIdHostedContentId(t *testing.T) {
	segments := MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
