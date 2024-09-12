package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointDeviceImageId{}

func TestNewDeviceManagementVirtualEndpointDeviceImageID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointDeviceImageID("cloudPCDeviceImageIdValue")

	if id.CloudPCDeviceImageId != "cloudPCDeviceImageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCDeviceImageId'", id.CloudPCDeviceImageId, "cloudPCDeviceImageIdValue")
	}
}

func TestFormatDeviceManagementVirtualEndpointDeviceImageID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointDeviceImageID("cloudPCDeviceImageIdValue").ID()
	expected := "/deviceManagement/virtualEndpoint/deviceImages/cloudPCDeviceImageIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointDeviceImageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointDeviceImageId
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
			Input: "/deviceManagement/virtualEndpoint/deviceImages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/deviceImages/cloudPCDeviceImageIdValue",
			Expected: &DeviceManagementVirtualEndpointDeviceImageId{
				CloudPCDeviceImageId: "cloudPCDeviceImageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/deviceImages/cloudPCDeviceImageIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointDeviceImageID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCDeviceImageId != v.Expected.CloudPCDeviceImageId {
			t.Fatalf("Expected %q but got %q for CloudPCDeviceImageId", v.Expected.CloudPCDeviceImageId, actual.CloudPCDeviceImageId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointDeviceImageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointDeviceImageId
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
			Input: "/deviceManagement/virtualEndpoint/deviceImages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/dEvIcEiMaGeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/deviceImages/cloudPCDeviceImageIdValue",
			Expected: &DeviceManagementVirtualEndpointDeviceImageId{
				CloudPCDeviceImageId: "cloudPCDeviceImageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/deviceImages/cloudPCDeviceImageIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/dEvIcEiMaGeS/cLoUdPcDeViCeImAgEiDvAlUe",
			Expected: &DeviceManagementVirtualEndpointDeviceImageId{
				CloudPCDeviceImageId: "cLoUdPcDeViCeImAgEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/dEvIcEiMaGeS/cLoUdPcDeViCeImAgEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointDeviceImageIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCDeviceImageId != v.Expected.CloudPCDeviceImageId {
			t.Fatalf("Expected %q but got %q for CloudPCDeviceImageId", v.Expected.CloudPCDeviceImageId, actual.CloudPCDeviceImageId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointDeviceImageId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointDeviceImageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointDeviceImageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
