package plannerall

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePlannerAllId{}

func TestNewMePlannerAllID(t *testing.T) {
	id := NewMePlannerAllID("plannerDeltaIdValue")

	if id.PlannerDeltaId != "plannerDeltaIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerDeltaId'", id.PlannerDeltaId, "plannerDeltaIdValue")
	}
}

func TestFormatMePlannerAllID(t *testing.T) {
	actual := NewMePlannerAllID("plannerDeltaIdValue").ID()
	expected := "/me/planner/all/plannerDeltaIdValue"
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
			Input: "/me/planner/all/plannerDeltaIdValue",
			Expected: &MePlannerAllId{
				PlannerDeltaId: "plannerDeltaIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/planner/all/plannerDeltaIdValue/extra",
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
			Input: "/me/planner/all/plannerDeltaIdValue",
			Expected: &MePlannerAllId{
				PlannerDeltaId: "plannerDeltaIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/planner/all/plannerDeltaIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/aLl/pLaNnErDeLtAiDvAlUe",
			Expected: &MePlannerAllId{
				PlannerDeltaId: "pLaNnErDeLtAiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/aLl/pLaNnErDeLtAiDvAlUe/extra",
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
