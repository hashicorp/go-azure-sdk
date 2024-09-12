package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointOnPremisesConnectionId{}

func TestNewDeviceManagementVirtualEndpointOnPremisesConnectionID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointOnPremisesConnectionID("cloudPCOnPremisesConnectionIdValue")

	if id.CloudPCOnPremisesConnectionId != "cloudPCOnPremisesConnectionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCOnPremisesConnectionId'", id.CloudPCOnPremisesConnectionId, "cloudPCOnPremisesConnectionIdValue")
	}
}

func TestFormatDeviceManagementVirtualEndpointOnPremisesConnectionID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointOnPremisesConnectionID("cloudPCOnPremisesConnectionIdValue").ID()
	expected := "/deviceManagement/virtualEndpoint/onPremisesConnections/cloudPCOnPremisesConnectionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointOnPremisesConnectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointOnPremisesConnectionId
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
			Input: "/deviceManagement/virtualEndpoint/onPremisesConnections",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/onPremisesConnections/cloudPCOnPremisesConnectionIdValue",
			Expected: &DeviceManagementVirtualEndpointOnPremisesConnectionId{
				CloudPCOnPremisesConnectionId: "cloudPCOnPremisesConnectionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/onPremisesConnections/cloudPCOnPremisesConnectionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointOnPremisesConnectionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCOnPremisesConnectionId != v.Expected.CloudPCOnPremisesConnectionId {
			t.Fatalf("Expected %q but got %q for CloudPCOnPremisesConnectionId", v.Expected.CloudPCOnPremisesConnectionId, actual.CloudPCOnPremisesConnectionId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointOnPremisesConnectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointOnPremisesConnectionId
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
			Input: "/deviceManagement/virtualEndpoint/onPremisesConnections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/oNpReMiSeScOnNeCtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/onPremisesConnections/cloudPCOnPremisesConnectionIdValue",
			Expected: &DeviceManagementVirtualEndpointOnPremisesConnectionId{
				CloudPCOnPremisesConnectionId: "cloudPCOnPremisesConnectionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/onPremisesConnections/cloudPCOnPremisesConnectionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/oNpReMiSeScOnNeCtIoNs/cLoUdPcOnPrEmIsEsCoNnEcTiOnIdVaLuE",
			Expected: &DeviceManagementVirtualEndpointOnPremisesConnectionId{
				CloudPCOnPremisesConnectionId: "cLoUdPcOnPrEmIsEsCoNnEcTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/oNpReMiSeScOnNeCtIoNs/cLoUdPcOnPrEmIsEsCoNnEcTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointOnPremisesConnectionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCOnPremisesConnectionId != v.Expected.CloudPCOnPremisesConnectionId {
			t.Fatalf("Expected %q but got %q for CloudPCOnPremisesConnectionId", v.Expected.CloudPCOnPremisesConnectionId, actual.CloudPCOnPremisesConnectionId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointOnPremisesConnectionId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointOnPremisesConnectionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointOnPremisesConnectionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
