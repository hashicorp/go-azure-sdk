package rolescopetag

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRoleScopeTagId{}

func TestNewDeviceManagementRoleScopeTagID(t *testing.T) {
	id := NewDeviceManagementRoleScopeTagID("roleScopeTagIdValue")

	if id.RoleScopeTagId != "roleScopeTagIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleScopeTagId'", id.RoleScopeTagId, "roleScopeTagIdValue")
	}
}

func TestFormatDeviceManagementRoleScopeTagID(t *testing.T) {
	actual := NewDeviceManagementRoleScopeTagID("roleScopeTagIdValue").ID()
	expected := "/deviceManagement/roleScopeTags/roleScopeTagIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementRoleScopeTagID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementRoleScopeTagId
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
			Input: "/deviceManagement/roleScopeTags",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/roleScopeTags/roleScopeTagIdValue",
			Expected: &DeviceManagementRoleScopeTagId{
				RoleScopeTagId: "roleScopeTagIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleScopeTags/roleScopeTagIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementRoleScopeTagID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RoleScopeTagId != v.Expected.RoleScopeTagId {
			t.Fatalf("Expected %q but got %q for RoleScopeTagId", v.Expected.RoleScopeTagId, actual.RoleScopeTagId)
		}

	}
}

func TestParseDeviceManagementRoleScopeTagIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementRoleScopeTagId
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
			Input: "/deviceManagement/roleScopeTags",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEsCoPeTaGs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/roleScopeTags/roleScopeTagIdValue",
			Expected: &DeviceManagementRoleScopeTagId{
				RoleScopeTagId: "roleScopeTagIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleScopeTags/roleScopeTagIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEsCoPeTaGs/rOlEsCoPeTaGiDvAlUe",
			Expected: &DeviceManagementRoleScopeTagId{
				RoleScopeTagId: "rOlEsCoPeTaGiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEsCoPeTaGs/rOlEsCoPeTaGiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementRoleScopeTagIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RoleScopeTagId != v.Expected.RoleScopeTagId {
			t.Fatalf("Expected %q but got %q for RoleScopeTagId", v.Expected.RoleScopeTagId, actual.RoleScopeTagId)
		}

	}
}

func TestSegmentsForDeviceManagementRoleScopeTagId(t *testing.T) {
	segments := DeviceManagementRoleScopeTagId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementRoleScopeTagId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
