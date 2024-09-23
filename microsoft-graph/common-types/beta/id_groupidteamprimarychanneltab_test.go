package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPrimaryChannelTabId{}

func TestNewGroupIdTeamPrimaryChannelTabID(t *testing.T) {
	id := NewGroupIdTeamPrimaryChannelTabID("groupId", "teamsTabId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.TeamsTabId != "teamsTabId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamsTabId'", id.TeamsTabId, "teamsTabId")
	}
}

func TestFormatGroupIdTeamPrimaryChannelTabID(t *testing.T) {
	actual := NewGroupIdTeamPrimaryChannelTabID("groupId", "teamsTabId").ID()
	expected := "/groups/groupId/team/primaryChannel/tabs/teamsTabId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamPrimaryChannelTabID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamPrimaryChannelTabId
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
			Input: "/groups/groupId/team/primaryChannel",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/primaryChannel/tabs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/primaryChannel/tabs/teamsTabId",
			Expected: &GroupIdTeamPrimaryChannelTabId{
				GroupId:    "groupId",
				TeamsTabId: "teamsTabId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/primaryChannel/tabs/teamsTabId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamPrimaryChannelTabID(v.Input)
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

		if actual.TeamsTabId != v.Expected.TeamsTabId {
			t.Fatalf("Expected %q but got %q for TeamsTabId", v.Expected.TeamsTabId, actual.TeamsTabId)
		}

	}
}

func TestParseGroupIdTeamPrimaryChannelTabIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamPrimaryChannelTabId
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
			Input: "/groups/groupId/team/primaryChannel",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/pRiMaRyChAnNeL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/primaryChannel/tabs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/pRiMaRyChAnNeL/tAbS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/primaryChannel/tabs/teamsTabId",
			Expected: &GroupIdTeamPrimaryChannelTabId{
				GroupId:    "groupId",
				TeamsTabId: "teamsTabId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/primaryChannel/tabs/teamsTabId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/pRiMaRyChAnNeL/tAbS/tEaMsTaBiD",
			Expected: &GroupIdTeamPrimaryChannelTabId{
				GroupId:    "gRoUpId",
				TeamsTabId: "tEaMsTaBiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/pRiMaRyChAnNeL/tAbS/tEaMsTaBiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamPrimaryChannelTabIDInsensitively(v.Input)
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

		if actual.TeamsTabId != v.Expected.TeamsTabId {
			t.Fatalf("Expected %q but got %q for TeamsTabId", v.Expected.TeamsTabId, actual.TeamsTabId)
		}

	}
}

func TestSegmentsForGroupIdTeamPrimaryChannelTabId(t *testing.T) {
	segments := GroupIdTeamPrimaryChannelTabId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamPrimaryChannelTabId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
