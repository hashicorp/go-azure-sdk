package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPrimaryChannelTabId{}

func TestNewGroupIdTeamPrimaryChannelTabID(t *testing.T) {
	id := NewGroupIdTeamPrimaryChannelTabID("groupIdValue", "teamsTabIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.TeamsTabId != "teamsTabIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamsTabId'", id.TeamsTabId, "teamsTabIdValue")
	}
}

func TestFormatGroupIdTeamPrimaryChannelTabID(t *testing.T) {
	actual := NewGroupIdTeamPrimaryChannelTabID("groupIdValue", "teamsTabIdValue").ID()
	expected := "/groups/groupIdValue/team/primaryChannel/tabs/teamsTabIdValue"
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
			Input: "/groups/groupIdValue/team/primaryChannel/tabs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/primaryChannel/tabs/teamsTabIdValue",
			Expected: &GroupIdTeamPrimaryChannelTabId{
				GroupId:    "groupIdValue",
				TeamsTabId: "teamsTabIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/primaryChannel/tabs/teamsTabIdValue/extra",
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
			Input: "/groups/groupIdValue/team/primaryChannel/tabs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/tAbS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/primaryChannel/tabs/teamsTabIdValue",
			Expected: &GroupIdTeamPrimaryChannelTabId{
				GroupId:    "groupIdValue",
				TeamsTabId: "teamsTabIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/primaryChannel/tabs/teamsTabIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/tAbS/tEaMsTaBiDvAlUe",
			Expected: &GroupIdTeamPrimaryChannelTabId{
				GroupId:    "gRoUpIdVaLuE",
				TeamsTabId: "tEaMsTaBiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/pRiMaRyChAnNeL/tAbS/tEaMsTaBiDvAlUe/extra",
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
