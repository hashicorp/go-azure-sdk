package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDefenderResourceNamespaceIdResourceActionId{}

func TestNewRoleManagementDefenderResourceNamespaceIdResourceActionID(t *testing.T) {
	id := NewRoleManagementDefenderResourceNamespaceIdResourceActionID("unifiedRbacResourceNamespaceId", "unifiedRbacResourceActionId")

	if id.UnifiedRbacResourceNamespaceId != "unifiedRbacResourceNamespaceId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRbacResourceNamespaceId'", id.UnifiedRbacResourceNamespaceId, "unifiedRbacResourceNamespaceId")
	}

	if id.UnifiedRbacResourceActionId != "unifiedRbacResourceActionId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRbacResourceActionId'", id.UnifiedRbacResourceActionId, "unifiedRbacResourceActionId")
	}
}

func TestFormatRoleManagementDefenderResourceNamespaceIdResourceActionID(t *testing.T) {
	actual := NewRoleManagementDefenderResourceNamespaceIdResourceActionID("unifiedRbacResourceNamespaceId", "unifiedRbacResourceActionId").ID()
	expected := "/roleManagement/defender/resourceNamespaces/unifiedRbacResourceNamespaceId/resourceActions/unifiedRbacResourceActionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementDefenderResourceNamespaceIdResourceActionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDefenderResourceNamespaceIdResourceActionId
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
			Input: "/roleManagement/defender",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/defender/resourceNamespaces",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/defender/resourceNamespaces/unifiedRbacResourceNamespaceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/defender/resourceNamespaces/unifiedRbacResourceNamespaceId/resourceActions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/defender/resourceNamespaces/unifiedRbacResourceNamespaceId/resourceActions/unifiedRbacResourceActionId",
			Expected: &RoleManagementDefenderResourceNamespaceIdResourceActionId{
				UnifiedRbacResourceNamespaceId: "unifiedRbacResourceNamespaceId",
				UnifiedRbacResourceActionId:    "unifiedRbacResourceActionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/defender/resourceNamespaces/unifiedRbacResourceNamespaceId/resourceActions/unifiedRbacResourceActionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDefenderResourceNamespaceIdResourceActionID(v.Input)
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

		if actual.UnifiedRbacResourceActionId != v.Expected.UnifiedRbacResourceActionId {
			t.Fatalf("Expected %q but got %q for UnifiedRbacResourceActionId", v.Expected.UnifiedRbacResourceActionId, actual.UnifiedRbacResourceActionId)
		}

	}
}

func TestParseRoleManagementDefenderResourceNamespaceIdResourceActionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDefenderResourceNamespaceIdResourceActionId
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
			Input: "/roleManagement/defender",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEfEnDeR",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/defender/resourceNamespaces",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEfEnDeR/rEsOuRcEnAmEsPaCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/defender/resourceNamespaces/unifiedRbacResourceNamespaceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEfEnDeR/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/defender/resourceNamespaces/unifiedRbacResourceNamespaceId/resourceActions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEfEnDeR/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiD/rEsOuRcEaCtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/defender/resourceNamespaces/unifiedRbacResourceNamespaceId/resourceActions/unifiedRbacResourceActionId",
			Expected: &RoleManagementDefenderResourceNamespaceIdResourceActionId{
				UnifiedRbacResourceNamespaceId: "unifiedRbacResourceNamespaceId",
				UnifiedRbacResourceActionId:    "unifiedRbacResourceActionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/defender/resourceNamespaces/unifiedRbacResourceNamespaceId/resourceActions/unifiedRbacResourceActionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEfEnDeR/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiD/rEsOuRcEaCtIoNs/uNiFiEdRbAcReSoUrCeAcTiOnId",
			Expected: &RoleManagementDefenderResourceNamespaceIdResourceActionId{
				UnifiedRbacResourceNamespaceId: "uNiFiEdRbAcReSoUrCeNaMeSpAcEiD",
				UnifiedRbacResourceActionId:    "uNiFiEdRbAcReSoUrCeAcTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEfEnDeR/rEsOuRcEnAmEsPaCeS/uNiFiEdRbAcReSoUrCeNaMeSpAcEiD/rEsOuRcEaCtIoNs/uNiFiEdRbAcReSoUrCeAcTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDefenderResourceNamespaceIdResourceActionIDInsensitively(v.Input)
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

		if actual.UnifiedRbacResourceActionId != v.Expected.UnifiedRbacResourceActionId {
			t.Fatalf("Expected %q but got %q for UnifiedRbacResourceActionId", v.Expected.UnifiedRbacResourceActionId, actual.UnifiedRbacResourceActionId)
		}

	}
}

func TestSegmentsForRoleManagementDefenderResourceNamespaceIdResourceActionId(t *testing.T) {
	segments := RoleManagementDefenderResourceNamespaceIdResourceActionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementDefenderResourceNamespaceIdResourceActionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
