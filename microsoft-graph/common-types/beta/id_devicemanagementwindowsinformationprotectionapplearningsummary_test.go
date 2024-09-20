package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsInformationProtectionAppLearningSummaryId{}

func TestNewDeviceManagementWindowsInformationProtectionAppLearningSummaryID(t *testing.T) {
	id := NewDeviceManagementWindowsInformationProtectionAppLearningSummaryID("windowsInformationProtectionAppLearningSummaryId")

	if id.WindowsInformationProtectionAppLearningSummaryId != "windowsInformationProtectionAppLearningSummaryId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsInformationProtectionAppLearningSummaryId'", id.WindowsInformationProtectionAppLearningSummaryId, "windowsInformationProtectionAppLearningSummaryId")
	}
}

func TestFormatDeviceManagementWindowsInformationProtectionAppLearningSummaryID(t *testing.T) {
	actual := NewDeviceManagementWindowsInformationProtectionAppLearningSummaryID("windowsInformationProtectionAppLearningSummaryId").ID()
	expected := "/deviceManagement/windowsInformationProtectionAppLearningSummaries/windowsInformationProtectionAppLearningSummaryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementWindowsInformationProtectionAppLearningSummaryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsInformationProtectionAppLearningSummaryId
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
			Input: "/deviceManagement/windowsInformationProtectionAppLearningSummaries",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsInformationProtectionAppLearningSummaries/windowsInformationProtectionAppLearningSummaryId",
			Expected: &DeviceManagementWindowsInformationProtectionAppLearningSummaryId{
				WindowsInformationProtectionAppLearningSummaryId: "windowsInformationProtectionAppLearningSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsInformationProtectionAppLearningSummaries/windowsInformationProtectionAppLearningSummaryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsInformationProtectionAppLearningSummaryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsInformationProtectionAppLearningSummaryId != v.Expected.WindowsInformationProtectionAppLearningSummaryId {
			t.Fatalf("Expected %q but got %q for WindowsInformationProtectionAppLearningSummaryId", v.Expected.WindowsInformationProtectionAppLearningSummaryId, actual.WindowsInformationProtectionAppLearningSummaryId)
		}

	}
}

func TestParseDeviceManagementWindowsInformationProtectionAppLearningSummaryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsInformationProtectionAppLearningSummaryId
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
			Input: "/deviceManagement/windowsInformationProtectionAppLearningSummaries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsInFoRmAtIoNpRoTeCtIoNaPpLeArNiNgSuMmArIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsInformationProtectionAppLearningSummaries/windowsInformationProtectionAppLearningSummaryId",
			Expected: &DeviceManagementWindowsInformationProtectionAppLearningSummaryId{
				WindowsInformationProtectionAppLearningSummaryId: "windowsInformationProtectionAppLearningSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsInformationProtectionAppLearningSummaries/windowsInformationProtectionAppLearningSummaryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsInFoRmAtIoNpRoTeCtIoNaPpLeArNiNgSuMmArIeS/wInDoWsInFoRmAtIoNpRoTeCtIoNaPpLeArNiNgSuMmArYiD",
			Expected: &DeviceManagementWindowsInformationProtectionAppLearningSummaryId{
				WindowsInformationProtectionAppLearningSummaryId: "wInDoWsInFoRmAtIoNpRoTeCtIoNaPpLeArNiNgSuMmArYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsInFoRmAtIoNpRoTeCtIoNaPpLeArNiNgSuMmArIeS/wInDoWsInFoRmAtIoNpRoTeCtIoNaPpLeArNiNgSuMmArYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsInformationProtectionAppLearningSummaryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsInformationProtectionAppLearningSummaryId != v.Expected.WindowsInformationProtectionAppLearningSummaryId {
			t.Fatalf("Expected %q but got %q for WindowsInformationProtectionAppLearningSummaryId", v.Expected.WindowsInformationProtectionAppLearningSummaryId, actual.WindowsInformationProtectionAppLearningSummaryId)
		}

	}
}

func TestSegmentsForDeviceManagementWindowsInformationProtectionAppLearningSummaryId(t *testing.T) {
	segments := DeviceManagementWindowsInformationProtectionAppLearningSummaryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementWindowsInformationProtectionAppLearningSummaryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
