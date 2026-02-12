package iotcentrals

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ModuleComponentId{}

func TestNewModuleComponentID(t *testing.T) {
	id := NewModuleComponentID("https://endpoint-url.example.com", "deviceId", "moduleName", "componentName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.DeviceId != "deviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceId'", id.DeviceId, "deviceId")
	}

	if id.ModuleName != "moduleName" {
		t.Fatalf("Expected %q but got %q for Segment 'ModuleName'", id.ModuleName, "moduleName")
	}

	if id.ComponentName != "componentName" {
		t.Fatalf("Expected %q but got %q for Segment 'ComponentName'", id.ComponentName, "componentName")
	}
}

func TestFormatModuleComponentID(t *testing.T) {
	actual := NewModuleComponentID("https://endpoint-url.example.com", "deviceId", "moduleName", "componentName").ID()
	expected := "https://endpoint-url.example.com/devices/deviceId/modules/moduleName/components/componentName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseModuleComponentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ModuleComponentId
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
			Input: "https://endpoint-url.example.com/devices/deviceId/modules",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices/deviceId/modules/moduleName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices/deviceId/modules/moduleName/components",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/devices/deviceId/modules/moduleName/components/componentName",
			Expected: &ModuleComponentId{
				BaseURI:       "https://endpoint-url.example.com",
				DeviceId:      "deviceId",
				ModuleName:    "moduleName",
				ComponentName: "componentName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/devices/deviceId/modules/moduleName/components/componentName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseModuleComponentID(v.Input)
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

		if actual.ModuleName != v.Expected.ModuleName {
			t.Fatalf("Expected %q but got %q for ModuleName", v.Expected.ModuleName, actual.ModuleName)
		}

		if actual.ComponentName != v.Expected.ComponentName {
			t.Fatalf("Expected %q but got %q for ComponentName", v.Expected.ComponentName, actual.ComponentName)
		}

	}
}

func TestParseModuleComponentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ModuleComponentId
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
			Input: "https://endpoint-url.example.com/devices/deviceId/modules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs/dEvIcEiD/mOdUlEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices/deviceId/modules/moduleName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs/dEvIcEiD/mOdUlEs/mOdUlEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices/deviceId/modules/moduleName/components",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs/dEvIcEiD/mOdUlEs/mOdUlEnAmE/cOmPoNeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/devices/deviceId/modules/moduleName/components/componentName",
			Expected: &ModuleComponentId{
				BaseURI:       "https://endpoint-url.example.com",
				DeviceId:      "deviceId",
				ModuleName:    "moduleName",
				ComponentName: "componentName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/devices/deviceId/modules/moduleName/components/componentName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs/dEvIcEiD/mOdUlEs/mOdUlEnAmE/cOmPoNeNtS/cOmPoNeNtNaMe",
			Expected: &ModuleComponentId{
				BaseURI:       "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				DeviceId:      "dEvIcEiD",
				ModuleName:    "mOdUlEnAmE",
				ComponentName: "cOmPoNeNtNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs/dEvIcEiD/mOdUlEs/mOdUlEnAmE/cOmPoNeNtS/cOmPoNeNtNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseModuleComponentIDInsensitively(v.Input)
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

		if actual.ModuleName != v.Expected.ModuleName {
			t.Fatalf("Expected %q but got %q for ModuleName", v.Expected.ModuleName, actual.ModuleName)
		}

		if actual.ComponentName != v.Expected.ComponentName {
			t.Fatalf("Expected %q but got %q for ComponentName", v.Expected.ComponentName, actual.ComponentName)
		}

	}
}

func TestSegmentsForModuleComponentId(t *testing.T) {
	segments := ModuleComponentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ModuleComponentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
