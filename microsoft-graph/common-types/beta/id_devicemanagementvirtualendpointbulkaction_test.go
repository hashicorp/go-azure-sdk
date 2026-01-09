package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointBulkActionId{}

func TestNewDeviceManagementVirtualEndpointBulkActionID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointBulkActionID("cloudPCBulkActionId")

	if id.CloudPCBulkActionId != "cloudPCBulkActionId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCBulkActionId'", id.CloudPCBulkActionId, "cloudPCBulkActionId")
	}
}

func TestFormatDeviceManagementVirtualEndpointBulkActionID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointBulkActionID("cloudPCBulkActionId").ID()
	expected := "/deviceManagement/virtualEndpoint/bulkActions/cloudPCBulkActionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointBulkActionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointBulkActionId
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
			Input: "/deviceManagement/virtualEndpoint/bulkActions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/bulkActions/cloudPCBulkActionId",
			Expected: &DeviceManagementVirtualEndpointBulkActionId{
				CloudPCBulkActionId: "cloudPCBulkActionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/bulkActions/cloudPCBulkActionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointBulkActionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCBulkActionId != v.Expected.CloudPCBulkActionId {
			t.Fatalf("Expected %q but got %q for CloudPCBulkActionId", v.Expected.CloudPCBulkActionId, actual.CloudPCBulkActionId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointBulkActionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointBulkActionId
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
			Input: "/deviceManagement/virtualEndpoint/bulkActions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/bUlKaCtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/bulkActions/cloudPCBulkActionId",
			Expected: &DeviceManagementVirtualEndpointBulkActionId{
				CloudPCBulkActionId: "cloudPCBulkActionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/bulkActions/cloudPCBulkActionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/bUlKaCtIoNs/cLoUdPcBuLkAcTiOnId",
			Expected: &DeviceManagementVirtualEndpointBulkActionId{
				CloudPCBulkActionId: "cLoUdPcBuLkAcTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/bUlKaCtIoNs/cLoUdPcBuLkAcTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointBulkActionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCBulkActionId != v.Expected.CloudPCBulkActionId {
			t.Fatalf("Expected %q but got %q for CloudPCBulkActionId", v.Expected.CloudPCBulkActionId, actual.CloudPCBulkActionId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointBulkActionId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointBulkActionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointBulkActionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
