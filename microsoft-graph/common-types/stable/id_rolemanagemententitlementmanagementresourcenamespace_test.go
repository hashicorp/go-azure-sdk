package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementResourceNamespaceId{}

func TestNewRoleManagementEntitlementManagementResourceNamespaceID(t *testing.T) {
	id := NewRoleManagementEntitlementManagementResourceNamespaceID("unifiedRbacResourceNamespaceId")

	if id.UnifiedRbacResourceNamespaceId != "unifiedRbacResourceNamespaceId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRbacResourceNamespaceId'", id.UnifiedRbacResourceNamespaceId, "unifiedRbacResourceNamespaceId")
	}
}

func TestFormatRoleManagementEntitlementManagementResourceNamespaceID(t *testing.T) {
	actual := NewRoleManagementEntitlementManagementResourceNamespaceID("unifiedRbacResourceNamespaceId").ID()
	expected := "/roleManagement/entitlementManagement/resourceNamespaces/unifiedRbacResourceNamespaceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEntitlementManagementResourceNamespaceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementResourceNamespaceId
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
			Input: "/roleManagement/entitlementManagement/resourceNamespaces",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/resourceNamespaces/unifiedRbacResourceNamespaceId",
			Expected: &RoleManagementEntitlementManagementResourceNamespaceId{
				UnifiedRbacResourceNamespaceId: "unifiedRbacResourceNamespaceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/resourceNamespaces/unifiedRbacResourceNamespaceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementResourceNamespaceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRbacResourceNamespaceId != v.Expected.UnifiedRbacResourceNamespaceId {
			t.Fatalf("Expected %q but got %q for UnifiedRbacResourceNamespaceId", v.Expected.UnifiedRbacResourceNamespaceId, actual.UnifiedRbacResourceNamespaceId)
		}

	}
}

func TestParseRoleManagementEntitlementManagementResourceNamespaceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementResourceNamespaceId
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
			Input: "/roleManagement/entitlementManagement/resourceNamespaces",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEnAmEsPaCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/resourceNamespaces/unifiedRbacResourceNamespaceId",
			Expected: &RoleManagementEntitlementManagementResourceNamespaceId{
				UnifiedRbacResourceNamespaceId: "unifiedRbacResourceNamespaceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/resourceNamespaces/unifiedRbacResourceNamespaceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiD",
			Expected: &RoleManagementEntitlementManagementResourceNamespaceId{
				UnifiedRbacResourceNamespaceId: "uNiFiEdRbAcReSoUrCeNaMeSpAcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementResourceNamespaceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRbacResourceNamespaceId != v.Expected.UnifiedRbacResourceNamespaceId {
			t.Fatalf("Expected %q but got %q for UnifiedRbacResourceNamespaceId", v.Expected.UnifiedRbacResourceNamespaceId, actual.UnifiedRbacResourceNamespaceId)
		}

	}
}

func TestSegmentsForRoleManagementEntitlementManagementResourceNamespaceId(t *testing.T) {
	segments := RoleManagementEntitlementManagementResourceNamespaceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEntitlementManagementResourceNamespaceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
