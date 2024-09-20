package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatIdMessageIdReplyId{}

func TestNewUserIdChatIdMessageIdReplyID(t *testing.T) {
	id := NewUserIdChatIdMessageIdReplyID("userId", "chatId", "chatMessageId", "chatMessageId1")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ChatId != "chatId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatId")
	}

	if id.ChatMessageId != "chatMessageId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId'", id.ChatMessageId, "chatMessageId")
	}

	if id.ChatMessageId1 != "chatMessageId1" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId1'", id.ChatMessageId1, "chatMessageId1")
	}
}

func TestFormatUserIdChatIdMessageIdReplyID(t *testing.T) {
	actual := NewUserIdChatIdMessageIdReplyID("userId", "chatId", "chatMessageId", "chatMessageId1").ID()
	expected := "/users/userId/chats/chatId/messages/chatMessageId/replies/chatMessageId1"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdChatIdMessageIdReplyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdChatIdMessageIdReplyId
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
			Input: "/users/userId/chats/chatId/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/chats/chatId/messages/chatMessageId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/chats/chatId/messages/chatMessageId/replies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/chats/chatId/messages/chatMessageId/replies/chatMessageId1",
			Expected: &UserIdChatIdMessageIdReplyId{
				UserId:         "userId",
				ChatId:         "chatId",
				ChatMessageId:  "chatMessageId",
				ChatMessageId1: "chatMessageId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/chats/chatId/messages/chatMessageId/replies/chatMessageId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdChatIdMessageIdReplyID(v.Input)
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

		if actual.ChatMessageId != v.Expected.ChatMessageId {
			t.Fatalf("Expected %q but got %q for ChatMessageId", v.Expected.ChatMessageId, actual.ChatMessageId)
		}

		if actual.ChatMessageId1 != v.Expected.ChatMessageId1 {
			t.Fatalf("Expected %q but got %q for ChatMessageId1", v.Expected.ChatMessageId1, actual.ChatMessageId1)
		}

	}
}

func TestParseUserIdChatIdMessageIdReplyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdChatIdMessageIdReplyId
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
			Input: "/users/userId/chats/chatId/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/chats/chatId/messages/chatMessageId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/mEsSaGeS/cHaTmEsSaGeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/chats/chatId/messages/chatMessageId/replies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/mEsSaGeS/cHaTmEsSaGeId/rEpLiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/chats/chatId/messages/chatMessageId/replies/chatMessageId1",
			Expected: &UserIdChatIdMessageIdReplyId{
				UserId:         "userId",
				ChatId:         "chatId",
				ChatMessageId:  "chatMessageId",
				ChatMessageId1: "chatMessageId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/chats/chatId/messages/chatMessageId/replies/chatMessageId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/mEsSaGeS/cHaTmEsSaGeId/rEpLiEs/cHaTmEsSaGeId1",
			Expected: &UserIdChatIdMessageIdReplyId{
				UserId:         "uSeRiD",
				ChatId:         "cHaTiD",
				ChatMessageId:  "cHaTmEsSaGeId",
				ChatMessageId1: "cHaTmEsSaGeId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/mEsSaGeS/cHaTmEsSaGeId/rEpLiEs/cHaTmEsSaGeId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdChatIdMessageIdReplyIDInsensitively(v.Input)
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

		if actual.ChatMessageId != v.Expected.ChatMessageId {
			t.Fatalf("Expected %q but got %q for ChatMessageId", v.Expected.ChatMessageId, actual.ChatMessageId)
		}

		if actual.ChatMessageId1 != v.Expected.ChatMessageId1 {
			t.Fatalf("Expected %q but got %q for ChatMessageId1", v.Expected.ChatMessageId1, actual.ChatMessageId1)
		}

	}
}

func TestSegmentsForUserIdChatIdMessageIdReplyId(t *testing.T) {
	segments := UserIdChatIdMessageIdReplyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdChatIdMessageIdReplyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
