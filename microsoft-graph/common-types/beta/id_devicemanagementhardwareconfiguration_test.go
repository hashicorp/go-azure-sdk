package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementHardwareConfigurationId{}

func TestNewDeviceManagementHardwareConfigurationID(t *testing.T) {
	id := NewDeviceManagementHardwareConfigurationID("hardwareConfigurationId")

	if id.HardwareConfigurationId != "hardwareConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwareConfigurationId'", id.HardwareConfigurationId, "hardwareConfigurationId")
	}
}

func TestFormatDeviceManagementHardwareConfigurationID(t *testing.T) {
	actual := NewDeviceManagementHardwareConfigurationID("hardwareConfigurationId").ID()
	expected := "/deviceManagement/hardwareConfigurations/hardwareConfigurationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementHardwareConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementHardwareConfigurationId
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
			Input: "/deviceManagement/hardwareConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationId",
			Expected: &DeviceManagementHardwareConfigurationId{
				HardwareConfigurationId: "hardwareConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementHardwareConfigurationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.HardwareConfigurationId != v.Expected.HardwareConfigurationId {
			t.Fatalf("Expected %q but got %q for HardwareConfigurationId", v.Expected.HardwareConfigurationId, actual.HardwareConfigurationId)
		}

	}
}

func TestParseDeviceManagementHardwareConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementHardwareConfigurationId
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
			Input: "/deviceManagement/hardwareConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationId",
			Expected: &DeviceManagementHardwareConfigurationId{
				HardwareConfigurationId: "hardwareConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEcOnFiGuRaTiOnS/hArDwArEcOnFiGuRaTiOnId",
			Expected: &DeviceManagementHardwareConfigurationId{
				HardwareConfigurationId: "hArDwArEcOnFiGuRaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEcOnFiGuRaTiOnS/hArDwArEcOnFiGuRaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementHardwareConfigurationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.HardwareConfigurationId != v.Expected.HardwareConfigurationId {
			t.Fatalf("Expected %q but got %q for HardwareConfigurationId", v.Expected.HardwareConfigurationId, actual.HardwareConfigurationId)
		}

	}
}

func TestSegmentsForDeviceManagementHardwareConfigurationId(t *testing.T) {
	segments := DeviceManagementHardwareConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementHardwareConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
