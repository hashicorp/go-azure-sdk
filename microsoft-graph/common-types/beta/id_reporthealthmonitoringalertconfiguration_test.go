package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportHealthMonitoringAlertConfigurationId{}

func TestNewReportHealthMonitoringAlertConfigurationID(t *testing.T) {
	id := NewReportHealthMonitoringAlertConfigurationID("alertConfigurationId")

	if id.AlertConfigurationId != "alertConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'AlertConfigurationId'", id.AlertConfigurationId, "alertConfigurationId")
	}
}

func TestFormatReportHealthMonitoringAlertConfigurationID(t *testing.T) {
	actual := NewReportHealthMonitoringAlertConfigurationID("alertConfigurationId").ID()
	expected := "/reports/healthMonitoring/alertConfigurations/alertConfigurationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportHealthMonitoringAlertConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportHealthMonitoringAlertConfigurationId
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
			Input: "/reports/healthMonitoring/alertConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/healthMonitoring/alertConfigurations/alertConfigurationId",
			Expected: &ReportHealthMonitoringAlertConfigurationId{
				AlertConfigurationId: "alertConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/healthMonitoring/alertConfigurations/alertConfigurationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportHealthMonitoringAlertConfigurationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AlertConfigurationId != v.Expected.AlertConfigurationId {
			t.Fatalf("Expected %q but got %q for AlertConfigurationId", v.Expected.AlertConfigurationId, actual.AlertConfigurationId)
		}

	}
}

func TestParseReportHealthMonitoringAlertConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportHealthMonitoringAlertConfigurationId
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
			Input: "/reports/healthMonitoring/alertConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/hEaLtHmOnItOrInG/aLeRtCoNfIgUrAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/healthMonitoring/alertConfigurations/alertConfigurationId",
			Expected: &ReportHealthMonitoringAlertConfigurationId{
				AlertConfigurationId: "alertConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/healthMonitoring/alertConfigurations/alertConfigurationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/hEaLtHmOnItOrInG/aLeRtCoNfIgUrAtIoNs/aLeRtCoNfIgUrAtIoNiD",
			Expected: &ReportHealthMonitoringAlertConfigurationId{
				AlertConfigurationId: "aLeRtCoNfIgUrAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/hEaLtHmOnItOrInG/aLeRtCoNfIgUrAtIoNs/aLeRtCoNfIgUrAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportHealthMonitoringAlertConfigurationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AlertConfigurationId != v.Expected.AlertConfigurationId {
			t.Fatalf("Expected %q but got %q for AlertConfigurationId", v.Expected.AlertConfigurationId, actual.AlertConfigurationId)
		}

	}
}

func TestSegmentsForReportHealthMonitoringAlertConfigurationId(t *testing.T) {
	segments := ReportHealthMonitoringAlertConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportHealthMonitoringAlertConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
