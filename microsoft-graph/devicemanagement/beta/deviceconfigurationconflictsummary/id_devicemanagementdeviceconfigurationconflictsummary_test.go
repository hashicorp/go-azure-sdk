package deviceconfigurationconflictsummary

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationConflictSummaryId{}

func TestNewDeviceManagementDeviceConfigurationConflictSummaryID(t *testing.T) {
	id := NewDeviceManagementDeviceConfigurationConflictSummaryID("deviceConfigurationConflictSummaryIdValue")

	if id.DeviceConfigurationConflictSummaryId != "deviceConfigurationConflictSummaryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationConflictSummaryId'", id.DeviceConfigurationConflictSummaryId, "deviceConfigurationConflictSummaryIdValue")
	}
}

func TestFormatDeviceManagementDeviceConfigurationConflictSummaryID(t *testing.T) {
	actual := NewDeviceManagementDeviceConfigurationConflictSummaryID("deviceConfigurationConflictSummaryIdValue").ID()
	expected := "/deviceManagement/deviceConfigurationConflictSummary/deviceConfigurationConflictSummaryIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceConfigurationConflictSummaryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationConflictSummaryId
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
			Input: "/deviceManagement/deviceConfigurationConflictSummary",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurationConflictSummary/deviceConfigurationConflictSummaryIdValue",
			Expected: &DeviceManagementDeviceConfigurationConflictSummaryId{
				DeviceConfigurationConflictSummaryId: "deviceConfigurationConflictSummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurationConflictSummary/deviceConfigurationConflictSummaryIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationConflictSummaryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceConfigurationConflictSummaryId != v.Expected.DeviceConfigurationConflictSummaryId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationConflictSummaryId", v.Expected.DeviceConfigurationConflictSummaryId, actual.DeviceConfigurationConflictSummaryId)
		}

	}
}

func TestParseDeviceManagementDeviceConfigurationConflictSummaryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationConflictSummaryId
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
			Input: "/deviceManagement/deviceConfigurationConflictSummary",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnCoNfLiCtSuMmArY",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurationConflictSummary/deviceConfigurationConflictSummaryIdValue",
			Expected: &DeviceManagementDeviceConfigurationConflictSummaryId{
				DeviceConfigurationConflictSummaryId: "deviceConfigurationConflictSummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurationConflictSummary/deviceConfigurationConflictSummaryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnCoNfLiCtSuMmArY/dEvIcEcOnFiGuRaTiOnCoNfLiCtSuMmArYiDvAlUe",
			Expected: &DeviceManagementDeviceConfigurationConflictSummaryId{
				DeviceConfigurationConflictSummaryId: "dEvIcEcOnFiGuRaTiOnCoNfLiCtSuMmArYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnCoNfLiCtSuMmArY/dEvIcEcOnFiGuRaTiOnCoNfLiCtSuMmArYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationConflictSummaryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceConfigurationConflictSummaryId != v.Expected.DeviceConfigurationConflictSummaryId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationConflictSummaryId", v.Expected.DeviceConfigurationConflictSummaryId, actual.DeviceConfigurationConflictSummaryId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceConfigurationConflictSummaryId(t *testing.T) {
	segments := DeviceManagementDeviceConfigurationConflictSummaryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceConfigurationConflictSummaryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
