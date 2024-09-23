package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId{}

func TestNewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestID(t *testing.T) {
	id := NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestID("unifiedRoleEligibilityScheduleRequestId")

	if id.UnifiedRoleEligibilityScheduleRequestId != "unifiedRoleEligibilityScheduleRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleEligibilityScheduleRequestId'", id.UnifiedRoleEligibilityScheduleRequestId, "unifiedRoleEligibilityScheduleRequestId")
	}
}

func TestFormatRoleManagementEntitlementManagementRoleEligibilityScheduleRequestID(t *testing.T) {
	actual := NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestID("unifiedRoleEligibilityScheduleRequestId").ID()
	expected := "/roleManagement/entitlementManagement/roleEligibilityScheduleRequests/unifiedRoleEligibilityScheduleRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEntitlementManagementRoleEligibilityScheduleRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId
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
			Input: "/roleManagement/entitlementManagement/roleEligibilityScheduleRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/roleEligibilityScheduleRequests/unifiedRoleEligibilityScheduleRequestId",
			Expected: &RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId{
				UnifiedRoleEligibilityScheduleRequestId: "unifiedRoleEligibilityScheduleRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/roleEligibilityScheduleRequests/unifiedRoleEligibilityScheduleRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementRoleEligibilityScheduleRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleEligibilityScheduleRequestId != v.Expected.UnifiedRoleEligibilityScheduleRequestId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleEligibilityScheduleRequestId", v.Expected.UnifiedRoleEligibilityScheduleRequestId, actual.UnifiedRoleEligibilityScheduleRequestId)
		}

	}
}

func TestParseRoleManagementEntitlementManagementRoleEligibilityScheduleRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId
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
			Input: "/roleManagement/entitlementManagement/roleEligibilityScheduleRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEeLiGiBiLiTyScHeDuLeReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/roleEligibilityScheduleRequests/unifiedRoleEligibilityScheduleRequestId",
			Expected: &RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId{
				UnifiedRoleEligibilityScheduleRequestId: "unifiedRoleEligibilityScheduleRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/roleEligibilityScheduleRequests/unifiedRoleEligibilityScheduleRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEeLiGiBiLiTyScHeDuLeReQuEsTs/uNiFiEdRoLeElIgIbIlItYsChEdUlErEqUeStId",
			Expected: &RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId{
				UnifiedRoleEligibilityScheduleRequestId: "uNiFiEdRoLeElIgIbIlItYsChEdUlErEqUeStId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEeLiGiBiLiTyScHeDuLeReQuEsTs/uNiFiEdRoLeElIgIbIlItYsChEdUlErEqUeStId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementRoleEligibilityScheduleRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleEligibilityScheduleRequestId != v.Expected.UnifiedRoleEligibilityScheduleRequestId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleEligibilityScheduleRequestId", v.Expected.UnifiedRoleEligibilityScheduleRequestId, actual.UnifiedRoleEligibilityScheduleRequestId)
		}

	}
}

func TestSegmentsForRoleManagementEntitlementManagementRoleEligibilityScheduleRequestId(t *testing.T) {
	segments := RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
