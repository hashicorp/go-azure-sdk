package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentID("accessPackageAssignmentResourceRoleId", "accessPackageAssignmentId")

	if id.AccessPackageAssignmentResourceRoleId != "accessPackageAssignmentResourceRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentResourceRoleId'", id.AccessPackageAssignmentResourceRoleId, "accessPackageAssignmentResourceRoleId")
	}

	if id.AccessPackageAssignmentId != "accessPackageAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentId'", id.AccessPackageAssignmentId, "accessPackageAssignmentId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentID("accessPackageAssignmentResourceRoleId", "accessPackageAssignmentId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageAssignments/accessPackageAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageAssignments/accessPackageAssignmentId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId{
				AccessPackageAssignmentResourceRoleId: "accessPackageAssignmentResourceRoleId",
				AccessPackageAssignmentId:             "accessPackageAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageAssignments/accessPackageAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentID(v.Input)
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

		if actual.AccessPackageAssignmentId != v.Expected.AccessPackageAssignmentId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentId", v.Expected.AccessPackageAssignmentId, actual.AccessPackageAssignmentId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeS/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeId/aCcEsSpAcKaGeAsSiGnMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageAssignments/accessPackageAssignmentId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId{
				AccessPackageAssignmentResourceRoleId: "accessPackageAssignmentResourceRoleId",
				AccessPackageAssignmentId:             "accessPackageAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/accessPackageAssignments/accessPackageAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeS/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeId/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId{
				AccessPackageAssignmentResourceRoleId: "aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeId",
				AccessPackageAssignmentId:             "aCcEsSpAcKaGeAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeS/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeId/aCcEsSpAcKaGeAsSiGnMeNtS/aCcEsSpAcKaGeAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentIDInsensitively(v.Input)
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

		if actual.AccessPackageAssignmentId != v.Expected.AccessPackageAssignmentId {
			t.Fatalf("Expected %q but got %q for AccessPackageAssignmentId", v.Expected.AccessPackageAssignmentId, actual.AccessPackageAssignmentId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
