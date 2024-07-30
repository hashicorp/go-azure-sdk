package plannerplanbuckettask

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdPlannerPlanIdBucketId{}

func TestNewGroupIdPlannerPlanIdBucketID(t *testing.T) {
	id := NewGroupIdPlannerPlanIdBucketID("groupIdValue", "plannerPlanIdValue", "plannerBucketIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.PlannerPlanId != "plannerPlanIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerPlanId'", id.PlannerPlanId, "plannerPlanIdValue")
	}

	if id.PlannerBucketId != "plannerBucketIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerBucketId'", id.PlannerBucketId, "plannerBucketIdValue")
	}
}

func TestFormatGroupIdPlannerPlanIdBucketID(t *testing.T) {
	actual := NewGroupIdPlannerPlanIdBucketID("groupIdValue", "plannerPlanIdValue", "plannerBucketIdValue").ID()
	expected := "/groups/groupIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue"
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
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/planner",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/planner/plans",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/planner/plans/plannerPlanIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/planner/plans/plannerPlanIdValue/buckets",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue",
			Expected: &GroupIdPlannerPlanIdBucketId{
				GroupId:         "groupIdValue",
				PlannerPlanId:   "plannerPlanIdValue",
				PlannerBucketId: "plannerBucketIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue/extra",
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
			Input: "/groups/groupIdValue/planner",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/pLaNnEr",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/planner/plans",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/pLaNnEr/pLaNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/planner/plans/plannerPlanIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/planner/plans/plannerPlanIdValue/buckets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE/bUcKeTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue",
			Expected: &GroupIdPlannerPlanIdBucketId{
				GroupId:         "groupIdValue",
				PlannerPlanId:   "plannerPlanIdValue",
				PlannerBucketId: "plannerBucketIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE/bUcKeTs/pLaNnErBuCkEtIdVaLuE",
			Expected: &GroupIdPlannerPlanIdBucketId{
				GroupId:         "gRoUpIdVaLuE",
				PlannerPlanId:   "pLaNnErPlAnIdVaLuE",
				PlannerBucketId: "pLaNnErBuCkEtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE/bUcKeTs/pLaNnErBuCkEtIdVaLuE/extra",
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
