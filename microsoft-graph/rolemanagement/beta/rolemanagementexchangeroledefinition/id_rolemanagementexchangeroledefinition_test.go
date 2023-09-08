package rolemanagementexchangeroledefinition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementExchangeRoleDefinitionId{}

func TestNewRoleManagementExchangeRoleDefinitionID(t *testing.T) {
	id := NewRoleManagementExchangeRoleDefinitionID("unifiedRoleDefinitionIdValue")

	if id.UnifiedRoleDefinitionId != "unifiedRoleDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleDefinitionId'", id.UnifiedRoleDefinitionId, "unifiedRoleDefinitionIdValue")
	}
}

func TestFormatRoleManagementExchangeRoleDefinitionID(t *testing.T) {
	actual := NewRoleManagementExchangeRoleDefinitionID("unifiedRoleDefinitionIdValue").ID()
	expected := "/roleManagement/exchange/roleDefinitions/unifiedRoleDefinitionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementExchangeRoleDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementExchangeRoleDefinitionId
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
			// Valid URI
			Input: "/roleManagement/exchange/roleDefinitions/unifiedRoleDefinitionIdValue",
			Expected: &RoleManagementExchangeRoleDefinitionId{
				UnifiedRoleDefinitionId: "unifiedRoleDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/exchange/roleDefinitions/unifiedRoleDefinitionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementExchangeRoleDefinitionID(v.Input)
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

	}
}

func TestParseRoleManagementExchangeRoleDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementExchangeRoleDefinitionId
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
			// Valid URI
			Input: "/roleManagement/exchange/roleDefinitions/unifiedRoleDefinitionIdValue",
			Expected: &RoleManagementExchangeRoleDefinitionId{
				UnifiedRoleDefinitionId: "unifiedRoleDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/exchange/roleDefinitions/unifiedRoleDefinitionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE/rOlEdEfInItIoNs/uNiFiEdRoLeDeFiNiTiOnIdVaLuE",
			Expected: &RoleManagementExchangeRoleDefinitionId{
				UnifiedRoleDefinitionId: "uNiFiEdRoLeDeFiNiTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE/rOlEdEfInItIoNs/uNiFiEdRoLeDeFiNiTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementExchangeRoleDefinitionIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForRoleManagementExchangeRoleDefinitionId(t *testing.T) {
	segments := RoleManagementExchangeRoleDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementExchangeRoleDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
