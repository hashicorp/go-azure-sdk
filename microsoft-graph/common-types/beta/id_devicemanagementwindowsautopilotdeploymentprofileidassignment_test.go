package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId{}

func TestNewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentID("windowsAutopilotDeploymentProfileId", "windowsAutopilotDeploymentProfileAssignmentId")

	if id.WindowsAutopilotDeploymentProfileId != "windowsAutopilotDeploymentProfileId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsAutopilotDeploymentProfileId'", id.WindowsAutopilotDeploymentProfileId, "windowsAutopilotDeploymentProfileId")
	}

	if id.WindowsAutopilotDeploymentProfileAssignmentId != "windowsAutopilotDeploymentProfileAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsAutopilotDeploymentProfileAssignmentId'", id.WindowsAutopilotDeploymentProfileAssignmentId, "windowsAutopilotDeploymentProfileAssignmentId")
	}
}

func TestFormatDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentID("windowsAutopilotDeploymentProfileId", "windowsAutopilotDeploymentProfileAssignmentId").ID()
	expected := "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileId/assignments/windowsAutopilotDeploymentProfileAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId
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
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileId/assignments/windowsAutopilotDeploymentProfileAssignmentId",
			Expected: &DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId{
				WindowsAutopilotDeploymentProfileId:           "windowsAutopilotDeploymentProfileId",
				WindowsAutopilotDeploymentProfileAssignmentId: "windowsAutopilotDeploymentProfileAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileId/assignments/windowsAutopilotDeploymentProfileAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentID(v.Input)
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

		if actual.WindowsAutopilotDeploymentProfileAssignmentId != v.Expected.WindowsAutopilotDeploymentProfileAssignmentId {
			t.Fatalf("Expected %q but got %q for WindowsAutopilotDeploymentProfileAssignmentId", v.Expected.WindowsAutopilotDeploymentProfileAssignmentId, actual.WindowsAutopilotDeploymentProfileAssignmentId)
		}

	}
}

func TestParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId
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
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeS/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeS/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeId/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileId/assignments/windowsAutopilotDeploymentProfileAssignmentId",
			Expected: &DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId{
				WindowsAutopilotDeploymentProfileId:           "windowsAutopilotDeploymentProfileId",
				WindowsAutopilotDeploymentProfileAssignmentId: "windowsAutopilotDeploymentProfileAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/windowsAutopilotDeploymentProfiles/windowsAutopilotDeploymentProfileId/assignments/windowsAutopilotDeploymentProfileAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeS/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeId/aSsIgNmEnTs/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeAsSiGnMeNtId",
			Expected: &DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId{
				WindowsAutopilotDeploymentProfileId:           "wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeId",
				WindowsAutopilotDeploymentProfileAssignmentId: "wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeS/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeId/aSsIgNmEnTs/wInDoWsAuToPiLoTdEpLoYmEnTpRoFiLeAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentIDInsensitively(v.Input)
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

		if actual.WindowsAutopilotDeploymentProfileAssignmentId != v.Expected.WindowsAutopilotDeploymentProfileAssignmentId {
			t.Fatalf("Expected %q but got %q for WindowsAutopilotDeploymentProfileAssignmentId", v.Expected.WindowsAutopilotDeploymentProfileAssignmentId, actual.WindowsAutopilotDeploymentProfileAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId(t *testing.T) {
	segments := DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
