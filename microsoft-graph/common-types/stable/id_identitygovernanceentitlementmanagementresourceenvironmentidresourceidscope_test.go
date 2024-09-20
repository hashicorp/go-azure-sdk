package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId{}

func TestNewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeID("accessPackageResourceEnvironmentId", "accessPackageResourceId", "accessPackageResourceScopeId")

	if id.AccessPackageResourceEnvironmentId != "accessPackageResourceEnvironmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceEnvironmentId'", id.AccessPackageResourceEnvironmentId, "accessPackageResourceEnvironmentId")
	}

	if id.AccessPackageResourceId != "accessPackageResourceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceId'", id.AccessPackageResourceId, "accessPackageResourceId")
	}

	if id.AccessPackageResourceScopeId != "accessPackageResourceScopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceScopeId'", id.AccessPackageResourceScopeId, "accessPackageResourceScopeId")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeID("accessPackageResourceEnvironmentId", "accessPackageResourceId", "accessPackageResourceScopeId").ID()
	expected := "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources/accessPackageResourceId/scopes/accessPackageResourceScopeId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId
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
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources/accessPackageResourceId/scopes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources/accessPackageResourceId/scopes/accessPackageResourceScopeId",
			Expected: &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId{
				AccessPackageResourceEnvironmentId: "accessPackageResourceEnvironmentId",
				AccessPackageResourceId:            "accessPackageResourceId",
				AccessPackageResourceScopeId:       "accessPackageResourceScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources/accessPackageResourceId/scopes/accessPackageResourceScopeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeID(v.Input)
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

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId
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
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources/accessPackageResourceId/scopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiD/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId/sCoPeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources/accessPackageResourceId/scopes/accessPackageResourceScopeId",
			Expected: &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId{
				AccessPackageResourceEnvironmentId: "accessPackageResourceEnvironmentId",
				AccessPackageResourceId:            "accessPackageResourceId",
				AccessPackageResourceScopeId:       "accessPackageResourceScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentId/resources/accessPackageResourceId/scopes/accessPackageResourceScopeId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiD/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiD",
			Expected: &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId{
				AccessPackageResourceEnvironmentId: "aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiD",
				AccessPackageResourceId:            "aCcEsSpAcKaGeReSoUrCeId",
				AccessPackageResourceScopeId:       "aCcEsSpAcKaGeReSoUrCeScOpEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiD/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeId/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIDInsensitively(v.Input)
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

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
