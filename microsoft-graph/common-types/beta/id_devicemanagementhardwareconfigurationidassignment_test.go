package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementHardwareConfigurationIdAssignmentId{}

func TestNewDeviceManagementHardwareConfigurationIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementHardwareConfigurationIdAssignmentID("hardwareConfigurationId", "hardwareConfigurationAssignmentId")

	if id.HardwareConfigurationId != "hardwareConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwareConfigurationId'", id.HardwareConfigurationId, "hardwareConfigurationId")
	}

	if id.HardwareConfigurationAssignmentId != "hardwareConfigurationAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwareConfigurationAssignmentId'", id.HardwareConfigurationAssignmentId, "hardwareConfigurationAssignmentId")
	}
}

func TestFormatDeviceManagementHardwareConfigurationIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementHardwareConfigurationIdAssignmentID("hardwareConfigurationId", "hardwareConfigurationAssignmentId").ID()
	expected := "/deviceManagement/hardwareConfigurations/hardwareConfigurationId/assignments/hardwareConfigurationAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementHardwareConfigurationIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementHardwareConfigurationIdAssignmentId
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
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationId/assignments/hardwareConfigurationAssignmentId",
			Expected: &DeviceManagementHardwareConfigurationIdAssignmentId{
				HardwareConfigurationId:           "hardwareConfigurationId",
				HardwareConfigurationAssignmentId: "hardwareConfigurationAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationId/assignments/hardwareConfigurationAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementHardwareConfigurationIdAssignmentID(v.Input)
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

		if actual.HardwareConfigurationAssignmentId != v.Expected.HardwareConfigurationAssignmentId {
			t.Fatalf("Expected %q but got %q for HardwareConfigurationAssignmentId", v.Expected.HardwareConfigurationAssignmentId, actual.HardwareConfigurationAssignmentId)
		}

	}
}

func TestParseDeviceManagementHardwareConfigurationIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementHardwareConfigurationIdAssignmentId
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
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEcOnFiGuRaTiOnS/hArDwArEcOnFiGuRaTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEcOnFiGuRaTiOnS/hArDwArEcOnFiGuRaTiOnId/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationId/assignments/hardwareConfigurationAssignmentId",
			Expected: &DeviceManagementHardwareConfigurationIdAssignmentId{
				HardwareConfigurationId:           "hardwareConfigurationId",
				HardwareConfigurationAssignmentId: "hardwareConfigurationAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/hardwareConfigurations/hardwareConfigurationId/assignments/hardwareConfigurationAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEcOnFiGuRaTiOnS/hArDwArEcOnFiGuRaTiOnId/aSsIgNmEnTs/hArDwArEcOnFiGuRaTiOnAsSiGnMeNtId",
			Expected: &DeviceManagementHardwareConfigurationIdAssignmentId{
				HardwareConfigurationId:           "hArDwArEcOnFiGuRaTiOnId",
				HardwareConfigurationAssignmentId: "hArDwArEcOnFiGuRaTiOnAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/hArDwArEcOnFiGuRaTiOnS/hArDwArEcOnFiGuRaTiOnId/aSsIgNmEnTs/hArDwArEcOnFiGuRaTiOnAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementHardwareConfigurationIdAssignmentIDInsensitively(v.Input)
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

		if actual.HardwareConfigurationAssignmentId != v.Expected.HardwareConfigurationAssignmentId {
			t.Fatalf("Expected %q but got %q for HardwareConfigurationAssignmentId", v.Expected.HardwareConfigurationAssignmentId, actual.HardwareConfigurationAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementHardwareConfigurationIdAssignmentId(t *testing.T) {
	segments := DeviceManagementHardwareConfigurationIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementHardwareConfigurationIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
