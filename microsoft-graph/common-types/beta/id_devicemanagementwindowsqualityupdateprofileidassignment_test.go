package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsQualityUpdateProfileIdAssignmentId{}

func TestNewDeviceManagementWindowsQualityUpdateProfileIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementWindowsQualityUpdateProfileIdAssignmentID("windowsQualityUpdateProfileId", "windowsQualityUpdateProfileAssignmentId")

	if id.WindowsQualityUpdateProfileId != "windowsQualityUpdateProfileId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsQualityUpdateProfileId'", id.WindowsQualityUpdateProfileId, "windowsQualityUpdateProfileId")
	}

	if id.WindowsQualityUpdateProfileAssignmentId != "windowsQualityUpdateProfileAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsQualityUpdateProfileAssignmentId'", id.WindowsQualityUpdateProfileAssignmentId, "windowsQualityUpdateProfileAssignmentId")
	}
}

func TestFormatDeviceManagementWindowsQualityUpdateProfileIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementWindowsQualityUpdateProfileIdAssignmentID("windowsQualityUpdateProfileId", "windowsQualityUpdateProfileAssignmentId").ID()
	expected := "/deviceManagement/windowsQualityUpdateProfiles/windowsQualityUpdateProfileId/assignments/windowsQualityUpdateProfileAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementWindowsQualityUpdateProfileIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsQualityUpdateProfileIdAssignmentId
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
			// Incomplete URI
			Input: "/deviceManagement/windowsQualityUpdateProfiles/windowsQualityUpdateProfileId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsQualityUpdateProfiles/windowsQualityUpdateProfileId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsQualityUpdateProfiles/windowsQualityUpdateProfileId/assignments/windowsQualityUpdateProfileAssignmentId",
			Expected: &DeviceManagementWindowsQualityUpdateProfileIdAssignmentId{
				WindowsQualityUpdateProfileId:           "windowsQualityUpdateProfileId",
				WindowsQualityUpdateProfileAssignmentId: "windowsQualityUpdateProfileAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsQualityUpdateProfiles/windowsQualityUpdateProfileId/assignments/windowsQualityUpdateProfileAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsQualityUpdateProfileIdAssignmentID(v.Input)
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

		if actual.WindowsQualityUpdateProfileAssignmentId != v.Expected.WindowsQualityUpdateProfileAssignmentId {
			t.Fatalf("Expected %q but got %q for WindowsQualityUpdateProfileAssignmentId", v.Expected.WindowsQualityUpdateProfileAssignmentId, actual.WindowsQualityUpdateProfileAssignmentId)
		}

	}
}

func TestParseDeviceManagementWindowsQualityUpdateProfileIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsQualityUpdateProfileIdAssignmentId
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
			// Incomplete URI
			Input: "/deviceManagement/windowsQualityUpdateProfiles/windowsQualityUpdateProfileId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsQuAlItYuPdAtEpRoFiLeS/wInDoWsQuAlItYuPdAtEpRoFiLeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsQualityUpdateProfiles/windowsQualityUpdateProfileId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsQuAlItYuPdAtEpRoFiLeS/wInDoWsQuAlItYuPdAtEpRoFiLeId/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsQualityUpdateProfiles/windowsQualityUpdateProfileId/assignments/windowsQualityUpdateProfileAssignmentId",
			Expected: &DeviceManagementWindowsQualityUpdateProfileIdAssignmentId{
				WindowsQualityUpdateProfileId:           "windowsQualityUpdateProfileId",
				WindowsQualityUpdateProfileAssignmentId: "windowsQualityUpdateProfileAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsQualityUpdateProfiles/windowsQualityUpdateProfileId/assignments/windowsQualityUpdateProfileAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsQuAlItYuPdAtEpRoFiLeS/wInDoWsQuAlItYuPdAtEpRoFiLeId/aSsIgNmEnTs/wInDoWsQuAlItYuPdAtEpRoFiLeAsSiGnMeNtId",
			Expected: &DeviceManagementWindowsQualityUpdateProfileIdAssignmentId{
				WindowsQualityUpdateProfileId:           "wInDoWsQuAlItYuPdAtEpRoFiLeId",
				WindowsQualityUpdateProfileAssignmentId: "wInDoWsQuAlItYuPdAtEpRoFiLeAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsQuAlItYuPdAtEpRoFiLeS/wInDoWsQuAlItYuPdAtEpRoFiLeId/aSsIgNmEnTs/wInDoWsQuAlItYuPdAtEpRoFiLeAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsQualityUpdateProfileIdAssignmentIDInsensitively(v.Input)
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

		if actual.WindowsQualityUpdateProfileAssignmentId != v.Expected.WindowsQualityUpdateProfileAssignmentId {
			t.Fatalf("Expected %q but got %q for WindowsQualityUpdateProfileAssignmentId", v.Expected.WindowsQualityUpdateProfileAssignmentId, actual.WindowsQualityUpdateProfileAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementWindowsQualityUpdateProfileIdAssignmentId(t *testing.T) {
	segments := DeviceManagementWindowsQualityUpdateProfileIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementWindowsQualityUpdateProfileIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
