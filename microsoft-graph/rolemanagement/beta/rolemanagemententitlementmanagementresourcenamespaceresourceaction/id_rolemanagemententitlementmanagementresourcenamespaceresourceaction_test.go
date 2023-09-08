package rolemanagemententitlementmanagementresourcenamespaceresourceaction

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEntitlementManagementResourceNamespaceResourceActionId{}

func TestNewRoleManagementEntitlementManagementResourceNamespaceResourceActionID(t *testing.T) {
	id := NewRoleManagementEntitlementManagementResourceNamespaceResourceActionID("unifiedRbacResourceNamespaceIdValue", "unifiedRbacResourceActionIdValue")

	if id.UnifiedRbacResourceNamespaceId != "unifiedRbacResourceNamespaceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRbacResourceNamespaceId'", id.UnifiedRbacResourceNamespaceId, "unifiedRbacResourceNamespaceIdValue")
	}

	if id.UnifiedRbacResourceActionId != "unifiedRbacResourceActionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRbacResourceActionId'", id.UnifiedRbacResourceActionId, "unifiedRbacResourceActionIdValue")
	}
}

func TestFormatRoleManagementEntitlementManagementResourceNamespaceResourceActionID(t *testing.T) {
	actual := NewRoleManagementEntitlementManagementResourceNamespaceResourceActionID("unifiedRbacResourceNamespaceIdValue", "unifiedRbacResourceActionIdValue").ID()
	expected := "/roleManagement/entitlementManagement/resourceNamespaces/unifiedRbacResourceNamespaceIdValue/resourceActions/unifiedRbacResourceActionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEntitlementManagementResourceNamespaceResourceActionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementResourceNamespaceResourceActionId
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
			// Incomplete URI
			Input: "/roleManagement/entitlementManagement/resourceNamespaces/unifiedRbacResourceNamespaceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/entitlementManagement/resourceNamespaces/unifiedRbacResourceNamespaceIdValue/resourceActions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/resourceNamespaces/unifiedRbacResourceNamespaceIdValue/resourceActions/unifiedRbacResourceActionIdValue",
			Expected: &RoleManagementEntitlementManagementResourceNamespaceResourceActionId{
				UnifiedRbacResourceNamespaceId: "unifiedRbacResourceNamespaceIdValue",
				UnifiedRbacResourceActionId:    "unifiedRbacResourceActionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/resourceNamespaces/unifiedRbacResourceNamespaceIdValue/resourceActions/unifiedRbacResourceActionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementResourceNamespaceResourceActionID(v.Input)
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

		if actual.UnifiedRbacResourceActionId != v.Expected.UnifiedRbacResourceActionId {
			t.Fatalf("Expected %q but got %q for UnifiedRbacResourceActionId", v.Expected.UnifiedRbacResourceActionId, actual.UnifiedRbacResourceActionId)
		}

	}
}

func TestParseRoleManagementEntitlementManagementResourceNamespaceResourceActionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementResourceNamespaceResourceActionId
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
			// Incomplete URI
			Input: "/roleManagement/entitlementManagement/resourceNamespaces/unifiedRbacResourceNamespaceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/entitlementManagement/resourceNamespaces/unifiedRbacResourceNamespaceIdValue/resourceActions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiDvAlUe/rEsOuRcEaCtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/resourceNamespaces/unifiedRbacResourceNamespaceIdValue/resourceActions/unifiedRbacResourceActionIdValue",
			Expected: &RoleManagementEntitlementManagementResourceNamespaceResourceActionId{
				UnifiedRbacResourceNamespaceId: "unifiedRbacResourceNamespaceIdValue",
				UnifiedRbacResourceActionId:    "unifiedRbacResourceActionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/resourceNamespaces/unifiedRbacResourceNamespaceIdValue/resourceActions/unifiedRbacResourceActionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiDvAlUe/rEsOuRcEaCtIoNs/uNiFiEdRbAcReSoUrCeAcTiOnIdVaLuE",
			Expected: &RoleManagementEntitlementManagementResourceNamespaceResourceActionId{
				UnifiedRbacResourceNamespaceId: "uNiFiEdRbAcReSoUrCeNaMeSpAcEiDvAlUe",
				UnifiedRbacResourceActionId:    "uNiFiEdRbAcReSoUrCeAcTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiDvAlUe/rEsOuRcEaCtIoNs/uNiFiEdRbAcReSoUrCeAcTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementResourceNamespaceResourceActionIDInsensitively(v.Input)
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

		if actual.UnifiedRbacResourceActionId != v.Expected.UnifiedRbacResourceActionId {
			t.Fatalf("Expected %q but got %q for UnifiedRbacResourceActionId", v.Expected.UnifiedRbacResourceActionId, actual.UnifiedRbacResourceActionId)
		}

	}
}

func TestSegmentsForRoleManagementEntitlementManagementResourceNamespaceResourceActionId(t *testing.T) {
	segments := RoleManagementEntitlementManagementResourceNamespaceResourceActionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEntitlementManagementResourceNamespaceResourceActionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
