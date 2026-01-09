package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID("accessPackageCatalogId", "customAccessPackageWorkflowExtensionId")

	if id.AccessPackageCatalogId != "accessPackageCatalogId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageCatalogId'", id.AccessPackageCatalogId, "accessPackageCatalogId")
	}

	if id.CustomAccessPackageWorkflowExtensionId != "customAccessPackageWorkflowExtensionId" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomAccessPackageWorkflowExtensionId'", id.CustomAccessPackageWorkflowExtensionId, "customAccessPackageWorkflowExtensionId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID("accessPackageCatalogId", "customAccessPackageWorkflowExtensionId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/customAccessPackageWorkflowExtensions/customAccessPackageWorkflowExtensionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/customAccessPackageWorkflowExtensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/customAccessPackageWorkflowExtensions/customAccessPackageWorkflowExtensionId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId{
				AccessPackageCatalogId:                 "accessPackageCatalogId",
				CustomAccessPackageWorkflowExtensionId: "customAccessPackageWorkflowExtensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/customAccessPackageWorkflowExtensions/customAccessPackageWorkflowExtensionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID(v.Input)
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

		if actual.CustomAccessPackageWorkflowExtensionId != v.Expected.CustomAccessPackageWorkflowExtensionId {
			t.Fatalf("Expected %q but got %q for CustomAccessPackageWorkflowExtensionId", v.Expected.CustomAccessPackageWorkflowExtensionId, actual.CustomAccessPackageWorkflowExtensionId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/customAccessPackageWorkflowExtensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD/cUsToMaCcEsSpAcKaGeWoRkFlOwExTeNsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/customAccessPackageWorkflowExtensions/customAccessPackageWorkflowExtensionId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId{
				AccessPackageCatalogId:                 "accessPackageCatalogId",
				CustomAccessPackageWorkflowExtensionId: "customAccessPackageWorkflowExtensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogId/customAccessPackageWorkflowExtensions/customAccessPackageWorkflowExtensionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD/cUsToMaCcEsSpAcKaGeWoRkFlOwExTeNsIoNs/cUsToMaCcEsSpAcKaGeWoRkFlOwExTeNsIoNiD",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId{
				AccessPackageCatalogId:                 "aCcEsSpAcKaGeCaTaLoGiD",
				CustomAccessPackageWorkflowExtensionId: "cUsToMaCcEsSpAcKaGeWoRkFlOwExTeNsIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiD/cUsToMaCcEsSpAcKaGeWoRkFlOwExTeNsIoNs/cUsToMaCcEsSpAcKaGeWoRkFlOwExTeNsIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionIDInsensitively(v.Input)
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

		if actual.CustomAccessPackageWorkflowExtensionId != v.Expected.CustomAccessPackageWorkflowExtensionId {
			t.Fatalf("Expected %q but got %q for CustomAccessPackageWorkflowExtensionId", v.Expected.CustomAccessPackageWorkflowExtensionId, actual.CustomAccessPackageWorkflowExtensionId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
