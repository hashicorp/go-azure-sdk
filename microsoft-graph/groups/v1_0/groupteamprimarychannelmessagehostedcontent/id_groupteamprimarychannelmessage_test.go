package groupteamprimarychannelmessagehostedcontent

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamPrimaryChannelMessageId{}

func TestNewGroupTeamPrimaryChannelMessageID(t *testing.T) {
	id := NewGroupTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.ChatMessageId != "chatMessageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId'", id.ChatMessageId, "chatMessageIdValue")
	}
}

func TestFormatGroupTeamPrimaryChannelMessageID(t *testing.T) {
	actual := NewGroupTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue").ID()
	expected := "/groups/groupIdValue/team/primaryChannel/messages/chatMessageIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupTeamPrimaryChannelMessageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamPrimaryChannelMessageId
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
			// Valid URI
			Input: "/groups/groupIdValue/team/primaryChannel/messages/chatMessageIdValue",
			Expected: &GroupTeamPrimaryChannelMessageId{
				GroupId:       "groupIdValue",
				ChatMessageId: "chatMessageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/primaryChannel/messages/chatMessageIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamPrimaryChannelMessageID(v.Input)
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

	}
}

func TestParseGroupTeamPrimaryChannelMessageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamPrimaryChannelMessageId
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
			// Valid URI
			Input: "/groups/groupIdValue/team/primaryChannel/messages/chatMessageIdValue",
			Expected: &GroupTeamPrimaryChannelMessageId{
				GroupId:       "groupIdValue",
				ChatMessageId: "chatMessageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/primaryChannel/messages/chatMessageIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE",
			Expected: &GroupTeamPrimaryChannelMessageId{
				GroupId:       "gRoUpIdVaLuE",
				ChatMessageId: "cHaTmEsSaGeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/mEsSaGeS/cHaTmEsSaGeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamPrimaryChannelMessageIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForGroupTeamPrimaryChannelMessageId(t *testing.T) {
	segments := GroupTeamPrimaryChannelMessageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupTeamPrimaryChannelMessageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
