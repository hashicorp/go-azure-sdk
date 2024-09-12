package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatIdPinnedMessageId{}

func TestNewUserIdChatIdPinnedMessageID(t *testing.T) {
	id := NewUserIdChatIdPinnedMessageID("userIdValue", "chatIdValue", "pinnedChatMessageInfoIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ChatId != "chatIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatIdValue")
	}

	if id.PinnedChatMessageInfoId != "pinnedChatMessageInfoIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PinnedChatMessageInfoId'", id.PinnedChatMessageInfoId, "pinnedChatMessageInfoIdValue")
	}
}

func TestFormatUserIdChatIdPinnedMessageID(t *testing.T) {
	actual := NewUserIdChatIdPinnedMessageID("userIdValue", "chatIdValue", "pinnedChatMessageInfoIdValue").ID()
	expected := "/users/userIdValue/chats/chatIdValue/pinnedMessages/pinnedChatMessageInfoIdValue"
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats/chatIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats/chatIdValue/pinnedMessages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/chats/chatIdValue/pinnedMessages/pinnedChatMessageInfoIdValue",
			Expected: &UserIdChatIdPinnedMessageId{
				UserId:                  "userIdValue",
				ChatId:                  "chatIdValue",
				PinnedChatMessageInfoId: "pinnedChatMessageInfoIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/chats/chatIdValue/pinnedMessages/pinnedChatMessageInfoIdValue/extra",
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
			Input: "/users/userIdValue/chats",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats/chatIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs/cHaTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats/chatIdValue/pinnedMessages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs/cHaTiDvAlUe/pInNeDmEsSaGeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/chats/chatIdValue/pinnedMessages/pinnedChatMessageInfoIdValue",
			Expected: &UserIdChatIdPinnedMessageId{
				UserId:                  "userIdValue",
				ChatId:                  "chatIdValue",
				PinnedChatMessageInfoId: "pinnedChatMessageInfoIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/chats/chatIdValue/pinnedMessages/pinnedChatMessageInfoIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs/cHaTiDvAlUe/pInNeDmEsSaGeS/pInNeDcHaTmEsSaGeInFoIdVaLuE",
			Expected: &UserIdChatIdPinnedMessageId{
				UserId:                  "uSeRiDvAlUe",
				ChatId:                  "cHaTiDvAlUe",
				PinnedChatMessageInfoId: "pInNeDcHaTmEsSaGeInFoIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs/cHaTiDvAlUe/pInNeDmEsSaGeS/pInNeDcHaTmEsSaGeInFoIdVaLuE/extra",
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
