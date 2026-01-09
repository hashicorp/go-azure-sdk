package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleID("accessPackageAssignmentResourceRoleId")

	if id.AccessPackageAssignmentResourceRoleId != "accessPackageAssignmentResourceRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageAssignmentResourceRoleId'", id.AccessPackageAssignmentResourceRoleId, "accessPackageAssignmentResourceRoleId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleID("accessPackageAssignmentResourceRoleId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId
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
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId{
				AccessPackageAssignmentResourceRoleId: "accessPackageAssignmentResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleID(v.Input)
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

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId
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
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId{
				AccessPackageAssignmentResourceRoleId: "accessPackageAssignmentResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/accessPackageAssignmentResourceRoleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeS/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId{
				AccessPackageAssignmentResourceRoleId: "aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeS/aCcEsSpAcKaGeAsSiGnMeNtReSoUrCeRoLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
