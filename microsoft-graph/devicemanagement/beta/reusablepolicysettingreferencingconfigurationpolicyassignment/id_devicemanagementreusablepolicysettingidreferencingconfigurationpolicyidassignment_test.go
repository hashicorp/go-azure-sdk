package reusablepolicysettingreferencingconfigurationpolicyassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId{}

func TestNewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID("deviceManagementReusablePolicySettingIdValue", "deviceManagementConfigurationPolicyIdValue", "deviceManagementConfigurationPolicyAssignmentIdValue")

	if id.DeviceManagementReusablePolicySettingId != "deviceManagementReusablePolicySettingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementReusablePolicySettingId'", id.DeviceManagementReusablePolicySettingId, "deviceManagementReusablePolicySettingIdValue")
	}

	if id.DeviceManagementConfigurationPolicyId != "deviceManagementConfigurationPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyId'", id.DeviceManagementConfigurationPolicyId, "deviceManagementConfigurationPolicyIdValue")
	}

	if id.DeviceManagementConfigurationPolicyAssignmentId != "deviceManagementConfigurationPolicyAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyAssignmentId'", id.DeviceManagementConfigurationPolicyAssignmentId, "deviceManagementConfigurationPolicyAssignmentIdValue")
	}
}

func TestFormatDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentID("deviceManagementReusablePolicySettingIdValue", "deviceManagementConfigurationPolicyIdValue", "deviceManagementConfigurationPolicyAssignmentIdValue").ID()
	expected := "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/assignments/deviceManagementConfigurationPolicyAssignmentIdValue"
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
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/assignments/deviceManagementConfigurationPolicyAssignmentIdValue",
			Expected: &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId{
				DeviceManagementReusablePolicySettingId:         "deviceManagementReusablePolicySettingIdValue",
				DeviceManagementConfigurationPolicyId:           "deviceManagementConfigurationPolicyIdValue",
				DeviceManagementConfigurationPolicyAssignmentId: "deviceManagementConfigurationPolicyAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/assignments/deviceManagementConfigurationPolicyAssignmentIdValue/extra",
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
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgIdVaLuE/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/assignments/deviceManagementConfigurationPolicyAssignmentIdValue",
			Expected: &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId{
				DeviceManagementReusablePolicySettingId:         "deviceManagementReusablePolicySettingIdValue",
				DeviceManagementConfigurationPolicyId:           "deviceManagementConfigurationPolicyIdValue",
				DeviceManagementConfigurationPolicyAssignmentId: "deviceManagementConfigurationPolicyAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reusablePolicySettings/deviceManagementReusablePolicySettingIdValue/referencingConfigurationPolicies/deviceManagementConfigurationPolicyIdValue/assignments/deviceManagementConfigurationPolicyAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgIdVaLuE/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE/aSsIgNmEnTs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyAsSiGnMeNtIdVaLuE",
			Expected: &DeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyIdAssignmentId{
				DeviceManagementReusablePolicySettingId:         "dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgIdVaLuE",
				DeviceManagementConfigurationPolicyId:           "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE",
				DeviceManagementConfigurationPolicyAssignmentId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyAsSiGnMeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEuSaBlEpOlIcYsEtTiNgS/dEvIcEmAnAgEmEnTrEuSaBlEpOlIcYsEtTiNgIdVaLuE/rEfErEnCiNgCoNfIgUrAtIoNpOlIcIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE/aSsIgNmEnTs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyAsSiGnMeNtIdVaLuE/extra",
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
