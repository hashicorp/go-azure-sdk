package userinsightmonthlysignup

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlySignUpId{}

func TestNewReportUserInsightMonthlySignUpID(t *testing.T) {
	id := NewReportUserInsightMonthlySignUpID("userSignUpMetricIdValue")

	if id.UserSignUpMetricId != "userSignUpMetricIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserSignUpMetricId'", id.UserSignUpMetricId, "userSignUpMetricIdValue")
	}
}

func TestFormatReportUserInsightMonthlySignUpID(t *testing.T) {
	actual := NewReportUserInsightMonthlySignUpID("userSignUpMetricIdValue").ID()
	expected := "/reports/userInsights/monthly/signUps/userSignUpMetricIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightMonthlySignUpID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlySignUpId
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
			Input: "/reports/userInsights/monthly/signUps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/signUps/userSignUpMetricIdValue",
			Expected: &ReportUserInsightMonthlySignUpId{
				UserSignUpMetricId: "userSignUpMetricIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/signUps/userSignUpMetricIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlySignUpID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserSignUpMetricId != v.Expected.UserSignUpMetricId {
			t.Fatalf("Expected %q but got %q for UserSignUpMetricId", v.Expected.UserSignUpMetricId, actual.UserSignUpMetricId)
		}

	}
}

func TestParseReportUserInsightMonthlySignUpIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlySignUpId
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
			Input: "/reports/userInsights/monthly/signUps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/sIgNuPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/signUps/userSignUpMetricIdValue",
			Expected: &ReportUserInsightMonthlySignUpId{
				UserSignUpMetricId: "userSignUpMetricIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/signUps/userSignUpMetricIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/sIgNuPs/uSeRsIgNuPmEtRiCiDvAlUe",
			Expected: &ReportUserInsightMonthlySignUpId{
				UserSignUpMetricId: "uSeRsIgNuPmEtRiCiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/sIgNuPs/uSeRsIgNuPmEtRiCiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlySignUpIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserSignUpMetricId != v.Expected.UserSignUpMetricId {
			t.Fatalf("Expected %q but got %q for UserSignUpMetricId", v.Expected.UserSignUpMetricId, actual.UserSignUpMetricId)
		}

	}
}

func TestSegmentsForReportUserInsightMonthlySignUpId(t *testing.T) {
	segments := ReportUserInsightMonthlySignUpId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightMonthlySignUpId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
