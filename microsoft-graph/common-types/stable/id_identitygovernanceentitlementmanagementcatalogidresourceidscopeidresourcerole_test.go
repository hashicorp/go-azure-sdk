package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId{}

func TestNewIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleID("accessPackageCatalogId", "accessPackageResourceId", "accessPackageResourceScopeId", "accessPackageResourceRoleId")

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

func TestFormatIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleID("accessPackageCatalogId", "accessPackageResourceId", "accessPackageResourceScopeId", "accessPackageResourceRoleId").ID()
	expected := "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources/accessPackageResourceId/scopes/accessPackageResourceScopeId/resource/roles/accessPackageResourceRoleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/catalogs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources/accessPackageResourceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources/accessPackageResourceId/scopes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources/accessPackageResourceId/scopes/accessPackageResourceScopeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources/accessPackageResourceId/scopes/accessPackageResourceScopeId/resource",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources/accessPackageResourceId/scopes/accessPackageResourceScopeId/resource/roles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources/accessPackageResourceId/scopes/accessPackageResourceScopeId/resource/roles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId{
				AccessPackageCatalogId:       "accessPackageCatalogId",
				AccessPackageResourceId:      "accessPackageResourceId",
				AccessPackageResourceScopeId: "accessPackageResourceScopeId",
				AccessPackageResourceRoleId:  "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources/accessPackageResourceId/scopes/accessPackageResourceScopeId/resource/roles/accessPackageResourceRoleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleID(v.Input)
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

func TestParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/catalogs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/rEsOuRcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources/accessPackageResourceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources/accessPackageResourceId/scopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId/sCoPeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources/accessPackageResourceId/scopes/accessPackageResourceScopeId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources/accessPackageResourceId/scopes/accessPackageResourceScopeId/resource",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiD/rEsOuRcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources/accessPackageResourceId/scopes/accessPackageResourceScopeId/resource/roles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiD/rEsOuRcE/rOlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources/accessPackageResourceId/scopes/accessPackageResourceScopeId/resource/roles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId{
				AccessPackageCatalogId:       "accessPackageCatalogId",
				AccessPackageResourceId:      "accessPackageResourceId",
				AccessPackageResourceScopeId: "accessPackageResourceScopeId",
				AccessPackageResourceRoleId:  "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resources/accessPackageResourceId/scopes/accessPackageResourceScopeId/resource/roles/accessPackageResourceRoleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiD/rEsOuRcE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId",
			Expected: &IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId{
				AccessPackageCatalogId:       "aCcEsSpAcKaGeCaTaLoGiD",
				AccessPackageResourceId:      "aCcEsSpAcKaGeReSoUrCeId",
				AccessPackageResourceScopeId: "aCcEsSpAcKaGeReSoUrCeScOpEiD",
				AccessPackageResourceRoleId:  "aCcEsSpAcKaGeReSoUrCeRoLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiD/rEsOuRcE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementCatalogIdResourceIdScopeIdResourceRoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
