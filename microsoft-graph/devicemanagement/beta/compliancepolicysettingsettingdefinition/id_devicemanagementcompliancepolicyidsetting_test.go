package compliancepolicysettingsettingdefinition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCompliancePolicyIdSettingId{}

func TestNewDeviceManagementCompliancePolicyIdSettingID(t *testing.T) {
	id := NewDeviceManagementCompliancePolicyIdSettingID("deviceManagementCompliancePolicyIdValue", "deviceManagementConfigurationSettingIdValue")

	if id.DeviceManagementCompliancePolicyId != "deviceManagementCompliancePolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementCompliancePolicyId'", id.DeviceManagementCompliancePolicyId, "deviceManagementCompliancePolicyIdValue")
	}

	if id.DeviceManagementConfigurationSettingId != "deviceManagementConfigurationSettingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationSettingId'", id.DeviceManagementConfigurationSettingId, "deviceManagementConfigurationSettingIdValue")
	}
}

func TestFormatDeviceManagementCompliancePolicyIdSettingID(t *testing.T) {
	actual := NewDeviceManagementCompliancePolicyIdSettingID("deviceManagementCompliancePolicyIdValue", "deviceManagementConfigurationSettingIdValue").ID()
	expected := "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/settings/deviceManagementConfigurationSettingIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementCompliancePolicyIdSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCompliancePolicyIdSettingId
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
			Input: "/deviceManagement/compliancePolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/settings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/settings/deviceManagementConfigurationSettingIdValue",
			Expected: &DeviceManagementCompliancePolicyIdSettingId{
				DeviceManagementCompliancePolicyId:     "deviceManagementCompliancePolicyIdValue",
				DeviceManagementConfigurationSettingId: "deviceManagementConfigurationSettingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCompliancePolicyIdSettingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementCompliancePolicyId != v.Expected.DeviceManagementCompliancePolicyId {
			t.Fatalf("Expected %q but got %q for DeviceManagementCompliancePolicyId", v.Expected.DeviceManagementCompliancePolicyId, actual.DeviceManagementCompliancePolicyId)
		}

		if actual.DeviceManagementConfigurationSettingId != v.Expected.DeviceManagementConfigurationSettingId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingId", v.Expected.DeviceManagementConfigurationSettingId, actual.DeviceManagementConfigurationSettingId)
		}

	}
}

func TestParseDeviceManagementCompliancePolicyIdSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCompliancePolicyIdSettingId
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
			Input: "/deviceManagement/compliancePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/settings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe/sEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/settings/deviceManagementConfigurationSettingIdValue",
			Expected: &DeviceManagementCompliancePolicyIdSettingId{
				DeviceManagementCompliancePolicyId:     "deviceManagementCompliancePolicyIdValue",
				DeviceManagementConfigurationSettingId: "deviceManagementConfigurationSettingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyIdValue/settings/deviceManagementConfigurationSettingIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe/sEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiDvAlUe",
			Expected: &DeviceManagementCompliancePolicyIdSettingId{
				DeviceManagementCompliancePolicyId:     "dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe",
				DeviceManagementConfigurationSettingId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiDvAlUe/sEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCompliancePolicyIdSettingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementCompliancePolicyId != v.Expected.DeviceManagementCompliancePolicyId {
			t.Fatalf("Expected %q but got %q for DeviceManagementCompliancePolicyId", v.Expected.DeviceManagementCompliancePolicyId, actual.DeviceManagementCompliancePolicyId)
		}

		if actual.DeviceManagementConfigurationSettingId != v.Expected.DeviceManagementConfigurationSettingId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingId", v.Expected.DeviceManagementConfigurationSettingId, actual.DeviceManagementConfigurationSettingId)
		}

	}
}

func TestSegmentsForDeviceManagementCompliancePolicyIdSettingId(t *testing.T) {
	segments := DeviceManagementCompliancePolicyIdSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementCompliancePolicyIdSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
