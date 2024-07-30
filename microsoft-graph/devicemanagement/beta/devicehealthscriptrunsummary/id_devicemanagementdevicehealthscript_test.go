package devicehealthscriptrunsummary

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceHealthScriptId{}

func TestNewDeviceManagementDeviceHealthScriptID(t *testing.T) {
	id := NewDeviceManagementDeviceHealthScriptID("deviceHealthScriptIdValue")

	if id.DeviceHealthScriptId != "deviceHealthScriptIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceHealthScriptId'", id.DeviceHealthScriptId, "deviceHealthScriptIdValue")
	}
}

func TestFormatDeviceManagementDeviceHealthScriptID(t *testing.T) {
	actual := NewDeviceManagementDeviceHealthScriptID("deviceHealthScriptIdValue").ID()
	expected := "/deviceManagement/deviceHealthScripts/deviceHealthScriptIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceHealthScriptID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceHealthScriptId
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
			// Valid URI
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptIdValue",
			Expected: &DeviceManagementDeviceHealthScriptId{
				DeviceHealthScriptId: "deviceHealthScriptIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceHealthScriptID(v.Input)
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

	}
}

func TestParseDeviceManagementDeviceHealthScriptIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceHealthScriptId
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
			// Valid URI
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptIdValue",
			Expected: &DeviceManagementDeviceHealthScriptId{
				DeviceHealthScriptId: "deviceHealthScriptIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceHealthScripts/deviceHealthScriptIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEhEaLtHsCrIpTs/dEvIcEhEaLtHsCrIpTiDvAlUe",
			Expected: &DeviceManagementDeviceHealthScriptId{
				DeviceHealthScriptId: "dEvIcEhEaLtHsCrIpTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEhEaLtHsCrIpTs/dEvIcEhEaLtHsCrIpTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceHealthScriptIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementDeviceHealthScriptId(t *testing.T) {
	segments := DeviceManagementDeviceHealthScriptId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceHealthScriptId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
