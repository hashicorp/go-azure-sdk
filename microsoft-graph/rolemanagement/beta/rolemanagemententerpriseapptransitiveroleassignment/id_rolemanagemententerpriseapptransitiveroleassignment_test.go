package rolemanagemententerpriseapptransitiveroleassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppTransitiveRoleAssignmentId{}

func TestNewRoleManagementEnterpriseAppTransitiveRoleAssignmentID(t *testing.T) {
	id := NewRoleManagementEnterpriseAppTransitiveRoleAssignmentID("rbacApplicationIdValue", "unifiedRoleAssignmentIdValue")

	if id.RbacApplicationId != "rbacApplicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RbacApplicationId'", id.RbacApplicationId, "rbacApplicationIdValue")
	}

	if id.UnifiedRoleAssignmentId != "unifiedRoleAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentId'", id.UnifiedRoleAssignmentId, "unifiedRoleAssignmentIdValue")
	}
}

func TestFormatRoleManagementEnterpriseAppTransitiveRoleAssignmentID(t *testing.T) {
	actual := NewRoleManagementEnterpriseAppTransitiveRoleAssignmentID("rbacApplicationIdValue", "unifiedRoleAssignmentIdValue").ID()
	expected := "/roleManagement/enterpriseApps/rbacApplicationIdValue/transitiveRoleAssignments/unifiedRoleAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEnterpriseAppTransitiveRoleAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppTransitiveRoleAssignmentId
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/transitiveRoleAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/transitiveRoleAssignments/unifiedRoleAssignmentIdValue",
			Expected: &RoleManagementEnterpriseAppTransitiveRoleAssignmentId{
				RbacApplicationId:       "rbacApplicationIdValue",
				UnifiedRoleAssignmentId: "unifiedRoleAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/transitiveRoleAssignments/unifiedRoleAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppTransitiveRoleAssignmentID(v.Input)
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

		if actual.UnifiedRoleAssignmentId != v.Expected.UnifiedRoleAssignmentId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentId", v.Expected.UnifiedRoleAssignmentId, actual.UnifiedRoleAssignmentId)
		}

	}
}

func TestParseRoleManagementEnterpriseAppTransitiveRoleAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppTransitiveRoleAssignmentId
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/transitiveRoleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/tRaNsItIvErOlEaSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/transitiveRoleAssignments/unifiedRoleAssignmentIdValue",
			Expected: &RoleManagementEnterpriseAppTransitiveRoleAssignmentId{
				RbacApplicationId:       "rbacApplicationIdValue",
				UnifiedRoleAssignmentId: "unifiedRoleAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/transitiveRoleAssignments/unifiedRoleAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/tRaNsItIvErOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtIdVaLuE",
			Expected: &RoleManagementEnterpriseAppTransitiveRoleAssignmentId{
				RbacApplicationId:       "rBaCaPpLiCaTiOnIdVaLuE",
				UnifiedRoleAssignmentId: "uNiFiEdRoLeAsSiGnMeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/tRaNsItIvErOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppTransitiveRoleAssignmentIDInsensitively(v.Input)
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

		if actual.UnifiedRoleAssignmentId != v.Expected.UnifiedRoleAssignmentId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentId", v.Expected.UnifiedRoleAssignmentId, actual.UnifiedRoleAssignmentId)
		}

	}
}

func TestSegmentsForRoleManagementEnterpriseAppTransitiveRoleAssignmentId(t *testing.T) {
	segments := RoleManagementEnterpriseAppTransitiveRoleAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEnterpriseAppTransitiveRoleAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
