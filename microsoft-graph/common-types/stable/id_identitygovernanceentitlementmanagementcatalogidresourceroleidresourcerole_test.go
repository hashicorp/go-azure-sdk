package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId{}

func TestNewIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleID("accessPackageCatalogId", "accessPackageResourceRoleId", "accessPackageResourceRoleId1")

	if id.AccessPackageCatalogId != "accessPackageCatalogId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageCatalogId'", id.AccessPackageCatalogId, "accessPackageCatalogId")
	}

	if id.AccessPackageResourceRoleId != "accessPackageResourceRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleId'", id.AccessPackageResourceRoleId, "accessPackageResourceRoleId")
	}

	if id.AccessPackageResourceRoleId1 != "accessPackageResourceRoleId1" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleId1'", id.AccessPackageResourceRoleId1, "accessPackageResourceRoleId1")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleID("accessPackageCatalogId", "accessPackageResourceRoleId", "accessPackageResourceRoleId1").ID()
	expected := "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resourceRoles/accessPackageResourceRoleId/resource/roles/accessPackageResourceRoleId1"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resourceRoles",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resourceRoles/accessPackageResourceRoleId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resourceRoles/accessPackageResourceRoleId/resource",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resourceRoles/accessPackageResourceRoleId/resource/roles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resourceRoles/accessPackageResourceRoleId/resource/roles/accessPackageResourceRoleId1",
			Expected: &IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId{
				AccessPackageCatalogId:       "accessPackageCatalogId",
				AccessPackageResourceRoleId:  "accessPackageResourceRoleId",
				AccessPackageResourceRoleId1: "accessPackageResourceRoleId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resourceRoles/accessPackageResourceRoleId/resource/roles/accessPackageResourceRoleId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleID(v.Input)
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

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

		if actual.AccessPackageResourceRoleId1 != v.Expected.AccessPackageResourceRoleId1 {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId1", v.Expected.AccessPackageResourceRoleId1, actual.AccessPackageResourceRoleId1)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resourceRoles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/rEsOuRcErOlEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resourceRoles/accessPackageResourceRoleId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/rEsOuRcErOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resourceRoles/accessPackageResourceRoleId/resource",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/rEsOuRcErOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId/rEsOuRcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resourceRoles/accessPackageResourceRoleId/resource/roles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/rEsOuRcErOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId/rEsOuRcE/rOlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resourceRoles/accessPackageResourceRoleId/resource/roles/accessPackageResourceRoleId1",
			Expected: &IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId{
				AccessPackageCatalogId:       "accessPackageCatalogId",
				AccessPackageResourceRoleId:  "accessPackageResourceRoleId",
				AccessPackageResourceRoleId1: "accessPackageResourceRoleId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/resourceRoles/accessPackageResourceRoleId/resource/roles/accessPackageResourceRoleId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/rEsOuRcErOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId/rEsOuRcE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId1",
			Expected: &IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId{
				AccessPackageCatalogId:       "aCcEsSpAcKaGeCaTaLoGiD",
				AccessPackageResourceRoleId:  "aCcEsSpAcKaGeReSoUrCeRoLeId",
				AccessPackageResourceRoleId1: "aCcEsSpAcKaGeReSoUrCeRoLeId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/rEsOuRcErOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId/rEsOuRcE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleIDInsensitively(v.Input)
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

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

		if actual.AccessPackageResourceRoleId1 != v.Expected.AccessPackageResourceRoleId1 {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId1", v.Expected.AccessPackageResourceRoleId1, actual.AccessPackageResourceRoleId1)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceRoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
