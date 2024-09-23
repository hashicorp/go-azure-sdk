package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleID("accessPackageAssignmentId", "accessPackageResourceRoleScopeId", "accessPackageResourceRoleId")

	if id.AccessPackageAssignmentId != "accessPackageAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentId'", id.AccessPackageAssignmentId, "accessPackageAssignmentId")
	}

	if id.AccessPackageResourceRoleScopeId != "accessPackageResourceRoleScopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleScopeId'", id.AccessPackageResourceRoleScopeId, "accessPackageResourceRoleScopeId")
	}

	if id.AccessPackageResourceRoleId != "accessPackageResourceRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleId'", id.AccessPackageResourceRoleId, "accessPackageResourceRoleId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleID("accessPackageAssignmentId", "accessPackageResourceRoleScopeId", "accessPackageResourceRoleId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageResourceRoleScopes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId{
				AccessPackageAssignmentId:        "accessPackageAssignmentId",
				AccessPackageResourceRoleScopeId: "accessPackageResourceRoleScopeId",
				AccessPackageResourceRoleId:      "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageAssignmentId != v.Expected.AccessPackageAssignmentId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentId", v.Expected.AccessPackageAssignmentId, actual.AccessPackageAssignmentId)
		}

		if actual.AccessPackageResourceRoleScopeId != v.Expected.AccessPackageResourceRoleScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleScopeId", v.Expected.AccessPackageResourceRoleScopeId, actual.AccessPackageResourceRoleScopeId)
		}

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageResourceRoleScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGe/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGe/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGe/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/aCcEsSpAcKaGeReSoUrCeRoLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGe/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/aCcEsSpAcKaGeReSoUrCeRoLe/aCcEsSpAcKaGeReSoUrCe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGe/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/aCcEsSpAcKaGeReSoUrCeRoLe/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeRoLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId{
				AccessPackageAssignmentId:        "accessPackageAssignmentId",
				AccessPackageResourceRoleScopeId: "accessPackageResourceRoleScopeId",
				AccessPackageResourceRoleId:      "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignments/accessPackageAssignmentId/accessPackage/accessPackageResourceRoleScopes/accessPackageResourceRoleScopeId/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGe/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/aCcEsSpAcKaGeReSoUrCeRoLe/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeRoLeS/aCcEsSpAcKaGeReSoUrCeRoLeId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId{
				AccessPackageAssignmentId:        "aCcEsSpAcKaGeAsSiGnMeNtId",
				AccessPackageResourceRoleScopeId: "aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD",
				AccessPackageResourceRoleId:      "aCcEsSpAcKaGeReSoUrCeRoLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/aCcEsSpAcKaGe/aCcEsSpAcKaGeReSoUrCeRoLeScOpEs/aCcEsSpAcKaGeReSoUrCeRoLeScOpEiD/aCcEsSpAcKaGeReSoUrCeRoLe/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeRoLeS/aCcEsSpAcKaGeReSoUrCeRoLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageAssignmentId != v.Expected.AccessPackageAssignmentId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentId", v.Expected.AccessPackageAssignmentId, actual.AccessPackageAssignmentId)
		}

		if actual.AccessPackageResourceRoleScopeId != v.Expected.AccessPackageResourceRoleScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleScopeId", v.Expected.AccessPackageResourceRoleScopeId, actual.AccessPackageResourceRoleScopeId)
		}

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
