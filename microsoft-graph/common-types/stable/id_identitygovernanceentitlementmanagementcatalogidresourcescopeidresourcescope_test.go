package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId{}

func TestNewIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeID("accessPackageCatalogIdValue", "accessPackageResourceScopeIdValue", "accessPackageResourceScopeId1Value")

	if id.AccessPackageCatalogId != "accessPackageCatalogIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageCatalogId'", id.AccessPackageCatalogId, "accessPackageCatalogIdValue")
	}

	if id.AccessPackageResourceScopeId != "accessPackageResourceScopeIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceScopeId'", id.AccessPackageResourceScopeId, "accessPackageResourceScopeIdValue")
	}

	if id.AccessPackageResourceScopeId1 != "accessPackageResourceScopeId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceScopeId1'", id.AccessPackageResourceScopeId1, "accessPackageResourceScopeId1Value")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeID("accessPackageCatalogIdValue", "accessPackageResourceScopeIdValue", "accessPackageResourceScopeId1Value").ID()
	expected := "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/resourceScopes/accessPackageResourceScopeIdValue/resource/scopes/accessPackageResourceScopeId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId
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
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/resourceScopes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/resourceScopes/accessPackageResourceScopeIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/resourceScopes/accessPackageResourceScopeIdValue/resource",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/resourceScopes/accessPackageResourceScopeIdValue/resource/scopes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/resourceScopes/accessPackageResourceScopeIdValue/resource/scopes/accessPackageResourceScopeId1Value",
			Expected: &IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId{
				AccessPackageCatalogId:        "accessPackageCatalogIdValue",
				AccessPackageResourceScopeId:  "accessPackageResourceScopeIdValue",
				AccessPackageResourceScopeId1: "accessPackageResourceScopeId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/resourceScopes/accessPackageResourceScopeIdValue/resource/scopes/accessPackageResourceScopeId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeID(v.Input)
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

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

		if actual.AccessPackageResourceScopeId1 != v.Expected.AccessPackageResourceScopeId1 {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId1", v.Expected.AccessPackageResourceScopeId1, actual.AccessPackageResourceScopeId1)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId
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
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/resourceScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiDvAlUe/rEsOuRcEsCoPeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/resourceScopes/accessPackageResourceScopeIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiDvAlUe/rEsOuRcEsCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/resourceScopes/accessPackageResourceScopeIdValue/resource",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiDvAlUe/rEsOuRcEsCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe/rEsOuRcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/resourceScopes/accessPackageResourceScopeIdValue/resource/scopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiDvAlUe/rEsOuRcEsCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe/rEsOuRcE/sCoPeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/resourceScopes/accessPackageResourceScopeIdValue/resource/scopes/accessPackageResourceScopeId1Value",
			Expected: &IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId{
				AccessPackageCatalogId:        "accessPackageCatalogIdValue",
				AccessPackageResourceScopeId:  "accessPackageResourceScopeIdValue",
				AccessPackageResourceScopeId1: "accessPackageResourceScopeId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/resourceScopes/accessPackageResourceScopeIdValue/resource/scopes/accessPackageResourceScopeId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiDvAlUe/rEsOuRcEsCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe/rEsOuRcE/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiD1VaLuE",
			Expected: &IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId{
				AccessPackageCatalogId:        "aCcEsSpAcKaGeCaTaLoGiDvAlUe",
				AccessPackageResourceScopeId:  "aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe",
				AccessPackageResourceScopeId1: "aCcEsSpAcKaGeReSoUrCeScOpEiD1VaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiDvAlUe/rEsOuRcEsCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe/rEsOuRcE/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiD1VaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeIDInsensitively(v.Input)
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

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

		if actual.AccessPackageResourceScopeId1 != v.Expected.AccessPackageResourceScopeId1 {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId1", v.Expected.AccessPackageResourceScopeId1, actual.AccessPackageResourceScopeId1)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceScopeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
