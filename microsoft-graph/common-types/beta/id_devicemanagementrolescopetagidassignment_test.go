package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementRoleScopeTagIdAssignmentId{}

func TestNewDeviceManagementRoleScopeTagIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementRoleScopeTagIdAssignmentID("roleScopeTagId", "roleScopeTagAutoAssignmentId")

	if id.RoleScopeTagId != "roleScopeTagId" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleScopeTagId'", id.RoleScopeTagId, "roleScopeTagId")
	}

	if id.RoleScopeTagAutoAssignmentId != "roleScopeTagAutoAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleScopeTagAutoAssignmentId'", id.RoleScopeTagAutoAssignmentId, "roleScopeTagAutoAssignmentId")
	}
}

func TestFormatDeviceManagementRoleScopeTagIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementRoleScopeTagIdAssignmentID("roleScopeTagId", "roleScopeTagAutoAssignmentId").ID()
	expected := "/deviceManagement/roleScopeTags/roleScopeTagId/assignments/roleScopeTagAutoAssignmentId"
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
			Input: "/deviceManagement/roleScopeTags/roleScopeTagId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleScopeTags/roleScopeTagId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/roleScopeTags/roleScopeTagId/assignments/roleScopeTagAutoAssignmentId",
			Expected: &DeviceManagementRoleScopeTagIdAssignmentId{
				RoleScopeTagId:               "roleScopeTagId",
				RoleScopeTagAutoAssignmentId: "roleScopeTagAutoAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleScopeTags/roleScopeTagId/assignments/roleScopeTagAutoAssignmentId/extra",
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
			Input: "/deviceManagement/roleScopeTags/roleScopeTagId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEsCoPeTaGs/rOlEsCoPeTaGiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/roleScopeTags/roleScopeTagId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEsCoPeTaGs/rOlEsCoPeTaGiD/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/roleScopeTags/roleScopeTagId/assignments/roleScopeTagAutoAssignmentId",
			Expected: &DeviceManagementRoleScopeTagIdAssignmentId{
				RoleScopeTagId:               "roleScopeTagId",
				RoleScopeTagAutoAssignmentId: "roleScopeTagAutoAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/roleScopeTags/roleScopeTagId/assignments/roleScopeTagAutoAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEsCoPeTaGs/rOlEsCoPeTaGiD/aSsIgNmEnTs/rOlEsCoPeTaGaUtOaSsIgNmEnTiD",
			Expected: &DeviceManagementRoleScopeTagIdAssignmentId{
				RoleScopeTagId:               "rOlEsCoPeTaGiD",
				RoleScopeTagAutoAssignmentId: "rOlEsCoPeTaGaUtOaSsIgNmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rOlEsCoPeTaGs/rOlEsCoPeTaGiD/aSsIgNmEnTs/rOlEsCoPeTaGaUtOaSsIgNmEnTiD/extra",
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
