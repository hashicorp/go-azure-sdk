package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdConversationIdThreadId{}

func TestNewGroupIdConversationIdThreadID(t *testing.T) {
	id := NewGroupIdConversationIdThreadID("groupId", "conversationId", "conversationThreadId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.ConversationId != "conversationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationId'", id.ConversationId, "conversationId")
	}

	if id.ConversationThreadId != "conversationThreadId" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationThreadId'", id.ConversationThreadId, "conversationThreadId")
	}
}

func TestFormatGroupIdConversationIdThreadID(t *testing.T) {
	actual := NewGroupIdConversationIdThreadID("groupId", "conversationId", "conversationThreadId").ID()
	expected := "/groups/groupId/conversations/conversationId/threads/conversationThreadId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdConversationIdThreadID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdConversationIdThreadId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/conversations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/conversations/conversationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/conversations/conversationId/threads",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId",
			Expected: &GroupIdConversationIdThreadId{
				GroupId:              "groupId",
				ConversationId:       "conversationId",
				ConversationThreadId: "conversationThreadId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdConversationIdThreadID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.ConversationId != v.Expected.ConversationId {
			t.Fatalf("Expected %q but got %q for ConversationId", v.Expected.ConversationId, actual.ConversationId)
		}

		if actual.ConversationThreadId != v.Expected.ConversationThreadId {
			t.Fatalf("Expected %q but got %q for ConversationThreadId", v.Expected.ConversationThreadId, actual.ConversationThreadId)
		}

	}
}

func TestParseGroupIdConversationIdThreadIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdConversationIdThreadId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/conversations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cOnVeRsAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/conversations/conversationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cOnVeRsAtIoNs/cOnVeRsAtIoNiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/conversations/conversationId/threads",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cOnVeRsAtIoNs/cOnVeRsAtIoNiD/tHrEaDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId",
			Expected: &GroupIdConversationIdThreadId{
				GroupId:              "groupId",
				ConversationId:       "conversationId",
				ConversationThreadId: "conversationThreadId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cOnVeRsAtIoNs/cOnVeRsAtIoNiD/tHrEaDs/cOnVeRsAtIoNtHrEaDiD",
			Expected: &GroupIdConversationIdThreadId{
				GroupId:              "gRoUpId",
				ConversationId:       "cOnVeRsAtIoNiD",
				ConversationThreadId: "cOnVeRsAtIoNtHrEaDiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cOnVeRsAtIoNs/cOnVeRsAtIoNiD/tHrEaDs/cOnVeRsAtIoNtHrEaDiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdConversationIdThreadIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.ConversationId != v.Expected.ConversationId {
			t.Fatalf("Expected %q but got %q for ConversationId", v.Expected.ConversationId, actual.ConversationId)
		}

		if actual.ConversationThreadId != v.Expected.ConversationThreadId {
			t.Fatalf("Expected %q but got %q for ConversationThreadId", v.Expected.ConversationThreadId, actual.ConversationThreadId)
		}

	}
}

func TestSegmentsForGroupIdConversationIdThreadId(t *testing.T) {
	segments := GroupIdConversationIdThreadId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdConversationIdThreadId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
