package devicemanagementroleassignmentappscope

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDeviceManagementRoleAssignmentIdAppScopeId{}

func TestNewRoleManagementDeviceManagementRoleAssignmentIdAppScopeID(t *testing.T) {
	id := NewRoleManagementDeviceManagementRoleAssignmentIdAppScopeID("unifiedRoleAssignmentMultipleIdValue", "appScopeIdValue")

	if id.UnifiedRoleAssignmentMultipleId != "unifiedRoleAssignmentMultipleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentMultipleId'", id.UnifiedRoleAssignmentMultipleId, "unifiedRoleAssignmentMultipleIdValue")
	}

	if id.AppScopeId != "appScopeIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AppScopeId'", id.AppScopeId, "appScopeIdValue")
	}
}

func TestFormatRoleManagementDeviceManagementRoleAssignmentIdAppScopeID(t *testing.T) {
	actual := NewRoleManagementDeviceManagementRoleAssignmentIdAppScopeID("unifiedRoleAssignmentMultipleIdValue", "appScopeIdValue").ID()
	expected := "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleIdValue/appScopes/appScopeIdValue"
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
			Input: "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleIdValue/appScopes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleIdValue/appScopes/appScopeIdValue",
			Expected: &RoleManagementDeviceManagementRoleAssignmentIdAppScopeId{
				UnifiedRoleAssignmentMultipleId: "unifiedRoleAssignmentMultipleIdValue",
				AppScopeId:                      "appScopeIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleIdValue/appScopes/appScopeIdValue/extra",
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
			Input: "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleIdValue/appScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeIdVaLuE/aPpScOpEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleIdValue/appScopes/appScopeIdValue",
			Expected: &RoleManagementDeviceManagementRoleAssignmentIdAppScopeId{
				UnifiedRoleAssignmentMultipleId: "unifiedRoleAssignmentMultipleIdValue",
				AppScopeId:                      "appScopeIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/deviceManagement/roleAssignments/unifiedRoleAssignmentMultipleIdValue/appScopes/appScopeIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeIdVaLuE/aPpScOpEs/aPpScOpEiDvAlUe",
			Expected: &RoleManagementDeviceManagementRoleAssignmentIdAppScopeId{
				UnifiedRoleAssignmentMultipleId: "uNiFiEdRoLeAsSiGnMeNtMuLtIpLeIdVaLuE",
				AppScopeId:                      "aPpScOpEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEvIcEmAnAgEmEnT/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeIdVaLuE/aPpScOpEs/aPpScOpEiDvAlUe/extra",
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
