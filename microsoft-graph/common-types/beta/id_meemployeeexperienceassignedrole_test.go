package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEmployeeExperienceAssignedRoleId{}

func TestNewMeEmployeeExperienceAssignedRoleID(t *testing.T) {
	id := NewMeEmployeeExperienceAssignedRoleID("engagementRoleId")

	if id.EngagementRoleId != "engagementRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'EngagementRoleId'", id.EngagementRoleId, "engagementRoleId")
	}
}

func TestFormatMeEmployeeExperienceAssignedRoleID(t *testing.T) {
	actual := NewMeEmployeeExperienceAssignedRoleID("engagementRoleId").ID()
	expected := "/me/employeeExperience/assignedRoles/engagementRoleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeEmployeeExperienceAssignedRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeEmployeeExperienceAssignedRoleId
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
			// Valid URI
			Input: "/me/employeeExperience/assignedRoles/engagementRoleId",
			Expected: &MeEmployeeExperienceAssignedRoleId{
				EngagementRoleId: "engagementRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/employeeExperience/assignedRoles/engagementRoleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeEmployeeExperienceAssignedRoleID(v.Input)
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

	}
}

func TestParseMeEmployeeExperienceAssignedRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeEmployeeExperienceAssignedRoleId
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
			// Valid URI
			Input: "/me/employeeExperience/assignedRoles/engagementRoleId",
			Expected: &MeEmployeeExperienceAssignedRoleId{
				EngagementRoleId: "engagementRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/employeeExperience/assignedRoles/engagementRoleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eMpLoYeEeXpErIeNcE/aSsIgNeDrOlEs/eNgAgEmEnTrOlEiD",
			Expected: &MeEmployeeExperienceAssignedRoleId{
				EngagementRoleId: "eNgAgEmEnTrOlEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/eMpLoYeEeXpErIeNcE/aSsIgNeDrOlEs/eNgAgEmEnTrOlEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeEmployeeExperienceAssignedRoleIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeEmployeeExperienceAssignedRoleId(t *testing.T) {
	segments := MeEmployeeExperienceAssignedRoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeEmployeeExperienceAssignedRoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
