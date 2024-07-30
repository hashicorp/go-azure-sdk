package devicemanagementscriptassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceManagementScriptId{}

func TestNewDeviceManagementDeviceManagementScriptID(t *testing.T) {
	id := NewDeviceManagementDeviceManagementScriptID("deviceManagementScriptIdValue")

	if id.DeviceManagementScriptId != "deviceManagementScriptIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementScriptId'", id.DeviceManagementScriptId, "deviceManagementScriptIdValue")
	}
}

func TestFormatDeviceManagementDeviceManagementScriptID(t *testing.T) {
	actual := NewDeviceManagementDeviceManagementScriptID("deviceManagementScriptIdValue").ID()
	expected := "/deviceManagement/deviceManagementScripts/deviceManagementScriptIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceManagementScriptID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceManagementScriptId
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
			Input: "/deviceManagement/deviceManagementScripts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptIdValue",
			Expected: &DeviceManagementDeviceManagementScriptId{
				DeviceManagementScriptId: "deviceManagementScriptIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceManagementScriptID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementScriptId != v.Expected.DeviceManagementScriptId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptId", v.Expected.DeviceManagementScriptId, actual.DeviceManagementScriptId)
		}

	}
}

func TestParseDeviceManagementDeviceManagementScriptIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceManagementScriptId
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
			Input: "/deviceManagement/deviceManagementScripts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEmAnAgEmEnTsCrIpTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptIdValue",
			Expected: &DeviceManagementDeviceManagementScriptId{
				DeviceManagementScriptId: "deviceManagementScriptIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEmAnAgEmEnTsCrIpTs/dEvIcEmAnAgEmEnTsCrIpTiDvAlUe",
			Expected: &DeviceManagementDeviceManagementScriptId{
				DeviceManagementScriptId: "dEvIcEmAnAgEmEnTsCrIpTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEmAnAgEmEnTsCrIpTs/dEvIcEmAnAgEmEnTsCrIpTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceManagementScriptIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementScriptId != v.Expected.DeviceManagementScriptId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptId", v.Expected.DeviceManagementScriptId, actual.DeviceManagementScriptId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceManagementScriptId(t *testing.T) {
	segments := DeviceManagementDeviceManagementScriptId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceManagementScriptId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
