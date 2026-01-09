package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePlannerFavoritePlanId{}

func TestNewMePlannerFavoritePlanID(t *testing.T) {
	id := NewMePlannerFavoritePlanID("plannerPlanId")

	if id.PlannerPlanId != "plannerPlanId" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerPlanId'", id.PlannerPlanId, "plannerPlanId")
	}
}

func TestFormatMePlannerFavoritePlanID(t *testing.T) {
	actual := NewMePlannerFavoritePlanID("plannerPlanId").ID()
	expected := "/me/planner/favoritePlans/plannerPlanId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMePlannerFavoritePlanID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePlannerFavoritePlanId
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
			Input: "/me/planner/favoritePlans",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/planner/favoritePlans/plannerPlanId",
			Expected: &MePlannerFavoritePlanId{
				PlannerPlanId: "plannerPlanId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/planner/favoritePlans/plannerPlanId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePlannerFavoritePlanID(v.Input)
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

func TestParseMePlannerFavoritePlanIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePlannerFavoritePlanId
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
			Input: "/me/planner/favoritePlans",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/fAvOrItEpLaNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/planner/favoritePlans/plannerPlanId",
			Expected: &MePlannerFavoritePlanId{
				PlannerPlanId: "plannerPlanId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/planner/favoritePlans/plannerPlanId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/fAvOrItEpLaNs/pLaNnErPlAnId",
			Expected: &MePlannerFavoritePlanId{
				PlannerPlanId: "pLaNnErPlAnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pLaNnEr/fAvOrItEpLaNs/pLaNnErPlAnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePlannerFavoritePlanIDInsensitively(v.Input)
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

func TestSegmentsForMePlannerFavoritePlanId(t *testing.T) {
	segments := MePlannerFavoritePlanId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MePlannerFavoritePlanId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
