package virtualendpointfrontlineserviceplan

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointFrontLineServicePlanId{}

func TestNewDeviceManagementVirtualEndpointFrontLineServicePlanID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointFrontLineServicePlanID("cloudPCFrontLineServicePlanIdValue")

	if id.CloudPCFrontLineServicePlanId != "cloudPCFrontLineServicePlanIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCFrontLineServicePlanId'", id.CloudPCFrontLineServicePlanId, "cloudPCFrontLineServicePlanIdValue")
	}
}

func TestFormatDeviceManagementVirtualEndpointFrontLineServicePlanID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointFrontLineServicePlanID("cloudPCFrontLineServicePlanIdValue").ID()
	expected := "/deviceManagement/virtualEndpoint/frontLineServicePlans/cloudPCFrontLineServicePlanIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointFrontLineServicePlanID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointFrontLineServicePlanId
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
			Input: "/deviceManagement/virtualEndpoint/frontLineServicePlans",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/frontLineServicePlans/cloudPCFrontLineServicePlanIdValue",
			Expected: &DeviceManagementVirtualEndpointFrontLineServicePlanId{
				CloudPCFrontLineServicePlanId: "cloudPCFrontLineServicePlanIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/frontLineServicePlans/cloudPCFrontLineServicePlanIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointFrontLineServicePlanID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCFrontLineServicePlanId != v.Expected.CloudPCFrontLineServicePlanId {
			t.Fatalf("Expected %q but got %q for CloudPCFrontLineServicePlanId", v.Expected.CloudPCFrontLineServicePlanId, actual.CloudPCFrontLineServicePlanId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointFrontLineServicePlanIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointFrontLineServicePlanId
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
			Input: "/deviceManagement/virtualEndpoint/frontLineServicePlans",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/fRoNtLiNeSeRvIcEpLaNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/frontLineServicePlans/cloudPCFrontLineServicePlanIdValue",
			Expected: &DeviceManagementVirtualEndpointFrontLineServicePlanId{
				CloudPCFrontLineServicePlanId: "cloudPCFrontLineServicePlanIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/frontLineServicePlans/cloudPCFrontLineServicePlanIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/fRoNtLiNeSeRvIcEpLaNs/cLoUdPcFrOnTlInEsErViCePlAnIdVaLuE",
			Expected: &DeviceManagementVirtualEndpointFrontLineServicePlanId{
				CloudPCFrontLineServicePlanId: "cLoUdPcFrOnTlInEsErViCePlAnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/fRoNtLiNeSeRvIcEpLaNs/cLoUdPcFrOnTlInEsErViCePlAnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointFrontLineServicePlanIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCFrontLineServicePlanId != v.Expected.CloudPCFrontLineServicePlanId {
			t.Fatalf("Expected %q but got %q for CloudPCFrontLineServicePlanId", v.Expected.CloudPCFrontLineServicePlanId, actual.CloudPCFrontLineServicePlanId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointFrontLineServicePlanId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointFrontLineServicePlanId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointFrontLineServicePlanId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
