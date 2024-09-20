package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceHealthScriptIdDeviceRunStateId{}

func TestNewDeviceManagementDeviceHealthScriptIdDeviceRunStateID(t *testing.T) {
	id := NewDeviceManagementDeviceHealthScriptIdDeviceRunStateID("deviceHealthScriptId", "deviceHealthScriptDeviceStateId")

	if id.DeviceHealthScriptId != "deviceHealthScriptId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceHealthScriptId'", id.DeviceHealthScriptId, "deviceHealthScriptId")
	}

	if id.DeviceHealthScriptDeviceStateId != "deviceHealthScriptDeviceStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceHealthScriptDeviceStateId'", id.DeviceHealthScriptDeviceStateId, "deviceHealthScriptDeviceStateId")
	}
}

func TestFormatDeviceManagementDeviceHealthScriptIdDeviceRunStateID(t *testing.T) {
	actual := NewDeviceManagementDeviceHealthScriptIdDeviceRunStateID("deviceHealthScriptId", "deviceHealthScriptDeviceStateId").ID()
	expected := "/deviceManagement/deviceHealthScripts/deviceHealthScriptId/deviceRunStates/deviceHealthScriptDeviceStateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceHealthScriptIdDeviceRunStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceHealthScriptIdDeviceRunStateId
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
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptId/deviceRunStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptId/deviceRunStates/deviceHealthScriptDeviceStateId",
			Expected: &DeviceManagementDeviceHealthScriptIdDeviceRunStateId{
				DeviceHealthScriptId:            "deviceHealthScriptId",
				DeviceHealthScriptDeviceStateId: "deviceHealthScriptDeviceStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptId/deviceRunStates/deviceHealthScriptDeviceStateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceHealthScriptIdDeviceRunStateID(v.Input)
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

		if actual.DeviceHealthScriptDeviceStateId != v.Expected.DeviceHealthScriptDeviceStateId {
			t.Fatalf("Expected %q but got %q for DeviceHealthScriptDeviceStateId", v.Expected.DeviceHealthScriptDeviceStateId, actual.DeviceHealthScriptDeviceStateId)
		}

	}
}

func TestParseDeviceManagementDeviceHealthScriptIdDeviceRunStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceHealthScriptIdDeviceRunStateId
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
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptId/deviceRunStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEhEaLtHsCrIpTs/dEvIcEhEaLtHsCrIpTiD/dEvIcErUnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptId/deviceRunStates/deviceHealthScriptDeviceStateId",
			Expected: &DeviceManagementDeviceHealthScriptIdDeviceRunStateId{
				DeviceHealthScriptId:            "deviceHealthScriptId",
				DeviceHealthScriptDeviceStateId: "deviceHealthScriptDeviceStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptId/deviceRunStates/deviceHealthScriptDeviceStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEhEaLtHsCrIpTs/dEvIcEhEaLtHsCrIpTiD/dEvIcErUnStAtEs/dEvIcEhEaLtHsCrIpTdEvIcEsTaTeId",
			Expected: &DeviceManagementDeviceHealthScriptIdDeviceRunStateId{
				DeviceHealthScriptId:            "dEvIcEhEaLtHsCrIpTiD",
				DeviceHealthScriptDeviceStateId: "dEvIcEhEaLtHsCrIpTdEvIcEsTaTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEhEaLtHsCrIpTs/dEvIcEhEaLtHsCrIpTiD/dEvIcErUnStAtEs/dEvIcEhEaLtHsCrIpTdEvIcEsTaTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceHealthScriptIdDeviceRunStateIDInsensitively(v.Input)
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

		if actual.DeviceHealthScriptDeviceStateId != v.Expected.DeviceHealthScriptDeviceStateId {
			t.Fatalf("Expected %q but got %q for DeviceHealthScriptDeviceStateId", v.Expected.DeviceHealthScriptDeviceStateId, actual.DeviceHealthScriptDeviceStateId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceHealthScriptIdDeviceRunStateId(t *testing.T) {
	segments := DeviceManagementDeviceHealthScriptIdDeviceRunStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceHealthScriptIdDeviceRunStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
