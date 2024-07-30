package devicecompliancescriptassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceComplianceScriptIdAssignmentId{}

func TestNewDeviceManagementDeviceComplianceScriptIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementDeviceComplianceScriptIdAssignmentID("deviceComplianceScriptIdValue", "deviceHealthScriptAssignmentIdValue")

	if id.DeviceComplianceScriptId != "deviceComplianceScriptIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceComplianceScriptId'", id.DeviceComplianceScriptId, "deviceComplianceScriptIdValue")
	}

	if id.DeviceHealthScriptAssignmentId != "deviceHealthScriptAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceHealthScriptAssignmentId'", id.DeviceHealthScriptAssignmentId, "deviceHealthScriptAssignmentIdValue")
	}
}

func TestFormatDeviceManagementDeviceComplianceScriptIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementDeviceComplianceScriptIdAssignmentID("deviceComplianceScriptIdValue", "deviceHealthScriptAssignmentIdValue").ID()
	expected := "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptIdValue/assignments/deviceHealthScriptAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceComplianceScriptIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceComplianceScriptIdAssignmentId
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
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptIdValue/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptIdValue/assignments/deviceHealthScriptAssignmentIdValue",
			Expected: &DeviceManagementDeviceComplianceScriptIdAssignmentId{
				DeviceComplianceScriptId:       "deviceComplianceScriptIdValue",
				DeviceHealthScriptAssignmentId: "deviceHealthScriptAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptIdValue/assignments/deviceHealthScriptAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceComplianceScriptIdAssignmentID(v.Input)
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

		if actual.DeviceHealthScriptAssignmentId != v.Expected.DeviceHealthScriptAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceHealthScriptAssignmentId", v.Expected.DeviceHealthScriptAssignmentId, actual.DeviceHealthScriptAssignmentId)
		}

	}
}

func TestParseDeviceManagementDeviceComplianceScriptIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceComplianceScriptIdAssignmentId
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
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEsCrIpTs/dEvIcEcOmPlIaNcEsCrIpTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptIdValue/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEsCrIpTs/dEvIcEcOmPlIaNcEsCrIpTiDvAlUe/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptIdValue/assignments/deviceHealthScriptAssignmentIdValue",
			Expected: &DeviceManagementDeviceComplianceScriptIdAssignmentId{
				DeviceComplianceScriptId:       "deviceComplianceScriptIdValue",
				DeviceHealthScriptAssignmentId: "deviceHealthScriptAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptIdValue/assignments/deviceHealthScriptAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEsCrIpTs/dEvIcEcOmPlIaNcEsCrIpTiDvAlUe/aSsIgNmEnTs/dEvIcEhEaLtHsCrIpTaSsIgNmEnTiDvAlUe",
			Expected: &DeviceManagementDeviceComplianceScriptIdAssignmentId{
				DeviceComplianceScriptId:       "dEvIcEcOmPlIaNcEsCrIpTiDvAlUe",
				DeviceHealthScriptAssignmentId: "dEvIcEhEaLtHsCrIpTaSsIgNmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEsCrIpTs/dEvIcEcOmPlIaNcEsCrIpTiDvAlUe/aSsIgNmEnTs/dEvIcEhEaLtHsCrIpTaSsIgNmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceComplianceScriptIdAssignmentIDInsensitively(v.Input)
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

		if actual.DeviceHealthScriptAssignmentId != v.Expected.DeviceHealthScriptAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceHealthScriptAssignmentId", v.Expected.DeviceHealthScriptAssignmentId, actual.DeviceHealthScriptAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceComplianceScriptIdAssignmentId(t *testing.T) {
	segments := DeviceManagementDeviceComplianceScriptIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceComplianceScriptIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
