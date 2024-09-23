package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID("accessPackageResourceId", "accessPackageResourceScopeId", "accessPackageResourceRoleId")

	if id.AccessPackageResourceId != "accessPackageResourceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceId'", id.AccessPackageResourceId, "accessPackageResourceId")
	}

	if id.AccessPackageResourceScopeId != "accessPackageResourceScopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceScopeId'", id.AccessPackageResourceScopeId, "accessPackageResourceScopeId")
	}

	if id.AccessPackageResourceRoleId != "accessPackageResourceRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleId'", id.AccessPackageResourceRoleId, "accessPackageResourceRoleId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID("accessPackageResourceId", "accessPackageResourceScopeId", "accessPackageResourceRoleId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource/accessPackageResourceRoles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{
				AccessPackageResourceId:      "accessPackageResourceId",
				AccessPackageResourceScopeId: "accessPackageResourceScopeId",
				AccessPackageResourceRoleId:  "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleID(v.Input)
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

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeScOpEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeScOpEs/aCcEsSpAcKaGeReSoUrCeScOpEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeScOpEs/aCcEsSpAcKaGeReSoUrCeScOpEiD/aCcEsSpAcKaGeReSoUrCe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource/accessPackageResourceRoles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeScOpEs/aCcEsSpAcKaGeReSoUrCeScOpEiD/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeRoLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{
				AccessPackageResourceId:      "accessPackageResourceId",
				AccessPackageResourceScopeId: "accessPackageResourceScopeId",
				AccessPackageResourceRoleId:  "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageResources/accessPackageResourceId/accessPackageResourceScopes/accessPackageResourceScopeId/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeScOpEs/aCcEsSpAcKaGeReSoUrCeScOpEiD/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeRoLeS/aCcEsSpAcKaGeReSoUrCeRoLeId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{
				AccessPackageResourceId:      "aCcEsSpAcKaGeReSoUrCeId",
				AccessPackageResourceScopeId: "aCcEsSpAcKaGeReSoUrCeScOpEiD",
				AccessPackageResourceRoleId:  "aCcEsSpAcKaGeReSoUrCeRoLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeId/aCcEsSpAcKaGeReSoUrCeScOpEs/aCcEsSpAcKaGeReSoUrCeScOpEiD/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeRoLeS/aCcEsSpAcKaGeReSoUrCeRoLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleIDInsensitively(v.Input)
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

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
