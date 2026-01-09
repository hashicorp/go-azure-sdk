package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId{}

func TestNewRoleManagementCloudPCRoleAssignmentIdDirectoryScopeID(t *testing.T) {
	id := NewRoleManagementCloudPCRoleAssignmentIdDirectoryScopeID("unifiedRoleAssignmentMultipleId", "directoryObjectId")

	if id.UnifiedRoleAssignmentMultipleId != "unifiedRoleAssignmentMultipleId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentMultipleId'", id.UnifiedRoleAssignmentMultipleId, "unifiedRoleAssignmentMultipleId")
	}

	if id.DirectoryObjectId != "directoryObjectId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectId")
	}
}

func TestFormatRoleManagementCloudPCRoleAssignmentIdDirectoryScopeID(t *testing.T) {
	actual := NewRoleManagementCloudPCRoleAssignmentIdDirectoryScopeID("unifiedRoleAssignmentMultipleId", "directoryObjectId").ID()
	expected := "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleId/directoryScopes/directoryObjectId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementCloudPCRoleAssignmentIdDirectoryScopeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId
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
			Input: "/roleManagement/cloudPC/roleAssignments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleId/directoryScopes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleId/directoryScopes/directoryObjectId",
			Expected: &RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId{
				UnifiedRoleAssignmentMultipleId: "unifiedRoleAssignmentMultipleId",
				DirectoryObjectId:               "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleId/directoryScopes/directoryObjectId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementCloudPCRoleAssignmentIdDirectoryScopeID(v.Input)
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

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParseRoleManagementCloudPCRoleAssignmentIdDirectoryScopeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId
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
			Input: "/roleManagement/cloudPC/roleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEaSsIgNmEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleId/directoryScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId/dIrEcToRyScOpEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleId/directoryScopes/directoryObjectId",
			Expected: &RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId{
				UnifiedRoleAssignmentMultipleId: "unifiedRoleAssignmentMultipleId",
				DirectoryObjectId:               "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/cloudPC/roleAssignments/unifiedRoleAssignmentMultipleId/directoryScopes/directoryObjectId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId/dIrEcToRyScOpEs/dIrEcToRyObJeCtId",
			Expected: &RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId{
				UnifiedRoleAssignmentMultipleId: "uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId",
				DirectoryObjectId:               "dIrEcToRyObJeCtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/cLoUdPc/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId/dIrEcToRyScOpEs/dIrEcToRyObJeCtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementCloudPCRoleAssignmentIdDirectoryScopeIDInsensitively(v.Input)
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

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForRoleManagementCloudPCRoleAssignmentIdDirectoryScopeId(t *testing.T) {
	segments := RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
