package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId{}

func TestNewDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentID(t *testing.T) {
	id := NewDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentID("deviceCustomAttributeShellScriptId", "deviceManagementScriptGroupAssignmentId")

	if id.DeviceCustomAttributeShellScriptId != "deviceCustomAttributeShellScriptId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCustomAttributeShellScriptId'", id.DeviceCustomAttributeShellScriptId, "deviceCustomAttributeShellScriptId")
	}

	if id.DeviceManagementScriptGroupAssignmentId != "deviceManagementScriptGroupAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementScriptGroupAssignmentId'", id.DeviceManagementScriptGroupAssignmentId, "deviceManagementScriptGroupAssignmentId")
	}
}

func TestFormatDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentID(t *testing.T) {
	actual := NewDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentID("deviceCustomAttributeShellScriptId", "deviceManagementScriptGroupAssignmentId").ID()
	expected := "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/groupAssignments/deviceManagementScriptGroupAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId
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
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/groupAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/groupAssignments/deviceManagementScriptGroupAssignmentId",
			Expected: &DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId{
				DeviceCustomAttributeShellScriptId:      "deviceCustomAttributeShellScriptId",
				DeviceManagementScriptGroupAssignmentId: "deviceManagementScriptGroupAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/groupAssignments/deviceManagementScriptGroupAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentID(v.Input)
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

		if actual.DeviceManagementScriptGroupAssignmentId != v.Expected.DeviceManagementScriptGroupAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptGroupAssignmentId", v.Expected.DeviceManagementScriptGroupAssignmentId, actual.DeviceManagementScriptGroupAssignmentId)
		}

	}
}

func TestParseDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId
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
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/groupAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTs/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD/gRoUpAsSiGnMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/groupAssignments/deviceManagementScriptGroupAssignmentId",
			Expected: &DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId{
				DeviceCustomAttributeShellScriptId:      "deviceCustomAttributeShellScriptId",
				DeviceManagementScriptGroupAssignmentId: "deviceManagementScriptGroupAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCustomAttributeShellScripts/deviceCustomAttributeShellScriptId/groupAssignments/deviceManagementScriptGroupAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTs/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD/gRoUpAsSiGnMeNtS/dEvIcEmAnAgEmEnTsCrIpTgRoUpAsSiGnMeNtId",
			Expected: &DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId{
				DeviceCustomAttributeShellScriptId:      "dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD",
				DeviceManagementScriptGroupAssignmentId: "dEvIcEmAnAgEmEnTsCrIpTgRoUpAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTs/dEvIcEcUsToMaTtRiBuTeShElLsCrIpTiD/gRoUpAsSiGnMeNtS/dEvIcEmAnAgEmEnTsCrIpTgRoUpAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentIDInsensitively(v.Input)
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

		if actual.DeviceManagementScriptGroupAssignmentId != v.Expected.DeviceManagementScriptGroupAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptGroupAssignmentId", v.Expected.DeviceManagementScriptGroupAssignmentId, actual.DeviceManagementScriptGroupAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId(t *testing.T) {
	segments := DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceCustomAttributeShellScriptIdGroupAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
