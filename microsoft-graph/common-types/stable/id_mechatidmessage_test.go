package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeChatIdMessageId{}

func TestNewMeChatIdMessageID(t *testing.T) {
	id := NewMeChatIdMessageID("chatId", "chatMessageId")

	if id.ChatId != "chatId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatId")
	}

	if id.ChatMessageId != "chatMessageId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId'", id.ChatMessageId, "chatMessageId")
	}
}

func TestFormatMeChatIdMessageID(t *testing.T) {
	actual := NewMeChatIdMessageID("chatId", "chatMessageId").ID()
	expected := "/me/chats/chatId/messages/chatMessageId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeChatIdMessageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatIdMessageId
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
			// Valid URI
			Input: "/me/chats/chatId/messages/chatMessageId",
			Expected: &MeChatIdMessageId{
				ChatId:        "chatId",
				ChatMessageId: "chatMessageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatId/messages/chatMessageId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatIdMessageID(v.Input)
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

	}
}

func TestParseMeChatIdMessageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatIdMessageId
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
			// Valid URI
			Input: "/me/chats/chatId/messages/chatMessageId",
			Expected: &MeChatIdMessageId{
				ChatId:        "chatId",
				ChatMessageId: "chatMessageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatId/messages/chatMessageId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD/mEsSaGeS/cHaTmEsSaGeId",
			Expected: &MeChatIdMessageId{
				ChatId:        "cHaTiD",
				ChatMessageId: "cHaTmEsSaGeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD/mEsSaGeS/cHaTmEsSaGeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatIdMessageIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeChatIdMessageId(t *testing.T) {
	segments := MeChatIdMessageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeChatIdMessageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
