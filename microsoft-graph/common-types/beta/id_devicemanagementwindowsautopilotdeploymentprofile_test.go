package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsAutopilotDeploymentProfileId{}

func TestNewDeviceManagementWindowsAutopilotDeploymentProfileID(t *testing.T) {
	id := NewDeviceManagementWindowsAutopilotDeploymentProfileID("windowsAutopilotDeploymentProfileId")

	if id.WindowsAutopilotDeploymentProfileId != "windowsAutopilotDeploymentProfileId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsAutopilotDeploymentProfileId'", id.WindowsAutopilotDeploymentProfileId, "windowsAutopilotDeploymentProfileId")
	}
}

func TestFormatDeviceManagementWindowsAutopilotDeploymentProfileID(t *testing.T) {
	actual := NewDeviceManagementWindowsAutopilotDeploymentProfileID("windowsAutopilotDeploymentProfileId").ID()
	expected := "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementWindowsAutopilotDeploymentProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsAutopilotDeploymentProfileId
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
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileId",
			Expected: &DeviceManagementWindowsAutopilotDeploymentProfileId{
				WindowsAutopilotDeploymentProfileId: "windowsAutopilotDeploymentProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsAutopilotDeploymentProfileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsAutopilotDeploymentProfileId != v.Expected.WindowsAutopilotDeploymentProfileId {
			t.Fatalf("Expected %q but got %q for WindowsAutopilotDeploymentProfileId", v.Expected.WindowsAutopilotDeploymentProfileId, actual.WindowsAutopilotDeploymentProfileId)
		}

	}
}

func TestParseDeviceManagementWindowsAutopilotDeploymentProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsAutopilotDeploymentProfileId
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
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileId",
			Expected: &DeviceManagementWindowsAutopilotDeploymentProfileId{
				WindowsAutopilotDeploymentProfileId: "windowsAutopilotDeploymentProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeS/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeId",
			Expected: &DeviceManagementWindowsAutopilotDeploymentProfileId{
				WindowsAutopilotDeploymentProfileId: "wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeS/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsAutopilotDeploymentProfileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsAutopilotDeploymentProfileId != v.Expected.WindowsAutopilotDeploymentProfileId {
			t.Fatalf("Expected %q but got %q for WindowsAutopilotDeploymentProfileId", v.Expected.WindowsAutopilotDeploymentProfileId, actual.WindowsAutopilotDeploymentProfileId)
		}

	}
}

func TestSegmentsForDeviceManagementWindowsAutopilotDeploymentProfileId(t *testing.T) {
	segments := DeviceManagementWindowsAutopilotDeploymentProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementWindowsAutopilotDeploymentProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
