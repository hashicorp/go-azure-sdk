package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdThreadIdPostIdInReplyToExtensionId{}

func TestNewGroupIdThreadIdPostIdInReplyToExtensionID(t *testing.T) {
	id := NewGroupIdThreadIdPostIdInReplyToExtensionID("groupId", "conversationThreadId", "postId", "extensionId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.ConversationThreadId != "conversationThreadId" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationThreadId'", id.ConversationThreadId, "conversationThreadId")
	}

	if id.PostId != "postId" {
		t.Fatalf("Expected %q but got %q for Segment 'PostId'", id.PostId, "postId")
	}

	if id.ExtensionId != "extensionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionId")
	}
}

func TestFormatGroupIdThreadIdPostIdInReplyToExtensionID(t *testing.T) {
	actual := NewGroupIdThreadIdPostIdInReplyToExtensionID("groupId", "conversationThreadId", "postId", "extensionId").ID()
	expected := "/groups/groupId/threads/conversationThreadId/posts/postId/inReplyTo/extensions/extensionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdThreadIdPostIdInReplyToExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdThreadIdPostIdInReplyToExtensionId
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
			Input: "/groups/groupId/threads",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/threads/conversationThreadId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/threads/conversationThreadId/posts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/threads/conversationThreadId/posts/postId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/threads/conversationThreadId/posts/postId/inReplyTo",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/threads/conversationThreadId/posts/postId/inReplyTo/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/threads/conversationThreadId/posts/postId/inReplyTo/extensions/extensionId",
			Expected: &GroupIdThreadIdPostIdInReplyToExtensionId{
				GroupId:              "groupId",
				ConversationThreadId: "conversationThreadId",
				PostId:               "postId",
				ExtensionId:          "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/threads/conversationThreadId/posts/postId/inReplyTo/extensions/extensionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdThreadIdPostIdInReplyToExtensionID(v.Input)
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

		if actual.PostId != v.Expected.PostId {
			t.Fatalf("Expected %q but got %q for PostId", v.Expected.PostId, actual.PostId)
		}

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestParseGroupIdThreadIdPostIdInReplyToExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdThreadIdPostIdInReplyToExtensionId
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
			Input: "/groups/groupId/threads",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tHrEaDs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/threads/conversationThreadId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tHrEaDs/cOnVeRsAtIoNtHrEaDiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/threads/conversationThreadId/posts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tHrEaDs/cOnVeRsAtIoNtHrEaDiD/pOsTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/threads/conversationThreadId/posts/postId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tHrEaDs/cOnVeRsAtIoNtHrEaDiD/pOsTs/pOsTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/threads/conversationThreadId/posts/postId/inReplyTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tHrEaDs/cOnVeRsAtIoNtHrEaDiD/pOsTs/pOsTiD/iNrEpLyTo",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/threads/conversationThreadId/posts/postId/inReplyTo/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tHrEaDs/cOnVeRsAtIoNtHrEaDiD/pOsTs/pOsTiD/iNrEpLyTo/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/threads/conversationThreadId/posts/postId/inReplyTo/extensions/extensionId",
			Expected: &GroupIdThreadIdPostIdInReplyToExtensionId{
				GroupId:              "groupId",
				ConversationThreadId: "conversationThreadId",
				PostId:               "postId",
				ExtensionId:          "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/threads/conversationThreadId/posts/postId/inReplyTo/extensions/extensionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tHrEaDs/cOnVeRsAtIoNtHrEaDiD/pOsTs/pOsTiD/iNrEpLyTo/eXtEnSiOnS/eXtEnSiOnId",
			Expected: &GroupIdThreadIdPostIdInReplyToExtensionId{
				GroupId:              "gRoUpId",
				ConversationThreadId: "cOnVeRsAtIoNtHrEaDiD",
				PostId:               "pOsTiD",
				ExtensionId:          "eXtEnSiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tHrEaDs/cOnVeRsAtIoNtHrEaDiD/pOsTs/pOsTiD/iNrEpLyTo/eXtEnSiOnS/eXtEnSiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdThreadIdPostIdInReplyToExtensionIDInsensitively(v.Input)
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

		if actual.PostId != v.Expected.PostId {
			t.Fatalf("Expected %q but got %q for PostId", v.Expected.PostId, actual.PostId)
		}

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestSegmentsForGroupIdThreadIdPostIdInReplyToExtensionId(t *testing.T) {
	segments := GroupIdThreadIdPostIdInReplyToExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdThreadIdPostIdInReplyToExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
