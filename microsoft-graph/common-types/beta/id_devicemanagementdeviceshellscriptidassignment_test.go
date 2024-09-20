package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceShellScriptIdAssignmentId{}

func TestNewDeviceManagementDeviceShellScriptIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementDeviceShellScriptIdAssignmentID("deviceShellScriptId", "deviceManagementScriptAssignmentId")

	if id.DeviceShellScriptId != "deviceShellScriptId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceShellScriptId'", id.DeviceShellScriptId, "deviceShellScriptId")
	}

	if id.DeviceManagementScriptAssignmentId != "deviceManagementScriptAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementScriptAssignmentId'", id.DeviceManagementScriptAssignmentId, "deviceManagementScriptAssignmentId")
	}
}

func TestFormatDeviceManagementDeviceShellScriptIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementDeviceShellScriptIdAssignmentID("deviceShellScriptId", "deviceManagementScriptAssignmentId").ID()
	expected := "/deviceManagement/deviceShellScripts/deviceShellScriptId/assignments/deviceManagementScriptAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceShellScriptIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceShellScriptIdAssignmentId
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
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/assignments/deviceManagementScriptAssignmentId",
			Expected: &DeviceManagementDeviceShellScriptIdAssignmentId{
				DeviceShellScriptId:                "deviceShellScriptId",
				DeviceManagementScriptAssignmentId: "deviceManagementScriptAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/assignments/deviceManagementScriptAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceShellScriptIdAssignmentID(v.Input)
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

		if actual.DeviceManagementScriptAssignmentId != v.Expected.DeviceManagementScriptAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptAssignmentId", v.Expected.DeviceManagementScriptAssignmentId, actual.DeviceManagementScriptAssignmentId)
		}

	}
}

func TestParseDeviceManagementDeviceShellScriptIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceShellScriptIdAssignmentId
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
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEsHeLlScRiPtS/dEvIcEsHeLlScRiPtId/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/assignments/deviceManagementScriptAssignmentId",
			Expected: &DeviceManagementDeviceShellScriptIdAssignmentId{
				DeviceShellScriptId:                "deviceShellScriptId",
				DeviceManagementScriptAssignmentId: "deviceManagementScriptAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/assignments/deviceManagementScriptAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEsHeLlScRiPtS/dEvIcEsHeLlScRiPtId/aSsIgNmEnTs/dEvIcEmAnAgEmEnTsCrIpTaSsIgNmEnTiD",
			Expected: &DeviceManagementDeviceShellScriptIdAssignmentId{
				DeviceShellScriptId:                "dEvIcEsHeLlScRiPtId",
				DeviceManagementScriptAssignmentId: "dEvIcEmAnAgEmEnTsCrIpTaSsIgNmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEsHeLlScRiPtS/dEvIcEsHeLlScRiPtId/aSsIgNmEnTs/dEvIcEmAnAgEmEnTsCrIpTaSsIgNmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceShellScriptIdAssignmentIDInsensitively(v.Input)
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

		if actual.DeviceManagementScriptAssignmentId != v.Expected.DeviceManagementScriptAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptAssignmentId", v.Expected.DeviceManagementScriptAssignmentId, actual.DeviceManagementScriptAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceShellScriptIdAssignmentId(t *testing.T) {
	segments := DeviceManagementDeviceShellScriptIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceShellScriptIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
