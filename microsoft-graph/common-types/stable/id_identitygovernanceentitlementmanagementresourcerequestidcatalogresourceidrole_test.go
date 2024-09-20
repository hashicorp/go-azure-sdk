package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId{}

func TestNewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleID("accessPackageResourceRequestId", "accessPackageResourceId", "accessPackageResourceRoleId")

	if id.AccessPackageResourceRequestId != "accessPackageResourceRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRequestId'", id.AccessPackageResourceRequestId, "accessPackageResourceRequestId")
	}

	if id.AccessPackageResourceId != "accessPackageResourceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceId'", id.AccessPackageResourceId, "accessPackageResourceId")
	}

	if id.AccessPackageResourceRoleId != "accessPackageResourceRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleId'", id.AccessPackageResourceRoleId, "accessPackageResourceRoleId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleID("accessPackageResourceRequestId", "accessPackageResourceId", "accessPackageResourceRoleId").ID()
	expected := "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestId/catalog/resources/accessPackageResourceId/roles/accessPackageResourceRoleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId
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
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestId/catalog",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestId/catalog/resources",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestId/catalog/resources/accessPackageResourceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestId/catalog/resources/accessPackageResourceId/roles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestId/catalog/resources/accessPackageResourceId/roles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId{
				AccessPackageResourceRequestId: "accessPackageResourceRequestId",
				AccessPackageResourceId:        "accessPackageResourceId",
				AccessPackageResourceRoleId:    "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestId/catalog/resources/accessPackageResourceId/roles/accessPackageResourceRoleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleID(v.Input)
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

		if actual.AccessPackageResourceId != v.Expected.AccessPackageResourceId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceId", v.Expected.AccessPackageResourceId, actual.AccessPackageResourceId)
		}

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId
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
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestId/catalog",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiD/cAtAlOg",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestId/catalog/resources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiD/cAtAlOg/rEsOuRcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestId/catalog/resources/accessPackageResourceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiD/cAtAlOg/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestId/catalog/resources/accessPackageResourceId/roles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiD/cAtAlOg/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId/rOlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestId/catalog/resources/accessPackageResourceId/roles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId{
				AccessPackageResourceRequestId: "accessPackageResourceRequestId",
				AccessPackageResourceId:        "accessPackageResourceId",
				AccessPackageResourceRoleId:    "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceRequests/accessPackageResourceRequestId/catalog/resources/accessPackageResourceId/roles/accessPackageResourceRoleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiD/cAtAlOg/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId",
			Expected: &IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId{
				AccessPackageResourceRequestId: "aCcEsSpAcKaGeReSoUrCeReQuEsTiD",
				AccessPackageResourceId:        "aCcEsSpAcKaGeReSoUrCeId",
				AccessPackageResourceRoleId:    "aCcEsSpAcKaGeReSoUrCeRoLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcErEqUeStS/aCcEsSpAcKaGeReSoUrCeReQuEsTiD/cAtAlOg/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIDInsensitively(v.Input)
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

		if actual.AccessPackageResourceId != v.Expected.AccessPackageResourceId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceId", v.Expected.AccessPackageResourceId, actual.AccessPackageResourceId)
		}

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
