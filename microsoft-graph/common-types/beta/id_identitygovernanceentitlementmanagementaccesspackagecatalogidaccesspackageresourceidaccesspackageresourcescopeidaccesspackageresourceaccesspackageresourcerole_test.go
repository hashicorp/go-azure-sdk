package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID("accessPackageCatalogId", "accessPackageResourceId", "accessPackageResourceScopeId", "accessPackageResourceRoleId")

	if id.AccessPackageCatalogId != "accessPackageCatalogId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageCatalogId'", id.AccessPackageCatalogId, "accessPackageCatalogId")
	}

	if id.AccessPackageResourceId != "accessPackageResourceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceId'", id.AccessPackageResourceId, "accessPackageResourceId")
	}

	if id.AccessPackageResourceScopeId != "accessPackageResourceScopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceScopeId'", id.AccessPackageResourceScopeId, "accessPackageResourceScopeId")
	}

	if id.AccessPackageResourceRoleId != "accessPackageResourceRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleId'", id.AccessPackageResourceRoleId, "accessPackageResourceRoleId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID("accessPackageCatalogId", "accessPackageResourceId", "accessPackageResourceScopeId", "accessPackageResourceRoleId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource/accessPackageResourceRoles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{
				AccessPackageCatalogId:       "accessPackageCatalogId",
				AccessPackageResourceId:      "accessPackageResourceId",
				AccessPackageResourceScopeId: "accessPackageResourceScopeId",
				AccessPackageResourceRoleId:  "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageCatalogId != v.Expected.AccessPackageCatalogId {
			t.Fatalf("Expected %q but got %q for AccessPackageCatalogId", v.Expected.AccessPackageCatalogId, actual.AccessPackageCatalogId)
		}

		if actual.AccessPackageResourceId != v.Expected.AccessPackageResourceId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceId", v.Expected.AccessPackageResourceId, actual.AccessPackageResourceId)
		}

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD/aCcEsSpAcKaGeReSoUrCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeScOpEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeScOpEs/aCcEsSpAcKaGeReSoUrCeScOpEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeScOpEs/aCcEsSpAcKaGeReSoUrCeScOpEiD/aCcEsSpAcKaGeReSoUrCe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource/accessPackageResourceRoles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeScOpEs/aCcEsSpAcKaGeReSoUrCeScOpEiD/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeRoLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{
				AccessPackageCatalogId:       "accessPackageCatalogId",
				AccessPackageResourceId:      "accessPackageResourceId",
				AccessPackageResourceScopeId: "accessPackageResourceScopeId",
				AccessPackageResourceRoleId:  "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeScOpEs/aCcEsSpAcKaGeReSoUrCeScOpEiD/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeRoLeS/aCcEsSpAcKaGeReSoUrCeRoLeId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{
				AccessPackageCatalogId:       "aCcEsSpAcKaGeCaTaLoGiD",
				AccessPackageResourceId:      "aCcEsSpAcKaGeReSoUrCeId",
				AccessPackageResourceScopeId: "aCcEsSpAcKaGeReSoUrCeScOpEiD",
				AccessPackageResourceRoleId:  "aCcEsSpAcKaGeReSoUrCeRoLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeScOpEs/aCcEsSpAcKaGeReSoUrCeScOpEiD/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeRoLeS/aCcEsSpAcKaGeReSoUrCeRoLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageCatalogId != v.Expected.AccessPackageCatalogId {
			t.Fatalf("Expected %q but got %q for AccessPackageCatalogId", v.Expected.AccessPackageCatalogId, actual.AccessPackageCatalogId)
		}

		if actual.AccessPackageResourceId != v.Expected.AccessPackageResourceId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceId", v.Expected.AccessPackageResourceId, actual.AccessPackageResourceId)
		}

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
