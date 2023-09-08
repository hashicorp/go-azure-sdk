package rolemanagementdirectoryroleassignmentscheduleinstanceprincipal

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementDirectoryRoleAssignmentScheduleInstanceId{}

func TestNewRoleManagementDirectoryRoleAssignmentScheduleInstanceID(t *testing.T) {
	id := NewRoleManagementDirectoryRoleAssignmentScheduleInstanceID("unifiedRoleAssignmentScheduleInstanceIdValue")

	if id.UnifiedRoleAssignmentScheduleInstanceId != "unifiedRoleAssignmentScheduleInstanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentScheduleInstanceId'", id.UnifiedRoleAssignmentScheduleInstanceId, "unifiedRoleAssignmentScheduleInstanceIdValue")
	}
}

func TestFormatRoleManagementDirectoryRoleAssignmentScheduleInstanceID(t *testing.T) {
	actual := NewRoleManagementDirectoryRoleAssignmentScheduleInstanceID("unifiedRoleAssignmentScheduleInstanceIdValue").ID()
	expected := "/roleManagement/directory/roleAssignmentScheduleInstances/unifiedRoleAssignmentScheduleInstanceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementDirectoryRoleAssignmentScheduleInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDirectoryRoleAssignmentScheduleInstanceId
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
			Input: "/roleManagement/directory/roleAssignmentScheduleInstances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/directory/roleAssignmentScheduleInstances/unifiedRoleAssignmentScheduleInstanceIdValue",
			Expected: &RoleManagementDirectoryRoleAssignmentScheduleInstanceId{
				UnifiedRoleAssignmentScheduleInstanceId: "unifiedRoleAssignmentScheduleInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/directory/roleAssignmentScheduleInstances/unifiedRoleAssignmentScheduleInstanceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDirectoryRoleAssignmentScheduleInstanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentScheduleInstanceId != v.Expected.UnifiedRoleAssignmentScheduleInstanceId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentScheduleInstanceId", v.Expected.UnifiedRoleAssignmentScheduleInstanceId, actual.UnifiedRoleAssignmentScheduleInstanceId)
		}

	}
}

func TestParseRoleManagementDirectoryRoleAssignmentScheduleInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDirectoryRoleAssignmentScheduleInstanceId
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
			Input: "/roleManagement/directory/roleAssignmentScheduleInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEaSsIgNmEnTsChEdUlEiNsTaNcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/directory/roleAssignmentScheduleInstances/unifiedRoleAssignmentScheduleInstanceIdValue",
			Expected: &RoleManagementDirectoryRoleAssignmentScheduleInstanceId{
				UnifiedRoleAssignmentScheduleInstanceId: "unifiedRoleAssignmentScheduleInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/directory/roleAssignmentScheduleInstances/unifiedRoleAssignmentScheduleInstanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEaSsIgNmEnTsChEdUlEiNsTaNcEs/uNiFiEdRoLeAsSiGnMeNtScHeDuLeInStAnCeIdVaLuE",
			Expected: &RoleManagementDirectoryRoleAssignmentScheduleInstanceId{
				UnifiedRoleAssignmentScheduleInstanceId: "uNiFiEdRoLeAsSiGnMeNtScHeDuLeInStAnCeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEaSsIgNmEnTsChEdUlEiNsTaNcEs/uNiFiEdRoLeAsSiGnMeNtScHeDuLeInStAnCeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDirectoryRoleAssignmentScheduleInstanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentScheduleInstanceId != v.Expected.UnifiedRoleAssignmentScheduleInstanceId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentScheduleInstanceId", v.Expected.UnifiedRoleAssignmentScheduleInstanceId, actual.UnifiedRoleAssignmentScheduleInstanceId)
		}

	}
}

func TestSegmentsForRoleManagementDirectoryRoleAssignmentScheduleInstanceId(t *testing.T) {
	segments := RoleManagementDirectoryRoleAssignmentScheduleInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementDirectoryRoleAssignmentScheduleInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
