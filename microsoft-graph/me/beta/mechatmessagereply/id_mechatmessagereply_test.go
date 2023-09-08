package mechatmessagereply

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeChatMessageReplyId{}

func TestNewMeChatMessageReplyID(t *testing.T) {
	id := NewMeChatMessageReplyID("chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

	if id.ChatId != "chatIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatIdValue")
	}

	if id.ChatMessageId != "chatMessageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId'", id.ChatMessageId, "chatMessageIdValue")
	}

	if id.ChatMessageId1 != "chatMessageId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId1'", id.ChatMessageId1, "chatMessageId1Value")
	}
}

func TestFormatMeChatMessageReplyID(t *testing.T) {
	actual := NewMeChatMessageReplyID("chatIdValue", "chatMessageIdValue", "chatMessageId1Value").ID()
	expected := "/me/chats/chatIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeChatMessageReplyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatMessageReplyId
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
			Input: "/me/chats/chatIdValue/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/chats/chatIdValue/messages/chatMessageIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/chats/chatIdValue/messages/chatMessageIdValue/replies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value",
			Expected: &MeChatMessageReplyId{
				ChatId:         "chatIdValue",
				ChatMessageId:  "chatMessageIdValue",
				ChatMessageId1: "chatMessageId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatMessageReplyID(v.Input)
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

func TestParseMeChatMessageReplyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatMessageReplyId
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
			Input: "/me/chats/chatIdValue/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/chats/chatIdValue/messages/chatMessageIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/mEsSaGeS/cHaTmEsSaGeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/chats/chatIdValue/messages/chatMessageIdValue/replies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value",
			Expected: &MeChatMessageReplyId{
				ChatId:         "chatIdValue",
				ChatMessageId:  "chatMessageIdValue",
				ChatMessageId1: "chatMessageId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe",
			Expected: &MeChatMessageReplyId{
				ChatId:         "cHaTiDvAlUe",
				ChatMessageId:  "cHaTmEsSaGeIdVaLuE",
				ChatMessageId1: "cHaTmEsSaGeId1vAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatMessageReplyIDInsensitively(v.Input)
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

func TestSegmentsForMeChatMessageReplyId(t *testing.T) {
	segments := MeChatMessageReplyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeChatMessageReplyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
