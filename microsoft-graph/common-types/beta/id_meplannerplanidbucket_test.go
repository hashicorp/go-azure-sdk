package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePlannerPlanIdBucketId{}

func TestNewMePlannerPlanIdBucketID(t *testing.T) {
	id := NewMePlannerPlanIdBucketID("plannerPlanId", "plannerBucketId")

	if id.PlannerPlanId != "plannerPlanId" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerPlanId'", id.PlannerPlanId, "plannerPlanId")
	}

	if id.PlannerBucketId != "plannerBucketId" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerBucketId'", id.PlannerBucketId, "plannerBucketId")
	}
}

func TestFormatMePlannerPlanIdBucketID(t *testing.T) {
	actual := NewMePlannerPlanIdBucketID("plannerPlanId", "plannerBucketId").ID()
	expected := "/me/planner/plans/plannerPlanId/buckets/plannerBucketId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMePlannerPlanIdBucketID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePlannerPlanIdBucketId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/planner",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/planner/plans",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/planner/plans/plannerPlanId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/planner/plans/plannerPlanId/buckets",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/planner/plans/plannerPlanId/buckets/plannerBucketId",
			Expected: &MePlannerPlanIdBucketId{
				PlannerPlanId:   "plannerPlanId",
				PlannerBucketId: "plannerBucketId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/planner/plans/plannerPlanId/buckets/plannerBucketId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePlannerPlanIdBucketID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PlannerPlanId != v.Expected.PlannerPlanId {
			t.Fatalf("Expected %q but got %q for PlannerPlanId", v.Expected.PlannerPlanId, actual.PlannerPlanId)
		}

		if actual.PlannerBucketId != v.Expected.PlannerBucketId {
			t.Fatalf("Expected %q but got %q for PlannerBucketId", v.Expected.PlannerBucketId, actual.PlannerBucketId)
		}

	}
}

func TestParseMePlannerPlanIdBucketIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePlannerPlanIdBucketId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/planner",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/planner/plans",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/pLaNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/planner/plans/plannerPlanId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/pLaNs/pLaNnErPlAnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/planner/plans/plannerPlanId/buckets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/pLaNs/pLaNnErPlAnId/bUcKeTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/planner/plans/plannerPlanId/buckets/plannerBucketId",
			Expected: &MePlannerPlanIdBucketId{
				PlannerPlanId:   "plannerPlanId",
				PlannerBucketId: "plannerBucketId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/planner/plans/plannerPlanId/buckets/plannerBucketId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/pLaNs/pLaNnErPlAnId/bUcKeTs/pLaNnErBuCkEtId",
			Expected: &MePlannerPlanIdBucketId{
				PlannerPlanId:   "pLaNnErPlAnId",
				PlannerBucketId: "pLaNnErBuCkEtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/pLaNs/pLaNnErPlAnId/bUcKeTs/pLaNnErBuCkEtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePlannerPlanIdBucketIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PlannerPlanId != v.Expected.PlannerPlanId {
			t.Fatalf("Expected %q but got %q for PlannerPlanId", v.Expected.PlannerPlanId, actual.PlannerPlanId)
		}

		if actual.PlannerBucketId != v.Expected.PlannerBucketId {
			t.Fatalf("Expected %q but got %q for PlannerBucketId", v.Expected.PlannerBucketId, actual.PlannerBucketId)
		}

	}
}

func TestSegmentsForMePlannerPlanIdBucketId(t *testing.T) {
	segments := MePlannerPlanIdBucketId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MePlannerPlanIdBucketId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
