package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointCloudPCId{}

func TestNewDeviceManagementVirtualEndpointCloudPCID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue")

	if id.CloudPCId != "cloudPCIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCId'", id.CloudPCId, "cloudPCIdValue")
	}
}

func TestFormatDeviceManagementVirtualEndpointCloudPCID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointCloudPCID("cloudPCIdValue").ID()
	expected := "/deviceManagement/virtualEndpoint/cloudPCs/cloudPCIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointCloudPCID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointCloudPCId
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
			Input: "/deviceManagement/virtualEndpoint/cloudPCs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/cloudPCs/cloudPCIdValue",
			Expected: &DeviceManagementVirtualEndpointCloudPCId{
				CloudPCId: "cloudPCIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/cloudPCs/cloudPCIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointCloudPCID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCId != v.Expected.CloudPCId {
			t.Fatalf("Expected %q but got %q for CloudPCId", v.Expected.CloudPCId, actual.CloudPCId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointCloudPCIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointCloudPCId
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
			Input: "/deviceManagement/virtualEndpoint/cloudPCs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/cLoUdPcS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/cloudPCs/cloudPCIdValue",
			Expected: &DeviceManagementVirtualEndpointCloudPCId{
				CloudPCId: "cloudPCIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/cloudPCs/cloudPCIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/cLoUdPcS/cLoUdPcIdVaLuE",
			Expected: &DeviceManagementVirtualEndpointCloudPCId{
				CloudPCId: "cLoUdPcIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/cLoUdPcS/cLoUdPcIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointCloudPCIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCId != v.Expected.CloudPCId {
			t.Fatalf("Expected %q but got %q for CloudPCId", v.Expected.CloudPCId, actual.CloudPCId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointCloudPCId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointCloudPCId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointCloudPCId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
