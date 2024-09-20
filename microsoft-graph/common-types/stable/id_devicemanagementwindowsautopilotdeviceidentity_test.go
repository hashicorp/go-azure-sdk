package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsAutopilotDeviceIdentityId{}

func TestNewDeviceManagementWindowsAutopilotDeviceIdentityID(t *testing.T) {
	id := NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityId")

	if id.WindowsAutopilotDeviceIdentityId != "windowsAutopilotDeviceIdentityId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsAutopilotDeviceIdentityId'", id.WindowsAutopilotDeviceIdentityId, "windowsAutopilotDeviceIdentityId")
	}
}

func TestFormatDeviceManagementWindowsAutopilotDeviceIdentityID(t *testing.T) {
	actual := NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityId").ID()
	expected := "/deviceManagement/windowsAutopilotDeviceIdentities/windowsAutopilotDeviceIdentityId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementWindowsAutopilotDeviceIdentityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsAutopilotDeviceIdentityId
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
			Input: "/deviceManagement/windowsAutopilotDeviceIdentities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsAutopilotDeviceIdentities/windowsAutopilotDeviceIdentityId",
			Expected: &DeviceManagementWindowsAutopilotDeviceIdentityId{
				WindowsAutopilotDeviceIdentityId: "windowsAutopilotDeviceIdentityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsAutopilotDeviceIdentities/windowsAutopilotDeviceIdentityId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsAutopilotDeviceIdentityID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsAutopilotDeviceIdentityId != v.Expected.WindowsAutopilotDeviceIdentityId {
			t.Fatalf("Expected %q but got %q for WindowsAutopilotDeviceIdentityId", v.Expected.WindowsAutopilotDeviceIdentityId, actual.WindowsAutopilotDeviceIdentityId)
		}

	}
}

func TestParseDeviceManagementWindowsAutopilotDeviceIdentityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsAutopilotDeviceIdentityId
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
			Input: "/deviceManagement/windowsAutopilotDeviceIdentities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsAuToPiLoTdEvIcEiDeNtItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsAutopilotDeviceIdentities/windowsAutopilotDeviceIdentityId",
			Expected: &DeviceManagementWindowsAutopilotDeviceIdentityId{
				WindowsAutopilotDeviceIdentityId: "windowsAutopilotDeviceIdentityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsAutopilotDeviceIdentities/windowsAutopilotDeviceIdentityId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsAuToPiLoTdEvIcEiDeNtItIeS/wInDoWsAuToPiLoTdEvIcEiDeNtItYiD",
			Expected: &DeviceManagementWindowsAutopilotDeviceIdentityId{
				WindowsAutopilotDeviceIdentityId: "wInDoWsAuToPiLoTdEvIcEiDeNtItYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsAuToPiLoTdEvIcEiDeNtItIeS/wInDoWsAuToPiLoTdEvIcEiDeNtItYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsAutopilotDeviceIdentityIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsAutopilotDeviceIdentityId != v.Expected.WindowsAutopilotDeviceIdentityId {
			t.Fatalf("Expected %q but got %q for WindowsAutopilotDeviceIdentityId", v.Expected.WindowsAutopilotDeviceIdentityId, actual.WindowsAutopilotDeviceIdentityId)
		}

	}
}

func TestSegmentsForDeviceManagementWindowsAutopilotDeviceIdentityId(t *testing.T) {
	segments := DeviceManagementWindowsAutopilotDeviceIdentityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementWindowsAutopilotDeviceIdentityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
