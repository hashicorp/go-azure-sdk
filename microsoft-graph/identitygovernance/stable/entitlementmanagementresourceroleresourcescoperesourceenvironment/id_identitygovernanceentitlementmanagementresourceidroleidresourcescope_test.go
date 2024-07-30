package entitlementmanagementresourceroleresourcescoperesourceenvironment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId{}

func TestNewIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeID(t *testing.T) {
	id := NewIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeID("accessPackageResourceIdValue", "accessPackageResourceRoleIdValue", "accessPackageResourceScopeIdValue")

	if id.AccessPackageResourceId != "accessPackageResourceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceId'", id.AccessPackageResourceId, "accessPackageResourceIdValue")
	}

	if id.AccessPackageResourceRoleId != "accessPackageResourceRoleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceRoleId'", id.AccessPackageResourceRoleId, "accessPackageResourceRoleIdValue")
	}

	if id.AccessPackageResourceScopeId != "accessPackageResourceScopeIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessPackageResourceScopeId'", id.AccessPackageResourceScopeId, "accessPackageResourceScopeIdValue")
	}
}

func TestFormatIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeID(t *testing.T) {
	actual := NewIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeID("accessPackageResourceIdValue", "accessPackageResourceRoleIdValue", "accessPackageResourceScopeIdValue").ID()
	expected := "/identityGovernance/entitlementManagement/resources/accessPackageResourceIdValue/roles/accessPackageResourceRoleIdValue/resource/scopes/accessPackageResourceScopeIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId
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
			Input: "/identityGovernance/entitlementManagement/resources",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resources/accessPackageResourceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resources/accessPackageResourceIdValue/roles",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resources/accessPackageResourceIdValue/roles/accessPackageResourceRoleIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resources/accessPackageResourceIdValue/roles/accessPackageResourceRoleIdValue/resource",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resources/accessPackageResourceIdValue/roles/accessPackageResourceRoleIdValue/resource/scopes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resources/accessPackageResourceIdValue/roles/accessPackageResourceRoleIdValue/resource/scopes/accessPackageResourceScopeIdValue",
			Expected: &IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId{
				AccessPackageResourceId:      "accessPackageResourceIdValue",
				AccessPackageResourceRoleId:  "accessPackageResourceRoleIdValue",
				AccessPackageResourceScopeId: "accessPackageResourceScopeIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resources/accessPackageResourceIdValue/roles/accessPackageResourceRoleIdValue/resource/scopes/accessPackageResourceScopeIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageResourceId != v.Expected.AccessPackageResourceId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceId", v.Expected.AccessPackageResourceId, actual.AccessPackageResourceId)
		}

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

	}
}

func TestParseIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId
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
			Input: "/identityGovernance/entitlementManagement/resources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resources/accessPackageResourceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resources/accessPackageResourceIdValue/roles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE/rOlEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resources/accessPackageResourceIdValue/roles/accessPackageResourceRoleIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resources/accessPackageResourceIdValue/roles/accessPackageResourceRoleIdValue/resource",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE/rEsOuRcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/entitlementManagement/resources/accessPackageResourceIdValue/roles/accessPackageResourceRoleIdValue/resource/scopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE/rEsOuRcE/sCoPeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/entitlementManagement/resources/accessPackageResourceIdValue/roles/accessPackageResourceRoleIdValue/resource/scopes/accessPackageResourceScopeIdValue",
			Expected: &IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId{
				AccessPackageResourceId:      "accessPackageResourceIdValue",
				AccessPackageResourceRoleId:  "accessPackageResourceRoleIdValue",
				AccessPackageResourceScopeId: "accessPackageResourceScopeIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/entitlementManagement/resources/accessPackageResourceIdValue/roles/accessPackageResourceRoleIdValue/resource/scopes/accessPackageResourceScopeIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE/rEsOuRcE/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe",
			Expected: &IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId{
				AccessPackageResourceId:      "aCcEsSpAcKaGeReSoUrCeIdVaLuE",
				AccessPackageResourceRoleId:  "aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE",
				AccessPackageResourceScopeId: "aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/eNtItLeMeNtMaNaGeMeNt/rEsOuRcEs/aCcEsSpAcKaGeReSoUrCeIdVaLuE/rOlEs/aCcEsSpAcKaGeReSoUrCeRoLeIdVaLuE/rEsOuRcE/sCoPeS/aCcEsSpAcKaGeReSoUrCeScOpEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessPackageResourceId != v.Expected.AccessPackageResourceId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceId", v.Expected.AccessPackageResourceId, actual.AccessPackageResourceId)
		}

		if actual.AccessPackageResourceRoleId != v.Expected.AccessPackageResourceRoleId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceRoleId", v.Expected.AccessPackageResourceRoleId, actual.AccessPackageResourceRoleId)
		}

		if actual.AccessPackageResourceScopeId != v.Expected.AccessPackageResourceScopeId {
			t.Fatalf("Expected %q but got %q for AccessPackageResourceScopeId", v.Expected.AccessPackageResourceScopeId, actual.AccessPackageResourceScopeId)
		}

	}
}

func TestSegmentsForIdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId(t *testing.T) {
	segments := IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
