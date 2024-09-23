package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdResourceNamespaceId{}

func TestNewRoleManagementEnterpriseAppIdResourceNamespaceID(t *testing.T) {
	id := NewRoleManagementEnterpriseAppIdResourceNamespaceID("rbacApplicationId", "unifiedRbacResourceNamespaceId")

	if id.RbacApplicationId != "rbacApplicationId" {
		t.Fatalf("Expected %q but got %q for Segment 'RbacApplicationId'", id.RbacApplicationId, "rbacApplicationId")
	}

	if id.UnifiedRbacResourceNamespaceId != "unifiedRbacResourceNamespaceId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRbacResourceNamespaceId'", id.UnifiedRbacResourceNamespaceId, "unifiedRbacResourceNamespaceId")
	}
}

func TestFormatRoleManagementEnterpriseAppIdResourceNamespaceID(t *testing.T) {
	actual := NewRoleManagementEnterpriseAppIdResourceNamespaceID("rbacApplicationId", "unifiedRbacResourceNamespaceId").ID()
	expected := "/roleManagement/enterpriseApps/rbacApplicationId/resourceNamespaces/unifiedRbacResourceNamespaceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEnterpriseAppIdResourceNamespaceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppIdResourceNamespaceId
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
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/resourceNamespaces/unifiedRbacResourceNamespaceId",
			Expected: &RoleManagementEnterpriseAppIdResourceNamespaceId{
				RbacApplicationId:              "rbacApplicationId",
				UnifiedRbacResourceNamespaceId: "unifiedRbacResourceNamespaceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/resourceNamespaces/unifiedRbacResourceNamespaceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppIdResourceNamespaceID(v.Input)
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

	}
}

func TestParseRoleManagementEnterpriseAppIdResourceNamespaceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppIdResourceNamespaceId
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
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/resourceNamespaces/unifiedRbacResourceNamespaceId",
			Expected: &RoleManagementEnterpriseAppIdResourceNamespaceId{
				RbacApplicationId:              "rbacApplicationId",
				UnifiedRbacResourceNamespaceId: "unifiedRbacResourceNamespaceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationId/resourceNamespaces/unifiedRbacResourceNamespaceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnId/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiD",
			Expected: &RoleManagementEnterpriseAppIdResourceNamespaceId{
				RbacApplicationId:              "rBaCaPpLiCaTiOnId",
				UnifiedRbacResourceNamespaceId: "uNiFiEdRbAcReSoUrCeNaMeSpAcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnId/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppIdResourceNamespaceIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForRoleManagementEnterpriseAppIdResourceNamespaceId(t *testing.T) {
	segments := RoleManagementEnterpriseAppIdResourceNamespaceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEnterpriseAppIdResourceNamespaceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
