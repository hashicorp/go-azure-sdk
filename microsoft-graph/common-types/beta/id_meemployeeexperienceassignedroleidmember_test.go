package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEmployeeExperienceAssignedRoleIdMemberId{}

func TestNewMeEmployeeExperienceAssignedRoleIdMemberID(t *testing.T) {
	id := NewMeEmployeeExperienceAssignedRoleIdMemberID("engagementRoleId", "engagementRoleMemberId")

	if id.EngagementRoleId != "engagementRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'EngagementRoleId'", id.EngagementRoleId, "engagementRoleId")
	}

	if id.EngagementRoleMemberId != "engagementRoleMemberId" {
		t.Fatalf("Expected %q but got %q for Segment 'EngagementRoleMemberId'", id.EngagementRoleMemberId, "engagementRoleMemberId")
	}
}

func TestFormatMeEmployeeExperienceAssignedRoleIdMemberID(t *testing.T) {
	actual := NewMeEmployeeExperienceAssignedRoleIdMemberID("engagementRoleId", "engagementRoleMemberId").ID()
	expected := "/me/employeeExperience/assignedRoles/engagementRoleId/members/engagementRoleMemberId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeEmployeeExperienceAssignedRoleIdMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeEmployeeExperienceAssignedRoleIdMemberId
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
			Input: "/me/employeeExperience",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/employeeExperience/assignedRoles",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/employeeExperience/assignedRoles/engagementRoleId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/employeeExperience/assignedRoles/engagementRoleId/members",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/employeeExperience/assignedRoles/engagementRoleId/members/engagementRoleMemberId",
			Expected: &MeEmployeeExperienceAssignedRoleIdMemberId{
				EngagementRoleId:       "engagementRoleId",
				EngagementRoleMemberId: "engagementRoleMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/employeeExperience/assignedRoles/engagementRoleId/members/engagementRoleMemberId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeEmployeeExperienceAssignedRoleIdMemberID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.EngagementRoleId != v.Expected.EngagementRoleId {
			t.Fatalf("Expected %q but got %q for EngagementRoleId", v.Expected.EngagementRoleId, actual.EngagementRoleId)
		}

		if actual.EngagementRoleMemberId != v.Expected.EngagementRoleMemberId {
			t.Fatalf("Expected %q but got %q for EngagementRoleMemberId", v.Expected.EngagementRoleMemberId, actual.EngagementRoleMemberId)
		}

	}
}

func TestParseMeEmployeeExperienceAssignedRoleIdMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeEmployeeExperienceAssignedRoleIdMemberId
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
			Input: "/me/employeeExperience",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eMpLoYeEeXpErIeNcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/employeeExperience/assignedRoles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eMpLoYeEeXpErIeNcE/aSsIgNeDrOlEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/employeeExperience/assignedRoles/engagementRoleId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eMpLoYeEeXpErIeNcE/aSsIgNeDrOlEs/eNgAgEmEnTrOlEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/employeeExperience/assignedRoles/engagementRoleId/members",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eMpLoYeEeXpErIeNcE/aSsIgNeDrOlEs/eNgAgEmEnTrOlEiD/mEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/employeeExperience/assignedRoles/engagementRoleId/members/engagementRoleMemberId",
			Expected: &MeEmployeeExperienceAssignedRoleIdMemberId{
				EngagementRoleId:       "engagementRoleId",
				EngagementRoleMemberId: "engagementRoleMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/employeeExperience/assignedRoles/engagementRoleId/members/engagementRoleMemberId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eMpLoYeEeXpErIeNcE/aSsIgNeDrOlEs/eNgAgEmEnTrOlEiD/mEmBeRs/eNgAgEmEnTrOlEmEmBeRiD",
			Expected: &MeEmployeeExperienceAssignedRoleIdMemberId{
				EngagementRoleId:       "eNgAgEmEnTrOlEiD",
				EngagementRoleMemberId: "eNgAgEmEnTrOlEmEmBeRiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/eMpLoYeEeXpErIeNcE/aSsIgNeDrOlEs/eNgAgEmEnTrOlEiD/mEmBeRs/eNgAgEmEnTrOlEmEmBeRiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeEmployeeExperienceAssignedRoleIdMemberIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.EngagementRoleId != v.Expected.EngagementRoleId {
			t.Fatalf("Expected %q but got %q for EngagementRoleId", v.Expected.EngagementRoleId, actual.EngagementRoleId)
		}

		if actual.EngagementRoleMemberId != v.Expected.EngagementRoleMemberId {
			t.Fatalf("Expected %q but got %q for EngagementRoleMemberId", v.Expected.EngagementRoleMemberId, actual.EngagementRoleMemberId)
		}

	}
}

func TestSegmentsForMeEmployeeExperienceAssignedRoleIdMemberId(t *testing.T) {
	segments := MeEmployeeExperienceAssignedRoleIdMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeEmployeeExperienceAssignedRoleIdMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
