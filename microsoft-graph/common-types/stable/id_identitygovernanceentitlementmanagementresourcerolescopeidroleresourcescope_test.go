package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId{}

func TestNewIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeID("accessPackageResourceRoleScopeId", "accessPackageResourceScopeId")

	if id.AccessPackageResourceRoleScopeId != "accessPackageResourceRoleScopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleScopeId'", id.AccessPackageResourceRoleScopeId, "accessPackageResourceRoleScopeId")
	}

	if id.AccessPackageResourceScopeId != "accessPackageResourceScopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceScopeId'", id.AccessPackageResourceScopeId, "accessPackageResourceScopeId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeID("accessPackageResourceRoleScopeId", "accessPackageResourceScopeId").ID()
	expected := "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource/scopes/accessPackageResourceScopeId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId
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
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/role",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource/scopes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource/scopes/accessPackageResourceScopeId",
			Expected: &IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId{
				AccessPackageResourceRoleScopeId: "accessPackageResourceRoleScopeId",
				AccessPackageResourceScopeId:     "accessPackageResourceScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource/scopes/accessPackageResourceScopeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageResourceRoleScopeId != v.Expected.AccessPackageResourceRoleScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleScopeId", v.Expected.AccessPackageResourceRoleScopeId, actual.AccessPackageResourceRoleScopeId)
		}

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId
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
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErOlEsCoPeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/role",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/rOlE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/rOlE/rEsOuRcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource/scopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/rOlE/rEsOuRcE/sCoPeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource/scopes/accessPackageResourceScopeId",
			Expected: &IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId{
				AccessPackageResourceRoleScopeId: "accessPackageResourceRoleScopeId",
				AccessPackageResourceScopeId:     "accessPackageResourceScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource/scopes/accessPackageResourceScopeId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/rOlE/rEsOuRcE/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiD",
			Expected: &IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId{
				AccessPackageResourceRoleScopeId: "aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD",
				AccessPackageResourceScopeId:     "aCcEsSpAcKaGeReSoUrCeScOpEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/rOlE/rEsOuRcE/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageResourceRoleScopeId != v.Expected.AccessPackageResourceRoleScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleScopeId", v.Expected.AccessPackageResourceRoleScopeId, actual.AccessPackageResourceRoleScopeId)
		}

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
