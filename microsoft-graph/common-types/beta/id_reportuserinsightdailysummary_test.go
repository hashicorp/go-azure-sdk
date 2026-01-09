package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailySummaryId{}

func TestNewReportUserInsightDailySummaryID(t *testing.T) {
	id := NewReportUserInsightDailySummaryID("insightSummaryId")

	if id.InsightSummaryId != "insightSummaryId" {
		t.Fatalf("Expected %q but got %q for Segment 'InsightSummaryId'", id.InsightSummaryId, "insightSummaryId")
	}
}

func TestFormatReportUserInsightDailySummaryID(t *testing.T) {
	actual := NewReportUserInsightDailySummaryID("insightSummaryId").ID()
	expected := "/reports/userInsights/daily/summary/insightSummaryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightDailySummaryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightDailySummaryId
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
			Input: "/reports/userInsights/daily/summary",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/daily/summary/insightSummaryId",
			Expected: &ReportUserInsightDailySummaryId{
				InsightSummaryId: "insightSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/daily/summary/insightSummaryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightDailySummaryID(v.Input)
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

func TestParseReportUserInsightDailySummaryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightDailySummaryId
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
			Input: "/reports/userInsights/daily/summary",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/sUmMaRy",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/daily/summary/insightSummaryId",
			Expected: &ReportUserInsightDailySummaryId{
				InsightSummaryId: "insightSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/daily/summary/insightSummaryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/sUmMaRy/iNsIgHtSuMmArYiD",
			Expected: &ReportUserInsightDailySummaryId{
				InsightSummaryId: "iNsIgHtSuMmArYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/sUmMaRy/iNsIgHtSuMmArYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightDailySummaryIDInsensitively(v.Input)
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

func TestSegmentsForReportUserInsightDailySummaryId(t *testing.T) {
	segments := ReportUserInsightDailySummaryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightDailySummaryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
