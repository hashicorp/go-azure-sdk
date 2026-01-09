package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMobileThreatDefenseConnectorId{}

func TestNewDeviceManagementMobileThreatDefenseConnectorID(t *testing.T) {
	id := NewDeviceManagementMobileThreatDefenseConnectorID("mobileThreatDefenseConnectorId")

	if id.MobileThreatDefenseConnectorId != "mobileThreatDefenseConnectorId" {
		t.Fatalf("Expected %q but got %q for Segment 'MobileThreatDefenseConnectorId'", id.MobileThreatDefenseConnectorId, "mobileThreatDefenseConnectorId")
	}
}

func TestFormatDeviceManagementMobileThreatDefenseConnectorID(t *testing.T) {
	actual := NewDeviceManagementMobileThreatDefenseConnectorID("mobileThreatDefenseConnectorId").ID()
	expected := "/deviceManagement/mobileThreatDefenseConnectors/mobileThreatDefenseConnectorId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementMobileThreatDefenseConnectorID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMobileThreatDefenseConnectorId
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
			Input: "/deviceManagement/mobileThreatDefenseConnectors",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/mobileThreatDefenseConnectors/mobileThreatDefenseConnectorId",
			Expected: &DeviceManagementMobileThreatDefenseConnectorId{
				MobileThreatDefenseConnectorId: "mobileThreatDefenseConnectorId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/mobileThreatDefenseConnectors/mobileThreatDefenseConnectorId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMobileThreatDefenseConnectorID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MobileThreatDefenseConnectorId != v.Expected.MobileThreatDefenseConnectorId {
			t.Fatalf("Expected %q but got %q for MobileThreatDefenseConnectorId", v.Expected.MobileThreatDefenseConnectorId, actual.MobileThreatDefenseConnectorId)
		}

	}
}

func TestParseDeviceManagementMobileThreatDefenseConnectorIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMobileThreatDefenseConnectorId
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
			Input: "/deviceManagement/mobileThreatDefenseConnectors",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mObIlEtHrEaTdEfEnSeCoNnEcToRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/mobileThreatDefenseConnectors/mobileThreatDefenseConnectorId",
			Expected: &DeviceManagementMobileThreatDefenseConnectorId{
				MobileThreatDefenseConnectorId: "mobileThreatDefenseConnectorId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/mobileThreatDefenseConnectors/mobileThreatDefenseConnectorId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mObIlEtHrEaTdEfEnSeCoNnEcToRs/mObIlEtHrEaTdEfEnSeCoNnEcToRiD",
			Expected: &DeviceManagementMobileThreatDefenseConnectorId{
				MobileThreatDefenseConnectorId: "mObIlEtHrEaTdEfEnSeCoNnEcToRiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mObIlEtHrEaTdEfEnSeCoNnEcToRs/mObIlEtHrEaTdEfEnSeCoNnEcToRiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMobileThreatDefenseConnectorIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MobileThreatDefenseConnectorId != v.Expected.MobileThreatDefenseConnectorId {
			t.Fatalf("Expected %q but got %q for MobileThreatDefenseConnectorId", v.Expected.MobileThreatDefenseConnectorId, actual.MobileThreatDefenseConnectorId)
		}

	}
}

func TestSegmentsForDeviceManagementMobileThreatDefenseConnectorId(t *testing.T) {
	segments := DeviceManagementMobileThreatDefenseConnectorId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementMobileThreatDefenseConnectorId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
