package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdChatIdPermissionGrantId{}

func TestNewUserIdChatIdPermissionGrantID(t *testing.T) {
	id := NewUserIdChatIdPermissionGrantID("userId", "chatId", "resourceSpecificPermissionGrantId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ChatId != "chatId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatId")
	}

	if id.ResourceSpecificPermissionGrantId != "resourceSpecificPermissionGrantId" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceSpecificPermissionGrantId'", id.ResourceSpecificPermissionGrantId, "resourceSpecificPermissionGrantId")
	}
}

func TestFormatUserIdChatIdPermissionGrantID(t *testing.T) {
	actual := NewUserIdChatIdPermissionGrantID("userId", "chatId", "resourceSpecificPermissionGrantId").ID()
	expected := "/users/userId/chats/chatId/permissionGrants/resourceSpecificPermissionGrantId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdChatIdPermissionGrantID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdChatIdPermissionGrantId
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
			Input: "/users/userId/chats/chatId/permissionGrants",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/chats/chatId/permissionGrants/resourceSpecificPermissionGrantId",
			Expected: &UserIdChatIdPermissionGrantId{
				UserId:                            "userId",
				ChatId:                            "chatId",
				ResourceSpecificPermissionGrantId: "resourceSpecificPermissionGrantId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/chats/chatId/permissionGrants/resourceSpecificPermissionGrantId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdChatIdPermissionGrantID(v.Input)
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

		if actual.ResourceSpecificPermissionGrantId != v.Expected.ResourceSpecificPermissionGrantId {
			t.Fatalf("Expected %q but got %q for ResourceSpecificPermissionGrantId", v.Expected.ResourceSpecificPermissionGrantId, actual.ResourceSpecificPermissionGrantId)
		}

	}
}

func TestParseUserIdChatIdPermissionGrantIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdChatIdPermissionGrantId
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
			Input: "/users/userId/chats/chatId/permissionGrants",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/pErMiSsIoNgRaNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/chats/chatId/permissionGrants/resourceSpecificPermissionGrantId",
			Expected: &UserIdChatIdPermissionGrantId{
				UserId:                            "userId",
				ChatId:                            "chatId",
				ResourceSpecificPermissionGrantId: "resourceSpecificPermissionGrantId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/chats/chatId/permissionGrants/resourceSpecificPermissionGrantId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/pErMiSsIoNgRaNtS/rEsOuRcEsPeCiFiCpErMiSsIoNgRaNtId",
			Expected: &UserIdChatIdPermissionGrantId{
				UserId:                            "uSeRiD",
				ChatId:                            "cHaTiD",
				ResourceSpecificPermissionGrantId: "rEsOuRcEsPeCiFiCpErMiSsIoNgRaNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cHaTs/cHaTiD/pErMiSsIoNgRaNtS/rEsOuRcEsPeCiFiCpErMiSsIoNgRaNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdChatIdPermissionGrantIDInsensitively(v.Input)
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

		if actual.ResourceSpecificPermissionGrantId != v.Expected.ResourceSpecificPermissionGrantId {
			t.Fatalf("Expected %q but got %q for ResourceSpecificPermissionGrantId", v.Expected.ResourceSpecificPermissionGrantId, actual.ResourceSpecificPermissionGrantId)
		}

	}
}

func TestSegmentsForUserIdChatIdPermissionGrantId(t *testing.T) {
	segments := UserIdChatIdPermissionGrantId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdChatIdPermissionGrantId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
