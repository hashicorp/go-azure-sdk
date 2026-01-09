package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId{}

func TestNewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingID(t *testing.T) {
	id := NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingID("deviceManagementReusablePolicySettingId", "deviceManagementConfigurationPolicyId", "deviceManagementConfigurationSettingId")

	if id.DeviceManagementReusablePolicySettingId != "deviceManagementReusablePolicySettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementReusablePolicySettingId'", id.DeviceManagementReusablePolicySettingId, "deviceManagementReusablePolicySettingId")
	}

	if id.DeviceManagementConfigurationPolicyId != "deviceManagementConfigurationPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyId'", id.DeviceManagementConfigurationPolicyId, "deviceManagementConfigurationPolicyId")
	}

	if id.DeviceManagementConfigurationSettingId != "deviceManagementConfigurationSettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationSettingId'", id.DeviceManagementConfigurationSettingId, "deviceManagementConfigurationSettingId")
	}
}

func TestFormatDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingID(t *testing.T) {
	actual := NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingID("deviceManagementReusablePolicySettingId", "deviceManagementConfigurationPolicyId", "deviceManagementConfigurationSettingId").ID()
	expected := "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId/settings/deviceManagementConfigurationSettingId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId
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
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId/settings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId/settings/deviceManagementConfigurationSettingId",
			Expected: &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId{
				DeviceManagementReusablePolicySettingId: "deviceManagementReusablePolicySettingId",
				DeviceManagementConfigurationPolicyId:   "deviceManagementConfigurationPolicyId",
				DeviceManagementConfigurationSettingId:  "deviceManagementConfigurationSettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId/settings/deviceManagementConfigurationSettingId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingID(v.Input)
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

	}
}

func TestParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId
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
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId/settings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId/sEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId/settings/deviceManagementConfigurationSettingId",
			Expected: &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId{
				DeviceManagementReusablePolicySettingId: "deviceManagementReusablePolicySettingId",
				DeviceManagementConfigurationPolicyId:   "deviceManagementConfigurationPolicyId",
				DeviceManagementConfigurationSettingId:  "deviceManagementConfigurationSettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId/settings/deviceManagementConfigurationSettingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId/sEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiD",
			Expected: &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId{
				DeviceManagementReusablePolicySettingId: "dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId",
				DeviceManagementConfigurationPolicyId:   "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId",
				DeviceManagementConfigurationSettingId:  "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId/sEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId(t *testing.T) {
	segments := DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
