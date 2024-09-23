package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId{}

func TestNewIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID("accessPackageCatalogId", "customCalloutExtensionId")

	if id.AccessPackageCatalogId != "accessPackageCatalogId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageCatalogId'", id.AccessPackageCatalogId, "accessPackageCatalogId")
	}

	if id.CustomCalloutExtensionId != "customCalloutExtensionId" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomCalloutExtensionId'", id.CustomCalloutExtensionId, "customCalloutExtensionId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID("accessPackageCatalogId", "customCalloutExtensionId").ID()
	expected := "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/customWorkflowExtensions/customCalloutExtensionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId
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
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/customWorkflowExtensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/customWorkflowExtensions/customCalloutExtensionId",
			Expected: &IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId{
				AccessPackageCatalogId:   "accessPackageCatalogId",
				CustomCalloutExtensionId: "customCalloutExtensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/customWorkflowExtensions/customCalloutExtensionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID(v.Input)
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

		if actual.CustomCalloutExtensionId != v.Expected.CustomCalloutExtensionId {
			t.Fatalf("Expected %q but got %q for CustomCalloutExtensionId", v.Expected.CustomCalloutExtensionId, actual.CustomCalloutExtensionId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId
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
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/customWorkflowExtensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/cUsToMwOrKfLoWeXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/customWorkflowExtensions/customCalloutExtensionId",
			Expected: &IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId{
				AccessPackageCatalogId:   "accessPackageCatalogId",
				CustomCalloutExtensionId: "customCalloutExtensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogId/customWorkflowExtensions/customCalloutExtensionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/cUsToMwOrKfLoWeXtEnSiOnS/cUsToMcAlLoUtExTeNsIoNiD",
			Expected: &IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId{
				AccessPackageCatalogId:   "aCcEsSpAcKaGeCaTaLoGiD",
				CustomCalloutExtensionId: "cUsToMcAlLoUtExTeNsIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiD/cUsToMwOrKfLoWeXtEnSiOnS/cUsToMcAlLoUtExTeNsIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionIDInsensitively(v.Input)
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

		if actual.CustomCalloutExtensionId != v.Expected.CustomCalloutExtensionId {
			t.Fatalf("Expected %q but got %q for CustomCalloutExtensionId", v.Expected.CustomCalloutExtensionId, actual.CustomCalloutExtensionId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
