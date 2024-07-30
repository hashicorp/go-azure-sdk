package windowsautopilotdeploymentprofileassigneddeviceintendeddeploymentprofile

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId{}

func TestNewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID(t *testing.T) {
	id := NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileIdValue", "windowsAutopilotDeviceIdentityIdValue")

	if id.WindowsAutopilotDeploymentProfileId != "windowsAutopilotDeploymentProfileIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsAutopilotDeploymentProfileId'", id.WindowsAutopilotDeploymentProfileId, "windowsAutopilotDeploymentProfileIdValue")
	}

	if id.WindowsAutopilotDeviceIdentityId != "windowsAutopilotDeviceIdentityIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsAutopilotDeviceIdentityId'", id.WindowsAutopilotDeviceIdentityId, "windowsAutopilotDeviceIdentityIdValue")
	}
}

func TestFormatDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID(t *testing.T) {
	actual := NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileIdValue", "windowsAutopilotDeviceIdentityIdValue").ID()
	expected := "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileIdValue/assignedDevices/windowsAutopilotDeviceIdentityIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId
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
			// Incomplete URI
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileIdValue/assignedDevices",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileIdValue/assignedDevices/windowsAutopilotDeviceIdentityIdValue",
			Expected: &DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId{
				WindowsAutopilotDeploymentProfileId: "windowsAutopilotDeploymentProfileIdValue",
				WindowsAutopilotDeviceIdentityId:    "windowsAutopilotDeviceIdentityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileIdValue/assignedDevices/windowsAutopilotDeviceIdentityIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID(v.Input)
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

		if actual.WindowsAutopilotDeviceIdentityId != v.Expected.WindowsAutopilotDeviceIdentityId {
			t.Fatalf("Expected %q but got %q for WindowsAutopilotDeviceIdentityId", v.Expected.WindowsAutopilotDeviceIdentityId, actual.WindowsAutopilotDeviceIdentityId)
		}

	}
}

func TestParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId
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
			// Incomplete URI
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeS/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileIdValue/assignedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeS/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeIdVaLuE/aSsIgNeDdEvIcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileIdValue/assignedDevices/windowsAutopilotDeviceIdentityIdValue",
			Expected: &DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId{
				WindowsAutopilotDeploymentProfileId: "windowsAutopilotDeploymentProfileIdValue",
				WindowsAutopilotDeviceIdentityId:    "windowsAutopilotDeviceIdentityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileIdValue/assignedDevices/windowsAutopilotDeviceIdentityIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeS/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeIdVaLuE/aSsIgNeDdEvIcEs/wInDoWsAuToPiLoTdEvIcEiDeNtItYiDvAlUe",
			Expected: &DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId{
				WindowsAutopilotDeploymentProfileId: "wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeIdVaLuE",
				WindowsAutopilotDeviceIdentityId:    "wInDoWsAuToPiLoTdEvIcEiDeNtItYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeS/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeIdVaLuE/aSsIgNeDdEvIcEs/wInDoWsAuToPiLoTdEvIcEiDeNtItYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceIDInsensitively(v.Input)
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

		if actual.WindowsAutopilotDeviceIdentityId != v.Expected.WindowsAutopilotDeviceIdentityId {
			t.Fatalf("Expected %q but got %q for WindowsAutopilotDeviceIdentityId", v.Expected.WindowsAutopilotDeviceIdentityId, actual.WindowsAutopilotDeviceIdentityId)
		}

	}
}

func TestSegmentsForDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId(t *testing.T) {
	segments := DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
