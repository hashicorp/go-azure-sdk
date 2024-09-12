package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ExportDeviceManagementVirtualEndpointReportsJobsId{}

func TestNewExportDeviceManagementVirtualEndpointReportsJobsID(t *testing.T) {
	id := NewExportDeviceManagementVirtualEndpointReportsJobsID("cloudPCExportJobIdValue")

	if id.CloudPCExportJobId != "cloudPCExportJobIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCExportJobId'", id.CloudPCExportJobId, "cloudPCExportJobIdValue")
	}
}

func TestFormatExportDeviceManagementVirtualEndpointReportsJobsID(t *testing.T) {
	actual := NewExportDeviceManagementVirtualEndpointReportsJobsID("cloudPCExportJobIdValue").ID()
	expected := "/deviceManagement/virtualEndpoint/reports/exportJobs/cloudPCExportJobIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseExportDeviceManagementVirtualEndpointReportsJobsID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExportDeviceManagementVirtualEndpointReportsJobsId
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
			Input: "/deviceManagement/virtualEndpoint",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/reports",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/reports/exportJobs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/reports/exportJobs/cloudPCExportJobIdValue",
			Expected: &ExportDeviceManagementVirtualEndpointReportsJobsId{
				CloudPCExportJobId: "cloudPCExportJobIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/reports/exportJobs/cloudPCExportJobIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExportDeviceManagementVirtualEndpointReportsJobsID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCExportJobId != v.Expected.CloudPCExportJobId {
			t.Fatalf("Expected %q but got %q for CloudPCExportJobId", v.Expected.CloudPCExportJobId, actual.CloudPCExportJobId)
		}

	}
}

func TestParseExportDeviceManagementVirtualEndpointReportsJobsIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExportDeviceManagementVirtualEndpointReportsJobsId
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
			Input: "/deviceManagement/virtualEndpoint",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/reports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/rEpOrTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/reports/exportJobs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/rEpOrTs/eXpOrTjObS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/reports/exportJobs/cloudPCExportJobIdValue",
			Expected: &ExportDeviceManagementVirtualEndpointReportsJobsId{
				CloudPCExportJobId: "cloudPCExportJobIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/reports/exportJobs/cloudPCExportJobIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/rEpOrTs/eXpOrTjObS/cLoUdPcExPoRtJoBiDvAlUe",
			Expected: &ExportDeviceManagementVirtualEndpointReportsJobsId{
				CloudPCExportJobId: "cLoUdPcExPoRtJoBiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/rEpOrTs/eXpOrTjObS/cLoUdPcExPoRtJoBiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExportDeviceManagementVirtualEndpointReportsJobsIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCExportJobId != v.Expected.CloudPCExportJobId {
			t.Fatalf("Expected %q but got %q for CloudPCExportJobId", v.Expected.CloudPCExportJobId, actual.CloudPCExportJobId)
		}

	}
}

func TestSegmentsForExportDeviceManagementVirtualEndpointReportsJobsId(t *testing.T) {
	segments := ExportDeviceManagementVirtualEndpointReportsJobsId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ExportDeviceManagementVirtualEndpointReportsJobsId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
