package userplannerplanbuckettaskdetail

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPlannerPlanBucketTaskId{}

func TestNewUserPlannerPlanBucketTaskID(t *testing.T) {
	id := NewUserPlannerPlanBucketTaskID("userIdValue", "plannerPlanIdValue", "plannerBucketIdValue", "plannerTaskIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.PlannerPlanId != "plannerPlanIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerPlanId'", id.PlannerPlanId, "plannerPlanIdValue")
	}

	if id.PlannerBucketId != "plannerBucketIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerBucketId'", id.PlannerBucketId, "plannerBucketIdValue")
	}

	if id.PlannerTaskId != "plannerTaskIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerTaskId'", id.PlannerTaskId, "plannerTaskIdValue")
	}
}

func TestFormatUserPlannerPlanBucketTaskID(t *testing.T) {
	actual := NewUserPlannerPlanBucketTaskID("userIdValue", "plannerPlanIdValue", "plannerBucketIdValue", "plannerTaskIdValue").ID()
	expected := "/users/userIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue/tasks/plannerTaskIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserPlannerPlanBucketTaskID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserPlannerPlanBucketTaskId
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
			Input: "/users/userIdValue/planner/plans",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/planner/plans/plannerPlanIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/planner/plans/plannerPlanIdValue/buckets",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue/tasks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue/tasks/plannerTaskIdValue",
			Expected: &UserPlannerPlanBucketTaskId{
				UserId:          "userIdValue",
				PlannerPlanId:   "plannerPlanIdValue",
				PlannerBucketId: "plannerBucketIdValue",
				PlannerTaskId:   "plannerTaskIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue/tasks/plannerTaskIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserPlannerPlanBucketTaskID(v.Input)
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

		if actual.PlannerBucketId != v.Expected.PlannerBucketId {
			t.Fatalf("Expected %q but got %q for PlannerBucketId", v.Expected.PlannerBucketId, actual.PlannerBucketId)
		}

		if actual.PlannerTaskId != v.Expected.PlannerTaskId {
			t.Fatalf("Expected %q but got %q for PlannerTaskId", v.Expected.PlannerTaskId, actual.PlannerTaskId)
		}

	}
}

func TestParseUserPlannerPlanBucketTaskIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserPlannerPlanBucketTaskId
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
			Input: "/users/userIdValue/planner/plans",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pLaNnEr/pLaNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/planner/plans/plannerPlanIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/planner/plans/plannerPlanIdValue/buckets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE/bUcKeTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE/bUcKeTs/pLaNnErBuCkEtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE/bUcKeTs/pLaNnErBuCkEtIdVaLuE/tAsKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue/tasks/plannerTaskIdValue",
			Expected: &UserPlannerPlanBucketTaskId{
				UserId:          "userIdValue",
				PlannerPlanId:   "plannerPlanIdValue",
				PlannerBucketId: "plannerBucketIdValue",
				PlannerTaskId:   "plannerTaskIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue/tasks/plannerTaskIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE/bUcKeTs/pLaNnErBuCkEtIdVaLuE/tAsKs/pLaNnErTaSkIdVaLuE",
			Expected: &UserPlannerPlanBucketTaskId{
				UserId:          "uSeRiDvAlUe",
				PlannerPlanId:   "pLaNnErPlAnIdVaLuE",
				PlannerBucketId: "pLaNnErBuCkEtIdVaLuE",
				PlannerTaskId:   "pLaNnErTaSkIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE/bUcKeTs/pLaNnErBuCkEtIdVaLuE/tAsKs/pLaNnErTaSkIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserPlannerPlanBucketTaskIDInsensitively(v.Input)
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

		if actual.PlannerBucketId != v.Expected.PlannerBucketId {
			t.Fatalf("Expected %q but got %q for PlannerBucketId", v.Expected.PlannerBucketId, actual.PlannerBucketId)
		}

		if actual.PlannerTaskId != v.Expected.PlannerTaskId {
			t.Fatalf("Expected %q but got %q for PlannerTaskId", v.Expected.PlannerTaskId, actual.PlannerTaskId)
		}

	}
}

func TestSegmentsForUserPlannerPlanBucketTaskId(t *testing.T) {
	segments := UserPlannerPlanBucketTaskId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserPlannerPlanBucketTaskId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
