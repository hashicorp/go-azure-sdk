package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId{}

func TestNewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID(t *testing.T) {
	id := NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID("deviceManagementReusablePolicySettingId", "deviceManagementConfigurationPolicyId")

	if id.DeviceManagementReusablePolicySettingId != "deviceManagementReusablePolicySettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementReusablePolicySettingId'", id.DeviceManagementReusablePolicySettingId, "deviceManagementReusablePolicySettingId")
	}

	if id.DeviceManagementConfigurationPolicyId != "deviceManagementConfigurationPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyId'", id.DeviceManagementConfigurationPolicyId, "deviceManagementConfigurationPolicyId")
	}
}

func TestFormatDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID(t *testing.T) {
	actual := NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID("deviceManagementReusablePolicySettingId", "deviceManagementConfigurationPolicyId").ID()
	expected := "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId
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
			// Valid URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId",
			Expected: &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId{
				DeviceManagementReusablePolicySettingId: "deviceManagementReusablePolicySettingId",
				DeviceManagementConfigurationPolicyId:   "deviceManagementConfigurationPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID(v.Input)
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

	}
}

func TestParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId
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
			// Valid URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId",
			Expected: &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId{
				DeviceManagementReusablePolicySettingId: "deviceManagementReusablePolicySettingId",
				DeviceManagementConfigurationPolicyId:   "deviceManagementConfigurationPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId",
			Expected: &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId{
				DeviceManagementReusablePolicySettingId: "dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId",
				DeviceManagementConfigurationPolicyId:   "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId(t *testing.T) {
	segments := DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
