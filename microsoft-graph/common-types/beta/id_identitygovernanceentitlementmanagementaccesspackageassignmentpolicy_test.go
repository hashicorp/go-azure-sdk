package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID("accessPackageAssignmentPolicyId")

	if id.AccessPackageAssignmentPolicyId != "accessPackageAssignmentPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentPolicyId'", id.AccessPackageAssignmentPolicyId, "accessPackageAssignmentPolicyId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID("accessPackageAssignmentPolicyId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId{
				AccessPackageAssignmentPolicyId: "accessPackageAssignmentPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyID(v.Input)
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

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtPoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId{
				AccessPackageAssignmentPolicyId: "accessPackageAssignmentPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentPolicies/accessPackageAssignmentPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtPoLiCiEs/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId{
				AccessPackageAssignmentPolicyId: "aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtPoLiCiEs/aCcEsSpAcKaGeAsSiGnMeNtPoLiCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageAssignmentPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
