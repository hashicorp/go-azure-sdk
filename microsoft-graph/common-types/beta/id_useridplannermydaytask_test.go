package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPlannerMyDayTaskId{}

func TestNewUserIdPlannerMyDayTaskID(t *testing.T) {
	id := NewUserIdPlannerMyDayTaskID("userId", "plannerTaskId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.PlannerTaskId != "plannerTaskId" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerTaskId'", id.PlannerTaskId, "plannerTaskId")
	}
}

func TestFormatUserIdPlannerMyDayTaskID(t *testing.T) {
	actual := NewUserIdPlannerMyDayTaskID("userId", "plannerTaskId").ID()
	expected := "/users/userId/planner/myDayTasks/plannerTaskId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdPlannerMyDayTaskID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPlannerMyDayTaskId
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
			Input: "/users/userId/planner/myDayTasks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/planner/myDayTasks/plannerTaskId",
			Expected: &UserIdPlannerMyDayTaskId{
				UserId:        "userId",
				PlannerTaskId: "plannerTaskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/planner/myDayTasks/plannerTaskId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPlannerMyDayTaskID(v.Input)
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

		if actual.PlannerTaskId != v.Expected.PlannerTaskId {
			t.Fatalf("Expected %q but got %q for PlannerTaskId", v.Expected.PlannerTaskId, actual.PlannerTaskId)
		}

	}
}

func TestParseUserIdPlannerMyDayTaskIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPlannerMyDayTaskId
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
			Input: "/users/userId/planner/myDayTasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pLaNnEr/mYdAyTaSkS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/planner/myDayTasks/plannerTaskId",
			Expected: &UserIdPlannerMyDayTaskId{
				UserId:        "userId",
				PlannerTaskId: "plannerTaskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/planner/myDayTasks/plannerTaskId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pLaNnEr/mYdAyTaSkS/pLaNnErTaSkId",
			Expected: &UserIdPlannerMyDayTaskId{
				UserId:        "uSeRiD",
				PlannerTaskId: "pLaNnErTaSkId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pLaNnEr/mYdAyTaSkS/pLaNnErTaSkId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPlannerMyDayTaskIDInsensitively(v.Input)
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

		if actual.PlannerTaskId != v.Expected.PlannerTaskId {
			t.Fatalf("Expected %q but got %q for PlannerTaskId", v.Expected.PlannerTaskId, actual.PlannerTaskId)
		}

	}
}

func TestSegmentsForUserIdPlannerMyDayTaskId(t *testing.T) {
	segments := UserIdPlannerMyDayTaskId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdPlannerMyDayTaskId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
