package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMonitoringAlertRecordId{}

func TestNewDeviceManagementMonitoringAlertRecordID(t *testing.T) {
	id := NewDeviceManagementMonitoringAlertRecordID("alertRecordIdValue")

	if id.AlertRecordId != "alertRecordIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AlertRecordId'", id.AlertRecordId, "alertRecordIdValue")
	}
}

func TestFormatDeviceManagementMonitoringAlertRecordID(t *testing.T) {
	actual := NewDeviceManagementMonitoringAlertRecordID("alertRecordIdValue").ID()
	expected := "/deviceManagement/monitoring/alertRecords/alertRecordIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementMonitoringAlertRecordID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMonitoringAlertRecordId
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
			Input: "/deviceManagement/monitoring",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/monitoring/alertRecords",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/monitoring/alertRecords/alertRecordIdValue",
			Expected: &DeviceManagementMonitoringAlertRecordId{
				AlertRecordId: "alertRecordIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/monitoring/alertRecords/alertRecordIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMonitoringAlertRecordID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AlertRecordId != v.Expected.AlertRecordId {
			t.Fatalf("Expected %q but got %q for AlertRecordId", v.Expected.AlertRecordId, actual.AlertRecordId)
		}

	}
}

func TestParseDeviceManagementMonitoringAlertRecordIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMonitoringAlertRecordId
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
			Input: "/deviceManagement/monitoring",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mOnItOrInG",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/monitoring/alertRecords",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mOnItOrInG/aLeRtReCoRdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/monitoring/alertRecords/alertRecordIdValue",
			Expected: &DeviceManagementMonitoringAlertRecordId{
				AlertRecordId: "alertRecordIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/monitoring/alertRecords/alertRecordIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mOnItOrInG/aLeRtReCoRdS/aLeRtReCoRdIdVaLuE",
			Expected: &DeviceManagementMonitoringAlertRecordId{
				AlertRecordId: "aLeRtReCoRdIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mOnItOrInG/aLeRtReCoRdS/aLeRtReCoRdIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMonitoringAlertRecordIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AlertRecordId != v.Expected.AlertRecordId {
			t.Fatalf("Expected %q but got %q for AlertRecordId", v.Expected.AlertRecordId, actual.AlertRecordId)
		}

	}
}

func TestSegmentsForDeviceManagementMonitoringAlertRecordId(t *testing.T) {
	segments := DeviceManagementMonitoringAlertRecordId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementMonitoringAlertRecordId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
