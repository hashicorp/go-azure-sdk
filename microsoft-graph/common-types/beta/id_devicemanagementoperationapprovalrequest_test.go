package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementOperationApprovalRequestId{}

func TestNewDeviceManagementOperationApprovalRequestID(t *testing.T) {
	id := NewDeviceManagementOperationApprovalRequestID("operationApprovalRequestIdValue")

	if id.OperationApprovalRequestId != "operationApprovalRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OperationApprovalRequestId'", id.OperationApprovalRequestId, "operationApprovalRequestIdValue")
	}
}

func TestFormatDeviceManagementOperationApprovalRequestID(t *testing.T) {
	actual := NewDeviceManagementOperationApprovalRequestID("operationApprovalRequestIdValue").ID()
	expected := "/deviceManagement/operationApprovalRequests/operationApprovalRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementOperationApprovalRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementOperationApprovalRequestId
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
			Input: "/deviceManagement/operationApprovalRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/operationApprovalRequests/operationApprovalRequestIdValue",
			Expected: &DeviceManagementOperationApprovalRequestId{
				OperationApprovalRequestId: "operationApprovalRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/operationApprovalRequests/operationApprovalRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementOperationApprovalRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OperationApprovalRequestId != v.Expected.OperationApprovalRequestId {
			t.Fatalf("Expected %q but got %q for OperationApprovalRequestId", v.Expected.OperationApprovalRequestId, actual.OperationApprovalRequestId)
		}

	}
}

func TestParseDeviceManagementOperationApprovalRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementOperationApprovalRequestId
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
			Input: "/deviceManagement/operationApprovalRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/oPeRaTiOnApPrOvAlReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/operationApprovalRequests/operationApprovalRequestIdValue",
			Expected: &DeviceManagementOperationApprovalRequestId{
				OperationApprovalRequestId: "operationApprovalRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/operationApprovalRequests/operationApprovalRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/oPeRaTiOnApPrOvAlReQuEsTs/oPeRaTiOnApPrOvAlReQuEsTiDvAlUe",
			Expected: &DeviceManagementOperationApprovalRequestId{
				OperationApprovalRequestId: "oPeRaTiOnApPrOvAlReQuEsTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/oPeRaTiOnApPrOvAlReQuEsTs/oPeRaTiOnApPrOvAlReQuEsTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementOperationApprovalRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OperationApprovalRequestId != v.Expected.OperationApprovalRequestId {
			t.Fatalf("Expected %q but got %q for OperationApprovalRequestId", v.Expected.OperationApprovalRequestId, actual.OperationApprovalRequestId)
		}

	}
}

func TestSegmentsForDeviceManagementOperationApprovalRequestId(t *testing.T) {
	segments := DeviceManagementOperationApprovalRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementOperationApprovalRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
