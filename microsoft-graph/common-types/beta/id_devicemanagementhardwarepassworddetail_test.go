package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementHardwarePasswordDetailId{}

func TestNewDeviceManagementHardwarePasswordDetailID(t *testing.T) {
	id := NewDeviceManagementHardwarePasswordDetailID("hardwarePasswordDetailId")

	if id.HardwarePasswordDetailId != "hardwarePasswordDetailId" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwarePasswordDetailId'", id.HardwarePasswordDetailId, "hardwarePasswordDetailId")
	}
}

func TestFormatDeviceManagementHardwarePasswordDetailID(t *testing.T) {
	actual := NewDeviceManagementHardwarePasswordDetailID("hardwarePasswordDetailId").ID()
	expected := "/deviceManagement/hardwarePasswordDetails/hardwarePasswordDetailId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementHardwarePasswordDetailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementHardwarePasswordDetailId
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
			Input: "/deviceManagement/hardwarePasswordDetails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/hardwarePasswordDetails/hardwarePasswordDetailId",
			Expected: &DeviceManagementHardwarePasswordDetailId{
				HardwarePasswordDetailId: "hardwarePasswordDetailId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/hardwarePasswordDetails/hardwarePasswordDetailId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementHardwarePasswordDetailID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.HardwarePasswordDetailId != v.Expected.HardwarePasswordDetailId {
			t.Fatalf("Expected %q but got %q for HardwarePasswordDetailId", v.Expected.HardwarePasswordDetailId, actual.HardwarePasswordDetailId)
		}

	}
}

func TestParseDeviceManagementHardwarePasswordDetailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementHardwarePasswordDetailId
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
			Input: "/deviceManagement/hardwarePasswordDetails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEpAsSwOrDdEtAiLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/hardwarePasswordDetails/hardwarePasswordDetailId",
			Expected: &DeviceManagementHardwarePasswordDetailId{
				HardwarePasswordDetailId: "hardwarePasswordDetailId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/hardwarePasswordDetails/hardwarePasswordDetailId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEpAsSwOrDdEtAiLs/hArDwArEpAsSwOrDdEtAiLiD",
			Expected: &DeviceManagementHardwarePasswordDetailId{
				HardwarePasswordDetailId: "hArDwArEpAsSwOrDdEtAiLiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEpAsSwOrDdEtAiLs/hArDwArEpAsSwOrDdEtAiLiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementHardwarePasswordDetailIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.HardwarePasswordDetailId != v.Expected.HardwarePasswordDetailId {
			t.Fatalf("Expected %q but got %q for HardwarePasswordDetailId", v.Expected.HardwarePasswordDetailId, actual.HardwarePasswordDetailId)
		}

	}
}

func TestSegmentsForDeviceManagementHardwarePasswordDetailId(t *testing.T) {
	segments := DeviceManagementHardwarePasswordDetailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementHardwarePasswordDetailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
