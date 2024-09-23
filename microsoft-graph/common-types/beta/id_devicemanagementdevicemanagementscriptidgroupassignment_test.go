package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceManagementScriptIdGroupAssignmentId{}

func TestNewDeviceManagementDeviceManagementScriptIdGroupAssignmentID(t *testing.T) {
	id := NewDeviceManagementDeviceManagementScriptIdGroupAssignmentID("deviceManagementScriptId", "deviceManagementScriptGroupAssignmentId")

	if id.DeviceManagementScriptId != "deviceManagementScriptId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementScriptId'", id.DeviceManagementScriptId, "deviceManagementScriptId")
	}

	if id.DeviceManagementScriptGroupAssignmentId != "deviceManagementScriptGroupAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementScriptGroupAssignmentId'", id.DeviceManagementScriptGroupAssignmentId, "deviceManagementScriptGroupAssignmentId")
	}
}

func TestFormatDeviceManagementDeviceManagementScriptIdGroupAssignmentID(t *testing.T) {
	actual := NewDeviceManagementDeviceManagementScriptIdGroupAssignmentID("deviceManagementScriptId", "deviceManagementScriptGroupAssignmentId").ID()
	expected := "/deviceManagement/deviceManagementScripts/deviceManagementScriptId/groupAssignments/deviceManagementScriptGroupAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceManagementScriptIdGroupAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceManagementScriptIdGroupAssignmentId
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
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptId/groupAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptId/groupAssignments/deviceManagementScriptGroupAssignmentId",
			Expected: &DeviceManagementDeviceManagementScriptIdGroupAssignmentId{
				DeviceManagementScriptId:                "deviceManagementScriptId",
				DeviceManagementScriptGroupAssignmentId: "deviceManagementScriptGroupAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptId/groupAssignments/deviceManagementScriptGroupAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceManagementScriptIdGroupAssignmentID(v.Input)
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

		if actual.DeviceManagementScriptGroupAssignmentId != v.Expected.DeviceManagementScriptGroupAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptGroupAssignmentId", v.Expected.DeviceManagementScriptGroupAssignmentId, actual.DeviceManagementScriptGroupAssignmentId)
		}

	}
}

func TestParseDeviceManagementDeviceManagementScriptIdGroupAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceManagementScriptIdGroupAssignmentId
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
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptId/groupAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEmAnAgEmEnTsCrIpTs/dEvIcEmAnAgEmEnTsCrIpTiD/gRoUpAsSiGnMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptId/groupAssignments/deviceManagementScriptGroupAssignmentId",
			Expected: &DeviceManagementDeviceManagementScriptIdGroupAssignmentId{
				DeviceManagementScriptId:                "deviceManagementScriptId",
				DeviceManagementScriptGroupAssignmentId: "deviceManagementScriptGroupAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceManagementScripts/deviceManagementScriptId/groupAssignments/deviceManagementScriptGroupAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEmAnAgEmEnTsCrIpTs/dEvIcEmAnAgEmEnTsCrIpTiD/gRoUpAsSiGnMeNtS/dEvIcEmAnAgEmEnTsCrIpTgRoUpAsSiGnMeNtId",
			Expected: &DeviceManagementDeviceManagementScriptIdGroupAssignmentId{
				DeviceManagementScriptId:                "dEvIcEmAnAgEmEnTsCrIpTiD",
				DeviceManagementScriptGroupAssignmentId: "dEvIcEmAnAgEmEnTsCrIpTgRoUpAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEmAnAgEmEnTsCrIpTs/dEvIcEmAnAgEmEnTsCrIpTiD/gRoUpAsSiGnMeNtS/dEvIcEmAnAgEmEnTsCrIpTgRoUpAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceManagementScriptIdGroupAssignmentIDInsensitively(v.Input)
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

		if actual.DeviceManagementScriptGroupAssignmentId != v.Expected.DeviceManagementScriptGroupAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementScriptGroupAssignmentId", v.Expected.DeviceManagementScriptGroupAssignmentId, actual.DeviceManagementScriptGroupAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceManagementScriptIdGroupAssignmentId(t *testing.T) {
	segments := DeviceManagementDeviceManagementScriptIdGroupAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceManagementScriptIdGroupAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
