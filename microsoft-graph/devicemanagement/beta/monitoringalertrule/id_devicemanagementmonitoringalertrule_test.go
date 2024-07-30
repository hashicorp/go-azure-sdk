package monitoringalertrule

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMonitoringAlertRuleId{}

func TestNewDeviceManagementMonitoringAlertRuleID(t *testing.T) {
	id := NewDeviceManagementMonitoringAlertRuleID("alertRuleIdValue")

	if id.AlertRuleId != "alertRuleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AlertRuleId'", id.AlertRuleId, "alertRuleIdValue")
	}
}

func TestFormatDeviceManagementMonitoringAlertRuleID(t *testing.T) {
	actual := NewDeviceManagementMonitoringAlertRuleID("alertRuleIdValue").ID()
	expected := "/deviceManagement/monitoring/alertRules/alertRuleIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementMonitoringAlertRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMonitoringAlertRuleId
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
			Input: "/deviceManagement/monitoring/alertRules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/monitoring/alertRules/alertRuleIdValue",
			Expected: &DeviceManagementMonitoringAlertRuleId{
				AlertRuleId: "alertRuleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/monitoring/alertRules/alertRuleIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMonitoringAlertRuleID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AlertRuleId != v.Expected.AlertRuleId {
			t.Fatalf("Expected %q but got %q for AlertRuleId", v.Expected.AlertRuleId, actual.AlertRuleId)
		}

	}
}

func TestParseDeviceManagementMonitoringAlertRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMonitoringAlertRuleId
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
			Input: "/deviceManagement/monitoring/alertRules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mOnItOrInG/aLeRtRuLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/monitoring/alertRules/alertRuleIdValue",
			Expected: &DeviceManagementMonitoringAlertRuleId{
				AlertRuleId: "alertRuleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/monitoring/alertRules/alertRuleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mOnItOrInG/aLeRtRuLeS/aLeRtRuLeIdVaLuE",
			Expected: &DeviceManagementMonitoringAlertRuleId{
				AlertRuleId: "aLeRtRuLeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mOnItOrInG/aLeRtRuLeS/aLeRtRuLeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMonitoringAlertRuleIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AlertRuleId != v.Expected.AlertRuleId {
			t.Fatalf("Expected %q but got %q for AlertRuleId", v.Expected.AlertRuleId, actual.AlertRuleId)
		}

	}
}

func TestSegmentsForDeviceManagementMonitoringAlertRuleId(t *testing.T) {
	segments := DeviceManagementMonitoringAlertRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementMonitoringAlertRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
