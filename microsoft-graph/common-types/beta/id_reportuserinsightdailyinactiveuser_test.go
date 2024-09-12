package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailyInactiveUserId{}

func TestNewReportUserInsightDailyInactiveUserID(t *testing.T) {
	id := NewReportUserInsightDailyInactiveUserID("dailyInactiveUsersMetricIdValue")

	if id.DailyInactiveUsersMetricId != "dailyInactiveUsersMetricIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DailyInactiveUsersMetricId'", id.DailyInactiveUsersMetricId, "dailyInactiveUsersMetricIdValue")
	}
}

func TestFormatReportUserInsightDailyInactiveUserID(t *testing.T) {
	actual := NewReportUserInsightDailyInactiveUserID("dailyInactiveUsersMetricIdValue").ID()
	expected := "/reports/userInsights/daily/inactiveUsers/dailyInactiveUsersMetricIdValue"
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
			Input: "/reports/userInsights/daily/inactiveUsers/dailyInactiveUsersMetricIdValue",
			Expected: &ReportUserInsightDailyInactiveUserId{
				DailyInactiveUsersMetricId: "dailyInactiveUsersMetricIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/daily/inactiveUsers/dailyInactiveUsersMetricIdValue/extra",
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
			Input: "/reports/userInsights/daily/inactiveUsers/dailyInactiveUsersMetricIdValue",
			Expected: &ReportUserInsightDailyInactiveUserId{
				DailyInactiveUsersMetricId: "dailyInactiveUsersMetricIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/daily/inactiveUsers/dailyInactiveUsersMetricIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/iNaCtIvEuSeRs/dAiLyInAcTiVeUsErSmEtRiCiDvAlUe",
			Expected: &ReportUserInsightDailyInactiveUserId{
				DailyInactiveUsersMetricId: "dAiLyInAcTiVeUsErSmEtRiCiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/iNaCtIvEuSeRs/dAiLyInAcTiVeUsErSmEtRiCiDvAlUe/extra",
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
