package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightDailySignUpId{}

func TestNewReportUserInsightDailySignUpID(t *testing.T) {
	id := NewReportUserInsightDailySignUpID("userSignUpMetricIdValue")

	if id.UserSignUpMetricId != "userSignUpMetricIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserSignUpMetricId'", id.UserSignUpMetricId, "userSignUpMetricIdValue")
	}
}

func TestFormatReportUserInsightDailySignUpID(t *testing.T) {
	actual := NewReportUserInsightDailySignUpID("userSignUpMetricIdValue").ID()
	expected := "/reports/userInsights/daily/signUps/userSignUpMetricIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightDailySignUpID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightDailySignUpId
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
			Input: "/reports/userInsights/daily/signUps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/daily/signUps/userSignUpMetricIdValue",
			Expected: &ReportUserInsightDailySignUpId{
				UserSignUpMetricId: "userSignUpMetricIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/daily/signUps/userSignUpMetricIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightDailySignUpID(v.Input)
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

func TestParseReportUserInsightDailySignUpIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightDailySignUpId
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
			Input: "/reports/userInsights/daily/signUps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/sIgNuPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/daily/signUps/userSignUpMetricIdValue",
			Expected: &ReportUserInsightDailySignUpId{
				UserSignUpMetricId: "userSignUpMetricIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/daily/signUps/userSignUpMetricIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/sIgNuPs/uSeRsIgNuPmEtRiCiDvAlUe",
			Expected: &ReportUserInsightDailySignUpId{
				UserSignUpMetricId: "uSeRsIgNuPmEtRiCiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/dAiLy/sIgNuPs/uSeRsIgNuPmEtRiCiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightDailySignUpIDInsensitively(v.Input)
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

func TestSegmentsForReportUserInsightDailySignUpId(t *testing.T) {
	segments := ReportUserInsightDailySignUpId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightDailySignUpId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
