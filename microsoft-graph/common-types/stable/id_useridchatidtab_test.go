package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatIdTabId{}

func TestNewUserIdChatIdTabID(t *testing.T) {
	id := NewUserIdChatIdTabID("userId", "chatId", "teamsTabId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ChatId != "chatId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatId")
	}

	if id.TeamsTabId != "teamsTabId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamsTabId'", id.TeamsTabId, "teamsTabId")
	}
}

func TestFormatUserIdChatIdTabID(t *testing.T) {
	actual := NewUserIdChatIdTabID("userId", "chatId", "teamsTabId").ID()
	expected := "/users/userId/chats/chatId/tabs/teamsTabId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdChatIdTabID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdChatIdTabId
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
			Input: "/users/userId/chats/chatId/tabs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/chats/chatId/tabs/teamsTabId",
			Expected: &UserIdChatIdTabId{
				UserId:     "userId",
				ChatId:     "chatId",
				TeamsTabId: "teamsTabId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/chats/chatId/tabs/teamsTabId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdChatIdTabID(v.Input)
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

		if actual.TeamsTabId != v.Expected.TeamsTabId {
			t.Fatalf("Expected %q but got %q for TeamsTabId", v.Expected.TeamsTabId, actual.TeamsTabId)
		}

	}
}

func TestParseUserIdChatIdTabIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdChatIdTabId
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
			Input: "/users/userId/chats/chatId/tabs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/tAbS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/chats/chatId/tabs/teamsTabId",
			Expected: &UserIdChatIdTabId{
				UserId:     "userId",
				ChatId:     "chatId",
				TeamsTabId: "teamsTabId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/chats/chatId/tabs/teamsTabId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/tAbS/tEaMsTaBiD",
			Expected: &UserIdChatIdTabId{
				UserId:     "uSeRiD",
				ChatId:     "cHaTiD",
				TeamsTabId: "tEaMsTaBiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/tAbS/tEaMsTaBiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdChatIdTabIDInsensitively(v.Input)
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

		if actual.TeamsTabId != v.Expected.TeamsTabId {
			t.Fatalf("Expected %q but got %q for TeamsTabId", v.Expected.TeamsTabId, actual.TeamsTabId)
		}

	}
}

func TestSegmentsForUserIdChatIdTabId(t *testing.T) {
	segments := UserIdChatIdTabId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdChatIdTabId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
