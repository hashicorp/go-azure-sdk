package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportMonthlyPrintUsageSummariesByPrinterId{}

func TestNewReportMonthlyPrintUsageSummariesByPrinterID(t *testing.T) {
	id := NewReportMonthlyPrintUsageSummariesByPrinterID("printUsageByPrinterId")

	if id.PrintUsageByPrinterId != "printUsageByPrinterId" {
		t.Fatalf("Expected %q but got %q for Segment 'PrintUsageByPrinterId'", id.PrintUsageByPrinterId, "printUsageByPrinterId")
	}
}

func TestFormatReportMonthlyPrintUsageSummariesByPrinterID(t *testing.T) {
	actual := NewReportMonthlyPrintUsageSummariesByPrinterID("printUsageByPrinterId").ID()
	expected := "/reports/monthlyPrintUsageSummariesByPrinter/printUsageByPrinterId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportMonthlyPrintUsageSummariesByPrinterID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportMonthlyPrintUsageSummariesByPrinterId
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
			Input: "/reports/monthlyPrintUsageSummariesByPrinter",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/monthlyPrintUsageSummariesByPrinter/printUsageByPrinterId",
			Expected: &ReportMonthlyPrintUsageSummariesByPrinterId{
				PrintUsageByPrinterId: "printUsageByPrinterId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/monthlyPrintUsageSummariesByPrinter/printUsageByPrinterId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportMonthlyPrintUsageSummariesByPrinterID(v.Input)
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

func TestParseReportMonthlyPrintUsageSummariesByPrinterIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportMonthlyPrintUsageSummariesByPrinterId
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
			Input: "/reports/monthlyPrintUsageSummariesByPrinter",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/mOnThLyPrInTuSaGeSuMmArIeSbYpRiNtEr",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/monthlyPrintUsageSummariesByPrinter/printUsageByPrinterId",
			Expected: &ReportMonthlyPrintUsageSummariesByPrinterId{
				PrintUsageByPrinterId: "printUsageByPrinterId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/monthlyPrintUsageSummariesByPrinter/printUsageByPrinterId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/mOnThLyPrInTuSaGeSuMmArIeSbYpRiNtEr/pRiNtUsAgEbYpRiNtErId",
			Expected: &ReportMonthlyPrintUsageSummariesByPrinterId{
				PrintUsageByPrinterId: "pRiNtUsAgEbYpRiNtErId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/mOnThLyPrInTuSaGeSuMmArIeSbYpRiNtEr/pRiNtUsAgEbYpRiNtErId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportMonthlyPrintUsageSummariesByPrinterIDInsensitively(v.Input)
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

func TestSegmentsForReportMonthlyPrintUsageSummariesByPrinterId(t *testing.T) {
	segments := ReportMonthlyPrintUsageSummariesByPrinterId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportMonthlyPrintUsageSummariesByPrinterId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
