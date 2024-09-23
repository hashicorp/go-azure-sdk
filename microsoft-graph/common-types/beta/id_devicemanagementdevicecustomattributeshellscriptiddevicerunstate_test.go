package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId{}

func TestNewDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateID(t *testing.T) {
	id := NewDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateID("deviceCustomAttributeShellScriptId", "deviceManagementScriptDeviceStateId")

	if id.DeviceCustomAttributeShellScriptId != "deviceCustomAttributeShellScriptId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCustomAttributeShellScriptId'", id.DeviceCustomAttributeShellScriptId, "deviceCustomAttributeShellScriptId")
	}

	if id.DeviceManagementScriptDeviceStateId != "deviceManagementScriptDeviceStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementScriptDeviceStateId'", id.DeviceManagementScriptDeviceStateId, "deviceManagementScriptDeviceStateId")
	}
}

func TestFormatDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateID(t *testing.T) {
	actual := NewDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateID("deviceCustomAttributeShellScriptId", "deviceManagementScriptDeviceStateId").ID()
	expected := "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/deviceRunStates/deviceManagementScriptDeviceStateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId
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
			Input: "/deviceManagement/deviceCustomAttributeShellScripts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/deviceRunStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/deviceRunStates/deviceManagementScriptDeviceStateId",
			Expected: &DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId{
				DeviceCustomAttributeShellScriptId:  "deviceCustomAttributeShellScriptId",
				DeviceManagementScriptDeviceStateId: "deviceManagementScriptDeviceStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/deviceRunStates/deviceManagementScriptDeviceStateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceCustomAttributeShellScriptId != v.Expected.DeviceCustomAttributeShellScriptId {
			t.Fatalf("Expected %q but got %q for DeviceCustomAttributeShellScriptId", v.Expected.DeviceCustomAttributeShellScriptId, actual.DeviceCustomAttributeShellScriptId)
		}

		if actual.DeviceManagementScriptDeviceStateId != v.Expected.DeviceManagementScriptDeviceStateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptDeviceStateId", v.Expected.DeviceManagementScriptDeviceStateId, actual.DeviceManagementScriptDeviceStateId)
		}

	}
}

func TestParseDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId
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
			Input: "/deviceManagement/deviceCustomAttributeShellScripts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTs/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/deviceRunStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTs/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD/dEvIcErUnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/deviceRunStates/deviceManagementScriptDeviceStateId",
			Expected: &DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId{
				DeviceCustomAttributeShellScriptId:  "deviceCustomAttributeShellScriptId",
				DeviceManagementScriptDeviceStateId: "deviceManagementScriptDeviceStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/deviceRunStates/deviceManagementScriptDeviceStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTs/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD/dEvIcErUnStAtEs/dEvIcEmAnAgEmEnTsCrIpTdEvIcEsTaTeId",
			Expected: &DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId{
				DeviceCustomAttributeShellScriptId:  "dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD",
				DeviceManagementScriptDeviceStateId: "dEvIcEmAnAgEmEnTsCrIpTdEvIcEsTaTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTs/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD/dEvIcErUnStAtEs/dEvIcEmAnAgEmEnTsCrIpTdEvIcEsTaTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceCustomAttributeShellScriptId != v.Expected.DeviceCustomAttributeShellScriptId {
			t.Fatalf("Expected %q but got %q for DeviceCustomAttributeShellScriptId", v.Expected.DeviceCustomAttributeShellScriptId, actual.DeviceCustomAttributeShellScriptId)
		}

		if actual.DeviceManagementScriptDeviceStateId != v.Expected.DeviceManagementScriptDeviceStateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptDeviceStateId", v.Expected.DeviceManagementScriptDeviceStateId, actual.DeviceManagementScriptDeviceStateId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId(t *testing.T) {
	segments := DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
