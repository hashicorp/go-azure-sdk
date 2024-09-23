package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceShellScriptIdUserRunStateId{}

func TestNewDeviceManagementDeviceShellScriptIdUserRunStateID(t *testing.T) {
	id := NewDeviceManagementDeviceShellScriptIdUserRunStateID("deviceShellScriptId", "deviceManagementScriptUserStateId")

	if id.DeviceShellScriptId != "deviceShellScriptId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceShellScriptId'", id.DeviceShellScriptId, "deviceShellScriptId")
	}

	if id.DeviceManagementScriptUserStateId != "deviceManagementScriptUserStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementScriptUserStateId'", id.DeviceManagementScriptUserStateId, "deviceManagementScriptUserStateId")
	}
}

func TestFormatDeviceManagementDeviceShellScriptIdUserRunStateID(t *testing.T) {
	actual := NewDeviceManagementDeviceShellScriptIdUserRunStateID("deviceShellScriptId", "deviceManagementScriptUserStateId").ID()
	expected := "/deviceManagement/deviceShellScripts/deviceShellScriptId/userRunStates/deviceManagementScriptUserStateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceShellScriptIdUserRunStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceShellScriptIdUserRunStateId
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
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/userRunStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/userRunStates/deviceManagementScriptUserStateId",
			Expected: &DeviceManagementDeviceShellScriptIdUserRunStateId{
				DeviceShellScriptId:               "deviceShellScriptId",
				DeviceManagementScriptUserStateId: "deviceManagementScriptUserStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/userRunStates/deviceManagementScriptUserStateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceShellScriptIdUserRunStateID(v.Input)
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

		if actual.DeviceManagementScriptUserStateId != v.Expected.DeviceManagementScriptUserStateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptUserStateId", v.Expected.DeviceManagementScriptUserStateId, actual.DeviceManagementScriptUserStateId)
		}

	}
}

func TestParseDeviceManagementDeviceShellScriptIdUserRunStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceShellScriptIdUserRunStateId
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
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/userRunStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEsHeLlScRiPtS/dEvIcEsHeLlScRiPtId/uSeRrUnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/userRunStates/deviceManagementScriptUserStateId",
			Expected: &DeviceManagementDeviceShellScriptIdUserRunStateId{
				DeviceShellScriptId:               "deviceShellScriptId",
				DeviceManagementScriptUserStateId: "deviceManagementScriptUserStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/userRunStates/deviceManagementScriptUserStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEsHeLlScRiPtS/dEvIcEsHeLlScRiPtId/uSeRrUnStAtEs/dEvIcEmAnAgEmEnTsCrIpTuSeRsTaTeId",
			Expected: &DeviceManagementDeviceShellScriptIdUserRunStateId{
				DeviceShellScriptId:               "dEvIcEsHeLlScRiPtId",
				DeviceManagementScriptUserStateId: "dEvIcEmAnAgEmEnTsCrIpTuSeRsTaTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEsHeLlScRiPtS/dEvIcEsHeLlScRiPtId/uSeRrUnStAtEs/dEvIcEmAnAgEmEnTsCrIpTuSeRsTaTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceShellScriptIdUserRunStateIDInsensitively(v.Input)
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

		if actual.DeviceManagementScriptUserStateId != v.Expected.DeviceManagementScriptUserStateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptUserStateId", v.Expected.DeviceManagementScriptUserStateId, actual.DeviceManagementScriptUserStateId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceShellScriptIdUserRunStateId(t *testing.T) {
	segments := DeviceManagementDeviceShellScriptIdUserRunStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceShellScriptIdUserRunStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
