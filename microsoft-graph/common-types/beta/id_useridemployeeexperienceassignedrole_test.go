package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdEmployeeExperienceAssignedRoleId{}

func TestNewUserIdEmployeeExperienceAssignedRoleID(t *testing.T) {
	id := NewUserIdEmployeeExperienceAssignedRoleID("userId", "engagementRoleId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.EngagementRoleId != "engagementRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'EngagementRoleId'", id.EngagementRoleId, "engagementRoleId")
	}
}

func TestFormatUserIdEmployeeExperienceAssignedRoleID(t *testing.T) {
	actual := NewUserIdEmployeeExperienceAssignedRoleID("userId", "engagementRoleId").ID()
	expected := "/users/userId/employeeExperience/assignedRoles/engagementRoleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdEmployeeExperienceAssignedRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdEmployeeExperienceAssignedRoleId
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
			Input: "/users/userId/employeeExperience",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/employeeExperience/assignedRoles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/employeeExperience/assignedRoles/engagementRoleId",
			Expected: &UserIdEmployeeExperienceAssignedRoleId{
				UserId:           "userId",
				EngagementRoleId: "engagementRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/employeeExperience/assignedRoles/engagementRoleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdEmployeeExperienceAssignedRoleID(v.Input)
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

		if actual.EngagementRoleId != v.Expected.EngagementRoleId {
			t.Fatalf("Expected %q but got %q for EngagementRoleId", v.Expected.EngagementRoleId, actual.EngagementRoleId)
		}

	}
}

func TestParseUserIdEmployeeExperienceAssignedRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdEmployeeExperienceAssignedRoleId
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
			Input: "/users/userId/employeeExperience",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eMpLoYeEeXpErIeNcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/employeeExperience/assignedRoles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eMpLoYeEeXpErIeNcE/aSsIgNeDrOlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/employeeExperience/assignedRoles/engagementRoleId",
			Expected: &UserIdEmployeeExperienceAssignedRoleId{
				UserId:           "userId",
				EngagementRoleId: "engagementRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/employeeExperience/assignedRoles/engagementRoleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eMpLoYeEeXpErIeNcE/aSsIgNeDrOlEs/eNgAgEmEnTrOlEiD",
			Expected: &UserIdEmployeeExperienceAssignedRoleId{
				UserId:           "uSeRiD",
				EngagementRoleId: "eNgAgEmEnTrOlEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eMpLoYeEeXpErIeNcE/aSsIgNeDrOlEs/eNgAgEmEnTrOlEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdEmployeeExperienceAssignedRoleIDInsensitively(v.Input)
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

		if actual.EngagementRoleId != v.Expected.EngagementRoleId {
			t.Fatalf("Expected %q but got %q for EngagementRoleId", v.Expected.EngagementRoleId, actual.EngagementRoleId)
		}

	}
}

func TestSegmentsForUserIdEmployeeExperienceAssignedRoleId(t *testing.T) {
	segments := UserIdEmployeeExperienceAssignedRoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdEmployeeExperienceAssignedRoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
