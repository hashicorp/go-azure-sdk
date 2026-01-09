package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailyActiveUserId{}

func TestNewReportUserInsightDailyActiveUserID(t *testing.T) {
	id := NewReportUserInsightDailyActiveUserID("activeUsersMetricId")

	if id.ActiveUsersMetricId != "activeUsersMetricId" {
		t.Fatalf("Expected %q but got %q for Segment 'ActiveUsersMetricId'", id.ActiveUsersMetricId, "activeUsersMetricId")
	}
}

func TestFormatReportUserInsightDailyActiveUserID(t *testing.T) {
	actual := NewReportUserInsightDailyActiveUserID("activeUsersMetricId").ID()
	expected := "/reports/userInsights/daily/activeUsers/activeUsersMetricId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightDailyActiveUserID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightDailyActiveUserId
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
			Input: "/reports/userInsights/daily/activeUsers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/daily/activeUsers/activeUsersMetricId",
			Expected: &ReportUserInsightDailyActiveUserId{
				ActiveUsersMetricId: "activeUsersMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/daily/activeUsers/activeUsersMetricId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightDailyActiveUserID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ActiveUsersMetricId != v.Expected.ActiveUsersMetricId {
			t.Fatalf("Expected %q but got %q for ActiveUsersMetricId", v.Expected.ActiveUsersMetricId, actual.ActiveUsersMetricId)
		}

	}
}

func TestParseReportUserInsightDailyActiveUserIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightDailyActiveUserId
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
			Input: "/reports/userInsights/daily/activeUsers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/aCtIvEuSeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/daily/activeUsers/activeUsersMetricId",
			Expected: &ReportUserInsightDailyActiveUserId{
				ActiveUsersMetricId: "activeUsersMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/daily/activeUsers/activeUsersMetricId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/aCtIvEuSeRs/aCtIvEuSeRsMeTrIcId",
			Expected: &ReportUserInsightDailyActiveUserId{
				ActiveUsersMetricId: "aCtIvEuSeRsMeTrIcId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/aCtIvEuSeRs/aCtIvEuSeRsMeTrIcId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightDailyActiveUserIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ActiveUsersMetricId != v.Expected.ActiveUsersMetricId {
			t.Fatalf("Expected %q but got %q for ActiveUsersMetricId", v.Expected.ActiveUsersMetricId, actual.ActiveUsersMetricId)
		}

	}
}

func TestSegmentsForReportUserInsightDailyActiveUserId(t *testing.T) {
	segments := ReportUserInsightDailyActiveUserId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightDailyActiveUserId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
