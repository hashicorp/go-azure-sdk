package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointProvisioningPolicyId{}

func TestNewDeviceManagementVirtualEndpointProvisioningPolicyID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointProvisioningPolicyID("cloudPCProvisioningPolicyId")

	if id.CloudPCProvisioningPolicyId != "cloudPCProvisioningPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCProvisioningPolicyId'", id.CloudPCProvisioningPolicyId, "cloudPCProvisioningPolicyId")
	}
}

func TestFormatDeviceManagementVirtualEndpointProvisioningPolicyID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointProvisioningPolicyID("cloudPCProvisioningPolicyId").ID()
	expected := "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointProvisioningPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointProvisioningPolicyId
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
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId",
			Expected: &DeviceManagementVirtualEndpointProvisioningPolicyId{
				CloudPCProvisioningPolicyId: "cloudPCProvisioningPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointProvisioningPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCProvisioningPolicyId != v.Expected.CloudPCProvisioningPolicyId {
			t.Fatalf("Expected %q but got %q for CloudPCProvisioningPolicyId", v.Expected.CloudPCProvisioningPolicyId, actual.CloudPCProvisioningPolicyId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointProvisioningPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointProvisioningPolicyId
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
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/pRoViSiOnInGpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId",
			Expected: &DeviceManagementVirtualEndpointProvisioningPolicyId{
				CloudPCProvisioningPolicyId: "cloudPCProvisioningPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/pRoViSiOnInGpOlIcIeS/cLoUdPcPrOvIsIoNiNgPoLiCyId",
			Expected: &DeviceManagementVirtualEndpointProvisioningPolicyId{
				CloudPCProvisioningPolicyId: "cLoUdPcPrOvIsIoNiNgPoLiCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/pRoViSiOnInGpOlIcIeS/cLoUdPcPrOvIsIoNiNgPoLiCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointProvisioningPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCProvisioningPolicyId != v.Expected.CloudPCProvisioningPolicyId {
			t.Fatalf("Expected %q but got %q for CloudPCProvisioningPolicyId", v.Expected.CloudPCProvisioningPolicyId, actual.CloudPCProvisioningPolicyId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointProvisioningPolicyId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointProvisioningPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointProvisioningPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
