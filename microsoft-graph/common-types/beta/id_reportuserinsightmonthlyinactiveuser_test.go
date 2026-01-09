package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlyInactiveUserId{}

func TestNewReportUserInsightMonthlyInactiveUserID(t *testing.T) {
	id := NewReportUserInsightMonthlyInactiveUserID("monthlyInactiveUsersMetricId")

	if id.MonthlyInactiveUsersMetricId != "monthlyInactiveUsersMetricId" {
		t.Fatalf("Expected %q but got %q for Segment 'MonthlyInactiveUsersMetricId'", id.MonthlyInactiveUsersMetricId, "monthlyInactiveUsersMetricId")
	}
}

func TestFormatReportUserInsightMonthlyInactiveUserID(t *testing.T) {
	actual := NewReportUserInsightMonthlyInactiveUserID("monthlyInactiveUsersMetricId").ID()
	expected := "/reports/userInsights/monthly/inactiveUsers/monthlyInactiveUsersMetricId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightMonthlyInactiveUserID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlyInactiveUserId
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
			Input: "/reports/userInsights/monthly/inactiveUsers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/inactiveUsers/monthlyInactiveUsersMetricId",
			Expected: &ReportUserInsightMonthlyInactiveUserId{
				MonthlyInactiveUsersMetricId: "monthlyInactiveUsersMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/inactiveUsers/monthlyInactiveUsersMetricId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlyInactiveUserID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MonthlyInactiveUsersMetricId != v.Expected.MonthlyInactiveUsersMetricId {
			t.Fatalf("Expected %q but got %q for MonthlyInactiveUsersMetricId", v.Expected.MonthlyInactiveUsersMetricId, actual.MonthlyInactiveUsersMetricId)
		}

	}
}

func TestParseReportUserInsightMonthlyInactiveUserIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlyInactiveUserId
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
			Input: "/reports/userInsights/monthly/inactiveUsers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/iNaCtIvEuSeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/inactiveUsers/monthlyInactiveUsersMetricId",
			Expected: &ReportUserInsightMonthlyInactiveUserId{
				MonthlyInactiveUsersMetricId: "monthlyInactiveUsersMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/inactiveUsers/monthlyInactiveUsersMetricId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/iNaCtIvEuSeRs/mOnThLyInAcTiVeUsErSmEtRiCiD",
			Expected: &ReportUserInsightMonthlyInactiveUserId{
				MonthlyInactiveUsersMetricId: "mOnThLyInAcTiVeUsErSmEtRiCiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/iNaCtIvEuSeRs/mOnThLyInAcTiVeUsErSmEtRiCiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlyInactiveUserIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MonthlyInactiveUsersMetricId != v.Expected.MonthlyInactiveUsersMetricId {
			t.Fatalf("Expected %q but got %q for MonthlyInactiveUsersMetricId", v.Expected.MonthlyInactiveUsersMetricId, actual.MonthlyInactiveUsersMetricId)
		}

	}
}

func TestSegmentsForReportUserInsightMonthlyInactiveUserId(t *testing.T) {
	segments := ReportUserInsightMonthlyInactiveUserId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightMonthlyInactiveUserId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
