package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId{}

func TestNewDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryID(t *testing.T) {
	id := NewDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryID("macOSSoftwareUpdateAccountSummaryIdValue", "macOSSoftwareUpdateCategorySummaryIdValue")

	if id.MacOSSoftwareUpdateAccountSummaryId != "macOSSoftwareUpdateAccountSummaryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MacOSSoftwareUpdateAccountSummaryId'", id.MacOSSoftwareUpdateAccountSummaryId, "macOSSoftwareUpdateAccountSummaryIdValue")
	}

	if id.MacOSSoftwareUpdateCategorySummaryId != "macOSSoftwareUpdateCategorySummaryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MacOSSoftwareUpdateCategorySummaryId'", id.MacOSSoftwareUpdateCategorySummaryId, "macOSSoftwareUpdateCategorySummaryIdValue")
	}
}

func TestFormatDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryID(t *testing.T) {
	actual := NewDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryID("macOSSoftwareUpdateAccountSummaryIdValue", "macOSSoftwareUpdateCategorySummaryIdValue").ID()
	expected := "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/categorySummaries/macOSSoftwareUpdateCategorySummaryIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId
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
			// Valid URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/categorySummaries/macOSSoftwareUpdateCategorySummaryIdValue",
			Expected: &DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId{
				MacOSSoftwareUpdateAccountSummaryId:  "macOSSoftwareUpdateAccountSummaryIdValue",
				MacOSSoftwareUpdateCategorySummaryId: "macOSSoftwareUpdateCategorySummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/categorySummaries/macOSSoftwareUpdateCategorySummaryIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryID(v.Input)
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

	}
}

func TestParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId
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
			// Valid URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/categorySummaries/macOSSoftwareUpdateCategorySummaryIdValue",
			Expected: &DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId{
				MacOSSoftwareUpdateAccountSummaryId:  "macOSSoftwareUpdateAccountSummaryIdValue",
				MacOSSoftwareUpdateCategorySummaryId: "macOSSoftwareUpdateCategorySummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/categorySummaries/macOSSoftwareUpdateCategorySummaryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyIdVaLuE/cAtEgOrYsUmMaRiEs/mAcOsSoFtWaReUpDaTeCaTeGoRySuMmArYiDvAlUe",
			Expected: &DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId{
				MacOSSoftwareUpdateAccountSummaryId:  "mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyIdVaLuE",
				MacOSSoftwareUpdateCategorySummaryId: "mAcOsSoFtWaReUpDaTeCaTeGoRySuMmArYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyIdVaLuE/cAtEgOrYsUmMaRiEs/mAcOsSoFtWaReUpDaTeCaTeGoRySuMmArYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId(t *testing.T) {
	segments := DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
