package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailyMfaTelecomFraudId{}

func TestNewReportUserInsightDailyMfaTelecomFraudID(t *testing.T) {
	id := NewReportUserInsightDailyMfaTelecomFraudID("mfaTelecomFraudMetricId")

	if id.MfaTelecomFraudMetricId != "mfaTelecomFraudMetricId" {
		t.Fatalf("Expected %q but got %q for Segment 'MfaTelecomFraudMetricId'", id.MfaTelecomFraudMetricId, "mfaTelecomFraudMetricId")
	}
}

func TestFormatReportUserInsightDailyMfaTelecomFraudID(t *testing.T) {
	actual := NewReportUserInsightDailyMfaTelecomFraudID("mfaTelecomFraudMetricId").ID()
	expected := "/reports/userInsights/daily/mfaTelecomFraud/mfaTelecomFraudMetricId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightDailyMfaTelecomFraudID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightDailyMfaTelecomFraudId
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
			Input: "/reports/userInsights/daily/mfaTelecomFraud",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/daily/mfaTelecomFraud/mfaTelecomFraudMetricId",
			Expected: &ReportUserInsightDailyMfaTelecomFraudId{
				MfaTelecomFraudMetricId: "mfaTelecomFraudMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/daily/mfaTelecomFraud/mfaTelecomFraudMetricId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightDailyMfaTelecomFraudID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MfaTelecomFraudMetricId != v.Expected.MfaTelecomFraudMetricId {
			t.Fatalf("Expected %q but got %q for MfaTelecomFraudMetricId", v.Expected.MfaTelecomFraudMetricId, actual.MfaTelecomFraudMetricId)
		}

	}
}

func TestParseReportUserInsightDailyMfaTelecomFraudIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightDailyMfaTelecomFraudId
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
			Input: "/reports/userInsights/daily/mfaTelecomFraud",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/mFaTeLeCoMfRaUd",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/daily/mfaTelecomFraud/mfaTelecomFraudMetricId",
			Expected: &ReportUserInsightDailyMfaTelecomFraudId{
				MfaTelecomFraudMetricId: "mfaTelecomFraudMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/daily/mfaTelecomFraud/mfaTelecomFraudMetricId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/mFaTeLeCoMfRaUd/mFaTeLeCoMfRaUdMeTrIcId",
			Expected: &ReportUserInsightDailyMfaTelecomFraudId{
				MfaTelecomFraudMetricId: "mFaTeLeCoMfRaUdMeTrIcId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/mFaTeLeCoMfRaUd/mFaTeLeCoMfRaUdMeTrIcId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightDailyMfaTelecomFraudIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MfaTelecomFraudMetricId != v.Expected.MfaTelecomFraudMetricId {
			t.Fatalf("Expected %q but got %q for MfaTelecomFraudMetricId", v.Expected.MfaTelecomFraudMetricId, actual.MfaTelecomFraudMetricId)
		}

	}
}

func TestSegmentsForReportUserInsightDailyMfaTelecomFraudId(t *testing.T) {
	segments := ReportUserInsightDailyMfaTelecomFraudId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightDailyMfaTelecomFraudId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
