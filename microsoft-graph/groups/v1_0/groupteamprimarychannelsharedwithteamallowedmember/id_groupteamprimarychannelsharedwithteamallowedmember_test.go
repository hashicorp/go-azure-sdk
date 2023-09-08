package groupteamprimarychannelsharedwithteamallowedmember

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId{}

func TestNewGroupTeamPrimaryChannelSharedWithTeamAllowedMemberID(t *testing.T) {
	id := NewGroupTeamPrimaryChannelSharedWithTeamAllowedMemberID("groupIdValue", "sharedWithChannelTeamInfoIdValue", "conversationMemberIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SharedWithChannelTeamInfoId != "sharedWithChannelTeamInfoIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SharedWithChannelTeamInfoId'", id.SharedWithChannelTeamInfoId, "sharedWithChannelTeamInfoIdValue")
	}

	if id.ConversationMemberId != "conversationMemberIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationMemberId'", id.ConversationMemberId, "conversationMemberIdValue")
	}
}

func TestFormatGroupTeamPrimaryChannelSharedWithTeamAllowedMemberID(t *testing.T) {
	actual := NewGroupTeamPrimaryChannelSharedWithTeamAllowedMemberID("groupIdValue", "sharedWithChannelTeamInfoIdValue", "conversationMemberIdValue").ID()
	expected := "/groups/groupIdValue/team/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue/allowedMembers/conversationMemberIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupTeamPrimaryChannelSharedWithTeamAllowedMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId
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
			Input: "/groups/groupIdValue/team/primaryChannel/sharedWithTeams",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue/allowedMembers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue/allowedMembers/conversationMemberIdValue",
			Expected: &GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId{
				GroupId:                     "groupIdValue",
				SharedWithChannelTeamInfoId: "sharedWithChannelTeamInfoIdValue",
				ConversationMemberId:        "conversationMemberIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue/allowedMembers/conversationMemberIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamPrimaryChannelSharedWithTeamAllowedMemberID(v.Input)
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

		if actual.SharedWithChannelTeamInfoId != v.Expected.SharedWithChannelTeamInfoId {
			t.Fatalf("Expected %q but got %q for SharedWithChannelTeamInfoId", v.Expected.SharedWithChannelTeamInfoId, actual.SharedWithChannelTeamInfoId)
		}

		if actual.ConversationMemberId != v.Expected.ConversationMemberId {
			t.Fatalf("Expected %q but got %q for ConversationMemberId", v.Expected.ConversationMemberId, actual.ConversationMemberId)
		}

	}
}

func TestParseGroupTeamPrimaryChannelSharedWithTeamAllowedMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId
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
			Input: "/groups/groupIdValue/team/primaryChannel/sharedWithTeams",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/sHaReDwItHtEaMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue/allowedMembers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoIdVaLuE/aLlOwEdMeMbErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue/allowedMembers/conversationMemberIdValue",
			Expected: &GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId{
				GroupId:                     "groupIdValue",
				SharedWithChannelTeamInfoId: "sharedWithChannelTeamInfoIdValue",
				ConversationMemberId:        "conversationMemberIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue/allowedMembers/conversationMemberIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoIdVaLuE/aLlOwEdMeMbErS/cOnVeRsAtIoNmEmBeRiDvAlUe",
			Expected: &GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId{
				GroupId:                     "gRoUpIdVaLuE",
				SharedWithChannelTeamInfoId: "sHaReDwItHcHaNnElTeAmInFoIdVaLuE",
				ConversationMemberId:        "cOnVeRsAtIoNmEmBeRiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoIdVaLuE/aLlOwEdMeMbErS/cOnVeRsAtIoNmEmBeRiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamPrimaryChannelSharedWithTeamAllowedMemberIDInsensitively(v.Input)
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

		if actual.SharedWithChannelTeamInfoId != v.Expected.SharedWithChannelTeamInfoId {
			t.Fatalf("Expected %q but got %q for SharedWithChannelTeamInfoId", v.Expected.SharedWithChannelTeamInfoId, actual.SharedWithChannelTeamInfoId)
		}

		if actual.ConversationMemberId != v.Expected.ConversationMemberId {
			t.Fatalf("Expected %q but got %q for ConversationMemberId", v.Expected.ConversationMemberId, actual.ConversationMemberId)
		}

	}
}

func TestSegmentsForGroupTeamPrimaryChannelSharedWithTeamAllowedMemberId(t *testing.T) {
	segments := GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
