package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamChannelIdMessageIdHostedContentId{}

func TestNewGroupIdTeamChannelIdMessageIdHostedContentID(t *testing.T) {
	id := NewGroupIdTeamChannelIdMessageIdHostedContentID("groupId", "channelId", "chatMessageId", "chatMessageHostedContentId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.ChannelId != "channelId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChannelId'", id.ChannelId, "channelId")
	}

	if id.ChatMessageId != "chatMessageId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId'", id.ChatMessageId, "chatMessageId")
	}

	if id.ChatMessageHostedContentId != "chatMessageHostedContentId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageHostedContentId'", id.ChatMessageHostedContentId, "chatMessageHostedContentId")
	}
}

func TestFormatGroupIdTeamChannelIdMessageIdHostedContentID(t *testing.T) {
	actual := NewGroupIdTeamChannelIdMessageIdHostedContentID("groupId", "channelId", "chatMessageId", "chatMessageHostedContentId").ID()
	expected := "/groups/groupId/team/channels/channelId/messages/chatMessageId/hostedContents/chatMessageHostedContentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamChannelIdMessageIdHostedContentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamChannelIdMessageIdHostedContentId
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
			Input: "/groups/groupId/team/channels/channelId/messages/chatMessageId/hostedContents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/channels/channelId/messages/chatMessageId/hostedContents/chatMessageHostedContentId",
			Expected: &GroupIdTeamChannelIdMessageIdHostedContentId{
				GroupId:                    "groupId",
				ChannelId:                  "channelId",
				ChatMessageId:              "chatMessageId",
				ChatMessageHostedContentId: "chatMessageHostedContentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/channels/channelId/messages/chatMessageId/hostedContents/chatMessageHostedContentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamChannelIdMessageIdHostedContentID(v.Input)
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

		if actual.ChatMessageHostedContentId != v.Expected.ChatMessageHostedContentId {
			t.Fatalf("Expected %q but got %q for ChatMessageHostedContentId", v.Expected.ChatMessageHostedContentId, actual.ChatMessageHostedContentId)
		}

	}
}

func TestParseGroupIdTeamChannelIdMessageIdHostedContentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamChannelIdMessageIdHostedContentId
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
			Input: "/groups/groupId/team/channels/channelId/messages/chatMessageId/hostedContents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/mEsSaGeS/cHaTmEsSaGeId/hOsTeDcOnTeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/channels/channelId/messages/chatMessageId/hostedContents/chatMessageHostedContentId",
			Expected: &GroupIdTeamChannelIdMessageIdHostedContentId{
				GroupId:                    "groupId",
				ChannelId:                  "channelId",
				ChatMessageId:              "chatMessageId",
				ChatMessageHostedContentId: "chatMessageHostedContentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/channels/channelId/messages/chatMessageId/hostedContents/chatMessageHostedContentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/mEsSaGeS/cHaTmEsSaGeId/hOsTeDcOnTeNtS/cHaTmEsSaGeHoStEdCoNtEnTiD",
			Expected: &GroupIdTeamChannelIdMessageIdHostedContentId{
				GroupId:                    "gRoUpId",
				ChannelId:                  "cHaNnElId",
				ChatMessageId:              "cHaTmEsSaGeId",
				ChatMessageHostedContentId: "cHaTmEsSaGeHoStEdCoNtEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/mEsSaGeS/cHaTmEsSaGeId/hOsTeDcOnTeNtS/cHaTmEsSaGeHoStEdCoNtEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamChannelIdMessageIdHostedContentIDInsensitively(v.Input)
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

		if actual.ChatMessageHostedContentId != v.Expected.ChatMessageHostedContentId {
			t.Fatalf("Expected %q but got %q for ChatMessageHostedContentId", v.Expected.ChatMessageHostedContentId, actual.ChatMessageHostedContentId)
		}

	}
}

func TestSegmentsForGroupIdTeamChannelIdMessageIdHostedContentId(t *testing.T) {
	segments := GroupIdTeamChannelIdMessageIdHostedContentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamChannelIdMessageIdHostedContentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
