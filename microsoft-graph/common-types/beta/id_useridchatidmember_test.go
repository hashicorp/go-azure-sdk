package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatIdMemberId{}

func TestNewUserIdChatIdMemberID(t *testing.T) {
	id := NewUserIdChatIdMemberID("userId", "chatId", "conversationMemberId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ChatId != "chatId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatId")
	}

	if id.ConversationMemberId != "conversationMemberId" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationMemberId'", id.ConversationMemberId, "conversationMemberId")
	}
}

func TestFormatUserIdChatIdMemberID(t *testing.T) {
	actual := NewUserIdChatIdMemberID("userId", "chatId", "conversationMemberId").ID()
	expected := "/users/userId/chats/chatId/members/conversationMemberId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdChatIdMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdChatIdMemberId
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
			Input: "/users/userId/chats",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/chats/chatId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/chats/chatId/members",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/chats/chatId/members/conversationMemberId",
			Expected: &UserIdChatIdMemberId{
				UserId:               "userId",
				ChatId:               "chatId",
				ConversationMemberId: "conversationMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/chats/chatId/members/conversationMemberId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdChatIdMemberID(v.Input)
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

		if actual.ChatId != v.Expected.ChatId {
			t.Fatalf("Expected %q but got %q for ChatId", v.Expected.ChatId, actual.ChatId)
		}

		if actual.ConversationMemberId != v.Expected.ConversationMemberId {
			t.Fatalf("Expected %q but got %q for ConversationMemberId", v.Expected.ConversationMemberId, actual.ConversationMemberId)
		}

	}
}

func TestParseUserIdChatIdMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdChatIdMemberId
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
			Input: "/users/userId/chats",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/chats/chatId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/chats/chatId/members",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/mEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/chats/chatId/members/conversationMemberId",
			Expected: &UserIdChatIdMemberId{
				UserId:               "userId",
				ChatId:               "chatId",
				ConversationMemberId: "conversationMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/chats/chatId/members/conversationMemberId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/mEmBeRs/cOnVeRsAtIoNmEmBeRiD",
			Expected: &UserIdChatIdMemberId{
				UserId:               "uSeRiD",
				ChatId:               "cHaTiD",
				ConversationMemberId: "cOnVeRsAtIoNmEmBeRiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/mEmBeRs/cOnVeRsAtIoNmEmBeRiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdChatIdMemberIDInsensitively(v.Input)
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

		if actual.ChatId != v.Expected.ChatId {
			t.Fatalf("Expected %q but got %q for ChatId", v.Expected.ChatId, actual.ChatId)
		}

		if actual.ConversationMemberId != v.Expected.ConversationMemberId {
			t.Fatalf("Expected %q but got %q for ConversationMemberId", v.Expected.ConversationMemberId, actual.ConversationMemberId)
		}

	}
}

func TestSegmentsForUserIdChatIdMemberId(t *testing.T) {
	segments := UserIdChatIdMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdChatIdMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
