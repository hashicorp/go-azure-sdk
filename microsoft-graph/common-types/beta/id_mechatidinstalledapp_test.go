package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeChatIdInstalledAppId{}

func TestNewMeChatIdInstalledAppID(t *testing.T) {
	id := NewMeChatIdInstalledAppID("chatId", "teamsAppInstallationId")

	if id.ChatId != "chatId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatId")
	}

	if id.TeamsAppInstallationId != "teamsAppInstallationId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamsAppInstallationId'", id.TeamsAppInstallationId, "teamsAppInstallationId")
	}
}

func TestFormatMeChatIdInstalledAppID(t *testing.T) {
	actual := NewMeChatIdInstalledAppID("chatId", "teamsAppInstallationId").ID()
	expected := "/me/chats/chatId/installedApps/teamsAppInstallationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeChatIdInstalledAppID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatIdInstalledAppId
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
			Input: "/me/chats/chatId/installedApps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatId/installedApps/teamsAppInstallationId",
			Expected: &MeChatIdInstalledAppId{
				ChatId:                 "chatId",
				TeamsAppInstallationId: "teamsAppInstallationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatId/installedApps/teamsAppInstallationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatIdInstalledAppID(v.Input)
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

		if actual.TeamsAppInstallationId != v.Expected.TeamsAppInstallationId {
			t.Fatalf("Expected %q but got %q for TeamsAppInstallationId", v.Expected.TeamsAppInstallationId, actual.TeamsAppInstallationId)
		}

	}
}

func TestParseMeChatIdInstalledAppIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatIdInstalledAppId
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
			Input: "/me/chats/chatId/installedApps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD/iNsTaLlEdApPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatId/installedApps/teamsAppInstallationId",
			Expected: &MeChatIdInstalledAppId{
				ChatId:                 "chatId",
				TeamsAppInstallationId: "teamsAppInstallationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatId/installedApps/teamsAppInstallationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD/iNsTaLlEdApPs/tEaMsApPiNsTaLlAtIoNiD",
			Expected: &MeChatIdInstalledAppId{
				ChatId:                 "cHaTiD",
				TeamsAppInstallationId: "tEaMsApPiNsTaLlAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiD/iNsTaLlEdApPs/tEaMsApPiNsTaLlAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatIdInstalledAppIDInsensitively(v.Input)
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

		if actual.TeamsAppInstallationId != v.Expected.TeamsAppInstallationId {
			t.Fatalf("Expected %q but got %q for TeamsAppInstallationId", v.Expected.TeamsAppInstallationId, actual.TeamsAppInstallationId)
		}

	}
}

func TestSegmentsForMeChatIdInstalledAppId(t *testing.T) {
	segments := MeChatIdInstalledAppId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeChatIdInstalledAppId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
