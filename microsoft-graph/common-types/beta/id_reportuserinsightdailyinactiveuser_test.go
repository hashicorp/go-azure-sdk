package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailyInactiveUserId{}

func TestNewReportUserInsightDailyInactiveUserID(t *testing.T) {
	id := NewReportUserInsightDailyInactiveUserID("dailyInactiveUsersMetricId")

	if id.DailyInactiveUsersMetricId != "dailyInactiveUsersMetricId" {
		t.Fatalf("Expected %q but got %q for Segment 'DailyInactiveUsersMetricId'", id.DailyInactiveUsersMetricId, "dailyInactiveUsersMetricId")
	}
}

func TestFormatReportUserInsightDailyInactiveUserID(t *testing.T) {
	actual := NewReportUserInsightDailyInactiveUserID("dailyInactiveUsersMetricId").ID()
	expected := "/reports/userInsights/daily/inactiveUsers/dailyInactiveUsersMetricId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightDailyInactiveUserID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightDailyInactiveUserId
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
			Input: "/reports/userInsights/daily/inactiveUsers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/daily/inactiveUsers/dailyInactiveUsersMetricId",
			Expected: &ReportUserInsightDailyInactiveUserId{
				DailyInactiveUsersMetricId: "dailyInactiveUsersMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/daily/inactiveUsers/dailyInactiveUsersMetricId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightDailyInactiveUserID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DailyInactiveUsersMetricId != v.Expected.DailyInactiveUsersMetricId {
			t.Fatalf("Expected %q but got %q for DailyInactiveUsersMetricId", v.Expected.DailyInactiveUsersMetricId, actual.DailyInactiveUsersMetricId)
		}

	}
}

func TestParseReportUserInsightDailyInactiveUserIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightDailyInactiveUserId
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
			Input: "/reports/userInsights/daily/inactiveUsers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/iNaCtIvEuSeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/daily/inactiveUsers/dailyInactiveUsersMetricId",
			Expected: &ReportUserInsightDailyInactiveUserId{
				DailyInactiveUsersMetricId: "dailyInactiveUsersMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/daily/inactiveUsers/dailyInactiveUsersMetricId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/iNaCtIvEuSeRs/dAiLyInAcTiVeUsErSmEtRiCiD",
			Expected: &ReportUserInsightDailyInactiveUserId{
				DailyInactiveUsersMetricId: "dAiLyInAcTiVeUsErSmEtRiCiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/iNaCtIvEuSeRs/dAiLyInAcTiVeUsErSmEtRiCiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightDailyInactiveUserIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DailyInactiveUsersMetricId != v.Expected.DailyInactiveUsersMetricId {
			t.Fatalf("Expected %q but got %q for DailyInactiveUsersMetricId", v.Expected.DailyInactiveUsersMetricId, actual.DailyInactiveUsersMetricId)
		}

	}
}

func TestSegmentsForReportUserInsightDailyInactiveUserId(t *testing.T) {
	segments := ReportUserInsightDailyInactiveUserId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightDailyInactiveUserId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
