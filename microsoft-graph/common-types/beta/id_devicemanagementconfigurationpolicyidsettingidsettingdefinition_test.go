package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId{}

func TestNewDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionID(t *testing.T) {
	id := NewDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionID("deviceManagementConfigurationPolicyIdValue", "deviceManagementConfigurationSettingIdValue", "deviceManagementConfigurationSettingDefinitionIdValue")

	if id.DeviceManagementConfigurationPolicyId != "deviceManagementConfigurationPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyId'", id.DeviceManagementConfigurationPolicyId, "deviceManagementConfigurationPolicyIdValue")
	}

	if id.DeviceManagementConfigurationSettingId != "deviceManagementConfigurationSettingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationSettingId'", id.DeviceManagementConfigurationSettingId, "deviceManagementConfigurationSettingIdValue")
	}

	if id.DeviceManagementConfigurationSettingDefinitionId != "deviceManagementConfigurationSettingDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationSettingDefinitionId'", id.DeviceManagementConfigurationSettingDefinitionId, "deviceManagementConfigurationSettingDefinitionIdValue")
	}
}

func TestFormatDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionID(t *testing.T) {
	actual := NewDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionID("deviceManagementConfigurationPolicyIdValue", "deviceManagementConfigurationSettingIdValue", "deviceManagementConfigurationSettingDefinitionIdValue").ID()
	expected := "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/settingDefinitions/deviceManagementConfigurationSettingDefinitionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId
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
			Input: "/deviceManagement/configurationPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/settingDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/settingDefinitions/deviceManagementConfigurationSettingDefinitionIdValue",
			Expected: &DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId{
				DeviceManagementConfigurationPolicyId:            "deviceManagementConfigurationPolicyIdValue",
				DeviceManagementConfigurationSettingId:           "deviceManagementConfigurationSettingIdValue",
				DeviceManagementConfigurationSettingDefinitionId: "deviceManagementConfigurationSettingDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/settingDefinitions/deviceManagementConfigurationSettingDefinitionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationPolicyId != v.Expected.DeviceManagementConfigurationPolicyId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationPolicyId", v.Expected.DeviceManagementConfigurationPolicyId, actual.DeviceManagementConfigurationPolicyId)
		}

		if actual.DeviceManagementConfigurationSettingId != v.Expected.DeviceManagementConfigurationSettingId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingId", v.Expected.DeviceManagementConfigurationSettingId, actual.DeviceManagementConfigurationSettingId)
		}

		if actual.DeviceManagementConfigurationSettingDefinitionId != v.Expected.DeviceManagementConfigurationSettingDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingDefinitionId", v.Expected.DeviceManagementConfigurationSettingDefinitionId, actual.DeviceManagementConfigurationSettingDefinitionId)
		}

	}
}

func TestParseDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId
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
			Input: "/deviceManagement/configurationPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE/sEtTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE/sEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/settingDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE/sEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiDvAlUe/sEtTiNgDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/settingDefinitions/deviceManagementConfigurationSettingDefinitionIdValue",
			Expected: &DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId{
				DeviceManagementConfigurationPolicyId:            "deviceManagementConfigurationPolicyIdValue",
				DeviceManagementConfigurationSettingId:           "deviceManagementConfigurationSettingIdValue",
				DeviceManagementConfigurationSettingDefinitionId: "deviceManagementConfigurationSettingDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/settingDefinitions/deviceManagementConfigurationSettingDefinitionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE/sEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiDvAlUe/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiDvAlUe",
			Expected: &DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId{
				DeviceManagementConfigurationPolicyId:            "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE",
				DeviceManagementConfigurationSettingId:           "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiDvAlUe",
				DeviceManagementConfigurationSettingDefinitionId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE/sEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiDvAlUe/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationPolicyId != v.Expected.DeviceManagementConfigurationPolicyId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationPolicyId", v.Expected.DeviceManagementConfigurationPolicyId, actual.DeviceManagementConfigurationPolicyId)
		}

		if actual.DeviceManagementConfigurationSettingId != v.Expected.DeviceManagementConfigurationSettingId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingId", v.Expected.DeviceManagementConfigurationSettingId, actual.DeviceManagementConfigurationSettingId)
		}

		if actual.DeviceManagementConfigurationSettingDefinitionId != v.Expected.DeviceManagementConfigurationSettingDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingDefinitionId", v.Expected.DeviceManagementConfigurationSettingDefinitionId, actual.DeviceManagementConfigurationSettingDefinitionId)
		}

	}
}

func TestSegmentsForDeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId(t *testing.T) {
	segments := DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementConfigurationPolicyIdSettingIdSettingDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
