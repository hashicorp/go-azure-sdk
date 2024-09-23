package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId{}

func TestNewIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleID("accessPackageResourceRoleScopeId", "accessPackageResourceRoleId")

	if id.AccessPackageResourceRoleScopeId != "accessPackageResourceRoleScopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleScopeId'", id.AccessPackageResourceRoleScopeId, "accessPackageResourceRoleScopeId")
	}

	if id.AccessPackageResourceRoleId != "accessPackageResourceRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleId'", id.AccessPackageResourceRoleId, "accessPackageResourceRoleId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleID("accessPackageResourceRoleScopeId", "accessPackageResourceRoleId").ID()
	expected := "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/scope/resource/roles/accessPackageResourceRoleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/scope",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/scope/resource",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/scope/resource/roles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/scope/resource/roles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId{
				AccessPackageResourceRoleScopeId: "accessPackageResourceRoleScopeId",
				AccessPackageResourceRoleId:      "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/scope/resource/roles/accessPackageResourceRoleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleID(v.Input)
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

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/scope",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/sCoPe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/scope/resource",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/sCoPe/rEsOuRcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/scope/resource/roles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/sCoPe/rEsOuRcE/rOlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/scope/resource/roles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId{
				AccessPackageResourceRoleScopeId: "accessPackageResourceRoleScopeId",
				AccessPackageResourceRoleId:      "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceRoleScopes/accessPackageResourceRoleScopeId/scope/resource/roles/accessPackageResourceRoleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/sCoPe/rEsOuRcE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId",
			Expected: &IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId{
				AccessPackageResourceRoleScopeId: "aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD",
				AccessPackageResourceRoleId:      "aCcEsSpAcKaGeReSoUrCeRoLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/sCoPe/rEsOuRcE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleIDInsensitively(v.Input)
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

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
