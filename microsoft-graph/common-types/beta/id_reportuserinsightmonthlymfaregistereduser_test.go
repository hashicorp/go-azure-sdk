package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlyMfaRegisteredUserId{}

func TestNewReportUserInsightMonthlyMfaRegisteredUserID(t *testing.T) {
	id := NewReportUserInsightMonthlyMfaRegisteredUserID("mfaUserCountMetricId")

	if id.MfaUserCountMetricId != "mfaUserCountMetricId" {
		t.Fatalf("Expected %q but got %q for Segment 'MfaUserCountMetricId'", id.MfaUserCountMetricId, "mfaUserCountMetricId")
	}
}

func TestFormatReportUserInsightMonthlyMfaRegisteredUserID(t *testing.T) {
	actual := NewReportUserInsightMonthlyMfaRegisteredUserID("mfaUserCountMetricId").ID()
	expected := "/reports/userInsights/monthly/mfaRegisteredUsers/mfaUserCountMetricId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightMonthlyMfaRegisteredUserID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlyMfaRegisteredUserId
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
			Input: "/reports/userInsights/monthly/mfaRegisteredUsers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/mfaRegisteredUsers/mfaUserCountMetricId",
			Expected: &ReportUserInsightMonthlyMfaRegisteredUserId{
				MfaUserCountMetricId: "mfaUserCountMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/mfaRegisteredUsers/mfaUserCountMetricId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlyMfaRegisteredUserID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MfaUserCountMetricId != v.Expected.MfaUserCountMetricId {
			t.Fatalf("Expected %q but got %q for MfaUserCountMetricId", v.Expected.MfaUserCountMetricId, actual.MfaUserCountMetricId)
		}

	}
}

func TestParseReportUserInsightMonthlyMfaRegisteredUserIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlyMfaRegisteredUserId
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
			Input: "/reports/userInsights/monthly/mfaRegisteredUsers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/mFaReGiStErEdUsErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/mfaRegisteredUsers/mfaUserCountMetricId",
			Expected: &ReportUserInsightMonthlyMfaRegisteredUserId{
				MfaUserCountMetricId: "mfaUserCountMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/mfaRegisteredUsers/mfaUserCountMetricId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/mFaReGiStErEdUsErS/mFaUsErCoUnTmEtRiCiD",
			Expected: &ReportUserInsightMonthlyMfaRegisteredUserId{
				MfaUserCountMetricId: "mFaUsErCoUnTmEtRiCiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/mFaReGiStErEdUsErS/mFaUsErCoUnTmEtRiCiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlyMfaRegisteredUserIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MfaUserCountMetricId != v.Expected.MfaUserCountMetricId {
			t.Fatalf("Expected %q but got %q for MfaUserCountMetricId", v.Expected.MfaUserCountMetricId, actual.MfaUserCountMetricId)
		}

	}
}

func TestSegmentsForReportUserInsightMonthlyMfaRegisteredUserId(t *testing.T) {
	segments := ReportUserInsightMonthlyMfaRegisteredUserId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightMonthlyMfaRegisteredUserId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
