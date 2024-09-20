package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPlannerRecentPlanId{}

func TestNewUserIdPlannerRecentPlanID(t *testing.T) {
	id := NewUserIdPlannerRecentPlanID("userId", "plannerPlanId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.PlannerPlanId != "plannerPlanId" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerPlanId'", id.PlannerPlanId, "plannerPlanId")
	}
}

func TestFormatUserIdPlannerRecentPlanID(t *testing.T) {
	actual := NewUserIdPlannerRecentPlanID("userId", "plannerPlanId").ID()
	expected := "/users/userId/planner/recentPlans/plannerPlanId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdPlannerRecentPlanID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPlannerRecentPlanId
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
			Input: "/users/userId/planner/recentPlans",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/planner/recentPlans/plannerPlanId",
			Expected: &UserIdPlannerRecentPlanId{
				UserId:        "userId",
				PlannerPlanId: "plannerPlanId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/planner/recentPlans/plannerPlanId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPlannerRecentPlanID(v.Input)
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

		if actual.PlannerPlanId != v.Expected.PlannerPlanId {
			t.Fatalf("Expected %q but got %q for PlannerPlanId", v.Expected.PlannerPlanId, actual.PlannerPlanId)
		}

	}
}

func TestParseUserIdPlannerRecentPlanIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPlannerRecentPlanId
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
			Input: "/users/userId/planner/recentPlans",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pLaNnEr/rEcEnTpLaNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/planner/recentPlans/plannerPlanId",
			Expected: &UserIdPlannerRecentPlanId{
				UserId:        "userId",
				PlannerPlanId: "plannerPlanId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/planner/recentPlans/plannerPlanId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pLaNnEr/rEcEnTpLaNs/pLaNnErPlAnId",
			Expected: &UserIdPlannerRecentPlanId{
				UserId:        "uSeRiD",
				PlannerPlanId: "pLaNnErPlAnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pLaNnEr/rEcEnTpLaNs/pLaNnErPlAnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPlannerRecentPlanIDInsensitively(v.Input)
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

		if actual.PlannerPlanId != v.Expected.PlannerPlanId {
			t.Fatalf("Expected %q but got %q for PlannerPlanId", v.Expected.PlannerPlanId, actual.PlannerPlanId)
		}

	}
}

func TestSegmentsForUserIdPlannerRecentPlanId(t *testing.T) {
	segments := UserIdPlannerRecentPlanId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdPlannerRecentPlanId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
