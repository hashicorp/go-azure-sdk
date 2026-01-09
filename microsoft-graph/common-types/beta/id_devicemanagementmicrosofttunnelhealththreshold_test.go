package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMicrosoftTunnelHealthThresholdId{}

func TestNewDeviceManagementMicrosoftTunnelHealthThresholdID(t *testing.T) {
	id := NewDeviceManagementMicrosoftTunnelHealthThresholdID("microsoftTunnelHealthThresholdId")

	if id.MicrosoftTunnelHealthThresholdId != "microsoftTunnelHealthThresholdId" {
		t.Fatalf("Expected %q but got %q for Segment 'MicrosoftTunnelHealthThresholdId'", id.MicrosoftTunnelHealthThresholdId, "microsoftTunnelHealthThresholdId")
	}
}

func TestFormatDeviceManagementMicrosoftTunnelHealthThresholdID(t *testing.T) {
	actual := NewDeviceManagementMicrosoftTunnelHealthThresholdID("microsoftTunnelHealthThresholdId").ID()
	expected := "/deviceManagement/microsoftTunnelHealthThresholds/microsoftTunnelHealthThresholdId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementMicrosoftTunnelHealthThresholdID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMicrosoftTunnelHealthThresholdId
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
			Input: "/deviceManagement/microsoftTunnelHealthThresholds",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/microsoftTunnelHealthThresholds/microsoftTunnelHealthThresholdId",
			Expected: &DeviceManagementMicrosoftTunnelHealthThresholdId{
				MicrosoftTunnelHealthThresholdId: "microsoftTunnelHealthThresholdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/microsoftTunnelHealthThresholds/microsoftTunnelHealthThresholdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMicrosoftTunnelHealthThresholdID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MicrosoftTunnelHealthThresholdId != v.Expected.MicrosoftTunnelHealthThresholdId {
			t.Fatalf("Expected %q but got %q for MicrosoftTunnelHealthThresholdId", v.Expected.MicrosoftTunnelHealthThresholdId, actual.MicrosoftTunnelHealthThresholdId)
		}

	}
}

func TestParseDeviceManagementMicrosoftTunnelHealthThresholdIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMicrosoftTunnelHealthThresholdId
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
			Input: "/deviceManagement/microsoftTunnelHealthThresholds",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mIcRoSoFtTuNnElHeAlThThReShOlDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/microsoftTunnelHealthThresholds/microsoftTunnelHealthThresholdId",
			Expected: &DeviceManagementMicrosoftTunnelHealthThresholdId{
				MicrosoftTunnelHealthThresholdId: "microsoftTunnelHealthThresholdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/microsoftTunnelHealthThresholds/microsoftTunnelHealthThresholdId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mIcRoSoFtTuNnElHeAlThThReShOlDs/mIcRoSoFtTuNnElHeAlThThReShOlDiD",
			Expected: &DeviceManagementMicrosoftTunnelHealthThresholdId{
				MicrosoftTunnelHealthThresholdId: "mIcRoSoFtTuNnElHeAlThThReShOlDiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mIcRoSoFtTuNnElHeAlThThReShOlDs/mIcRoSoFtTuNnElHeAlThThReShOlDiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMicrosoftTunnelHealthThresholdIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MicrosoftTunnelHealthThresholdId != v.Expected.MicrosoftTunnelHealthThresholdId {
			t.Fatalf("Expected %q but got %q for MicrosoftTunnelHealthThresholdId", v.Expected.MicrosoftTunnelHealthThresholdId, actual.MicrosoftTunnelHealthThresholdId)
		}

	}
}

func TestSegmentsForDeviceManagementMicrosoftTunnelHealthThresholdId(t *testing.T) {
	segments := DeviceManagementMicrosoftTunnelHealthThresholdId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementMicrosoftTunnelHealthThresholdId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
