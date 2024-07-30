package dailyprintusagesummariesbyuser

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportDailyPrintUsageSummariesByUserId{}

func TestNewReportDailyPrintUsageSummariesByUserID(t *testing.T) {
	id := NewReportDailyPrintUsageSummariesByUserID("printUsageByUserIdValue")

	if id.PrintUsageByUserId != "printUsageByUserIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PrintUsageByUserId'", id.PrintUsageByUserId, "printUsageByUserIdValue")
	}
}

func TestFormatReportDailyPrintUsageSummariesByUserID(t *testing.T) {
	actual := NewReportDailyPrintUsageSummariesByUserID("printUsageByUserIdValue").ID()
	expected := "/reports/dailyPrintUsageSummariesByUser/printUsageByUserIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportDailyPrintUsageSummariesByUserID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportDailyPrintUsageSummariesByUserId
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
			Input: "/reports/dailyPrintUsageSummariesByUser",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/dailyPrintUsageSummariesByUser/printUsageByUserIdValue",
			Expected: &ReportDailyPrintUsageSummariesByUserId{
				PrintUsageByUserId: "printUsageByUserIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/dailyPrintUsageSummariesByUser/printUsageByUserIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportDailyPrintUsageSummariesByUserID(v.Input)
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

func TestParseReportDailyPrintUsageSummariesByUserIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportDailyPrintUsageSummariesByUserId
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
			Input: "/reports/dailyPrintUsageSummariesByUser",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/dAiLyPrInTuSaGeSuMmArIeSbYuSeR",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/dailyPrintUsageSummariesByUser/printUsageByUserIdValue",
			Expected: &ReportDailyPrintUsageSummariesByUserId{
				PrintUsageByUserId: "printUsageByUserIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/dailyPrintUsageSummariesByUser/printUsageByUserIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/dAiLyPrInTuSaGeSuMmArIeSbYuSeR/pRiNtUsAgEbYuSeRiDvAlUe",
			Expected: &ReportDailyPrintUsageSummariesByUserId{
				PrintUsageByUserId: "pRiNtUsAgEbYuSeRiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/dAiLyPrInTuSaGeSuMmArIeSbYuSeR/pRiNtUsAgEbYuSeRiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportDailyPrintUsageSummariesByUserIDInsensitively(v.Input)
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

func TestSegmentsForReportDailyPrintUsageSummariesByUserId(t *testing.T) {
	segments := ReportDailyPrintUsageSummariesByUserId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportDailyPrintUsageSummariesByUserId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
