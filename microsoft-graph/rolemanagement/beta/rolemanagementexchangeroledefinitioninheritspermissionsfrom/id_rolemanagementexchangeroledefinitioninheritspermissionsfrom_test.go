package rolemanagementexchangeroledefinitioninheritspermissionsfrom

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementExchangeRoleDefinitionInheritsPermissionsFromId{}

func TestNewRoleManagementExchangeRoleDefinitionInheritsPermissionsFromID(t *testing.T) {
	id := NewRoleManagementExchangeRoleDefinitionInheritsPermissionsFromID("unifiedRoleDefinitionIdValue", "unifiedRoleDefinitionId1Value")

	if id.UnifiedRoleDefinitionId != "unifiedRoleDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleDefinitionId'", id.UnifiedRoleDefinitionId, "unifiedRoleDefinitionIdValue")
	}

	if id.UnifiedRoleDefinitionId1 != "unifiedRoleDefinitionId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleDefinitionId1'", id.UnifiedRoleDefinitionId1, "unifiedRoleDefinitionId1Value")
	}
}

func TestFormatRoleManagementExchangeRoleDefinitionInheritsPermissionsFromID(t *testing.T) {
	actual := NewRoleManagementExchangeRoleDefinitionInheritsPermissionsFromID("unifiedRoleDefinitionIdValue", "unifiedRoleDefinitionId1Value").ID()
	expected := "/roleManagement/exchange/roleDefinitions/unifiedRoleDefinitionIdValue/inheritsPermissionsFrom/unifiedRoleDefinitionId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementExchangeRoleDefinitionInheritsPermissionsFromID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementExchangeRoleDefinitionInheritsPermissionsFromId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/exchange",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/exchange/roleDefinitions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/exchange/roleDefinitions/unifiedRoleDefinitionIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/exchange/roleDefinitions/unifiedRoleDefinitionIdValue/inheritsPermissionsFrom",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/exchange/roleDefinitions/unifiedRoleDefinitionIdValue/inheritsPermissionsFrom/unifiedRoleDefinitionId1Value",
			Expected: &RoleManagementExchangeRoleDefinitionInheritsPermissionsFromId{
				UnifiedRoleDefinitionId:  "unifiedRoleDefinitionIdValue",
				UnifiedRoleDefinitionId1: "unifiedRoleDefinitionId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/exchange/roleDefinitions/unifiedRoleDefinitionIdValue/inheritsPermissionsFrom/unifiedRoleDefinitionId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementExchangeRoleDefinitionInheritsPermissionsFromID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleDefinitionId != v.Expected.UnifiedRoleDefinitionId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleDefinitionId", v.Expected.UnifiedRoleDefinitionId, actual.UnifiedRoleDefinitionId)
		}

		if actual.UnifiedRoleDefinitionId1 != v.Expected.UnifiedRoleDefinitionId1 {
			t.Fatalf("Expected %q but got %q for UnifiedRoleDefinitionId1", v.Expected.UnifiedRoleDefinitionId1, actual.UnifiedRoleDefinitionId1)
		}

	}
}

func TestParseRoleManagementExchangeRoleDefinitionInheritsPermissionsFromIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementExchangeRoleDefinitionInheritsPermissionsFromId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/exchange",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/exchange/roleDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE/rOlEdEfInItIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/exchange/roleDefinitions/unifiedRoleDefinitionIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE/rOlEdEfInItIoNs/uNiFiEdRoLeDeFiNiTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/exchange/roleDefinitions/unifiedRoleDefinitionIdValue/inheritsPermissionsFrom",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE/rOlEdEfInItIoNs/uNiFiEdRoLeDeFiNiTiOnIdVaLuE/iNhErItSpErMiSsIoNsFrOm",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/exchange/roleDefinitions/unifiedRoleDefinitionIdValue/inheritsPermissionsFrom/unifiedRoleDefinitionId1Value",
			Expected: &RoleManagementExchangeRoleDefinitionInheritsPermissionsFromId{
				UnifiedRoleDefinitionId:  "unifiedRoleDefinitionIdValue",
				UnifiedRoleDefinitionId1: "unifiedRoleDefinitionId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/exchange/roleDefinitions/unifiedRoleDefinitionIdValue/inheritsPermissionsFrom/unifiedRoleDefinitionId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE/rOlEdEfInItIoNs/uNiFiEdRoLeDeFiNiTiOnIdVaLuE/iNhErItSpErMiSsIoNsFrOm/uNiFiEdRoLeDeFiNiTiOnId1vAlUe",
			Expected: &RoleManagementExchangeRoleDefinitionInheritsPermissionsFromId{
				UnifiedRoleDefinitionId:  "uNiFiEdRoLeDeFiNiTiOnIdVaLuE",
				UnifiedRoleDefinitionId1: "uNiFiEdRoLeDeFiNiTiOnId1vAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE/rOlEdEfInItIoNs/uNiFiEdRoLeDeFiNiTiOnIdVaLuE/iNhErItSpErMiSsIoNsFrOm/uNiFiEdRoLeDeFiNiTiOnId1vAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementExchangeRoleDefinitionInheritsPermissionsFromIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleDefinitionId != v.Expected.UnifiedRoleDefinitionId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleDefinitionId", v.Expected.UnifiedRoleDefinitionId, actual.UnifiedRoleDefinitionId)
		}

		if actual.UnifiedRoleDefinitionId1 != v.Expected.UnifiedRoleDefinitionId1 {
			t.Fatalf("Expected %q but got %q for UnifiedRoleDefinitionId1", v.Expected.UnifiedRoleDefinitionId1, actual.UnifiedRoleDefinitionId1)
		}

	}
}

func TestSegmentsForRoleManagementExchangeRoleDefinitionInheritsPermissionsFromId(t *testing.T) {
	segments := RoleManagementExchangeRoleDefinitionInheritsPermissionsFromId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementExchangeRoleDefinitionInheritsPermissionsFromId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
