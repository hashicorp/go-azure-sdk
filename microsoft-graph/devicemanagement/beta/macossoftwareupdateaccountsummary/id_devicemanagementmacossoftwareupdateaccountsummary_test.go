package macossoftwareupdateaccountsummary

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMacOSSoftwareUpdateAccountSummaryId{}

func TestNewDeviceManagementMacOSSoftwareUpdateAccountSummaryID(t *testing.T) {
	id := NewDeviceManagementMacOSSoftwareUpdateAccountSummaryID("macOSSoftwareUpdateAccountSummaryIdValue")

	if id.MacOSSoftwareUpdateAccountSummaryId != "macOSSoftwareUpdateAccountSummaryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MacOSSoftwareUpdateAccountSummaryId'", id.MacOSSoftwareUpdateAccountSummaryId, "macOSSoftwareUpdateAccountSummaryIdValue")
	}
}

func TestFormatDeviceManagementMacOSSoftwareUpdateAccountSummaryID(t *testing.T) {
	actual := NewDeviceManagementMacOSSoftwareUpdateAccountSummaryID("macOSSoftwareUpdateAccountSummaryIdValue").ID()
	expected := "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementMacOSSoftwareUpdateAccountSummaryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMacOSSoftwareUpdateAccountSummaryId
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
			// Valid URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue",
			Expected: &DeviceManagementMacOSSoftwareUpdateAccountSummaryId{
				MacOSSoftwareUpdateAccountSummaryId: "macOSSoftwareUpdateAccountSummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryID(v.Input)
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

	}
}

func TestParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementMacOSSoftwareUpdateAccountSummaryId
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
			// Valid URI
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue",
			Expected: &DeviceManagementMacOSSoftwareUpdateAccountSummaryId{
				MacOSSoftwareUpdateAccountSummaryId: "macOSSoftwareUpdateAccountSummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/macOSSoftwareUpdateAccountSummaries/macOSSoftwareUpdateAccountSummaryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyIdVaLuE",
			Expected: &DeviceManagementMacOSSoftwareUpdateAccountSummaryId{
				MacOSSoftwareUpdateAccountSummaryId: "mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRiEs/mAcOsSoFtWaReUpDaTeAcCoUnTsUmMaRyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementMacOSSoftwareUpdateAccountSummaryIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementMacOSSoftwareUpdateAccountSummaryId(t *testing.T) {
	segments := DeviceManagementMacOSSoftwareUpdateAccountSummaryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementMacOSSoftwareUpdateAccountSummaryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
