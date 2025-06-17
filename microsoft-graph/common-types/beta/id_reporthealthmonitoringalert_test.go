package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportHealthMonitoringAlertId{}

func TestNewReportHealthMonitoringAlertID(t *testing.T) {
	id := NewReportHealthMonitoringAlertID("alertId")

	if id.AlertId != "alertId" {
		t.Fatalf("Expected %q but got %q for Segment 'AlertId'", id.AlertId, "alertId")
	}
}

func TestFormatReportHealthMonitoringAlertID(t *testing.T) {
	actual := NewReportHealthMonitoringAlertID("alertId").ID()
	expected := "/reports/healthMonitoring/alerts/alertId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportHealthMonitoringAlertID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportHealthMonitoringAlertId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/healthMonitoring",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/healthMonitoring/alerts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/healthMonitoring/alerts/alertId",
			Expected: &ReportHealthMonitoringAlertId{
				AlertId: "alertId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/healthMonitoring/alerts/alertId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportHealthMonitoringAlertID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AlertId != v.Expected.AlertId {
			t.Fatalf("Expected %q but got %q for AlertId", v.Expected.AlertId, actual.AlertId)
		}

	}
}

func TestParseReportHealthMonitoringAlertIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportHealthMonitoringAlertId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/healthMonitoring",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/hEaLtHmOnItOrInG",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/healthMonitoring/alerts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/hEaLtHmOnItOrInG/aLeRtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/healthMonitoring/alerts/alertId",
			Expected: &ReportHealthMonitoringAlertId{
				AlertId: "alertId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/healthMonitoring/alerts/alertId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/hEaLtHmOnItOrInG/aLeRtS/aLeRtId",
			Expected: &ReportHealthMonitoringAlertId{
				AlertId: "aLeRtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/hEaLtHmOnItOrInG/aLeRtS/aLeRtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportHealthMonitoringAlertIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AlertId != v.Expected.AlertId {
			t.Fatalf("Expected %q but got %q for AlertId", v.Expected.AlertId, actual.AlertId)
		}

	}
}

func TestSegmentsForReportHealthMonitoringAlertId(t *testing.T) {
	segments := ReportHealthMonitoringAlertId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportHealthMonitoringAlertId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
