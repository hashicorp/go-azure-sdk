package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId{}

func TestNewDeviceManagementConfigurationPolicyTemplateIdSettingTemplateID(t *testing.T) {
	id := NewDeviceManagementConfigurationPolicyTemplateIdSettingTemplateID("deviceManagementConfigurationPolicyTemplateId", "deviceManagementConfigurationSettingTemplateId")

	if id.DeviceManagementConfigurationPolicyTemplateId != "deviceManagementConfigurationPolicyTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyTemplateId'", id.DeviceManagementConfigurationPolicyTemplateId, "deviceManagementConfigurationPolicyTemplateId")
	}

	if id.DeviceManagementConfigurationSettingTemplateId != "deviceManagementConfigurationSettingTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationSettingTemplateId'", id.DeviceManagementConfigurationSettingTemplateId, "deviceManagementConfigurationSettingTemplateId")
	}
}

func TestFormatDeviceManagementConfigurationPolicyTemplateIdSettingTemplateID(t *testing.T) {
	actual := NewDeviceManagementConfigurationPolicyTemplateIdSettingTemplateID("deviceManagementConfigurationPolicyTemplateId", "deviceManagementConfigurationSettingTemplateId").ID()
	expected := "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/settingTemplates/deviceManagementConfigurationSettingTemplateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId
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
			// Valid URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/settingTemplates/deviceManagementConfigurationSettingTemplateId",
			Expected: &DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId{
				DeviceManagementConfigurationPolicyTemplateId:  "deviceManagementConfigurationPolicyTemplateId",
				DeviceManagementConfigurationSettingTemplateId: "deviceManagementConfigurationSettingTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/settingTemplates/deviceManagementConfigurationSettingTemplateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateID(v.Input)
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

	}
}

func TestParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId
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
			// Valid URI
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/settingTemplates/deviceManagementConfigurationSettingTemplateId",
			Expected: &DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId{
				DeviceManagementConfigurationPolicyTemplateId:  "deviceManagementConfigurationPolicyTemplateId",
				DeviceManagementConfigurationSettingTemplateId: "deviceManagementConfigurationSettingTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationPolicyTemplates/deviceManagementConfigurationPolicyTemplateId/settingTemplates/deviceManagementConfigurationSettingTemplateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeId/sEtTiNgTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiD",
			Expected: &DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId{
				DeviceManagementConfigurationPolicyTemplateId:  "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeId",
				DeviceManagementConfigurationSettingTemplateId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCyTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyTeMpLaTeId/sEtTiNgTeMpLaTeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigurationPolicyTemplateIdSettingTemplateIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementConfigurationPolicyTemplateIdSettingTemplateId(t *testing.T) {
	segments := DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
