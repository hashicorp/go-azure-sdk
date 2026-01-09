package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePlannerAllId{}

func TestNewMePlannerAllID(t *testing.T) {
	id := NewMePlannerAllID("plannerDeltaId")

	if id.PlannerDeltaId != "plannerDeltaId" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerDeltaId'", id.PlannerDeltaId, "plannerDeltaId")
	}
}

func TestFormatMePlannerAllID(t *testing.T) {
	actual := NewMePlannerAllID("plannerDeltaId").ID()
	expected := "/me/planner/all/plannerDeltaId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMePlannerAllID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePlannerAllId
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
			Input: "/me/planner/all",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/planner/all/plannerDeltaId",
			Expected: &MePlannerAllId{
				PlannerDeltaId: "plannerDeltaId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/planner/all/plannerDeltaId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePlannerAllID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PlannerDeltaId != v.Expected.PlannerDeltaId {
			t.Fatalf("Expected %q but got %q for PlannerDeltaId", v.Expected.PlannerDeltaId, actual.PlannerDeltaId)
		}

	}
}

func TestParseMePlannerAllIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePlannerAllId
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
			Input: "/me/planner/all",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/aLl",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/planner/all/plannerDeltaId",
			Expected: &MePlannerAllId{
				PlannerDeltaId: "plannerDeltaId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/planner/all/plannerDeltaId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/aLl/pLaNnErDeLtAiD",
			Expected: &MePlannerAllId{
				PlannerDeltaId: "pLaNnErDeLtAiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/aLl/pLaNnErDeLtAiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePlannerAllIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PlannerDeltaId != v.Expected.PlannerDeltaId {
			t.Fatalf("Expected %q but got %q for PlannerDeltaId", v.Expected.PlannerDeltaId, actual.PlannerDeltaId)
		}

	}
}

func TestSegmentsForMePlannerAllId(t *testing.T) {
	segments := MePlannerAllId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MePlannerAllId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
