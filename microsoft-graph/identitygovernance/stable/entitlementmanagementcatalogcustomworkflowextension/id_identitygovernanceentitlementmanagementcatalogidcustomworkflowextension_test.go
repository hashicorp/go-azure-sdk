package entitlementmanagementcatalogcustomworkflowextension

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId{}

func TestNewIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID("accessPackageCatalogIdValue", "customCalloutExtensionIdValue")

	if id.AccessPackageCatalogId != "accessPackageCatalogIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageCatalogId'", id.AccessPackageCatalogId, "accessPackageCatalogIdValue")
	}

	if id.CustomCalloutExtensionId != "customCalloutExtensionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomCalloutExtensionId'", id.CustomCalloutExtensionId, "customCalloutExtensionIdValue")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionID("accessPackageCatalogIdValue", "customCalloutExtensionIdValue").ID()
	expected := "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/customWorkflowExtensions/customCalloutExtensionIdValue"
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
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/customWorkflowExtensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/customWorkflowExtensions/customCalloutExtensionIdValue",
			Expected: &IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId{
				AccessPackageCatalogId:   "accessPackageCatalogIdValue",
				CustomCalloutExtensionId: "customCalloutExtensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/customWorkflowExtensions/customCalloutExtensionIdValue/extra",
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
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/customWorkflowExtensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiDvAlUe/cUsToMwOrKfLoWeXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/customWorkflowExtensions/customCalloutExtensionIdValue",
			Expected: &IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId{
				AccessPackageCatalogId:   "accessPackageCatalogIdValue",
				CustomCalloutExtensionId: "customCalloutExtensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/catalogs/accessPackageCatalogIdValue/customWorkflowExtensions/customCalloutExtensionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiDvAlUe/cUsToMwOrKfLoWeXtEnSiOnS/cUsToMcAlLoUtExTeNsIoNiDvAlUe",
			Expected: &IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId{
				AccessPackageCatalogId:   "aCcEsSpAcKaGeCaTaLoGiDvAlUe",
				CustomCalloutExtensionId: "cUsToMcAlLoUtExTeNsIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/cAtAlOgS/aCcEsSpAcKaGeCaTaLoGiDvAlUe/cUsToMwOrKfLoWeXtEnSiOnS/cUsToMcAlLoUtExTeNsIoNiDvAlUe/extra",
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
