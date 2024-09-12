package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryRoleAssignmentScheduleId{}

func TestNewRoleManagementDirectoryRoleAssignmentScheduleID(t *testing.T) {
	id := NewRoleManagementDirectoryRoleAssignmentScheduleID("unifiedRoleAssignmentScheduleIdValue")

	if id.UnifiedRoleAssignmentScheduleId != "unifiedRoleAssignmentScheduleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentScheduleId'", id.UnifiedRoleAssignmentScheduleId, "unifiedRoleAssignmentScheduleIdValue")
	}
}

func TestFormatRoleManagementDirectoryRoleAssignmentScheduleID(t *testing.T) {
	actual := NewRoleManagementDirectoryRoleAssignmentScheduleID("unifiedRoleAssignmentScheduleIdValue").ID()
	expected := "/roleManagement/directory/roleAssignmentSchedules/unifiedRoleAssignmentScheduleIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementDirectoryRoleAssignmentScheduleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDirectoryRoleAssignmentScheduleId
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
			Input: "/roleManagement/directory/roleAssignmentSchedules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/directory/roleAssignmentSchedules/unifiedRoleAssignmentScheduleIdValue",
			Expected: &RoleManagementDirectoryRoleAssignmentScheduleId{
				UnifiedRoleAssignmentScheduleId: "unifiedRoleAssignmentScheduleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/directory/roleAssignmentSchedules/unifiedRoleAssignmentScheduleIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDirectoryRoleAssignmentScheduleID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentScheduleId != v.Expected.UnifiedRoleAssignmentScheduleId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentScheduleId", v.Expected.UnifiedRoleAssignmentScheduleId, actual.UnifiedRoleAssignmentScheduleId)
		}

	}
}

func TestParseRoleManagementDirectoryRoleAssignmentScheduleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDirectoryRoleAssignmentScheduleId
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
			Input: "/roleManagement/directory/roleAssignmentSchedules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEaSsIgNmEnTsChEdUlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/directory/roleAssignmentSchedules/unifiedRoleAssignmentScheduleIdValue",
			Expected: &RoleManagementDirectoryRoleAssignmentScheduleId{
				UnifiedRoleAssignmentScheduleId: "unifiedRoleAssignmentScheduleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/directory/roleAssignmentSchedules/unifiedRoleAssignmentScheduleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEaSsIgNmEnTsChEdUlEs/uNiFiEdRoLeAsSiGnMeNtScHeDuLeIdVaLuE",
			Expected: &RoleManagementDirectoryRoleAssignmentScheduleId{
				UnifiedRoleAssignmentScheduleId: "uNiFiEdRoLeAsSiGnMeNtScHeDuLeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEaSsIgNmEnTsChEdUlEs/uNiFiEdRoLeAsSiGnMeNtScHeDuLeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDirectoryRoleAssignmentScheduleIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentScheduleId != v.Expected.UnifiedRoleAssignmentScheduleId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentScheduleId", v.Expected.UnifiedRoleAssignmentScheduleId, actual.UnifiedRoleAssignmentScheduleId)
		}

	}
}

func TestSegmentsForRoleManagementDirectoryRoleAssignmentScheduleId(t *testing.T) {
	segments := RoleManagementDirectoryRoleAssignmentScheduleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementDirectoryRoleAssignmentScheduleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
