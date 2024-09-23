package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepID("approvalId", "approvalStepId")

	if id.ApprovalId != "approvalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalId'", id.ApprovalId, "approvalId")
	}

	if id.ApprovalStepId != "approvalStepId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalStepId'", id.ApprovalStepId, "approvalStepId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepID("approvalId", "approvalStepId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/approvalId/steps/approvalStepId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/approvalId/steps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/approvalId/steps/approvalStepId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId{
				ApprovalId:     "approvalId",
				ApprovalStepId: "approvalStepId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/approvalId/steps/approvalStepId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepID(v.Input)
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

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/approvalId/steps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtApPrOvAlS/aPpRoVaLiD/sTePs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/approvalId/steps/approvalStepId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId{
				ApprovalId:     "approvalId",
				ApprovalStepId: "approvalStepId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/approvalId/steps/approvalStepId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtApPrOvAlS/aPpRoVaLiD/sTePs/aPpRoVaLsTePiD",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId{
				ApprovalId:     "aPpRoVaLiD",
				ApprovalStepId: "aPpRoVaLsTePiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtApPrOvAlS/aPpRoVaLiD/sTePs/aPpRoVaLsTePiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
