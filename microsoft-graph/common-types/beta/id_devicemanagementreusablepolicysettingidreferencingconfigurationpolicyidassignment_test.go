package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId{}

func TestNewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID("deviceManagementReusablePolicySettingId", "deviceManagementConfigurationPolicyId", "deviceManagementConfigurationPolicyAssignmentId")

	if id.DeviceManagementReusablePolicySettingId != "deviceManagementReusablePolicySettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementReusablePolicySettingId'", id.DeviceManagementReusablePolicySettingId, "deviceManagementReusablePolicySettingId")
	}

	if id.DeviceManagementConfigurationPolicyId != "deviceManagementConfigurationPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyId'", id.DeviceManagementConfigurationPolicyId, "deviceManagementConfigurationPolicyId")
	}

	if id.DeviceManagementConfigurationPolicyAssignmentId != "deviceManagementConfigurationPolicyAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyAssignmentId'", id.DeviceManagementConfigurationPolicyAssignmentId, "deviceManagementConfigurationPolicyAssignmentId")
	}
}

func TestFormatDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID("deviceManagementReusablePolicySettingId", "deviceManagementConfigurationPolicyId", "deviceManagementConfigurationPolicyAssignmentId").ID()
	expected := "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId/assignments/deviceManagementConfigurationPolicyAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId
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
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId/assignments/deviceManagementConfigurationPolicyAssignmentId",
			Expected: &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId{
				DeviceManagementReusablePolicySettingId:         "deviceManagementReusablePolicySettingId",
				DeviceManagementConfigurationPolicyId:           "deviceManagementConfigurationPolicyId",
				DeviceManagementConfigurationPolicyAssignmentId: "deviceManagementConfigurationPolicyAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId/assignments/deviceManagementConfigurationPolicyAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID(v.Input)
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

		if actual.DeviceManagementConfigurationPolicyAssignmentId != v.Expected.DeviceManagementConfigurationPolicyAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationPolicyAssignmentId", v.Expected.DeviceManagementConfigurationPolicyAssignmentId, actual.DeviceManagementConfigurationPolicyAssignmentId)
		}

	}
}

func TestParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId
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
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId/assignments/deviceManagementConfigurationPolicyAssignmentId",
			Expected: &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId{
				DeviceManagementReusablePolicySettingId:         "deviceManagementReusablePolicySettingId",
				DeviceManagementConfigurationPolicyId:           "deviceManagementConfigurationPolicyId",
				DeviceManagementConfigurationPolicyAssignmentId: "deviceManagementConfigurationPolicyAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingId/referencingConfigurationPolicies/deviceManagementConfigurationPolicyId/assignments/deviceManagementConfigurationPolicyAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId/aSsIgNmEnTs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyAsSiGnMeNtId",
			Expected: &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId{
				DeviceManagementReusablePolicySettingId:         "dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId",
				DeviceManagementConfigurationPolicyId:           "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId",
				DeviceManagementConfigurationPolicyAssignmentId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgId/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId/aSsIgNmEnTs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentIDInsensitively(v.Input)
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

		if actual.DeviceManagementConfigurationPolicyAssignmentId != v.Expected.DeviceManagementConfigurationPolicyAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationPolicyAssignmentId", v.Expected.DeviceManagementConfigurationPolicyAssignmentId, actual.DeviceManagementConfigurationPolicyAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId(t *testing.T) {
	segments := DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
