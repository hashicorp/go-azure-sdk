package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId{}

func TestNewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionID("accessPackageResourceRequestIdValue", "customCalloutExtensionIdValue")

	if id.AccessPackageResourceRequestId != "accessPackageResourceRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRequestId'", id.AccessPackageResourceRequestId, "accessPackageResourceRequestIdValue")
	}

	if id.CustomCalloutExtensionId != "customCalloutExtensionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomCalloutExtensionId'", id.CustomCalloutExtensionId, "customCalloutExtensionIdValue")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionID("accessPackageResourceRequestIdValue", "customCalloutExtensionIdValue").ID()
	expected := "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/customWorkflowExtensions/customCalloutExtensionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId
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
			Input: "/identityGovernance/entitlementManagement/resourceRequests",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/customWorkflowExtensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/customWorkflowExtensions/customCalloutExtensionIdValue",
			Expected: &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId{
				AccessPackageResourceRequestId: "accessPackageResourceRequestIdValue",
				CustomCalloutExtensionId:       "customCalloutExtensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/customWorkflowExtensions/customCalloutExtensionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageResourceRequestId != v.Expected.AccessPackageResourceRequestId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRequestId", v.Expected.AccessPackageResourceRequestId, actual.AccessPackageResourceRequestId)
		}

		if actual.CustomCalloutExtensionId != v.Expected.CustomCalloutExtensionId {
			t.Fatalf("Expected %q but got %q for CustomCalloutExtensionId", v.Expected.CustomCalloutExtensionId, actual.CustomCalloutExtensionId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId
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
			Input: "/identityGovernance/entitlementManagement/resourceRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe/cAtAlOg",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/customWorkflowExtensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe/cAtAlOg/cUsToMwOrKfLoWeXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/customWorkflowExtensions/customCalloutExtensionIdValue",
			Expected: &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId{
				AccessPackageResourceRequestId: "accessPackageResourceRequestIdValue",
				CustomCalloutExtensionId:       "customCalloutExtensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/customWorkflowExtensions/customCalloutExtensionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe/cAtAlOg/cUsToMwOrKfLoWeXtEnSiOnS/cUsToMcAlLoUtExTeNsIoNiDvAlUe",
			Expected: &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId{
				AccessPackageResourceRequestId: "aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe",
				CustomCalloutExtensionId:       "cUsToMcAlLoUtExTeNsIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe/cAtAlOg/cUsToMwOrKfLoWeXtEnSiOnS/cUsToMcAlLoUtExTeNsIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageResourceRequestId != v.Expected.AccessPackageResourceRequestId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRequestId", v.Expected.AccessPackageResourceRequestId, actual.AccessPackageResourceRequestId)
		}

		if actual.CustomCalloutExtensionId != v.Expected.CustomCalloutExtensionId {
			t.Fatalf("Expected %q but got %q for CustomCalloutExtensionId", v.Expected.CustomCalloutExtensionId, actual.CustomCalloutExtensionId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
