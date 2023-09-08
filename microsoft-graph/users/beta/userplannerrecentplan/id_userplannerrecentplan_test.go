package userplannerrecentplan

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPlannerRecentPlanId{}

func TestNewUserPlannerRecentPlanID(t *testing.T) {
	id := NewUserPlannerRecentPlanID("userIdValue", "plannerPlanIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.PlannerPlanId != "plannerPlanIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerPlanId'", id.PlannerPlanId, "plannerPlanIdValue")
	}
}

func TestFormatUserPlannerRecentPlanID(t *testing.T) {
	actual := NewUserPlannerRecentPlanID("userIdValue", "plannerPlanIdValue").ID()
	expected := "/users/userIdValue/planner/recentPlans/plannerPlanIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserPlannerRecentPlanID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserPlannerRecentPlanId
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/planner",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/planner/recentPlans",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/planner/recentPlans/plannerPlanIdValue",
			Expected: &UserPlannerRecentPlanId{
				UserId:        "userIdValue",
				PlannerPlanId: "plannerPlanIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/planner/recentPlans/plannerPlanIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserPlannerRecentPlanID(v.Input)
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

func TestParseUserPlannerRecentPlanIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserPlannerRecentPlanId
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/planner",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pLaNnEr",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/planner/recentPlans",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pLaNnEr/rEcEnTpLaNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/planner/recentPlans/plannerPlanIdValue",
			Expected: &UserPlannerRecentPlanId{
				UserId:        "userIdValue",
				PlannerPlanId: "plannerPlanIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/planner/recentPlans/plannerPlanIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pLaNnEr/rEcEnTpLaNs/pLaNnErPlAnIdVaLuE",
			Expected: &UserPlannerRecentPlanId{
				UserId:        "uSeRiDvAlUe",
				PlannerPlanId: "pLaNnErPlAnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pLaNnEr/rEcEnTpLaNs/pLaNnErPlAnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserPlannerRecentPlanIDInsensitively(v.Input)
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

func TestSegmentsForUserPlannerRecentPlanId(t *testing.T) {
	segments := UserPlannerRecentPlanId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserPlannerRecentPlanId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
