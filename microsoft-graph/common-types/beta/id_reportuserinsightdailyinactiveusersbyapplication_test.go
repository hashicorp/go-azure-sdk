package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailyInactiveUsersByApplicationId{}

func TestNewReportUserInsightDailyInactiveUsersByApplicationID(t *testing.T) {
	id := NewReportUserInsightDailyInactiveUsersByApplicationID("dailyInactiveUsersByApplicationMetricIdValue")

	if id.DailyInactiveUsersByApplicationMetricId != "dailyInactiveUsersByApplicationMetricIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DailyInactiveUsersByApplicationMetricId'", id.DailyInactiveUsersByApplicationMetricId, "dailyInactiveUsersByApplicationMetricIdValue")
	}
}

func TestFormatReportUserInsightDailyInactiveUsersByApplicationID(t *testing.T) {
	actual := NewReportUserInsightDailyInactiveUsersByApplicationID("dailyInactiveUsersByApplicationMetricIdValue").ID()
	expected := "/reports/userInsights/daily/inactiveUsersByApplication/dailyInactiveUsersByApplicationMetricIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightDailyInactiveUsersByApplicationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightDailyInactiveUsersByApplicationId
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
			Input: "/reports/userInsights",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/userInsights/daily",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/userInsights/daily/inactiveUsersByApplication",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/daily/inactiveUsersByApplication/dailyInactiveUsersByApplicationMetricIdValue",
			Expected: &ReportUserInsightDailyInactiveUsersByApplicationId{
				DailyInactiveUsersByApplicationMetricId: "dailyInactiveUsersByApplicationMetricIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/daily/inactiveUsersByApplication/dailyInactiveUsersByApplicationMetricIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightDailyInactiveUsersByApplicationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DailyInactiveUsersByApplicationMetricId != v.Expected.DailyInactiveUsersByApplicationMetricId {
			t.Fatalf("Expected %q but got %q for DailyInactiveUsersByApplicationMetricId", v.Expected.DailyInactiveUsersByApplicationMetricId, actual.DailyInactiveUsersByApplicationMetricId)
		}

	}
}

func TestParseReportUserInsightDailyInactiveUsersByApplicationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightDailyInactiveUsersByApplicationId
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
			Input: "/reports/userInsights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/userInsights/daily",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/userInsights/daily/inactiveUsersByApplication",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/iNaCtIvEuSeRsByApPlIcAtIoN",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/daily/inactiveUsersByApplication/dailyInactiveUsersByApplicationMetricIdValue",
			Expected: &ReportUserInsightDailyInactiveUsersByApplicationId{
				DailyInactiveUsersByApplicationMetricId: "dailyInactiveUsersByApplicationMetricIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/daily/inactiveUsersByApplication/dailyInactiveUsersByApplicationMetricIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/iNaCtIvEuSeRsByApPlIcAtIoN/dAiLyInAcTiVeUsErSbYaPpLiCaTiOnMeTrIcIdVaLuE",
			Expected: &ReportUserInsightDailyInactiveUsersByApplicationId{
				DailyInactiveUsersByApplicationMetricId: "dAiLyInAcTiVeUsErSbYaPpLiCaTiOnMeTrIcIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/iNaCtIvEuSeRsByApPlIcAtIoN/dAiLyInAcTiVeUsErSbYaPpLiCaTiOnMeTrIcIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightDailyInactiveUsersByApplicationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DailyInactiveUsersByApplicationMetricId != v.Expected.DailyInactiveUsersByApplicationMetricId {
			t.Fatalf("Expected %q but got %q for DailyInactiveUsersByApplicationMetricId", v.Expected.DailyInactiveUsersByApplicationMetricId, actual.DailyInactiveUsersByApplicationMetricId)
		}

	}
}

func TestSegmentsForReportUserInsightDailyInactiveUsersByApplicationId(t *testing.T) {
	segments := ReportUserInsightDailyInactiveUsersByApplicationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightDailyInactiveUsersByApplicationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
