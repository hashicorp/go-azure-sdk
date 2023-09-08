package rolemanagementcloudpcroledefinitioninheritspermissionsfrom

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId{}

func TestNewRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromID(t *testing.T) {
	id := NewRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromID("unifiedRoleDefinitionIdValue", "unifiedRoleDefinitionId1Value")

	if id.UnifiedRoleDefinitionId != "unifiedRoleDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleDefinitionId'", id.UnifiedRoleDefinitionId, "unifiedRoleDefinitionIdValue")
	}

	if id.UnifiedRoleDefinitionId1 != "unifiedRoleDefinitionId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleDefinitionId1'", id.UnifiedRoleDefinitionId1, "unifiedRoleDefinitionId1Value")
	}
}

func TestFormatRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromID(t *testing.T) {
	actual := NewRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromID("unifiedRoleDefinitionIdValue", "unifiedRoleDefinitionId1Value").ID()
	expected := "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionIdValue/inheritsPermissionsFrom/unifiedRoleDefinitionId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId
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
			Input: "/roleManagement/cloudPC",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleDefinitions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionIdValue/inheritsPermissionsFrom",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionIdValue/inheritsPermissionsFrom/unifiedRoleDefinitionId1Value",
			Expected: &RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId{
				UnifiedRoleDefinitionId:  "unifiedRoleDefinitionIdValue",
				UnifiedRoleDefinitionId1: "unifiedRoleDefinitionId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionIdValue/inheritsPermissionsFrom/unifiedRoleDefinitionId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromID(v.Input)
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

func TestParseRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId
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
			Input: "/roleManagement/cloudPC",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEdEfInItIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEdEfInItIoNs/uNiFiEdRoLeDeFiNiTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionIdValue/inheritsPermissionsFrom",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEdEfInItIoNs/uNiFiEdRoLeDeFiNiTiOnIdVaLuE/iNhErItSpErMiSsIoNsFrOm",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionIdValue/inheritsPermissionsFrom/unifiedRoleDefinitionId1Value",
			Expected: &RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId{
				UnifiedRoleDefinitionId:  "unifiedRoleDefinitionIdValue",
				UnifiedRoleDefinitionId1: "unifiedRoleDefinitionId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionIdValue/inheritsPermissionsFrom/unifiedRoleDefinitionId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEdEfInItIoNs/uNiFiEdRoLeDeFiNiTiOnIdVaLuE/iNhErItSpErMiSsIoNsFrOm/uNiFiEdRoLeDeFiNiTiOnId1vAlUe",
			Expected: &RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId{
				UnifiedRoleDefinitionId:  "uNiFiEdRoLeDeFiNiTiOnIdVaLuE",
				UnifiedRoleDefinitionId1: "uNiFiEdRoLeDeFiNiTiOnId1vAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEdEfInItIoNs/uNiFiEdRoLeDeFiNiTiOnIdVaLuE/iNhErItSpErMiSsIoNsFrOm/uNiFiEdRoLeDeFiNiTiOnId1vAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromIDInsensitively(v.Input)
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

func TestSegmentsForRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId(t *testing.T) {
	segments := RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
