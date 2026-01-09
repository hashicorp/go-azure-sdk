package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsDriverUpdateProfileId{}

func TestNewDeviceManagementWindowsDriverUpdateProfileID(t *testing.T) {
	id := NewDeviceManagementWindowsDriverUpdateProfileID("windowsDriverUpdateProfileId")

	if id.WindowsDriverUpdateProfileId != "windowsDriverUpdateProfileId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsDriverUpdateProfileId'", id.WindowsDriverUpdateProfileId, "windowsDriverUpdateProfileId")
	}
}

func TestFormatDeviceManagementWindowsDriverUpdateProfileID(t *testing.T) {
	actual := NewDeviceManagementWindowsDriverUpdateProfileID("windowsDriverUpdateProfileId").ID()
	expected := "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementWindowsDriverUpdateProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsDriverUpdateProfileId
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
			Input: "/deviceManagement/windowsDriverUpdateProfiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId",
			Expected: &DeviceManagementWindowsDriverUpdateProfileId{
				WindowsDriverUpdateProfileId: "windowsDriverUpdateProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsDriverUpdateProfileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsDriverUpdateProfileId != v.Expected.WindowsDriverUpdateProfileId {
			t.Fatalf("Expected %q but got %q for WindowsDriverUpdateProfileId", v.Expected.WindowsDriverUpdateProfileId, actual.WindowsDriverUpdateProfileId)
		}

	}
}

func TestParseDeviceManagementWindowsDriverUpdateProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsDriverUpdateProfileId
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
			Input: "/deviceManagement/windowsDriverUpdateProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsDrIvErUpDaTePrOfIlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId",
			Expected: &DeviceManagementWindowsDriverUpdateProfileId{
				WindowsDriverUpdateProfileId: "windowsDriverUpdateProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsDrIvErUpDaTePrOfIlEs/wInDoWsDrIvErUpDaTePrOfIlEiD",
			Expected: &DeviceManagementWindowsDriverUpdateProfileId{
				WindowsDriverUpdateProfileId: "wInDoWsDrIvErUpDaTePrOfIlEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsDrIvErUpDaTePrOfIlEs/wInDoWsDrIvErUpDaTePrOfIlEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsDriverUpdateProfileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsDriverUpdateProfileId != v.Expected.WindowsDriverUpdateProfileId {
			t.Fatalf("Expected %q but got %q for WindowsDriverUpdateProfileId", v.Expected.WindowsDriverUpdateProfileId, actual.WindowsDriverUpdateProfileId)
		}

	}
}

func TestSegmentsForDeviceManagementWindowsDriverUpdateProfileId(t *testing.T) {
	segments := DeviceManagementWindowsDriverUpdateProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementWindowsDriverUpdateProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
