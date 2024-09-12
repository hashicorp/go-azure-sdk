package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID("accessPackageCatalogIdValue", "customAccessPackageWorkflowExtensionIdValue")

	if id.AccessPackageCatalogId != "accessPackageCatalogIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageCatalogId'", id.AccessPackageCatalogId, "accessPackageCatalogIdValue")
	}

	if id.CustomAccessPackageWorkflowExtensionId != "customAccessPackageWorkflowExtensionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomAccessPackageWorkflowExtensionId'", id.CustomAccessPackageWorkflowExtensionId, "customAccessPackageWorkflowExtensionIdValue")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionID("accessPackageCatalogIdValue", "customAccessPackageWorkflowExtensionIdValue").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogIdValue/customAccessPackageWorkflowExtensions/customAccessPackageWorkflowExtensionIdValue"
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
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogIdValue/customAccessPackageWorkflowExtensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogIdValue/customAccessPackageWorkflowExtensions/customAccessPackageWorkflowExtensionIdValue",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId{
				AccessPackageCatalogId:                 "accessPackageCatalogIdValue",
				CustomAccessPackageWorkflowExtensionId: "customAccessPackageWorkflowExtensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogIdValue/customAccessPackageWorkflowExtensions/customAccessPackageWorkflowExtensionIdValue/extra",
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
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogIdValue/customAccessPackageWorkflowExtensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiDvAlUe/cUsToMaCcEsSpAcKaGeWoRkFlOwExTeNsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogIdValue/customAccessPackageWorkflowExtensions/customAccessPackageWorkflowExtensionIdValue",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId{
				AccessPackageCatalogId:                 "accessPackageCatalogIdValue",
				CustomAccessPackageWorkflowExtensionId: "customAccessPackageWorkflowExtensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageCatalogs/accessPackageCatalogIdValue/customAccessPackageWorkflowExtensions/customAccessPackageWorkflowExtensionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiDvAlUe/cUsToMaCcEsSpAcKaGeWoRkFlOwExTeNsIoNs/cUsToMaCcEsSpAcKaGeWoRkFlOwExTeNsIoNiDvAlUe",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId{
				AccessPackageCatalogId:                 "aCcEsSpAcKaGeCaTaLoGiDvAlUe",
				CustomAccessPackageWorkflowExtensionId: "cUsToMaCcEsSpAcKaGeWoRkFlOwExTeNsIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeCaTaLoGs/aCcEsSpAcKaGeCaTaLoGiDvAlUe/cUsToMaCcEsSpAcKaGeWoRkFlOwExTeNsIoNs/cUsToMaCcEsSpAcKaGeWoRkFlOwExTeNsIoNiDvAlUe/extra",
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
