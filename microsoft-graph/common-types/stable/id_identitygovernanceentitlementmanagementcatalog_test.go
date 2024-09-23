package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogId{}

func TestNewIdentityGovernanceEntitlementManagementCatalogID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementCatalogID("accessPackageCatalogId")

	if id.AccessPackageCatalogId != "accessPackageCatalogId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageCatalogId'", id.AccessPackageCatalogId, "accessPackageCatalogId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementCatalogID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementCatalogID("accessPackageCatalogId").ID()
	expected := "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementCatalogID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementCatalogId
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
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId",
			Expected: &IdentityGovernanceEntitlementManagementCatalogId{
				AccessPackageCatalogId: "accessPackageCatalogId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementCatalogID(v.Input)
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

	}
}

func TestParseIdentityGovernanceEntitlementManagementCatalogIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementCatalogId
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
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId",
			Expected: &IdentityGovernanceEntitlementManagementCatalogId{
				AccessPackageCatalogId: "accessPackageCatalogId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD",
			Expected: &IdentityGovernanceEntitlementManagementCatalogId{
				AccessPackageCatalogId: "aCcEsSpAcKaGeCaTaLoGiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementCatalogIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementCatalogId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementCatalogId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementCatalogId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
