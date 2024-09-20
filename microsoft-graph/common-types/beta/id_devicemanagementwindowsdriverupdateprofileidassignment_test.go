package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsDriverUpdateProfileIdAssignmentId{}

func TestNewDeviceManagementWindowsDriverUpdateProfileIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementWindowsDriverUpdateProfileIdAssignmentID("windowsDriverUpdateProfileId", "windowsDriverUpdateProfileAssignmentId")

	if id.WindowsDriverUpdateProfileId != "windowsDriverUpdateProfileId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsDriverUpdateProfileId'", id.WindowsDriverUpdateProfileId, "windowsDriverUpdateProfileId")
	}

	if id.WindowsDriverUpdateProfileAssignmentId != "windowsDriverUpdateProfileAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsDriverUpdateProfileAssignmentId'", id.WindowsDriverUpdateProfileAssignmentId, "windowsDriverUpdateProfileAssignmentId")
	}
}

func TestFormatDeviceManagementWindowsDriverUpdateProfileIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementWindowsDriverUpdateProfileIdAssignmentID("windowsDriverUpdateProfileId", "windowsDriverUpdateProfileAssignmentId").ID()
	expected := "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId/assignments/windowsDriverUpdateProfileAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementWindowsDriverUpdateProfileIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsDriverUpdateProfileIdAssignmentId
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
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId/assignments/windowsDriverUpdateProfileAssignmentId",
			Expected: &DeviceManagementWindowsDriverUpdateProfileIdAssignmentId{
				WindowsDriverUpdateProfileId:           "windowsDriverUpdateProfileId",
				WindowsDriverUpdateProfileAssignmentId: "windowsDriverUpdateProfileAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId/assignments/windowsDriverUpdateProfileAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsDriverUpdateProfileIdAssignmentID(v.Input)
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

		if actual.WindowsDriverUpdateProfileAssignmentId != v.Expected.WindowsDriverUpdateProfileAssignmentId {
			t.Fatalf("Expected %q but got %q for WindowsDriverUpdateProfileAssignmentId", v.Expected.WindowsDriverUpdateProfileAssignmentId, actual.WindowsDriverUpdateProfileAssignmentId)
		}

	}
}

func TestParseDeviceManagementWindowsDriverUpdateProfileIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsDriverUpdateProfileIdAssignmentId
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
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsDrIvErUpDaTePrOfIlEs/wInDoWsDrIvErUpDaTePrOfIlEiD/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId/assignments/windowsDriverUpdateProfileAssignmentId",
			Expected: &DeviceManagementWindowsDriverUpdateProfileIdAssignmentId{
				WindowsDriverUpdateProfileId:           "windowsDriverUpdateProfileId",
				WindowsDriverUpdateProfileAssignmentId: "windowsDriverUpdateProfileAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsDriverUpdateProfiles/windowsDriverUpdateProfileId/assignments/windowsDriverUpdateProfileAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsDrIvErUpDaTePrOfIlEs/wInDoWsDrIvErUpDaTePrOfIlEiD/aSsIgNmEnTs/wInDoWsDrIvErUpDaTePrOfIlEaSsIgNmEnTiD",
			Expected: &DeviceManagementWindowsDriverUpdateProfileIdAssignmentId{
				WindowsDriverUpdateProfileId:           "wInDoWsDrIvErUpDaTePrOfIlEiD",
				WindowsDriverUpdateProfileAssignmentId: "wInDoWsDrIvErUpDaTePrOfIlEaSsIgNmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsDrIvErUpDaTePrOfIlEs/wInDoWsDrIvErUpDaTePrOfIlEiD/aSsIgNmEnTs/wInDoWsDrIvErUpDaTePrOfIlEaSsIgNmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsDriverUpdateProfileIdAssignmentIDInsensitively(v.Input)
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

		if actual.WindowsDriverUpdateProfileAssignmentId != v.Expected.WindowsDriverUpdateProfileAssignmentId {
			t.Fatalf("Expected %q but got %q for WindowsDriverUpdateProfileAssignmentId", v.Expected.WindowsDriverUpdateProfileAssignmentId, actual.WindowsDriverUpdateProfileAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementWindowsDriverUpdateProfileIdAssignmentId(t *testing.T) {
	segments := DeviceManagementWindowsDriverUpdateProfileIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementWindowsDriverUpdateProfileIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
