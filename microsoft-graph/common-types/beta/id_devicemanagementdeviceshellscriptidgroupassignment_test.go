package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceShellScriptIdGroupAssignmentId{}

func TestNewDeviceManagementDeviceShellScriptIdGroupAssignmentID(t *testing.T) {
	id := NewDeviceManagementDeviceShellScriptIdGroupAssignmentID("deviceShellScriptId", "deviceManagementScriptGroupAssignmentId")

	if id.DeviceShellScriptId != "deviceShellScriptId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceShellScriptId'", id.DeviceShellScriptId, "deviceShellScriptId")
	}

	if id.DeviceManagementScriptGroupAssignmentId != "deviceManagementScriptGroupAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementScriptGroupAssignmentId'", id.DeviceManagementScriptGroupAssignmentId, "deviceManagementScriptGroupAssignmentId")
	}
}

func TestFormatDeviceManagementDeviceShellScriptIdGroupAssignmentID(t *testing.T) {
	actual := NewDeviceManagementDeviceShellScriptIdGroupAssignmentID("deviceShellScriptId", "deviceManagementScriptGroupAssignmentId").ID()
	expected := "/deviceManagement/deviceShellScripts/deviceShellScriptId/groupAssignments/deviceManagementScriptGroupAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceShellScriptIdGroupAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceShellScriptIdGroupAssignmentId
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
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/groupAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/groupAssignments/deviceManagementScriptGroupAssignmentId",
			Expected: &DeviceManagementDeviceShellScriptIdGroupAssignmentId{
				DeviceShellScriptId:                     "deviceShellScriptId",
				DeviceManagementScriptGroupAssignmentId: "deviceManagementScriptGroupAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/groupAssignments/deviceManagementScriptGroupAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceShellScriptIdGroupAssignmentID(v.Input)
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

		if actual.DeviceManagementScriptGroupAssignmentId != v.Expected.DeviceManagementScriptGroupAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptGroupAssignmentId", v.Expected.DeviceManagementScriptGroupAssignmentId, actual.DeviceManagementScriptGroupAssignmentId)
		}

	}
}

func TestParseDeviceManagementDeviceShellScriptIdGroupAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceShellScriptIdGroupAssignmentId
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
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/groupAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEsHeLlScRiPtS/dEvIcEsHeLlScRiPtId/gRoUpAsSiGnMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/groupAssignments/deviceManagementScriptGroupAssignmentId",
			Expected: &DeviceManagementDeviceShellScriptIdGroupAssignmentId{
				DeviceShellScriptId:                     "deviceShellScriptId",
				DeviceManagementScriptGroupAssignmentId: "deviceManagementScriptGroupAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceShellScripts/deviceShellScriptId/groupAssignments/deviceManagementScriptGroupAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEsHeLlScRiPtS/dEvIcEsHeLlScRiPtId/gRoUpAsSiGnMeNtS/dEvIcEmAnAgEmEnTsCrIpTgRoUpAsSiGnMeNtId",
			Expected: &DeviceManagementDeviceShellScriptIdGroupAssignmentId{
				DeviceShellScriptId:                     "dEvIcEsHeLlScRiPtId",
				DeviceManagementScriptGroupAssignmentId: "dEvIcEmAnAgEmEnTsCrIpTgRoUpAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEsHeLlScRiPtS/dEvIcEsHeLlScRiPtId/gRoUpAsSiGnMeNtS/dEvIcEmAnAgEmEnTsCrIpTgRoUpAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceShellScriptIdGroupAssignmentIDInsensitively(v.Input)
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

		if actual.DeviceManagementScriptGroupAssignmentId != v.Expected.DeviceManagementScriptGroupAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptGroupAssignmentId", v.Expected.DeviceManagementScriptGroupAssignmentId, actual.DeviceManagementScriptGroupAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceShellScriptIdGroupAssignmentId(t *testing.T) {
	segments := DeviceManagementDeviceShellScriptIdGroupAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceShellScriptIdGroupAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
