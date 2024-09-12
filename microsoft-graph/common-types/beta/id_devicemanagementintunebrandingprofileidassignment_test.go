package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntuneBrandingProfileIdAssignmentId{}

func TestNewDeviceManagementIntuneBrandingProfileIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementIntuneBrandingProfileIdAssignmentID("intuneBrandingProfileIdValue", "intuneBrandingProfileAssignmentIdValue")

	if id.IntuneBrandingProfileId != "intuneBrandingProfileIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'IntuneBrandingProfileId'", id.IntuneBrandingProfileId, "intuneBrandingProfileIdValue")
	}

	if id.IntuneBrandingProfileAssignmentId != "intuneBrandingProfileAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'IntuneBrandingProfileAssignmentId'", id.IntuneBrandingProfileAssignmentId, "intuneBrandingProfileAssignmentIdValue")
	}
}

func TestFormatDeviceManagementIntuneBrandingProfileIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementIntuneBrandingProfileIdAssignmentID("intuneBrandingProfileIdValue", "intuneBrandingProfileAssignmentIdValue").ID()
	expected := "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileIdValue/assignments/intuneBrandingProfileAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementIntuneBrandingProfileIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntuneBrandingProfileIdAssignmentId
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
			// Incomplete URI
			Input: "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileIdValue/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileIdValue/assignments/intuneBrandingProfileAssignmentIdValue",
			Expected: &DeviceManagementIntuneBrandingProfileIdAssignmentId{
				IntuneBrandingProfileId:           "intuneBrandingProfileIdValue",
				IntuneBrandingProfileAssignmentId: "intuneBrandingProfileAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileIdValue/assignments/intuneBrandingProfileAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntuneBrandingProfileIdAssignmentID(v.Input)
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

		if actual.IntuneBrandingProfileAssignmentId != v.Expected.IntuneBrandingProfileAssignmentId {
			t.Fatalf("Expected %q but got %q for IntuneBrandingProfileAssignmentId", v.Expected.IntuneBrandingProfileAssignmentId, actual.IntuneBrandingProfileAssignmentId)
		}

	}
}

func TestParseDeviceManagementIntuneBrandingProfileIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntuneBrandingProfileIdAssignmentId
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
			// Incomplete URI
			Input: "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtUnEbRaNdInGpRoFiLeS/iNtUnEbRaNdInGpRoFiLeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileIdValue/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtUnEbRaNdInGpRoFiLeS/iNtUnEbRaNdInGpRoFiLeIdVaLuE/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileIdValue/assignments/intuneBrandingProfileAssignmentIdValue",
			Expected: &DeviceManagementIntuneBrandingProfileIdAssignmentId{
				IntuneBrandingProfileId:           "intuneBrandingProfileIdValue",
				IntuneBrandingProfileAssignmentId: "intuneBrandingProfileAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intuneBrandingProfiles/intuneBrandingProfileIdValue/assignments/intuneBrandingProfileAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtUnEbRaNdInGpRoFiLeS/iNtUnEbRaNdInGpRoFiLeIdVaLuE/aSsIgNmEnTs/iNtUnEbRaNdInGpRoFiLeAsSiGnMeNtIdVaLuE",
			Expected: &DeviceManagementIntuneBrandingProfileIdAssignmentId{
				IntuneBrandingProfileId:           "iNtUnEbRaNdInGpRoFiLeIdVaLuE",
				IntuneBrandingProfileAssignmentId: "iNtUnEbRaNdInGpRoFiLeAsSiGnMeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtUnEbRaNdInGpRoFiLeS/iNtUnEbRaNdInGpRoFiLeIdVaLuE/aSsIgNmEnTs/iNtUnEbRaNdInGpRoFiLeAsSiGnMeNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntuneBrandingProfileIdAssignmentIDInsensitively(v.Input)
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

		if actual.IntuneBrandingProfileAssignmentId != v.Expected.IntuneBrandingProfileAssignmentId {
			t.Fatalf("Expected %q but got %q for IntuneBrandingProfileAssignmentId", v.Expected.IntuneBrandingProfileAssignmentId, actual.IntuneBrandingProfileAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementIntuneBrandingProfileIdAssignmentId(t *testing.T) {
	segments := DeviceManagementIntuneBrandingProfileIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementIntuneBrandingProfileIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
