package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId{}

func TestNewIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageID("accessPackageId", "accessPackageId1")

	if id.AccessPackageId != "accessPackageId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageId'", id.AccessPackageId, "accessPackageId")
	}

	if id.AccessPackageId1 != "accessPackageId1" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageId1'", id.AccessPackageId1, "accessPackageId1")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageID("accessPackageId", "accessPackageId1").ID()
	expected := "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/incompatibleAccessPackages/accessPackageId1"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId
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
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/incompatibleAccessPackages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/incompatibleAccessPackages/accessPackageId1",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId{
				AccessPackageId:  "accessPackageId",
				AccessPackageId1: "accessPackageId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/incompatibleAccessPackages/accessPackageId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageID(v.Input)
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

func TestParseIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId
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
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/incompatibleAccessPackages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/iNcOmPaTiBlEaCcEsSpAcKaGeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/incompatibleAccessPackages/accessPackageId1",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId{
				AccessPackageId:  "accessPackageId",
				AccessPackageId1: "accessPackageId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/accessPackages/accessPackageId/incompatibleAccessPackages/accessPackageId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/iNcOmPaTiBlEaCcEsSpAcKaGeS/aCcEsSpAcKaGeId1",
			Expected: &IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId{
				AccessPackageId:  "aCcEsSpAcKaGeId",
				AccessPackageId1: "aCcEsSpAcKaGeId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/aCcEsSpAcKaGeS/aCcEsSpAcKaGeId/iNcOmPaTiBlEaCcEsSpAcKaGeS/aCcEsSpAcKaGeId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
