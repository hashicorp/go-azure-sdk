package groupteamprimarychannelmessagereplyhostedcontent

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamPrimaryChannelMessageReplyId{}

func TestNewGroupTeamPrimaryChannelMessageReplyID(t *testing.T) {
	id := NewGroupTeamPrimaryChannelMessageReplyID("groupIdValue", "chatMessageIdValue", "chatMessageId1Value")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.ChatMessageId != "chatMessageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId'", id.ChatMessageId, "chatMessageIdValue")
	}

	if id.ChatMessageId1 != "chatMessageId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId1'", id.ChatMessageId1, "chatMessageId1Value")
	}
}

func TestFormatGroupTeamPrimaryChannelMessageReplyID(t *testing.T) {
	actual := NewGroupTeamPrimaryChannelMessageReplyID("groupIdValue", "chatMessageIdValue", "chatMessageId1Value").ID()
	expected := "/groups/groupIdValue/team/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupTeamPrimaryChannelMessageReplyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamPrimaryChannelMessageReplyId
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
			Input: "/groups/groupIdValue/team",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/primaryChannel",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/primaryChannel/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/primaryChannel/messages/chatMessageIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/primaryChannel/messages/chatMessageIdValue/replies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value",
			Expected: &GroupTeamPrimaryChannelMessageReplyId{
				GroupId:        "groupIdValue",
				ChatMessageId:  "chatMessageIdValue",
				ChatMessageId1: "chatMessageId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamPrimaryChannelMessageReplyID(v.Input)
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

		if actual.ChatMessageId != v.Expected.ChatMessageId {
			t.Fatalf("Expected %q but got %q for ChatMessageId", v.Expected.ChatMessageId, actual.ChatMessageId)
		}

		if actual.ChatMessageId1 != v.Expected.ChatMessageId1 {
			t.Fatalf("Expected %q but got %q for ChatMessageId1", v.Expected.ChatMessageId1, actual.ChatMessageId1)
		}

	}
}

func TestParseGroupTeamPrimaryChannelMessageReplyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamPrimaryChannelMessageReplyId
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
			Input: "/groups/groupIdValue/team",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/primaryChannel",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/primaryChannel/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/primaryChannel/messages/chatMessageIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/primaryChannel/messages/chatMessageIdValue/replies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value",
			Expected: &GroupTeamPrimaryChannelMessageReplyId{
				GroupId:        "groupIdValue",
				ChatMessageId:  "chatMessageIdValue",
				ChatMessageId1: "chatMessageId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/primaryChannel/messages/chatMessageIdValue/replies/chatMessageId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe",
			Expected: &GroupTeamPrimaryChannelMessageReplyId{
				GroupId:        "gRoUpIdVaLuE",
				ChatMessageId:  "cHaTmEsSaGeIdVaLuE",
				ChatMessageId1: "cHaTmEsSaGeId1vAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamPrimaryChannelMessageReplyIDInsensitively(v.Input)
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

		if actual.ChatMessageId != v.Expected.ChatMessageId {
			t.Fatalf("Expected %q but got %q for ChatMessageId", v.Expected.ChatMessageId, actual.ChatMessageId)
		}

		if actual.ChatMessageId1 != v.Expected.ChatMessageId1 {
			t.Fatalf("Expected %q but got %q for ChatMessageId1", v.Expected.ChatMessageId1, actual.ChatMessageId1)
		}

	}
}

func TestSegmentsForGroupTeamPrimaryChannelMessageReplyId(t *testing.T) {
	segments := GroupTeamPrimaryChannelMessageReplyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupTeamPrimaryChannelMessageReplyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
