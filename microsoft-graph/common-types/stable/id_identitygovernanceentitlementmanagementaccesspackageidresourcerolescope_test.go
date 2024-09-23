package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeID("accessPackageId", "accessPackageResourceRoleScopeId")

	if id.AccessPackageId != "accessPackageId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageId'", id.AccessPackageId, "accessPackageId")
	}

	if id.AccessPackageResourceRoleScopeId != "accessPackageResourceRoleScopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleScopeId'", id.AccessPackageResourceRoleScopeId, "accessPackageResourceRoleScopeId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeID("accessPackageId", "accessPackageResourceRoleScopeId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId
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
			Input: "/identityGovernance/entitlementManagement/accessPackages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId{
				AccessPackageId:                  "accessPackageId",
				AccessPackageResourceRoleScopeId: "accessPackageResourceRoleScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageId != v.Expected.AccessPackageId {
			t.Fatalf("Expected %q but got %q for AccessPackageId", v.Expected.AccessPackageId, actual.AccessPackageId)
		}

		if actual.AccessPackageResourceRoleScopeId != v.Expected.AccessPackageResourceRoleScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleScopeId", v.Expected.AccessPackageResourceRoleScopeId, actual.AccessPackageResourceRoleScopeId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId
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
			Input: "/identityGovernance/entitlementManagement/accessPackages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/rEsOuRcErOlEsCoPeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId{
				AccessPackageId:                  "accessPackageId",
				AccessPackageResourceRoleScopeId: "accessPackageResourceRoleScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId{
				AccessPackageId:                  "aCcEsSpAcKaGeId",
				AccessPackageResourceRoleScopeId: "aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageId != v.Expected.AccessPackageId {
			t.Fatalf("Expected %q but got %q for AccessPackageId", v.Expected.AccessPackageId, actual.AccessPackageId)
		}

		if actual.AccessPackageResourceRoleScopeId != v.Expected.AccessPackageResourceRoleScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleScopeId", v.Expected.AccessPackageResourceRoleScopeId, actual.AccessPackageResourceRoleScopeId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
