package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementImportedWindowsAutopilotDeviceIdentityId{}

func TestNewDeviceManagementImportedWindowsAutopilotDeviceIdentityID(t *testing.T) {
	id := NewDeviceManagementImportedWindowsAutopilotDeviceIdentityID("importedWindowsAutopilotDeviceIdentityId")

	if id.ImportedWindowsAutopilotDeviceIdentityId != "importedWindowsAutopilotDeviceIdentityId" {
		t.Fatalf("Expected %q but got %q for Segment 'ImportedWindowsAutopilotDeviceIdentityId'", id.ImportedWindowsAutopilotDeviceIdentityId, "importedWindowsAutopilotDeviceIdentityId")
	}
}

func TestFormatDeviceManagementImportedWindowsAutopilotDeviceIdentityID(t *testing.T) {
	actual := NewDeviceManagementImportedWindowsAutopilotDeviceIdentityID("importedWindowsAutopilotDeviceIdentityId").ID()
	expected := "/deviceManagement/importedWindowsAutopilotDeviceIdentities/importedWindowsAutopilotDeviceIdentityId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementImportedWindowsAutopilotDeviceIdentityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementImportedWindowsAutopilotDeviceIdentityId
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
			Input: "/deviceManagement/importedWindowsAutopilotDeviceIdentities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/importedWindowsAutopilotDeviceIdentities/importedWindowsAutopilotDeviceIdentityId",
			Expected: &DeviceManagementImportedWindowsAutopilotDeviceIdentityId{
				ImportedWindowsAutopilotDeviceIdentityId: "importedWindowsAutopilotDeviceIdentityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/importedWindowsAutopilotDeviceIdentities/importedWindowsAutopilotDeviceIdentityId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementImportedWindowsAutopilotDeviceIdentityID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ImportedWindowsAutopilotDeviceIdentityId != v.Expected.ImportedWindowsAutopilotDeviceIdentityId {
			t.Fatalf("Expected %q but got %q for ImportedWindowsAutopilotDeviceIdentityId", v.Expected.ImportedWindowsAutopilotDeviceIdentityId, actual.ImportedWindowsAutopilotDeviceIdentityId)
		}

	}
}

func TestParseDeviceManagementImportedWindowsAutopilotDeviceIdentityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementImportedWindowsAutopilotDeviceIdentityId
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
			Input: "/deviceManagement/importedWindowsAutopilotDeviceIdentities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iMpOrTeDwInDoWsAuToPiLoTdEvIcEiDeNtItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/importedWindowsAutopilotDeviceIdentities/importedWindowsAutopilotDeviceIdentityId",
			Expected: &DeviceManagementImportedWindowsAutopilotDeviceIdentityId{
				ImportedWindowsAutopilotDeviceIdentityId: "importedWindowsAutopilotDeviceIdentityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/importedWindowsAutopilotDeviceIdentities/importedWindowsAutopilotDeviceIdentityId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iMpOrTeDwInDoWsAuToPiLoTdEvIcEiDeNtItIeS/iMpOrTeDwInDoWsAuToPiLoTdEvIcEiDeNtItYiD",
			Expected: &DeviceManagementImportedWindowsAutopilotDeviceIdentityId{
				ImportedWindowsAutopilotDeviceIdentityId: "iMpOrTeDwInDoWsAuToPiLoTdEvIcEiDeNtItYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iMpOrTeDwInDoWsAuToPiLoTdEvIcEiDeNtItIeS/iMpOrTeDwInDoWsAuToPiLoTdEvIcEiDeNtItYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementImportedWindowsAutopilotDeviceIdentityIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ImportedWindowsAutopilotDeviceIdentityId != v.Expected.ImportedWindowsAutopilotDeviceIdentityId {
			t.Fatalf("Expected %q but got %q for ImportedWindowsAutopilotDeviceIdentityId", v.Expected.ImportedWindowsAutopilotDeviceIdentityId, actual.ImportedWindowsAutopilotDeviceIdentityId)
		}

	}
}

func TestSegmentsForDeviceManagementImportedWindowsAutopilotDeviceIdentityId(t *testing.T) {
	segments := DeviceManagementImportedWindowsAutopilotDeviceIdentityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementImportedWindowsAutopilotDeviceIdentityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
