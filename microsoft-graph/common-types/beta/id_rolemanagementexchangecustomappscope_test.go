package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementExchangeCustomAppScopeId{}

func TestNewRoleManagementExchangeCustomAppScopeID(t *testing.T) {
	id := NewRoleManagementExchangeCustomAppScopeID("customAppScopeIdValue")

	if id.CustomAppScopeId != "customAppScopeIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomAppScopeId'", id.CustomAppScopeId, "customAppScopeIdValue")
	}
}

func TestFormatRoleManagementExchangeCustomAppScopeID(t *testing.T) {
	actual := NewRoleManagementExchangeCustomAppScopeID("customAppScopeIdValue").ID()
	expected := "/roleManagement/exchange/customAppScopes/customAppScopeIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementExchangeCustomAppScopeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementExchangeCustomAppScopeId
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
			Input: "/roleManagement/exchange/customAppScopes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/exchange/customAppScopes/customAppScopeIdValue",
			Expected: &RoleManagementExchangeCustomAppScopeId{
				CustomAppScopeId: "customAppScopeIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/exchange/customAppScopes/customAppScopeIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementExchangeCustomAppScopeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CustomAppScopeId != v.Expected.CustomAppScopeId {
			t.Fatalf("Expected %q but got %q for CustomAppScopeId", v.Expected.CustomAppScopeId, actual.CustomAppScopeId)
		}

	}
}

func TestParseRoleManagementExchangeCustomAppScopeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementExchangeCustomAppScopeId
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
			Input: "/roleManagement/exchange/customAppScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE/cUsToMaPpScOpEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/exchange/customAppScopes/customAppScopeIdValue",
			Expected: &RoleManagementExchangeCustomAppScopeId{
				CustomAppScopeId: "customAppScopeIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/exchange/customAppScopes/customAppScopeIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE/cUsToMaPpScOpEs/cUsToMaPpScOpEiDvAlUe",
			Expected: &RoleManagementExchangeCustomAppScopeId{
				CustomAppScopeId: "cUsToMaPpScOpEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE/cUsToMaPpScOpEs/cUsToMaPpScOpEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementExchangeCustomAppScopeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CustomAppScopeId != v.Expected.CustomAppScopeId {
			t.Fatalf("Expected %q but got %q for CustomAppScopeId", v.Expected.CustomAppScopeId, actual.CustomAppScopeId)
		}

	}
}

func TestSegmentsForRoleManagementExchangeCustomAppScopeId(t *testing.T) {
	segments := RoleManagementExchangeCustomAppScopeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementExchangeCustomAppScopeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
