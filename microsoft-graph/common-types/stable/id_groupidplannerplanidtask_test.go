package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdPlannerPlanIdTaskId{}

func TestNewGroupIdPlannerPlanIdTaskID(t *testing.T) {
	id := NewGroupIdPlannerPlanIdTaskID("groupId", "plannerPlanId", "plannerTaskId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.PlannerPlanId != "plannerPlanId" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerPlanId'", id.PlannerPlanId, "plannerPlanId")
	}

	if id.PlannerTaskId != "plannerTaskId" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerTaskId'", id.PlannerTaskId, "plannerTaskId")
	}
}

func TestFormatGroupIdPlannerPlanIdTaskID(t *testing.T) {
	actual := NewGroupIdPlannerPlanIdTaskID("groupId", "plannerPlanId", "plannerTaskId").ID()
	expected := "/groups/groupId/planner/plans/plannerPlanId/tasks/plannerTaskId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdPlannerPlanIdTaskID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdPlannerPlanIdTaskId
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
			Input: "/groups/groupId/planner",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/planner/plans",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/planner/plans/plannerPlanId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/planner/plans/plannerPlanId/tasks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/planner/plans/plannerPlanId/tasks/plannerTaskId",
			Expected: &GroupIdPlannerPlanIdTaskId{
				GroupId:       "groupId",
				PlannerPlanId: "plannerPlanId",
				PlannerTaskId: "plannerTaskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/planner/plans/plannerPlanId/tasks/plannerTaskId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdPlannerPlanIdTaskID(v.Input)
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

		if actual.PlannerTaskId != v.Expected.PlannerTaskId {
			t.Fatalf("Expected %q but got %q for PlannerTaskId", v.Expected.PlannerTaskId, actual.PlannerTaskId)
		}

	}
}

func TestParseGroupIdPlannerPlanIdTaskIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdPlannerPlanIdTaskId
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
			Input: "/groups/groupId/planner",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/pLaNnEr",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/planner/plans",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/pLaNnEr/pLaNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/planner/plans/plannerPlanId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/pLaNnEr/pLaNs/pLaNnErPlAnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/planner/plans/plannerPlanId/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/pLaNnEr/pLaNs/pLaNnErPlAnId/tAsKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/planner/plans/plannerPlanId/tasks/plannerTaskId",
			Expected: &GroupIdPlannerPlanIdTaskId{
				GroupId:       "groupId",
				PlannerPlanId: "plannerPlanId",
				PlannerTaskId: "plannerTaskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/planner/plans/plannerPlanId/tasks/plannerTaskId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/pLaNnEr/pLaNs/pLaNnErPlAnId/tAsKs/pLaNnErTaSkId",
			Expected: &GroupIdPlannerPlanIdTaskId{
				GroupId:       "gRoUpId",
				PlannerPlanId: "pLaNnErPlAnId",
				PlannerTaskId: "pLaNnErTaSkId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/pLaNnEr/pLaNs/pLaNnErPlAnId/tAsKs/pLaNnErTaSkId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdPlannerPlanIdTaskIDInsensitively(v.Input)
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

		if actual.PlannerTaskId != v.Expected.PlannerTaskId {
			t.Fatalf("Expected %q but got %q for PlannerTaskId", v.Expected.PlannerTaskId, actual.PlannerTaskId)
		}

	}
}

func TestSegmentsForGroupIdPlannerPlanIdTaskId(t *testing.T) {
	segments := GroupIdPlannerPlanIdTaskId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdPlannerPlanIdTaskId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
