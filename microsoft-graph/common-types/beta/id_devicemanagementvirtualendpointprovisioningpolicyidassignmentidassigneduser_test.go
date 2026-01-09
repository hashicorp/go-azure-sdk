package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId{}

func TestNewDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserID(t *testing.T) {
	id := NewDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserID("cloudPCProvisioningPolicyId", "cloudPCProvisioningPolicyAssignmentId", "userId")

	if id.CloudPCProvisioningPolicyId != "cloudPCProvisioningPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCProvisioningPolicyId'", id.CloudPCProvisioningPolicyId, "cloudPCProvisioningPolicyId")
	}

	if id.CloudPCProvisioningPolicyAssignmentId != "cloudPCProvisioningPolicyAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCProvisioningPolicyAssignmentId'", id.CloudPCProvisioningPolicyAssignmentId, "cloudPCProvisioningPolicyAssignmentId")
	}

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}
}

func TestFormatDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserID(t *testing.T) {
	actual := NewDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserID("cloudPCProvisioningPolicyId", "cloudPCProvisioningPolicyAssignmentId", "userId").ID()
	expected := "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/assignments/cloudPCProvisioningPolicyAssignmentId/assignedUsers/userId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId
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
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/assignments/cloudPCProvisioningPolicyAssignmentId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/assignments/cloudPCProvisioningPolicyAssignmentId/assignedUsers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/assignments/cloudPCProvisioningPolicyAssignmentId/assignedUsers/userId",
			Expected: &DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId{
				CloudPCProvisioningPolicyId:           "cloudPCProvisioningPolicyId",
				CloudPCProvisioningPolicyAssignmentId: "cloudPCProvisioningPolicyAssignmentId",
				UserId:                                "userId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/assignments/cloudPCProvisioningPolicyAssignmentId/assignedUsers/userId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserID(v.Input)
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

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

	}
}

func TestParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId
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
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/assignments/cloudPCProvisioningPolicyAssignmentId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/pRoViSiOnInGpOlIcIeS/cLoUdPcPrOvIsIoNiNgPoLiCyId/aSsIgNmEnTs/cLoUdPcPrOvIsIoNiNgPoLiCyAsSiGnMeNtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/assignments/cloudPCProvisioningPolicyAssignmentId/assignedUsers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/pRoViSiOnInGpOlIcIeS/cLoUdPcPrOvIsIoNiNgPoLiCyId/aSsIgNmEnTs/cLoUdPcPrOvIsIoNiNgPoLiCyAsSiGnMeNtId/aSsIgNeDuSeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/assignments/cloudPCProvisioningPolicyAssignmentId/assignedUsers/userId",
			Expected: &DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId{
				CloudPCProvisioningPolicyId:           "cloudPCProvisioningPolicyId",
				CloudPCProvisioningPolicyAssignmentId: "cloudPCProvisioningPolicyAssignmentId",
				UserId:                                "userId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/virtualEndpoint/provisioningPolicies/cloudPCProvisioningPolicyId/assignments/cloudPCProvisioningPolicyAssignmentId/assignedUsers/userId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/pRoViSiOnInGpOlIcIeS/cLoUdPcPrOvIsIoNiNgPoLiCyId/aSsIgNmEnTs/cLoUdPcPrOvIsIoNiNgPoLiCyAsSiGnMeNtId/aSsIgNeDuSeRs/uSeRiD",
			Expected: &DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId{
				CloudPCProvisioningPolicyId:           "cLoUdPcPrOvIsIoNiNgPoLiCyId",
				CloudPCProvisioningPolicyAssignmentId: "cLoUdPcPrOvIsIoNiNgPoLiCyAsSiGnMeNtId",
				UserId:                                "uSeRiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/vIrTuAlEnDpOiNt/pRoViSiOnInGpOlIcIeS/cLoUdPcPrOvIsIoNiNgPoLiCyId/aSsIgNmEnTs/cLoUdPcPrOvIsIoNiNgPoLiCyAsSiGnMeNtId/aSsIgNeDuSeRs/uSeRiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserIDInsensitively(v.Input)
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

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

	}
}

func TestSegmentsForDeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId(t *testing.T) {
	segments := DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementVirtualEndpointProvisioningPolicyIdAssignmentIdAssignedUserId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
