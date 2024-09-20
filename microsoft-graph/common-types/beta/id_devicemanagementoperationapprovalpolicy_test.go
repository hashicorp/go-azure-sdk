package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementOperationApprovalPolicyId{}

func TestNewDeviceManagementOperationApprovalPolicyID(t *testing.T) {
	id := NewDeviceManagementOperationApprovalPolicyID("operationApprovalPolicyId")

	if id.OperationApprovalPolicyId != "operationApprovalPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'OperationApprovalPolicyId'", id.OperationApprovalPolicyId, "operationApprovalPolicyId")
	}
}

func TestFormatDeviceManagementOperationApprovalPolicyID(t *testing.T) {
	actual := NewDeviceManagementOperationApprovalPolicyID("operationApprovalPolicyId").ID()
	expected := "/deviceManagement/operationApprovalPolicies/operationApprovalPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementOperationApprovalPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementOperationApprovalPolicyId
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
			Input: "/deviceManagement/operationApprovalPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/operationApprovalPolicies/operationApprovalPolicyId",
			Expected: &DeviceManagementOperationApprovalPolicyId{
				OperationApprovalPolicyId: "operationApprovalPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/operationApprovalPolicies/operationApprovalPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementOperationApprovalPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OperationApprovalPolicyId != v.Expected.OperationApprovalPolicyId {
			t.Fatalf("Expected %q but got %q for OperationApprovalPolicyId", v.Expected.OperationApprovalPolicyId, actual.OperationApprovalPolicyId)
		}

	}
}

func TestParseDeviceManagementOperationApprovalPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementOperationApprovalPolicyId
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
			Input: "/deviceManagement/operationApprovalPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/oPeRaTiOnApPrOvAlPoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/operationApprovalPolicies/operationApprovalPolicyId",
			Expected: &DeviceManagementOperationApprovalPolicyId{
				OperationApprovalPolicyId: "operationApprovalPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/operationApprovalPolicies/operationApprovalPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/oPeRaTiOnApPrOvAlPoLiCiEs/oPeRaTiOnApPrOvAlPoLiCyId",
			Expected: &DeviceManagementOperationApprovalPolicyId{
				OperationApprovalPolicyId: "oPeRaTiOnApPrOvAlPoLiCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/oPeRaTiOnApPrOvAlPoLiCiEs/oPeRaTiOnApPrOvAlPoLiCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementOperationApprovalPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OperationApprovalPolicyId != v.Expected.OperationApprovalPolicyId {
			t.Fatalf("Expected %q but got %q for OperationApprovalPolicyId", v.Expected.OperationApprovalPolicyId, actual.OperationApprovalPolicyId)
		}

	}
}

func TestSegmentsForDeviceManagementOperationApprovalPolicyId(t *testing.T) {
	segments := DeviceManagementOperationApprovalPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementOperationApprovalPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
