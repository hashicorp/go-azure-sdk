package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePlannerPlanId{}

func TestNewMePlannerPlanID(t *testing.T) {
	id := NewMePlannerPlanID("plannerPlanId")

	if id.PlannerPlanId != "plannerPlanId" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerPlanId'", id.PlannerPlanId, "plannerPlanId")
	}
}

func TestFormatMePlannerPlanID(t *testing.T) {
	actual := NewMePlannerPlanID("plannerPlanId").ID()
	expected := "/me/planner/plans/plannerPlanId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMePlannerPlanID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePlannerPlanId
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
			// Valid URI
			Input: "/me/planner/plans/plannerPlanId",
			Expected: &MePlannerPlanId{
				PlannerPlanId: "plannerPlanId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/planner/plans/plannerPlanId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePlannerPlanID(v.Input)
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

	}
}

func TestParseMePlannerPlanIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePlannerPlanId
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
			// Valid URI
			Input: "/me/planner/plans/plannerPlanId",
			Expected: &MePlannerPlanId{
				PlannerPlanId: "plannerPlanId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/planner/plans/plannerPlanId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/pLaNs/pLaNnErPlAnId",
			Expected: &MePlannerPlanId{
				PlannerPlanId: "pLaNnErPlAnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/pLaNs/pLaNnErPlAnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePlannerPlanIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMePlannerPlanId(t *testing.T) {
	segments := MePlannerPlanId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MePlannerPlanId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
