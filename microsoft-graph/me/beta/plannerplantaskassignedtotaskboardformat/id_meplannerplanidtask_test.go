package plannerplantaskassignedtotaskboardformat

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePlannerPlanIdTaskId{}

func TestNewMePlannerPlanIdTaskID(t *testing.T) {
	id := NewMePlannerPlanIdTaskID("plannerPlanIdValue", "plannerTaskIdValue")

	if id.PlannerPlanId != "plannerPlanIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerPlanId'", id.PlannerPlanId, "plannerPlanIdValue")
	}

	if id.PlannerTaskId != "plannerTaskIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerTaskId'", id.PlannerTaskId, "plannerTaskIdValue")
	}
}

func TestFormatMePlannerPlanIdTaskID(t *testing.T) {
	actual := NewMePlannerPlanIdTaskID("plannerPlanIdValue", "plannerTaskIdValue").ID()
	expected := "/me/planner/plans/plannerPlanIdValue/tasks/plannerTaskIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMePlannerPlanIdTaskID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePlannerPlanIdTaskId
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
			Input: "/me/planner/plans/plannerPlanIdValue/tasks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/planner/plans/plannerPlanIdValue/tasks/plannerTaskIdValue",
			Expected: &MePlannerPlanIdTaskId{
				PlannerPlanId: "plannerPlanIdValue",
				PlannerTaskId: "plannerTaskIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/planner/plans/plannerPlanIdValue/tasks/plannerTaskIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePlannerPlanIdTaskID(v.Input)
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

		if actual.PlannerTaskId != v.Expected.PlannerTaskId {
			t.Fatalf("Expected %q but got %q for PlannerTaskId", v.Expected.PlannerTaskId, actual.PlannerTaskId)
		}

	}
}

func TestParseMePlannerPlanIdTaskIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePlannerPlanIdTaskId
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
			Input: "/me/planner/plans/plannerPlanIdValue/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE/tAsKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/planner/plans/plannerPlanIdValue/tasks/plannerTaskIdValue",
			Expected: &MePlannerPlanIdTaskId{
				PlannerPlanId: "plannerPlanIdValue",
				PlannerTaskId: "plannerTaskIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/planner/plans/plannerPlanIdValue/tasks/plannerTaskIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE/tAsKs/pLaNnErTaSkIdVaLuE",
			Expected: &MePlannerPlanIdTaskId{
				PlannerPlanId: "pLaNnErPlAnIdVaLuE",
				PlannerTaskId: "pLaNnErTaSkIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE/tAsKs/pLaNnErTaSkIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePlannerPlanIdTaskIDInsensitively(v.Input)
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

		if actual.PlannerTaskId != v.Expected.PlannerTaskId {
			t.Fatalf("Expected %q but got %q for PlannerTaskId", v.Expected.PlannerTaskId, actual.PlannerTaskId)
		}

	}
}

func TestSegmentsForMePlannerPlanIdTaskId(t *testing.T) {
	segments := MePlannerPlanIdTaskId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MePlannerPlanIdTaskId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
