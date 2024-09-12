package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdConversationIdThreadIdPostIdInReplyToMentionId{}

func TestNewGroupIdConversationIdThreadIdPostIdInReplyToMentionID(t *testing.T) {
	id := NewGroupIdConversationIdThreadIdPostIdInReplyToMentionID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue", "mentionIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.ConversationId != "conversationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationId'", id.ConversationId, "conversationIdValue")
	}

	if id.ConversationThreadId != "conversationThreadIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationThreadId'", id.ConversationThreadId, "conversationThreadIdValue")
	}

	if id.PostId != "postIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PostId'", id.PostId, "postIdValue")
	}

	if id.MentionId != "mentionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MentionId'", id.MentionId, "mentionIdValue")
	}
}

func TestFormatGroupIdConversationIdThreadIdPostIdInReplyToMentionID(t *testing.T) {
	actual := NewGroupIdConversationIdThreadIdPostIdInReplyToMentionID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue", "mentionIdValue").ID()
	expected := "/groups/groupIdValue/conversations/conversationIdValue/threads/conversationThreadIdValue/posts/postIdValue/inReplyTo/mentions/mentionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdConversationIdThreadIdPostIdInReplyToMentionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdConversationIdThreadIdPostIdInReplyToMentionId
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
			// Incomplete URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue/threads",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue/threads/conversationThreadIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue/threads/conversationThreadIdValue/posts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue/threads/conversationThreadIdValue/posts/postIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue/threads/conversationThreadIdValue/posts/postIdValue/inReplyTo",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue/threads/conversationThreadIdValue/posts/postIdValue/inReplyTo/mentions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue/threads/conversationThreadIdValue/posts/postIdValue/inReplyTo/mentions/mentionIdValue",
			Expected: &GroupIdConversationIdThreadIdPostIdInReplyToMentionId{
				GroupId:              "groupIdValue",
				ConversationId:       "conversationIdValue",
				ConversationThreadId: "conversationThreadIdValue",
				PostId:               "postIdValue",
				MentionId:            "mentionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/conversations/conversationIdValue/threads/conversationThreadIdValue/posts/postIdValue/inReplyTo/mentions/mentionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdConversationIdThreadIdPostIdInReplyToMentionID(v.Input)
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

		if actual.MentionId != v.Expected.MentionId {
			t.Fatalf("Expected %q but got %q for MentionId", v.Expected.MentionId, actual.MentionId)
		}

	}
}

func TestParseGroupIdConversationIdThreadIdPostIdInReplyToMentionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdConversationIdThreadIdPostIdInReplyToMentionId
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
			// Incomplete URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cOnVeRsAtIoNs/cOnVeRsAtIoNiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue/threads",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cOnVeRsAtIoNs/cOnVeRsAtIoNiDvAlUe/tHrEaDs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue/threads/conversationThreadIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cOnVeRsAtIoNs/cOnVeRsAtIoNiDvAlUe/tHrEaDs/cOnVeRsAtIoNtHrEaDiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue/threads/conversationThreadIdValue/posts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cOnVeRsAtIoNs/cOnVeRsAtIoNiDvAlUe/tHrEaDs/cOnVeRsAtIoNtHrEaDiDvAlUe/pOsTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue/threads/conversationThreadIdValue/posts/postIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cOnVeRsAtIoNs/cOnVeRsAtIoNiDvAlUe/tHrEaDs/cOnVeRsAtIoNtHrEaDiDvAlUe/pOsTs/pOsTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue/threads/conversationThreadIdValue/posts/postIdValue/inReplyTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cOnVeRsAtIoNs/cOnVeRsAtIoNiDvAlUe/tHrEaDs/cOnVeRsAtIoNtHrEaDiDvAlUe/pOsTs/pOsTiDvAlUe/iNrEpLyTo",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue/threads/conversationThreadIdValue/posts/postIdValue/inReplyTo/mentions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cOnVeRsAtIoNs/cOnVeRsAtIoNiDvAlUe/tHrEaDs/cOnVeRsAtIoNtHrEaDiDvAlUe/pOsTs/pOsTiDvAlUe/iNrEpLyTo/mEnTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/conversations/conversationIdValue/threads/conversationThreadIdValue/posts/postIdValue/inReplyTo/mentions/mentionIdValue",
			Expected: &GroupIdConversationIdThreadIdPostIdInReplyToMentionId{
				GroupId:              "groupIdValue",
				ConversationId:       "conversationIdValue",
				ConversationThreadId: "conversationThreadIdValue",
				PostId:               "postIdValue",
				MentionId:            "mentionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/conversations/conversationIdValue/threads/conversationThreadIdValue/posts/postIdValue/inReplyTo/mentions/mentionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cOnVeRsAtIoNs/cOnVeRsAtIoNiDvAlUe/tHrEaDs/cOnVeRsAtIoNtHrEaDiDvAlUe/pOsTs/pOsTiDvAlUe/iNrEpLyTo/mEnTiOnS/mEnTiOnIdVaLuE",
			Expected: &GroupIdConversationIdThreadIdPostIdInReplyToMentionId{
				GroupId:              "gRoUpIdVaLuE",
				ConversationId:       "cOnVeRsAtIoNiDvAlUe",
				ConversationThreadId: "cOnVeRsAtIoNtHrEaDiDvAlUe",
				PostId:               "pOsTiDvAlUe",
				MentionId:            "mEnTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cOnVeRsAtIoNs/cOnVeRsAtIoNiDvAlUe/tHrEaDs/cOnVeRsAtIoNtHrEaDiDvAlUe/pOsTs/pOsTiDvAlUe/iNrEpLyTo/mEnTiOnS/mEnTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdConversationIdThreadIdPostIdInReplyToMentionIDInsensitively(v.Input)
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

		if actual.MentionId != v.Expected.MentionId {
			t.Fatalf("Expected %q but got %q for MentionId", v.Expected.MentionId, actual.MentionId)
		}

	}
}

func TestSegmentsForGroupIdConversationIdThreadIdPostIdInReplyToMentionId(t *testing.T) {
	segments := GroupIdConversationIdThreadIdPostIdInReplyToMentionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdConversationIdThreadIdPostIdInReplyToMentionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
