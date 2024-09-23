package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId{}

func TestNewDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID(t *testing.T) {
	id := NewDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID("windowsDriverUpdateProfileId", "windowsDriverUpdateInventoryId")

	if id.WindowsDriverUpdateProfileId != "windowsDriverUpdateProfileId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsDriverUpdateProfileId'", id.WindowsDriverUpdateProfileId, "windowsDriverUpdateProfileId")
	}

	if id.WindowsDriverUpdateInventoryId != "windowsDriverUpdateInventoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsDriverUpdateInventoryId'", id.WindowsDriverUpdateInventoryId, "windowsDriverUpdateInventoryId")
	}
}

func TestFormatDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID(t *testing.T) {
	actual := NewDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID("windowsDriverUpdateProfileId", "windowsDriverUpdateInventoryId").ID()
	expected := "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId/driverInventories/windowsDriverUpdateInventoryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId
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
			// Incomplete URI
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId/driverInventories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId/driverInventories/windowsDriverUpdateInventoryId",
			Expected: &DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId{
				WindowsDriverUpdateProfileId:   "windowsDriverUpdateProfileId",
				WindowsDriverUpdateInventoryId: "windowsDriverUpdateInventoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId/driverInventories/windowsDriverUpdateInventoryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID(v.Input)
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

		if actual.WindowsDriverUpdateInventoryId != v.Expected.WindowsDriverUpdateInventoryId {
			t.Fatalf("Expected %q but got %q for WindowsDriverUpdateInventoryId", v.Expected.WindowsDriverUpdateInventoryId, actual.WindowsDriverUpdateInventoryId)
		}

	}
}

func TestParseDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId
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
			// Incomplete URI
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsDrIvErUpDaTePrOfIlEs/wInDoWsDrIvErUpDaTePrOfIlEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId/driverInventories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsDrIvErUpDaTePrOfIlEs/wInDoWsDrIvErUpDaTePrOfIlEiD/dRiVeRiNvEnToRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId/driverInventories/windowsDriverUpdateInventoryId",
			Expected: &DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId{
				WindowsDriverUpdateProfileId:   "windowsDriverUpdateProfileId",
				WindowsDriverUpdateInventoryId: "windowsDriverUpdateInventoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId/driverInventories/windowsDriverUpdateInventoryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsDrIvErUpDaTePrOfIlEs/wInDoWsDrIvErUpDaTePrOfIlEiD/dRiVeRiNvEnToRiEs/wInDoWsDrIvErUpDaTeInVeNtOrYiD",
			Expected: &DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId{
				WindowsDriverUpdateProfileId:   "wInDoWsDrIvErUpDaTePrOfIlEiD",
				WindowsDriverUpdateInventoryId: "wInDoWsDrIvErUpDaTeInVeNtOrYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsDrIvErUpDaTePrOfIlEs/wInDoWsDrIvErUpDaTePrOfIlEiD/dRiVeRiNvEnToRiEs/wInDoWsDrIvErUpDaTeInVeNtOrYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryIDInsensitively(v.Input)
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

		if actual.WindowsDriverUpdateInventoryId != v.Expected.WindowsDriverUpdateInventoryId {
			t.Fatalf("Expected %q but got %q for WindowsDriverUpdateInventoryId", v.Expected.WindowsDriverUpdateInventoryId, actual.WindowsDriverUpdateInventoryId)
		}

	}
}

func TestSegmentsForDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId(t *testing.T) {
	segments := DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
