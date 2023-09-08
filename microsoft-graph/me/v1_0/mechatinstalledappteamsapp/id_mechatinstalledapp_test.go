package mechatinstalledappteamsapp

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeChatInstalledAppId{}

func TestNewMeChatInstalledAppID(t *testing.T) {
	id := NewMeChatInstalledAppID("chatIdValue", "teamsAppInstallationIdValue")

	if id.ChatId != "chatIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatIdValue")
	}

	if id.TeamsAppInstallationId != "teamsAppInstallationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamsAppInstallationId'", id.TeamsAppInstallationId, "teamsAppInstallationIdValue")
	}
}

func TestFormatMeChatInstalledAppID(t *testing.T) {
	actual := NewMeChatInstalledAppID("chatIdValue", "teamsAppInstallationIdValue").ID()
	expected := "/me/chats/chatIdValue/installedApps/teamsAppInstallationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeChatInstalledAppID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatInstalledAppId
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
			Input: "/me/chats/chatIdValue/installedApps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatIdValue/installedApps/teamsAppInstallationIdValue",
			Expected: &MeChatInstalledAppId{
				ChatId:                 "chatIdValue",
				TeamsAppInstallationId: "teamsAppInstallationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatIdValue/installedApps/teamsAppInstallationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatInstalledAppID(v.Input)
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

func TestParseMeChatInstalledAppIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatInstalledAppId
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
			Input: "/me/chats/chatIdValue/installedApps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/iNsTaLlEdApPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatIdValue/installedApps/teamsAppInstallationIdValue",
			Expected: &MeChatInstalledAppId{
				ChatId:                 "chatIdValue",
				TeamsAppInstallationId: "teamsAppInstallationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatIdValue/installedApps/teamsAppInstallationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/iNsTaLlEdApPs/tEaMsApPiNsTaLlAtIoNiDvAlUe",
			Expected: &MeChatInstalledAppId{
				ChatId:                 "cHaTiDvAlUe",
				TeamsAppInstallationId: "tEaMsApPiNsTaLlAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/iNsTaLlEdApPs/tEaMsApPiNsTaLlAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatInstalledAppIDInsensitively(v.Input)
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

func TestSegmentsForMeChatInstalledAppId(t *testing.T) {
	segments := MeChatInstalledAppId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeChatInstalledAppId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
