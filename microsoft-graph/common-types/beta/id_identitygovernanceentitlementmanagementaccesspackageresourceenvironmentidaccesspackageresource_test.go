package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceID("accessPackageResourceEnvironmentIdValue", "accessPackageResourceIdValue")

	if id.AccessPackageResourceEnvironmentId != "accessPackageResourceEnvironmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceEnvironmentId'", id.AccessPackageResourceEnvironmentId, "accessPackageResourceEnvironmentIdValue")
	}

	if id.AccessPackageResourceId != "accessPackageResourceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceId'", id.AccessPackageResourceId, "accessPackageResourceIdValue")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceID("accessPackageResourceEnvironmentIdValue", "accessPackageResourceIdValue").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/accessPackageResourceEnvironmentIdValue/accessPackageResources/accessPackageResourceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId
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
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/accessPackageResourceEnvironmentIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/accessPackageResourceEnvironmentIdValue/accessPackageResources",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/accessPackageResourceEnvironmentIdValue/accessPackageResources/accessPackageResourceIdValue",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId{
				AccessPackageResourceEnvironmentId: "accessPackageResourceEnvironmentIdValue",
				AccessPackageResourceId:            "accessPackageResourceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/accessPackageResourceEnvironmentIdValue/accessPackageResources/accessPackageResourceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceID(v.Input)
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

		if actual.AccessPackageResourceId != v.Expected.AccessPackageResourceId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceId", v.Expected.AccessPackageResourceId, actual.AccessPackageResourceId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId
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
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/accessPackageResourceEnvironmentIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTs/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/accessPackageResourceEnvironmentIdValue/accessPackageResources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTs/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe/aCcEsSpAcKaGeReSoUrCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/accessPackageResourceEnvironmentIdValue/accessPackageResources/accessPackageResourceIdValue",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId{
				AccessPackageResourceEnvironmentId: "accessPackageResourceEnvironmentIdValue",
				AccessPackageResourceId:            "accessPackageResourceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/accessPackageResourceEnvironmentIdValue/accessPackageResources/accessPackageResourceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTs/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeIdVaLuE",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId{
				AccessPackageResourceEnvironmentId: "aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe",
				AccessPackageResourceId:            "aCcEsSpAcKaGeReSoUrCeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTs/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe/aCcEsSpAcKaGeReSoUrCeS/aCcEsSpAcKaGeReSoUrCeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceIDInsensitively(v.Input)
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

		if actual.AccessPackageResourceId != v.Expected.AccessPackageResourceId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceId", v.Expected.AccessPackageResourceId, actual.AccessPackageResourceId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentIdAccessPackageResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
