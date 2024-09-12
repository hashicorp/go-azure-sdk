package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementHardwareConfigurationIdUserRunStateId{}

func TestNewDeviceManagementHardwareConfigurationIdUserRunStateID(t *testing.T) {
	id := NewDeviceManagementHardwareConfigurationIdUserRunStateID("hardwareConfigurationIdValue", "hardwareConfigurationUserStateIdValue")

	if id.HardwareConfigurationId != "hardwareConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwareConfigurationId'", id.HardwareConfigurationId, "hardwareConfigurationIdValue")
	}

	if id.HardwareConfigurationUserStateId != "hardwareConfigurationUserStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwareConfigurationUserStateId'", id.HardwareConfigurationUserStateId, "hardwareConfigurationUserStateIdValue")
	}
}

func TestFormatDeviceManagementHardwareConfigurationIdUserRunStateID(t *testing.T) {
	actual := NewDeviceManagementHardwareConfigurationIdUserRunStateID("hardwareConfigurationIdValue", "hardwareConfigurationUserStateIdValue").ID()
	expected := "/deviceManagement/hardwareConfigurations/hardwareConfigurationIdValue/userRunStates/hardwareConfigurationUserStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementHardwareConfigurationIdUserRunStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementHardwareConfigurationIdUserRunStateId
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
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationIdValue/userRunStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationIdValue/userRunStates/hardwareConfigurationUserStateIdValue",
			Expected: &DeviceManagementHardwareConfigurationIdUserRunStateId{
				HardwareConfigurationId:          "hardwareConfigurationIdValue",
				HardwareConfigurationUserStateId: "hardwareConfigurationUserStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationIdValue/userRunStates/hardwareConfigurationUserStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementHardwareConfigurationIdUserRunStateID(v.Input)
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

		if actual.HardwareConfigurationUserStateId != v.Expected.HardwareConfigurationUserStateId {
			t.Fatalf("Expected %q but got %q for HardwareConfigurationUserStateId", v.Expected.HardwareConfigurationUserStateId, actual.HardwareConfigurationUserStateId)
		}

	}
}

func TestParseDeviceManagementHardwareConfigurationIdUserRunStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementHardwareConfigurationIdUserRunStateId
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
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationIdValue/userRunStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEcOnFiGuRaTiOnS/hArDwArEcOnFiGuRaTiOnIdVaLuE/uSeRrUnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationIdValue/userRunStates/hardwareConfigurationUserStateIdValue",
			Expected: &DeviceManagementHardwareConfigurationIdUserRunStateId{
				HardwareConfigurationId:          "hardwareConfigurationIdValue",
				HardwareConfigurationUserStateId: "hardwareConfigurationUserStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationIdValue/userRunStates/hardwareConfigurationUserStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEcOnFiGuRaTiOnS/hArDwArEcOnFiGuRaTiOnIdVaLuE/uSeRrUnStAtEs/hArDwArEcOnFiGuRaTiOnUsErStAtEiDvAlUe",
			Expected: &DeviceManagementHardwareConfigurationIdUserRunStateId{
				HardwareConfigurationId:          "hArDwArEcOnFiGuRaTiOnIdVaLuE",
				HardwareConfigurationUserStateId: "hArDwArEcOnFiGuRaTiOnUsErStAtEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEcOnFiGuRaTiOnS/hArDwArEcOnFiGuRaTiOnIdVaLuE/uSeRrUnStAtEs/hArDwArEcOnFiGuRaTiOnUsErStAtEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementHardwareConfigurationIdUserRunStateIDInsensitively(v.Input)
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

		if actual.HardwareConfigurationUserStateId != v.Expected.HardwareConfigurationUserStateId {
			t.Fatalf("Expected %q but got %q for HardwareConfigurationUserStateId", v.Expected.HardwareConfigurationUserStateId, actual.HardwareConfigurationUserStateId)
		}

	}
}

func TestSegmentsForDeviceManagementHardwareConfigurationIdUserRunStateId(t *testing.T) {
	segments := DeviceManagementHardwareConfigurationIdUserRunStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementHardwareConfigurationIdUserRunStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
