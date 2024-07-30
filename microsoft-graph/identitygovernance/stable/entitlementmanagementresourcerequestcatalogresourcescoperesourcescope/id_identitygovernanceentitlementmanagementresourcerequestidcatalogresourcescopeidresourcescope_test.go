package entitlementmanagementresourcerequestcatalogresourcescoperesourcescope

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId{}

func TestNewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeID("accessPackageResourceRequestIdValue", "accessPackageResourceScopeIdValue", "accessPackageResourceScopeId1Value")

	if id.AccessPackageResourceRequestId != "accessPackageResourceRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRequestId'", id.AccessPackageResourceRequestId, "accessPackageResourceRequestIdValue")
	}

	if id.AccessPackageResourceScopeId != "accessPackageResourceScopeIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceScopeId'", id.AccessPackageResourceScopeId, "accessPackageResourceScopeIdValue")
	}

	if id.AccessPackageResourceScopeId1 != "accessPackageResourceScopeId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceScopeId1'", id.AccessPackageResourceScopeId1, "accessPackageResourceScopeId1Value")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeID("accessPackageResourceRequestIdValue", "accessPackageResourceScopeIdValue", "accessPackageResourceScopeId1Value").ID()
	expected := "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceScopes/accessPackageResourceScopeIdValue/resource/scopes/accessPackageResourceScopeId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId
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
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceScopes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceScopes/accessPackageResourceScopeIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceScopes/accessPackageResourceScopeIdValue/resource",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceScopes/accessPackageResourceScopeIdValue/resource/scopes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceScopes/accessPackageResourceScopeIdValue/resource/scopes/accessPackageResourceScopeId1Value",
			Expected: &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId{
				AccessPackageResourceRequestId: "accessPackageResourceRequestIdValue",
				AccessPackageResourceScopeId:   "accessPackageResourceScopeIdValue",
				AccessPackageResourceScopeId1:  "accessPackageResourceScopeId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceScopes/accessPackageResourceScopeIdValue/resource/scopes/accessPackageResourceScopeId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeID(v.Input)
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

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

		if actual.AccessPackageResourceScopeId1 != v.Expected.AccessPackageResourceScopeId1 {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId1", v.Expected.AccessPackageResourceScopeId1, actual.AccessPackageResourceScopeId1)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId
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
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe/cAtAlOg/rEsOuRcEsCoPeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceScopes/accessPackageResourceScopeIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe/cAtAlOg/rEsOuRcEsCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceScopes/accessPackageResourceScopeIdValue/resource",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe/cAtAlOg/rEsOuRcEsCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe/rEsOuRcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceScopes/accessPackageResourceScopeIdValue/resource/scopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe/cAtAlOg/rEsOuRcEsCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe/rEsOuRcE/sCoPeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceScopes/accessPackageResourceScopeIdValue/resource/scopes/accessPackageResourceScopeId1Value",
			Expected: &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId{
				AccessPackageResourceRequestId: "accessPackageResourceRequestIdValue",
				AccessPackageResourceScopeId:   "accessPackageResourceScopeIdValue",
				AccessPackageResourceScopeId1:  "accessPackageResourceScopeId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceScopes/accessPackageResourceScopeIdValue/resource/scopes/accessPackageResourceScopeId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe/cAtAlOg/rEsOuRcEsCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe/rEsOuRcE/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiD1VaLuE",
			Expected: &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId{
				AccessPackageResourceRequestId: "aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe",
				AccessPackageResourceScopeId:   "aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe",
				AccessPackageResourceScopeId1:  "aCcEsSpAcKaGeReSoUrCeScOpEiD1VaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe/cAtAlOg/rEsOuRcEsCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe/rEsOuRcE/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiD1VaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeIDInsensitively(v.Input)
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

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

		if actual.AccessPackageResourceScopeId1 != v.Expected.AccessPackageResourceScopeId1 {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId1", v.Expected.AccessPackageResourceScopeId1, actual.AccessPackageResourceScopeId1)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceScopeIdResourceScopeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
