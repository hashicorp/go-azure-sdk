package updatedevicemanagementmacossoftwareupdateaccountsummarycategorysummarystatesummaries

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId{}

func TestNewUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesID(t *testing.T) {
	id := NewUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesID("macOSSoftwareUpdateAccountSummaryIdValue", "macOSSoftwareUpdateCategorySummaryIdValue", "macOSSoftwareUpdateStateSummaryIdValue")

	if id.MacOSSoftwareUpdateAccountSummaryId != "macOSSoftwareUpdateAccountSummaryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MacOSSoftwareUpdateAccountSummaryId'", id.MacOSSoftwareUpdateAccountSummaryId, "macOSSoftwareUpdateAccountSummaryIdValue")
	}

	if id.MacOSSoftwareUpdateCategorySummaryId != "macOSSoftwareUpdateCategorySummaryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MacOSSoftwareUpdateCategorySummaryId'", id.MacOSSoftwareUpdateCategorySummaryId, "macOSSoftwareUpdateCategorySummaryIdValue")
	}

	if id.MacOSSoftwareUpdateStateSummaryId != "macOSSoftwareUpdateStateSummaryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MacOSSoftwareUpdateStateSummaryId'", id.MacOSSoftwareUpdateStateSummaryId, "macOSSoftwareUpdateStateSummaryIdValue")
	}
}

func TestFormatUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesID(t *testing.T) {
	actual := NewUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesID("macOSSoftwareUpdateAccountSummaryIdValue", "macOSSoftwareUpdateCategorySummaryIdValue", "macOSSoftwareUpdateStateSummaryIdValue").ID()
	expected := "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/categorySummaries/macOSSoftwareUpdateCategorySummaryIdValue/updateStateSummaries/macOSSoftwareUpdateStateSummaryIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/categorySummaries",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/categorySummaries/macOSSoftwareUpdateCategorySummaryIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/categorySummaries/macOSSoftwareUpdateCategorySummaryIdValue/updateStateSummaries",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/categorySummaries/macOSSoftwareUpdateCategorySummaryIdValue/updateStateSummaries/macOSSoftwareUpdateStateSummaryIdValue",
			Expected: &UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId{
				MacOSSoftwareUpdateAccountSummaryId:  "macOSSoftwareUpdateAccountSummaryIdValue",
				MacOSSoftwareUpdateCategorySummaryId: "macOSSoftwareUpdateCategorySummaryIdValue",
				MacOSSoftwareUpdateStateSummaryId:    "macOSSoftwareUpdateStateSummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/categorySummaries/macOSSoftwareUpdateCategorySummaryIdValue/updateStateSummaries/macOSSoftwareUpdateStateSummaryIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MacOSSoftwareUpdateAccountSummaryId != v.Expected.MacOSSoftwareUpdateAccountSummaryId {
			t.Fatalf("Expected %q but got %q for MacOSSoftwareUpdateAccountSummaryId", v.Expected.MacOSSoftwareUpdateAccountSummaryId, actual.MacOSSoftwareUpdateAccountSummaryId)
		}

		if actual.MacOSSoftwareUpdateCategorySummaryId != v.Expected.MacOSSoftwareUpdateCategorySummaryId {
			t.Fatalf("Expected %q but got %q for MacOSSoftwareUpdateCategorySummaryId", v.Expected.MacOSSoftwareUpdateCategorySummaryId, actual.MacOSSoftwareUpdateCategorySummaryId)
		}

		if actual.MacOSSoftwareUpdateStateSummaryId != v.Expected.MacOSSoftwareUpdateStateSummaryId {
			t.Fatalf("Expected %q but got %q for MacOSSoftwareUpdateStateSummaryId", v.Expected.MacOSSoftwareUpdateStateSummaryId, actual.MacOSSoftwareUpdateStateSummaryId)
		}

	}
}

func TestParseUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/categorySummaries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyIdVaLuE/cAtEgOrYsUmMaRiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/categorySummaries/macOSSoftwareUpdateCategorySummaryIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyIdVaLuE/cAtEgOrYsUmMaRiEs/mAcOsSoFtWaReUpDaTeCaTeGoRySuMmArYiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/categorySummaries/macOSSoftwareUpdateCategorySummaryIdValue/updateStateSummaries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyIdVaLuE/cAtEgOrYsUmMaRiEs/mAcOsSoFtWaReUpDaTeCaTeGoRySuMmArYiDvAlUe/uPdAtEsTaTeSuMmArIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/categorySummaries/macOSSoftwareUpdateCategorySummaryIdValue/updateStateSummaries/macOSSoftwareUpdateStateSummaryIdValue",
			Expected: &UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId{
				MacOSSoftwareUpdateAccountSummaryId:  "macOSSoftwareUpdateAccountSummaryIdValue",
				MacOSSoftwareUpdateCategorySummaryId: "macOSSoftwareUpdateCategorySummaryIdValue",
				MacOSSoftwareUpdateStateSummaryId:    "macOSSoftwareUpdateStateSummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/categorySummaries/macOSSoftwareUpdateCategorySummaryIdValue/updateStateSummaries/macOSSoftwareUpdateStateSummaryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyIdVaLuE/cAtEgOrYsUmMaRiEs/mAcOsSoFtWaReUpDaTeCaTeGoRySuMmArYiDvAlUe/uPdAtEsTaTeSuMmArIeS/mAcOsSoFtWaReUpDaTeStAtEsUmMaRyIdVaLuE",
			Expected: &UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId{
				MacOSSoftwareUpdateAccountSummaryId:  "mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyIdVaLuE",
				MacOSSoftwareUpdateCategorySummaryId: "mAcOsSoFtWaReUpDaTeCaTeGoRySuMmArYiDvAlUe",
				MacOSSoftwareUpdateStateSummaryId:    "mAcOsSoFtWaReUpDaTeStAtEsUmMaRyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyIdVaLuE/cAtEgOrYsUmMaRiEs/mAcOsSoFtWaReUpDaTeCaTeGoRySuMmArYiDvAlUe/uPdAtEsTaTeSuMmArIeS/mAcOsSoFtWaReUpDaTeStAtEsUmMaRyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MacOSSoftwareUpdateAccountSummaryId != v.Expected.MacOSSoftwareUpdateAccountSummaryId {
			t.Fatalf("Expected %q but got %q for MacOSSoftwareUpdateAccountSummaryId", v.Expected.MacOSSoftwareUpdateAccountSummaryId, actual.MacOSSoftwareUpdateAccountSummaryId)
		}

		if actual.MacOSSoftwareUpdateCategorySummaryId != v.Expected.MacOSSoftwareUpdateCategorySummaryId {
			t.Fatalf("Expected %q but got %q for MacOSSoftwareUpdateCategorySummaryId", v.Expected.MacOSSoftwareUpdateCategorySummaryId, actual.MacOSSoftwareUpdateCategorySummaryId)
		}

		if actual.MacOSSoftwareUpdateStateSummaryId != v.Expected.MacOSSoftwareUpdateStateSummaryId {
			t.Fatalf("Expected %q but got %q for MacOSSoftwareUpdateStateSummaryId", v.Expected.MacOSSoftwareUpdateStateSummaryId, actual.MacOSSoftwareUpdateStateSummaryId)
		}

	}
}

func TestSegmentsForUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId(t *testing.T) {
	segments := UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
