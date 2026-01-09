package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId{}

func TestNewDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID(t *testing.T) {
	id := NewDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID("deviceCompliancePolicyId", "settingStateDeviceSummaryId")

	if id.DeviceCompliancePolicyId != "deviceCompliancePolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCompliancePolicyId'", id.DeviceCompliancePolicyId, "deviceCompliancePolicyId")
	}

	if id.SettingStateDeviceSummaryId != "settingStateDeviceSummaryId" {
		t.Fatalf("Expected %q but got %q for Segment 'SettingStateDeviceSummaryId'", id.SettingStateDeviceSummaryId, "settingStateDeviceSummaryId")
	}
}

func TestFormatDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID(t *testing.T) {
	actual := NewDeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryID("deviceCompliancePolicyId", "settingStateDeviceSummaryId").ID()
	expected := "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/deviceSettingStateSummaries/settingStateDeviceSummaryId"
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
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/deviceSettingStateSummaries",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/deviceSettingStateSummaries/settingStateDeviceSummaryId",
			Expected: &DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId{
				DeviceCompliancePolicyId:    "deviceCompliancePolicyId",
				SettingStateDeviceSummaryId: "settingStateDeviceSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/deviceSettingStateSummaries/settingStateDeviceSummaryId/extra",
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
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/deviceSettingStateSummaries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD/dEvIcEsEtTiNgStAtEsUmMaRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/deviceSettingStateSummaries/settingStateDeviceSummaryId",
			Expected: &DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId{
				DeviceCompliancePolicyId:    "deviceCompliancePolicyId",
				SettingStateDeviceSummaryId: "settingStateDeviceSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/deviceSettingStateSummaries/settingStateDeviceSummaryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD/dEvIcEsEtTiNgStAtEsUmMaRiEs/sEtTiNgStAtEdEvIcEsUmMaRyId",
			Expected: &DeviceManagementDeviceCompliancePolicyIdDeviceSettingStateSummaryId{
				DeviceCompliancePolicyId:    "dEvIcEcOmPlIaNcEpOlIcYiD",
				SettingStateDeviceSummaryId: "sEtTiNgStAtEdEvIcEsUmMaRyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD/dEvIcEsEtTiNgStAtEsUmMaRiEs/sEtTiNgStAtEdEvIcEsUmMaRyId/extra",
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
