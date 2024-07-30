package devicecompliancepolicysettingstatesummarydevicecompliancesettingstate

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicySettingStateSummaryId{}

func TestNewDeviceManagementDeviceCompliancePolicySettingStateSummaryID(t *testing.T) {
	id := NewDeviceManagementDeviceCompliancePolicySettingStateSummaryID("deviceCompliancePolicySettingStateSummaryIdValue")

	if id.DeviceCompliancePolicySettingStateSummaryId != "deviceCompliancePolicySettingStateSummaryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCompliancePolicySettingStateSummaryId'", id.DeviceCompliancePolicySettingStateSummaryId, "deviceCompliancePolicySettingStateSummaryIdValue")
	}
}

func TestFormatDeviceManagementDeviceCompliancePolicySettingStateSummaryID(t *testing.T) {
	actual := NewDeviceManagementDeviceCompliancePolicySettingStateSummaryID("deviceCompliancePolicySettingStateSummaryIdValue").ID()
	expected := "/deviceManagement/deviceCompliancePolicySettingStateSummaries/deviceCompliancePolicySettingStateSummaryIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceCompliancePolicySettingStateSummaryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicySettingStateSummaryId
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
			Input: "/deviceManagement/deviceCompliancePolicySettingStateSummaries",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicySettingStateSummaries/deviceCompliancePolicySettingStateSummaryIdValue",
			Expected: &DeviceManagementDeviceCompliancePolicySettingStateSummaryId{
				DeviceCompliancePolicySettingStateSummaryId: "deviceCompliancePolicySettingStateSummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicySettingStateSummaries/deviceCompliancePolicySettingStateSummaryIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicySettingStateSummaryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceCompliancePolicySettingStateSummaryId != v.Expected.DeviceCompliancePolicySettingStateSummaryId {
			t.Fatalf("Expected %q but got %q for DeviceCompliancePolicySettingStateSummaryId", v.Expected.DeviceCompliancePolicySettingStateSummaryId, actual.DeviceCompliancePolicySettingStateSummaryId)
		}

	}
}

func TestParseDeviceManagementDeviceCompliancePolicySettingStateSummaryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicySettingStateSummaryId
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
			Input: "/deviceManagement/deviceCompliancePolicySettingStateSummaries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcYsEtTiNgStAtEsUmMaRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicySettingStateSummaries/deviceCompliancePolicySettingStateSummaryIdValue",
			Expected: &DeviceManagementDeviceCompliancePolicySettingStateSummaryId{
				DeviceCompliancePolicySettingStateSummaryId: "deviceCompliancePolicySettingStateSummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicySettingStateSummaries/deviceCompliancePolicySettingStateSummaryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcYsEtTiNgStAtEsUmMaRiEs/dEvIcEcOmPlIaNcEpOlIcYsEtTiNgStAtEsUmMaRyIdVaLuE",
			Expected: &DeviceManagementDeviceCompliancePolicySettingStateSummaryId{
				DeviceCompliancePolicySettingStateSummaryId: "dEvIcEcOmPlIaNcEpOlIcYsEtTiNgStAtEsUmMaRyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcYsEtTiNgStAtEsUmMaRiEs/dEvIcEcOmPlIaNcEpOlIcYsEtTiNgStAtEsUmMaRyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicySettingStateSummaryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceCompliancePolicySettingStateSummaryId != v.Expected.DeviceCompliancePolicySettingStateSummaryId {
			t.Fatalf("Expected %q but got %q for DeviceCompliancePolicySettingStateSummaryId", v.Expected.DeviceCompliancePolicySettingStateSummaryId, actual.DeviceCompliancePolicySettingStateSummaryId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceCompliancePolicySettingStateSummaryId(t *testing.T) {
	segments := DeviceManagementDeviceCompliancePolicySettingStateSummaryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceCompliancePolicySettingStateSummaryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
