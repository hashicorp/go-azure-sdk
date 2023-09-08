package rolemanagementexchangeresourcenamespace

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementExchangeResourceNamespaceId{}

func TestNewRoleManagementExchangeResourceNamespaceID(t *testing.T) {
	id := NewRoleManagementExchangeResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

	if id.UnifiedRbacResourceNamespaceId != "unifiedRbacResourceNamespaceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRbacResourceNamespaceId'", id.UnifiedRbacResourceNamespaceId, "unifiedRbacResourceNamespaceIdValue")
	}
}

func TestFormatRoleManagementExchangeResourceNamespaceID(t *testing.T) {
	actual := NewRoleManagementExchangeResourceNamespaceID("unifiedRbacResourceNamespaceIdValue").ID()
	expected := "/roleManagement/exchange/resourceNamespaces/unifiedRbacResourceNamespaceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementExchangeResourceNamespaceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementExchangeResourceNamespaceId
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
			Input: "/roleManagement/exchange/resourceNamespaces",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/exchange/resourceNamespaces/unifiedRbacResourceNamespaceIdValue",
			Expected: &RoleManagementExchangeResourceNamespaceId{
				UnifiedRbacResourceNamespaceId: "unifiedRbacResourceNamespaceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/exchange/resourceNamespaces/unifiedRbacResourceNamespaceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementExchangeResourceNamespaceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRbacResourceNamespaceId != v.Expected.UnifiedRbacResourceNamespaceId {
			t.Fatalf("Expected %q but got %q for UnifiedRbacResourceNamespaceId", v.Expected.UnifiedRbacResourceNamespaceId, actual.UnifiedRbacResourceNamespaceId)
		}

	}
}

func TestParseRoleManagementExchangeResourceNamespaceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementExchangeResourceNamespaceId
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
			Input: "/roleManagement/exchange/resourceNamespaces",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE/rEsOuRcEnAmEsPaCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/exchange/resourceNamespaces/unifiedRbacResourceNamespaceIdValue",
			Expected: &RoleManagementExchangeResourceNamespaceId{
				UnifiedRbacResourceNamespaceId: "unifiedRbacResourceNamespaceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/exchange/resourceNamespaces/unifiedRbacResourceNamespaceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiDvAlUe",
			Expected: &RoleManagementExchangeResourceNamespaceId{
				UnifiedRbacResourceNamespaceId: "uNiFiEdRbAcReSoUrCeNaMeSpAcEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementExchangeResourceNamespaceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRbacResourceNamespaceId != v.Expected.UnifiedRbacResourceNamespaceId {
			t.Fatalf("Expected %q but got %q for UnifiedRbacResourceNamespaceId", v.Expected.UnifiedRbacResourceNamespaceId, actual.UnifiedRbacResourceNamespaceId)
		}

	}
}

func TestSegmentsForRoleManagementExchangeResourceNamespaceId(t *testing.T) {
	segments := RoleManagementExchangeResourceNamespaceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementExchangeResourceNamespaceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
