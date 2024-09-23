package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId{}

func TestNewIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepID(t *testing.T) {
	id := NewIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepID("approvalId", "approvalStepId")

	if id.ApprovalId != "approvalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalId'", id.ApprovalId, "approvalId")
	}

	if id.ApprovalStepId != "approvalStepId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalStepId'", id.ApprovalStepId, "approvalStepId")
	}
}

func TestFormatIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepID(t *testing.T) {
	actual := NewIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepID("approvalId", "approvalStepId").ID()
	expected := "/identityGovernance/privilegedAccess/group/assignmentApprovals/approvalId/steps/approvalStepId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId
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
			Input: "/identityGovernance/privilegedAccess",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/privilegedAccess/group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/privilegedAccess/group/assignmentApprovals",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/privilegedAccess/group/assignmentApprovals/approvalId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/privilegedAccess/group/assignmentApprovals/approvalId/steps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/privilegedAccess/group/assignmentApprovals/approvalId/steps/approvalStepId",
			Expected: &IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId{
				ApprovalId:     "approvalId",
				ApprovalStepId: "approvalStepId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/assignmentApprovals/approvalId/steps/approvalStepId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepID(v.Input)
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

func TestParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId
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
			Input: "/identityGovernance/privilegedAccess",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/privilegedAccess/group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/privilegedAccess/group/assignmentApprovals",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/aSsIgNmEnTaPpRoVaLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/privilegedAccess/group/assignmentApprovals/approvalId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/aSsIgNmEnTaPpRoVaLs/aPpRoVaLiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/privilegedAccess/group/assignmentApprovals/approvalId/steps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/aSsIgNmEnTaPpRoVaLs/aPpRoVaLiD/sTePs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/privilegedAccess/group/assignmentApprovals/approvalId/steps/approvalStepId",
			Expected: &IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId{
				ApprovalId:     "approvalId",
				ApprovalStepId: "approvalStepId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/privilegedAccess/group/assignmentApprovals/approvalId/steps/approvalStepId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/aSsIgNmEnTaPpRoVaLs/aPpRoVaLiD/sTePs/aPpRoVaLsTePiD",
			Expected: &IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId{
				ApprovalId:     "aPpRoVaLiD",
				ApprovalStepId: "aPpRoVaLsTePiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pRiViLeGeDaCcEsS/gRoUp/aSsIgNmEnTaPpRoVaLs/aPpRoVaLiD/sTePs/aPpRoVaLsTePiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId(t *testing.T) {
	segments := IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStepId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
