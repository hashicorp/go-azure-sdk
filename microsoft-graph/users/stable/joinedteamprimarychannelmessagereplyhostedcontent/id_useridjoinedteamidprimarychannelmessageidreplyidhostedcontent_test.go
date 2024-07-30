package joinedteamprimarychannelmessagereplyhostedcontent

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId{}

func TestNewUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentID(t *testing.T) {
	id := NewUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentID("userIdValue", "teamIdValue", "chatMessageIdValue", "chatMessageId1Value", "chatMessageHostedContentIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
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

func TestFormatUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentID(t *testing.T) {
	actual := NewUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentID("userIdValue", "teamIdValue", "chatMessageIdValue", "chatMessageId1Value", "chatMessageHostedContentIdValue").ID()
	expected := "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents/chatMessageHostedContentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents/chatMessageHostedContentIdValue",
			Expected: &UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId{
				UserId:                     "userIdValue",
				TeamId:                     "teamIdValue",
				ChatMessageId:              "chatMessageIdValue",
				ChatMessageId1:             "chatMessageId1Value",
				ChatMessageHostedContentId: "chatMessageHostedContentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents/chatMessageHostedContentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentID(v.Input)
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

func TestParseUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe/hOsTeDcOnTeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents/chatMessageHostedContentIdValue",
			Expected: &UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId{
				UserId:                     "userIdValue",
				TeamId:                     "teamIdValue",
				ChatMessageId:              "chatMessageIdValue",
				ChatMessageId1:             "chatMessageId1Value",
				ChatMessageHostedContentId: "chatMessageHostedContentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents/chatMessageHostedContentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe/hOsTeDcOnTeNtS/cHaTmEsSaGeHoStEdCoNtEnTiDvAlUe",
			Expected: &UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId{
				UserId:                     "uSeRiDvAlUe",
				TeamId:                     "tEaMiDvAlUe",
				ChatMessageId:              "cHaTmEsSaGeIdVaLuE",
				ChatMessageId1:             "cHaTmEsSaGeId1vAlUe",
				ChatMessageHostedContentId: "cHaTmEsSaGeHoStEdCoNtEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe/hOsTeDcOnTeNtS/cHaTmEsSaGeHoStEdCoNtEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentIDInsensitively(v.Input)
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

func TestSegmentsForUserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId(t *testing.T) {
	segments := UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdJoinedTeamIdPrimaryChannelMessageIdReplyIdHostedContentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
