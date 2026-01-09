package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId{}

func TestNewDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentID("deviceCustomAttributeShellScriptId", "deviceManagementScriptAssignmentId")

	if id.DeviceCustomAttributeShellScriptId != "deviceCustomAttributeShellScriptId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCustomAttributeShellScriptId'", id.DeviceCustomAttributeShellScriptId, "deviceCustomAttributeShellScriptId")
	}

	if id.DeviceManagementScriptAssignmentId != "deviceManagementScriptAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementScriptAssignmentId'", id.DeviceManagementScriptAssignmentId, "deviceManagementScriptAssignmentId")
	}
}

func TestFormatDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentID("deviceCustomAttributeShellScriptId", "deviceManagementScriptAssignmentId").ID()
	expected := "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/assignments/deviceManagementScriptAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId
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
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/assignments/deviceManagementScriptAssignmentId",
			Expected: &DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId{
				DeviceCustomAttributeShellScriptId: "deviceCustomAttributeShellScriptId",
				DeviceManagementScriptAssignmentId: "deviceManagementScriptAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/assignments/deviceManagementScriptAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentID(v.Input)
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

		if actual.DeviceManagementScriptAssignmentId != v.Expected.DeviceManagementScriptAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptAssignmentId", v.Expected.DeviceManagementScriptAssignmentId, actual.DeviceManagementScriptAssignmentId)
		}

	}
}

func TestParseDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId
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
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTs/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/assignments/deviceManagementScriptAssignmentId",
			Expected: &DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId{
				DeviceCustomAttributeShellScriptId: "deviceCustomAttributeShellScriptId",
				DeviceManagementScriptAssignmentId: "deviceManagementScriptAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/assignments/deviceManagementScriptAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTs/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD/aSsIgNmEnTs/dEvIcEmAnAgEmEnTsCrIpTaSsIgNmEnTiD",
			Expected: &DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId{
				DeviceCustomAttributeShellScriptId: "dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD",
				DeviceManagementScriptAssignmentId: "dEvIcEmAnAgEmEnTsCrIpTaSsIgNmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTs/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD/aSsIgNmEnTs/dEvIcEmAnAgEmEnTsCrIpTaSsIgNmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentIDInsensitively(v.Input)
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

		if actual.DeviceManagementScriptAssignmentId != v.Expected.DeviceManagementScriptAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptAssignmentId", v.Expected.DeviceManagementScriptAssignmentId, actual.DeviceManagementScriptAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId(t *testing.T) {
	segments := DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceCustomAttributeShellScriptIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
