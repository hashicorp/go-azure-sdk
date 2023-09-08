package mechatpermissiongrant

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeChatPermissionGrantId{}

func TestNewMeChatPermissionGrantID(t *testing.T) {
	id := NewMeChatPermissionGrantID("chatIdValue", "resourceSpecificPermissionGrantIdValue")

	if id.ChatId != "chatIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatIdValue")
	}

	if id.ResourceSpecificPermissionGrantId != "resourceSpecificPermissionGrantIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceSpecificPermissionGrantId'", id.ResourceSpecificPermissionGrantId, "resourceSpecificPermissionGrantIdValue")
	}
}

func TestFormatMeChatPermissionGrantID(t *testing.T) {
	actual := NewMeChatPermissionGrantID("chatIdValue", "resourceSpecificPermissionGrantIdValue").ID()
	expected := "/me/chats/chatIdValue/permissionGrants/resourceSpecificPermissionGrantIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeChatPermissionGrantID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatPermissionGrantId
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
			Input: "/me/chats/chatIdValue/permissionGrants",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatIdValue/permissionGrants/resourceSpecificPermissionGrantIdValue",
			Expected: &MeChatPermissionGrantId{
				ChatId:                            "chatIdValue",
				ResourceSpecificPermissionGrantId: "resourceSpecificPermissionGrantIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatIdValue/permissionGrants/resourceSpecificPermissionGrantIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatPermissionGrantID(v.Input)
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

		if actual.ResourceSpecificPermissionGrantId != v.Expected.ResourceSpecificPermissionGrantId {
			t.Fatalf("Expected %q but got %q for ResourceSpecificPermissionGrantId", v.Expected.ResourceSpecificPermissionGrantId, actual.ResourceSpecificPermissionGrantId)
		}

	}
}

func TestParseMeChatPermissionGrantIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatPermissionGrantId
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
			Input: "/me/chats/chatIdValue/permissionGrants",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/pErMiSsIoNgRaNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatIdValue/permissionGrants/resourceSpecificPermissionGrantIdValue",
			Expected: &MeChatPermissionGrantId{
				ChatId:                            "chatIdValue",
				ResourceSpecificPermissionGrantId: "resourceSpecificPermissionGrantIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatIdValue/permissionGrants/resourceSpecificPermissionGrantIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/pErMiSsIoNgRaNtS/rEsOuRcEsPeCiFiCpErMiSsIoNgRaNtIdVaLuE",
			Expected: &MeChatPermissionGrantId{
				ChatId:                            "cHaTiDvAlUe",
				ResourceSpecificPermissionGrantId: "rEsOuRcEsPeCiFiCpErMiSsIoNgRaNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/pErMiSsIoNgRaNtS/rEsOuRcEsPeCiFiCpErMiSsIoNgRaNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatPermissionGrantIDInsensitively(v.Input)
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

		if actual.ResourceSpecificPermissionGrantId != v.Expected.ResourceSpecificPermissionGrantId {
			t.Fatalf("Expected %q but got %q for ResourceSpecificPermissionGrantId", v.Expected.ResourceSpecificPermissionGrantId, actual.ResourceSpecificPermissionGrantId)
		}

	}
}

func TestSegmentsForMeChatPermissionGrantId(t *testing.T) {
	segments := MeChatPermissionGrantId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeChatPermissionGrantId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
