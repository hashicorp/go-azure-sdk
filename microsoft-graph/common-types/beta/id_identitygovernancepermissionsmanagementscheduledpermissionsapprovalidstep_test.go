package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId{}

func TestNewIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID(t *testing.T) {
	id := NewIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID("approvalId", "approvalStepId")

	if id.ApprovalId != "approvalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalId'", id.ApprovalId, "approvalId")
	}

	if id.ApprovalStepId != "approvalStepId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalStepId'", id.ApprovalStepId, "approvalStepId")
	}
}

func TestFormatIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID(t *testing.T) {
	actual := NewIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID("approvalId", "approvalStepId").ID()
	expected := "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalId/steps/approvalStepId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId
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
			Input: "/identityGovernance/permissionsManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalId/steps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalId/steps/approvalStepId",
			Expected: &IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId{
				ApprovalId:     "approvalId",
				ApprovalStepId: "approvalStepId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalId/steps/approvalStepId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID(v.Input)
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

func TestParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId
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
			Input: "/identityGovernance/permissionsManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsMaNaGeMeNt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsMaNaGeMeNt/sChEdUlEdPeRmIsSiOnSaPpRoVaLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsMaNaGeMeNt/sChEdUlEdPeRmIsSiOnSaPpRoVaLs/aPpRoVaLiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalId/steps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsMaNaGeMeNt/sChEdUlEdPeRmIsSiOnSaPpRoVaLs/aPpRoVaLiD/sTePs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalId/steps/approvalStepId",
			Expected: &IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId{
				ApprovalId:     "approvalId",
				ApprovalStepId: "approvalStepId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalId/steps/approvalStepId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsMaNaGeMeNt/sChEdUlEdPeRmIsSiOnSaPpRoVaLs/aPpRoVaLiD/sTePs/aPpRoVaLsTePiD",
			Expected: &IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId{
				ApprovalId:     "aPpRoVaLiD",
				ApprovalStepId: "aPpRoVaLsTePiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsMaNaGeMeNt/sChEdUlEdPeRmIsSiOnSaPpRoVaLs/aPpRoVaLiD/sTePs/aPpRoVaLsTePiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId(t *testing.T) {
	segments := IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
