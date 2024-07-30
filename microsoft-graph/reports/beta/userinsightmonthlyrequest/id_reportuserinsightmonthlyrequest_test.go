package userinsightmonthlyrequest

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlyRequestId{}

func TestNewReportUserInsightMonthlyRequestID(t *testing.T) {
	id := NewReportUserInsightMonthlyRequestID("userRequestsMetricIdValue")

	if id.UserRequestsMetricId != "userRequestsMetricIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserRequestsMetricId'", id.UserRequestsMetricId, "userRequestsMetricIdValue")
	}
}

func TestFormatReportUserInsightMonthlyRequestID(t *testing.T) {
	actual := NewReportUserInsightMonthlyRequestID("userRequestsMetricIdValue").ID()
	expected := "/reports/userInsights/monthly/requests/userRequestsMetricIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightMonthlyRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlyRequestId
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
			Input: "/reports/userInsights/monthly/requests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/requests/userRequestsMetricIdValue",
			Expected: &ReportUserInsightMonthlyRequestId{
				UserRequestsMetricId: "userRequestsMetricIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/requests/userRequestsMetricIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlyRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserRequestsMetricId != v.Expected.UserRequestsMetricId {
			t.Fatalf("Expected %q but got %q for UserRequestsMetricId", v.Expected.UserRequestsMetricId, actual.UserRequestsMetricId)
		}

	}
}

func TestParseReportUserInsightMonthlyRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlyRequestId
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
			Input: "/reports/userInsights/monthly/requests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/rEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/requests/userRequestsMetricIdValue",
			Expected: &ReportUserInsightMonthlyRequestId{
				UserRequestsMetricId: "userRequestsMetricIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/requests/userRequestsMetricIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/rEqUeStS/uSeRrEqUeStSmEtRiCiDvAlUe",
			Expected: &ReportUserInsightMonthlyRequestId{
				UserRequestsMetricId: "uSeRrEqUeStSmEtRiCiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/rEqUeStS/uSeRrEqUeStSmEtRiCiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlyRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserRequestsMetricId != v.Expected.UserRequestsMetricId {
			t.Fatalf("Expected %q but got %q for UserRequestsMetricId", v.Expected.UserRequestsMetricId, actual.UserRequestsMetricId)
		}

	}
}

func TestSegmentsForReportUserInsightMonthlyRequestId(t *testing.T) {
	segments := ReportUserInsightMonthlyRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightMonthlyRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
