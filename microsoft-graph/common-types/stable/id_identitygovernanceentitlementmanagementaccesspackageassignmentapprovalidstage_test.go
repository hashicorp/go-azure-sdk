package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageID("approvalId", "approvalStageId")

	if id.ApprovalId != "approvalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalId'", id.ApprovalId, "approvalId")
	}

	if id.ApprovalStageId != "approvalStageId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalStageId'", id.ApprovalStageId, "approvalStageId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageID("approvalId", "approvalStageId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/approvalId/stages/approvalStageId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/approvalId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/approvalId/stages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/approvalId/stages/approvalStageId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId{
				ApprovalId:      "approvalId",
				ApprovalStageId: "approvalStageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/approvalId/stages/approvalStageId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageID(v.Input)
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

		if actual.ApprovalStageId != v.Expected.ApprovalStageId {
			t.Fatalf("Expected %q but got %q for ApprovalStageId", v.Expected.ApprovalStageId, actual.ApprovalStageId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtApPrOvAlS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/approvalId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtApPrOvAlS/aPpRoVaLiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/approvalId/stages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtApPrOvAlS/aPpRoVaLiD/sTaGeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/approvalId/stages/approvalStageId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId{
				ApprovalId:      "approvalId",
				ApprovalStageId: "approvalStageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/approvalId/stages/approvalStageId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtApPrOvAlS/aPpRoVaLiD/sTaGeS/aPpRoVaLsTaGeId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId{
				ApprovalId:      "aPpRoVaLiD",
				ApprovalStageId: "aPpRoVaLsTaGeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtApPrOvAlS/aPpRoVaLiD/sTaGeS/aPpRoVaLsTaGeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageIDInsensitively(v.Input)
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

		if actual.ApprovalStageId != v.Expected.ApprovalStageId {
			t.Fatalf("Expected %q but got %q for ApprovalStageId", v.Expected.ApprovalStageId, actual.ApprovalStageId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
