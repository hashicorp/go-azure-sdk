package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromId{}

func TestNewRoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromID(t *testing.T) {
	id := NewRoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromID("unifiedRoleDefinitionId", "unifiedRoleDefinitionId1")

	if id.UnifiedRoleDefinitionId != "unifiedRoleDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleDefinitionId'", id.UnifiedRoleDefinitionId, "unifiedRoleDefinitionId")
	}

	if id.UnifiedRoleDefinitionId1 != "unifiedRoleDefinitionId1" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleDefinitionId1'", id.UnifiedRoleDefinitionId1, "unifiedRoleDefinitionId1")
	}
}

func TestFormatRoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromID(t *testing.T) {
	actual := NewRoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromID("unifiedRoleDefinitionId", "unifiedRoleDefinitionId1").ID()
	expected := "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionId/inheritsPermissionsFrom/unifiedRoleDefinitionId1"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromId
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
			Input: "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionId/inheritsPermissionsFrom",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionId/inheritsPermissionsFrom/unifiedRoleDefinitionId1",
			Expected: &RoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromId{
				UnifiedRoleDefinitionId:  "unifiedRoleDefinitionId",
				UnifiedRoleDefinitionId1: "unifiedRoleDefinitionId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionId/inheritsPermissionsFrom/unifiedRoleDefinitionId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromID(v.Input)
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

func TestParseRoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromId
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
			Input: "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEdEfInItIoNs/uNiFiEdRoLeDeFiNiTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionId/inheritsPermissionsFrom",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEdEfInItIoNs/uNiFiEdRoLeDeFiNiTiOnId/iNhErItSpErMiSsIoNsFrOm",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionId/inheritsPermissionsFrom/unifiedRoleDefinitionId1",
			Expected: &RoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromId{
				UnifiedRoleDefinitionId:  "unifiedRoleDefinitionId",
				UnifiedRoleDefinitionId1: "unifiedRoleDefinitionId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/cloudPC/roleDefinitions/unifiedRoleDefinitionId/inheritsPermissionsFrom/unifiedRoleDefinitionId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEdEfInItIoNs/uNiFiEdRoLeDeFiNiTiOnId/iNhErItSpErMiSsIoNsFrOm/uNiFiEdRoLeDeFiNiTiOnId1",
			Expected: &RoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromId{
				UnifiedRoleDefinitionId:  "uNiFiEdRoLeDeFiNiTiOnId",
				UnifiedRoleDefinitionId1: "uNiFiEdRoLeDeFiNiTiOnId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEdEfInItIoNs/uNiFiEdRoLeDeFiNiTiOnId/iNhErItSpErMiSsIoNsFrOm/uNiFiEdRoLeDeFiNiTiOnId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromIDInsensitively(v.Input)
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

func TestSegmentsForRoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromId(t *testing.T) {
	segments := RoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementCloudPCRoleDefinitionIdInheritsPermissionsFromId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
