package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceShellScriptIdDeviceRunStateId{}

func TestNewDeviceManagementDeviceShellScriptIdDeviceRunStateID(t *testing.T) {
	id := NewDeviceManagementDeviceShellScriptIdDeviceRunStateID("deviceShellScriptId", "deviceManagementScriptDeviceStateId")

	if id.DeviceShellScriptId != "deviceShellScriptId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceShellScriptId'", id.DeviceShellScriptId, "deviceShellScriptId")
	}

	if id.DeviceManagementScriptDeviceStateId != "deviceManagementScriptDeviceStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementScriptDeviceStateId'", id.DeviceManagementScriptDeviceStateId, "deviceManagementScriptDeviceStateId")
	}
}

func TestFormatDeviceManagementDeviceShellScriptIdDeviceRunStateID(t *testing.T) {
	actual := NewDeviceManagementDeviceShellScriptIdDeviceRunStateID("deviceShellScriptId", "deviceManagementScriptDeviceStateId").ID()
	expected := "/deviceManagement/deviceShellScripts/deviceShellScriptId/deviceRunStates/deviceManagementScriptDeviceStateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceShellScriptIdDeviceRunStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceShellScriptIdDeviceRunStateId
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
			// Incomplete URI
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/deviceRunStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/deviceRunStates/deviceManagementScriptDeviceStateId",
			Expected: &DeviceManagementDeviceShellScriptIdDeviceRunStateId{
				DeviceShellScriptId:                 "deviceShellScriptId",
				DeviceManagementScriptDeviceStateId: "deviceManagementScriptDeviceStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/deviceRunStates/deviceManagementScriptDeviceStateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceShellScriptIdDeviceRunStateID(v.Input)
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

		if actual.DeviceManagementScriptDeviceStateId != v.Expected.DeviceManagementScriptDeviceStateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptDeviceStateId", v.Expected.DeviceManagementScriptDeviceStateId, actual.DeviceManagementScriptDeviceStateId)
		}

	}
}

func TestParseDeviceManagementDeviceShellScriptIdDeviceRunStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceShellScriptIdDeviceRunStateId
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
			// Incomplete URI
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEsHeLlScRiPtS/dEvIcEsHeLlScRiPtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/deviceRunStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEsHeLlScRiPtS/dEvIcEsHeLlScRiPtId/dEvIcErUnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/deviceRunStates/deviceManagementScriptDeviceStateId",
			Expected: &DeviceManagementDeviceShellScriptIdDeviceRunStateId{
				DeviceShellScriptId:                 "deviceShellScriptId",
				DeviceManagementScriptDeviceStateId: "deviceManagementScriptDeviceStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/deviceRunStates/deviceManagementScriptDeviceStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEsHeLlScRiPtS/dEvIcEsHeLlScRiPtId/dEvIcErUnStAtEs/dEvIcEmAnAgEmEnTsCrIpTdEvIcEsTaTeId",
			Expected: &DeviceManagementDeviceShellScriptIdDeviceRunStateId{
				DeviceShellScriptId:                 "dEvIcEsHeLlScRiPtId",
				DeviceManagementScriptDeviceStateId: "dEvIcEmAnAgEmEnTsCrIpTdEvIcEsTaTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEsHeLlScRiPtS/dEvIcEsHeLlScRiPtId/dEvIcErUnStAtEs/dEvIcEmAnAgEmEnTsCrIpTdEvIcEsTaTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceShellScriptIdDeviceRunStateIDInsensitively(v.Input)
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

		if actual.DeviceManagementScriptDeviceStateId != v.Expected.DeviceManagementScriptDeviceStateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptDeviceStateId", v.Expected.DeviceManagementScriptDeviceStateId, actual.DeviceManagementScriptDeviceStateId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceShellScriptIdDeviceRunStateId(t *testing.T) {
	segments := DeviceManagementDeviceShellScriptIdDeviceRunStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceShellScriptIdDeviceRunStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
