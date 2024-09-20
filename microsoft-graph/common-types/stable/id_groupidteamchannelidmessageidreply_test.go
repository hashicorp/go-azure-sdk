package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamChannelIdMessageIdReplyId{}

func TestNewGroupIdTeamChannelIdMessageIdReplyID(t *testing.T) {
	id := NewGroupIdTeamChannelIdMessageIdReplyID("groupId", "channelId", "chatMessageId", "chatMessageId1")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.ChannelId != "channelId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChannelId'", id.ChannelId, "channelId")
	}

	if id.ChatMessageId != "chatMessageId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId'", id.ChatMessageId, "chatMessageId")
	}

	if id.ChatMessageId1 != "chatMessageId1" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId1'", id.ChatMessageId1, "chatMessageId1")
	}
}

func TestFormatGroupIdTeamChannelIdMessageIdReplyID(t *testing.T) {
	actual := NewGroupIdTeamChannelIdMessageIdReplyID("groupId", "channelId", "chatMessageId", "chatMessageId1").ID()
	expected := "/groups/groupId/team/channels/channelId/messages/chatMessageId/replies/chatMessageId1"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamChannelIdMessageIdReplyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamChannelIdMessageIdReplyId
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
			Input: "/groups/groupId/team",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/messages/chatMessageId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/messages/chatMessageId/replies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/channels/channelId/messages/chatMessageId/replies/chatMessageId1",
			Expected: &GroupIdTeamChannelIdMessageIdReplyId{
				GroupId:        "groupId",
				ChannelId:      "channelId",
				ChatMessageId:  "chatMessageId",
				ChatMessageId1: "chatMessageId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/channels/channelId/messages/chatMessageId/replies/chatMessageId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamChannelIdMessageIdReplyID(v.Input)
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

		if actual.ChannelId != v.Expected.ChannelId {
			t.Fatalf("Expected %q but got %q for ChannelId", v.Expected.ChannelId, actual.ChannelId)
		}

		if actual.ChatMessageId != v.Expected.ChatMessageId {
			t.Fatalf("Expected %q but got %q for ChatMessageId", v.Expected.ChatMessageId, actual.ChatMessageId)
		}

		if actual.ChatMessageId1 != v.Expected.ChatMessageId1 {
			t.Fatalf("Expected %q but got %q for ChatMessageId1", v.Expected.ChatMessageId1, actual.ChatMessageId1)
		}

	}
}

func TestParseGroupIdTeamChannelIdMessageIdReplyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamChannelIdMessageIdReplyId
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
			Input: "/groups/groupId/team",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/messages/chatMessageId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/mEsSaGeS/cHaTmEsSaGeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/messages/chatMessageId/replies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/mEsSaGeS/cHaTmEsSaGeId/rEpLiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/channels/channelId/messages/chatMessageId/replies/chatMessageId1",
			Expected: &GroupIdTeamChannelIdMessageIdReplyId{
				GroupId:        "groupId",
				ChannelId:      "channelId",
				ChatMessageId:  "chatMessageId",
				ChatMessageId1: "chatMessageId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/channels/channelId/messages/chatMessageId/replies/chatMessageId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/mEsSaGeS/cHaTmEsSaGeId/rEpLiEs/cHaTmEsSaGeId1",
			Expected: &GroupIdTeamChannelIdMessageIdReplyId{
				GroupId:        "gRoUpId",
				ChannelId:      "cHaNnElId",
				ChatMessageId:  "cHaTmEsSaGeId",
				ChatMessageId1: "cHaTmEsSaGeId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/mEsSaGeS/cHaTmEsSaGeId/rEpLiEs/cHaTmEsSaGeId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamChannelIdMessageIdReplyIDInsensitively(v.Input)
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

		if actual.ChannelId != v.Expected.ChannelId {
			t.Fatalf("Expected %q but got %q for ChannelId", v.Expected.ChannelId, actual.ChannelId)
		}

		if actual.ChatMessageId != v.Expected.ChatMessageId {
			t.Fatalf("Expected %q but got %q for ChatMessageId", v.Expected.ChatMessageId, actual.ChatMessageId)
		}

		if actual.ChatMessageId1 != v.Expected.ChatMessageId1 {
			t.Fatalf("Expected %q but got %q for ChatMessageId1", v.Expected.ChatMessageId1, actual.ChatMessageId1)
		}

	}
}

func TestSegmentsForGroupIdTeamChannelIdMessageIdReplyId(t *testing.T) {
	segments := GroupIdTeamChannelIdMessageIdReplyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamChannelIdMessageIdReplyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
