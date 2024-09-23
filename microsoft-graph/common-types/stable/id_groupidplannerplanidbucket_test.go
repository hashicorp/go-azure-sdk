package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdPlannerPlanIdBucketId{}

func TestNewGroupIdPlannerPlanIdBucketID(t *testing.T) {
	id := NewGroupIdPlannerPlanIdBucketID("groupId", "plannerPlanId", "plannerBucketId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.PlannerPlanId != "plannerPlanId" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerPlanId'", id.PlannerPlanId, "plannerPlanId")
	}

	if id.PlannerBucketId != "plannerBucketId" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerBucketId'", id.PlannerBucketId, "plannerBucketId")
	}
}

func TestFormatGroupIdPlannerPlanIdBucketID(t *testing.T) {
	actual := NewGroupIdPlannerPlanIdBucketID("groupId", "plannerPlanId", "plannerBucketId").ID()
	expected := "/groups/groupId/planner/plans/plannerPlanId/buckets/plannerBucketId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdPlannerPlanIdBucketID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdPlannerPlanIdBucketId
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
			Input: "/groups/groupId/planner/plans/plannerPlanId/buckets",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/planner/plans/plannerPlanId/buckets/plannerBucketId",
			Expected: &GroupIdPlannerPlanIdBucketId{
				GroupId:         "groupId",
				PlannerPlanId:   "plannerPlanId",
				PlannerBucketId: "plannerBucketId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/planner/plans/plannerPlanId/buckets/plannerBucketId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdPlannerPlanIdBucketID(v.Input)
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

		if actual.PlannerBucketId != v.Expected.PlannerBucketId {
			t.Fatalf("Expected %q but got %q for PlannerBucketId", v.Expected.PlannerBucketId, actual.PlannerBucketId)
		}

	}
}

func TestParseGroupIdPlannerPlanIdBucketIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdPlannerPlanIdBucketId
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
			Input: "/groups/groupId/planner/plans/plannerPlanId/buckets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/pLaNnEr/pLaNs/pLaNnErPlAnId/bUcKeTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/planner/plans/plannerPlanId/buckets/plannerBucketId",
			Expected: &GroupIdPlannerPlanIdBucketId{
				GroupId:         "groupId",
				PlannerPlanId:   "plannerPlanId",
				PlannerBucketId: "plannerBucketId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/planner/plans/plannerPlanId/buckets/plannerBucketId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/pLaNnEr/pLaNs/pLaNnErPlAnId/bUcKeTs/pLaNnErBuCkEtId",
			Expected: &GroupIdPlannerPlanIdBucketId{
				GroupId:         "gRoUpId",
				PlannerPlanId:   "pLaNnErPlAnId",
				PlannerBucketId: "pLaNnErBuCkEtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/pLaNnEr/pLaNs/pLaNnErPlAnId/bUcKeTs/pLaNnErBuCkEtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdPlannerPlanIdBucketIDInsensitively(v.Input)
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

		if actual.PlannerBucketId != v.Expected.PlannerBucketId {
			t.Fatalf("Expected %q but got %q for PlannerBucketId", v.Expected.PlannerBucketId, actual.PlannerBucketId)
		}

	}
}

func TestSegmentsForGroupIdPlannerPlanIdBucketId(t *testing.T) {
	segments := GroupIdPlannerPlanIdBucketId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdPlannerPlanIdBucketId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
