package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceID("accessPackageCatalogId", "accessPackageResourceId")

	if id.AccessPackageCatalogId != "accessPackageCatalogId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageCatalogId'", id.AccessPackageCatalogId, "accessPackageCatalogId")
	}

	if id.AccessPackageResourceId != "accessPackageResourceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceId'", id.AccessPackageResourceId, "accessPackageResourceId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceID("accessPackageCatalogId", "accessPackageResourceId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceId
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
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceId{
				AccessPackageCatalogId:  "accessPackageCatalogId",
				AccessPackageResourceId: "accessPackageResourceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceID(v.Input)
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

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceId
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
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceId{
				AccessPackageCatalogId:  "accessPackageCatalogId",
				AccessPackageResourceId: "accessPackageResourceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageResources/accessPackageResourceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceId{
				AccessPackageCatalogId:  "aCcEsSpAcKaGeCaTaLoGiD",
				AccessPackageResourceId: "aCcEsSpAcKaGeReSoUrCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
