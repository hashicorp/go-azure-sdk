package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryTransitiveRoleAssignmentId{}

func TestNewRoleManagementDirectoryTransitiveRoleAssignmentID(t *testing.T) {
	id := NewRoleManagementDirectoryTransitiveRoleAssignmentID("unifiedRoleAssignmentId")

	if id.UnifiedRoleAssignmentId != "unifiedRoleAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentId'", id.UnifiedRoleAssignmentId, "unifiedRoleAssignmentId")
	}
}

func TestFormatRoleManagementDirectoryTransitiveRoleAssignmentID(t *testing.T) {
	actual := NewRoleManagementDirectoryTransitiveRoleAssignmentID("unifiedRoleAssignmentId").ID()
	expected := "/roleManagement/directory/transitiveRoleAssignments/unifiedRoleAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementDirectoryTransitiveRoleAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDirectoryTransitiveRoleAssignmentId
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
			Input: "/roleManagement/directory/transitiveRoleAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/directory/transitiveRoleAssignments/unifiedRoleAssignmentId",
			Expected: &RoleManagementDirectoryTransitiveRoleAssignmentId{
				UnifiedRoleAssignmentId: "unifiedRoleAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/directory/transitiveRoleAssignments/unifiedRoleAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDirectoryTransitiveRoleAssignmentID(v.Input)
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

func TestParseRoleManagementDirectoryTransitiveRoleAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDirectoryTransitiveRoleAssignmentId
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
			Input: "/roleManagement/directory/transitiveRoleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/tRaNsItIvErOlEaSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/directory/transitiveRoleAssignments/unifiedRoleAssignmentId",
			Expected: &RoleManagementDirectoryTransitiveRoleAssignmentId{
				UnifiedRoleAssignmentId: "unifiedRoleAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/directory/transitiveRoleAssignments/unifiedRoleAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/tRaNsItIvErOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtId",
			Expected: &RoleManagementDirectoryTransitiveRoleAssignmentId{
				UnifiedRoleAssignmentId: "uNiFiEdRoLeAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/tRaNsItIvErOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDirectoryTransitiveRoleAssignmentIDInsensitively(v.Input)
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

func TestSegmentsForRoleManagementDirectoryTransitiveRoleAssignmentId(t *testing.T) {
	segments := RoleManagementDirectoryTransitiveRoleAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementDirectoryTransitiveRoleAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
