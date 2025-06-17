package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDefenderRoleAssignmentIdAppScopeId{}

func TestNewRoleManagementDefenderRoleAssignmentIdAppScopeID(t *testing.T) {
	id := NewRoleManagementDefenderRoleAssignmentIdAppScopeID("unifiedRoleAssignmentMultipleId", "appScopeId")

	if id.UnifiedRoleAssignmentMultipleId != "unifiedRoleAssignmentMultipleId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentMultipleId'", id.UnifiedRoleAssignmentMultipleId, "unifiedRoleAssignmentMultipleId")
	}

	if id.AppScopeId != "appScopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'AppScopeId'", id.AppScopeId, "appScopeId")
	}
}

func TestFormatRoleManagementDefenderRoleAssignmentIdAppScopeID(t *testing.T) {
	actual := NewRoleManagementDefenderRoleAssignmentIdAppScopeID("unifiedRoleAssignmentMultipleId", "appScopeId").ID()
	expected := "/roleManagement/defender/roleAssignments/unifiedRoleAssignmentMultipleId/appScopes/appScopeId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementDefenderRoleAssignmentIdAppScopeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDefenderRoleAssignmentIdAppScopeId
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
			Input: "/roleManagement/defender/roleAssignments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/defender/roleAssignments/unifiedRoleAssignmentMultipleId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/defender/roleAssignments/unifiedRoleAssignmentMultipleId/appScopes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/defender/roleAssignments/unifiedRoleAssignmentMultipleId/appScopes/appScopeId",
			Expected: &RoleManagementDefenderRoleAssignmentIdAppScopeId{
				UnifiedRoleAssignmentMultipleId: "unifiedRoleAssignmentMultipleId",
				AppScopeId:                      "appScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/defender/roleAssignments/unifiedRoleAssignmentMultipleId/appScopes/appScopeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDefenderRoleAssignmentIdAppScopeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentMultipleId != v.Expected.UnifiedRoleAssignmentMultipleId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentMultipleId", v.Expected.UnifiedRoleAssignmentMultipleId, actual.UnifiedRoleAssignmentMultipleId)
		}

		if actual.AppScopeId != v.Expected.AppScopeId {
			t.Fatalf("Expected %q but got %q for AppScopeId", v.Expected.AppScopeId, actual.AppScopeId)
		}

	}
}

func TestParseRoleManagementDefenderRoleAssignmentIdAppScopeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDefenderRoleAssignmentIdAppScopeId
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
			Input: "/roleManagement/defender/roleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEfEnDeR/rOlEaSsIgNmEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/defender/roleAssignments/unifiedRoleAssignmentMultipleId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEfEnDeR/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/defender/roleAssignments/unifiedRoleAssignmentMultipleId/appScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEfEnDeR/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId/aPpScOpEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/defender/roleAssignments/unifiedRoleAssignmentMultipleId/appScopes/appScopeId",
			Expected: &RoleManagementDefenderRoleAssignmentIdAppScopeId{
				UnifiedRoleAssignmentMultipleId: "unifiedRoleAssignmentMultipleId",
				AppScopeId:                      "appScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/defender/roleAssignments/unifiedRoleAssignmentMultipleId/appScopes/appScopeId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEfEnDeR/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId/aPpScOpEs/aPpScOpEiD",
			Expected: &RoleManagementDefenderRoleAssignmentIdAppScopeId{
				UnifiedRoleAssignmentMultipleId: "uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId",
				AppScopeId:                      "aPpScOpEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEfEnDeR/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId/aPpScOpEs/aPpScOpEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDefenderRoleAssignmentIdAppScopeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentMultipleId != v.Expected.UnifiedRoleAssignmentMultipleId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentMultipleId", v.Expected.UnifiedRoleAssignmentMultipleId, actual.UnifiedRoleAssignmentMultipleId)
		}

		if actual.AppScopeId != v.Expected.AppScopeId {
			t.Fatalf("Expected %q but got %q for AppScopeId", v.Expected.AppScopeId, actual.AppScopeId)
		}

	}
}

func TestSegmentsForRoleManagementDefenderRoleAssignmentIdAppScopeId(t *testing.T) {
	segments := RoleManagementDefenderRoleAssignmentIdAppScopeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementDefenderRoleAssignmentIdAppScopeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
