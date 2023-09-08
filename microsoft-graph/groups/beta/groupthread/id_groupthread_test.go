package groupthread

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupThreadId{}

func TestNewGroupThreadID(t *testing.T) {
	id := NewGroupThreadID("groupIdValue", "conversationThreadIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.ConversationThreadId != "conversationThreadIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationThreadId'", id.ConversationThreadId, "conversationThreadIdValue")
	}
}

func TestFormatGroupThreadID(t *testing.T) {
	actual := NewGroupThreadID("groupIdValue", "conversationThreadIdValue").ID()
	expected := "/groups/groupIdValue/threads/conversationThreadIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupThreadID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupThreadId
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
			Input: "/groups/groupIdValue/threads",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/threads/conversationThreadIdValue",
			Expected: &GroupThreadId{
				GroupId:              "groupIdValue",
				ConversationThreadId: "conversationThreadIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/threads/conversationThreadIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupThreadID(v.Input)
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

		if actual.ConversationThreadId != v.Expected.ConversationThreadId {
			t.Fatalf("Expected %q but got %q for ConversationThreadId", v.Expected.ConversationThreadId, actual.ConversationThreadId)
		}

	}
}

func TestParseGroupThreadIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupThreadId
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
			Input: "/groups/groupIdValue/threads",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tHrEaDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/threads/conversationThreadIdValue",
			Expected: &GroupThreadId{
				GroupId:              "groupIdValue",
				ConversationThreadId: "conversationThreadIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/threads/conversationThreadIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tHrEaDs/cOnVeRsAtIoNtHrEaDiDvAlUe",
			Expected: &GroupThreadId{
				GroupId:              "gRoUpIdVaLuE",
				ConversationThreadId: "cOnVeRsAtIoNtHrEaDiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tHrEaDs/cOnVeRsAtIoNtHrEaDiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupThreadIDInsensitively(v.Input)
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

		if actual.ConversationThreadId != v.Expected.ConversationThreadId {
			t.Fatalf("Expected %q but got %q for ConversationThreadId", v.Expected.ConversationThreadId, actual.ConversationThreadId)
		}

	}
}

func TestSegmentsForGroupThreadId(t *testing.T) {
	segments := GroupThreadId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupThreadId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
