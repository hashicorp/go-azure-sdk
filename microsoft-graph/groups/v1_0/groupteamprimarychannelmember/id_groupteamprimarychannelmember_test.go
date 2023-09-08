package groupteamprimarychannelmember

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamPrimaryChannelMemberId{}

func TestNewGroupTeamPrimaryChannelMemberID(t *testing.T) {
	id := NewGroupTeamPrimaryChannelMemberID("groupIdValue", "conversationMemberIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.ConversationMemberId != "conversationMemberIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationMemberId'", id.ConversationMemberId, "conversationMemberIdValue")
	}
}

func TestFormatGroupTeamPrimaryChannelMemberID(t *testing.T) {
	actual := NewGroupTeamPrimaryChannelMemberID("groupIdValue", "conversationMemberIdValue").ID()
	expected := "/groups/groupIdValue/team/primaryChannel/members/conversationMemberIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupTeamPrimaryChannelMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamPrimaryChannelMemberId
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
			Input: "/groups/groupIdValue/team/primaryChannel/members",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/primaryChannel/members/conversationMemberIdValue",
			Expected: &GroupTeamPrimaryChannelMemberId{
				GroupId:              "groupIdValue",
				ConversationMemberId: "conversationMemberIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/primaryChannel/members/conversationMemberIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamPrimaryChannelMemberID(v.Input)
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

		if actual.ConversationMemberId != v.Expected.ConversationMemberId {
			t.Fatalf("Expected %q but got %q for ConversationMemberId", v.Expected.ConversationMemberId, actual.ConversationMemberId)
		}

	}
}

func TestParseGroupTeamPrimaryChannelMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamPrimaryChannelMemberId
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
			Input: "/groups/groupIdValue/team/primaryChannel/members",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/mEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/primaryChannel/members/conversationMemberIdValue",
			Expected: &GroupTeamPrimaryChannelMemberId{
				GroupId:              "groupIdValue",
				ConversationMemberId: "conversationMemberIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/primaryChannel/members/conversationMemberIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/mEmBeRs/cOnVeRsAtIoNmEmBeRiDvAlUe",
			Expected: &GroupTeamPrimaryChannelMemberId{
				GroupId:              "gRoUpIdVaLuE",
				ConversationMemberId: "cOnVeRsAtIoNmEmBeRiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/mEmBeRs/cOnVeRsAtIoNmEmBeRiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamPrimaryChannelMemberIDInsensitively(v.Input)
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

		if actual.ConversationMemberId != v.Expected.ConversationMemberId {
			t.Fatalf("Expected %q but got %q for ConversationMemberId", v.Expected.ConversationMemberId, actual.ConversationMemberId)
		}

	}
}

func TestSegmentsForGroupTeamPrimaryChannelMemberId(t *testing.T) {
	segments := GroupTeamPrimaryChannelMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupTeamPrimaryChannelMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
