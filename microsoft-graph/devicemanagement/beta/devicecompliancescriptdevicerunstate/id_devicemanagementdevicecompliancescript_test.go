package devicecompliancescriptdevicerunstate

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceComplianceScriptId{}

func TestNewDeviceManagementDeviceComplianceScriptID(t *testing.T) {
	id := NewDeviceManagementDeviceComplianceScriptID("deviceComplianceScriptIdValue")

	if id.DeviceComplianceScriptId != "deviceComplianceScriptIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceComplianceScriptId'", id.DeviceComplianceScriptId, "deviceComplianceScriptIdValue")
	}
}

func TestFormatDeviceManagementDeviceComplianceScriptID(t *testing.T) {
	actual := NewDeviceManagementDeviceComplianceScriptID("deviceComplianceScriptIdValue").ID()
	expected := "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceComplianceScriptID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceComplianceScriptId
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
			// Valid URI
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptIdValue",
			Expected: &DeviceManagementDeviceComplianceScriptId{
				DeviceComplianceScriptId: "deviceComplianceScriptIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceComplianceScriptID(v.Input)
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

	}
}

func TestParseDeviceManagementDeviceComplianceScriptIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceComplianceScriptId
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
			// Valid URI
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptIdValue",
			Expected: &DeviceManagementDeviceComplianceScriptId{
				DeviceComplianceScriptId: "deviceComplianceScriptIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceComplianceScripts/deviceComplianceScriptIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEsCrIpTs/dEvIcEcOmPlIaNcEsCrIpTiDvAlUe",
			Expected: &DeviceManagementDeviceComplianceScriptId{
				DeviceComplianceScriptId: "dEvIcEcOmPlIaNcEsCrIpTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEsCrIpTs/dEvIcEcOmPlIaNcEsCrIpTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceComplianceScriptIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementDeviceComplianceScriptId(t *testing.T) {
	segments := DeviceManagementDeviceComplianceScriptId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceComplianceScriptId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
