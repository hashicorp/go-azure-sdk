package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointSnapshotId{}

func TestNewDeviceManagementVirtualEndpointSnapshotID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointSnapshotID("cloudPCSnapshotId")

	if id.CloudPCSnapshotId != "cloudPCSnapshotId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCSnapshotId'", id.CloudPCSnapshotId, "cloudPCSnapshotId")
	}
}

func TestFormatDeviceManagementVirtualEndpointSnapshotID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointSnapshotID("cloudPCSnapshotId").ID()
	expected := "/deviceManagement/virtualEndpoint/snapshots/cloudPCSnapshotId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointSnapshotID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointSnapshotId
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
			Input: "/deviceManagement/virtualEndpoint/snapshots",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/snapshots/cloudPCSnapshotId",
			Expected: &DeviceManagementVirtualEndpointSnapshotId{
				CloudPCSnapshotId: "cloudPCSnapshotId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/snapshots/cloudPCSnapshotId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointSnapshotID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCSnapshotId != v.Expected.CloudPCSnapshotId {
			t.Fatalf("Expected %q but got %q for CloudPCSnapshotId", v.Expected.CloudPCSnapshotId, actual.CloudPCSnapshotId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointSnapshotIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointSnapshotId
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
			Input: "/deviceManagement/virtualEndpoint/snapshots",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/sNaPsHoTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/snapshots/cloudPCSnapshotId",
			Expected: &DeviceManagementVirtualEndpointSnapshotId{
				CloudPCSnapshotId: "cloudPCSnapshotId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/snapshots/cloudPCSnapshotId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/sNaPsHoTs/cLoUdPcSnApShOtId",
			Expected: &DeviceManagementVirtualEndpointSnapshotId{
				CloudPCSnapshotId: "cLoUdPcSnApShOtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/sNaPsHoTs/cLoUdPcSnApShOtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointSnapshotIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCSnapshotId != v.Expected.CloudPCSnapshotId {
			t.Fatalf("Expected %q but got %q for CloudPCSnapshotId", v.Expected.CloudPCSnapshotId, actual.CloudPCSnapshotId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointSnapshotId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointSnapshotId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointSnapshotId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
