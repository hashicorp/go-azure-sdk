package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointSupportedRegionId{}

func TestNewDeviceManagementVirtualEndpointSupportedRegionID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointSupportedRegionID("cloudPCSupportedRegionId")

	if id.CloudPCSupportedRegionId != "cloudPCSupportedRegionId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCSupportedRegionId'", id.CloudPCSupportedRegionId, "cloudPCSupportedRegionId")
	}
}

func TestFormatDeviceManagementVirtualEndpointSupportedRegionID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointSupportedRegionID("cloudPCSupportedRegionId").ID()
	expected := "/deviceManagement/virtualEndpoint/supportedRegions/cloudPCSupportedRegionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointSupportedRegionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointSupportedRegionId
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
			Input: "/deviceManagement/virtualEndpoint/supportedRegions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/supportedRegions/cloudPCSupportedRegionId",
			Expected: &DeviceManagementVirtualEndpointSupportedRegionId{
				CloudPCSupportedRegionId: "cloudPCSupportedRegionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/supportedRegions/cloudPCSupportedRegionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointSupportedRegionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCSupportedRegionId != v.Expected.CloudPCSupportedRegionId {
			t.Fatalf("Expected %q but got %q for CloudPCSupportedRegionId", v.Expected.CloudPCSupportedRegionId, actual.CloudPCSupportedRegionId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointSupportedRegionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointSupportedRegionId
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
			Input: "/deviceManagement/virtualEndpoint/supportedRegions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/sUpPoRtEdReGiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/supportedRegions/cloudPCSupportedRegionId",
			Expected: &DeviceManagementVirtualEndpointSupportedRegionId{
				CloudPCSupportedRegionId: "cloudPCSupportedRegionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/supportedRegions/cloudPCSupportedRegionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/sUpPoRtEdReGiOnS/cLoUdPcSuPpOrTeDrEgIoNiD",
			Expected: &DeviceManagementVirtualEndpointSupportedRegionId{
				CloudPCSupportedRegionId: "cLoUdPcSuPpOrTeDrEgIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/sUpPoRtEdReGiOnS/cLoUdPcSuPpOrTeDrEgIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointSupportedRegionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCSupportedRegionId != v.Expected.CloudPCSupportedRegionId {
			t.Fatalf("Expected %q but got %q for CloudPCSupportedRegionId", v.Expected.CloudPCSupportedRegionId, actual.CloudPCSupportedRegionId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointSupportedRegionId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointSupportedRegionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointSupportedRegionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
