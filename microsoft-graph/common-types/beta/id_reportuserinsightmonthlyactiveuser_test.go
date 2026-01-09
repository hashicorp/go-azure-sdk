package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlyActiveUserId{}

func TestNewReportUserInsightMonthlyActiveUserID(t *testing.T) {
	id := NewReportUserInsightMonthlyActiveUserID("activeUsersMetricId")

	if id.ActiveUsersMetricId != "activeUsersMetricId" {
		t.Fatalf("Expected %q but got %q for Segment 'ActiveUsersMetricId'", id.ActiveUsersMetricId, "activeUsersMetricId")
	}
}

func TestFormatReportUserInsightMonthlyActiveUserID(t *testing.T) {
	actual := NewReportUserInsightMonthlyActiveUserID("activeUsersMetricId").ID()
	expected := "/reports/userInsights/monthly/activeUsers/activeUsersMetricId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightMonthlyActiveUserID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlyActiveUserId
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
			Input: "/reports/userInsights/monthly/activeUsers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/activeUsers/activeUsersMetricId",
			Expected: &ReportUserInsightMonthlyActiveUserId{
				ActiveUsersMetricId: "activeUsersMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/activeUsers/activeUsersMetricId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlyActiveUserID(v.Input)
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

func TestParseReportUserInsightMonthlyActiveUserIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlyActiveUserId
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
			Input: "/reports/userInsights/monthly/activeUsers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/aCtIvEuSeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/activeUsers/activeUsersMetricId",
			Expected: &ReportUserInsightMonthlyActiveUserId{
				ActiveUsersMetricId: "activeUsersMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/activeUsers/activeUsersMetricId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/aCtIvEuSeRs/aCtIvEuSeRsMeTrIcId",
			Expected: &ReportUserInsightMonthlyActiveUserId{
				ActiveUsersMetricId: "aCtIvEuSeRsMeTrIcId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/aCtIvEuSeRs/aCtIvEuSeRsMeTrIcId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlyActiveUserIDInsensitively(v.Input)
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

func TestSegmentsForReportUserInsightMonthlyActiveUserId(t *testing.T) {
	segments := ReportUserInsightMonthlyActiveUserId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightMonthlyActiveUserId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
