package rolescopetagassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRoleScopeTagIdAssignmentId{}

func TestNewDeviceManagementRoleScopeTagIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementRoleScopeTagIdAssignmentID("roleScopeTagIdValue", "roleScopeTagAutoAssignmentIdValue")

	if id.RoleScopeTagId != "roleScopeTagIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleScopeTagId'", id.RoleScopeTagId, "roleScopeTagIdValue")
	}

	if id.RoleScopeTagAutoAssignmentId != "roleScopeTagAutoAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleScopeTagAutoAssignmentId'", id.RoleScopeTagAutoAssignmentId, "roleScopeTagAutoAssignmentIdValue")
	}
}

func TestFormatDeviceManagementRoleScopeTagIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementRoleScopeTagIdAssignmentID("roleScopeTagIdValue", "roleScopeTagAutoAssignmentIdValue").ID()
	expected := "/deviceManagement/roleScopeTags/roleScopeTagIdValue/assignments/roleScopeTagAutoAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementRoleScopeTagIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementRoleScopeTagIdAssignmentId
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
			// Incomplete URI
			Input: "/deviceManagement/roleScopeTags/roleScopeTagIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleScopeTags/roleScopeTagIdValue/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/roleScopeTags/roleScopeTagIdValue/assignments/roleScopeTagAutoAssignmentIdValue",
			Expected: &DeviceManagementRoleScopeTagIdAssignmentId{
				RoleScopeTagId:               "roleScopeTagIdValue",
				RoleScopeTagAutoAssignmentId: "roleScopeTagAutoAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleScopeTags/roleScopeTagIdValue/assignments/roleScopeTagAutoAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementRoleScopeTagIdAssignmentID(v.Input)
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

		if actual.RoleScopeTagAutoAssignmentId != v.Expected.RoleScopeTagAutoAssignmentId {
			t.Fatalf("Expected %q but got %q for RoleScopeTagAutoAssignmentId", v.Expected.RoleScopeTagAutoAssignmentId, actual.RoleScopeTagAutoAssignmentId)
		}

	}
}

func TestParseDeviceManagementRoleScopeTagIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementRoleScopeTagIdAssignmentId
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
			// Incomplete URI
			Input: "/deviceManagement/roleScopeTags/roleScopeTagIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEsCoPeTaGs/rOlEsCoPeTaGiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleScopeTags/roleScopeTagIdValue/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEsCoPeTaGs/rOlEsCoPeTaGiDvAlUe/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/roleScopeTags/roleScopeTagIdValue/assignments/roleScopeTagAutoAssignmentIdValue",
			Expected: &DeviceManagementRoleScopeTagIdAssignmentId{
				RoleScopeTagId:               "roleScopeTagIdValue",
				RoleScopeTagAutoAssignmentId: "roleScopeTagAutoAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleScopeTags/roleScopeTagIdValue/assignments/roleScopeTagAutoAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEsCoPeTaGs/rOlEsCoPeTaGiDvAlUe/aSsIgNmEnTs/rOlEsCoPeTaGaUtOaSsIgNmEnTiDvAlUe",
			Expected: &DeviceManagementRoleScopeTagIdAssignmentId{
				RoleScopeTagId:               "rOlEsCoPeTaGiDvAlUe",
				RoleScopeTagAutoAssignmentId: "rOlEsCoPeTaGaUtOaSsIgNmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEsCoPeTaGs/rOlEsCoPeTaGiDvAlUe/aSsIgNmEnTs/rOlEsCoPeTaGaUtOaSsIgNmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementRoleScopeTagIdAssignmentIDInsensitively(v.Input)
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

		if actual.RoleScopeTagAutoAssignmentId != v.Expected.RoleScopeTagAutoAssignmentId {
			t.Fatalf("Expected %q but got %q for RoleScopeTagAutoAssignmentId", v.Expected.RoleScopeTagAutoAssignmentId, actual.RoleScopeTagAutoAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementRoleScopeTagIdAssignmentId(t *testing.T) {
	segments := DeviceManagementRoleScopeTagIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementRoleScopeTagIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
