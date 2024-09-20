package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId{}

func TestNewGroupIdConversationIdThreadIdPostIdInReplyToAttachmentID(t *testing.T) {
	id := NewGroupIdConversationIdThreadIdPostIdInReplyToAttachmentID("groupId", "conversationId", "conversationThreadId", "postId", "attachmentId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.ConversationId != "conversationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationId'", id.ConversationId, "conversationId")
	}

	if id.ConversationThreadId != "conversationThreadId" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationThreadId'", id.ConversationThreadId, "conversationThreadId")
	}

	if id.PostId != "postId" {
		t.Fatalf("Expected %q but got %q for Segment 'PostId'", id.PostId, "postId")
	}

	if id.AttachmentId != "attachmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentId'", id.AttachmentId, "attachmentId")
	}
}

func TestFormatGroupIdConversationIdThreadIdPostIdInReplyToAttachmentID(t *testing.T) {
	actual := NewGroupIdConversationIdThreadIdPostIdInReplyToAttachmentID("groupId", "conversationId", "conversationThreadId", "postId", "attachmentId").ID()
	expected := "/groups/groupId/conversations/conversationId/threads/conversationThreadId/posts/postId/inReplyTo/attachments/attachmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdConversationIdThreadIdPostIdInReplyToAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId
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
			// Incomplete URI
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId/posts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId/posts/postId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId/posts/postId/inReplyTo",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId/posts/postId/inReplyTo/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId/posts/postId/inReplyTo/attachments/attachmentId",
			Expected: &GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId{
				GroupId:              "groupId",
				ConversationId:       "conversationId",
				ConversationThreadId: "conversationThreadId",
				PostId:               "postId",
				AttachmentId:         "attachmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId/posts/postId/inReplyTo/attachments/attachmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdConversationIdThreadIdPostIdInReplyToAttachmentID(v.Input)
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

		if actual.PostId != v.Expected.PostId {
			t.Fatalf("Expected %q but got %q for PostId", v.Expected.PostId, actual.PostId)
		}

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestParseGroupIdConversationIdThreadIdPostIdInReplyToAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId
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
			// Incomplete URI
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cOnVeRsAtIoNs/cOnVeRsAtIoNiD/tHrEaDs/cOnVeRsAtIoNtHrEaDiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId/posts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cOnVeRsAtIoNs/cOnVeRsAtIoNiD/tHrEaDs/cOnVeRsAtIoNtHrEaDiD/pOsTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId/posts/postId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cOnVeRsAtIoNs/cOnVeRsAtIoNiD/tHrEaDs/cOnVeRsAtIoNtHrEaDiD/pOsTs/pOsTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId/posts/postId/inReplyTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cOnVeRsAtIoNs/cOnVeRsAtIoNiD/tHrEaDs/cOnVeRsAtIoNtHrEaDiD/pOsTs/pOsTiD/iNrEpLyTo",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId/posts/postId/inReplyTo/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cOnVeRsAtIoNs/cOnVeRsAtIoNiD/tHrEaDs/cOnVeRsAtIoNtHrEaDiD/pOsTs/pOsTiD/iNrEpLyTo/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId/posts/postId/inReplyTo/attachments/attachmentId",
			Expected: &GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId{
				GroupId:              "groupId",
				ConversationId:       "conversationId",
				ConversationThreadId: "conversationThreadId",
				PostId:               "postId",
				AttachmentId:         "attachmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/conversations/conversationId/threads/conversationThreadId/posts/postId/inReplyTo/attachments/attachmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cOnVeRsAtIoNs/cOnVeRsAtIoNiD/tHrEaDs/cOnVeRsAtIoNtHrEaDiD/pOsTs/pOsTiD/iNrEpLyTo/aTtAcHmEnTs/aTtAcHmEnTiD",
			Expected: &GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId{
				GroupId:              "gRoUpId",
				ConversationId:       "cOnVeRsAtIoNiD",
				ConversationThreadId: "cOnVeRsAtIoNtHrEaDiD",
				PostId:               "pOsTiD",
				AttachmentId:         "aTtAcHmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cOnVeRsAtIoNs/cOnVeRsAtIoNiD/tHrEaDs/cOnVeRsAtIoNtHrEaDiD/pOsTs/pOsTiD/iNrEpLyTo/aTtAcHmEnTs/aTtAcHmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdConversationIdThreadIdPostIdInReplyToAttachmentIDInsensitively(v.Input)
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

		if actual.PostId != v.Expected.PostId {
			t.Fatalf("Expected %q but got %q for PostId", v.Expected.PostId, actual.PostId)
		}

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestSegmentsForGroupIdConversationIdThreadIdPostIdInReplyToAttachmentId(t *testing.T) {
	segments := GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
