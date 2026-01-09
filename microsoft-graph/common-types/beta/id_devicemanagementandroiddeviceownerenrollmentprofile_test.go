package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAndroidDeviceOwnerEnrollmentProfileId{}

func TestNewDeviceManagementAndroidDeviceOwnerEnrollmentProfileID(t *testing.T) {
	id := NewDeviceManagementAndroidDeviceOwnerEnrollmentProfileID("androidDeviceOwnerEnrollmentProfileId")

	if id.AndroidDeviceOwnerEnrollmentProfileId != "androidDeviceOwnerEnrollmentProfileId" {
		t.Fatalf("Expected %q but got %q for Segment 'AndroidDeviceOwnerEnrollmentProfileId'", id.AndroidDeviceOwnerEnrollmentProfileId, "androidDeviceOwnerEnrollmentProfileId")
	}
}

func TestFormatDeviceManagementAndroidDeviceOwnerEnrollmentProfileID(t *testing.T) {
	actual := NewDeviceManagementAndroidDeviceOwnerEnrollmentProfileID("androidDeviceOwnerEnrollmentProfileId").ID()
	expected := "/deviceManagement/androidDeviceOwnerEnrollmentProfiles/androidDeviceOwnerEnrollmentProfileId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementAndroidDeviceOwnerEnrollmentProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAndroidDeviceOwnerEnrollmentProfileId
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
			Input: "/deviceManagement/androidDeviceOwnerEnrollmentProfiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/androidDeviceOwnerEnrollmentProfiles/androidDeviceOwnerEnrollmentProfileId",
			Expected: &DeviceManagementAndroidDeviceOwnerEnrollmentProfileId{
				AndroidDeviceOwnerEnrollmentProfileId: "androidDeviceOwnerEnrollmentProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/androidDeviceOwnerEnrollmentProfiles/androidDeviceOwnerEnrollmentProfileId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAndroidDeviceOwnerEnrollmentProfileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AndroidDeviceOwnerEnrollmentProfileId != v.Expected.AndroidDeviceOwnerEnrollmentProfileId {
			t.Fatalf("Expected %q but got %q for AndroidDeviceOwnerEnrollmentProfileId", v.Expected.AndroidDeviceOwnerEnrollmentProfileId, actual.AndroidDeviceOwnerEnrollmentProfileId)
		}

	}
}

func TestParseDeviceManagementAndroidDeviceOwnerEnrollmentProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAndroidDeviceOwnerEnrollmentProfileId
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
			Input: "/deviceManagement/androidDeviceOwnerEnrollmentProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aNdRoIdDeViCeOwNeReNrOlLmEnTpRoFiLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/androidDeviceOwnerEnrollmentProfiles/androidDeviceOwnerEnrollmentProfileId",
			Expected: &DeviceManagementAndroidDeviceOwnerEnrollmentProfileId{
				AndroidDeviceOwnerEnrollmentProfileId: "androidDeviceOwnerEnrollmentProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/androidDeviceOwnerEnrollmentProfiles/androidDeviceOwnerEnrollmentProfileId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aNdRoIdDeViCeOwNeReNrOlLmEnTpRoFiLeS/aNdRoIdDeViCeOwNeReNrOlLmEnTpRoFiLeId",
			Expected: &DeviceManagementAndroidDeviceOwnerEnrollmentProfileId{
				AndroidDeviceOwnerEnrollmentProfileId: "aNdRoIdDeViCeOwNeReNrOlLmEnTpRoFiLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aNdRoIdDeViCeOwNeReNrOlLmEnTpRoFiLeS/aNdRoIdDeViCeOwNeReNrOlLmEnTpRoFiLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAndroidDeviceOwnerEnrollmentProfileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AndroidDeviceOwnerEnrollmentProfileId != v.Expected.AndroidDeviceOwnerEnrollmentProfileId {
			t.Fatalf("Expected %q but got %q for AndroidDeviceOwnerEnrollmentProfileId", v.Expected.AndroidDeviceOwnerEnrollmentProfileId, actual.AndroidDeviceOwnerEnrollmentProfileId)
		}

	}
}

func TestSegmentsForDeviceManagementAndroidDeviceOwnerEnrollmentProfileId(t *testing.T) {
	segments := DeviceManagementAndroidDeviceOwnerEnrollmentProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementAndroidDeviceOwnerEnrollmentProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
