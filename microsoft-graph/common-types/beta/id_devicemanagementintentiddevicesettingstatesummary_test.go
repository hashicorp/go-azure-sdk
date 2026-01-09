package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdDeviceSettingStateSummaryId{}

func TestNewDeviceManagementIntentIdDeviceSettingStateSummaryID(t *testing.T) {
	id := NewDeviceManagementIntentIdDeviceSettingStateSummaryID("deviceManagementIntentId", "deviceManagementIntentDeviceSettingStateSummaryId")

	if id.DeviceManagementIntentId != "deviceManagementIntentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentId'", id.DeviceManagementIntentId, "deviceManagementIntentId")
	}

	if id.DeviceManagementIntentDeviceSettingStateSummaryId != "deviceManagementIntentDeviceSettingStateSummaryId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentDeviceSettingStateSummaryId'", id.DeviceManagementIntentDeviceSettingStateSummaryId, "deviceManagementIntentDeviceSettingStateSummaryId")
	}
}

func TestFormatDeviceManagementIntentIdDeviceSettingStateSummaryID(t *testing.T) {
	actual := NewDeviceManagementIntentIdDeviceSettingStateSummaryID("deviceManagementIntentId", "deviceManagementIntentDeviceSettingStateSummaryId").ID()
	expected := "/deviceManagement/intents/deviceManagementIntentId/deviceSettingStateSummaries/deviceManagementIntentDeviceSettingStateSummaryId"
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
			Input: "/deviceManagement/intents/deviceManagementIntentId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/deviceSettingStateSummaries",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/deviceSettingStateSummaries/deviceManagementIntentDeviceSettingStateSummaryId",
			Expected: &DeviceManagementIntentIdDeviceSettingStateSummaryId{
				DeviceManagementIntentId:                          "deviceManagementIntentId",
				DeviceManagementIntentDeviceSettingStateSummaryId: "deviceManagementIntentDeviceSettingStateSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentId/deviceSettingStateSummaries/deviceManagementIntentDeviceSettingStateSummaryId/extra",
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
			Input: "/deviceManagement/intents/deviceManagementIntentId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/deviceSettingStateSummaries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiD/dEvIcEsEtTiNgStAtEsUmMaRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/deviceSettingStateSummaries/deviceManagementIntentDeviceSettingStateSummaryId",
			Expected: &DeviceManagementIntentIdDeviceSettingStateSummaryId{
				DeviceManagementIntentId:                          "deviceManagementIntentId",
				DeviceManagementIntentDeviceSettingStateSummaryId: "deviceManagementIntentDeviceSettingStateSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentId/deviceSettingStateSummaries/deviceManagementIntentDeviceSettingStateSummaryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiD/dEvIcEsEtTiNgStAtEsUmMaRiEs/dEvIcEmAnAgEmEnTiNtEnTdEvIcEsEtTiNgStAtEsUmMaRyId",
			Expected: &DeviceManagementIntentIdDeviceSettingStateSummaryId{
				DeviceManagementIntentId:                          "dEvIcEmAnAgEmEnTiNtEnTiD",
				DeviceManagementIntentDeviceSettingStateSummaryId: "dEvIcEmAnAgEmEnTiNtEnTdEvIcEsEtTiNgStAtEsUmMaRyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiD/dEvIcEsEtTiNgStAtEsUmMaRiEs/dEvIcEmAnAgEmEnTiNtEnTdEvIcEsEtTiNgStAtEsUmMaRyId/extra",
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
