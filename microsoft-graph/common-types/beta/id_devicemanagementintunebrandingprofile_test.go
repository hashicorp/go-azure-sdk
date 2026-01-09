package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntuneBrandingProfileId{}

func TestNewDeviceManagementIntuneBrandingProfileID(t *testing.T) {
	id := NewDeviceManagementIntuneBrandingProfileID("intuneBrandingProfileId")

	if id.IntuneBrandingProfileId != "intuneBrandingProfileId" {
		t.Fatalf("Expected %q but got %q for Segment 'IntuneBrandingProfileId'", id.IntuneBrandingProfileId, "intuneBrandingProfileId")
	}
}

func TestFormatDeviceManagementIntuneBrandingProfileID(t *testing.T) {
	actual := NewDeviceManagementIntuneBrandingProfileID("intuneBrandingProfileId").ID()
	expected := "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementIntuneBrandingProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntuneBrandingProfileId
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
			Input: "/deviceManagement/intuneBrandingProfiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileId",
			Expected: &DeviceManagementIntuneBrandingProfileId{
				IntuneBrandingProfileId: "intuneBrandingProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntuneBrandingProfileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.IntuneBrandingProfileId != v.Expected.IntuneBrandingProfileId {
			t.Fatalf("Expected %q but got %q for IntuneBrandingProfileId", v.Expected.IntuneBrandingProfileId, actual.IntuneBrandingProfileId)
		}

	}
}

func TestParseDeviceManagementIntuneBrandingProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntuneBrandingProfileId
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
			Input: "/deviceManagement/intuneBrandingProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtUnEbRaNdInGpRoFiLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileId",
			Expected: &DeviceManagementIntuneBrandingProfileId{
				IntuneBrandingProfileId: "intuneBrandingProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtUnEbRaNdInGpRoFiLeS/iNtUnEbRaNdInGpRoFiLeId",
			Expected: &DeviceManagementIntuneBrandingProfileId{
				IntuneBrandingProfileId: "iNtUnEbRaNdInGpRoFiLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtUnEbRaNdInGpRoFiLeS/iNtUnEbRaNdInGpRoFiLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntuneBrandingProfileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.IntuneBrandingProfileId != v.Expected.IntuneBrandingProfileId {
			t.Fatalf("Expected %q but got %q for IntuneBrandingProfileId", v.Expected.IntuneBrandingProfileId, actual.IntuneBrandingProfileId)
		}

	}
}

func TestSegmentsForDeviceManagementIntuneBrandingProfileId(t *testing.T) {
	segments := DeviceManagementIntuneBrandingProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementIntuneBrandingProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
