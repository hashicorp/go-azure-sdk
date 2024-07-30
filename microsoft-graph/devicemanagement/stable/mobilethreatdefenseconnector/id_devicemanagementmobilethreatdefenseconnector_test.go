package mobilethreatdefenseconnector

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMobileThreatDefenseConnectorId{}

func TestNewDeviceManagementMobileThreatDefenseConnectorID(t *testing.T) {
	id := NewDeviceManagementMobileThreatDefenseConnectorID("mobileThreatDefenseConnectorIdValue")

	if id.MobileThreatDefenseConnectorId != "mobileThreatDefenseConnectorIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MobileThreatDefenseConnectorId'", id.MobileThreatDefenseConnectorId, "mobileThreatDefenseConnectorIdValue")
	}
}

func TestFormatDeviceManagementMobileThreatDefenseConnectorID(t *testing.T) {
	actual := NewDeviceManagementMobileThreatDefenseConnectorID("mobileThreatDefenseConnectorIdValue").ID()
	expected := "/deviceManagement/mobileThreatDefenseConnectors/mobileThreatDefenseConnectorIdValue"
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
			Input: "/deviceManagement/mobileThreatDefenseConnectors/mobileThreatDefenseConnectorIdValue",
			Expected: &DeviceManagementMobileThreatDefenseConnectorId{
				MobileThreatDefenseConnectorId: "mobileThreatDefenseConnectorIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/mobileThreatDefenseConnectors/mobileThreatDefenseConnectorIdValue/extra",
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
			Input: "/deviceManagement/mobileThreatDefenseConnectors/mobileThreatDefenseConnectorIdValue",
			Expected: &DeviceManagementMobileThreatDefenseConnectorId{
				MobileThreatDefenseConnectorId: "mobileThreatDefenseConnectorIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/mobileThreatDefenseConnectors/mobileThreatDefenseConnectorIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mObIlEtHrEaTdEfEnSeCoNnEcToRs/mObIlEtHrEaTdEfEnSeCoNnEcToRiDvAlUe",
			Expected: &DeviceManagementMobileThreatDefenseConnectorId{
				MobileThreatDefenseConnectorId: "mObIlEtHrEaTdEfEnSeCoNnEcToRiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mObIlEtHrEaTdEfEnSeCoNnEcToRs/mObIlEtHrEaTdEfEnSeCoNnEcToRiDvAlUe/extra",
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
