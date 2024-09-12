package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdConversationId{}

func TestNewGroupIdConversationID(t *testing.T) {
	id := NewGroupIdConversationID("groupIdValue", "conversationIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.ConversationId != "conversationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationId'", id.ConversationId, "conversationIdValue")
	}
}

func TestFormatGroupIdConversationID(t *testing.T) {
	actual := NewGroupIdConversationID("groupIdValue", "conversationIdValue").ID()
	expected := "/groups/groupIdValue/conversations/conversationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdConversationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdConversationId
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
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/conversations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue",
			Expected: &GroupIdConversationId{
				GroupId:        "groupIdValue",
				ConversationId: "conversationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/conversations/conversationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdConversationID(v.Input)
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

	}
}

func TestParseGroupIdConversationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdConversationId
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
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/conversations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cOnVeRsAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue",
			Expected: &GroupIdConversationId{
				GroupId:        "groupIdValue",
				ConversationId: "conversationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/conversations/conversationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cOnVeRsAtIoNs/cOnVeRsAtIoNiDvAlUe",
			Expected: &GroupIdConversationId{
				GroupId:        "gRoUpIdVaLuE",
				ConversationId: "cOnVeRsAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cOnVeRsAtIoNs/cOnVeRsAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdConversationIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForGroupIdConversationId(t *testing.T) {
	segments := GroupIdConversationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdConversationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
