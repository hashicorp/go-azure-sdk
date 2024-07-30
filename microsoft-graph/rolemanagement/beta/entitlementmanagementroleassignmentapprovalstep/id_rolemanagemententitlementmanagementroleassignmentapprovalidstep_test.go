package entitlementmanagementroleassignmentapprovalstep

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId{}

func TestNewRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepID(t *testing.T) {
	id := NewRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepID("approvalIdValue", "approvalStepIdValue")

	if id.ApprovalId != "approvalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalId'", id.ApprovalId, "approvalIdValue")
	}

	if id.ApprovalStepId != "approvalStepIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalStepId'", id.ApprovalStepId, "approvalStepIdValue")
	}
}

func TestFormatRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepID(t *testing.T) {
	actual := NewRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepID("approvalIdValue", "approvalStepIdValue").ID()
	expected := "/roleManagement/entitlementManagement/roleAssignmentApprovals/approvalIdValue/steps/approvalStepIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId
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
			Input: "/roleManagement/entitlementManagement/roleAssignmentApprovals",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/entitlementManagement/roleAssignmentApprovals/approvalIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/entitlementManagement/roleAssignmentApprovals/approvalIdValue/steps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/roleAssignmentApprovals/approvalIdValue/steps/approvalStepIdValue",
			Expected: &RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId{
				ApprovalId:     "approvalIdValue",
				ApprovalStepId: "approvalStepIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/roleAssignmentApprovals/approvalIdValue/steps/approvalStepIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApprovalId != v.Expected.ApprovalId {
			t.Fatalf("Expected %q but got %q for ApprovalId", v.Expected.ApprovalId, actual.ApprovalId)
		}

		if actual.ApprovalStepId != v.Expected.ApprovalStepId {
			t.Fatalf("Expected %q but got %q for ApprovalStepId", v.Expected.ApprovalStepId, actual.ApprovalStepId)
		}

	}
}

func TestParseRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId
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
			Input: "/roleManagement/entitlementManagement/roleAssignmentApprovals",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEaSsIgNmEnTaPpRoVaLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/entitlementManagement/roleAssignmentApprovals/approvalIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEaSsIgNmEnTaPpRoVaLs/aPpRoVaLiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/entitlementManagement/roleAssignmentApprovals/approvalIdValue/steps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEaSsIgNmEnTaPpRoVaLs/aPpRoVaLiDvAlUe/sTePs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/entitlementManagement/roleAssignmentApprovals/approvalIdValue/steps/approvalStepIdValue",
			Expected: &RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId{
				ApprovalId:     "approvalIdValue",
				ApprovalStepId: "approvalStepIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/entitlementManagement/roleAssignmentApprovals/approvalIdValue/steps/approvalStepIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEaSsIgNmEnTaPpRoVaLs/aPpRoVaLiDvAlUe/sTePs/aPpRoVaLsTePiDvAlUe",
			Expected: &RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId{
				ApprovalId:     "aPpRoVaLiDvAlUe",
				ApprovalStepId: "aPpRoVaLsTePiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eNtItLeMeNtMaNaGeMeNt/rOlEaSsIgNmEnTaPpRoVaLs/aPpRoVaLiDvAlUe/sTePs/aPpRoVaLsTePiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApprovalId != v.Expected.ApprovalId {
			t.Fatalf("Expected %q but got %q for ApprovalId", v.Expected.ApprovalId, actual.ApprovalId)
		}

		if actual.ApprovalStepId != v.Expected.ApprovalStepId {
			t.Fatalf("Expected %q but got %q for ApprovalStepId", v.Expected.ApprovalStepId, actual.ApprovalStepId)
		}

	}
}

func TestSegmentsForRoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId(t *testing.T) {
	segments := RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementEntitlementManagementRoleAssignmentApprovalIdStepId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
