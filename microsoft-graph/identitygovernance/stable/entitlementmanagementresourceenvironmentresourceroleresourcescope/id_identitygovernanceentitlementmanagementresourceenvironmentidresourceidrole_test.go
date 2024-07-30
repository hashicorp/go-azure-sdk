package entitlementmanagementresourceenvironmentresourceroleresourcescope

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId{}

func TestNewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID("accessPackageResourceEnvironmentIdValue", "accessPackageResourceIdValue", "accessPackageResourceRoleIdValue")

	if id.AccessPackageResourceEnvironmentId != "accessPackageResourceEnvironmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceEnvironmentId'", id.AccessPackageResourceEnvironmentId, "accessPackageResourceEnvironmentIdValue")
	}

	if id.AccessPackageResourceId != "accessPackageResourceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceId'", id.AccessPackageResourceId, "accessPackageResourceIdValue")
	}

	if id.AccessPackageResourceRoleId != "accessPackageResourceRoleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleId'", id.AccessPackageResourceRoleId, "accessPackageResourceRoleIdValue")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleID("accessPackageResourceEnvironmentIdValue", "accessPackageResourceIdValue", "accessPackageResourceRoleIdValue").ID()
	expected := "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/roles/accessPackageResourceRoleIdValue"
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
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/roles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/roles/accessPackageResourceRoleIdValue",
			Expected: &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId{
				AccessPackageResourceEnvironmentId: "accessPackageResourceEnvironmentIdValue",
				AccessPackageResourceId:            "accessPackageResourceIdValue",
				AccessPackageResourceRoleId:        "accessPackageResourceRoleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/roles/accessPackageResourceRoleIdValue/extra",
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
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/roles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE/rOlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/roles/accessPackageResourceRoleIdValue",
			Expected: &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId{
				AccessPackageResourceEnvironmentId: "accessPackageResourceEnvironmentIdValue",
				AccessPackageResourceId:            "accessPackageResourceIdValue",
				AccessPackageResourceRoleId:        "accessPackageResourceRoleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resourceEnvironments/accessPackageResourceEnvironmentIdValue/resources/accessPackageResourceIdValue/roles/accessPackageResourceRoleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE",
			Expected: &IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId{
				AccessPackageResourceEnvironmentId: "aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe",
				AccessPackageResourceId:            "aCcEsSpAcKaGeReSoUrCeIdVaLuE",
				AccessPackageResourceRoleId:        "aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEeNvIrOnMeNtS/aCcEsSpAcKaGeReSoUrCeEnViRoNmEnTiDvAlUe/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE/extra",
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
