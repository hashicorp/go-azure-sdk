package windowsqualityupdateprofileassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsQualityUpdateProfileId{}

func TestNewDeviceManagementWindowsQualityUpdateProfileID(t *testing.T) {
	id := NewDeviceManagementWindowsQualityUpdateProfileID("windowsQualityUpdateProfileIdValue")

	if id.WindowsQualityUpdateProfileId != "windowsQualityUpdateProfileIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsQualityUpdateProfileId'", id.WindowsQualityUpdateProfileId, "windowsQualityUpdateProfileIdValue")
	}
}

func TestFormatDeviceManagementWindowsQualityUpdateProfileID(t *testing.T) {
	actual := NewDeviceManagementWindowsQualityUpdateProfileID("windowsQualityUpdateProfileIdValue").ID()
	expected := "/deviceManagement/windowsQualityUpdateProfiles/windowsQualityUpdateProfileIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementWindowsQualityUpdateProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsQualityUpdateProfileId
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
			Input: "/deviceManagement/windowsQualityUpdateProfiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsQualityUpdateProfiles/windowsQualityUpdateProfileIdValue",
			Expected: &DeviceManagementWindowsQualityUpdateProfileId{
				WindowsQualityUpdateProfileId: "windowsQualityUpdateProfileIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsQualityUpdateProfiles/windowsQualityUpdateProfileIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsQualityUpdateProfileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsQualityUpdateProfileId != v.Expected.WindowsQualityUpdateProfileId {
			t.Fatalf("Expected %q but got %q for WindowsQualityUpdateProfileId", v.Expected.WindowsQualityUpdateProfileId, actual.WindowsQualityUpdateProfileId)
		}

	}
}

func TestParseDeviceManagementWindowsQualityUpdateProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsQualityUpdateProfileId
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
			Input: "/deviceManagement/windowsQualityUpdateProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsQuAlItYuPdAtEpRoFiLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsQualityUpdateProfiles/windowsQualityUpdateProfileIdValue",
			Expected: &DeviceManagementWindowsQualityUpdateProfileId{
				WindowsQualityUpdateProfileId: "windowsQualityUpdateProfileIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsQualityUpdateProfiles/windowsQualityUpdateProfileIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsQuAlItYuPdAtEpRoFiLeS/wInDoWsQuAlItYuPdAtEpRoFiLeIdVaLuE",
			Expected: &DeviceManagementWindowsQualityUpdateProfileId{
				WindowsQualityUpdateProfileId: "wInDoWsQuAlItYuPdAtEpRoFiLeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsQuAlItYuPdAtEpRoFiLeS/wInDoWsQuAlItYuPdAtEpRoFiLeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsQualityUpdateProfileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsQualityUpdateProfileId != v.Expected.WindowsQualityUpdateProfileId {
			t.Fatalf("Expected %q but got %q for WindowsQualityUpdateProfileId", v.Expected.WindowsQualityUpdateProfileId, actual.WindowsQualityUpdateProfileId)
		}

	}
}

func TestSegmentsForDeviceManagementWindowsQualityUpdateProfileId(t *testing.T) {
	segments := DeviceManagementWindowsQualityUpdateProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementWindowsQualityUpdateProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
