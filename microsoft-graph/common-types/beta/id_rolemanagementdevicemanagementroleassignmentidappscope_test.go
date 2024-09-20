package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDeviceManagementRoleAssignmentIdAppScopeId{}

func TestNewRoleManagementDeviceManagementRoleAssignmentIdAppScopeID(t *testing.T) {
	id := NewRoleManagementDeviceManagementRoleAssignmentIdAppScopeID("unifiedRoleAssignmentMultipleId", "appScopeId")

	if id.UnifiedRoleAssignmentMultipleId != "unifiedRoleAssignmentMultipleId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentMultipleId'", id.UnifiedRoleAssignmentMultipleId, "unifiedRoleAssignmentMultipleId")
	}

	if id.AppScopeId != "appScopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'AppScopeId'", id.AppScopeId, "appScopeId")
	}
}

func TestFormatRoleManagementDeviceManagementRoleAssignmentIdAppScopeID(t *testing.T) {
	actual := NewRoleManagementDeviceManagementRoleAssignmentIdAppScopeID("unifiedRoleAssignmentMultipleId", "appScopeId").ID()
	expected := "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleId/appScopes/appScopeId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementDeviceManagementRoleAssignmentIdAppScopeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDeviceManagementRoleAssignmentIdAppScopeId
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
			Input: "/roleManagement/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/deviceManagement/roleAssignments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleId/appScopes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleId/appScopes/appScopeId",
			Expected: &RoleManagementDeviceManagementRoleAssignmentIdAppScopeId{
				UnifiedRoleAssignmentMultipleId: "unifiedRoleAssignmentMultipleId",
				AppScopeId:                      "appScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleId/appScopes/appScopeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDeviceManagementRoleAssignmentIdAppScopeID(v.Input)
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

func TestParseRoleManagementDeviceManagementRoleAssignmentIdAppScopeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDeviceManagementRoleAssignmentIdAppScopeId
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
			Input: "/roleManagement/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEvIcEmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/deviceManagement/roleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleId/appScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId/aPpScOpEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleId/appScopes/appScopeId",
			Expected: &RoleManagementDeviceManagementRoleAssignmentIdAppScopeId{
				UnifiedRoleAssignmentMultipleId: "unifiedRoleAssignmentMultipleId",
				AppScopeId:                      "appScopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleId/appScopes/appScopeId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId/aPpScOpEs/aPpScOpEiD",
			Expected: &RoleManagementDeviceManagementRoleAssignmentIdAppScopeId{
				UnifiedRoleAssignmentMultipleId: "uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId",
				AppScopeId:                      "aPpScOpEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId/aPpScOpEs/aPpScOpEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDeviceManagementRoleAssignmentIdAppScopeIDInsensitively(v.Input)
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

func TestSegmentsForRoleManagementDeviceManagementRoleAssignmentIdAppScopeId(t *testing.T) {
	segments := RoleManagementDeviceManagementRoleAssignmentIdAppScopeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementDeviceManagementRoleAssignmentIdAppScopeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
