package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdDeviceSettingStateSummaryId{}

func TestNewDeviceManagementIntentIdDeviceSettingStateSummaryID(t *testing.T) {
	id := NewDeviceManagementIntentIdDeviceSettingStateSummaryID("deviceManagementIntentIdValue", "deviceManagementIntentDeviceSettingStateSummaryIdValue")

	if id.DeviceManagementIntentId != "deviceManagementIntentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentId'", id.DeviceManagementIntentId, "deviceManagementIntentIdValue")
	}

	if id.DeviceManagementIntentDeviceSettingStateSummaryId != "deviceManagementIntentDeviceSettingStateSummaryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentDeviceSettingStateSummaryId'", id.DeviceManagementIntentDeviceSettingStateSummaryId, "deviceManagementIntentDeviceSettingStateSummaryIdValue")
	}
}

func TestFormatDeviceManagementIntentIdDeviceSettingStateSummaryID(t *testing.T) {
	actual := NewDeviceManagementIntentIdDeviceSettingStateSummaryID("deviceManagementIntentIdValue", "deviceManagementIntentDeviceSettingStateSummaryIdValue").ID()
	expected := "/deviceManagement/intents/deviceManagementIntentIdValue/deviceSettingStateSummaries/deviceManagementIntentDeviceSettingStateSummaryIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementIntentIdDeviceSettingStateSummaryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntentIdDeviceSettingStateSummaryId
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
			Input: "/deviceManagement/intents",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/deviceSettingStateSummaries",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/deviceSettingStateSummaries/deviceManagementIntentDeviceSettingStateSummaryIdValue",
			Expected: &DeviceManagementIntentIdDeviceSettingStateSummaryId{
				DeviceManagementIntentId:                          "deviceManagementIntentIdValue",
				DeviceManagementIntentDeviceSettingStateSummaryId: "deviceManagementIntentDeviceSettingStateSummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/deviceSettingStateSummaries/deviceManagementIntentDeviceSettingStateSummaryIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntentIdDeviceSettingStateSummaryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementIntentId != v.Expected.DeviceManagementIntentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentId", v.Expected.DeviceManagementIntentId, actual.DeviceManagementIntentId)
		}

		if actual.DeviceManagementIntentDeviceSettingStateSummaryId != v.Expected.DeviceManagementIntentDeviceSettingStateSummaryId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentDeviceSettingStateSummaryId", v.Expected.DeviceManagementIntentDeviceSettingStateSummaryId, actual.DeviceManagementIntentDeviceSettingStateSummaryId)
		}

	}
}

func TestParseDeviceManagementIntentIdDeviceSettingStateSummaryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntentIdDeviceSettingStateSummaryId
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
			Input: "/deviceManagement/intents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/deviceSettingStateSummaries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/dEvIcEsEtTiNgStAtEsUmMaRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/deviceSettingStateSummaries/deviceManagementIntentDeviceSettingStateSummaryIdValue",
			Expected: &DeviceManagementIntentIdDeviceSettingStateSummaryId{
				DeviceManagementIntentId:                          "deviceManagementIntentIdValue",
				DeviceManagementIntentDeviceSettingStateSummaryId: "deviceManagementIntentDeviceSettingStateSummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/deviceSettingStateSummaries/deviceManagementIntentDeviceSettingStateSummaryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/dEvIcEsEtTiNgStAtEsUmMaRiEs/dEvIcEmAnAgEmEnTiNtEnTdEvIcEsEtTiNgStAtEsUmMaRyIdVaLuE",
			Expected: &DeviceManagementIntentIdDeviceSettingStateSummaryId{
				DeviceManagementIntentId:                          "dEvIcEmAnAgEmEnTiNtEnTiDvAlUe",
				DeviceManagementIntentDeviceSettingStateSummaryId: "dEvIcEmAnAgEmEnTiNtEnTdEvIcEsEtTiNgStAtEsUmMaRyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/dEvIcEsEtTiNgStAtEsUmMaRiEs/dEvIcEmAnAgEmEnTiNtEnTdEvIcEsEtTiNgStAtEsUmMaRyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntentIdDeviceSettingStateSummaryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementIntentId != v.Expected.DeviceManagementIntentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentId", v.Expected.DeviceManagementIntentId, actual.DeviceManagementIntentId)
		}

		if actual.DeviceManagementIntentDeviceSettingStateSummaryId != v.Expected.DeviceManagementIntentDeviceSettingStateSummaryId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentDeviceSettingStateSummaryId", v.Expected.DeviceManagementIntentDeviceSettingStateSummaryId, actual.DeviceManagementIntentDeviceSettingStateSummaryId)
		}

	}
}

func TestSegmentsForDeviceManagementIntentIdDeviceSettingStateSummaryId(t *testing.T) {
	segments := DeviceManagementIntentIdDeviceSettingStateSummaryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementIntentIdDeviceSettingStateSummaryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
