package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId{}

func TestNewDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentID("cloudPCProvisioningPolicyId", "cloudPCProvisioningPolicyAssignmentId")

	if id.CloudPCProvisioningPolicyId != "cloudPCProvisioningPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCProvisioningPolicyId'", id.CloudPCProvisioningPolicyId, "cloudPCProvisioningPolicyId")
	}

	if id.CloudPCProvisioningPolicyAssignmentId != "cloudPCProvisioningPolicyAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCProvisioningPolicyAssignmentId'", id.CloudPCProvisioningPolicyAssignmentId, "cloudPCProvisioningPolicyAssignmentId")
	}
}

func TestFormatDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentID("cloudPCProvisioningPolicyId", "cloudPCProvisioningPolicyAssignmentId").ID()
	expected := "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/assignments/cloudPCProvisioningPolicyAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId
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
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/assignments/cloudPCProvisioningPolicyAssignmentId",
			Expected: &DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId{
				CloudPCProvisioningPolicyId:           "cloudPCProvisioningPolicyId",
				CloudPCProvisioningPolicyAssignmentId: "cloudPCProvisioningPolicyAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/assignments/cloudPCProvisioningPolicyAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentID(v.Input)
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

		if actual.CloudPCProvisioningPolicyAssignmentId != v.Expected.CloudPCProvisioningPolicyAssignmentId {
			t.Fatalf("Expected %q but got %q for CloudPCProvisioningPolicyAssignmentId", v.Expected.CloudPCProvisioningPolicyAssignmentId, actual.CloudPCProvisioningPolicyAssignmentId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId
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
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/pRoViSiOnInGpOlIcIeS/cLoUdPcPrOvIsIoNiNgPoLiCyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/pRoViSiOnInGpOlIcIeS/cLoUdPcPrOvIsIoNiNgPoLiCyId/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/assignments/cloudPCProvisioningPolicyAssignmentId",
			Expected: &DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId{
				CloudPCProvisioningPolicyId:           "cloudPCProvisioningPolicyId",
				CloudPCProvisioningPolicyAssignmentId: "cloudPCProvisioningPolicyAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/assignments/cloudPCProvisioningPolicyAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/pRoViSiOnInGpOlIcIeS/cLoUdPcPrOvIsIoNiNgPoLiCyId/aSsIgNmEnTs/cLoUdPcPrOvIsIoNiNgPoLiCyAsSiGnMeNtId",
			Expected: &DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId{
				CloudPCProvisioningPolicyId:           "cLoUdPcPrOvIsIoNiNgPoLiCyId",
				CloudPCProvisioningPolicyAssignmentId: "cLoUdPcPrOvIsIoNiNgPoLiCyAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/pRoViSiOnInGpOlIcIeS/cLoUdPcPrOvIsIoNiNgPoLiCyId/aSsIgNmEnTs/cLoUdPcPrOvIsIoNiNgPoLiCyAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIDInsensitively(v.Input)
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

		if actual.CloudPCProvisioningPolicyAssignmentId != v.Expected.CloudPCProvisioningPolicyAssignmentId {
			t.Fatalf("Expected %q but got %q for CloudPCProvisioningPolicyAssignmentId", v.Expected.CloudPCProvisioningPolicyAssignmentId, actual.CloudPCProvisioningPolicyAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
