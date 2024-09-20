package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationPolicyIdSettingId{}

func TestNewDeviceManagementConfigurationPolicyIdSettingID(t *testing.T) {
	id := NewDeviceManagementConfigurationPolicyIdSettingID("deviceManagementConfigurationPolicyId", "deviceManagementConfigurationSettingId")

	if id.DeviceManagementConfigurationPolicyId != "deviceManagementConfigurationPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyId'", id.DeviceManagementConfigurationPolicyId, "deviceManagementConfigurationPolicyId")
	}

	if id.DeviceManagementConfigurationSettingId != "deviceManagementConfigurationSettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationSettingId'", id.DeviceManagementConfigurationSettingId, "deviceManagementConfigurationSettingId")
	}
}

func TestFormatDeviceManagementConfigurationPolicyIdSettingID(t *testing.T) {
	actual := NewDeviceManagementConfigurationPolicyIdSettingID("deviceManagementConfigurationPolicyId", "deviceManagementConfigurationSettingId").ID()
	expected := "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId/settings/deviceManagementConfigurationSettingId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementConfigurationPolicyIdSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigurationPolicyIdSettingId
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
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId/settings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId/settings/deviceManagementConfigurationSettingId",
			Expected: &DeviceManagementConfigurationPolicyIdSettingId{
				DeviceManagementConfigurationPolicyId:  "deviceManagementConfigurationPolicyId",
				DeviceManagementConfigurationSettingId: "deviceManagementConfigurationSettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId/settings/deviceManagementConfigurationSettingId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigurationPolicyIdSettingID(v.Input)
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

	}
}

func TestParseDeviceManagementConfigurationPolicyIdSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigurationPolicyIdSettingId
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
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId/settings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId/sEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId/settings/deviceManagementConfigurationSettingId",
			Expected: &DeviceManagementConfigurationPolicyIdSettingId{
				DeviceManagementConfigurationPolicyId:  "deviceManagementConfigurationPolicyId",
				DeviceManagementConfigurationSettingId: "deviceManagementConfigurationSettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyId/settings/deviceManagementConfigurationSettingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId/sEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiD",
			Expected: &DeviceManagementConfigurationPolicyIdSettingId{
				DeviceManagementConfigurationPolicyId:  "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId",
				DeviceManagementConfigurationSettingId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyId/sEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigurationPolicyIdSettingIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementConfigurationPolicyIdSettingId(t *testing.T) {
	segments := DeviceManagementConfigurationPolicyIdSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementConfigurationPolicyIdSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
