package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicyIdAssignmentId{}

func TestNewDeviceManagementDeviceCompliancePolicyIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementDeviceCompliancePolicyIdAssignmentID("deviceCompliancePolicyId", "deviceCompliancePolicyAssignmentId")

	if id.DeviceCompliancePolicyId != "deviceCompliancePolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCompliancePolicyId'", id.DeviceCompliancePolicyId, "deviceCompliancePolicyId")
	}

	if id.DeviceCompliancePolicyAssignmentId != "deviceCompliancePolicyAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCompliancePolicyAssignmentId'", id.DeviceCompliancePolicyAssignmentId, "deviceCompliancePolicyAssignmentId")
	}
}

func TestFormatDeviceManagementDeviceCompliancePolicyIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementDeviceCompliancePolicyIdAssignmentID("deviceCompliancePolicyId", "deviceCompliancePolicyAssignmentId").ID()
	expected := "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/assignments/deviceCompliancePolicyAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceCompliancePolicyIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicyIdAssignmentId
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
			Input: "/deviceManagement/deviceCompliancePolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/assignments/deviceCompliancePolicyAssignmentId",
			Expected: &DeviceManagementDeviceCompliancePolicyIdAssignmentId{
				DeviceCompliancePolicyId:           "deviceCompliancePolicyId",
				DeviceCompliancePolicyAssignmentId: "deviceCompliancePolicyAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/assignments/deviceCompliancePolicyAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicyIdAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceCompliancePolicyId != v.Expected.DeviceCompliancePolicyId {
			t.Fatalf("Expected %q but got %q for DeviceCompliancePolicyId", v.Expected.DeviceCompliancePolicyId, actual.DeviceCompliancePolicyId)
		}

		if actual.DeviceCompliancePolicyAssignmentId != v.Expected.DeviceCompliancePolicyAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceCompliancePolicyAssignmentId", v.Expected.DeviceCompliancePolicyAssignmentId, actual.DeviceCompliancePolicyAssignmentId)
		}

	}
}

func TestParseDeviceManagementDeviceCompliancePolicyIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicyIdAssignmentId
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
			Input: "/deviceManagement/deviceCompliancePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/assignments/deviceCompliancePolicyAssignmentId",
			Expected: &DeviceManagementDeviceCompliancePolicyIdAssignmentId{
				DeviceCompliancePolicyId:           "deviceCompliancePolicyId",
				DeviceCompliancePolicyAssignmentId: "deviceCompliancePolicyAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/assignments/deviceCompliancePolicyAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD/aSsIgNmEnTs/dEvIcEcOmPlIaNcEpOlIcYaSsIgNmEnTiD",
			Expected: &DeviceManagementDeviceCompliancePolicyIdAssignmentId{
				DeviceCompliancePolicyId:           "dEvIcEcOmPlIaNcEpOlIcYiD",
				DeviceCompliancePolicyAssignmentId: "dEvIcEcOmPlIaNcEpOlIcYaSsIgNmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD/aSsIgNmEnTs/dEvIcEcOmPlIaNcEpOlIcYaSsIgNmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicyIdAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceCompliancePolicyId != v.Expected.DeviceCompliancePolicyId {
			t.Fatalf("Expected %q but got %q for DeviceCompliancePolicyId", v.Expected.DeviceCompliancePolicyId, actual.DeviceCompliancePolicyId)
		}

		if actual.DeviceCompliancePolicyAssignmentId != v.Expected.DeviceCompliancePolicyAssignmentId {
			t.Fatalf("Expected %q but got %q for DeviceCompliancePolicyAssignmentId", v.Expected.DeviceCompliancePolicyAssignmentId, actual.DeviceCompliancePolicyAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceCompliancePolicyIdAssignmentId(t *testing.T) {
	segments := DeviceManagementDeviceCompliancePolicyIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceCompliancePolicyIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
