package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlyMfaCompletionId{}

func TestNewReportUserInsightMonthlyMfaCompletionID(t *testing.T) {
	id := NewReportUserInsightMonthlyMfaCompletionID("mfaCompletionMetricId")

	if id.MfaCompletionMetricId != "mfaCompletionMetricId" {
		t.Fatalf("Expected %q but got %q for Segment 'MfaCompletionMetricId'", id.MfaCompletionMetricId, "mfaCompletionMetricId")
	}
}

func TestFormatReportUserInsightMonthlyMfaCompletionID(t *testing.T) {
	actual := NewReportUserInsightMonthlyMfaCompletionID("mfaCompletionMetricId").ID()
	expected := "/reports/userInsights/monthly/mfaCompletions/mfaCompletionMetricId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightMonthlyMfaCompletionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlyMfaCompletionId
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
			Input: "/reports/userInsights/monthly/mfaCompletions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/mfaCompletions/mfaCompletionMetricId",
			Expected: &ReportUserInsightMonthlyMfaCompletionId{
				MfaCompletionMetricId: "mfaCompletionMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/mfaCompletions/mfaCompletionMetricId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlyMfaCompletionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MfaCompletionMetricId != v.Expected.MfaCompletionMetricId {
			t.Fatalf("Expected %q but got %q for MfaCompletionMetricId", v.Expected.MfaCompletionMetricId, actual.MfaCompletionMetricId)
		}

	}
}

func TestParseReportUserInsightMonthlyMfaCompletionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlyMfaCompletionId
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
			Input: "/reports/userInsights/monthly/mfaCompletions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/mFaCoMpLeTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/mfaCompletions/mfaCompletionMetricId",
			Expected: &ReportUserInsightMonthlyMfaCompletionId{
				MfaCompletionMetricId: "mfaCompletionMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/mfaCompletions/mfaCompletionMetricId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/mFaCoMpLeTiOnS/mFaCoMpLeTiOnMeTrIcId",
			Expected: &ReportUserInsightMonthlyMfaCompletionId{
				MfaCompletionMetricId: "mFaCoMpLeTiOnMeTrIcId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/mFaCoMpLeTiOnS/mFaCoMpLeTiOnMeTrIcId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlyMfaCompletionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MfaCompletionMetricId != v.Expected.MfaCompletionMetricId {
			t.Fatalf("Expected %q but got %q for MfaCompletionMetricId", v.Expected.MfaCompletionMetricId, actual.MfaCompletionMetricId)
		}

	}
}

func TestSegmentsForReportUserInsightMonthlyMfaCompletionId(t *testing.T) {
	segments := ReportUserInsightMonthlyMfaCompletionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightMonthlyMfaCompletionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
