package directoryroleassignmentprincipal

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryRoleAssignmentId{}

func TestNewRoleManagementDirectoryRoleAssignmentID(t *testing.T) {
	id := NewRoleManagementDirectoryRoleAssignmentID("unifiedRoleAssignmentIdValue")

	if id.UnifiedRoleAssignmentId != "unifiedRoleAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentId'", id.UnifiedRoleAssignmentId, "unifiedRoleAssignmentIdValue")
	}
}

func TestFormatRoleManagementDirectoryRoleAssignmentID(t *testing.T) {
	actual := NewRoleManagementDirectoryRoleAssignmentID("unifiedRoleAssignmentIdValue").ID()
	expected := "/roleManagement/directory/roleAssignments/unifiedRoleAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementDirectoryRoleAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDirectoryRoleAssignmentId
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
			Input: "/roleManagement/directory",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/directory/roleAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/directory/roleAssignments/unifiedRoleAssignmentIdValue",
			Expected: &RoleManagementDirectoryRoleAssignmentId{
				UnifiedRoleAssignmentId: "unifiedRoleAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/directory/roleAssignments/unifiedRoleAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDirectoryRoleAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentId != v.Expected.UnifiedRoleAssignmentId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentId", v.Expected.UnifiedRoleAssignmentId, actual.UnifiedRoleAssignmentId)
		}

	}
}

func TestParseRoleManagementDirectoryRoleAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDirectoryRoleAssignmentId
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
			Input: "/roleManagement/directory",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/directory/roleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEaSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/directory/roleAssignments/unifiedRoleAssignmentIdValue",
			Expected: &RoleManagementDirectoryRoleAssignmentId{
				UnifiedRoleAssignmentId: "unifiedRoleAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/directory/roleAssignments/unifiedRoleAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtIdVaLuE",
			Expected: &RoleManagementDirectoryRoleAssignmentId{
				UnifiedRoleAssignmentId: "uNiFiEdRoLeAsSiGnMeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDirectoryRoleAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentId != v.Expected.UnifiedRoleAssignmentId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentId", v.Expected.UnifiedRoleAssignmentId, actual.UnifiedRoleAssignmentId)
		}

	}
}

func TestSegmentsForRoleManagementDirectoryRoleAssignmentId(t *testing.T) {
	segments := RoleManagementDirectoryRoleAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementDirectoryRoleAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
