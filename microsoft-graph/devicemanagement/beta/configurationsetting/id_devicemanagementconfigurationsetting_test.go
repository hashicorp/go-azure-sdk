package configurationsetting

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationSettingId{}

func TestNewDeviceManagementConfigurationSettingID(t *testing.T) {
	id := NewDeviceManagementConfigurationSettingID("deviceManagementConfigurationSettingDefinitionIdValue")

	if id.DeviceManagementConfigurationSettingDefinitionId != "deviceManagementConfigurationSettingDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationSettingDefinitionId'", id.DeviceManagementConfigurationSettingDefinitionId, "deviceManagementConfigurationSettingDefinitionIdValue")
	}
}

func TestFormatDeviceManagementConfigurationSettingID(t *testing.T) {
	actual := NewDeviceManagementConfigurationSettingID("deviceManagementConfigurationSettingDefinitionIdValue").ID()
	expected := "/deviceManagement/configurationSettings/deviceManagementConfigurationSettingDefinitionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementConfigurationSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigurationSettingId
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
			Input: "/deviceManagement/configurationSettings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationSettings/deviceManagementConfigurationSettingDefinitionIdValue",
			Expected: &DeviceManagementConfigurationSettingId{
				DeviceManagementConfigurationSettingDefinitionId: "deviceManagementConfigurationSettingDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationSettings/deviceManagementConfigurationSettingDefinitionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigurationSettingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationSettingDefinitionId != v.Expected.DeviceManagementConfigurationSettingDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingDefinitionId", v.Expected.DeviceManagementConfigurationSettingDefinitionId, actual.DeviceManagementConfigurationSettingDefinitionId)
		}

	}
}

func TestParseDeviceManagementConfigurationSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigurationSettingId
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
			Input: "/deviceManagement/configurationSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnSeTtInGs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationSettings/deviceManagementConfigurationSettingDefinitionIdValue",
			Expected: &DeviceManagementConfigurationSettingId{
				DeviceManagementConfigurationSettingDefinitionId: "deviceManagementConfigurationSettingDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationSettings/deviceManagementConfigurationSettingDefinitionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnSeTtInGs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiDvAlUe",
			Expected: &DeviceManagementConfigurationSettingId{
				DeviceManagementConfigurationSettingDefinitionId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnSeTtInGs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigurationSettingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationSettingDefinitionId != v.Expected.DeviceManagementConfigurationSettingDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingDefinitionId", v.Expected.DeviceManagementConfigurationSettingDefinitionId, actual.DeviceManagementConfigurationSettingDefinitionId)
		}

	}
}

func TestSegmentsForDeviceManagementConfigurationSettingId(t *testing.T) {
	segments := DeviceManagementConfigurationSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementConfigurationSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
