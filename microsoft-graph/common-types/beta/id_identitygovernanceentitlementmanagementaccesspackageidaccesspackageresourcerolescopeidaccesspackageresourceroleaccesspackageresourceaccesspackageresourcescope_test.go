package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeID("accessPackageId", "accessPackageResourceRoleScopeId", "accessPackageResourceScopeId")

	if id.AccessPackageId != "accessPackageId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageId'", id.AccessPackageId, "accessPackageId")
	}

	if id.AccessPackageResourceRoleScopeId != "accessPackageResourceRoleScopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleScopeId'", id.AccessPackageResourceRoleScopeId, "accessPackageResourceRoleScopeId")
	}

	if id.AccessPackageResourceScopeId != "accessPackageResourceScopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceScopeId'", id.AccessPackageResourceScopeId, "accessPackageResourceScopeId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeID("accessPackageId", "accessPackageResourceRoleScopeId", "accessPackageResourceScopeId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource/accessPackageResourceScopes/accessPackageResourceScopeId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId
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
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageResourceRoleScopes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource/accessPackageResourceScopes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource/accessPackageResourceScopes/accessPackageResourceScopeId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId{
				AccessPackageId:                  "accessPackageId",
				AccessPackageResourceRoleScopeId: "accessPackageResourceRoleScopeId",
				AccessPackageResourceScopeId:     "accessPackageResourceScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource/accessPackageResourceScopes/accessPackageResourceScopeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeID(v.Input)
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

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId
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
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageResourceRoleScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/aCcEsSpAcKaGeReSoUrCeRoLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/aCcEsSpAcKaGeReSoUrCeRoLe/aCcEsSpAcKaGeReSoUrCe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource/accessPackageResourceScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/aCcEsSpAcKaGeReSoUrCeRoLe/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeScOpEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource/accessPackageResourceScopes/accessPackageResourceScopeId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId{
				AccessPackageId:                  "accessPackageId",
				AccessPackageResourceRoleScopeId: "accessPackageResourceRoleScopeId",
				AccessPackageResourceScopeId:     "accessPackageResourceScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource/accessPackageResourceScopes/accessPackageResourceScopeId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/aCcEsSpAcKaGeReSoUrCeRoLe/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeScOpEs/aCcEsSpAcKaGeReSoUrCeScOpEiD",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId{
				AccessPackageId:                  "aCcEsSpAcKaGeId",
				AccessPackageResourceRoleScopeId: "aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD",
				AccessPackageResourceScopeId:     "aCcEsSpAcKaGeReSoUrCeScOpEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/aCcEsSpAcKaGeReSoUrCeRoLe/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeScOpEs/aCcEsSpAcKaGeReSoUrCeScOpEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeIDInsensitively(v.Input)
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

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
