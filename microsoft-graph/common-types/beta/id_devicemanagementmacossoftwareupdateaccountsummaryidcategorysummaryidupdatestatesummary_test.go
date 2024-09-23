package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId{}

func TestNewDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryID(t *testing.T) {
	id := NewDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryID("macOSSoftwareUpdateAccountSummaryId", "macOSSoftwareUpdateCategorySummaryId", "macOSSoftwareUpdateStateSummaryId")

	if id.MacOSSoftwareUpdateAccountSummaryId != "macOSSoftwareUpdateAccountSummaryId" {
		t.Fatalf("Expected %q but got %q for Segment 'MacOSSoftwareUpdateAccountSummaryId'", id.MacOSSoftwareUpdateAccountSummaryId, "macOSSoftwareUpdateAccountSummaryId")
	}

	if id.MacOSSoftwareUpdateCategorySummaryId != "macOSSoftwareUpdateCategorySummaryId" {
		t.Fatalf("Expected %q but got %q for Segment 'MacOSSoftwareUpdateCategorySummaryId'", id.MacOSSoftwareUpdateCategorySummaryId, "macOSSoftwareUpdateCategorySummaryId")
	}

	if id.MacOSSoftwareUpdateStateSummaryId != "macOSSoftwareUpdateStateSummaryId" {
		t.Fatalf("Expected %q but got %q for Segment 'MacOSSoftwareUpdateStateSummaryId'", id.MacOSSoftwareUpdateStateSummaryId, "macOSSoftwareUpdateStateSummaryId")
	}
}

func TestFormatDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryID(t *testing.T) {
	actual := NewDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryID("macOSSoftwareUpdateAccountSummaryId", "macOSSoftwareUpdateCategorySummaryId", "macOSSoftwareUpdateStateSummaryId").ID()
	expected := "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryId/categorySummaries/macOSSoftwareUpdateCategorySummaryId/updateStateSummaries/macOSSoftwareUpdateStateSummaryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId
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
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryId/categorySummaries",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryId/categorySummaries/macOSSoftwareUpdateCategorySummaryId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryId/categorySummaries/macOSSoftwareUpdateCategorySummaryId/updateStateSummaries",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryId/categorySummaries/macOSSoftwareUpdateCategorySummaryId/updateStateSummaries/macOSSoftwareUpdateStateSummaryId",
			Expected: &DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId{
				MacOSSoftwareUpdateAccountSummaryId:  "macOSSoftwareUpdateAccountSummaryId",
				MacOSSoftwareUpdateCategorySummaryId: "macOSSoftwareUpdateCategorySummaryId",
				MacOSSoftwareUpdateStateSummaryId:    "macOSSoftwareUpdateStateSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryId/categorySummaries/macOSSoftwareUpdateCategorySummaryId/updateStateSummaries/macOSSoftwareUpdateStateSummaryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryID(v.Input)
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

func TestParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId
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
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryId/categorySummaries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyId/cAtEgOrYsUmMaRiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryId/categorySummaries/macOSSoftwareUpdateCategorySummaryId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyId/cAtEgOrYsUmMaRiEs/mAcOsSoFtWaReUpDaTeCaTeGoRySuMmArYiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryId/categorySummaries/macOSSoftwareUpdateCategorySummaryId/updateStateSummaries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyId/cAtEgOrYsUmMaRiEs/mAcOsSoFtWaReUpDaTeCaTeGoRySuMmArYiD/uPdAtEsTaTeSuMmArIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryId/categorySummaries/macOSSoftwareUpdateCategorySummaryId/updateStateSummaries/macOSSoftwareUpdateStateSummaryId",
			Expected: &DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId{
				MacOSSoftwareUpdateAccountSummaryId:  "macOSSoftwareUpdateAccountSummaryId",
				MacOSSoftwareUpdateCategorySummaryId: "macOSSoftwareUpdateCategorySummaryId",
				MacOSSoftwareUpdateStateSummaryId:    "macOSSoftwareUpdateStateSummaryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryId/categorySummaries/macOSSoftwareUpdateCategorySummaryId/updateStateSummaries/macOSSoftwareUpdateStateSummaryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyId/cAtEgOrYsUmMaRiEs/mAcOsSoFtWaReUpDaTeCaTeGoRySuMmArYiD/uPdAtEsTaTeSuMmArIeS/mAcOsSoFtWaReUpDaTeStAtEsUmMaRyId",
			Expected: &DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId{
				MacOSSoftwareUpdateAccountSummaryId:  "mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyId",
				MacOSSoftwareUpdateCategorySummaryId: "mAcOsSoFtWaReUpDaTeCaTeGoRySuMmArYiD",
				MacOSSoftwareUpdateStateSummaryId:    "mAcOsSoFtWaReUpDaTeStAtEsUmMaRyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyId/cAtEgOrYsUmMaRiEs/mAcOsSoFtWaReUpDaTeCaTeGoRySuMmArYiD/uPdAtEsTaTeSuMmArIeS/mAcOsSoFtWaReUpDaTeStAtEsUmMaRyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryIDInsensitively(v.Input)
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

func TestSegmentsForDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId(t *testing.T) {
	segments := DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdUpdateStateSummaryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
