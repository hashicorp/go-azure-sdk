package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementHardwarePasswordInfoId{}

func TestNewDeviceManagementHardwarePasswordInfoID(t *testing.T) {
	id := NewDeviceManagementHardwarePasswordInfoID("hardwarePasswordInfoIdValue")

	if id.HardwarePasswordInfoId != "hardwarePasswordInfoIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwarePasswordInfoId'", id.HardwarePasswordInfoId, "hardwarePasswordInfoIdValue")
	}
}

func TestFormatDeviceManagementHardwarePasswordInfoID(t *testing.T) {
	actual := NewDeviceManagementHardwarePasswordInfoID("hardwarePasswordInfoIdValue").ID()
	expected := "/deviceManagement/hardwarePasswordInfo/hardwarePasswordInfoIdValue"
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
			Input: "/deviceManagement/hardwarePasswordInfo/hardwarePasswordInfoIdValue",
			Expected: &DeviceManagementHardwarePasswordInfoId{
				HardwarePasswordInfoId: "hardwarePasswordInfoIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/hardwarePasswordInfo/hardwarePasswordInfoIdValue/extra",
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
			Input: "/deviceManagement/hardwarePasswordInfo/hardwarePasswordInfoIdValue",
			Expected: &DeviceManagementHardwarePasswordInfoId{
				HardwarePasswordInfoId: "hardwarePasswordInfoIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/hardwarePasswordInfo/hardwarePasswordInfoIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEpAsSwOrDiNfO/hArDwArEpAsSwOrDiNfOiDvAlUe",
			Expected: &DeviceManagementHardwarePasswordInfoId{
				HardwarePasswordInfoId: "hArDwArEpAsSwOrDiNfOiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEpAsSwOrDiNfO/hArDwArEpAsSwOrDiNfOiDvAlUe/extra",
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
