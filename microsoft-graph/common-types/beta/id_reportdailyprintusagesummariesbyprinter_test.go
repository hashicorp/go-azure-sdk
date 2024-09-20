package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportDailyPrintUsageSummariesByPrinterId{}

func TestNewReportDailyPrintUsageSummariesByPrinterID(t *testing.T) {
	id := NewReportDailyPrintUsageSummariesByPrinterID("printUsageByPrinterId")

	if id.PrintUsageByPrinterId != "printUsageByPrinterId" {
		t.Fatalf("Expected %q but got %q for Segment 'PrintUsageByPrinterId'", id.PrintUsageByPrinterId, "printUsageByPrinterId")
	}
}

func TestFormatReportDailyPrintUsageSummariesByPrinterID(t *testing.T) {
	actual := NewReportDailyPrintUsageSummariesByPrinterID("printUsageByPrinterId").ID()
	expected := "/reports/dailyPrintUsageSummariesByPrinter/printUsageByPrinterId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportDailyPrintUsageSummariesByPrinterID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportDailyPrintUsageSummariesByPrinterId
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
			Input: "/reports/dailyPrintUsageSummariesByPrinter",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/dailyPrintUsageSummariesByPrinter/printUsageByPrinterId",
			Expected: &ReportDailyPrintUsageSummariesByPrinterId{
				PrintUsageByPrinterId: "printUsageByPrinterId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/dailyPrintUsageSummariesByPrinter/printUsageByPrinterId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportDailyPrintUsageSummariesByPrinterID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrintUsageByPrinterId != v.Expected.PrintUsageByPrinterId {
			t.Fatalf("Expected %q but got %q for PrintUsageByPrinterId", v.Expected.PrintUsageByPrinterId, actual.PrintUsageByPrinterId)
		}

	}
}

func TestParseReportDailyPrintUsageSummariesByPrinterIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportDailyPrintUsageSummariesByPrinterId
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
			Input: "/reports/dailyPrintUsageSummariesByPrinter",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/dAiLyPrInTuSaGeSuMmArIeSbYpRiNtEr",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/dailyPrintUsageSummariesByPrinter/printUsageByPrinterId",
			Expected: &ReportDailyPrintUsageSummariesByPrinterId{
				PrintUsageByPrinterId: "printUsageByPrinterId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/dailyPrintUsageSummariesByPrinter/printUsageByPrinterId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/dAiLyPrInTuSaGeSuMmArIeSbYpRiNtEr/pRiNtUsAgEbYpRiNtErId",
			Expected: &ReportDailyPrintUsageSummariesByPrinterId{
				PrintUsageByPrinterId: "pRiNtUsAgEbYpRiNtErId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/dAiLyPrInTuSaGeSuMmArIeSbYpRiNtEr/pRiNtUsAgEbYpRiNtErId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportDailyPrintUsageSummariesByPrinterIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrintUsageByPrinterId != v.Expected.PrintUsageByPrinterId {
			t.Fatalf("Expected %q but got %q for PrintUsageByPrinterId", v.Expected.PrintUsageByPrinterId, actual.PrintUsageByPrinterId)
		}

	}
}

func TestSegmentsForReportDailyPrintUsageSummariesByPrinterId(t *testing.T) {
	segments := ReportDailyPrintUsageSummariesByPrinterId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportDailyPrintUsageSummariesByPrinterId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
