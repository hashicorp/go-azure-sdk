package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdThreadIdPostIdInReplyToMentionId{}

func TestNewGroupIdThreadIdPostIdInReplyToMentionID(t *testing.T) {
	id := NewGroupIdThreadIdPostIdInReplyToMentionID("groupId", "conversationThreadId", "postId", "mentionId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.ConversationThreadId != "conversationThreadId" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationThreadId'", id.ConversationThreadId, "conversationThreadId")
	}

	if id.PostId != "postId" {
		t.Fatalf("Expected %q but got %q for Segment 'PostId'", id.PostId, "postId")
	}

	if id.MentionId != "mentionId" {
		t.Fatalf("Expected %q but got %q for Segment 'MentionId'", id.MentionId, "mentionId")
	}
}

func TestFormatGroupIdThreadIdPostIdInReplyToMentionID(t *testing.T) {
	actual := NewGroupIdThreadIdPostIdInReplyToMentionID("groupId", "conversationThreadId", "postId", "mentionId").ID()
	expected := "/groups/groupId/threads/conversationThreadId/posts/postId/inReplyTo/mentions/mentionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdThreadIdPostIdInReplyToMentionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdThreadIdPostIdInReplyToMentionId
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
			Input: "/groups/groupId/threads/conversationThreadId/posts/postId/inReplyTo/mentions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/threads/conversationThreadId/posts/postId/inReplyTo/mentions/mentionId",
			Expected: &GroupIdThreadIdPostIdInReplyToMentionId{
				GroupId:              "groupId",
				ConversationThreadId: "conversationThreadId",
				PostId:               "postId",
				MentionId:            "mentionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/threads/conversationThreadId/posts/postId/inReplyTo/mentions/mentionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdThreadIdPostIdInReplyToMentionID(v.Input)
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

		if actual.MentionId != v.Expected.MentionId {
			t.Fatalf("Expected %q but got %q for MentionId", v.Expected.MentionId, actual.MentionId)
		}

	}
}

func TestParseGroupIdThreadIdPostIdInReplyToMentionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdThreadIdPostIdInReplyToMentionId
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
			Input: "/groups/groupId/threads/conversationThreadId/posts/postId/inReplyTo/mentions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tHrEaDs/cOnVeRsAtIoNtHrEaDiD/pOsTs/pOsTiD/iNrEpLyTo/mEnTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/threads/conversationThreadId/posts/postId/inReplyTo/mentions/mentionId",
			Expected: &GroupIdThreadIdPostIdInReplyToMentionId{
				GroupId:              "groupId",
				ConversationThreadId: "conversationThreadId",
				PostId:               "postId",
				MentionId:            "mentionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/threads/conversationThreadId/posts/postId/inReplyTo/mentions/mentionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tHrEaDs/cOnVeRsAtIoNtHrEaDiD/pOsTs/pOsTiD/iNrEpLyTo/mEnTiOnS/mEnTiOnId",
			Expected: &GroupIdThreadIdPostIdInReplyToMentionId{
				GroupId:              "gRoUpId",
				ConversationThreadId: "cOnVeRsAtIoNtHrEaDiD",
				PostId:               "pOsTiD",
				MentionId:            "mEnTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tHrEaDs/cOnVeRsAtIoNtHrEaDiD/pOsTs/pOsTiD/iNrEpLyTo/mEnTiOnS/mEnTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdThreadIdPostIdInReplyToMentionIDInsensitively(v.Input)
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

		if actual.MentionId != v.Expected.MentionId {
			t.Fatalf("Expected %q but got %q for MentionId", v.Expected.MentionId, actual.MentionId)
		}

	}
}

func TestSegmentsForGroupIdThreadIdPostIdInReplyToMentionId(t *testing.T) {
	segments := GroupIdThreadIdPostIdInReplyToMentionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdThreadIdPostIdInReplyToMentionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
