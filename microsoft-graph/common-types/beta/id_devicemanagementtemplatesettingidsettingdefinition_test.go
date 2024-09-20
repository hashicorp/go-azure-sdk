package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateSettingIdSettingDefinitionId{}

func TestNewDeviceManagementTemplateSettingIdSettingDefinitionID(t *testing.T) {
	id := NewDeviceManagementTemplateSettingIdSettingDefinitionID("deviceManagementConfigurationSettingTemplateId", "deviceManagementConfigurationSettingDefinitionId")

	if id.DeviceManagementConfigurationSettingTemplateId != "deviceManagementConfigurationSettingTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationSettingTemplateId'", id.DeviceManagementConfigurationSettingTemplateId, "deviceManagementConfigurationSettingTemplateId")
	}

	if id.DeviceManagementConfigurationSettingDefinitionId != "deviceManagementConfigurationSettingDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationSettingDefinitionId'", id.DeviceManagementConfigurationSettingDefinitionId, "deviceManagementConfigurationSettingDefinitionId")
	}
}

func TestFormatDeviceManagementTemplateSettingIdSettingDefinitionID(t *testing.T) {
	actual := NewDeviceManagementTemplateSettingIdSettingDefinitionID("deviceManagementConfigurationSettingTemplateId", "deviceManagementConfigurationSettingDefinitionId").ID()
	expected := "/deviceManagement/templateSettings/deviceManagementConfigurationSettingTemplateId/settingDefinitions/deviceManagementConfigurationSettingDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementTemplateSettingIdSettingDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateSettingIdSettingDefinitionId
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
			Input: "/deviceManagement/templateSettings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templateSettings/deviceManagementConfigurationSettingTemplateId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templateSettings/deviceManagementConfigurationSettingTemplateId/settingDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templateSettings/deviceManagementConfigurationSettingTemplateId/settingDefinitions/deviceManagementConfigurationSettingDefinitionId",
			Expected: &DeviceManagementTemplateSettingIdSettingDefinitionId{
				DeviceManagementConfigurationSettingTemplateId:   "deviceManagementConfigurationSettingTemplateId",
				DeviceManagementConfigurationSettingDefinitionId: "deviceManagementConfigurationSettingDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templateSettings/deviceManagementConfigurationSettingTemplateId/settingDefinitions/deviceManagementConfigurationSettingDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateSettingIdSettingDefinitionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationSettingTemplateId != v.Expected.DeviceManagementConfigurationSettingTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingTemplateId", v.Expected.DeviceManagementConfigurationSettingTemplateId, actual.DeviceManagementConfigurationSettingTemplateId)
		}

		if actual.DeviceManagementConfigurationSettingDefinitionId != v.Expected.DeviceManagementConfigurationSettingDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingDefinitionId", v.Expected.DeviceManagementConfigurationSettingDefinitionId, actual.DeviceManagementConfigurationSettingDefinitionId)
		}

	}
}

func TestParseDeviceManagementTemplateSettingIdSettingDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateSettingIdSettingDefinitionId
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
			Input: "/deviceManagement/templateSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEsEtTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templateSettings/deviceManagementConfigurationSettingTemplateId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEsEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templateSettings/deviceManagementConfigurationSettingTemplateId/settingDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEsEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiD/sEtTiNgDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templateSettings/deviceManagementConfigurationSettingTemplateId/settingDefinitions/deviceManagementConfigurationSettingDefinitionId",
			Expected: &DeviceManagementTemplateSettingIdSettingDefinitionId{
				DeviceManagementConfigurationSettingTemplateId:   "deviceManagementConfigurationSettingTemplateId",
				DeviceManagementConfigurationSettingDefinitionId: "deviceManagementConfigurationSettingDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templateSettings/deviceManagementConfigurationSettingTemplateId/settingDefinitions/deviceManagementConfigurationSettingDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEsEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiD/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiD",
			Expected: &DeviceManagementTemplateSettingIdSettingDefinitionId{
				DeviceManagementConfigurationSettingTemplateId:   "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiD",
				DeviceManagementConfigurationSettingDefinitionId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEsEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiD/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateSettingIdSettingDefinitionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationSettingTemplateId != v.Expected.DeviceManagementConfigurationSettingTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingTemplateId", v.Expected.DeviceManagementConfigurationSettingTemplateId, actual.DeviceManagementConfigurationSettingTemplateId)
		}

		if actual.DeviceManagementConfigurationSettingDefinitionId != v.Expected.DeviceManagementConfigurationSettingDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingDefinitionId", v.Expected.DeviceManagementConfigurationSettingDefinitionId, actual.DeviceManagementConfigurationSettingDefinitionId)
		}

	}
}

func TestSegmentsForDeviceManagementTemplateSettingIdSettingDefinitionId(t *testing.T) {
	segments := DeviceManagementTemplateSettingIdSettingDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementTemplateSettingIdSettingDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
