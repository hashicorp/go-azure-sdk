package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeChatIdPinnedMessageId{}

func TestNewMeChatIdPinnedMessageID(t *testing.T) {
	id := NewMeChatIdPinnedMessageID("chatId", "pinnedChatMessageInfoId")

	if id.ChatId != "chatId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatId")
	}

	if id.PinnedChatMessageInfoId != "pinnedChatMessageInfoId" {
		t.Fatalf("Expected %q but got %q for Segment 'PinnedChatMessageInfoId'", id.PinnedChatMessageInfoId, "pinnedChatMessageInfoId")
	}
}

func TestFormatMeChatIdPinnedMessageID(t *testing.T) {
	actual := NewMeChatIdPinnedMessageID("chatId", "pinnedChatMessageInfoId").ID()
	expected := "/me/chats/chatId/pinnedMessages/pinnedChatMessageInfoId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeChatIdPinnedMessageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatIdPinnedMessageId
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
			Input: "/me/chats/chatId/pinnedMessages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatId/pinnedMessages/pinnedChatMessageInfoId",
			Expected: &MeChatIdPinnedMessageId{
				ChatId:                  "chatId",
				PinnedChatMessageInfoId: "pinnedChatMessageInfoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatId/pinnedMessages/pinnedChatMessageInfoId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatIdPinnedMessageID(v.Input)
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

		if actual.PinnedChatMessageInfoId != v.Expected.PinnedChatMessageInfoId {
			t.Fatalf("Expected %q but got %q for PinnedChatMessageInfoId", v.Expected.PinnedChatMessageInfoId, actual.PinnedChatMessageInfoId)
		}

	}
}

func TestParseMeChatIdPinnedMessageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatIdPinnedMessageId
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
			Input: "/me/chats/chatId/pinnedMessages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD/pInNeDmEsSaGeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatId/pinnedMessages/pinnedChatMessageInfoId",
			Expected: &MeChatIdPinnedMessageId{
				ChatId:                  "chatId",
				PinnedChatMessageInfoId: "pinnedChatMessageInfoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatId/pinnedMessages/pinnedChatMessageInfoId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD/pInNeDmEsSaGeS/pInNeDcHaTmEsSaGeInFoId",
			Expected: &MeChatIdPinnedMessageId{
				ChatId:                  "cHaTiD",
				PinnedChatMessageInfoId: "pInNeDcHaTmEsSaGeInFoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD/pInNeDmEsSaGeS/pInNeDcHaTmEsSaGeInFoId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatIdPinnedMessageIDInsensitively(v.Input)
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

		if actual.PinnedChatMessageInfoId != v.Expected.PinnedChatMessageInfoId {
			t.Fatalf("Expected %q but got %q for PinnedChatMessageInfoId", v.Expected.PinnedChatMessageInfoId, actual.PinnedChatMessageInfoId)
		}

	}
}

func TestSegmentsForMeChatIdPinnedMessageId(t *testing.T) {
	segments := MeChatIdPinnedMessageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeChatIdPinnedMessageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
