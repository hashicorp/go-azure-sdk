package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId{}

func TestNewDeviceManagementWindowsFeatureUpdateProfileIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementWindowsFeatureUpdateProfileIdAssignmentID("windowsFeatureUpdateProfileId", "windowsFeatureUpdateProfileAssignmentId")

	if id.WindowsFeatureUpdateProfileId != "windowsFeatureUpdateProfileId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsFeatureUpdateProfileId'", id.WindowsFeatureUpdateProfileId, "windowsFeatureUpdateProfileId")
	}

	if id.WindowsFeatureUpdateProfileAssignmentId != "windowsFeatureUpdateProfileAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsFeatureUpdateProfileAssignmentId'", id.WindowsFeatureUpdateProfileAssignmentId, "windowsFeatureUpdateProfileAssignmentId")
	}
}

func TestFormatDeviceManagementWindowsFeatureUpdateProfileIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementWindowsFeatureUpdateProfileIdAssignmentID("windowsFeatureUpdateProfileId", "windowsFeatureUpdateProfileAssignmentId").ID()
	expected := "/deviceManagement/windowsFeatureUpdateProfiles/windowsFeatureUpdateProfileId/assignments/windowsFeatureUpdateProfileAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementWindowsFeatureUpdateProfileIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId
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
			Input: "/deviceManagement/windowsFeatureUpdateProfiles",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsFeatureUpdateProfiles/windowsFeatureUpdateProfileId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsFeatureUpdateProfiles/windowsFeatureUpdateProfileId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsFeatureUpdateProfiles/windowsFeatureUpdateProfileId/assignments/windowsFeatureUpdateProfileAssignmentId",
			Expected: &DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId{
				WindowsFeatureUpdateProfileId:           "windowsFeatureUpdateProfileId",
				WindowsFeatureUpdateProfileAssignmentId: "windowsFeatureUpdateProfileAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsFeatureUpdateProfiles/windowsFeatureUpdateProfileId/assignments/windowsFeatureUpdateProfileAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsFeatureUpdateProfileIdAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsFeatureUpdateProfileId != v.Expected.WindowsFeatureUpdateProfileId {
			t.Fatalf("Expected %q but got %q for WindowsFeatureUpdateProfileId", v.Expected.WindowsFeatureUpdateProfileId, actual.WindowsFeatureUpdateProfileId)
		}

		if actual.WindowsFeatureUpdateProfileAssignmentId != v.Expected.WindowsFeatureUpdateProfileAssignmentId {
			t.Fatalf("Expected %q but got %q for WindowsFeatureUpdateProfileAssignmentId", v.Expected.WindowsFeatureUpdateProfileAssignmentId, actual.WindowsFeatureUpdateProfileAssignmentId)
		}

	}
}

func TestParseDeviceManagementWindowsFeatureUpdateProfileIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId
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
			Input: "/deviceManagement/windowsFeatureUpdateProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsFeAtUrEuPdAtEpRoFiLeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsFeatureUpdateProfiles/windowsFeatureUpdateProfileId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsFeAtUrEuPdAtEpRoFiLeS/wInDoWsFeAtUrEuPdAtEpRoFiLeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsFeatureUpdateProfiles/windowsFeatureUpdateProfileId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsFeAtUrEuPdAtEpRoFiLeS/wInDoWsFeAtUrEuPdAtEpRoFiLeId/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsFeatureUpdateProfiles/windowsFeatureUpdateProfileId/assignments/windowsFeatureUpdateProfileAssignmentId",
			Expected: &DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId{
				WindowsFeatureUpdateProfileId:           "windowsFeatureUpdateProfileId",
				WindowsFeatureUpdateProfileAssignmentId: "windowsFeatureUpdateProfileAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsFeatureUpdateProfiles/windowsFeatureUpdateProfileId/assignments/windowsFeatureUpdateProfileAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsFeAtUrEuPdAtEpRoFiLeS/wInDoWsFeAtUrEuPdAtEpRoFiLeId/aSsIgNmEnTs/wInDoWsFeAtUrEuPdAtEpRoFiLeAsSiGnMeNtId",
			Expected: &DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId{
				WindowsFeatureUpdateProfileId:           "wInDoWsFeAtUrEuPdAtEpRoFiLeId",
				WindowsFeatureUpdateProfileAssignmentId: "wInDoWsFeAtUrEuPdAtEpRoFiLeAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsFeAtUrEuPdAtEpRoFiLeS/wInDoWsFeAtUrEuPdAtEpRoFiLeId/aSsIgNmEnTs/wInDoWsFeAtUrEuPdAtEpRoFiLeAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsFeatureUpdateProfileIdAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsFeatureUpdateProfileId != v.Expected.WindowsFeatureUpdateProfileId {
			t.Fatalf("Expected %q but got %q for WindowsFeatureUpdateProfileId", v.Expected.WindowsFeatureUpdateProfileId, actual.WindowsFeatureUpdateProfileId)
		}

		if actual.WindowsFeatureUpdateProfileAssignmentId != v.Expected.WindowsFeatureUpdateProfileAssignmentId {
			t.Fatalf("Expected %q but got %q for WindowsFeatureUpdateProfileAssignmentId", v.Expected.WindowsFeatureUpdateProfileAssignmentId, actual.WindowsFeatureUpdateProfileAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementWindowsFeatureUpdateProfileIdAssignmentId(t *testing.T) {
	segments := DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementWindowsFeatureUpdateProfileIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
