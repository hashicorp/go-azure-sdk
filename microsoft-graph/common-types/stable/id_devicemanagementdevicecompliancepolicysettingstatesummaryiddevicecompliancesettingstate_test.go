package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId{}

func TestNewDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateID(t *testing.T) {
	id := NewDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateID("deviceCompliancePolicySettingStateSummaryIdValue", "deviceComplianceSettingStateIdValue")

	if id.DeviceCompliancePolicySettingStateSummaryId != "deviceCompliancePolicySettingStateSummaryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCompliancePolicySettingStateSummaryId'", id.DeviceCompliancePolicySettingStateSummaryId, "deviceCompliancePolicySettingStateSummaryIdValue")
	}

	if id.DeviceComplianceSettingStateId != "deviceComplianceSettingStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceComplianceSettingStateId'", id.DeviceComplianceSettingStateId, "deviceComplianceSettingStateIdValue")
	}
}

func TestFormatDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateID(t *testing.T) {
	actual := NewDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateID("deviceCompliancePolicySettingStateSummaryIdValue", "deviceComplianceSettingStateIdValue").ID()
	expected := "/deviceManagement/deviceCompliancePolicySettingStateSummaries/deviceCompliancePolicySettingStateSummaryIdValue/deviceComplianceSettingStates/deviceComplianceSettingStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId
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
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicySettingStateSummaries/deviceCompliancePolicySettingStateSummaryIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicySettingStateSummaries/deviceCompliancePolicySettingStateSummaryIdValue/deviceComplianceSettingStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicySettingStateSummaries/deviceCompliancePolicySettingStateSummaryIdValue/deviceComplianceSettingStates/deviceComplianceSettingStateIdValue",
			Expected: &DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId{
				DeviceCompliancePolicySettingStateSummaryId: "deviceCompliancePolicySettingStateSummaryIdValue",
				DeviceComplianceSettingStateId:              "deviceComplianceSettingStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicySettingStateSummaries/deviceCompliancePolicySettingStateSummaryIdValue/deviceComplianceSettingStates/deviceComplianceSettingStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateID(v.Input)
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

		if actual.DeviceComplianceSettingStateId != v.Expected.DeviceComplianceSettingStateId {
			t.Fatalf("Expected %q but got %q for DeviceComplianceSettingStateId", v.Expected.DeviceComplianceSettingStateId, actual.DeviceComplianceSettingStateId)
		}

	}
}

func TestParseDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId
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
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicySettingStateSummaries/deviceCompliancePolicySettingStateSummaryIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcYsEtTiNgStAtEsUmMaRiEs/dEvIcEcOmPlIaNcEpOlIcYsEtTiNgStAtEsUmMaRyIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicySettingStateSummaries/deviceCompliancePolicySettingStateSummaryIdValue/deviceComplianceSettingStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcYsEtTiNgStAtEsUmMaRiEs/dEvIcEcOmPlIaNcEpOlIcYsEtTiNgStAtEsUmMaRyIdVaLuE/dEvIcEcOmPlIaNcEsEtTiNgStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicySettingStateSummaries/deviceCompliancePolicySettingStateSummaryIdValue/deviceComplianceSettingStates/deviceComplianceSettingStateIdValue",
			Expected: &DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId{
				DeviceCompliancePolicySettingStateSummaryId: "deviceCompliancePolicySettingStateSummaryIdValue",
				DeviceComplianceSettingStateId:              "deviceComplianceSettingStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicySettingStateSummaries/deviceCompliancePolicySettingStateSummaryIdValue/deviceComplianceSettingStates/deviceComplianceSettingStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcYsEtTiNgStAtEsUmMaRiEs/dEvIcEcOmPlIaNcEpOlIcYsEtTiNgStAtEsUmMaRyIdVaLuE/dEvIcEcOmPlIaNcEsEtTiNgStAtEs/dEvIcEcOmPlIaNcEsEtTiNgStAtEiDvAlUe",
			Expected: &DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId{
				DeviceCompliancePolicySettingStateSummaryId: "dEvIcEcOmPlIaNcEpOlIcYsEtTiNgStAtEsUmMaRyIdVaLuE",
				DeviceComplianceSettingStateId:              "dEvIcEcOmPlIaNcEsEtTiNgStAtEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcYsEtTiNgStAtEsUmMaRiEs/dEvIcEcOmPlIaNcEpOlIcYsEtTiNgStAtEsUmMaRyIdVaLuE/dEvIcEcOmPlIaNcEsEtTiNgStAtEs/dEvIcEcOmPlIaNcEsEtTiNgStAtEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateIDInsensitively(v.Input)
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

		if actual.DeviceComplianceSettingStateId != v.Expected.DeviceComplianceSettingStateId {
			t.Fatalf("Expected %q but got %q for DeviceComplianceSettingStateId", v.Expected.DeviceComplianceSettingStateId, actual.DeviceComplianceSettingStateId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId(t *testing.T) {
	segments := DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceCompliancePolicySettingStateSummaryIdDeviceComplianceSettingStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
