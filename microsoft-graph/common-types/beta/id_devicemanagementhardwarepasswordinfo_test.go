package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementHardwarePasswordInfoId{}

func TestNewDeviceManagementHardwarePasswordInfoID(t *testing.T) {
	id := NewDeviceManagementHardwarePasswordInfoID("hardwarePasswordInfoId")

	if id.HardwarePasswordInfoId != "hardwarePasswordInfoId" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwarePasswordInfoId'", id.HardwarePasswordInfoId, "hardwarePasswordInfoId")
	}
}

func TestFormatDeviceManagementHardwarePasswordInfoID(t *testing.T) {
	actual := NewDeviceManagementHardwarePasswordInfoID("hardwarePasswordInfoId").ID()
	expected := "/deviceManagement/hardwarePasswordInfo/hardwarePasswordInfoId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementHardwarePasswordInfoID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementHardwarePasswordInfoId
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
			Input: "/deviceManagement/hardwarePasswordInfo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/hardwarePasswordInfo/hardwarePasswordInfoId",
			Expected: &DeviceManagementHardwarePasswordInfoId{
				HardwarePasswordInfoId: "hardwarePasswordInfoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/hardwarePasswordInfo/hardwarePasswordInfoId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementHardwarePasswordInfoID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.HardwarePasswordInfoId != v.Expected.HardwarePasswordInfoId {
			t.Fatalf("Expected %q but got %q for HardwarePasswordInfoId", v.Expected.HardwarePasswordInfoId, actual.HardwarePasswordInfoId)
		}

	}
}

func TestParseDeviceManagementHardwarePasswordInfoIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementHardwarePasswordInfoId
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
			Input: "/deviceManagement/hardwarePasswordInfo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEpAsSwOrDiNfO",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/hardwarePasswordInfo/hardwarePasswordInfoId",
			Expected: &DeviceManagementHardwarePasswordInfoId{
				HardwarePasswordInfoId: "hardwarePasswordInfoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/hardwarePasswordInfo/hardwarePasswordInfoId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEpAsSwOrDiNfO/hArDwArEpAsSwOrDiNfOiD",
			Expected: &DeviceManagementHardwarePasswordInfoId{
				HardwarePasswordInfoId: "hArDwArEpAsSwOrDiNfOiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEpAsSwOrDiNfO/hArDwArEpAsSwOrDiNfOiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementHardwarePasswordInfoIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.HardwarePasswordInfoId != v.Expected.HardwarePasswordInfoId {
			t.Fatalf("Expected %q but got %q for HardwarePasswordInfoId", v.Expected.HardwarePasswordInfoId, actual.HardwarePasswordInfoId)
		}

	}
}

func TestSegmentsForDeviceManagementHardwarePasswordInfoId(t *testing.T) {
	segments := DeviceManagementHardwarePasswordInfoId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementHardwarePasswordInfoId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
