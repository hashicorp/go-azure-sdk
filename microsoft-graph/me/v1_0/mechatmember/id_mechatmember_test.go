package mechatmember

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeChatMemberId{}

func TestNewMeChatMemberID(t *testing.T) {
	id := NewMeChatMemberID("chatIdValue", "conversationMemberIdValue")

	if id.ChatId != "chatIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatIdValue")
	}

	if id.ConversationMemberId != "conversationMemberIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationMemberId'", id.ConversationMemberId, "conversationMemberIdValue")
	}
}

func TestFormatMeChatMemberID(t *testing.T) {
	actual := NewMeChatMemberID("chatIdValue", "conversationMemberIdValue").ID()
	expected := "/me/chats/chatIdValue/members/conversationMemberIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeChatMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatMemberId
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
			Input: "/me/chats/chatIdValue/members",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatIdValue/members/conversationMemberIdValue",
			Expected: &MeChatMemberId{
				ChatId:               "chatIdValue",
				ConversationMemberId: "conversationMemberIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatIdValue/members/conversationMemberIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatMemberID(v.Input)
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

		if actual.ConversationMemberId != v.Expected.ConversationMemberId {
			t.Fatalf("Expected %q but got %q for ConversationMemberId", v.Expected.ConversationMemberId, actual.ConversationMemberId)
		}

	}
}

func TestParseMeChatMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeChatMemberId
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
			Input: "/me/chats/chatIdValue/members",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/mEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/chats/chatIdValue/members/conversationMemberIdValue",
			Expected: &MeChatMemberId{
				ChatId:               "chatIdValue",
				ConversationMemberId: "conversationMemberIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/chats/chatIdValue/members/conversationMemberIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/mEmBeRs/cOnVeRsAtIoNmEmBeRiDvAlUe",
			Expected: &MeChatMemberId{
				ChatId:               "cHaTiDvAlUe",
				ConversationMemberId: "cOnVeRsAtIoNmEmBeRiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cHaTs/cHaTiDvAlUe/mEmBeRs/cOnVeRsAtIoNmEmBeRiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeChatMemberIDInsensitively(v.Input)
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

		if actual.ConversationMemberId != v.Expected.ConversationMemberId {
			t.Fatalf("Expected %q but got %q for ConversationMemberId", v.Expected.ConversationMemberId, actual.ConversationMemberId)
		}

	}
}

func TestSegmentsForMeChatMemberId(t *testing.T) {
	segments := MeChatMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeChatMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
