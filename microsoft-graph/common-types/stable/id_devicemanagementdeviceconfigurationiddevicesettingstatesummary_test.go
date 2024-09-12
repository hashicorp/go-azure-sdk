package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId{}

func TestNewDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryID(t *testing.T) {
	id := NewDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryID("deviceConfigurationIdValue", "settingStateDeviceSummaryIdValue")

	if id.DeviceConfigurationId != "deviceConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationId'", id.DeviceConfigurationId, "deviceConfigurationIdValue")
	}

	if id.SettingStateDeviceSummaryId != "settingStateDeviceSummaryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SettingStateDeviceSummaryId'", id.SettingStateDeviceSummaryId, "settingStateDeviceSummaryIdValue")
	}
}

func TestFormatDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryID(t *testing.T) {
	actual := NewDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryID("deviceConfigurationIdValue", "settingStateDeviceSummaryIdValue").ID()
	expected := "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/deviceSettingStateSummaries/settingStateDeviceSummaryIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId
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
			Input: "/deviceManagement/deviceConfigurations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/deviceSettingStateSummaries",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/deviceSettingStateSummaries/settingStateDeviceSummaryIdValue",
			Expected: &DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId{
				DeviceConfigurationId:       "deviceConfigurationIdValue",
				SettingStateDeviceSummaryId: "settingStateDeviceSummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/deviceSettingStateSummaries/settingStateDeviceSummaryIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceConfigurationId != v.Expected.DeviceConfigurationId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationId", v.Expected.DeviceConfigurationId, actual.DeviceConfigurationId)
		}

		if actual.SettingStateDeviceSummaryId != v.Expected.SettingStateDeviceSummaryId {
			t.Fatalf("Expected %q but got %q for SettingStateDeviceSummaryId", v.Expected.SettingStateDeviceSummaryId, actual.SettingStateDeviceSummaryId)
		}

	}
}

func TestParseDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId
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
			Input: "/deviceManagement/deviceConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/deviceSettingStateSummaries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnIdVaLuE/dEvIcEsEtTiNgStAtEsUmMaRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/deviceSettingStateSummaries/settingStateDeviceSummaryIdValue",
			Expected: &DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId{
				DeviceConfigurationId:       "deviceConfigurationIdValue",
				SettingStateDeviceSummaryId: "settingStateDeviceSummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurations/deviceConfigurationIdValue/deviceSettingStateSummaries/settingStateDeviceSummaryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnIdVaLuE/dEvIcEsEtTiNgStAtEsUmMaRiEs/sEtTiNgStAtEdEvIcEsUmMaRyIdVaLuE",
			Expected: &DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId{
				DeviceConfigurationId:       "dEvIcEcOnFiGuRaTiOnIdVaLuE",
				SettingStateDeviceSummaryId: "sEtTiNgStAtEdEvIcEsUmMaRyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnS/dEvIcEcOnFiGuRaTiOnIdVaLuE/dEvIcEsEtTiNgStAtEsUmMaRiEs/sEtTiNgStAtEdEvIcEsUmMaRyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceConfigurationId != v.Expected.DeviceConfigurationId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationId", v.Expected.DeviceConfigurationId, actual.DeviceConfigurationId)
		}

		if actual.SettingStateDeviceSummaryId != v.Expected.SettingStateDeviceSummaryId {
			t.Fatalf("Expected %q but got %q for SettingStateDeviceSummaryId", v.Expected.SettingStateDeviceSummaryId, actual.SettingStateDeviceSummaryId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId(t *testing.T) {
	segments := DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceConfigurationIdDeviceSettingStateSummaryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
