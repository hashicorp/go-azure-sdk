package userplannerplanbucket

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPlannerPlanBucketId{}

func TestNewUserPlannerPlanBucketID(t *testing.T) {
	id := NewUserPlannerPlanBucketID("userIdValue", "plannerPlanIdValue", "plannerBucketIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.PlannerPlanId != "plannerPlanIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerPlanId'", id.PlannerPlanId, "plannerPlanIdValue")
	}

	if id.PlannerBucketId != "plannerBucketIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerBucketId'", id.PlannerBucketId, "plannerBucketIdValue")
	}
}

func TestFormatUserPlannerPlanBucketID(t *testing.T) {
	actual := NewUserPlannerPlanBucketID("userIdValue", "plannerPlanIdValue", "plannerBucketIdValue").ID()
	expected := "/users/userIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserPlannerPlanBucketID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserPlannerPlanBucketId
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
			// Valid URI
			Input: "/users/userIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue",
			Expected: &UserPlannerPlanBucketId{
				UserId:          "userIdValue",
				PlannerPlanId:   "plannerPlanIdValue",
				PlannerBucketId: "plannerBucketIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserPlannerPlanBucketID(v.Input)
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

	}
}

func TestParseUserPlannerPlanBucketIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserPlannerPlanBucketId
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
			// Valid URI
			Input: "/users/userIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue",
			Expected: &UserPlannerPlanBucketId{
				UserId:          "userIdValue",
				PlannerPlanId:   "plannerPlanIdValue",
				PlannerBucketId: "plannerBucketIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/planner/plans/plannerPlanIdValue/buckets/plannerBucketIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE/bUcKeTs/pLaNnErBuCkEtIdVaLuE",
			Expected: &UserPlannerPlanBucketId{
				UserId:          "uSeRiDvAlUe",
				PlannerPlanId:   "pLaNnErPlAnIdVaLuE",
				PlannerBucketId: "pLaNnErBuCkEtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pLaNnEr/pLaNs/pLaNnErPlAnIdVaLuE/bUcKeTs/pLaNnErBuCkEtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserPlannerPlanBucketIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForUserPlannerPlanBucketId(t *testing.T) {
	segments := UserPlannerPlanBucketId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserPlannerPlanBucketId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
