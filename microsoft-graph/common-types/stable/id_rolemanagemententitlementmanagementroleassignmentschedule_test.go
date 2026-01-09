package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementRoleAssignmentScheduleId{}

func TestNewRoleManagementEntitlementManagementRoleAssignmentScheduleID(t *testing.T) {
	id := NewRoleManagementEntitlementManagementRoleAssignmentScheduleID("unifiedRoleAssignmentScheduleId")

	if id.UnifiedRoleAssignmentScheduleId != "unifiedRoleAssignmentScheduleId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentScheduleId'", id.UnifiedRoleAssignmentScheduleId, "unifiedRoleAssignmentScheduleId")
	}
}

func TestFormatRoleManagementEntitlementManagementRoleAssignmentScheduleID(t *testing.T) {
	actual := NewRoleManagementEntitlementManagementRoleAssignmentScheduleID("unifiedRoleAssignmentScheduleId").ID()
	expected := "/roleManagement/entitlementManagement/roleAssignmentSchedules/unifiedRoleAssignmentScheduleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEntitlementManagementRoleAssignmentScheduleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementRoleAssignmentScheduleId
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
			Input: "/roleManagement/entitlementManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/entitlementManagement/roleAssignmentSchedules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/roleAssignmentSchedules/unifiedRoleAssignmentScheduleId",
			Expected: &RoleManagementEntitlementManagementRoleAssignmentScheduleId{
				UnifiedRoleAssignmentScheduleId: "unifiedRoleAssignmentScheduleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/roleAssignmentSchedules/unifiedRoleAssignmentScheduleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementRoleAssignmentScheduleID(v.Input)
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

func TestParseRoleManagementEntitlementManagementRoleAssignmentScheduleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementRoleAssignmentScheduleId
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
			Input: "/roleManagement/entitlementManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/entitlementManagement/roleAssignmentSchedules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEaSsIgNmEnTsChEdUlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/roleAssignmentSchedules/unifiedRoleAssignmentScheduleId",
			Expected: &RoleManagementEntitlementManagementRoleAssignmentScheduleId{
				UnifiedRoleAssignmentScheduleId: "unifiedRoleAssignmentScheduleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/roleAssignmentSchedules/unifiedRoleAssignmentScheduleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEaSsIgNmEnTsChEdUlEs/uNiFiEdRoLeAsSiGnMeNtScHeDuLeId",
			Expected: &RoleManagementEntitlementManagementRoleAssignmentScheduleId{
				UnifiedRoleAssignmentScheduleId: "uNiFiEdRoLeAsSiGnMeNtScHeDuLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEaSsIgNmEnTsChEdUlEs/uNiFiEdRoLeAsSiGnMeNtScHeDuLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementRoleAssignmentScheduleIDInsensitively(v.Input)
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

func TestSegmentsForRoleManagementEntitlementManagementRoleAssignmentScheduleId(t *testing.T) {
	segments := RoleManagementEntitlementManagementRoleAssignmentScheduleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEntitlementManagementRoleAssignmentScheduleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
