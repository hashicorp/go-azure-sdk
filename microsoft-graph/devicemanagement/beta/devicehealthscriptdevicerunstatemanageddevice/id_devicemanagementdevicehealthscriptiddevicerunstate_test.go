package devicehealthscriptdevicerunstatemanageddevice

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceHealthScriptIdDeviceRunStateId{}

func TestNewDeviceManagementDeviceHealthScriptIdDeviceRunStateID(t *testing.T) {
	id := NewDeviceManagementDeviceHealthScriptIdDeviceRunStateID("deviceHealthScriptIdValue", "deviceHealthScriptDeviceStateIdValue")

	if id.DeviceHealthScriptId != "deviceHealthScriptIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceHealthScriptId'", id.DeviceHealthScriptId, "deviceHealthScriptIdValue")
	}

	if id.DeviceHealthScriptDeviceStateId != "deviceHealthScriptDeviceStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceHealthScriptDeviceStateId'", id.DeviceHealthScriptDeviceStateId, "deviceHealthScriptDeviceStateIdValue")
	}
}

func TestFormatDeviceManagementDeviceHealthScriptIdDeviceRunStateID(t *testing.T) {
	actual := NewDeviceManagementDeviceHealthScriptIdDeviceRunStateID("deviceHealthScriptIdValue", "deviceHealthScriptDeviceStateIdValue").ID()
	expected := "/deviceManagement/deviceHealthScripts/deviceHealthScriptIdValue/deviceRunStates/deviceHealthScriptDeviceStateIdValue"
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
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptIdValue/deviceRunStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptIdValue/deviceRunStates/deviceHealthScriptDeviceStateIdValue",
			Expected: &DeviceManagementDeviceHealthScriptIdDeviceRunStateId{
				DeviceHealthScriptId:            "deviceHealthScriptIdValue",
				DeviceHealthScriptDeviceStateId: "deviceHealthScriptDeviceStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptIdValue/deviceRunStates/deviceHealthScriptDeviceStateIdValue/extra",
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
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEhEaLtHsCrIpTs/dEvIcEhEaLtHsCrIpTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptIdValue/deviceRunStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEhEaLtHsCrIpTs/dEvIcEhEaLtHsCrIpTiDvAlUe/dEvIcErUnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptIdValue/deviceRunStates/deviceHealthScriptDeviceStateIdValue",
			Expected: &DeviceManagementDeviceHealthScriptIdDeviceRunStateId{
				DeviceHealthScriptId:            "deviceHealthScriptIdValue",
				DeviceHealthScriptDeviceStateId: "deviceHealthScriptDeviceStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptIdValue/deviceRunStates/deviceHealthScriptDeviceStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEhEaLtHsCrIpTs/dEvIcEhEaLtHsCrIpTiDvAlUe/dEvIcErUnStAtEs/dEvIcEhEaLtHsCrIpTdEvIcEsTaTeIdVaLuE",
			Expected: &DeviceManagementDeviceHealthScriptIdDeviceRunStateId{
				DeviceHealthScriptId:            "dEvIcEhEaLtHsCrIpTiDvAlUe",
				DeviceHealthScriptDeviceStateId: "dEvIcEhEaLtHsCrIpTdEvIcEsTaTeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEhEaLtHsCrIpTs/dEvIcEhEaLtHsCrIpTiDvAlUe/dEvIcErUnStAtEs/dEvIcEhEaLtHsCrIpTdEvIcEsTaTeIdVaLuE/extra",
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
