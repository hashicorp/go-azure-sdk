package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointReportExportJobId{}

func TestNewDeviceManagementVirtualEndpointReportExportJobID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointReportExportJobID("cloudPCExportJobId")

	if id.CloudPCExportJobId != "cloudPCExportJobId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCExportJobId'", id.CloudPCExportJobId, "cloudPCExportJobId")
	}
}

func TestFormatDeviceManagementVirtualEndpointReportExportJobID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointReportExportJobID("cloudPCExportJobId").ID()
	expected := "/deviceManagement/virtualEndpoint/reports/exportJobs/cloudPCExportJobId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointReportExportJobID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointReportExportJobId
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
			Input: "/deviceManagement/virtualEndpoint/reports/exportJobs/cloudPCExportJobId",
			Expected: &DeviceManagementVirtualEndpointReportExportJobId{
				CloudPCExportJobId: "cloudPCExportJobId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/reports/exportJobs/cloudPCExportJobId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointReportExportJobID(v.Input)
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

func TestParseDeviceManagementVirtualEndpointReportExportJobIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointReportExportJobId
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
			Input: "/deviceManagement/virtualEndpoint/reports/exportJobs/cloudPCExportJobId",
			Expected: &DeviceManagementVirtualEndpointReportExportJobId{
				CloudPCExportJobId: "cloudPCExportJobId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/reports/exportJobs/cloudPCExportJobId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/rEpOrTs/eXpOrTjObS/cLoUdPcExPoRtJoBiD",
			Expected: &DeviceManagementVirtualEndpointReportExportJobId{
				CloudPCExportJobId: "cLoUdPcExPoRtJoBiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/rEpOrTs/eXpOrTjObS/cLoUdPcExPoRtJoBiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointReportExportJobIDInsensitively(v.Input)
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

func TestSegmentsForDeviceManagementVirtualEndpointReportExportJobId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointReportExportJobId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointReportExportJobId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
