package iotcentrals

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ComponentTelemetryId{}

func TestNewComponentTelemetryID(t *testing.T) {
	id := NewComponentTelemetryID("https://endpoint-url.example.com", "deviceId", "componentName", "telemetryName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.DeviceId != "deviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceId'", id.DeviceId, "deviceId")
	}

	if id.ComponentName != "componentName" {
		t.Fatalf("Expected %q but got %q for Segment 'ComponentName'", id.ComponentName, "componentName")
	}

	if id.TelemetryName != "telemetryName" {
		t.Fatalf("Expected %q but got %q for Segment 'TelemetryName'", id.TelemetryName, "telemetryName")
	}
}

func TestFormatComponentTelemetryID(t *testing.T) {
	actual := NewComponentTelemetryID("https://endpoint-url.example.com", "deviceId", "componentName", "telemetryName").ID()
	expected := "https://endpoint-url.example.com/devices/deviceId/components/componentName/telemetry/telemetryName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseComponentTelemetryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ComponentTelemetryId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices/deviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices/deviceId/components",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices/deviceId/components/componentName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices/deviceId/components/componentName/telemetry",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/devices/deviceId/components/componentName/telemetry/telemetryName",
			Expected: &ComponentTelemetryId{
				BaseURI:       "https://endpoint-url.example.com",
				DeviceId:      "deviceId",
				ComponentName: "componentName",
				TelemetryName: "telemetryName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/devices/deviceId/components/componentName/telemetry/telemetryName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseComponentTelemetryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.DeviceId != v.Expected.DeviceId {
			t.Fatalf("Expected %q but got %q for DeviceId", v.Expected.DeviceId, actual.DeviceId)
		}

		if actual.ComponentName != v.Expected.ComponentName {
			t.Fatalf("Expected %q but got %q for ComponentName", v.Expected.ComponentName, actual.ComponentName)
		}

		if actual.TelemetryName != v.Expected.TelemetryName {
			t.Fatalf("Expected %q but got %q for TelemetryName", v.Expected.TelemetryName, actual.TelemetryName)
		}

	}
}

func TestParseComponentTelemetryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ComponentTelemetryId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices/deviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs/dEvIcEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices/deviceId/components",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs/dEvIcEiD/cOmPoNeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices/deviceId/components/componentName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs/dEvIcEiD/cOmPoNeNtS/cOmPoNeNtNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices/deviceId/components/componentName/telemetry",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs/dEvIcEiD/cOmPoNeNtS/cOmPoNeNtNaMe/tElEmEtRy",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/devices/deviceId/components/componentName/telemetry/telemetryName",
			Expected: &ComponentTelemetryId{
				BaseURI:       "https://endpoint-url.example.com",
				DeviceId:      "deviceId",
				ComponentName: "componentName",
				TelemetryName: "telemetryName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/devices/deviceId/components/componentName/telemetry/telemetryName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs/dEvIcEiD/cOmPoNeNtS/cOmPoNeNtNaMe/tElEmEtRy/tElEmEtRyNaMe",
			Expected: &ComponentTelemetryId{
				BaseURI:       "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				DeviceId:      "dEvIcEiD",
				ComponentName: "cOmPoNeNtNaMe",
				TelemetryName: "tElEmEtRyNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs/dEvIcEiD/cOmPoNeNtS/cOmPoNeNtNaMe/tElEmEtRy/tElEmEtRyNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseComponentTelemetryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.DeviceId != v.Expected.DeviceId {
			t.Fatalf("Expected %q but got %q for DeviceId", v.Expected.DeviceId, actual.DeviceId)
		}

		if actual.ComponentName != v.Expected.ComponentName {
			t.Fatalf("Expected %q but got %q for ComponentName", v.Expected.ComponentName, actual.ComponentName)
		}

		if actual.TelemetryName != v.Expected.TelemetryName {
			t.Fatalf("Expected %q but got %q for TelemetryName", v.Expected.TelemetryName, actual.TelemetryName)
		}

	}
}

func TestSegmentsForComponentTelemetryId(t *testing.T) {
	segments := ComponentTelemetryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ComponentTelemetryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
