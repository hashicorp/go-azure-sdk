package userchatoperation

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserChatOperationId{}

func TestNewUserChatOperationID(t *testing.T) {
	id := NewUserChatOperationID("userIdValue", "chatIdValue", "teamsAsyncOperationIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ChatId != "chatIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatIdValue")
	}

	if id.TeamsAsyncOperationId != "teamsAsyncOperationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamsAsyncOperationId'", id.TeamsAsyncOperationId, "teamsAsyncOperationIdValue")
	}
}

func TestFormatUserChatOperationID(t *testing.T) {
	actual := NewUserChatOperationID("userIdValue", "chatIdValue", "teamsAsyncOperationIdValue").ID()
	expected := "/users/userIdValue/chats/chatIdValue/operations/teamsAsyncOperationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserChatOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserChatOperationId
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
			Input: "/users/userIdValue/chats/chatIdValue/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/chats/chatIdValue/operations/teamsAsyncOperationIdValue",
			Expected: &UserChatOperationId{
				UserId:                "userIdValue",
				ChatId:                "chatIdValue",
				TeamsAsyncOperationId: "teamsAsyncOperationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/chats/chatIdValue/operations/teamsAsyncOperationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserChatOperationID(v.Input)
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

		if actual.TeamsAsyncOperationId != v.Expected.TeamsAsyncOperationId {
			t.Fatalf("Expected %q but got %q for TeamsAsyncOperationId", v.Expected.TeamsAsyncOperationId, actual.TeamsAsyncOperationId)
		}

	}
}

func TestParseUserChatOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserChatOperationId
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
			Input: "/users/userIdValue/chats/chatIdValue/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs/cHaTiDvAlUe/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/chats/chatIdValue/operations/teamsAsyncOperationIdValue",
			Expected: &UserChatOperationId{
				UserId:                "userIdValue",
				ChatId:                "chatIdValue",
				TeamsAsyncOperationId: "teamsAsyncOperationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/chats/chatIdValue/operations/teamsAsyncOperationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs/cHaTiDvAlUe/oPeRaTiOnS/tEaMsAsYnCoPeRaTiOnIdVaLuE",
			Expected: &UserChatOperationId{
				UserId:                "uSeRiDvAlUe",
				ChatId:                "cHaTiDvAlUe",
				TeamsAsyncOperationId: "tEaMsAsYnCoPeRaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs/cHaTiDvAlUe/oPeRaTiOnS/tEaMsAsYnCoPeRaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserChatOperationIDInsensitively(v.Input)
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

		if actual.TeamsAsyncOperationId != v.Expected.TeamsAsyncOperationId {
			t.Fatalf("Expected %q but got %q for TeamsAsyncOperationId", v.Expected.TeamsAsyncOperationId, actual.TeamsAsyncOperationId)
		}

	}
}

func TestSegmentsForUserChatOperationId(t *testing.T) {
	segments := UserChatOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserChatOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
