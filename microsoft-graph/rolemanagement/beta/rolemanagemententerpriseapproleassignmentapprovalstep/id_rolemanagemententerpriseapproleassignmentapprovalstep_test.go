package rolemanagemententerpriseapproleassignmentapprovalstep

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppRoleAssignmentApprovalStepId{}

func TestNewRoleManagementEnterpriseAppRoleAssignmentApprovalStepID(t *testing.T) {
	id := NewRoleManagementEnterpriseAppRoleAssignmentApprovalStepID("rbacApplicationIdValue", "approvalIdValue", "approvalStepIdValue")

	if id.RbacApplicationId != "rbacApplicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RbacApplicationId'", id.RbacApplicationId, "rbacApplicationIdValue")
	}

	if id.ApprovalId != "approvalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalId'", id.ApprovalId, "approvalIdValue")
	}

	if id.ApprovalStepId != "approvalStepIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalStepId'", id.ApprovalStepId, "approvalStepIdValue")
	}
}

func TestFormatRoleManagementEnterpriseAppRoleAssignmentApprovalStepID(t *testing.T) {
	actual := NewRoleManagementEnterpriseAppRoleAssignmentApprovalStepID("rbacApplicationIdValue", "approvalIdValue", "approvalStepIdValue").ID()
	expected := "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentApprovals/approvalIdValue/steps/approvalStepIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEnterpriseAppRoleAssignmentApprovalStepID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppRoleAssignmentApprovalStepId
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
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentApprovals/approvalIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentApprovals/approvalIdValue/steps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentApprovals/approvalIdValue/steps/approvalStepIdValue",
			Expected: &RoleManagementEnterpriseAppRoleAssignmentApprovalStepId{
				RbacApplicationId: "rbacApplicationIdValue",
				ApprovalId:        "approvalIdValue",
				ApprovalStepId:    "approvalStepIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentApprovals/approvalIdValue/steps/approvalStepIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppRoleAssignmentApprovalStepID(v.Input)
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

		if actual.ApprovalStepId != v.Expected.ApprovalStepId {
			t.Fatalf("Expected %q but got %q for ApprovalStepId", v.Expected.ApprovalStepId, actual.ApprovalStepId)
		}

	}
}

func TestParseRoleManagementEnterpriseAppRoleAssignmentApprovalStepIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEnterpriseAppRoleAssignmentApprovalStepId
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
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentApprovals/approvalIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEaSsIgNmEnTaPpRoVaLs/aPpRoVaLiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentApprovals/approvalIdValue/steps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEaSsIgNmEnTaPpRoVaLs/aPpRoVaLiDvAlUe/sTePs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentApprovals/approvalIdValue/steps/approvalStepIdValue",
			Expected: &RoleManagementEnterpriseAppRoleAssignmentApprovalStepId{
				RbacApplicationId: "rbacApplicationIdValue",
				ApprovalId:        "approvalIdValue",
				ApprovalStepId:    "approvalStepIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/enterpriseApps/rbacApplicationIdValue/roleAssignmentApprovals/approvalIdValue/steps/approvalStepIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEaSsIgNmEnTaPpRoVaLs/aPpRoVaLiDvAlUe/sTePs/aPpRoVaLsTePiDvAlUe",
			Expected: &RoleManagementEnterpriseAppRoleAssignmentApprovalStepId{
				RbacApplicationId: "rBaCaPpLiCaTiOnIdVaLuE",
				ApprovalId:        "aPpRoVaLiDvAlUe",
				ApprovalStepId:    "aPpRoVaLsTePiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtErPrIsEaPpS/rBaCaPpLiCaTiOnIdVaLuE/rOlEaSsIgNmEnTaPpRoVaLs/aPpRoVaLiDvAlUe/sTePs/aPpRoVaLsTePiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEnterpriseAppRoleAssignmentApprovalStepIDInsensitively(v.Input)
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

		if actual.ApprovalStepId != v.Expected.ApprovalStepId {
			t.Fatalf("Expected %q but got %q for ApprovalStepId", v.Expected.ApprovalStepId, actual.ApprovalStepId)
		}

	}
}

func TestSegmentsForRoleManagementEnterpriseAppRoleAssignmentApprovalStepId(t *testing.T) {
	segments := RoleManagementEnterpriseAppRoleAssignmentApprovalStepId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEnterpriseAppRoleAssignmentApprovalStepId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
