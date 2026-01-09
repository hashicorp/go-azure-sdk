package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeChatIdOperationId{}

func TestNewMeChatIdOperationID(t *testing.T) {
	id := NewMeChatIdOperationID("chatId", "teamsAsyncOperationId")

	if id.ChatId != "chatId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatId")
	}

	if id.TeamsAsyncOperationId != "teamsAsyncOperationId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamsAsyncOperationId'", id.TeamsAsyncOperationId, "teamsAsyncOperationId")
	}
}

func TestFormatMeChatIdOperationID(t *testing.T) {
	actual := NewMeChatIdOperationID("chatId", "teamsAsyncOperationId").ID()
	expected := "/me/chats/chatId/operations/teamsAsyncOperationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeChatIdOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatIdOperationId
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
			Input: "/me/chats/chatId/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatId/operations/teamsAsyncOperationId",
			Expected: &MeChatIdOperationId{
				ChatId:                "chatId",
				TeamsAsyncOperationId: "teamsAsyncOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatId/operations/teamsAsyncOperationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatIdOperationID(v.Input)
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

		if actual.TeamsAsyncOperationId != v.Expected.TeamsAsyncOperationId {
			t.Fatalf("Expected %q but got %q for TeamsAsyncOperationId", v.Expected.TeamsAsyncOperationId, actual.TeamsAsyncOperationId)
		}

	}
}

func TestParseMeChatIdOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatIdOperationId
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
			Input: "/me/chats/chatId/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatId/operations/teamsAsyncOperationId",
			Expected: &MeChatIdOperationId{
				ChatId:                "chatId",
				TeamsAsyncOperationId: "teamsAsyncOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatId/operations/teamsAsyncOperationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD/oPeRaTiOnS/tEaMsAsYnCoPeRaTiOnId",
			Expected: &MeChatIdOperationId{
				ChatId:                "cHaTiD",
				TeamsAsyncOperationId: "tEaMsAsYnCoPeRaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD/oPeRaTiOnS/tEaMsAsYnCoPeRaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatIdOperationIDInsensitively(v.Input)
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

		if actual.TeamsAsyncOperationId != v.Expected.TeamsAsyncOperationId {
			t.Fatalf("Expected %q but got %q for TeamsAsyncOperationId", v.Expected.TeamsAsyncOperationId, actual.TeamsAsyncOperationId)
		}

	}
}

func TestSegmentsForMeChatIdOperationId(t *testing.T) {
	segments := MeChatIdOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeChatIdOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
