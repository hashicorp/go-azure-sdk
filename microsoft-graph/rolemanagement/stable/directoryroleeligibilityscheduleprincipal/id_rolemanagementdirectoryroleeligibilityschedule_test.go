package directoryroleeligibilityscheduleprincipal

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryRoleEligibilityScheduleId{}

func TestNewRoleManagementDirectoryRoleEligibilityScheduleID(t *testing.T) {
	id := NewRoleManagementDirectoryRoleEligibilityScheduleID("unifiedRoleEligibilityScheduleIdValue")

	if id.UnifiedRoleEligibilityScheduleId != "unifiedRoleEligibilityScheduleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleEligibilityScheduleId'", id.UnifiedRoleEligibilityScheduleId, "unifiedRoleEligibilityScheduleIdValue")
	}
}

func TestFormatRoleManagementDirectoryRoleEligibilityScheduleID(t *testing.T) {
	actual := NewRoleManagementDirectoryRoleEligibilityScheduleID("unifiedRoleEligibilityScheduleIdValue").ID()
	expected := "/roleManagement/directory/roleEligibilitySchedules/unifiedRoleEligibilityScheduleIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementDirectoryRoleEligibilityScheduleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDirectoryRoleEligibilityScheduleId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/directory",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/directory/roleEligibilitySchedules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/directory/roleEligibilitySchedules/unifiedRoleEligibilityScheduleIdValue",
			Expected: &RoleManagementDirectoryRoleEligibilityScheduleId{
				UnifiedRoleEligibilityScheduleId: "unifiedRoleEligibilityScheduleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/directory/roleEligibilitySchedules/unifiedRoleEligibilityScheduleIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDirectoryRoleEligibilityScheduleID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleEligibilityScheduleId != v.Expected.UnifiedRoleEligibilityScheduleId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleEligibilityScheduleId", v.Expected.UnifiedRoleEligibilityScheduleId, actual.UnifiedRoleEligibilityScheduleId)
		}

	}
}

func TestParseRoleManagementDirectoryRoleEligibilityScheduleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDirectoryRoleEligibilityScheduleId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/directory",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/directory/roleEligibilitySchedules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEeLiGiBiLiTyScHeDuLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/directory/roleEligibilitySchedules/unifiedRoleEligibilityScheduleIdValue",
			Expected: &RoleManagementDirectoryRoleEligibilityScheduleId{
				UnifiedRoleEligibilityScheduleId: "unifiedRoleEligibilityScheduleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/directory/roleEligibilitySchedules/unifiedRoleEligibilityScheduleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEeLiGiBiLiTyScHeDuLeS/uNiFiEdRoLeElIgIbIlItYsChEdUlEiDvAlUe",
			Expected: &RoleManagementDirectoryRoleEligibilityScheduleId{
				UnifiedRoleEligibilityScheduleId: "uNiFiEdRoLeElIgIbIlItYsChEdUlEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEeLiGiBiLiTyScHeDuLeS/uNiFiEdRoLeElIgIbIlItYsChEdUlEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDirectoryRoleEligibilityScheduleIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleEligibilityScheduleId != v.Expected.UnifiedRoleEligibilityScheduleId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleEligibilityScheduleId", v.Expected.UnifiedRoleEligibilityScheduleId, actual.UnifiedRoleEligibilityScheduleId)
		}

	}
}

func TestSegmentsForRoleManagementDirectoryRoleEligibilityScheduleId(t *testing.T) {
	segments := RoleManagementDirectoryRoleEligibilityScheduleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementDirectoryRoleEligibilityScheduleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
