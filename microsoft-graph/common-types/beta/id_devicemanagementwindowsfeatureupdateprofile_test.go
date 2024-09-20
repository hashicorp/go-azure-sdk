package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsFeatureUpdateProfileId{}

func TestNewDeviceManagementWindowsFeatureUpdateProfileID(t *testing.T) {
	id := NewDeviceManagementWindowsFeatureUpdateProfileID("windowsFeatureUpdateProfileId")

	if id.WindowsFeatureUpdateProfileId != "windowsFeatureUpdateProfileId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsFeatureUpdateProfileId'", id.WindowsFeatureUpdateProfileId, "windowsFeatureUpdateProfileId")
	}
}

func TestFormatDeviceManagementWindowsFeatureUpdateProfileID(t *testing.T) {
	actual := NewDeviceManagementWindowsFeatureUpdateProfileID("windowsFeatureUpdateProfileId").ID()
	expected := "/deviceManagement/windowsFeatureUpdateProfiles/windowsFeatureUpdateProfileId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementWindowsFeatureUpdateProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsFeatureUpdateProfileId
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
			// Valid URI
			Input: "/deviceManagement/windowsFeatureUpdateProfiles/windowsFeatureUpdateProfileId",
			Expected: &DeviceManagementWindowsFeatureUpdateProfileId{
				WindowsFeatureUpdateProfileId: "windowsFeatureUpdateProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsFeatureUpdateProfiles/windowsFeatureUpdateProfileId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsFeatureUpdateProfileID(v.Input)
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

	}
}

func TestParseDeviceManagementWindowsFeatureUpdateProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsFeatureUpdateProfileId
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
			// Valid URI
			Input: "/deviceManagement/windowsFeatureUpdateProfiles/windowsFeatureUpdateProfileId",
			Expected: &DeviceManagementWindowsFeatureUpdateProfileId{
				WindowsFeatureUpdateProfileId: "windowsFeatureUpdateProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsFeatureUpdateProfiles/windowsFeatureUpdateProfileId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsFeAtUrEuPdAtEpRoFiLeS/wInDoWsFeAtUrEuPdAtEpRoFiLeId",
			Expected: &DeviceManagementWindowsFeatureUpdateProfileId{
				WindowsFeatureUpdateProfileId: "wInDoWsFeAtUrEuPdAtEpRoFiLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsFeAtUrEuPdAtEpRoFiLeS/wInDoWsFeAtUrEuPdAtEpRoFiLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsFeatureUpdateProfileIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementWindowsFeatureUpdateProfileId(t *testing.T) {
	segments := DeviceManagementWindowsFeatureUpdateProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementWindowsFeatureUpdateProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
