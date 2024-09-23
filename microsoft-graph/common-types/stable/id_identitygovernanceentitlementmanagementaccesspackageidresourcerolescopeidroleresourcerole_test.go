package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleID("accessPackageId", "accessPackageResourceRoleScopeId", "accessPackageResourceRoleId")

	if id.AccessPackageId != "accessPackageId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageId'", id.AccessPackageId, "accessPackageId")
	}

	if id.AccessPackageResourceRoleScopeId != "accessPackageResourceRoleScopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleScopeId'", id.AccessPackageResourceRoleScopeId, "accessPackageResourceRoleScopeId")
	}

	if id.AccessPackageResourceRoleId != "accessPackageResourceRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleId'", id.AccessPackageResourceRoleId, "accessPackageResourceRoleId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleID("accessPackageId", "accessPackageResourceRoleScopeId", "accessPackageResourceRoleId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource/roles/accessPackageResourceRoleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/accessPackages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId/role",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource/roles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource/roles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId{
				AccessPackageId:                  "accessPackageId",
				AccessPackageResourceRoleScopeId: "accessPackageResourceRoleScopeId",
				AccessPackageResourceRoleId:      "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource/roles/accessPackageResourceRoleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageId != v.Expected.AccessPackageId {
			t.Fatalf("Expected %q but got %q for AccessPackageId", v.Expected.AccessPackageId, actual.AccessPackageId)
		}

		if actual.AccessPackageResourceRoleScopeId != v.Expected.AccessPackageResourceRoleScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleScopeId", v.Expected.AccessPackageResourceRoleScopeId, actual.AccessPackageResourceRoleScopeId)
		}

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/accessPackages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/rEsOuRcErOlEsCoPeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId/role",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/rOlE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/rOlE/rEsOuRcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource/roles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/rOlE/rEsOuRcE/rOlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource/roles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId{
				AccessPackageId:                  "accessPackageId",
				AccessPackageResourceRoleScopeId: "accessPackageResourceRoleScopeId",
				AccessPackageResourceRoleId:      "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/resourceRoleScopes/accessPackageResourceRoleScopeId/role/resource/roles/accessPackageResourceRoleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/rOlE/rEsOuRcE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId{
				AccessPackageId:                  "aCcEsSpAcKaGeId",
				AccessPackageResourceRoleScopeId: "aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD",
				AccessPackageResourceRoleId:      "aCcEsSpAcKaGeReSoUrCeRoLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/rEsOuRcErOlEsCoPeS/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/rOlE/rEsOuRcE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageId != v.Expected.AccessPackageId {
			t.Fatalf("Expected %q but got %q for AccessPackageId", v.Expected.AccessPackageId, actual.AccessPackageId)
		}

		if actual.AccessPackageResourceRoleScopeId != v.Expected.AccessPackageResourceRoleScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleScopeId", v.Expected.AccessPackageResourceRoleScopeId, actual.AccessPackageResourceRoleScopeId)
		}

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceRoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
