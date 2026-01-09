package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlyInactiveUsersByApplicationId{}

func TestNewReportUserInsightMonthlyInactiveUsersByApplicationID(t *testing.T) {
	id := NewReportUserInsightMonthlyInactiveUsersByApplicationID("monthlyInactiveUsersByApplicationMetricId")

	if id.MonthlyInactiveUsersByApplicationMetricId != "monthlyInactiveUsersByApplicationMetricId" {
		t.Fatalf("Expected %q but got %q for Segment 'MonthlyInactiveUsersByApplicationMetricId'", id.MonthlyInactiveUsersByApplicationMetricId, "monthlyInactiveUsersByApplicationMetricId")
	}
}

func TestFormatReportUserInsightMonthlyInactiveUsersByApplicationID(t *testing.T) {
	actual := NewReportUserInsightMonthlyInactiveUsersByApplicationID("monthlyInactiveUsersByApplicationMetricId").ID()
	expected := "/reports/userInsights/monthly/inactiveUsersByApplication/monthlyInactiveUsersByApplicationMetricId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightMonthlyInactiveUsersByApplicationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlyInactiveUsersByApplicationId
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
			Input: "/reports/userInsights/monthly",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/userInsights/monthly/inactiveUsersByApplication",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/inactiveUsersByApplication/monthlyInactiveUsersByApplicationMetricId",
			Expected: &ReportUserInsightMonthlyInactiveUsersByApplicationId{
				MonthlyInactiveUsersByApplicationMetricId: "monthlyInactiveUsersByApplicationMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/inactiveUsersByApplication/monthlyInactiveUsersByApplicationMetricId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlyInactiveUsersByApplicationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MonthlyInactiveUsersByApplicationMetricId != v.Expected.MonthlyInactiveUsersByApplicationMetricId {
			t.Fatalf("Expected %q but got %q for MonthlyInactiveUsersByApplicationMetricId", v.Expected.MonthlyInactiveUsersByApplicationMetricId, actual.MonthlyInactiveUsersByApplicationMetricId)
		}

	}
}

func TestParseReportUserInsightMonthlyInactiveUsersByApplicationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlyInactiveUsersByApplicationId
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
			Input: "/reports/userInsights/monthly",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/userInsights/monthly/inactiveUsersByApplication",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/iNaCtIvEuSeRsByApPlIcAtIoN",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/inactiveUsersByApplication/monthlyInactiveUsersByApplicationMetricId",
			Expected: &ReportUserInsightMonthlyInactiveUsersByApplicationId{
				MonthlyInactiveUsersByApplicationMetricId: "monthlyInactiveUsersByApplicationMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/inactiveUsersByApplication/monthlyInactiveUsersByApplicationMetricId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/iNaCtIvEuSeRsByApPlIcAtIoN/mOnThLyInAcTiVeUsErSbYaPpLiCaTiOnMeTrIcId",
			Expected: &ReportUserInsightMonthlyInactiveUsersByApplicationId{
				MonthlyInactiveUsersByApplicationMetricId: "mOnThLyInAcTiVeUsErSbYaPpLiCaTiOnMeTrIcId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/iNaCtIvEuSeRsByApPlIcAtIoN/mOnThLyInAcTiVeUsErSbYaPpLiCaTiOnMeTrIcId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlyInactiveUsersByApplicationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MonthlyInactiveUsersByApplicationMetricId != v.Expected.MonthlyInactiveUsersByApplicationMetricId {
			t.Fatalf("Expected %q but got %q for MonthlyInactiveUsersByApplicationMetricId", v.Expected.MonthlyInactiveUsersByApplicationMetricId, actual.MonthlyInactiveUsersByApplicationMetricId)
		}

	}
}

func TestSegmentsForReportUserInsightMonthlyInactiveUsersByApplicationId(t *testing.T) {
	segments := ReportUserInsightMonthlyInactiveUsersByApplicationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightMonthlyInactiveUsersByApplicationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
