package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId{}

func TestNewIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID(t *testing.T) {
	id := NewIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID("approvalIdValue", "approvalStepIdValue")

	if id.ApprovalId != "approvalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalId'", id.ApprovalId, "approvalIdValue")
	}

	if id.ApprovalStepId != "approvalStepIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalStepId'", id.ApprovalStepId, "approvalStepIdValue")
	}
}

func TestFormatIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID(t *testing.T) {
	actual := NewIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepID("approvalIdValue", "approvalStepIdValue").ID()
	expected := "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalIdValue/steps/approvalStepIdValue"
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
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalIdValue/steps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalIdValue/steps/approvalStepIdValue",
			Expected: &IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId{
				ApprovalId:     "approvalIdValue",
				ApprovalStepId: "approvalStepIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalIdValue/steps/approvalStepIdValue/extra",
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
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsMaNaGeMeNt/sChEdUlEdPeRmIsSiOnSaPpRoVaLs/aPpRoVaLiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalIdValue/steps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsMaNaGeMeNt/sChEdUlEdPeRmIsSiOnSaPpRoVaLs/aPpRoVaLiDvAlUe/sTePs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalIdValue/steps/approvalStepIdValue",
			Expected: &IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId{
				ApprovalId:     "approvalIdValue",
				ApprovalStepId: "approvalStepIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalIdValue/steps/approvalStepIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsMaNaGeMeNt/sChEdUlEdPeRmIsSiOnSaPpRoVaLs/aPpRoVaLiDvAlUe/sTePs/aPpRoVaLsTePiDvAlUe",
			Expected: &IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId{
				ApprovalId:     "aPpRoVaLiDvAlUe",
				ApprovalStepId: "aPpRoVaLsTePiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsMaNaGeMeNt/sChEdUlEdPeRmIsSiOnSaPpRoVaLs/aPpRoVaLiDvAlUe/sTePs/aPpRoVaLsTePiDvAlUe/extra",
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
