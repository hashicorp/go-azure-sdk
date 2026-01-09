package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionID("accessPackageCatalogId", "customCalloutExtensionId")

	if id.AccessPackageCatalogId != "accessPackageCatalogId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageCatalogId'", id.AccessPackageCatalogId, "accessPackageCatalogId")
	}

	if id.CustomCalloutExtensionId != "customCalloutExtensionId" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomCalloutExtensionId'", id.CustomCalloutExtensionId, "customCalloutExtensionId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionID("accessPackageCatalogId", "customCalloutExtensionId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageCustomWorkflowExtensions/customCalloutExtensionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageCustomWorkflowExtensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageCustomWorkflowExtensions/customCalloutExtensionId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId{
				AccessPackageCatalogId:   "accessPackageCatalogId",
				CustomCalloutExtensionId: "customCalloutExtensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageCustomWorkflowExtensions/customCalloutExtensionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionID(v.Input)
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

func TestParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageCustomWorkflowExtensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD/aCcEsSpAcKaGeCuStOmWoRkFlOwExTeNsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageCustomWorkflowExtensions/customCalloutExtensionId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId{
				AccessPackageCatalogId:   "accessPackageCatalogId",
				CustomCalloutExtensionId: "customCalloutExtensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/accessPackageCustomWorkflowExtensions/customCalloutExtensionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD/aCcEsSpAcKaGeCuStOmWoRkFlOwExTeNsIoNs/cUsToMcAlLoUtExTeNsIoNiD",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId{
				AccessPackageCatalogId:   "aCcEsSpAcKaGeCaTaLoGiD",
				CustomCalloutExtensionId: "cUsToMcAlLoUtExTeNsIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD/aCcEsSpAcKaGeCuStOmWoRkFlOwExTeNsIoNs/cUsToMcAlLoUtExTeNsIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageCustomWorkflowExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
