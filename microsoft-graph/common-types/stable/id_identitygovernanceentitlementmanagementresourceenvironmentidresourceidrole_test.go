package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId{}

func TestNewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID("accessPackageResourceEnvironmentId", "accessPackageResourceId", "accessPackageResourceRoleId")

	if id.AccessPackageResourceEnvironmentId != "accessPackageResourceEnvironmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceEnvironmentId'", id.AccessPackageResourceEnvironmentId, "accessPackageResourceEnvironmentId")
	}

	if id.AccessPackageResourceId != "accessPackageResourceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceId'", id.AccessPackageResourceId, "accessPackageResourceId")
	}

	if id.AccessPackageResourceRoleId != "accessPackageResourceRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleId'", id.AccessPackageResourceRoleId, "accessPackageResourceRoleId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID("accessPackageResourceEnvironmentId", "accessPackageResourceId", "accessPackageResourceRoleId").ID()
	expected := "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources/accessPackageResourceId/roles/accessPackageResourceRoleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId
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
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources/accessPackageResourceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources/accessPackageResourceId/roles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources/accessPackageResourceId/roles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId{
				AccessPackageResourceEnvironmentId: "accessPackageResourceEnvironmentId",
				AccessPackageResourceId:            "accessPackageResourceId",
				AccessPackageResourceRoleId:        "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources/accessPackageResourceId/roles/accessPackageResourceRoleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID(v.Input)
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

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId
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
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiD/rEsOuRcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources/accessPackageResourceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiD/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources/accessPackageResourceId/roles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiD/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId/rOlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources/accessPackageResourceId/roles/accessPackageResourceRoleId",
			Expected: &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId{
				AccessPackageResourceEnvironmentId: "accessPackageResourceEnvironmentId",
				AccessPackageResourceId:            "accessPackageResourceId",
				AccessPackageResourceRoleId:        "accessPackageResourceRoleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources/accessPackageResourceId/roles/accessPackageResourceRoleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiD/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId",
			Expected: &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId{
				AccessPackageResourceEnvironmentId: "aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiD",
				AccessPackageResourceId:            "aCcEsSpAcKaGeReSoUrCeId",
				AccessPackageResourceRoleId:        "aCcEsSpAcKaGeReSoUrCeRoLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiD/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIDInsensitively(v.Input)
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

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
