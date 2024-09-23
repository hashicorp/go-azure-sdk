package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleID("accessPackageAssignmentResourceRoleId", "accessPackageResourceRoleId")

	if id.AccessPackageAssignmentResourceRoleId != "accessPackageAssignmentResourceRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentResourceRoleId'", id.AccessPackageAssignmentResourceRoleId, "accessPackageAssignmentResourceRoleId")
	}

	if id.AccessPackageResourceRoleId != "accessPackageResourceRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleId'", id.AccessPackageResourceRoleId, "accessPackageResourceRoleId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleID("accessPackageAssignmentResourceRoleId", "accessPackageResourceRoleId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageResourceRole",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageResourceRole/accessPackageResource",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId{
				AccessPackageAssignmentResourceRoleId: "accessPackageAssignmentResourceRoleId",
				AccessPackageResourceRoleId:           "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageAssignmentResourceRoleId != v.Expected.AccessPackageAssignmentResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentResourceRoleId", v.Expected.AccessPackageAssignmentResourceRoleId, actual.AccessPackageAssignmentResourceRoleId)
		}

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeS/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageResourceRole",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeS/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeId/aCcEsSpAcKaGeReSoUrCeRoLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageResourceRole/accessPackageResource",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeS/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeId/aCcEsSpAcKaGeReSoUrCeRoLe/aCcEsSpAcKaGeReSoUrCe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeS/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeId/aCcEsSpAcKaGeReSoUrCeRoLe/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeRoLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId{
				AccessPackageAssignmentResourceRoleId: "accessPackageAssignmentResourceRoleId",
				AccessPackageResourceRoleId:           "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles/accessPackageResourceRoleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeS/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeId/aCcEsSpAcKaGeReSoUrCeRoLe/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeRoLeS/aCcEsSpAcKaGeReSoUrCeRoLeId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId{
				AccessPackageAssignmentResourceRoleId: "aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeId",
				AccessPackageResourceRoleId:           "aCcEsSpAcKaGeReSoUrCeRoLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeS/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeId/aCcEsSpAcKaGeReSoUrCeRoLe/aCcEsSpAcKaGeReSoUrCe/aCcEsSpAcKaGeReSoUrCeRoLeS/aCcEsSpAcKaGeReSoUrCeRoLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageAssignmentResourceRoleId != v.Expected.AccessPackageAssignmentResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentResourceRoleId", v.Expected.AccessPackageAssignmentResourceRoleId, actual.AccessPackageAssignmentResourceRoleId)
		}

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
