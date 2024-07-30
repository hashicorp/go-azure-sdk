package reusablepolicysettingreferencingconfigurationpolicysettingsettingdefinition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId{}

func TestNewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionID(t *testing.T) {
	id := NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionID("deviceManagementReusablePolicySettingIdValue", "deviceManagementConfigurationPolicyIdValue", "deviceManagementConfigurationSettingIdValue", "deviceManagementConfigurationSettingDefinitionIdValue")

	if id.DeviceManagementReusablePolicySettingId != "deviceManagementReusablePolicySettingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementReusablePolicySettingId'", id.DeviceManagementReusablePolicySettingId, "deviceManagementReusablePolicySettingIdValue")
	}

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

func TestFormatDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionID(t *testing.T) {
	actual := NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionID("deviceManagementReusablePolicySettingIdValue", "deviceManagementConfigurationPolicyIdValue", "deviceManagementConfigurationSettingIdValue", "deviceManagementConfigurationSettingDefinitionIdValue").ID()
	expected := "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/settingDefinitions/deviceManagementConfigurationSettingDefinitionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId
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
			Input: "/deviceManagement/reusablePolicySettings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/settings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/settingDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/settingDefinitions/deviceManagementConfigurationSettingDefinitionIdValue",
			Expected: &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId{
				DeviceManagementReusablePolicySettingId:          "deviceManagementReusablePolicySettingIdValue",
				DeviceManagementConfigurationPolicyId:            "deviceManagementConfigurationPolicyIdValue",
				DeviceManagementConfigurationSettingId:           "deviceManagementConfigurationSettingIdValue",
				DeviceManagementConfigurationSettingDefinitionId: "deviceManagementConfigurationSettingDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/settingDefinitions/deviceManagementConfigurationSettingDefinitionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementReusablePolicySettingId != v.Expected.DeviceManagementReusablePolicySettingId {
			t.Fatalf("Expected %q but got %q for DeviceManagementReusablePolicySettingId", v.Expected.DeviceManagementReusablePolicySettingId, actual.DeviceManagementReusablePolicySettingId)
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

func TestParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId
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
			Input: "/deviceManagement/reusablePolicySettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgIdVaLuE/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgIdVaLuE/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/settings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgIdVaLuE/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE/sEtTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgIdVaLuE/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE/sEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/settingDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgIdVaLuE/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE/sEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiDvAlUe/sEtTiNgDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/settingDefinitions/deviceManagementConfigurationSettingDefinitionIdValue",
			Expected: &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId{
				DeviceManagementReusablePolicySettingId:          "deviceManagementReusablePolicySettingIdValue",
				DeviceManagementConfigurationPolicyId:            "deviceManagementConfigurationPolicyIdValue",
				DeviceManagementConfigurationSettingId:           "deviceManagementConfigurationSettingIdValue",
				DeviceManagementConfigurationSettingDefinitionId: "deviceManagementConfigurationSettingDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/settingDefinitions/deviceManagementConfigurationSettingDefinitionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgIdVaLuE/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE/sEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiDvAlUe/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiDvAlUe",
			Expected: &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId{
				DeviceManagementReusablePolicySettingId:          "dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgIdVaLuE",
				DeviceManagementConfigurationPolicyId:            "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE",
				DeviceManagementConfigurationSettingId:           "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiDvAlUe",
				DeviceManagementConfigurationSettingDefinitionId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgIdVaLuE/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE/sEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiDvAlUe/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGdEfInItIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementReusablePolicySettingId != v.Expected.DeviceManagementReusablePolicySettingId {
			t.Fatalf("Expected %q but got %q for DeviceManagementReusablePolicySettingId", v.Expected.DeviceManagementReusablePolicySettingId, actual.DeviceManagementReusablePolicySettingId)
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

func TestSegmentsForDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId(t *testing.T) {
	segments := DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIdSettingDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
