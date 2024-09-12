package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId{}

func TestNewIdentityGovernancePermissionsManagementScheduledPermissionsApprovalID(t *testing.T) {
	id := NewIdentityGovernancePermissionsManagementScheduledPermissionsApprovalID("approvalIdValue")

	if id.ApprovalId != "approvalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalId'", id.ApprovalId, "approvalIdValue")
	}
}

func TestFormatIdentityGovernancePermissionsManagementScheduledPermissionsApprovalID(t *testing.T) {
	actual := NewIdentityGovernancePermissionsManagementScheduledPermissionsApprovalID("approvalIdValue").ID()
	expected := "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId
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
			// Valid URI
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalIdValue",
			Expected: &IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId{
				ApprovalId: "approvalIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalID(v.Input)
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

	}
}

func TestParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId
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
			// Valid URI
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalIdValue",
			Expected: &IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId{
				ApprovalId: "approvalIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/permissionsManagement/scheduledPermissionsApprovals/approvalIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsMaNaGeMeNt/sChEdUlEdPeRmIsSiOnSaPpRoVaLs/aPpRoVaLiDvAlUe",
			Expected: &IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId{
				ApprovalId: "aPpRoVaLiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/pErMiSsIoNsMaNaGeMeNt/sChEdUlEdPeRmIsSiOnSaPpRoVaLs/aPpRoVaLiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernancePermissionsManagementScheduledPermissionsApprovalIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForIdentityGovernancePermissionsManagementScheduledPermissionsApprovalId(t *testing.T) {
	segments := IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernancePermissionsManagementScheduledPermissionsApprovalId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
