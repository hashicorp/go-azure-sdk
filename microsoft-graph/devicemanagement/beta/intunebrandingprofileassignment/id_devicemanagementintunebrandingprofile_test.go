package intunebrandingprofileassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntuneBrandingProfileId{}

func TestNewDeviceManagementIntuneBrandingProfileID(t *testing.T) {
	id := NewDeviceManagementIntuneBrandingProfileID("intuneBrandingProfileIdValue")

	if id.IntuneBrandingProfileId != "intuneBrandingProfileIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'IntuneBrandingProfileId'", id.IntuneBrandingProfileId, "intuneBrandingProfileIdValue")
	}
}

func TestFormatDeviceManagementIntuneBrandingProfileID(t *testing.T) {
	actual := NewDeviceManagementIntuneBrandingProfileID("intuneBrandingProfileIdValue").ID()
	expected := "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileIdValue"
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
			Input: "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileIdValue",
			Expected: &DeviceManagementIntuneBrandingProfileId{
				IntuneBrandingProfileId: "intuneBrandingProfileIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileIdValue/extra",
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
			Input: "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileIdValue",
			Expected: &DeviceManagementIntuneBrandingProfileId{
				IntuneBrandingProfileId: "intuneBrandingProfileIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtUnEbRaNdInGpRoFiLeS/iNtUnEbRaNdInGpRoFiLeIdVaLuE",
			Expected: &DeviceManagementIntuneBrandingProfileId{
				IntuneBrandingProfileId: "iNtUnEbRaNdInGpRoFiLeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtUnEbRaNdInGpRoFiLeS/iNtUnEbRaNdInGpRoFiLeIdVaLuE/extra",
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
