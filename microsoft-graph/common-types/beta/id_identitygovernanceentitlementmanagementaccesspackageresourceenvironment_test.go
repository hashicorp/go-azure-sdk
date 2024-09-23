package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentID("accessPackageResourceEnvironmentId")

	if id.AccessPackageResourceEnvironmentId != "accessPackageResourceEnvironmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceEnvironmentId'", id.AccessPackageResourceEnvironmentId, "accessPackageResourceEnvironmentId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentID("accessPackageResourceEnvironmentId").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/accessPackageResourceEnvironmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/accessPackageResourceEnvironmentId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId{
				AccessPackageResourceEnvironmentId: "accessPackageResourceEnvironmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/accessPackageResourceEnvironmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageResourceEnvironmentId != v.Expected.AccessPackageResourceEnvironmentId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceEnvironmentId", v.Expected.AccessPackageResourceEnvironmentId, actual.AccessPackageResourceEnvironmentId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId
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
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/accessPackageResourceEnvironmentId",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId{
				AccessPackageResourceEnvironmentId: "accessPackageResourceEnvironmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/accessPackageResourceEnvironmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTs/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiD",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId{
				AccessPackageResourceEnvironmentId: "aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTs/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageResourceEnvironmentId != v.Expected.AccessPackageResourceEnvironmentId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceEnvironmentId", v.Expected.AccessPackageResourceEnvironmentId, actual.AccessPackageResourceEnvironmentId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
