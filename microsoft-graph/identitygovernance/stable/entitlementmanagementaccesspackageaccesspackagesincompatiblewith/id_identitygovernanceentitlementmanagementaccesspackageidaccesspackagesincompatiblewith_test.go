package entitlementmanagementaccesspackageaccesspackagesincompatiblewith

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithID("accessPackageIdValue", "accessPackageId1Value")

	if id.AccessPackageId != "accessPackageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageId'", id.AccessPackageId, "accessPackageIdValue")
	}

	if id.AccessPackageId1 != "accessPackageId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageId1'", id.AccessPackageId1, "accessPackageId1Value")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithID("accessPackageIdValue", "accessPackageId1Value").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/accessPackagesIncompatibleWith/accessPackageId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId
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
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/accessPackagesIncompatibleWith",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/accessPackagesIncompatibleWith/accessPackageId1Value",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId{
				AccessPackageId:  "accessPackageIdValue",
				AccessPackageId1: "accessPackageId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/accessPackagesIncompatibleWith/accessPackageId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithID(v.Input)
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

		if actual.AccessPackageId1 != v.Expected.AccessPackageId1 {
			t.Fatalf("Expected %q but got %q for AccessPackageId1", v.Expected.AccessPackageId1, actual.AccessPackageId1)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId
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
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/accessPackagesIncompatibleWith",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeIdVaLuE/aCcEsSpAcKaGeSiNcOmPaTiBlEwItH",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/accessPackagesIncompatibleWith/accessPackageId1Value",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId{
				AccessPackageId:  "accessPackageIdValue",
				AccessPackageId1: "accessPackageId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageIdValue/accessPackagesIncompatibleWith/accessPackageId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeIdVaLuE/aCcEsSpAcKaGeSiNcOmPaTiBlEwItH/aCcEsSpAcKaGeId1vAlUe",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId{
				AccessPackageId:  "aCcEsSpAcKaGeIdVaLuE",
				AccessPackageId1: "aCcEsSpAcKaGeId1vAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeIdVaLuE/aCcEsSpAcKaGeSiNcOmPaTiBlEwItH/aCcEsSpAcKaGeId1vAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithIDInsensitively(v.Input)
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

		if actual.AccessPackageId1 != v.Expected.AccessPackageId1 {
			t.Fatalf("Expected %q but got %q for AccessPackageId1", v.Expected.AccessPackageId1, actual.AccessPackageId1)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
