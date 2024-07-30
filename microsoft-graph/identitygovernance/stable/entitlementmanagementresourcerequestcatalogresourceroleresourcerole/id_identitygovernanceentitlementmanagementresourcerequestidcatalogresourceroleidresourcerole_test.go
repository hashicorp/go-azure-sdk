package entitlementmanagementresourcerequestcatalogresourceroleresourcerole

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId{}

func TestNewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleID("accessPackageResourceRequestIdValue", "accessPackageResourceRoleIdValue", "accessPackageResourceRoleId1Value")

	if id.AccessPackageResourceRequestId != "accessPackageResourceRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRequestId'", id.AccessPackageResourceRequestId, "accessPackageResourceRequestIdValue")
	}

	if id.AccessPackageResourceRoleId != "accessPackageResourceRoleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleId'", id.AccessPackageResourceRoleId, "accessPackageResourceRoleIdValue")
	}

	if id.AccessPackageResourceRoleId1 != "accessPackageResourceRoleId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleId1'", id.AccessPackageResourceRoleId1, "accessPackageResourceRoleId1Value")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleID("accessPackageResourceRequestIdValue", "accessPackageResourceRoleIdValue", "accessPackageResourceRoleId1Value").ID()
	expected := "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceRoles/accessPackageResourceRoleIdValue/resource/roles/accessPackageResourceRoleId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceRoles",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceRoles/accessPackageResourceRoleIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceRoles/accessPackageResourceRoleIdValue/resource",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceRoles/accessPackageResourceRoleIdValue/resource/roles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceRoles/accessPackageResourceRoleIdValue/resource/roles/accessPackageResourceRoleId1Value",
			Expected: &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId{
				AccessPackageResourceRequestId: "accessPackageResourceRequestIdValue",
				AccessPackageResourceRoleId:    "accessPackageResourceRoleIdValue",
				AccessPackageResourceRoleId1:   "accessPackageResourceRoleId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceRoles/accessPackageResourceRoleIdValue/resource/roles/accessPackageResourceRoleId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleID(v.Input)
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

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

		if actual.AccessPackageResourceRoleId1 != v.Expected.AccessPackageResourceRoleId1 {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId1", v.Expected.AccessPackageResourceRoleId1, actual.AccessPackageResourceRoleId1)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceRoles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe/cAtAlOg/rEsOuRcErOlEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceRoles/accessPackageResourceRoleIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe/cAtAlOg/rEsOuRcErOlEs/aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceRoles/accessPackageResourceRoleIdValue/resource",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe/cAtAlOg/rEsOuRcErOlEs/aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE/rEsOuRcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceRoles/accessPackageResourceRoleIdValue/resource/roles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe/cAtAlOg/rEsOuRcErOlEs/aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE/rEsOuRcE/rOlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceRoles/accessPackageResourceRoleIdValue/resource/roles/accessPackageResourceRoleId1Value",
			Expected: &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId{
				AccessPackageResourceRequestId: "accessPackageResourceRequestIdValue",
				AccessPackageResourceRoleId:    "accessPackageResourceRoleIdValue",
				AccessPackageResourceRoleId1:   "accessPackageResourceRoleId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestIdValue/catalog/resourceRoles/accessPackageResourceRoleIdValue/resource/roles/accessPackageResourceRoleId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe/cAtAlOg/rEsOuRcErOlEs/aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE/rEsOuRcE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId1vAlUe",
			Expected: &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId{
				AccessPackageResourceRequestId: "aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe",
				AccessPackageResourceRoleId:    "aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE",
				AccessPackageResourceRoleId1:   "aCcEsSpAcKaGeReSoUrCeRoLeId1vAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiDvAlUe/cAtAlOg/rEsOuRcErOlEs/aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE/rEsOuRcE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId1vAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleIDInsensitively(v.Input)
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

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

		if actual.AccessPackageResourceRoleId1 != v.Expected.AccessPackageResourceRoleId1 {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId1", v.Expected.AccessPackageResourceRoleId1, actual.AccessPackageResourceRoleId1)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceRoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
