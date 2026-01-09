package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryRoleAssignmentScheduleRequestId{}

func TestNewRoleManagementDirectoryRoleAssignmentScheduleRequestID(t *testing.T) {
	id := NewRoleManagementDirectoryRoleAssignmentScheduleRequestID("unifiedRoleAssignmentScheduleRequestId")

	if id.UnifiedRoleAssignmentScheduleRequestId != "unifiedRoleAssignmentScheduleRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentScheduleRequestId'", id.UnifiedRoleAssignmentScheduleRequestId, "unifiedRoleAssignmentScheduleRequestId")
	}
}

func TestFormatRoleManagementDirectoryRoleAssignmentScheduleRequestID(t *testing.T) {
	actual := NewRoleManagementDirectoryRoleAssignmentScheduleRequestID("unifiedRoleAssignmentScheduleRequestId").ID()
	expected := "/roleManagement/directory/roleAssignmentScheduleRequests/unifiedRoleAssignmentScheduleRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementDirectoryRoleAssignmentScheduleRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDirectoryRoleAssignmentScheduleRequestId
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
			Input: "/roleManagement/directory/roleAssignmentScheduleRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/directory/roleAssignmentScheduleRequests/unifiedRoleAssignmentScheduleRequestId",
			Expected: &RoleManagementDirectoryRoleAssignmentScheduleRequestId{
				UnifiedRoleAssignmentScheduleRequestId: "unifiedRoleAssignmentScheduleRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/directory/roleAssignmentScheduleRequests/unifiedRoleAssignmentScheduleRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDirectoryRoleAssignmentScheduleRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentScheduleRequestId != v.Expected.UnifiedRoleAssignmentScheduleRequestId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentScheduleRequestId", v.Expected.UnifiedRoleAssignmentScheduleRequestId, actual.UnifiedRoleAssignmentScheduleRequestId)
		}

	}
}

func TestParseRoleManagementDirectoryRoleAssignmentScheduleRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDirectoryRoleAssignmentScheduleRequestId
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
			Input: "/roleManagement/directory/roleAssignmentScheduleRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEaSsIgNmEnTsChEdUlErEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/directory/roleAssignmentScheduleRequests/unifiedRoleAssignmentScheduleRequestId",
			Expected: &RoleManagementDirectoryRoleAssignmentScheduleRequestId{
				UnifiedRoleAssignmentScheduleRequestId: "unifiedRoleAssignmentScheduleRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/directory/roleAssignmentScheduleRequests/unifiedRoleAssignmentScheduleRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEaSsIgNmEnTsChEdUlErEqUeStS/uNiFiEdRoLeAsSiGnMeNtScHeDuLeReQuEsTiD",
			Expected: &RoleManagementDirectoryRoleAssignmentScheduleRequestId{
				UnifiedRoleAssignmentScheduleRequestId: "uNiFiEdRoLeAsSiGnMeNtScHeDuLeReQuEsTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEaSsIgNmEnTsChEdUlErEqUeStS/uNiFiEdRoLeAsSiGnMeNtScHeDuLeReQuEsTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDirectoryRoleAssignmentScheduleRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentScheduleRequestId != v.Expected.UnifiedRoleAssignmentScheduleRequestId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentScheduleRequestId", v.Expected.UnifiedRoleAssignmentScheduleRequestId, actual.UnifiedRoleAssignmentScheduleRequestId)
		}

	}
}

func TestSegmentsForRoleManagementDirectoryRoleAssignmentScheduleRequestId(t *testing.T) {
	segments := RoleManagementDirectoryRoleAssignmentScheduleRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementDirectoryRoleAssignmentScheduleRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
