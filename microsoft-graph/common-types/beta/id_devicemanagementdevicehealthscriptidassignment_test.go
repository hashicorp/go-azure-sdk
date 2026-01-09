package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceHealthScriptIdAssignmentId{}

func TestNewDeviceManagementDeviceHealthScriptIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementDeviceHealthScriptIdAssignmentID("deviceHealthScriptId", "deviceHealthScriptAssignmentId")

	if id.DeviceHealthScriptId != "deviceHealthScriptId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceHealthScriptId'", id.DeviceHealthScriptId, "deviceHealthScriptId")
	}

	if id.DeviceHealthScriptAssignmentId != "deviceHealthScriptAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceHealthScriptAssignmentId'", id.DeviceHealthScriptAssignmentId, "deviceHealthScriptAssignmentId")
	}
}

func TestFormatDeviceManagementDeviceHealthScriptIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementDeviceHealthScriptIdAssignmentID("deviceHealthScriptId", "deviceHealthScriptAssignmentId").ID()
	expected := "/deviceManagement/deviceHealthScripts/deviceHealthScriptId/assignments/deviceHealthScriptAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceHealthScriptIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceHealthScriptIdAssignmentId
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
			Input: "/deviceManagement/deviceHealthScripts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptId/assignments/deviceHealthScriptAssignmentId",
			Expected: &DeviceManagementDeviceHealthScriptIdAssignmentId{
				DeviceHealthScriptId:           "deviceHealthScriptId",
				DeviceHealthScriptAssignmentId: "deviceHealthScriptAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptId/assignments/deviceHealthScriptAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceHealthScriptIdAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceHealthScriptId != v.Expected.DeviceHealthScriptId {
			t.Fatalf("Expected %q but got %q for DeviceHealthScriptId", v.Expected.DeviceHealthScriptId, actual.DeviceHealthScriptId)
		}

		if actual.DeviceHealthScriptAssignmentId != v.Expected.DeviceHealthScriptAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceHealthScriptAssignmentId", v.Expected.DeviceHealthScriptAssignmentId, actual.DeviceHealthScriptAssignmentId)
		}

	}
}

func TestParseDeviceManagementDeviceHealthScriptIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceHealthScriptIdAssignmentId
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
			Input: "/deviceManagement/deviceHealthScripts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEhEaLtHsCrIpTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEhEaLtHsCrIpTs/dEvIcEhEaLtHsCrIpTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEhEaLtHsCrIpTs/dEvIcEhEaLtHsCrIpTiD/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptId/assignments/deviceHealthScriptAssignmentId",
			Expected: &DeviceManagementDeviceHealthScriptIdAssignmentId{
				DeviceHealthScriptId:           "deviceHealthScriptId",
				DeviceHealthScriptAssignmentId: "deviceHealthScriptAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptId/assignments/deviceHealthScriptAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEhEaLtHsCrIpTs/dEvIcEhEaLtHsCrIpTiD/aSsIgNmEnTs/dEvIcEhEaLtHsCrIpTaSsIgNmEnTiD",
			Expected: &DeviceManagementDeviceHealthScriptIdAssignmentId{
				DeviceHealthScriptId:           "dEvIcEhEaLtHsCrIpTiD",
				DeviceHealthScriptAssignmentId: "dEvIcEhEaLtHsCrIpTaSsIgNmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEhEaLtHsCrIpTs/dEvIcEhEaLtHsCrIpTiD/aSsIgNmEnTs/dEvIcEhEaLtHsCrIpTaSsIgNmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceHealthScriptIdAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceHealthScriptId != v.Expected.DeviceHealthScriptId {
			t.Fatalf("Expected %q but got %q for DeviceHealthScriptId", v.Expected.DeviceHealthScriptId, actual.DeviceHealthScriptId)
		}

		if actual.DeviceHealthScriptAssignmentId != v.Expected.DeviceHealthScriptAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceHealthScriptAssignmentId", v.Expected.DeviceHealthScriptAssignmentId, actual.DeviceHealthScriptAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceHealthScriptIdAssignmentId(t *testing.T) {
	segments := DeviceManagementDeviceHealthScriptIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceHealthScriptIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
