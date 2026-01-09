package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportDailyPrintUsageByUserId{}

func TestNewReportDailyPrintUsageByUserID(t *testing.T) {
	id := NewReportDailyPrintUsageByUserID("printUsageByUserId")

	if id.PrintUsageByUserId != "printUsageByUserId" {
		t.Fatalf("Expected %q but got %q for Segment 'PrintUsageByUserId'", id.PrintUsageByUserId, "printUsageByUserId")
	}
}

func TestFormatReportDailyPrintUsageByUserID(t *testing.T) {
	actual := NewReportDailyPrintUsageByUserID("printUsageByUserId").ID()
	expected := "/reports/dailyPrintUsageByUser/printUsageByUserId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportDailyPrintUsageByUserID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportDailyPrintUsageByUserId
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
			Input: "/reports/dailyPrintUsageByUser",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/dailyPrintUsageByUser/printUsageByUserId",
			Expected: &ReportDailyPrintUsageByUserId{
				PrintUsageByUserId: "printUsageByUserId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/dailyPrintUsageByUser/printUsageByUserId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportDailyPrintUsageByUserID(v.Input)
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

func TestParseReportDailyPrintUsageByUserIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportDailyPrintUsageByUserId
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
			Input: "/reports/dailyPrintUsageByUser",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/dAiLyPrInTuSaGeByUsEr",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/dailyPrintUsageByUser/printUsageByUserId",
			Expected: &ReportDailyPrintUsageByUserId{
				PrintUsageByUserId: "printUsageByUserId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/dailyPrintUsageByUser/printUsageByUserId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/dAiLyPrInTuSaGeByUsEr/pRiNtUsAgEbYuSeRiD",
			Expected: &ReportDailyPrintUsageByUserId{
				PrintUsageByUserId: "pRiNtUsAgEbYuSeRiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/dAiLyPrInTuSaGeByUsEr/pRiNtUsAgEbYuSeRiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportDailyPrintUsageByUserIDInsensitively(v.Input)
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

func TestSegmentsForReportDailyPrintUsageByUserId(t *testing.T) {
	segments := ReportDailyPrintUsageByUserId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportDailyPrintUsageByUserId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
