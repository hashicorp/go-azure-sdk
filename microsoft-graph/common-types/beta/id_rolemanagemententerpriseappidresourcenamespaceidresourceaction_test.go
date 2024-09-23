package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId{}

func TestNewRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionID(t *testing.T) {
	id := NewRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionID("rbacApplicationId", "unifiedRbacResourceNamespaceId", "unifiedRbacResourceActionId")

	if id.RbacApplicationId != "rbacApplicationId" {
		t.Fatalf("Expected %q but got %q for Segment 'RbacApplicationId'", id.RbacApplicationId, "rbacApplicationId")
	}

	if id.UnifiedRbacResourceNamespaceId != "unifiedRbacResourceNamespaceId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRbacResourceNamespaceId'", id.UnifiedRbacResourceNamespaceId, "unifiedRbacResourceNamespaceId")
	}

	if id.UnifiedRbacResourceActionId != "unifiedRbacResourceActionId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRbacResourceActionId'", id.UnifiedRbacResourceActionId, "unifiedRbacResourceActionId")
	}
}

func TestFormatRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionID(t *testing.T) {
	actual := NewRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionID("rbacApplicationId", "unifiedRbacResourceNamespaceId", "unifiedRbacResourceActionId").ID()
	expected := "/roleManagement/enterpriseApps/rbacApplicationId/resourceNamespaces/unifiedRbacResourceNamespaceId/resourceActions/unifiedRbacResourceActionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId
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
			Input: "/roleManagement/enterpriseApps",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/resourceNamespaces",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/resourceNamespaces/unifiedRbacResourceNamespaceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/resourceNamespaces/unifiedRbacResourceNamespaceId/resourceActions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/resourceNamespaces/unifiedRbacResourceNamespaceId/resourceActions/unifiedRbacResourceActionId",
			Expected: &RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId{
				RbacApplicationId:              "rbacApplicationId",
				UnifiedRbacResourceNamespaceId: "unifiedRbacResourceNamespaceId",
				UnifiedRbacResourceActionId:    "unifiedRbacResourceActionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/resourceNamespaces/unifiedRbacResourceNamespaceId/resourceActions/unifiedRbacResourceActionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RbacApplicationId != v.Expected.RbacApplicationId {
			t.Fatalf("Expected %q but got %q for RbacApplicationId", v.Expected.RbacApplicationId, actual.RbacApplicationId)
		}

		if actual.UnifiedRbacResourceNamespaceId != v.Expected.UnifiedRbacResourceNamespaceId {
			t.Fatalf("Expected %q but got %q for UnifiedRbacResourceNamespaceId", v.Expected.UnifiedRbacResourceNamespaceId, actual.UnifiedRbacResourceNamespaceId)
		}

		if actual.UnifiedRbacResourceActionId != v.Expected.UnifiedRbacResourceActionId {
			t.Fatalf("Expected %q but got %q for UnifiedRbacResourceActionId", v.Expected.UnifiedRbacResourceActionId, actual.UnifiedRbacResourceActionId)
		}

	}
}

func TestParseRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId
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
			Input: "/roleManagement/enterpriseApps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/resourceNamespaces",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnId/rEsOuRcEnAmEsPaCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/resourceNamespaces/unifiedRbacResourceNamespaceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnId/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/resourceNamespaces/unifiedRbacResourceNamespaceId/resourceActions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnId/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiD/rEsOuRcEaCtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/resourceNamespaces/unifiedRbacResourceNamespaceId/resourceActions/unifiedRbacResourceActionId",
			Expected: &RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId{
				RbacApplicationId:              "rbacApplicationId",
				UnifiedRbacResourceNamespaceId: "unifiedRbacResourceNamespaceId",
				UnifiedRbacResourceActionId:    "unifiedRbacResourceActionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/resourceNamespaces/unifiedRbacResourceNamespaceId/resourceActions/unifiedRbacResourceActionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnId/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiD/rEsOuRcEaCtIoNs/uNiFiEdRbAcReSoUrCeAcTiOnId",
			Expected: &RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId{
				RbacApplicationId:              "rBaCaPpLiCaTiOnId",
				UnifiedRbacResourceNamespaceId: "uNiFiEdRbAcReSoUrCeNaMeSpAcEiD",
				UnifiedRbacResourceActionId:    "uNiFiEdRbAcReSoUrCeAcTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnId/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiD/rEsOuRcEaCtIoNs/uNiFiEdRbAcReSoUrCeAcTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RbacApplicationId != v.Expected.RbacApplicationId {
			t.Fatalf("Expected %q but got %q for RbacApplicationId", v.Expected.RbacApplicationId, actual.RbacApplicationId)
		}

		if actual.UnifiedRbacResourceNamespaceId != v.Expected.UnifiedRbacResourceNamespaceId {
			t.Fatalf("Expected %q but got %q for UnifiedRbacResourceNamespaceId", v.Expected.UnifiedRbacResourceNamespaceId, actual.UnifiedRbacResourceNamespaceId)
		}

		if actual.UnifiedRbacResourceActionId != v.Expected.UnifiedRbacResourceActionId {
			t.Fatalf("Expected %q but got %q for UnifiedRbacResourceActionId", v.Expected.UnifiedRbacResourceActionId, actual.UnifiedRbacResourceActionId)
		}

	}
}

func TestSegmentsForRoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId(t *testing.T) {
	segments := RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
