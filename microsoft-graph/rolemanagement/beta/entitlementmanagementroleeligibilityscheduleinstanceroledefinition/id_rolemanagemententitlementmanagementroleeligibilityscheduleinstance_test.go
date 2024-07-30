package entitlementmanagementroleeligibilityscheduleinstanceroledefinition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId{}

func TestNewRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceID(t *testing.T) {
	id := NewRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceID("unifiedRoleEligibilityScheduleInstanceIdValue")

	if id.UnifiedRoleEligibilityScheduleInstanceId != "unifiedRoleEligibilityScheduleInstanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleEligibilityScheduleInstanceId'", id.UnifiedRoleEligibilityScheduleInstanceId, "unifiedRoleEligibilityScheduleInstanceIdValue")
	}
}

func TestFormatRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceID(t *testing.T) {
	actual := NewRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceID("unifiedRoleEligibilityScheduleInstanceIdValue").ID()
	expected := "/roleManagement/entitlementManagement/roleEligibilityScheduleInstances/unifiedRoleEligibilityScheduleInstanceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId
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
			Input: "/roleManagement/entitlementManagement/roleEligibilityScheduleInstances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/roleEligibilityScheduleInstances/unifiedRoleEligibilityScheduleInstanceIdValue",
			Expected: &RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId{
				UnifiedRoleEligibilityScheduleInstanceId: "unifiedRoleEligibilityScheduleInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/roleEligibilityScheduleInstances/unifiedRoleEligibilityScheduleInstanceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleEligibilityScheduleInstanceId != v.Expected.UnifiedRoleEligibilityScheduleInstanceId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleEligibilityScheduleInstanceId", v.Expected.UnifiedRoleEligibilityScheduleInstanceId, actual.UnifiedRoleEligibilityScheduleInstanceId)
		}

	}
}

func TestParseRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId
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
			Input: "/roleManagement/entitlementManagement/roleEligibilityScheduleInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEeLiGiBiLiTyScHeDuLeInStAnCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/roleEligibilityScheduleInstances/unifiedRoleEligibilityScheduleInstanceIdValue",
			Expected: &RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId{
				UnifiedRoleEligibilityScheduleInstanceId: "unifiedRoleEligibilityScheduleInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/roleEligibilityScheduleInstances/unifiedRoleEligibilityScheduleInstanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEeLiGiBiLiTyScHeDuLeInStAnCeS/uNiFiEdRoLeElIgIbIlItYsChEdUlEiNsTaNcEiDvAlUe",
			Expected: &RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId{
				UnifiedRoleEligibilityScheduleInstanceId: "uNiFiEdRoLeElIgIbIlItYsChEdUlEiNsTaNcEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEeLiGiBiLiTyScHeDuLeInStAnCeS/uNiFiEdRoLeElIgIbIlItYsChEdUlEiNsTaNcEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleEligibilityScheduleInstanceId != v.Expected.UnifiedRoleEligibilityScheduleInstanceId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleEligibilityScheduleInstanceId", v.Expected.UnifiedRoleEligibilityScheduleInstanceId, actual.UnifiedRoleEligibilityScheduleInstanceId)
		}

	}
}

func TestSegmentsForRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId(t *testing.T) {
	segments := RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
