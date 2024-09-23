package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointGalleryImageId{}

func TestNewDeviceManagementVirtualEndpointGalleryImageID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointGalleryImageID("cloudPCGalleryImageId")

	if id.CloudPCGalleryImageId != "cloudPCGalleryImageId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCGalleryImageId'", id.CloudPCGalleryImageId, "cloudPCGalleryImageId")
	}
}

func TestFormatDeviceManagementVirtualEndpointGalleryImageID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointGalleryImageID("cloudPCGalleryImageId").ID()
	expected := "/deviceManagement/virtualEndpoint/galleryImages/cloudPCGalleryImageId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointGalleryImageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointGalleryImageId
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
			Input: "/deviceManagement/virtualEndpoint/galleryImages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/galleryImages/cloudPCGalleryImageId",
			Expected: &DeviceManagementVirtualEndpointGalleryImageId{
				CloudPCGalleryImageId: "cloudPCGalleryImageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/galleryImages/cloudPCGalleryImageId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointGalleryImageID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCGalleryImageId != v.Expected.CloudPCGalleryImageId {
			t.Fatalf("Expected %q but got %q for CloudPCGalleryImageId", v.Expected.CloudPCGalleryImageId, actual.CloudPCGalleryImageId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointGalleryImageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointGalleryImageId
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
			Input: "/deviceManagement/virtualEndpoint/galleryImages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/gAlLeRyImAgEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/galleryImages/cloudPCGalleryImageId",
			Expected: &DeviceManagementVirtualEndpointGalleryImageId{
				CloudPCGalleryImageId: "cloudPCGalleryImageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/galleryImages/cloudPCGalleryImageId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/gAlLeRyImAgEs/cLoUdPcGaLlErYiMaGeId",
			Expected: &DeviceManagementVirtualEndpointGalleryImageId{
				CloudPCGalleryImageId: "cLoUdPcGaLlErYiMaGeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/gAlLeRyImAgEs/cLoUdPcGaLlErYiMaGeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointGalleryImageIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCGalleryImageId != v.Expected.CloudPCGalleryImageId {
			t.Fatalf("Expected %q but got %q for CloudPCGalleryImageId", v.Expected.CloudPCGalleryImageId, actual.CloudPCGalleryImageId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointGalleryImageId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointGalleryImageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointGalleryImageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
