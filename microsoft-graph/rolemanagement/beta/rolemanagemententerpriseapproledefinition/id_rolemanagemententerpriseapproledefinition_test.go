package rolemanagemententerpriseapproledefinition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppRoleDefinitionId{}

func TestNewRoleManagementEnterpriseAppRoleDefinitionID(t *testing.T) {
	id := NewRoleManagementEnterpriseAppRoleDefinitionID("rbacApplicationIdValue", "unifiedRoleDefinitionIdValue")

	if id.RbacApplicationId != "rbacApplicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RbacApplicationId'", id.RbacApplicationId, "rbacApplicationIdValue")
	}

	if id.UnifiedRoleDefinitionId != "unifiedRoleDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleDefinitionId'", id.UnifiedRoleDefinitionId, "unifiedRoleDefinitionIdValue")
	}
}

func TestFormatRoleManagementEnterpriseAppRoleDefinitionID(t *testing.T) {
	actual := NewRoleManagementEnterpriseAppRoleDefinitionID("rbacApplicationIdValue", "unifiedRoleDefinitionIdValue").ID()
	expected := "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleDefinitions/unifiedRoleDefinitionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEnterpriseAppRoleDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppRoleDefinitionId
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleDefinitions/unifiedRoleDefinitionIdValue",
			Expected: &RoleManagementEnterpriseAppRoleDefinitionId{
				RbacApplicationId:       "rbacApplicationIdValue",
				UnifiedRoleDefinitionId: "unifiedRoleDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleDefinitions/unifiedRoleDefinitionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppRoleDefinitionID(v.Input)
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

		if actual.UnifiedRoleDefinitionId != v.Expected.UnifiedRoleDefinitionId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleDefinitionId", v.Expected.UnifiedRoleDefinitionId, actual.UnifiedRoleDefinitionId)
		}

	}
}

func TestParseRoleManagementEnterpriseAppRoleDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppRoleDefinitionId
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEdEfInItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleDefinitions/unifiedRoleDefinitionIdValue",
			Expected: &RoleManagementEnterpriseAppRoleDefinitionId{
				RbacApplicationId:       "rbacApplicationIdValue",
				UnifiedRoleDefinitionId: "unifiedRoleDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleDefinitions/unifiedRoleDefinitionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEdEfInItIoNs/uNiFiEdRoLeDeFiNiTiOnIdVaLuE",
			Expected: &RoleManagementEnterpriseAppRoleDefinitionId{
				RbacApplicationId:       "rBaCaPpLiCaTiOnIdVaLuE",
				UnifiedRoleDefinitionId: "uNiFiEdRoLeDeFiNiTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEdEfInItIoNs/uNiFiEdRoLeDeFiNiTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppRoleDefinitionIDInsensitively(v.Input)
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

		if actual.UnifiedRoleDefinitionId != v.Expected.UnifiedRoleDefinitionId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleDefinitionId", v.Expected.UnifiedRoleDefinitionId, actual.UnifiedRoleDefinitionId)
		}

	}
}

func TestSegmentsForRoleManagementEnterpriseAppRoleDefinitionId(t *testing.T) {
	segments := RoleManagementEnterpriseAppRoleDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEnterpriseAppRoleDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
