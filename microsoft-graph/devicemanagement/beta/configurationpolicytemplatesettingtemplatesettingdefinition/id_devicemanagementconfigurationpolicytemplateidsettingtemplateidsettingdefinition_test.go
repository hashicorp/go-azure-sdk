package configurationpolicytemplatesettingtemplatesettingdefinition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId{}

func TestNewDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID(t *testing.T) {
	id := NewDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID("deviceManagementConfigurationPolicyTemplateIdValue", "deviceManagementConfigurationSettingTemplateIdValue", "deviceManagementConfigurationSettingDefinitionIdValue")

	if id.DeviceManagementConfigurationPolicyTemplateId != "deviceManagementConfigurationPolicyTemplateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyTemplateId'", id.DeviceManagementConfigurationPolicyTemplateId, "deviceManagementConfigurationPolicyTemplateIdValue")
	}

	if id.DeviceManagementConfigurationSettingTemplateId != "deviceManagementConfigurationSettingTemplateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationSettingTemplateId'", id.DeviceManagementConfigurationSettingTemplateId, "deviceManagementConfigurationSettingTemplateIdValue")
	}

	if id.DeviceManagementConfigurationSettingDefinitionId != "deviceManagementConfigurationSettingDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationSettingDefinitionId'", id.DeviceManagementConfigurationSettingDefinitionId, "deviceManagementConfigurationSettingDefinitionIdValue")
	}
}

func TestFormatDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID(t *testing.T) {
	actual := NewDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID("deviceManagementConfigurationPolicyTemplateIdValue", "deviceManagementConfigurationSettingTemplateIdValue", "deviceManagementConfigurationSettingDefinitionIdValue").ID()
	expected := "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateIdValue/settingTemplates/deviceManagementConfigurationSettingTemplateIdValue/settingDefinitions/deviceManagementConfigurationSettingDefinitionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId
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
			Input: "/deviceManagement/configurationPolicyTemplates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateIdValue/settingTemplates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateIdValue/settingTemplates/deviceManagementConfigurationSettingTemplateIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateIdValue/settingTemplates/deviceManagementConfigurationSettingTemplateIdValue/settingDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateIdValue/settingTemplates/deviceManagementConfigurationSettingTemplateIdValue/settingDefinitions/deviceManagementConfigurationSettingDefinitionIdValue",
			Expected: &DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId{
				DeviceManagementConfigurationPolicyTemplateId:    "deviceManagementConfigurationPolicyTemplateIdValue",
				DeviceManagementConfigurationSettingTemplateId:   "deviceManagementConfigurationSettingTemplateIdValue",
				DeviceManagementConfigurationSettingDefinitionId: "deviceManagementConfigurationSettingDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateIdValue/settingTemplates/deviceManagementConfigurationSettingTemplateIdValue/settingDefinitions/deviceManagementConfigurationSettingDefinitionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationPolicyTemplateId != v.Expected.DeviceManagementConfigurationPolicyTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationPolicyTemplateId", v.Expected.DeviceManagementConfigurationPolicyTemplateId, actual.DeviceManagementConfigurationPolicyTemplateId)
		}

		if actual.DeviceManagementConfigurationSettingTemplateId != v.Expected.DeviceManagementConfigurationSettingTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingTemplateId", v.Expected.DeviceManagementConfigurationSettingTemplateId, actual.DeviceManagementConfigurationSettingTemplateId)
		}

		if actual.DeviceManagementConfigurationSettingDefinitionId != v.Expected.DeviceManagementConfigurationSettingDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingDefinitionId", v.Expected.DeviceManagementConfigurationSettingDefinitionId, actual.DeviceManagementConfigurationSettingDefinitionId)
		}

	}
}

func TestParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId
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
			Input: "/deviceManagement/configurationPolicyTemplates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateIdValue/settingTemplates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeIdVaLuE/sEtTiNgTeMpLaTeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateIdValue/settingTemplates/deviceManagementConfigurationSettingTemplateIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeIdVaLuE/sEtTiNgTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateIdValue/settingTemplates/deviceManagementConfigurationSettingTemplateIdValue/settingDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeIdVaLuE/sEtTiNgTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiDvAlUe/sEtTiNgDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateIdValue/settingTemplates/deviceManagementConfigurationSettingTemplateIdValue/settingDefinitions/deviceManagementConfigurationSettingDefinitionIdValue",
			Expected: &DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId{
				DeviceManagementConfigurationPolicyTemplateId:    "deviceManagementConfigurationPolicyTemplateIdValue",
				DeviceManagementConfigurationSettingTemplateId:   "deviceManagementConfigurationSettingTemplateIdValue",
				DeviceManagementConfigurationSettingDefinitionId: "deviceManagementConfigurationSettingDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateIdValue/settingTemplates/deviceManagementConfigurationSettingTemplateIdValue/settingDefinitions/deviceManagementConfigurationSettingDefinitionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeIdVaLuE/sEtTiNgTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiDvAlUe/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiDvAlUe",
			Expected: &DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId{
				DeviceManagementConfigurationPolicyTemplateId:    "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeIdVaLuE",
				DeviceManagementConfigurationSettingTemplateId:   "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiDvAlUe",
				DeviceManagementConfigurationSettingDefinitionId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeIdVaLuE/sEtTiNgTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiDvAlUe/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationPolicyTemplateId != v.Expected.DeviceManagementConfigurationPolicyTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationPolicyTemplateId", v.Expected.DeviceManagementConfigurationPolicyTemplateId, actual.DeviceManagementConfigurationPolicyTemplateId)
		}

		if actual.DeviceManagementConfigurationSettingTemplateId != v.Expected.DeviceManagementConfigurationSettingTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingTemplateId", v.Expected.DeviceManagementConfigurationSettingTemplateId, actual.DeviceManagementConfigurationSettingTemplateId)
		}

		if actual.DeviceManagementConfigurationSettingDefinitionId != v.Expected.DeviceManagementConfigurationSettingDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingDefinitionId", v.Expected.DeviceManagementConfigurationSettingDefinitionId, actual.DeviceManagementConfigurationSettingDefinitionId)
		}

	}
}

func TestSegmentsForDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId(t *testing.T) {
	segments := DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
