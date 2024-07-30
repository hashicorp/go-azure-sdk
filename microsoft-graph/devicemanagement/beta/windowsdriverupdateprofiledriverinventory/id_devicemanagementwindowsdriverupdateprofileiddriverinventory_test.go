package windowsdriverupdateprofiledriverinventory

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId{}

func TestNewDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID(t *testing.T) {
	id := NewDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID("windowsDriverUpdateProfileIdValue", "windowsDriverUpdateInventoryIdValue")

	if id.WindowsDriverUpdateProfileId != "windowsDriverUpdateProfileIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsDriverUpdateProfileId'", id.WindowsDriverUpdateProfileId, "windowsDriverUpdateProfileIdValue")
	}

	if id.WindowsDriverUpdateInventoryId != "windowsDriverUpdateInventoryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsDriverUpdateInventoryId'", id.WindowsDriverUpdateInventoryId, "windowsDriverUpdateInventoryIdValue")
	}
}

func TestFormatDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID(t *testing.T) {
	actual := NewDeviceManagementWindowsDriverUpdateProfileIdDriverInventoryID("windowsDriverUpdateProfileIdValue", "windowsDriverUpdateInventoryIdValue").ID()
	expected := "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileIdValue/driverInventories/windowsDriverUpdateInventoryIdValue"
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
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileIdValue/driverInventories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileIdValue/driverInventories/windowsDriverUpdateInventoryIdValue",
			Expected: &DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId{
				WindowsDriverUpdateProfileId:   "windowsDriverUpdateProfileIdValue",
				WindowsDriverUpdateInventoryId: "windowsDriverUpdateInventoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileIdValue/driverInventories/windowsDriverUpdateInventoryIdValue/extra",
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
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsDrIvErUpDaTePrOfIlEs/wInDoWsDrIvErUpDaTePrOfIlEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileIdValue/driverInventories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsDrIvErUpDaTePrOfIlEs/wInDoWsDrIvErUpDaTePrOfIlEiDvAlUe/dRiVeRiNvEnToRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileIdValue/driverInventories/windowsDriverUpdateInventoryIdValue",
			Expected: &DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId{
				WindowsDriverUpdateProfileId:   "windowsDriverUpdateProfileIdValue",
				WindowsDriverUpdateInventoryId: "windowsDriverUpdateInventoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileIdValue/driverInventories/windowsDriverUpdateInventoryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsDrIvErUpDaTePrOfIlEs/wInDoWsDrIvErUpDaTePrOfIlEiDvAlUe/dRiVeRiNvEnToRiEs/wInDoWsDrIvErUpDaTeInVeNtOrYiDvAlUe",
			Expected: &DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId{
				WindowsDriverUpdateProfileId:   "wInDoWsDrIvErUpDaTePrOfIlEiDvAlUe",
				WindowsDriverUpdateInventoryId: "wInDoWsDrIvErUpDaTeInVeNtOrYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsDrIvErUpDaTePrOfIlEs/wInDoWsDrIvErUpDaTePrOfIlEiDvAlUe/dRiVeRiNvEnToRiEs/wInDoWsDrIvErUpDaTeInVeNtOrYiDvAlUe/extra",
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
