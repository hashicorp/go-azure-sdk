package entitlementmanagementroleeligibilityscheduleappscope

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementRoleEligibilityScheduleId{}

func TestNewRoleManagementEntitlementManagementRoleEligibilityScheduleID(t *testing.T) {
	id := NewRoleManagementEntitlementManagementRoleEligibilityScheduleID("unifiedRoleEligibilityScheduleIdValue")

	if id.UnifiedRoleEligibilityScheduleId != "unifiedRoleEligibilityScheduleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleEligibilityScheduleId'", id.UnifiedRoleEligibilityScheduleId, "unifiedRoleEligibilityScheduleIdValue")
	}
}

func TestFormatRoleManagementEntitlementManagementRoleEligibilityScheduleID(t *testing.T) {
	actual := NewRoleManagementEntitlementManagementRoleEligibilityScheduleID("unifiedRoleEligibilityScheduleIdValue").ID()
	expected := "/roleManagement/entitlementManagement/roleEligibilitySchedules/unifiedRoleEligibilityScheduleIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEntitlementManagementRoleEligibilityScheduleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementRoleEligibilityScheduleId
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
			Input: "/roleManagement/entitlementManagement/roleEligibilitySchedules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/roleEligibilitySchedules/unifiedRoleEligibilityScheduleIdValue",
			Expected: &RoleManagementEntitlementManagementRoleEligibilityScheduleId{
				UnifiedRoleEligibilityScheduleId: "unifiedRoleEligibilityScheduleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/roleEligibilitySchedules/unifiedRoleEligibilityScheduleIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementRoleEligibilityScheduleID(v.Input)
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

func TestParseRoleManagementEntitlementManagementRoleEligibilityScheduleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementRoleEligibilityScheduleId
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
			Input: "/roleManagement/entitlementManagement/roleEligibilitySchedules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEeLiGiBiLiTyScHeDuLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/roleEligibilitySchedules/unifiedRoleEligibilityScheduleIdValue",
			Expected: &RoleManagementEntitlementManagementRoleEligibilityScheduleId{
				UnifiedRoleEligibilityScheduleId: "unifiedRoleEligibilityScheduleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/roleEligibilitySchedules/unifiedRoleEligibilityScheduleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEeLiGiBiLiTyScHeDuLeS/uNiFiEdRoLeElIgIbIlItYsChEdUlEiDvAlUe",
			Expected: &RoleManagementEntitlementManagementRoleEligibilityScheduleId{
				UnifiedRoleEligibilityScheduleId: "uNiFiEdRoLeElIgIbIlItYsChEdUlEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEeLiGiBiLiTyScHeDuLeS/uNiFiEdRoLeElIgIbIlItYsChEdUlEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementRoleEligibilityScheduleIDInsensitively(v.Input)
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

func TestSegmentsForRoleManagementEntitlementManagementRoleEligibilityScheduleId(t *testing.T) {
	segments := RoleManagementEntitlementManagementRoleEligibilityScheduleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEntitlementManagementRoleEligibilityScheduleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
