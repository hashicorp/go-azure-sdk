package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeChatIdMessageIdReplyId{}

func TestNewMeChatIdMessageIdReplyID(t *testing.T) {
	id := NewMeChatIdMessageIdReplyID("chatId", "chatMessageId", "chatMessageId1")

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

func TestFormatMeChatIdMessageIdReplyID(t *testing.T) {
	actual := NewMeChatIdMessageIdReplyID("chatId", "chatMessageId", "chatMessageId1").ID()
	expected := "/me/chats/chatId/messages/chatMessageId/replies/chatMessageId1"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeChatIdMessageIdReplyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatIdMessageIdReplyId
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
			Input: "/me/chats",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/chats/chatId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/chats/chatId/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/chats/chatId/messages/chatMessageId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/chats/chatId/messages/chatMessageId/replies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatId/messages/chatMessageId/replies/chatMessageId1",
			Expected: &MeChatIdMessageIdReplyId{
				ChatId:         "chatId",
				ChatMessageId:  "chatMessageId",
				ChatMessageId1: "chatMessageId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatId/messages/chatMessageId/replies/chatMessageId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatIdMessageIdReplyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
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

func TestParseMeChatIdMessageIdReplyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatIdMessageIdReplyId
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
			Input: "/me/chats",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/chats/chatId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/chats/chatId/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/chats/chatId/messages/chatMessageId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD/mEsSaGeS/cHaTmEsSaGeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/chats/chatId/messages/chatMessageId/replies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD/mEsSaGeS/cHaTmEsSaGeId/rEpLiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatId/messages/chatMessageId/replies/chatMessageId1",
			Expected: &MeChatIdMessageIdReplyId{
				ChatId:         "chatId",
				ChatMessageId:  "chatMessageId",
				ChatMessageId1: "chatMessageId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatId/messages/chatMessageId/replies/chatMessageId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD/mEsSaGeS/cHaTmEsSaGeId/rEpLiEs/cHaTmEsSaGeId1",
			Expected: &MeChatIdMessageIdReplyId{
				ChatId:         "cHaTiD",
				ChatMessageId:  "cHaTmEsSaGeId",
				ChatMessageId1: "cHaTmEsSaGeId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD/mEsSaGeS/cHaTmEsSaGeId/rEpLiEs/cHaTmEsSaGeId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatIdMessageIdReplyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
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

func TestSegmentsForMeChatIdMessageIdReplyId(t *testing.T) {
	segments := MeChatIdMessageIdReplyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeChatIdMessageIdReplyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
