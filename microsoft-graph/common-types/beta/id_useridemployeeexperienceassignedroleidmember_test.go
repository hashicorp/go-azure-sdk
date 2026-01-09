package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdEmployeeExperienceAssignedRoleIdMemberId{}

func TestNewUserIdEmployeeExperienceAssignedRoleIdMemberID(t *testing.T) {
	id := NewUserIdEmployeeExperienceAssignedRoleIdMemberID("userId", "engagementRoleId", "engagementRoleMemberId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.EngagementRoleId != "engagementRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'EngagementRoleId'", id.EngagementRoleId, "engagementRoleId")
	}

	if id.EngagementRoleMemberId != "engagementRoleMemberId" {
		t.Fatalf("Expected %q but got %q for Segment 'EngagementRoleMemberId'", id.EngagementRoleMemberId, "engagementRoleMemberId")
	}
}

func TestFormatUserIdEmployeeExperienceAssignedRoleIdMemberID(t *testing.T) {
	actual := NewUserIdEmployeeExperienceAssignedRoleIdMemberID("userId", "engagementRoleId", "engagementRoleMemberId").ID()
	expected := "/users/userId/employeeExperience/assignedRoles/engagementRoleId/members/engagementRoleMemberId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdEmployeeExperienceAssignedRoleIdMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdEmployeeExperienceAssignedRoleIdMemberId
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
			// Incomplete URI
			Input: "/users/userId/employeeExperience/assignedRoles/engagementRoleId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/employeeExperience/assignedRoles/engagementRoleId/members",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/employeeExperience/assignedRoles/engagementRoleId/members/engagementRoleMemberId",
			Expected: &UserIdEmployeeExperienceAssignedRoleIdMemberId{
				UserId:                 "userId",
				EngagementRoleId:       "engagementRoleId",
				EngagementRoleMemberId: "engagementRoleMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/employeeExperience/assignedRoles/engagementRoleId/members/engagementRoleMemberId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdEmployeeExperienceAssignedRoleIdMemberID(v.Input)
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

		if actual.EngagementRoleMemberId != v.Expected.EngagementRoleMemberId {
			t.Fatalf("Expected %q but got %q for EngagementRoleMemberId", v.Expected.EngagementRoleMemberId, actual.EngagementRoleMemberId)
		}

	}
}

func TestParseUserIdEmployeeExperienceAssignedRoleIdMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdEmployeeExperienceAssignedRoleIdMemberId
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
			// Incomplete URI
			Input: "/users/userId/employeeExperience/assignedRoles/engagementRoleId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eMpLoYeEeXpErIeNcE/aSsIgNeDrOlEs/eNgAgEmEnTrOlEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/employeeExperience/assignedRoles/engagementRoleId/members",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eMpLoYeEeXpErIeNcE/aSsIgNeDrOlEs/eNgAgEmEnTrOlEiD/mEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/employeeExperience/assignedRoles/engagementRoleId/members/engagementRoleMemberId",
			Expected: &UserIdEmployeeExperienceAssignedRoleIdMemberId{
				UserId:                 "userId",
				EngagementRoleId:       "engagementRoleId",
				EngagementRoleMemberId: "engagementRoleMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/employeeExperience/assignedRoles/engagementRoleId/members/engagementRoleMemberId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eMpLoYeEeXpErIeNcE/aSsIgNeDrOlEs/eNgAgEmEnTrOlEiD/mEmBeRs/eNgAgEmEnTrOlEmEmBeRiD",
			Expected: &UserIdEmployeeExperienceAssignedRoleIdMemberId{
				UserId:                 "uSeRiD",
				EngagementRoleId:       "eNgAgEmEnTrOlEiD",
				EngagementRoleMemberId: "eNgAgEmEnTrOlEmEmBeRiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eMpLoYeEeXpErIeNcE/aSsIgNeDrOlEs/eNgAgEmEnTrOlEiD/mEmBeRs/eNgAgEmEnTrOlEmEmBeRiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdEmployeeExperienceAssignedRoleIdMemberIDInsensitively(v.Input)
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

		if actual.EngagementRoleMemberId != v.Expected.EngagementRoleMemberId {
			t.Fatalf("Expected %q but got %q for EngagementRoleMemberId", v.Expected.EngagementRoleMemberId, actual.EngagementRoleMemberId)
		}

	}
}

func TestSegmentsForUserIdEmployeeExperienceAssignedRoleIdMemberId(t *testing.T) {
	segments := UserIdEmployeeExperienceAssignedRoleIdMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdEmployeeExperienceAssignedRoleIdMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
