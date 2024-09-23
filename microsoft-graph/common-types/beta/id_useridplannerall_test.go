package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPlannerAllId{}

func TestNewUserIdPlannerAllID(t *testing.T) {
	id := NewUserIdPlannerAllID("userId", "plannerDeltaId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.PlannerDeltaId != "plannerDeltaId" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerDeltaId'", id.PlannerDeltaId, "plannerDeltaId")
	}
}

func TestFormatUserIdPlannerAllID(t *testing.T) {
	actual := NewUserIdPlannerAllID("userId", "plannerDeltaId").ID()
	expected := "/users/userId/planner/all/plannerDeltaId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdPlannerAllID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPlannerAllId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/planner",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/planner/all",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/planner/all/plannerDeltaId",
			Expected: &UserIdPlannerAllId{
				UserId:         "userId",
				PlannerDeltaId: "plannerDeltaId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/planner/all/plannerDeltaId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPlannerAllID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.PlannerDeltaId != v.Expected.PlannerDeltaId {
			t.Fatalf("Expected %q but got %q for PlannerDeltaId", v.Expected.PlannerDeltaId, actual.PlannerDeltaId)
		}

	}
}

func TestParseUserIdPlannerAllIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPlannerAllId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/planner",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pLaNnEr",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/planner/all",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pLaNnEr/aLl",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/planner/all/plannerDeltaId",
			Expected: &UserIdPlannerAllId{
				UserId:         "userId",
				PlannerDeltaId: "plannerDeltaId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/planner/all/plannerDeltaId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pLaNnEr/aLl/pLaNnErDeLtAiD",
			Expected: &UserIdPlannerAllId{
				UserId:         "uSeRiD",
				PlannerDeltaId: "pLaNnErDeLtAiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pLaNnEr/aLl/pLaNnErDeLtAiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPlannerAllIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.PlannerDeltaId != v.Expected.PlannerDeltaId {
			t.Fatalf("Expected %q but got %q for PlannerDeltaId", v.Expected.PlannerDeltaId, actual.PlannerDeltaId)
		}

	}
}

func TestSegmentsForUserIdPlannerAllId(t *testing.T) {
	segments := UserIdPlannerAllId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdPlannerAllId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
