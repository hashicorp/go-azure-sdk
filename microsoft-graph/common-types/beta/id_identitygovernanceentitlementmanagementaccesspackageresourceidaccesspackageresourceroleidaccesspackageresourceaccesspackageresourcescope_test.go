package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeID("accessPackageResourceId", "accessPackageResourceRoleId", "accessPackageResourceScopeId")

	if id.AccessPackageResourceId != "accessPackageResourceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceId'", id.AccessPackageResourceId, "accessPackageResourceId")
	}

	if id.AccessPackageResourceRoleId != "accessPackageResourceRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleId'", id.AccessPackageResourceRoleId, "accessPackageResourceRoleId")
	}

	if id.AccessPackageResourceScopeId != "accessPackageResourceScopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceScopeId'", id.AccessPackageResourceScopeId, "accessPackageResourceScopeId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeID("accessPackageResourceId", "accessPackageResourceRoleId", "accessPackageResourceScopeId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceRoles/accessPackageResourceRoleId/accessPackageResource/accessPackageResourceScopes/accessPackageResourceScopeId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageResources",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceRoles",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceRoles/accessPackageResourceRoleId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceRoles/accessPackageResourceRoleId/accessPackageResource",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceRoles/accessPackageResourceRoleId/accessPackageResource/accessPackageResourceScopes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceRoles/accessPackageResourceRoleId/accessPackageResource/accessPackageResourceScopes/accessPackageResourceScopeId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId{
				AccessPackageResourceId:      "accessPackageResourceId",
				AccessPackageResourceRoleId:  "accessPackageResourceRoleId",
				AccessPackageResourceScopeId: "accessPackageResourceScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceRoles/accessPackageResourceRoleId/accessPackageResource/accessPackageResourceScopes/accessPackageResourceScopeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageResourceId != v.Expected.AccessPackageResourceId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceId", v.Expected.AccessPackageResourceId, actual.AccessPackageResourceId)
		}

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageResources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceRoles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeRoLeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceRoles/accessPackageResourceRoleId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeRoLeS/aCcEsSpAcKaGeReSoUrCeRoLeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceRoles/accessPackageResourceRoleId/accessPackageResource",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeRoLeS/aCcEsSpAcKaGeReSoUrCeRoLeId/aCcEsSpAcKaGeReSoUrCe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceRoles/accessPackageResourceRoleId/accessPackageResource/accessPackageResourceScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeRoLeS/aCcEsSpAcKaGeReSoUrCeRoLeId/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeScOpEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceRoles/accessPackageResourceRoleId/accessPackageResource/accessPackageResourceScopes/accessPackageResourceScopeId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId{
				AccessPackageResourceId:      "accessPackageResourceId",
				AccessPackageResourceRoleId:  "accessPackageResourceRoleId",
				AccessPackageResourceScopeId: "accessPackageResourceScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceRoles/accessPackageResourceRoleId/accessPackageResource/accessPackageResourceScopes/accessPackageResourceScopeId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeRoLeS/aCcEsSpAcKaGeReSoUrCeRoLeId/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeScOpEs/aCcEsSpAcKaGeReSoUrCeScOpEiD",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId{
				AccessPackageResourceId:      "aCcEsSpAcKaGeReSoUrCeId",
				AccessPackageResourceRoleId:  "aCcEsSpAcKaGeReSoUrCeRoLeId",
				AccessPackageResourceScopeId: "aCcEsSpAcKaGeReSoUrCeScOpEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeRoLeS/aCcEsSpAcKaGeReSoUrCeRoLeId/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeScOpEs/aCcEsSpAcKaGeReSoUrCeScOpEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageResourceId != v.Expected.AccessPackageResourceId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceId", v.Expected.AccessPackageResourceId, actual.AccessPackageResourceId)
		}

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
