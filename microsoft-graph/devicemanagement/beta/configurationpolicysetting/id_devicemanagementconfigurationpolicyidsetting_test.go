package configurationpolicysetting

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationPolicyIdSettingId{}

func TestNewDeviceManagementConfigurationPolicyIdSettingID(t *testing.T) {
	id := NewDeviceManagementConfigurationPolicyIdSettingID("deviceManagementConfigurationPolicyIdValue", "deviceManagementConfigurationSettingIdValue")

	if id.DeviceManagementConfigurationPolicyId != "deviceManagementConfigurationPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationPolicyId'", id.DeviceManagementConfigurationPolicyId, "deviceManagementConfigurationPolicyIdValue")
	}

	if id.DeviceManagementConfigurationSettingId != "deviceManagementConfigurationSettingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationSettingId'", id.DeviceManagementConfigurationSettingId, "deviceManagementConfigurationSettingIdValue")
	}
}

func TestFormatDeviceManagementConfigurationPolicyIdSettingID(t *testing.T) {
	actual := NewDeviceManagementConfigurationPolicyIdSettingID("deviceManagementConfigurationPolicyIdValue", "deviceManagementConfigurationSettingIdValue").ID()
	expected := "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue"
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
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue",
			Expected: &DeviceManagementConfigurationPolicyIdSettingId{
				DeviceManagementConfigurationPolicyId:  "deviceManagementConfigurationPolicyIdValue",
				DeviceManagementConfigurationSettingId: "deviceManagementConfigurationSettingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/extra",
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
			// Valid URI
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue",
			Expected: &DeviceManagementConfigurationPolicyIdSettingId{
				DeviceManagementConfigurationPolicyId:  "deviceManagementConfigurationPolicyIdValue",
				DeviceManagementConfigurationSettingId: "deviceManagementConfigurationSettingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationPolicies/deviceManagementConfigurationPolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE/sEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiDvAlUe",
			Expected: &DeviceManagementConfigurationPolicyIdSettingId{
				DeviceManagementConfigurationPolicyId:  "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE",
				DeviceManagementConfigurationSettingId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnPoLiCiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnPoLiCyIdVaLuE/sEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiDvAlUe/extra",
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
