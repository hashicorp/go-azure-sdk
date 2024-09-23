package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeID("accessPackageResourceRoleScopeId")

	if id.AccessPackageResourceRoleScopeId != "accessPackageResourceRoleScopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleScopeId'", id.AccessPackageResourceRoleScopeId, "accessPackageResourceRoleScopeId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeID("accessPackageResourceRoleScopeId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceRoleScopes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId{
				AccessPackageResourceRoleScopeId: "accessPackageResourceRoleScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageResourceRoleScopeId != v.Expected.AccessPackageResourceRoleScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleScopeId", v.Expected.AccessPackageResourceRoleScopeId, actual.AccessPackageResourceRoleScopeId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceRoleScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId{
				AccessPackageResourceRoleScopeId: "accessPackageResourceRoleScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId{
				AccessPackageResourceRoleScopeId: "aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageResourceRoleScopeId != v.Expected.AccessPackageResourceRoleScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleScopeId", v.Expected.AccessPackageResourceRoleScopeId, actual.AccessPackageResourceRoleScopeId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
