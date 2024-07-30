package monthlyprintusagesummariesbyuser

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportMonthlyPrintUsageSummariesByUserId{}

func TestNewReportMonthlyPrintUsageSummariesByUserID(t *testing.T) {
	id := NewReportMonthlyPrintUsageSummariesByUserID("printUsageByUserIdValue")

	if id.PrintUsageByUserId != "printUsageByUserIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PrintUsageByUserId'", id.PrintUsageByUserId, "printUsageByUserIdValue")
	}
}

func TestFormatReportMonthlyPrintUsageSummariesByUserID(t *testing.T) {
	actual := NewReportMonthlyPrintUsageSummariesByUserID("printUsageByUserIdValue").ID()
	expected := "/reports/monthlyPrintUsageSummariesByUser/printUsageByUserIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportMonthlyPrintUsageSummariesByUserID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportMonthlyPrintUsageSummariesByUserId
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
			Input: "/reports/monthlyPrintUsageSummariesByUser",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/monthlyPrintUsageSummariesByUser/printUsageByUserIdValue",
			Expected: &ReportMonthlyPrintUsageSummariesByUserId{
				PrintUsageByUserId: "printUsageByUserIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/monthlyPrintUsageSummariesByUser/printUsageByUserIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportMonthlyPrintUsageSummariesByUserID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrintUsageByUserId != v.Expected.PrintUsageByUserId {
			t.Fatalf("Expected %q but got %q for PrintUsageByUserId", v.Expected.PrintUsageByUserId, actual.PrintUsageByUserId)
		}

	}
}

func TestParseReportMonthlyPrintUsageSummariesByUserIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportMonthlyPrintUsageSummariesByUserId
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
			Input: "/reports/monthlyPrintUsageSummariesByUser",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/mOnThLyPrInTuSaGeSuMmArIeSbYuSeR",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/monthlyPrintUsageSummariesByUser/printUsageByUserIdValue",
			Expected: &ReportMonthlyPrintUsageSummariesByUserId{
				PrintUsageByUserId: "printUsageByUserIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/monthlyPrintUsageSummariesByUser/printUsageByUserIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/mOnThLyPrInTuSaGeSuMmArIeSbYuSeR/pRiNtUsAgEbYuSeRiDvAlUe",
			Expected: &ReportMonthlyPrintUsageSummariesByUserId{
				PrintUsageByUserId: "pRiNtUsAgEbYuSeRiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/mOnThLyPrInTuSaGeSuMmArIeSbYuSeR/pRiNtUsAgEbYuSeRiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportMonthlyPrintUsageSummariesByUserIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrintUsageByUserId != v.Expected.PrintUsageByUserId {
			t.Fatalf("Expected %q but got %q for PrintUsageByUserId", v.Expected.PrintUsageByUserId, actual.PrintUsageByUserId)
		}

	}
}

func TestSegmentsForReportMonthlyPrintUsageSummariesByUserId(t *testing.T) {
	segments := ReportMonthlyPrintUsageSummariesByUserId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportMonthlyPrintUsageSummariesByUserId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
