package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId{}

func TestNewDeviceManagementWindowsInformationProtectionNetworkLearningSummaryID(t *testing.T) {
	id := NewDeviceManagementWindowsInformationProtectionNetworkLearningSummaryID("windowsInformationProtectionNetworkLearningSummaryId")

	if id.WindowsInformationProtectionNetworkLearningSummaryId != "windowsInformationProtectionNetworkLearningSummaryId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsInformationProtectionNetworkLearningSummaryId'", id.WindowsInformationProtectionNetworkLearningSummaryId, "windowsInformationProtectionNetworkLearningSummaryId")
	}
}

func TestFormatDeviceManagementWindowsInformationProtectionNetworkLearningSummaryID(t *testing.T) {
	actual := NewDeviceManagementWindowsInformationProtectionNetworkLearningSummaryID("windowsInformationProtectionNetworkLearningSummaryId").ID()
	expected := "/deviceManagement/windowsInformationProtectionNetworkLearningSummaries/windowsInformationProtectionNetworkLearningSummaryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementWindowsInformationProtectionNetworkLearningSummaryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId
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
			Input: "/deviceManagement/windowsInformationProtectionNetworkLearningSummaries",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsInformationProtectionNetworkLearningSummaries/windowsInformationProtectionNetworkLearningSummaryId",
			Expected: &DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId{
				WindowsInformationProtectionNetworkLearningSummaryId: "windowsInformationProtectionNetworkLearningSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsInformationProtectionNetworkLearningSummaries/windowsInformationProtectionNetworkLearningSummaryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsInformationProtectionNetworkLearningSummaryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsInformationProtectionNetworkLearningSummaryId != v.Expected.WindowsInformationProtectionNetworkLearningSummaryId {
			t.Fatalf("Expected %q but got %q for WindowsInformationProtectionNetworkLearningSummaryId", v.Expected.WindowsInformationProtectionNetworkLearningSummaryId, actual.WindowsInformationProtectionNetworkLearningSummaryId)
		}

	}
}

func TestParseDeviceManagementWindowsInformationProtectionNetworkLearningSummaryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId
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
			Input: "/deviceManagement/windowsInformationProtectionNetworkLearningSummaries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsInFoRmAtIoNpRoTeCtIoNnEtWoRkLeArNiNgSuMmArIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsInformationProtectionNetworkLearningSummaries/windowsInformationProtectionNetworkLearningSummaryId",
			Expected: &DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId{
				WindowsInformationProtectionNetworkLearningSummaryId: "windowsInformationProtectionNetworkLearningSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsInformationProtectionNetworkLearningSummaries/windowsInformationProtectionNetworkLearningSummaryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsInFoRmAtIoNpRoTeCtIoNnEtWoRkLeArNiNgSuMmArIeS/wInDoWsInFoRmAtIoNpRoTeCtIoNnEtWoRkLeArNiNgSuMmArYiD",
			Expected: &DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId{
				WindowsInformationProtectionNetworkLearningSummaryId: "wInDoWsInFoRmAtIoNpRoTeCtIoNnEtWoRkLeArNiNgSuMmArYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsInFoRmAtIoNpRoTeCtIoNnEtWoRkLeArNiNgSuMmArIeS/wInDoWsInFoRmAtIoNpRoTeCtIoNnEtWoRkLeArNiNgSuMmArYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsInformationProtectionNetworkLearningSummaryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsInformationProtectionNetworkLearningSummaryId != v.Expected.WindowsInformationProtectionNetworkLearningSummaryId {
			t.Fatalf("Expected %q but got %q for WindowsInformationProtectionNetworkLearningSummaryId", v.Expected.WindowsInformationProtectionNetworkLearningSummaryId, actual.WindowsInformationProtectionNetworkLearningSummaryId)
		}

	}
}

func TestSegmentsForDeviceManagementWindowsInformationProtectionNetworkLearningSummaryId(t *testing.T) {
	segments := DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementWindowsInformationProtectionNetworkLearningSummaryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
