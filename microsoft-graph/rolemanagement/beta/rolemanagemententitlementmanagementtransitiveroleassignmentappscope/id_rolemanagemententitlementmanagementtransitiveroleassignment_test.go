package rolemanagemententitlementmanagementtransitiveroleassignmentappscope

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEntitlementManagementTransitiveRoleAssignmentId{}

func TestNewRoleManagementEntitlementManagementTransitiveRoleAssignmentID(t *testing.T) {
	id := NewRoleManagementEntitlementManagementTransitiveRoleAssignmentID("unifiedRoleAssignmentIdValue")

	if id.UnifiedRoleAssignmentId != "unifiedRoleAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentId'", id.UnifiedRoleAssignmentId, "unifiedRoleAssignmentIdValue")
	}
}

func TestFormatRoleManagementEntitlementManagementTransitiveRoleAssignmentID(t *testing.T) {
	actual := NewRoleManagementEntitlementManagementTransitiveRoleAssignmentID("unifiedRoleAssignmentIdValue").ID()
	expected := "/roleManagement/entitlementManagement/transitiveRoleAssignments/unifiedRoleAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEntitlementManagementTransitiveRoleAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementTransitiveRoleAssignmentId
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
			Input: "/roleManagement/entitlementManagement/transitiveRoleAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/transitiveRoleAssignments/unifiedRoleAssignmentIdValue",
			Expected: &RoleManagementEntitlementManagementTransitiveRoleAssignmentId{
				UnifiedRoleAssignmentId: "unifiedRoleAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/transitiveRoleAssignments/unifiedRoleAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementTransitiveRoleAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentId != v.Expected.UnifiedRoleAssignmentId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentId", v.Expected.UnifiedRoleAssignmentId, actual.UnifiedRoleAssignmentId)
		}

	}
}

func TestParseRoleManagementEntitlementManagementTransitiveRoleAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementTransitiveRoleAssignmentId
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
			Input: "/roleManagement/entitlementManagement/transitiveRoleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/tRaNsItIvErOlEaSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/transitiveRoleAssignments/unifiedRoleAssignmentIdValue",
			Expected: &RoleManagementEntitlementManagementTransitiveRoleAssignmentId{
				UnifiedRoleAssignmentId: "unifiedRoleAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/transitiveRoleAssignments/unifiedRoleAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/tRaNsItIvErOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtIdVaLuE",
			Expected: &RoleManagementEntitlementManagementTransitiveRoleAssignmentId{
				UnifiedRoleAssignmentId: "uNiFiEdRoLeAsSiGnMeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/tRaNsItIvErOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementTransitiveRoleAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentId != v.Expected.UnifiedRoleAssignmentId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentId", v.Expected.UnifiedRoleAssignmentId, actual.UnifiedRoleAssignmentId)
		}

	}
}

func TestSegmentsForRoleManagementEntitlementManagementTransitiveRoleAssignmentId(t *testing.T) {
	segments := RoleManagementEntitlementManagementTransitiveRoleAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEntitlementManagementTransitiveRoleAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
