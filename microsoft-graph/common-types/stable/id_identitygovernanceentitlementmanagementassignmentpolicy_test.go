package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAssignmentPolicyId{}

func TestNewIdentityGovernanceEntitlementManagementAssignmentPolicyID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAssignmentPolicyID("accessPackageAssignmentPolicyId")

	if id.AccessPackageAssignmentPolicyId != "accessPackageAssignmentPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentPolicyId'", id.AccessPackageAssignmentPolicyId, "accessPackageAssignmentPolicyId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAssignmentPolicyID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAssignmentPolicyID("accessPackageAssignmentPolicyId").ID()
	expected := "/identityGovernance/entitlementManagement/assignmentPolicies/accessPackageAssignmentPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAssignmentPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAssignmentPolicyId
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
			Input: "/identityGovernance/entitlementManagement/assignmentPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/assignmentPolicies/accessPackageAssignmentPolicyId",
			Expected: &IdentityGovernanceEntitlementManagementAssignmentPolicyId{
				AccessPackageAssignmentPolicyId: "accessPackageAssignmentPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/assignmentPolicies/accessPackageAssignmentPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAssignmentPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageAssignmentPolicyId != v.Expected.AccessPackageAssignmentPolicyId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentPolicyId", v.Expected.AccessPackageAssignmentPolicyId, actual.AccessPackageAssignmentPolicyId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAssignmentPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAssignmentPolicyId
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
			Input: "/identityGovernance/entitlementManagement/assignmentPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aSsIgNmEnTpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/assignmentPolicies/accessPackageAssignmentPolicyId",
			Expected: &IdentityGovernanceEntitlementManagementAssignmentPolicyId{
				AccessPackageAssignmentPolicyId: "accessPackageAssignmentPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/assignmentPolicies/accessPackageAssignmentPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aSsIgNmEnTpOlIcIeS/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId",
			Expected: &IdentityGovernanceEntitlementManagementAssignmentPolicyId{
				AccessPackageAssignmentPolicyId: "aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aSsIgNmEnTpOlIcIeS/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAssignmentPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageAssignmentPolicyId != v.Expected.AccessPackageAssignmentPolicyId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentPolicyId", v.Expected.AccessPackageAssignmentPolicyId, actual.AccessPackageAssignmentPolicyId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAssignmentPolicyId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAssignmentPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAssignmentPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
