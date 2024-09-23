package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceShellScriptId{}

func TestNewDeviceManagementDeviceShellScriptID(t *testing.T) {
	id := NewDeviceManagementDeviceShellScriptID("deviceShellScriptId")

	if id.DeviceShellScriptId != "deviceShellScriptId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceShellScriptId'", id.DeviceShellScriptId, "deviceShellScriptId")
	}
}

func TestFormatDeviceManagementDeviceShellScriptID(t *testing.T) {
	actual := NewDeviceManagementDeviceShellScriptID("deviceShellScriptId").ID()
	expected := "/deviceManagement/deviceShellScripts/deviceShellScriptId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceShellScriptID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceShellScriptId
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
			Input: "/deviceManagement/deviceShellScripts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId",
			Expected: &DeviceManagementDeviceShellScriptId{
				DeviceShellScriptId: "deviceShellScriptId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceShellScriptID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceShellScriptId != v.Expected.DeviceShellScriptId {
			t.Fatalf("Expected %q but got %q for DeviceShellScriptId", v.Expected.DeviceShellScriptId, actual.DeviceShellScriptId)
		}

	}
}

func TestParseDeviceManagementDeviceShellScriptIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceShellScriptId
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
			Input: "/deviceManagement/deviceShellScripts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEsHeLlScRiPtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId",
			Expected: &DeviceManagementDeviceShellScriptId{
				DeviceShellScriptId: "deviceShellScriptId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEsHeLlScRiPtS/dEvIcEsHeLlScRiPtId",
			Expected: &DeviceManagementDeviceShellScriptId{
				DeviceShellScriptId: "dEvIcEsHeLlScRiPtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEsHeLlScRiPtS/dEvIcEsHeLlScRiPtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceShellScriptIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceShellScriptId != v.Expected.DeviceShellScriptId {
			t.Fatalf("Expected %q but got %q for DeviceShellScriptId", v.Expected.DeviceShellScriptId, actual.DeviceShellScriptId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceShellScriptId(t *testing.T) {
	segments := DeviceManagementDeviceShellScriptId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceShellScriptId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
