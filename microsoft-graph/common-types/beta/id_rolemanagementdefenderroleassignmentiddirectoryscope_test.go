package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDefenderRoleAssignmentIdDirectoryScopeId{}

func TestNewRoleManagementDefenderRoleAssignmentIdDirectoryScopeID(t *testing.T) {
	id := NewRoleManagementDefenderRoleAssignmentIdDirectoryScopeID("unifiedRoleAssignmentMultipleId", "directoryObjectId")

	if id.UnifiedRoleAssignmentMultipleId != "unifiedRoleAssignmentMultipleId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentMultipleId'", id.UnifiedRoleAssignmentMultipleId, "unifiedRoleAssignmentMultipleId")
	}

	if id.DirectoryObjectId != "directoryObjectId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectId")
	}
}

func TestFormatRoleManagementDefenderRoleAssignmentIdDirectoryScopeID(t *testing.T) {
	actual := NewRoleManagementDefenderRoleAssignmentIdDirectoryScopeID("unifiedRoleAssignmentMultipleId", "directoryObjectId").ID()
	expected := "/roleManagement/defender/roleAssignments/unifiedRoleAssignmentMultipleId/directoryScopes/directoryObjectId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementDefenderRoleAssignmentIdDirectoryScopeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDefenderRoleAssignmentIdDirectoryScopeId
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
			Input: "/roleManagement/defender/roleAssignments/unifiedRoleAssignmentMultipleId/directoryScopes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/defender/roleAssignments/unifiedRoleAssignmentMultipleId/directoryScopes/directoryObjectId",
			Expected: &RoleManagementDefenderRoleAssignmentIdDirectoryScopeId{
				UnifiedRoleAssignmentMultipleId: "unifiedRoleAssignmentMultipleId",
				DirectoryObjectId:               "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/defender/roleAssignments/unifiedRoleAssignmentMultipleId/directoryScopes/directoryObjectId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDefenderRoleAssignmentIdDirectoryScopeID(v.Input)
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

func TestParseRoleManagementDefenderRoleAssignmentIdDirectoryScopeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDefenderRoleAssignmentIdDirectoryScopeId
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
			Input: "/roleManagement/defender/roleAssignments/unifiedRoleAssignmentMultipleId/directoryScopes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEfEnDeR/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId/dIrEcToRyScOpEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/defender/roleAssignments/unifiedRoleAssignmentMultipleId/directoryScopes/directoryObjectId",
			Expected: &RoleManagementDefenderRoleAssignmentIdDirectoryScopeId{
				UnifiedRoleAssignmentMultipleId: "unifiedRoleAssignmentMultipleId",
				DirectoryObjectId:               "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/defender/roleAssignments/unifiedRoleAssignmentMultipleId/directoryScopes/directoryObjectId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEfEnDeR/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId/dIrEcToRyScOpEs/dIrEcToRyObJeCtId",
			Expected: &RoleManagementDefenderRoleAssignmentIdDirectoryScopeId{
				UnifiedRoleAssignmentMultipleId: "uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId",
				DirectoryObjectId:               "dIrEcToRyObJeCtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dEfEnDeR/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtMuLtIpLeId/dIrEcToRyScOpEs/dIrEcToRyObJeCtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDefenderRoleAssignmentIdDirectoryScopeIDInsensitively(v.Input)
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

func TestSegmentsForRoleManagementDefenderRoleAssignmentIdDirectoryScopeId(t *testing.T) {
	segments := RoleManagementDefenderRoleAssignmentIdDirectoryScopeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementDefenderRoleAssignmentIdDirectoryScopeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
