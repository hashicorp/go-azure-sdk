package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId{}

func TestNewDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID(t *testing.T) {
	id := NewDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID("deviceCompliancePolicyIdValue", "settingStateDeviceSummaryIdValue")

	if id.DeviceCompliancePolicyId != "deviceCompliancePolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCompliancePolicyId'", id.DeviceCompliancePolicyId, "deviceCompliancePolicyIdValue")
	}

	if id.SettingStateDeviceSummaryId != "settingStateDeviceSummaryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SettingStateDeviceSummaryId'", id.SettingStateDeviceSummaryId, "settingStateDeviceSummaryIdValue")
	}
}

func TestFormatDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID(t *testing.T) {
	actual := NewDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID("deviceCompliancePolicyIdValue", "settingStateDeviceSummaryIdValue").ID()
	expected := "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/deviceSettingStateSummaries/settingStateDeviceSummaryIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId
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
			Input: "/deviceManagement/deviceCompliancePolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/deviceSettingStateSummaries",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/deviceSettingStateSummaries/settingStateDeviceSummaryIdValue",
			Expected: &DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId{
				DeviceCompliancePolicyId:    "deviceCompliancePolicyIdValue",
				SettingStateDeviceSummaryId: "settingStateDeviceSummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/deviceSettingStateSummaries/settingStateDeviceSummaryIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceCompliancePolicyId != v.Expected.DeviceCompliancePolicyId {
			t.Fatalf("Expected %q but got %q for DeviceCompliancePolicyId", v.Expected.DeviceCompliancePolicyId, actual.DeviceCompliancePolicyId)
		}

		if actual.SettingStateDeviceSummaryId != v.Expected.SettingStateDeviceSummaryId {
			t.Fatalf("Expected %q but got %q for SettingStateDeviceSummaryId", v.Expected.SettingStateDeviceSummaryId, actual.SettingStateDeviceSummaryId)
		}

	}
}

func TestParseDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId
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
			Input: "/deviceManagement/deviceCompliancePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/deviceSettingStateSummaries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiDvAlUe/dEvIcEsEtTiNgStAtEsUmMaRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/deviceSettingStateSummaries/settingStateDeviceSummaryIdValue",
			Expected: &DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId{
				DeviceCompliancePolicyId:    "deviceCompliancePolicyIdValue",
				SettingStateDeviceSummaryId: "settingStateDeviceSummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/deviceSettingStateSummaries/settingStateDeviceSummaryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiDvAlUe/dEvIcEsEtTiNgStAtEsUmMaRiEs/sEtTiNgStAtEdEvIcEsUmMaRyIdVaLuE",
			Expected: &DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId{
				DeviceCompliancePolicyId:    "dEvIcEcOmPlIaNcEpOlIcYiDvAlUe",
				SettingStateDeviceSummaryId: "sEtTiNgStAtEdEvIcEsUmMaRyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiDvAlUe/dEvIcEsEtTiNgStAtEsUmMaRiEs/sEtTiNgStAtEdEvIcEsUmMaRyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceCompliancePolicyId != v.Expected.DeviceCompliancePolicyId {
			t.Fatalf("Expected %q but got %q for DeviceCompliancePolicyId", v.Expected.DeviceCompliancePolicyId, actual.DeviceCompliancePolicyId)
		}

		if actual.SettingStateDeviceSummaryId != v.Expected.SettingStateDeviceSummaryId {
			t.Fatalf("Expected %q but got %q for SettingStateDeviceSummaryId", v.Expected.SettingStateDeviceSummaryId, actual.SettingStateDeviceSummaryId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId(t *testing.T) {
	segments := DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
