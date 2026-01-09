package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlySummaryId{}

func TestNewReportUserInsightMonthlySummaryID(t *testing.T) {
	id := NewReportUserInsightMonthlySummaryID("insightSummaryId")

	if id.InsightSummaryId != "insightSummaryId" {
		t.Fatalf("Expected %q but got %q for Segment 'InsightSummaryId'", id.InsightSummaryId, "insightSummaryId")
	}
}

func TestFormatReportUserInsightMonthlySummaryID(t *testing.T) {
	actual := NewReportUserInsightMonthlySummaryID("insightSummaryId").ID()
	expected := "/reports/userInsights/monthly/summary/insightSummaryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightMonthlySummaryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlySummaryId
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
			Input: "/reports/userInsights/monthly/summary",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/summary/insightSummaryId",
			Expected: &ReportUserInsightMonthlySummaryId{
				InsightSummaryId: "insightSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/summary/insightSummaryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlySummaryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.InsightSummaryId != v.Expected.InsightSummaryId {
			t.Fatalf("Expected %q but got %q for InsightSummaryId", v.Expected.InsightSummaryId, actual.InsightSummaryId)
		}

	}
}

func TestParseReportUserInsightMonthlySummaryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlySummaryId
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
			Input: "/reports/userInsights/monthly/summary",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/sUmMaRy",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/summary/insightSummaryId",
			Expected: &ReportUserInsightMonthlySummaryId{
				InsightSummaryId: "insightSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/summary/insightSummaryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/sUmMaRy/iNsIgHtSuMmArYiD",
			Expected: &ReportUserInsightMonthlySummaryId{
				InsightSummaryId: "iNsIgHtSuMmArYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/sUmMaRy/iNsIgHtSuMmArYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlySummaryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.InsightSummaryId != v.Expected.InsightSummaryId {
			t.Fatalf("Expected %q but got %q for InsightSummaryId", v.Expected.InsightSummaryId, actual.InsightSummaryId)
		}

	}
}

func TestSegmentsForReportUserInsightMonthlySummaryId(t *testing.T) {
	segments := ReportUserInsightMonthlySummaryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightMonthlySummaryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
