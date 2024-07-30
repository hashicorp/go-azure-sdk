package userinsightdailymfacompletion

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailyMfaCompletionId{}

func TestNewReportUserInsightDailyMfaCompletionID(t *testing.T) {
	id := NewReportUserInsightDailyMfaCompletionID("mfaCompletionMetricIdValue")

	if id.MfaCompletionMetricId != "mfaCompletionMetricIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MfaCompletionMetricId'", id.MfaCompletionMetricId, "mfaCompletionMetricIdValue")
	}
}

func TestFormatReportUserInsightDailyMfaCompletionID(t *testing.T) {
	actual := NewReportUserInsightDailyMfaCompletionID("mfaCompletionMetricIdValue").ID()
	expected := "/reports/userInsights/daily/mfaCompletions/mfaCompletionMetricIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightDailyMfaCompletionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightDailyMfaCompletionId
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
			Input: "/reports/userInsights/daily/mfaCompletions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/daily/mfaCompletions/mfaCompletionMetricIdValue",
			Expected: &ReportUserInsightDailyMfaCompletionId{
				MfaCompletionMetricId: "mfaCompletionMetricIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/daily/mfaCompletions/mfaCompletionMetricIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightDailyMfaCompletionID(v.Input)
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

func TestParseReportUserInsightDailyMfaCompletionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightDailyMfaCompletionId
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
			Input: "/reports/userInsights/daily/mfaCompletions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/mFaCoMpLeTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/daily/mfaCompletions/mfaCompletionMetricIdValue",
			Expected: &ReportUserInsightDailyMfaCompletionId{
				MfaCompletionMetricId: "mfaCompletionMetricIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/daily/mfaCompletions/mfaCompletionMetricIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/mFaCoMpLeTiOnS/mFaCoMpLeTiOnMeTrIcIdVaLuE",
			Expected: &ReportUserInsightDailyMfaCompletionId{
				MfaCompletionMetricId: "mFaCoMpLeTiOnMeTrIcIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/mFaCoMpLeTiOnS/mFaCoMpLeTiOnMeTrIcIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightDailyMfaCompletionIDInsensitively(v.Input)
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

func TestSegmentsForReportUserInsightDailyMfaCompletionId(t *testing.T) {
	segments := ReportUserInsightDailyMfaCompletionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightDailyMfaCompletionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
