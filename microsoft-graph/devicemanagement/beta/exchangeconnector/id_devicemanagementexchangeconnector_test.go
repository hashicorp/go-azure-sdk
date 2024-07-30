package exchangeconnector

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementExchangeConnectorId{}

func TestNewDeviceManagementExchangeConnectorID(t *testing.T) {
	id := NewDeviceManagementExchangeConnectorID("deviceManagementExchangeConnectorIdValue")

	if id.DeviceManagementExchangeConnectorId != "deviceManagementExchangeConnectorIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementExchangeConnectorId'", id.DeviceManagementExchangeConnectorId, "deviceManagementExchangeConnectorIdValue")
	}
}

func TestFormatDeviceManagementExchangeConnectorID(t *testing.T) {
	actual := NewDeviceManagementExchangeConnectorID("deviceManagementExchangeConnectorIdValue").ID()
	expected := "/deviceManagement/exchangeConnectors/deviceManagementExchangeConnectorIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementExchangeConnectorID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementExchangeConnectorId
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
			Input: "/deviceManagement/exchangeConnectors",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/exchangeConnectors/deviceManagementExchangeConnectorIdValue",
			Expected: &DeviceManagementExchangeConnectorId{
				DeviceManagementExchangeConnectorId: "deviceManagementExchangeConnectorIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/exchangeConnectors/deviceManagementExchangeConnectorIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementExchangeConnectorID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementExchangeConnectorId != v.Expected.DeviceManagementExchangeConnectorId {
			t.Fatalf("Expected %q but got %q for DeviceManagementExchangeConnectorId", v.Expected.DeviceManagementExchangeConnectorId, actual.DeviceManagementExchangeConnectorId)
		}

	}
}

func TestParseDeviceManagementExchangeConnectorIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementExchangeConnectorId
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
			Input: "/deviceManagement/exchangeConnectors",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eXcHaNgEcOnNeCtOrS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/exchangeConnectors/deviceManagementExchangeConnectorIdValue",
			Expected: &DeviceManagementExchangeConnectorId{
				DeviceManagementExchangeConnectorId: "deviceManagementExchangeConnectorIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/exchangeConnectors/deviceManagementExchangeConnectorIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eXcHaNgEcOnNeCtOrS/dEvIcEmAnAgEmEnTeXcHaNgEcOnNeCtOrIdVaLuE",
			Expected: &DeviceManagementExchangeConnectorId{
				DeviceManagementExchangeConnectorId: "dEvIcEmAnAgEmEnTeXcHaNgEcOnNeCtOrIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eXcHaNgEcOnNeCtOrS/dEvIcEmAnAgEmEnTeXcHaNgEcOnNeCtOrIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementExchangeConnectorIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementExchangeConnectorId != v.Expected.DeviceManagementExchangeConnectorId {
			t.Fatalf("Expected %q but got %q for DeviceManagementExchangeConnectorId", v.Expected.DeviceManagementExchangeConnectorId, actual.DeviceManagementExchangeConnectorId)
		}

	}
}

func TestSegmentsForDeviceManagementExchangeConnectorId(t *testing.T) {
	segments := DeviceManagementExchangeConnectorId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementExchangeConnectorId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
