package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamPrimaryChannelPlannerPlanId{}

func TestNewGroupIdTeamPrimaryChannelPlannerPlanID(t *testing.T) {
	id := NewGroupIdTeamPrimaryChannelPlannerPlanID("groupId", "plannerPlanId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.PlannerPlanId != "plannerPlanId" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerPlanId'", id.PlannerPlanId, "plannerPlanId")
	}
}

func TestFormatGroupIdTeamPrimaryChannelPlannerPlanID(t *testing.T) {
	actual := NewGroupIdTeamPrimaryChannelPlannerPlanID("groupId", "plannerPlanId").ID()
	expected := "/groups/groupId/team/primaryChannel/planner/plans/plannerPlanId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamPrimaryChannelPlannerPlanID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamPrimaryChannelPlannerPlanId
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
			Input: "/groups/groupId/team/primaryChannel/planner",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/primaryChannel/planner/plans",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/primaryChannel/planner/plans/plannerPlanId",
			Expected: &GroupIdTeamPrimaryChannelPlannerPlanId{
				GroupId:       "groupId",
				PlannerPlanId: "plannerPlanId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/primaryChannel/planner/plans/plannerPlanId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamPrimaryChannelPlannerPlanID(v.Input)
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

		if actual.PlannerPlanId != v.Expected.PlannerPlanId {
			t.Fatalf("Expected %q but got %q for PlannerPlanId", v.Expected.PlannerPlanId, actual.PlannerPlanId)
		}

	}
}

func TestParseGroupIdTeamPrimaryChannelPlannerPlanIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamPrimaryChannelPlannerPlanId
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
			Input: "/groups/groupId/team/primaryChannel/planner",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/pRiMaRyChAnNeL/pLaNnEr",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/primaryChannel/planner/plans",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/pRiMaRyChAnNeL/pLaNnEr/pLaNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/primaryChannel/planner/plans/plannerPlanId",
			Expected: &GroupIdTeamPrimaryChannelPlannerPlanId{
				GroupId:       "groupId",
				PlannerPlanId: "plannerPlanId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/primaryChannel/planner/plans/plannerPlanId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/pRiMaRyChAnNeL/pLaNnEr/pLaNs/pLaNnErPlAnId",
			Expected: &GroupIdTeamPrimaryChannelPlannerPlanId{
				GroupId:       "gRoUpId",
				PlannerPlanId: "pLaNnErPlAnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/pRiMaRyChAnNeL/pLaNnEr/pLaNs/pLaNnErPlAnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamPrimaryChannelPlannerPlanIDInsensitively(v.Input)
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

		if actual.PlannerPlanId != v.Expected.PlannerPlanId {
			t.Fatalf("Expected %q but got %q for PlannerPlanId", v.Expected.PlannerPlanId, actual.PlannerPlanId)
		}

	}
}

func TestSegmentsForGroupIdTeamPrimaryChannelPlannerPlanId(t *testing.T) {
	segments := GroupIdTeamPrimaryChannelPlannerPlanId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamPrimaryChannelPlannerPlanId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
