package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAndroidForWorkEnrollmentProfileId{}

func TestNewDeviceManagementAndroidForWorkEnrollmentProfileID(t *testing.T) {
	id := NewDeviceManagementAndroidForWorkEnrollmentProfileID("androidForWorkEnrollmentProfileId")

	if id.AndroidForWorkEnrollmentProfileId != "androidForWorkEnrollmentProfileId" {
		t.Fatalf("Expected %q but got %q for Segment 'AndroidForWorkEnrollmentProfileId'", id.AndroidForWorkEnrollmentProfileId, "androidForWorkEnrollmentProfileId")
	}
}

func TestFormatDeviceManagementAndroidForWorkEnrollmentProfileID(t *testing.T) {
	actual := NewDeviceManagementAndroidForWorkEnrollmentProfileID("androidForWorkEnrollmentProfileId").ID()
	expected := "/deviceManagement/androidForWorkEnrollmentProfiles/androidForWorkEnrollmentProfileId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementAndroidForWorkEnrollmentProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAndroidForWorkEnrollmentProfileId
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
			Input: "/deviceManagement/androidForWorkEnrollmentProfiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/androidForWorkEnrollmentProfiles/androidForWorkEnrollmentProfileId",
			Expected: &DeviceManagementAndroidForWorkEnrollmentProfileId{
				AndroidForWorkEnrollmentProfileId: "androidForWorkEnrollmentProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/androidForWorkEnrollmentProfiles/androidForWorkEnrollmentProfileId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAndroidForWorkEnrollmentProfileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AndroidForWorkEnrollmentProfileId != v.Expected.AndroidForWorkEnrollmentProfileId {
			t.Fatalf("Expected %q but got %q for AndroidForWorkEnrollmentProfileId", v.Expected.AndroidForWorkEnrollmentProfileId, actual.AndroidForWorkEnrollmentProfileId)
		}

	}
}

func TestParseDeviceManagementAndroidForWorkEnrollmentProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAndroidForWorkEnrollmentProfileId
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
			Input: "/deviceManagement/androidForWorkEnrollmentProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aNdRoIdFoRwOrKeNrOlLmEnTpRoFiLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/androidForWorkEnrollmentProfiles/androidForWorkEnrollmentProfileId",
			Expected: &DeviceManagementAndroidForWorkEnrollmentProfileId{
				AndroidForWorkEnrollmentProfileId: "androidForWorkEnrollmentProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/androidForWorkEnrollmentProfiles/androidForWorkEnrollmentProfileId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aNdRoIdFoRwOrKeNrOlLmEnTpRoFiLeS/aNdRoIdFoRwOrKeNrOlLmEnTpRoFiLeId",
			Expected: &DeviceManagementAndroidForWorkEnrollmentProfileId{
				AndroidForWorkEnrollmentProfileId: "aNdRoIdFoRwOrKeNrOlLmEnTpRoFiLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aNdRoIdFoRwOrKeNrOlLmEnTpRoFiLeS/aNdRoIdFoRwOrKeNrOlLmEnTpRoFiLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAndroidForWorkEnrollmentProfileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AndroidForWorkEnrollmentProfileId != v.Expected.AndroidForWorkEnrollmentProfileId {
			t.Fatalf("Expected %q but got %q for AndroidForWorkEnrollmentProfileId", v.Expected.AndroidForWorkEnrollmentProfileId, actual.AndroidForWorkEnrollmentProfileId)
		}

	}
}

func TestSegmentsForDeviceManagementAndroidForWorkEnrollmentProfileId(t *testing.T) {
	segments := DeviceManagementAndroidForWorkEnrollmentProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementAndroidForWorkEnrollmentProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
