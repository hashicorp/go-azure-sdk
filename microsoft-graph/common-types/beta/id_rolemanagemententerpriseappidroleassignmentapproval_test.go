package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEnterpriseAppIdRoleAssignmentApprovalId{}

func TestNewRoleManagementEnterpriseAppIdRoleAssignmentApprovalID(t *testing.T) {
	id := NewRoleManagementEnterpriseAppIdRoleAssignmentApprovalID("rbacApplicationIdValue", "approvalIdValue")

	if id.RbacApplicationId != "rbacApplicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RbacApplicationId'", id.RbacApplicationId, "rbacApplicationIdValue")
	}

	if id.ApprovalId != "approvalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalId'", id.ApprovalId, "approvalIdValue")
	}
}

func TestFormatRoleManagementEnterpriseAppIdRoleAssignmentApprovalID(t *testing.T) {
	actual := NewRoleManagementEnterpriseAppIdRoleAssignmentApprovalID("rbacApplicationIdValue", "approvalIdValue").ID()
	expected := "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentApprovals/approvalIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEnterpriseAppIdRoleAssignmentApprovalID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppIdRoleAssignmentApprovalId
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentApprovals",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentApprovals/approvalIdValue",
			Expected: &RoleManagementEnterpriseAppIdRoleAssignmentApprovalId{
				RbacApplicationId: "rbacApplicationIdValue",
				ApprovalId:        "approvalIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentApprovals/approvalIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppIdRoleAssignmentApprovalID(v.Input)
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

		if actual.ApprovalId != v.Expected.ApprovalId {
			t.Fatalf("Expected %q but got %q for ApprovalId", v.Expected.ApprovalId, actual.ApprovalId)
		}

	}
}

func TestParseRoleManagementEnterpriseAppIdRoleAssignmentApprovalIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppIdRoleAssignmentApprovalId
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
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentApprovals",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEaSsIgNmEnTaPpRoVaLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentApprovals/approvalIdValue",
			Expected: &RoleManagementEnterpriseAppIdRoleAssignmentApprovalId{
				RbacApplicationId: "rbacApplicationIdValue",
				ApprovalId:        "approvalIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentApprovals/approvalIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEaSsIgNmEnTaPpRoVaLs/aPpRoVaLiDvAlUe",
			Expected: &RoleManagementEnterpriseAppIdRoleAssignmentApprovalId{
				RbacApplicationId: "rBaCaPpLiCaTiOnIdVaLuE",
				ApprovalId:        "aPpRoVaLiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEaSsIgNmEnTaPpRoVaLs/aPpRoVaLiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppIdRoleAssignmentApprovalIDInsensitively(v.Input)
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

		if actual.ApprovalId != v.Expected.ApprovalId {
			t.Fatalf("Expected %q but got %q for ApprovalId", v.Expected.ApprovalId, actual.ApprovalId)
		}

	}
}

func TestSegmentsForRoleManagementEnterpriseAppIdRoleAssignmentApprovalId(t *testing.T) {
	segments := RoleManagementEnterpriseAppIdRoleAssignmentApprovalId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEnterpriseAppIdRoleAssignmentApprovalId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
