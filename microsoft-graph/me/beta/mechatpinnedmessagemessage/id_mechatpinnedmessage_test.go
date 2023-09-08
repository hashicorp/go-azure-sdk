package mechatpinnedmessagemessage

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeChatPinnedMessageId{}

func TestNewMeChatPinnedMessageID(t *testing.T) {
	id := NewMeChatPinnedMessageID("chatIdValue", "pinnedChatMessageInfoIdValue")

	if id.ChatId != "chatIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatIdValue")
	}

	if id.PinnedChatMessageInfoId != "pinnedChatMessageInfoIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PinnedChatMessageInfoId'", id.PinnedChatMessageInfoId, "pinnedChatMessageInfoIdValue")
	}
}

func TestFormatMeChatPinnedMessageID(t *testing.T) {
	actual := NewMeChatPinnedMessageID("chatIdValue", "pinnedChatMessageInfoIdValue").ID()
	expected := "/me/chats/chatIdValue/pinnedMessages/pinnedChatMessageInfoIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeChatPinnedMessageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatPinnedMessageId
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
			Input: "/me/chats/chatIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/chats/chatIdValue/pinnedMessages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatIdValue/pinnedMessages/pinnedChatMessageInfoIdValue",
			Expected: &MeChatPinnedMessageId{
				ChatId:                  "chatIdValue",
				PinnedChatMessageInfoId: "pinnedChatMessageInfoIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatIdValue/pinnedMessages/pinnedChatMessageInfoIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatPinnedMessageID(v.Input)
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

func TestParseMeChatPinnedMessageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatPinnedMessageId
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
			Input: "/me/chats/chatIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/chats/chatIdValue/pinnedMessages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/pInNeDmEsSaGeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatIdValue/pinnedMessages/pinnedChatMessageInfoIdValue",
			Expected: &MeChatPinnedMessageId{
				ChatId:                  "chatIdValue",
				PinnedChatMessageInfoId: "pinnedChatMessageInfoIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatIdValue/pinnedMessages/pinnedChatMessageInfoIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/pInNeDmEsSaGeS/pInNeDcHaTmEsSaGeInFoIdVaLuE",
			Expected: &MeChatPinnedMessageId{
				ChatId:                  "cHaTiDvAlUe",
				PinnedChatMessageInfoId: "pInNeDcHaTmEsSaGeInFoIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/pInNeDmEsSaGeS/pInNeDcHaTmEsSaGeInFoIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatPinnedMessageIDInsensitively(v.Input)
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

func TestSegmentsForMeChatPinnedMessageId(t *testing.T) {
	segments := MeChatPinnedMessageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeChatPinnedMessageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
