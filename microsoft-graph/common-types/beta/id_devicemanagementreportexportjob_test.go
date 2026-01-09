package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementReportExportJobId{}

func TestNewDeviceManagementReportExportJobID(t *testing.T) {
	id := NewDeviceManagementReportExportJobID("deviceManagementExportJobId")

	if id.DeviceManagementExportJobId != "deviceManagementExportJobId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementExportJobId'", id.DeviceManagementExportJobId, "deviceManagementExportJobId")
	}
}

func TestFormatDeviceManagementReportExportJobID(t *testing.T) {
	actual := NewDeviceManagementReportExportJobID("deviceManagementExportJobId").ID()
	expected := "/deviceManagement/reports/exportJobs/deviceManagementExportJobId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementReportExportJobID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementReportExportJobId
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
			Input: "/deviceManagement/reports/exportJobs/deviceManagementExportJobId",
			Expected: &DeviceManagementReportExportJobId{
				DeviceManagementExportJobId: "deviceManagementExportJobId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reports/exportJobs/deviceManagementExportJobId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementReportExportJobID(v.Input)
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

func TestParseDeviceManagementReportExportJobIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementReportExportJobId
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
			Input: "/deviceManagement/reports/exportJobs/deviceManagementExportJobId",
			Expected: &DeviceManagementReportExportJobId{
				DeviceManagementExportJobId: "deviceManagementExportJobId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/reports/exportJobs/deviceManagementExportJobId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEpOrTs/eXpOrTjObS/dEvIcEmAnAgEmEnTeXpOrTjObId",
			Expected: &DeviceManagementReportExportJobId{
				DeviceManagementExportJobId: "dEvIcEmAnAgEmEnTeXpOrTjObId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEpOrTs/eXpOrTjObS/dEvIcEmAnAgEmEnTeXpOrTjObId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementReportExportJobIDInsensitively(v.Input)
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

func TestSegmentsForDeviceManagementReportExportJobId(t *testing.T) {
	segments := DeviceManagementReportExportJobId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementReportExportJobId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
