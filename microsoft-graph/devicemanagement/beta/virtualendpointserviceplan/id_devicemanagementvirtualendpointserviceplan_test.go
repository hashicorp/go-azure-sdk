package virtualendpointserviceplan

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointServicePlanId{}

func TestNewDeviceManagementVirtualEndpointServicePlanID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointServicePlanID("cloudPCServicePlanIdValue")

	if id.CloudPCServicePlanId != "cloudPCServicePlanIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCServicePlanId'", id.CloudPCServicePlanId, "cloudPCServicePlanIdValue")
	}
}

func TestFormatDeviceManagementVirtualEndpointServicePlanID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointServicePlanID("cloudPCServicePlanIdValue").ID()
	expected := "/deviceManagement/virtualEndpoint/servicePlans/cloudPCServicePlanIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointServicePlanID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointServicePlanId
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
			Input: "/deviceManagement/virtualEndpoint/servicePlans",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/servicePlans/cloudPCServicePlanIdValue",
			Expected: &DeviceManagementVirtualEndpointServicePlanId{
				CloudPCServicePlanId: "cloudPCServicePlanIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/servicePlans/cloudPCServicePlanIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointServicePlanID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCServicePlanId != v.Expected.CloudPCServicePlanId {
			t.Fatalf("Expected %q but got %q for CloudPCServicePlanId", v.Expected.CloudPCServicePlanId, actual.CloudPCServicePlanId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointServicePlanIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointServicePlanId
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
			Input: "/deviceManagement/virtualEndpoint/servicePlans",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/sErViCePlAnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/servicePlans/cloudPCServicePlanIdValue",
			Expected: &DeviceManagementVirtualEndpointServicePlanId{
				CloudPCServicePlanId: "cloudPCServicePlanIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/servicePlans/cloudPCServicePlanIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/sErViCePlAnS/cLoUdPcSeRvIcEpLaNiDvAlUe",
			Expected: &DeviceManagementVirtualEndpointServicePlanId{
				CloudPCServicePlanId: "cLoUdPcSeRvIcEpLaNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/sErViCePlAnS/cLoUdPcSeRvIcEpLaNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointServicePlanIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCServicePlanId != v.Expected.CloudPCServicePlanId {
			t.Fatalf("Expected %q but got %q for CloudPCServicePlanId", v.Expected.CloudPCServicePlanId, actual.CloudPCServicePlanId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointServicePlanId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointServicePlanId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointServicePlanId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
