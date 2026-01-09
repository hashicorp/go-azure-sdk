package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId{}

func TestNewDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID(t *testing.T) {
	id := NewDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID("deviceManagementConfigurationPolicyTemplateId", "deviceManagementConfigurationSettingTemplateId", "deviceManagementConfigurationSettingDefinitionId")

	if id.DeviceManagementConfigurationPolicyTemplateId != "deviceManagementConfigurationPolicyTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyTemplateId'", id.DeviceManagementConfigurationPolicyTemplateId, "deviceManagementConfigurationPolicyTemplateId")
	}

	if id.DeviceManagementConfigurationSettingTemplateId != "deviceManagementConfigurationSettingTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationSettingTemplateId'", id.DeviceManagementConfigurationSettingTemplateId, "deviceManagementConfigurationSettingTemplateId")
	}

	if id.DeviceManagementConfigurationSettingDefinitionId != "deviceManagementConfigurationSettingDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationSettingDefinitionId'", id.DeviceManagementConfigurationSettingDefinitionId, "deviceManagementConfigurationSettingDefinitionId")
	}
}

func TestFormatDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID(t *testing.T) {
	actual := NewDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionID("deviceManagementConfigurationPolicyTemplateId", "deviceManagementConfigurationSettingTemplateId", "deviceManagementConfigurationSettingDefinitionId").ID()
	expected := "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/settingTemplates/deviceManagementConfigurationSettingTemplateId/settingDefinitions/deviceManagementConfigurationSettingDefinitionId"
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
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/settingTemplates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/settingTemplates/deviceManagementConfigurationSettingTemplateId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/settingTemplates/deviceManagementConfigurationSettingTemplateId/settingDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/settingTemplates/deviceManagementConfigurationSettingTemplateId/settingDefinitions/deviceManagementConfigurationSettingDefinitionId",
			Expected: &DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId{
				DeviceManagementConfigurationPolicyTemplateId:    "deviceManagementConfigurationPolicyTemplateId",
				DeviceManagementConfigurationSettingTemplateId:   "deviceManagementConfigurationSettingTemplateId",
				DeviceManagementConfigurationSettingDefinitionId: "deviceManagementConfigurationSettingDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/settingTemplates/deviceManagementConfigurationSettingTemplateId/settingDefinitions/deviceManagementConfigurationSettingDefinitionId/extra",
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
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/settingTemplates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeId/sEtTiNgTeMpLaTeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/settingTemplates/deviceManagementConfigurationSettingTemplateId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeId/sEtTiNgTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/settingTemplates/deviceManagementConfigurationSettingTemplateId/settingDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeId/sEtTiNgTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiD/sEtTiNgDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/settingTemplates/deviceManagementConfigurationSettingTemplateId/settingDefinitions/deviceManagementConfigurationSettingDefinitionId",
			Expected: &DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId{
				DeviceManagementConfigurationPolicyTemplateId:    "deviceManagementConfigurationPolicyTemplateId",
				DeviceManagementConfigurationSettingTemplateId:   "deviceManagementConfigurationSettingTemplateId",
				DeviceManagementConfigurationSettingDefinitionId: "deviceManagementConfigurationSettingDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/settingTemplates/deviceManagementConfigurationSettingTemplateId/settingDefinitions/deviceManagementConfigurationSettingDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeId/sEtTiNgTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiD/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiD",
			Expected: &DeviceManagementConfigurationPolicyTemplateIdSettingTemplateIdSettingDefinitionId{
				DeviceManagementConfigurationPolicyTemplateId:    "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeId",
				DeviceManagementConfigurationSettingTemplateId:   "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiD",
				DeviceManagementConfigurationSettingDefinitionId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeId/sEtTiNgTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiD/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiD/extra",
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
