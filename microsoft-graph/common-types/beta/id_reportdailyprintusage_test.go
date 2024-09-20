package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportDailyPrintUsageId{}

func TestNewReportDailyPrintUsageID(t *testing.T) {
	id := NewReportDailyPrintUsageID("printUsageId")

	if id.PrintUsageId != "printUsageId" {
		t.Fatalf("Expected %q but got %q for Segment 'PrintUsageId'", id.PrintUsageId, "printUsageId")
	}
}

func TestFormatReportDailyPrintUsageID(t *testing.T) {
	actual := NewReportDailyPrintUsageID("printUsageId").ID()
	expected := "/reports/dailyPrintUsage/printUsageId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportDailyPrintUsageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportDailyPrintUsageId
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
			Input: "/reports/dailyPrintUsage",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/dailyPrintUsage/printUsageId",
			Expected: &ReportDailyPrintUsageId{
				PrintUsageId: "printUsageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/dailyPrintUsage/printUsageId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportDailyPrintUsageID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrintUsageId != v.Expected.PrintUsageId {
			t.Fatalf("Expected %q but got %q for PrintUsageId", v.Expected.PrintUsageId, actual.PrintUsageId)
		}

	}
}

func TestParseReportDailyPrintUsageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportDailyPrintUsageId
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
			Input: "/reports/dailyPrintUsage",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/dAiLyPrInTuSaGe",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/dailyPrintUsage/printUsageId",
			Expected: &ReportDailyPrintUsageId{
				PrintUsageId: "printUsageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/dailyPrintUsage/printUsageId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/dAiLyPrInTuSaGe/pRiNtUsAgEiD",
			Expected: &ReportDailyPrintUsageId{
				PrintUsageId: "pRiNtUsAgEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/dAiLyPrInTuSaGe/pRiNtUsAgEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportDailyPrintUsageIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PrintUsageId != v.Expected.PrintUsageId {
			t.Fatalf("Expected %q but got %q for PrintUsageId", v.Expected.PrintUsageId, actual.PrintUsageId)
		}

	}
}

func TestSegmentsForReportDailyPrintUsageId(t *testing.T) {
	segments := ReportDailyPrintUsageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportDailyPrintUsageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
