package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ExportDeviceManagementReportsJobsId{}

func TestNewExportDeviceManagementReportsJobsID(t *testing.T) {
	id := NewExportDeviceManagementReportsJobsID("deviceManagementExportJobIdValue")

	if id.DeviceManagementExportJobId != "deviceManagementExportJobIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementExportJobId'", id.DeviceManagementExportJobId, "deviceManagementExportJobIdValue")
	}
}

func TestFormatExportDeviceManagementReportsJobsID(t *testing.T) {
	actual := NewExportDeviceManagementReportsJobsID("deviceManagementExportJobIdValue").ID()
	expected := "/deviceManagement/reports/exportJobs/deviceManagementExportJobIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseExportDeviceManagementReportsJobsID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExportDeviceManagementReportsJobsId
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
			Input: "/deviceManagement/reports",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reports/exportJobs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/reports/exportJobs/deviceManagementExportJobIdValue",
			Expected: &ExportDeviceManagementReportsJobsId{
				DeviceManagementExportJobId: "deviceManagementExportJobIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reports/exportJobs/deviceManagementExportJobIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExportDeviceManagementReportsJobsID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementExportJobId != v.Expected.DeviceManagementExportJobId {
			t.Fatalf("Expected %q but got %q for DeviceManagementExportJobId", v.Expected.DeviceManagementExportJobId, actual.DeviceManagementExportJobId)
		}

	}
}

func TestParseExportDeviceManagementReportsJobsIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExportDeviceManagementReportsJobsId
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
			Input: "/deviceManagement/reports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEpOrTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/reports/exportJobs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEpOrTs/eXpOrTjObS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/reports/exportJobs/deviceManagementExportJobIdValue",
			Expected: &ExportDeviceManagementReportsJobsId{
				DeviceManagementExportJobId: "deviceManagementExportJobIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reports/exportJobs/deviceManagementExportJobIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEpOrTs/eXpOrTjObS/dEvIcEmAnAgEmEnTeXpOrTjObIdVaLuE",
			Expected: &ExportDeviceManagementReportsJobsId{
				DeviceManagementExportJobId: "dEvIcEmAnAgEmEnTeXpOrTjObIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEpOrTs/eXpOrTjObS/dEvIcEmAnAgEmEnTeXpOrTjObIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExportDeviceManagementReportsJobsIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementExportJobId != v.Expected.DeviceManagementExportJobId {
			t.Fatalf("Expected %q but got %q for DeviceManagementExportJobId", v.Expected.DeviceManagementExportJobId, actual.DeviceManagementExportJobId)
		}

	}
}

func TestSegmentsForExportDeviceManagementReportsJobsId(t *testing.T) {
	segments := ExportDeviceManagementReportsJobsId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ExportDeviceManagementReportsJobsId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
