package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamTagId{}

func TestNewGroupIdTeamTagID(t *testing.T) {
	id := NewGroupIdTeamTagID("groupId", "teamworkTagId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.TeamworkTagId != "teamworkTagId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamworkTagId'", id.TeamworkTagId, "teamworkTagId")
	}
}

func TestFormatGroupIdTeamTagID(t *testing.T) {
	actual := NewGroupIdTeamTagID("groupId", "teamworkTagId").ID()
	expected := "/groups/groupId/team/tags/teamworkTagId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamTagID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamTagId
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
			// Valid URI
			Input: "/groups/groupId/team/tags/teamworkTagId",
			Expected: &GroupIdTeamTagId{
				GroupId:       "groupId",
				TeamworkTagId: "teamworkTagId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/tags/teamworkTagId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamTagID(v.Input)
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

	}
}

func TestParseGroupIdTeamTagIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamTagId
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
			// Valid URI
			Input: "/groups/groupId/team/tags/teamworkTagId",
			Expected: &GroupIdTeamTagId{
				GroupId:       "groupId",
				TeamworkTagId: "teamworkTagId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/tags/teamworkTagId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/tAgS/tEaMwOrKtAgId",
			Expected: &GroupIdTeamTagId{
				GroupId:       "gRoUpId",
				TeamworkTagId: "tEaMwOrKtAgId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/tAgS/tEaMwOrKtAgId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamTagIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForGroupIdTeamTagId(t *testing.T) {
	segments := GroupIdTeamTagId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamTagId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
