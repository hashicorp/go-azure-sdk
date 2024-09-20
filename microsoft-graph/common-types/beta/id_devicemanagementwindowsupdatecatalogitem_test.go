package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsUpdateCatalogItemId{}

func TestNewDeviceManagementWindowsUpdateCatalogItemID(t *testing.T) {
	id := NewDeviceManagementWindowsUpdateCatalogItemID("windowsUpdateCatalogItemId")

	if id.WindowsUpdateCatalogItemId != "windowsUpdateCatalogItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsUpdateCatalogItemId'", id.WindowsUpdateCatalogItemId, "windowsUpdateCatalogItemId")
	}
}

func TestFormatDeviceManagementWindowsUpdateCatalogItemID(t *testing.T) {
	actual := NewDeviceManagementWindowsUpdateCatalogItemID("windowsUpdateCatalogItemId").ID()
	expected := "/deviceManagement/windowsUpdateCatalogItems/windowsUpdateCatalogItemId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementWindowsUpdateCatalogItemID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsUpdateCatalogItemId
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
			Input: "/deviceManagement/windowsUpdateCatalogItems",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsUpdateCatalogItems/windowsUpdateCatalogItemId",
			Expected: &DeviceManagementWindowsUpdateCatalogItemId{
				WindowsUpdateCatalogItemId: "windowsUpdateCatalogItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsUpdateCatalogItems/windowsUpdateCatalogItemId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsUpdateCatalogItemID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsUpdateCatalogItemId != v.Expected.WindowsUpdateCatalogItemId {
			t.Fatalf("Expected %q but got %q for WindowsUpdateCatalogItemId", v.Expected.WindowsUpdateCatalogItemId, actual.WindowsUpdateCatalogItemId)
		}

	}
}

func TestParseDeviceManagementWindowsUpdateCatalogItemIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsUpdateCatalogItemId
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
			Input: "/deviceManagement/windowsUpdateCatalogItems",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsUpDaTeCaTaLoGiTeMs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsUpdateCatalogItems/windowsUpdateCatalogItemId",
			Expected: &DeviceManagementWindowsUpdateCatalogItemId{
				WindowsUpdateCatalogItemId: "windowsUpdateCatalogItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsUpdateCatalogItems/windowsUpdateCatalogItemId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsUpDaTeCaTaLoGiTeMs/wInDoWsUpDaTeCaTaLoGiTeMiD",
			Expected: &DeviceManagementWindowsUpdateCatalogItemId{
				WindowsUpdateCatalogItemId: "wInDoWsUpDaTeCaTaLoGiTeMiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsUpDaTeCaTaLoGiTeMs/wInDoWsUpDaTeCaTaLoGiTeMiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsUpdateCatalogItemIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsUpdateCatalogItemId != v.Expected.WindowsUpdateCatalogItemId {
			t.Fatalf("Expected %q but got %q for WindowsUpdateCatalogItemId", v.Expected.WindowsUpdateCatalogItemId, actual.WindowsUpdateCatalogItemId)
		}

	}
}

func TestSegmentsForDeviceManagementWindowsUpdateCatalogItemId(t *testing.T) {
	segments := DeviceManagementWindowsUpdateCatalogItemId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementWindowsUpdateCatalogItemId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
