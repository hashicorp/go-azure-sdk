package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatIdPinnedMessageId{}

func TestNewUserIdChatIdPinnedMessageID(t *testing.T) {
	id := NewUserIdChatIdPinnedMessageID("userId", "chatId", "pinnedChatMessageInfoId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ChatId != "chatId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatId")
	}

	if id.PinnedChatMessageInfoId != "pinnedChatMessageInfoId" {
		t.Fatalf("Expected %q but got %q for Segment 'PinnedChatMessageInfoId'", id.PinnedChatMessageInfoId, "pinnedChatMessageInfoId")
	}
}

func TestFormatUserIdChatIdPinnedMessageID(t *testing.T) {
	actual := NewUserIdChatIdPinnedMessageID("userId", "chatId", "pinnedChatMessageInfoId").ID()
	expected := "/users/userId/chats/chatId/pinnedMessages/pinnedChatMessageInfoId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdChatIdPinnedMessageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdChatIdPinnedMessageId
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
			Input: "/users/userId/chats/chatId/pinnedMessages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/chats/chatId/pinnedMessages/pinnedChatMessageInfoId",
			Expected: &UserIdChatIdPinnedMessageId{
				UserId:                  "userId",
				ChatId:                  "chatId",
				PinnedChatMessageInfoId: "pinnedChatMessageInfoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/chats/chatId/pinnedMessages/pinnedChatMessageInfoId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdChatIdPinnedMessageID(v.Input)
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

		if actual.PinnedChatMessageInfoId != v.Expected.PinnedChatMessageInfoId {
			t.Fatalf("Expected %q but got %q for PinnedChatMessageInfoId", v.Expected.PinnedChatMessageInfoId, actual.PinnedChatMessageInfoId)
		}

	}
}

func TestParseUserIdChatIdPinnedMessageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdChatIdPinnedMessageId
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
			Input: "/users/userId/chats/chatId/pinnedMessages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/pInNeDmEsSaGeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/chats/chatId/pinnedMessages/pinnedChatMessageInfoId",
			Expected: &UserIdChatIdPinnedMessageId{
				UserId:                  "userId",
				ChatId:                  "chatId",
				PinnedChatMessageInfoId: "pinnedChatMessageInfoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/chats/chatId/pinnedMessages/pinnedChatMessageInfoId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/pInNeDmEsSaGeS/pInNeDcHaTmEsSaGeInFoId",
			Expected: &UserIdChatIdPinnedMessageId{
				UserId:                  "uSeRiD",
				ChatId:                  "cHaTiD",
				PinnedChatMessageInfoId: "pInNeDcHaTmEsSaGeInFoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/pInNeDmEsSaGeS/pInNeDcHaTmEsSaGeInFoId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdChatIdPinnedMessageIDInsensitively(v.Input)
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

		if actual.PinnedChatMessageInfoId != v.Expected.PinnedChatMessageInfoId {
			t.Fatalf("Expected %q but got %q for PinnedChatMessageInfoId", v.Expected.PinnedChatMessageInfoId, actual.PinnedChatMessageInfoId)
		}

	}
}

func TestSegmentsForUserIdChatIdPinnedMessageId(t *testing.T) {
	segments := UserIdChatIdPinnedMessageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdChatIdPinnedMessageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
