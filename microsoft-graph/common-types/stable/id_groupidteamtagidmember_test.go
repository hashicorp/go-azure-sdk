package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamTagIdMemberId{}

func TestNewGroupIdTeamTagIdMemberID(t *testing.T) {
	id := NewGroupIdTeamTagIdMemberID("groupId", "teamworkTagId", "teamworkTagMemberId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.TeamworkTagId != "teamworkTagId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamworkTagId'", id.TeamworkTagId, "teamworkTagId")
	}

	if id.TeamworkTagMemberId != "teamworkTagMemberId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamworkTagMemberId'", id.TeamworkTagMemberId, "teamworkTagMemberId")
	}
}

func TestFormatGroupIdTeamTagIdMemberID(t *testing.T) {
	actual := NewGroupIdTeamTagIdMemberID("groupId", "teamworkTagId", "teamworkTagMemberId").ID()
	expected := "/groups/groupId/team/tags/teamworkTagId/members/teamworkTagMemberId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamTagIdMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamTagIdMemberId
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
			Input: "/groups/groupId/team/tags",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/tags/teamworkTagId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/tags/teamworkTagId/members",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/tags/teamworkTagId/members/teamworkTagMemberId",
			Expected: &GroupIdTeamTagIdMemberId{
				GroupId:             "groupId",
				TeamworkTagId:       "teamworkTagId",
				TeamworkTagMemberId: "teamworkTagMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/tags/teamworkTagId/members/teamworkTagMemberId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamTagIdMemberID(v.Input)
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

		if actual.TeamworkTagId != v.Expected.TeamworkTagId {
			t.Fatalf("Expected %q but got %q for TeamworkTagId", v.Expected.TeamworkTagId, actual.TeamworkTagId)
		}

		if actual.TeamworkTagMemberId != v.Expected.TeamworkTagMemberId {
			t.Fatalf("Expected %q but got %q for TeamworkTagMemberId", v.Expected.TeamworkTagMemberId, actual.TeamworkTagMemberId)
		}

	}
}

func TestParseGroupIdTeamTagIdMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamTagIdMemberId
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
			Input: "/groups/groupId/team/tags",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/tAgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/tags/teamworkTagId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/tAgS/tEaMwOrKtAgId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/tags/teamworkTagId/members",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/tAgS/tEaMwOrKtAgId/mEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/tags/teamworkTagId/members/teamworkTagMemberId",
			Expected: &GroupIdTeamTagIdMemberId{
				GroupId:             "groupId",
				TeamworkTagId:       "teamworkTagId",
				TeamworkTagMemberId: "teamworkTagMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/tags/teamworkTagId/members/teamworkTagMemberId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/tAgS/tEaMwOrKtAgId/mEmBeRs/tEaMwOrKtAgMeMbErId",
			Expected: &GroupIdTeamTagIdMemberId{
				GroupId:             "gRoUpId",
				TeamworkTagId:       "tEaMwOrKtAgId",
				TeamworkTagMemberId: "tEaMwOrKtAgMeMbErId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/tAgS/tEaMwOrKtAgId/mEmBeRs/tEaMwOrKtAgMeMbErId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamTagIdMemberIDInsensitively(v.Input)
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

		if actual.TeamworkTagId != v.Expected.TeamworkTagId {
			t.Fatalf("Expected %q but got %q for TeamworkTagId", v.Expected.TeamworkTagId, actual.TeamworkTagId)
		}

		if actual.TeamworkTagMemberId != v.Expected.TeamworkTagMemberId {
			t.Fatalf("Expected %q but got %q for TeamworkTagMemberId", v.Expected.TeamworkTagMemberId, actual.TeamworkTagMemberId)
		}

	}
}

func TestSegmentsForGroupIdTeamTagIdMemberId(t *testing.T) {
	segments := GroupIdTeamTagIdMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamTagIdMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
