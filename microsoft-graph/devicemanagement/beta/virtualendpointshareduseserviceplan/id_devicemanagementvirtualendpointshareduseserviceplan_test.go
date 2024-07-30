package virtualendpointshareduseserviceplan

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointSharedUseServicePlanId{}

func TestNewDeviceManagementVirtualEndpointSharedUseServicePlanID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointSharedUseServicePlanID("cloudPCSharedUseServicePlanIdValue")

	if id.CloudPCSharedUseServicePlanId != "cloudPCSharedUseServicePlanIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCSharedUseServicePlanId'", id.CloudPCSharedUseServicePlanId, "cloudPCSharedUseServicePlanIdValue")
	}
}

func TestFormatDeviceManagementVirtualEndpointSharedUseServicePlanID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointSharedUseServicePlanID("cloudPCSharedUseServicePlanIdValue").ID()
	expected := "/deviceManagement/virtualEndpoint/sharedUseServicePlans/cloudPCSharedUseServicePlanIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointSharedUseServicePlanID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointSharedUseServicePlanId
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
			Input: "/deviceManagement/virtualEndpoint/sharedUseServicePlans",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/sharedUseServicePlans/cloudPCSharedUseServicePlanIdValue",
			Expected: &DeviceManagementVirtualEndpointSharedUseServicePlanId{
				CloudPCSharedUseServicePlanId: "cloudPCSharedUseServicePlanIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/sharedUseServicePlans/cloudPCSharedUseServicePlanIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointSharedUseServicePlanID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCSharedUseServicePlanId != v.Expected.CloudPCSharedUseServicePlanId {
			t.Fatalf("Expected %q but got %q for CloudPCSharedUseServicePlanId", v.Expected.CloudPCSharedUseServicePlanId, actual.CloudPCSharedUseServicePlanId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointSharedUseServicePlanIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointSharedUseServicePlanId
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
			Input: "/deviceManagement/virtualEndpoint/sharedUseServicePlans",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/sHaReDuSeSeRvIcEpLaNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/sharedUseServicePlans/cloudPCSharedUseServicePlanIdValue",
			Expected: &DeviceManagementVirtualEndpointSharedUseServicePlanId{
				CloudPCSharedUseServicePlanId: "cloudPCSharedUseServicePlanIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/sharedUseServicePlans/cloudPCSharedUseServicePlanIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/sHaReDuSeSeRvIcEpLaNs/cLoUdPcShArEdUsEsErViCePlAnIdVaLuE",
			Expected: &DeviceManagementVirtualEndpointSharedUseServicePlanId{
				CloudPCSharedUseServicePlanId: "cLoUdPcShArEdUsEsErViCePlAnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/sHaReDuSeSeRvIcEpLaNs/cLoUdPcShArEdUsEsErViCePlAnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointSharedUseServicePlanIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCSharedUseServicePlanId != v.Expected.CloudPCSharedUseServicePlanId {
			t.Fatalf("Expected %q but got %q for CloudPCSharedUseServicePlanId", v.Expected.CloudPCSharedUseServicePlanId, actual.CloudPCSharedUseServicePlanId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointSharedUseServicePlanId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointSharedUseServicePlanId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointSharedUseServicePlanId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
