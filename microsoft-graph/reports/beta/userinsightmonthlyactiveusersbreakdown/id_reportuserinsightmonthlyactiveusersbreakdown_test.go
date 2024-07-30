package userinsightmonthlyactiveusersbreakdown

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlyActiveUsersBreakdownId{}

func TestNewReportUserInsightMonthlyActiveUsersBreakdownID(t *testing.T) {
	id := NewReportUserInsightMonthlyActiveUsersBreakdownID("activeUsersBreakdownMetricIdValue")

	if id.ActiveUsersBreakdownMetricId != "activeUsersBreakdownMetricIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ActiveUsersBreakdownMetricId'", id.ActiveUsersBreakdownMetricId, "activeUsersBreakdownMetricIdValue")
	}
}

func TestFormatReportUserInsightMonthlyActiveUsersBreakdownID(t *testing.T) {
	actual := NewReportUserInsightMonthlyActiveUsersBreakdownID("activeUsersBreakdownMetricIdValue").ID()
	expected := "/reports/userInsights/monthly/activeUsersBreakdown/activeUsersBreakdownMetricIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightMonthlyActiveUsersBreakdownID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlyActiveUsersBreakdownId
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
			Input: "/reports/userInsights/monthly/activeUsersBreakdown",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/activeUsersBreakdown/activeUsersBreakdownMetricIdValue",
			Expected: &ReportUserInsightMonthlyActiveUsersBreakdownId{
				ActiveUsersBreakdownMetricId: "activeUsersBreakdownMetricIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/activeUsersBreakdown/activeUsersBreakdownMetricIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlyActiveUsersBreakdownID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ActiveUsersBreakdownMetricId != v.Expected.ActiveUsersBreakdownMetricId {
			t.Fatalf("Expected %q but got %q for ActiveUsersBreakdownMetricId", v.Expected.ActiveUsersBreakdownMetricId, actual.ActiveUsersBreakdownMetricId)
		}

	}
}

func TestParseReportUserInsightMonthlyActiveUsersBreakdownIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlyActiveUsersBreakdownId
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
			Input: "/reports/userInsights/monthly/activeUsersBreakdown",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/aCtIvEuSeRsBrEaKdOwN",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/activeUsersBreakdown/activeUsersBreakdownMetricIdValue",
			Expected: &ReportUserInsightMonthlyActiveUsersBreakdownId{
				ActiveUsersBreakdownMetricId: "activeUsersBreakdownMetricIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/activeUsersBreakdown/activeUsersBreakdownMetricIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/aCtIvEuSeRsBrEaKdOwN/aCtIvEuSeRsBrEaKdOwNmEtRiCiDvAlUe",
			Expected: &ReportUserInsightMonthlyActiveUsersBreakdownId{
				ActiveUsersBreakdownMetricId: "aCtIvEuSeRsBrEaKdOwNmEtRiCiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/aCtIvEuSeRsBrEaKdOwN/aCtIvEuSeRsBrEaKdOwNmEtRiCiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlyActiveUsersBreakdownIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ActiveUsersBreakdownMetricId != v.Expected.ActiveUsersBreakdownMetricId {
			t.Fatalf("Expected %q but got %q for ActiveUsersBreakdownMetricId", v.Expected.ActiveUsersBreakdownMetricId, actual.ActiveUsersBreakdownMetricId)
		}

	}
}

func TestSegmentsForReportUserInsightMonthlyActiveUsersBreakdownId(t *testing.T) {
	segments := ReportUserInsightMonthlyActiveUsersBreakdownId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightMonthlyActiveUsersBreakdownId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
