package meplannerplanbucket

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePlannerPlanBucketId{}

func TestNewMePlannerPlanBucketID(t *testing.T) {
	id := NewMePlannerPlanBucketID("plannerPlanIdValue", "plannerBucketIdValue")

	if id.PlannerPlanId != "plannerPlanIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerPlanId'", id.PlannerPlanId, "plannerPlanIdValue")
	}

	if id.PlannerBucketId != "plannerBucketIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerBucketId'", id.PlannerBucketId, "plannerBucketIdValue")
	}
}

func TestFormatMePlannerPlanBucketID(t *testing.T) {
	actual := NewMePlannerPlanBucketID("plannerPlanIdValue", "plannerBucketIdValue").ID()
	expected := "/me/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMePlannerPlanBucketID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePlannerPlanBucketId
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
			Input: "/me/planner/plans/plannerPlanIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/planner/plans/plannerPlanIdValue/buckets",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue",
			Expected: &MePlannerPlanBucketId{
				PlannerPlanId:   "plannerPlanIdValue",
				PlannerBucketId: "plannerBucketIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePlannerPlanBucketID(v.Input)
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

func TestParseMePlannerPlanBucketIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePlannerPlanBucketId
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
			Input: "/me/planner/plans/plannerPlanIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/planner/plans/plannerPlanIdValue/buckets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE/bUcKeTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue",
			Expected: &MePlannerPlanBucketId{
				PlannerPlanId:   "plannerPlanIdValue",
				PlannerBucketId: "plannerBucketIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE/bUcKeTs/pLaNnErBuCkEtIdVaLuE",
			Expected: &MePlannerPlanBucketId{
				PlannerPlanId:   "pLaNnErPlAnIdVaLuE",
				PlannerBucketId: "pLaNnErBuCkEtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE/bUcKeTs/pLaNnErBuCkEtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePlannerPlanBucketIDInsensitively(v.Input)
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

func TestSegmentsForMePlannerPlanBucketId(t *testing.T) {
	segments := MePlannerPlanBucketId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MePlannerPlanBucketId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
