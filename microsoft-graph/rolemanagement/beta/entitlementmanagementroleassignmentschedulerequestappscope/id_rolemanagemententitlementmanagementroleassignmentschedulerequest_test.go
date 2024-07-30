package entitlementmanagementroleassignmentschedulerequestappscope

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId{}

func TestNewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestID(t *testing.T) {
	id := NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestID("unifiedRoleAssignmentScheduleRequestIdValue")

	if id.UnifiedRoleAssignmentScheduleRequestId != "unifiedRoleAssignmentScheduleRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentScheduleRequestId'", id.UnifiedRoleAssignmentScheduleRequestId, "unifiedRoleAssignmentScheduleRequestIdValue")
	}
}

func TestFormatRoleManagementEntitlementManagementRoleAssignmentScheduleRequestID(t *testing.T) {
	actual := NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestID("unifiedRoleAssignmentScheduleRequestIdValue").ID()
	expected := "/roleManagement/entitlementManagement/roleAssignmentScheduleRequests/unifiedRoleAssignmentScheduleRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEntitlementManagementRoleAssignmentScheduleRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId
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
			Input: "/roleManagement/entitlementManagement/roleAssignmentScheduleRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/roleAssignmentScheduleRequests/unifiedRoleAssignmentScheduleRequestIdValue",
			Expected: &RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId{
				UnifiedRoleAssignmentScheduleRequestId: "unifiedRoleAssignmentScheduleRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/roleAssignmentScheduleRequests/unifiedRoleAssignmentScheduleRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementRoleAssignmentScheduleRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentScheduleRequestId != v.Expected.UnifiedRoleAssignmentScheduleRequestId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentScheduleRequestId", v.Expected.UnifiedRoleAssignmentScheduleRequestId, actual.UnifiedRoleAssignmentScheduleRequestId)
		}

	}
}

func TestParseRoleManagementEntitlementManagementRoleAssignmentScheduleRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId
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
			Input: "/roleManagement/entitlementManagement/roleAssignmentScheduleRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEaSsIgNmEnTsChEdUlErEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/roleAssignmentScheduleRequests/unifiedRoleAssignmentScheduleRequestIdValue",
			Expected: &RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId{
				UnifiedRoleAssignmentScheduleRequestId: "unifiedRoleAssignmentScheduleRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/roleAssignmentScheduleRequests/unifiedRoleAssignmentScheduleRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEaSsIgNmEnTsChEdUlErEqUeStS/uNiFiEdRoLeAsSiGnMeNtScHeDuLeReQuEsTiDvAlUe",
			Expected: &RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId{
				UnifiedRoleAssignmentScheduleRequestId: "uNiFiEdRoLeAsSiGnMeNtScHeDuLeReQuEsTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEaSsIgNmEnTsChEdUlErEqUeStS/uNiFiEdRoLeAsSiGnMeNtScHeDuLeReQuEsTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementRoleAssignmentScheduleRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentScheduleRequestId != v.Expected.UnifiedRoleAssignmentScheduleRequestId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentScheduleRequestId", v.Expected.UnifiedRoleAssignmentScheduleRequestId, actual.UnifiedRoleAssignmentScheduleRequestId)
		}

	}
}

func TestSegmentsForRoleManagementEntitlementManagementRoleAssignmentScheduleRequestId(t *testing.T) {
	segments := RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
