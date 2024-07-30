package entitlementmanagementresourceenvironmentresourcescoperesourceroleresourceenvironment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId{}

func TestNewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleID("accessPackageResourceEnvironmentIdValue", "accessPackageResourceIdValue", "accessPackageResourceScopeIdValue", "accessPackageResourceRoleIdValue")

	if id.AccessPackageResourceEnvironmentId != "accessPackageResourceEnvironmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceEnvironmentId'", id.AccessPackageResourceEnvironmentId, "accessPackageResourceEnvironmentIdValue")
	}

	if id.AccessPackageResourceId != "accessPackageResourceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceId'", id.AccessPackageResourceId, "accessPackageResourceIdValue")
	}

	if id.AccessPackageResourceScopeId != "accessPackageResourceScopeIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceScopeId'", id.AccessPackageResourceScopeId, "accessPackageResourceScopeIdValue")
	}

	if id.AccessPackageResourceRoleId != "accessPackageResourceRoleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleId'", id.AccessPackageResourceRoleId, "accessPackageResourceRoleIdValue")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleID("accessPackageResourceEnvironmentIdValue", "accessPackageResourceIdValue", "accessPackageResourceScopeIdValue", "accessPackageResourceRoleIdValue").ID()
	expected := "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/scopes/accessPackageResourceScopeIdValue/resource/roles/accessPackageResourceRoleIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/scopes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/scopes/accessPackageResourceScopeIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/scopes/accessPackageResourceScopeIdValue/resource",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/scopes/accessPackageResourceScopeIdValue/resource/roles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/scopes/accessPackageResourceScopeIdValue/resource/roles/accessPackageResourceRoleIdValue",
			Expected: &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId{
				AccessPackageResourceEnvironmentId: "accessPackageResourceEnvironmentIdValue",
				AccessPackageResourceId:            "accessPackageResourceIdValue",
				AccessPackageResourceScopeId:       "accessPackageResourceScopeIdValue",
				AccessPackageResourceRoleId:        "accessPackageResourceRoleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/scopes/accessPackageResourceScopeIdValue/resource/roles/accessPackageResourceRoleIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleID(v.Input)
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

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId
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
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe/rEsOuRcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/scopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE/sCoPeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/scopes/accessPackageResourceScopeIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/scopes/accessPackageResourceScopeIdValue/resource",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe/rEsOuRcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/scopes/accessPackageResourceScopeIdValue/resource/roles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe/rEsOuRcE/rOlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/scopes/accessPackageResourceScopeIdValue/resource/roles/accessPackageResourceRoleIdValue",
			Expected: &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId{
				AccessPackageResourceEnvironmentId: "accessPackageResourceEnvironmentIdValue",
				AccessPackageResourceId:            "accessPackageResourceIdValue",
				AccessPackageResourceScopeId:       "accessPackageResourceScopeIdValue",
				AccessPackageResourceRoleId:        "accessPackageResourceRoleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/scopes/accessPackageResourceScopeIdValue/resource/roles/accessPackageResourceRoleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe/rEsOuRcE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE",
			Expected: &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId{
				AccessPackageResourceEnvironmentId: "aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe",
				AccessPackageResourceId:            "aCcEsSpAcKaGeReSoUrCeIdVaLuE",
				AccessPackageResourceScopeId:       "aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe",
				AccessPackageResourceRoleId:        "aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe/rEsOuRcE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleIDInsensitively(v.Input)
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

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
