package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementHardwareConfigurationIdDeviceRunStateId{}

func TestNewDeviceManagementHardwareConfigurationIdDeviceRunStateID(t *testing.T) {
	id := NewDeviceManagementHardwareConfigurationIdDeviceRunStateID("hardwareConfigurationIdValue", "hardwareConfigurationDeviceStateIdValue")

	if id.HardwareConfigurationId != "hardwareConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwareConfigurationId'", id.HardwareConfigurationId, "hardwareConfigurationIdValue")
	}

	if id.HardwareConfigurationDeviceStateId != "hardwareConfigurationDeviceStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwareConfigurationDeviceStateId'", id.HardwareConfigurationDeviceStateId, "hardwareConfigurationDeviceStateIdValue")
	}
}

func TestFormatDeviceManagementHardwareConfigurationIdDeviceRunStateID(t *testing.T) {
	actual := NewDeviceManagementHardwareConfigurationIdDeviceRunStateID("hardwareConfigurationIdValue", "hardwareConfigurationDeviceStateIdValue").ID()
	expected := "/deviceManagement/hardwareConfigurations/hardwareConfigurationIdValue/deviceRunStates/hardwareConfigurationDeviceStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementHardwareConfigurationIdDeviceRunStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementHardwareConfigurationIdDeviceRunStateId
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
			// Incomplete URI
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationIdValue/deviceRunStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationIdValue/deviceRunStates/hardwareConfigurationDeviceStateIdValue",
			Expected: &DeviceManagementHardwareConfigurationIdDeviceRunStateId{
				HardwareConfigurationId:            "hardwareConfigurationIdValue",
				HardwareConfigurationDeviceStateId: "hardwareConfigurationDeviceStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationIdValue/deviceRunStates/hardwareConfigurationDeviceStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementHardwareConfigurationIdDeviceRunStateID(v.Input)
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

		if actual.HardwareConfigurationDeviceStateId != v.Expected.HardwareConfigurationDeviceStateId {
			t.Fatalf("Expected %q but got %q for HardwareConfigurationDeviceStateId", v.Expected.HardwareConfigurationDeviceStateId, actual.HardwareConfigurationDeviceStateId)
		}

	}
}

func TestParseDeviceManagementHardwareConfigurationIdDeviceRunStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementHardwareConfigurationIdDeviceRunStateId
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
			// Incomplete URI
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEcOnFiGuRaTiOnS/hArDwArEcOnFiGuRaTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationIdValue/deviceRunStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEcOnFiGuRaTiOnS/hArDwArEcOnFiGuRaTiOnIdVaLuE/dEvIcErUnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationIdValue/deviceRunStates/hardwareConfigurationDeviceStateIdValue",
			Expected: &DeviceManagementHardwareConfigurationIdDeviceRunStateId{
				HardwareConfigurationId:            "hardwareConfigurationIdValue",
				HardwareConfigurationDeviceStateId: "hardwareConfigurationDeviceStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationIdValue/deviceRunStates/hardwareConfigurationDeviceStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEcOnFiGuRaTiOnS/hArDwArEcOnFiGuRaTiOnIdVaLuE/dEvIcErUnStAtEs/hArDwArEcOnFiGuRaTiOnDeViCeStAtEiDvAlUe",
			Expected: &DeviceManagementHardwareConfigurationIdDeviceRunStateId{
				HardwareConfigurationId:            "hArDwArEcOnFiGuRaTiOnIdVaLuE",
				HardwareConfigurationDeviceStateId: "hArDwArEcOnFiGuRaTiOnDeViCeStAtEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEcOnFiGuRaTiOnS/hArDwArEcOnFiGuRaTiOnIdVaLuE/dEvIcErUnStAtEs/hArDwArEcOnFiGuRaTiOnDeViCeStAtEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementHardwareConfigurationIdDeviceRunStateIDInsensitively(v.Input)
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

		if actual.HardwareConfigurationDeviceStateId != v.Expected.HardwareConfigurationDeviceStateId {
			t.Fatalf("Expected %q but got %q for HardwareConfigurationDeviceStateId", v.Expected.HardwareConfigurationDeviceStateId, actual.HardwareConfigurationDeviceStateId)
		}

	}
}

func TestSegmentsForDeviceManagementHardwareConfigurationIdDeviceRunStateId(t *testing.T) {
	segments := DeviceManagementHardwareConfigurationIdDeviceRunStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementHardwareConfigurationIdDeviceRunStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
