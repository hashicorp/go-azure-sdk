package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceComplianceScriptIdDeviceRunStateId{}

func TestNewDeviceManagementDeviceComplianceScriptIdDeviceRunStateID(t *testing.T) {
	id := NewDeviceManagementDeviceComplianceScriptIdDeviceRunStateID("deviceComplianceScriptId", "deviceComplianceScriptDeviceStateId")

	if id.DeviceComplianceScriptId != "deviceComplianceScriptId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceComplianceScriptId'", id.DeviceComplianceScriptId, "deviceComplianceScriptId")
	}

	if id.DeviceComplianceScriptDeviceStateId != "deviceComplianceScriptDeviceStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceComplianceScriptDeviceStateId'", id.DeviceComplianceScriptDeviceStateId, "deviceComplianceScriptDeviceStateId")
	}
}

func TestFormatDeviceManagementDeviceComplianceScriptIdDeviceRunStateID(t *testing.T) {
	actual := NewDeviceManagementDeviceComplianceScriptIdDeviceRunStateID("deviceComplianceScriptId", "deviceComplianceScriptDeviceStateId").ID()
	expected := "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptId/deviceRunStates/deviceComplianceScriptDeviceStateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceComplianceScriptIdDeviceRunStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceComplianceScriptIdDeviceRunStateId
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
			Input: "/deviceManagement/deviceComplianceScripts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptId/deviceRunStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptId/deviceRunStates/deviceComplianceScriptDeviceStateId",
			Expected: &DeviceManagementDeviceComplianceScriptIdDeviceRunStateId{
				DeviceComplianceScriptId:            "deviceComplianceScriptId",
				DeviceComplianceScriptDeviceStateId: "deviceComplianceScriptDeviceStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptId/deviceRunStates/deviceComplianceScriptDeviceStateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceComplianceScriptIdDeviceRunStateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceComplianceScriptId != v.Expected.DeviceComplianceScriptId {
			t.Fatalf("Expected %q but got %q for DeviceComplianceScriptId", v.Expected.DeviceComplianceScriptId, actual.DeviceComplianceScriptId)
		}

		if actual.DeviceComplianceScriptDeviceStateId != v.Expected.DeviceComplianceScriptDeviceStateId {
			t.Fatalf("Expected %q but got %q for DeviceComplianceScriptDeviceStateId", v.Expected.DeviceComplianceScriptDeviceStateId, actual.DeviceComplianceScriptDeviceStateId)
		}

	}
}

func TestParseDeviceManagementDeviceComplianceScriptIdDeviceRunStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceComplianceScriptIdDeviceRunStateId
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
			Input: "/deviceManagement/deviceComplianceScripts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEsCrIpTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEsCrIpTs/dEvIcEcOmPlIaNcEsCrIpTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptId/deviceRunStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEsCrIpTs/dEvIcEcOmPlIaNcEsCrIpTiD/dEvIcErUnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptId/deviceRunStates/deviceComplianceScriptDeviceStateId",
			Expected: &DeviceManagementDeviceComplianceScriptIdDeviceRunStateId{
				DeviceComplianceScriptId:            "deviceComplianceScriptId",
				DeviceComplianceScriptDeviceStateId: "deviceComplianceScriptDeviceStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptId/deviceRunStates/deviceComplianceScriptDeviceStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEsCrIpTs/dEvIcEcOmPlIaNcEsCrIpTiD/dEvIcErUnStAtEs/dEvIcEcOmPlIaNcEsCrIpTdEvIcEsTaTeId",
			Expected: &DeviceManagementDeviceComplianceScriptIdDeviceRunStateId{
				DeviceComplianceScriptId:            "dEvIcEcOmPlIaNcEsCrIpTiD",
				DeviceComplianceScriptDeviceStateId: "dEvIcEcOmPlIaNcEsCrIpTdEvIcEsTaTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEsCrIpTs/dEvIcEcOmPlIaNcEsCrIpTiD/dEvIcErUnStAtEs/dEvIcEcOmPlIaNcEsCrIpTdEvIcEsTaTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceComplianceScriptIdDeviceRunStateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceComplianceScriptId != v.Expected.DeviceComplianceScriptId {
			t.Fatalf("Expected %q but got %q for DeviceComplianceScriptId", v.Expected.DeviceComplianceScriptId, actual.DeviceComplianceScriptId)
		}

		if actual.DeviceComplianceScriptDeviceStateId != v.Expected.DeviceComplianceScriptDeviceStateId {
			t.Fatalf("Expected %q but got %q for DeviceComplianceScriptDeviceStateId", v.Expected.DeviceComplianceScriptDeviceStateId, actual.DeviceComplianceScriptDeviceStateId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceComplianceScriptIdDeviceRunStateId(t *testing.T) {
	segments := DeviceManagementDeviceComplianceScriptIdDeviceRunStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceComplianceScriptIdDeviceRunStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
