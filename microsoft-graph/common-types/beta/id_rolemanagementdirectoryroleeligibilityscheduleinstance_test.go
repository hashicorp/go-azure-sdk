package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementDirectoryRoleEligibilityScheduleInstanceId{}

func TestNewRoleManagementDirectoryRoleEligibilityScheduleInstanceID(t *testing.T) {
	id := NewRoleManagementDirectoryRoleEligibilityScheduleInstanceID("unifiedRoleEligibilityScheduleInstanceId")

	if id.UnifiedRoleEligibilityScheduleInstanceId != "unifiedRoleEligibilityScheduleInstanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleEligibilityScheduleInstanceId'", id.UnifiedRoleEligibilityScheduleInstanceId, "unifiedRoleEligibilityScheduleInstanceId")
	}
}

func TestFormatRoleManagementDirectoryRoleEligibilityScheduleInstanceID(t *testing.T) {
	actual := NewRoleManagementDirectoryRoleEligibilityScheduleInstanceID("unifiedRoleEligibilityScheduleInstanceId").ID()
	expected := "/roleManagement/directory/roleEligibilityScheduleInstances/unifiedRoleEligibilityScheduleInstanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementDirectoryRoleEligibilityScheduleInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDirectoryRoleEligibilityScheduleInstanceId
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
			Input: "/roleManagement/directory/roleEligibilityScheduleInstances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/directory/roleEligibilityScheduleInstances/unifiedRoleEligibilityScheduleInstanceId",
			Expected: &RoleManagementDirectoryRoleEligibilityScheduleInstanceId{
				UnifiedRoleEligibilityScheduleInstanceId: "unifiedRoleEligibilityScheduleInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/directory/roleEligibilityScheduleInstances/unifiedRoleEligibilityScheduleInstanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDirectoryRoleEligibilityScheduleInstanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleEligibilityScheduleInstanceId != v.Expected.UnifiedRoleEligibilityScheduleInstanceId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleEligibilityScheduleInstanceId", v.Expected.UnifiedRoleEligibilityScheduleInstanceId, actual.UnifiedRoleEligibilityScheduleInstanceId)
		}

	}
}

func TestParseRoleManagementDirectoryRoleEligibilityScheduleInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementDirectoryRoleEligibilityScheduleInstanceId
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
			Input: "/roleManagement/directory/roleEligibilityScheduleInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEeLiGiBiLiTyScHeDuLeInStAnCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/directory/roleEligibilityScheduleInstances/unifiedRoleEligibilityScheduleInstanceId",
			Expected: &RoleManagementDirectoryRoleEligibilityScheduleInstanceId{
				UnifiedRoleEligibilityScheduleInstanceId: "unifiedRoleEligibilityScheduleInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/directory/roleEligibilityScheduleInstances/unifiedRoleEligibilityScheduleInstanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEeLiGiBiLiTyScHeDuLeInStAnCeS/uNiFiEdRoLeElIgIbIlItYsChEdUlEiNsTaNcEiD",
			Expected: &RoleManagementDirectoryRoleEligibilityScheduleInstanceId{
				UnifiedRoleEligibilityScheduleInstanceId: "uNiFiEdRoLeElIgIbIlItYsChEdUlEiNsTaNcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/dIrEcToRy/rOlEeLiGiBiLiTyScHeDuLeInStAnCeS/uNiFiEdRoLeElIgIbIlItYsChEdUlEiNsTaNcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementDirectoryRoleEligibilityScheduleInstanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleEligibilityScheduleInstanceId != v.Expected.UnifiedRoleEligibilityScheduleInstanceId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleEligibilityScheduleInstanceId", v.Expected.UnifiedRoleEligibilityScheduleInstanceId, actual.UnifiedRoleEligibilityScheduleInstanceId)
		}

	}
}

func TestSegmentsForRoleManagementDirectoryRoleEligibilityScheduleInstanceId(t *testing.T) {
	segments := RoleManagementDirectoryRoleEligibilityScheduleInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementDirectoryRoleEligibilityScheduleInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
