package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId{}

func TestNewDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateID(t *testing.T) {
	id := NewDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateID("deviceCustomAttributeShellScriptId", "deviceManagementScriptUserStateId", "deviceManagementScriptDeviceStateId")

	if id.DeviceCustomAttributeShellScriptId != "deviceCustomAttributeShellScriptId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCustomAttributeShellScriptId'", id.DeviceCustomAttributeShellScriptId, "deviceCustomAttributeShellScriptId")
	}

	if id.DeviceManagementScriptUserStateId != "deviceManagementScriptUserStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementScriptUserStateId'", id.DeviceManagementScriptUserStateId, "deviceManagementScriptUserStateId")
	}

	if id.DeviceManagementScriptDeviceStateId != "deviceManagementScriptDeviceStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementScriptDeviceStateId'", id.DeviceManagementScriptDeviceStateId, "deviceManagementScriptDeviceStateId")
	}
}

func TestFormatDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateID(t *testing.T) {
	actual := NewDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateID("deviceCustomAttributeShellScriptId", "deviceManagementScriptUserStateId", "deviceManagementScriptDeviceStateId").ID()
	expected := "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/userRunStates/deviceManagementScriptUserStateId/deviceRunStates/deviceManagementScriptDeviceStateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId
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
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/userRunStates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/userRunStates/deviceManagementScriptUserStateId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/userRunStates/deviceManagementScriptUserStateId/deviceRunStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/userRunStates/deviceManagementScriptUserStateId/deviceRunStates/deviceManagementScriptDeviceStateId",
			Expected: &DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId{
				DeviceCustomAttributeShellScriptId:  "deviceCustomAttributeShellScriptId",
				DeviceManagementScriptUserStateId:   "deviceManagementScriptUserStateId",
				DeviceManagementScriptDeviceStateId: "deviceManagementScriptDeviceStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/userRunStates/deviceManagementScriptUserStateId/deviceRunStates/deviceManagementScriptDeviceStateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateID(v.Input)
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

		if actual.DeviceManagementScriptUserStateId != v.Expected.DeviceManagementScriptUserStateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptUserStateId", v.Expected.DeviceManagementScriptUserStateId, actual.DeviceManagementScriptUserStateId)
		}

		if actual.DeviceManagementScriptDeviceStateId != v.Expected.DeviceManagementScriptDeviceStateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptDeviceStateId", v.Expected.DeviceManagementScriptDeviceStateId, actual.DeviceManagementScriptDeviceStateId)
		}

	}
}

func TestParseDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId
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
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/userRunStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTs/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD/uSeRrUnStAtEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/userRunStates/deviceManagementScriptUserStateId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTs/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD/uSeRrUnStAtEs/dEvIcEmAnAgEmEnTsCrIpTuSeRsTaTeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/userRunStates/deviceManagementScriptUserStateId/deviceRunStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTs/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD/uSeRrUnStAtEs/dEvIcEmAnAgEmEnTsCrIpTuSeRsTaTeId/dEvIcErUnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/userRunStates/deviceManagementScriptUserStateId/deviceRunStates/deviceManagementScriptDeviceStateId",
			Expected: &DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId{
				DeviceCustomAttributeShellScriptId:  "deviceCustomAttributeShellScriptId",
				DeviceManagementScriptUserStateId:   "deviceManagementScriptUserStateId",
				DeviceManagementScriptDeviceStateId: "deviceManagementScriptDeviceStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/userRunStates/deviceManagementScriptUserStateId/deviceRunStates/deviceManagementScriptDeviceStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTs/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD/uSeRrUnStAtEs/dEvIcEmAnAgEmEnTsCrIpTuSeRsTaTeId/dEvIcErUnStAtEs/dEvIcEmAnAgEmEnTsCrIpTdEvIcEsTaTeId",
			Expected: &DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId{
				DeviceCustomAttributeShellScriptId:  "dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD",
				DeviceManagementScriptUserStateId:   "dEvIcEmAnAgEmEnTsCrIpTuSeRsTaTeId",
				DeviceManagementScriptDeviceStateId: "dEvIcEmAnAgEmEnTsCrIpTdEvIcEsTaTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTs/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD/uSeRrUnStAtEs/dEvIcEmAnAgEmEnTsCrIpTuSeRsTaTeId/dEvIcErUnStAtEs/dEvIcEmAnAgEmEnTsCrIpTdEvIcEsTaTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateIDInsensitively(v.Input)
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

		if actual.DeviceManagementScriptUserStateId != v.Expected.DeviceManagementScriptUserStateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptUserStateId", v.Expected.DeviceManagementScriptUserStateId, actual.DeviceManagementScriptUserStateId)
		}

		if actual.DeviceManagementScriptDeviceStateId != v.Expected.DeviceManagementScriptDeviceStateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptDeviceStateId", v.Expected.DeviceManagementScriptDeviceStateId, actual.DeviceManagementScriptDeviceStateId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId(t *testing.T) {
	segments := DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateIdDeviceRunStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
