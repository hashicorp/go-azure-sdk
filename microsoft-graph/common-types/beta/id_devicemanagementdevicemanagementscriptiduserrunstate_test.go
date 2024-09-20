package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceManagementScriptIdUserRunStateId{}

func TestNewDeviceManagementDeviceManagementScriptIdUserRunStateID(t *testing.T) {
	id := NewDeviceManagementDeviceManagementScriptIdUserRunStateID("deviceManagementScriptId", "deviceManagementScriptUserStateId")

	if id.DeviceManagementScriptId != "deviceManagementScriptId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementScriptId'", id.DeviceManagementScriptId, "deviceManagementScriptId")
	}

	if id.DeviceManagementScriptUserStateId != "deviceManagementScriptUserStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementScriptUserStateId'", id.DeviceManagementScriptUserStateId, "deviceManagementScriptUserStateId")
	}
}

func TestFormatDeviceManagementDeviceManagementScriptIdUserRunStateID(t *testing.T) {
	actual := NewDeviceManagementDeviceManagementScriptIdUserRunStateID("deviceManagementScriptId", "deviceManagementScriptUserStateId").ID()
	expected := "/deviceManagement/deviceManagementScripts/deviceManagementScriptId/userRunStates/deviceManagementScriptUserStateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceManagementScriptIdUserRunStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceManagementScriptIdUserRunStateId
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
			// Incomplete URI
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptId/userRunStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptId/userRunStates/deviceManagementScriptUserStateId",
			Expected: &DeviceManagementDeviceManagementScriptIdUserRunStateId{
				DeviceManagementScriptId:          "deviceManagementScriptId",
				DeviceManagementScriptUserStateId: "deviceManagementScriptUserStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptId/userRunStates/deviceManagementScriptUserStateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceManagementScriptIdUserRunStateID(v.Input)
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

		if actual.DeviceManagementScriptUserStateId != v.Expected.DeviceManagementScriptUserStateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptUserStateId", v.Expected.DeviceManagementScriptUserStateId, actual.DeviceManagementScriptUserStateId)
		}

	}
}

func TestParseDeviceManagementDeviceManagementScriptIdUserRunStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceManagementScriptIdUserRunStateId
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
			// Incomplete URI
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEmAnAgEmEnTsCrIpTs/dEvIcEmAnAgEmEnTsCrIpTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptId/userRunStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEmAnAgEmEnTsCrIpTs/dEvIcEmAnAgEmEnTsCrIpTiD/uSeRrUnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptId/userRunStates/deviceManagementScriptUserStateId",
			Expected: &DeviceManagementDeviceManagementScriptIdUserRunStateId{
				DeviceManagementScriptId:          "deviceManagementScriptId",
				DeviceManagementScriptUserStateId: "deviceManagementScriptUserStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptId/userRunStates/deviceManagementScriptUserStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEmAnAgEmEnTsCrIpTs/dEvIcEmAnAgEmEnTsCrIpTiD/uSeRrUnStAtEs/dEvIcEmAnAgEmEnTsCrIpTuSeRsTaTeId",
			Expected: &DeviceManagementDeviceManagementScriptIdUserRunStateId{
				DeviceManagementScriptId:          "dEvIcEmAnAgEmEnTsCrIpTiD",
				DeviceManagementScriptUserStateId: "dEvIcEmAnAgEmEnTsCrIpTuSeRsTaTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEmAnAgEmEnTsCrIpTs/dEvIcEmAnAgEmEnTsCrIpTiD/uSeRrUnStAtEs/dEvIcEmAnAgEmEnTsCrIpTuSeRsTaTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceManagementScriptIdUserRunStateIDInsensitively(v.Input)
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

		if actual.DeviceManagementScriptUserStateId != v.Expected.DeviceManagementScriptUserStateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptUserStateId", v.Expected.DeviceManagementScriptUserStateId, actual.DeviceManagementScriptUserStateId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceManagementScriptIdUserRunStateId(t *testing.T) {
	segments := DeviceManagementDeviceManagementScriptIdUserRunStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceManagementScriptIdUserRunStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
