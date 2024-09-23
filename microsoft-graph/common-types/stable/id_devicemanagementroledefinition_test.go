package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRoleDefinitionId{}

func TestNewDeviceManagementRoleDefinitionID(t *testing.T) {
	id := NewDeviceManagementRoleDefinitionID("roleDefinitionId")

	if id.RoleDefinitionId != "roleDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleDefinitionId'", id.RoleDefinitionId, "roleDefinitionId")
	}
}

func TestFormatDeviceManagementRoleDefinitionID(t *testing.T) {
	actual := NewDeviceManagementRoleDefinitionID("roleDefinitionId").ID()
	expected := "/deviceManagement/roleDefinitions/roleDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementRoleDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementRoleDefinitionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/roleDefinitions/roleDefinitionId",
			Expected: &DeviceManagementRoleDefinitionId{
				RoleDefinitionId: "roleDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleDefinitions/roleDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementRoleDefinitionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RoleDefinitionId != v.Expected.RoleDefinitionId {
			t.Fatalf("Expected %q but got %q for RoleDefinitionId", v.Expected.RoleDefinitionId, actual.RoleDefinitionId)
		}

	}
}

func TestParseDeviceManagementRoleDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementRoleDefinitionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEdEfInItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/roleDefinitions/roleDefinitionId",
			Expected: &DeviceManagementRoleDefinitionId{
				RoleDefinitionId: "roleDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleDefinitions/roleDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEdEfInItIoNs/rOlEdEfInItIoNiD",
			Expected: &DeviceManagementRoleDefinitionId{
				RoleDefinitionId: "rOlEdEfInItIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEdEfInItIoNs/rOlEdEfInItIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementRoleDefinitionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RoleDefinitionId != v.Expected.RoleDefinitionId {
			t.Fatalf("Expected %q but got %q for RoleDefinitionId", v.Expected.RoleDefinitionId, actual.RoleDefinitionId)
		}

	}
}

func TestSegmentsForDeviceManagementRoleDefinitionId(t *testing.T) {
	segments := DeviceManagementRoleDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementRoleDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
