package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicyIdDeviceStatusId{}

func TestNewDeviceManagementDeviceCompliancePolicyIdDeviceStatusID(t *testing.T) {
	id := NewDeviceManagementDeviceCompliancePolicyIdDeviceStatusID("deviceCompliancePolicyId", "deviceComplianceDeviceStatusId")

	if id.DeviceCompliancePolicyId != "deviceCompliancePolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCompliancePolicyId'", id.DeviceCompliancePolicyId, "deviceCompliancePolicyId")
	}

	if id.DeviceComplianceDeviceStatusId != "deviceComplianceDeviceStatusId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceComplianceDeviceStatusId'", id.DeviceComplianceDeviceStatusId, "deviceComplianceDeviceStatusId")
	}
}

func TestFormatDeviceManagementDeviceCompliancePolicyIdDeviceStatusID(t *testing.T) {
	actual := NewDeviceManagementDeviceCompliancePolicyIdDeviceStatusID("deviceCompliancePolicyId", "deviceComplianceDeviceStatusId").ID()
	expected := "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/deviceStatuses/deviceComplianceDeviceStatusId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceCompliancePolicyIdDeviceStatusID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicyIdDeviceStatusId
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
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/deviceStatuses",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/deviceStatuses/deviceComplianceDeviceStatusId",
			Expected: &DeviceManagementDeviceCompliancePolicyIdDeviceStatusId{
				DeviceCompliancePolicyId:       "deviceCompliancePolicyId",
				DeviceComplianceDeviceStatusId: "deviceComplianceDeviceStatusId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/deviceStatuses/deviceComplianceDeviceStatusId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicyIdDeviceStatusID(v.Input)
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

		if actual.DeviceComplianceDeviceStatusId != v.Expected.DeviceComplianceDeviceStatusId {
			t.Fatalf("Expected %q but got %q for DeviceComplianceDeviceStatusId", v.Expected.DeviceComplianceDeviceStatusId, actual.DeviceComplianceDeviceStatusId)
		}

	}
}

func TestParseDeviceManagementDeviceCompliancePolicyIdDeviceStatusIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicyIdDeviceStatusId
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
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/deviceStatuses",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD/dEvIcEsTaTuSeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/deviceStatuses/deviceComplianceDeviceStatusId",
			Expected: &DeviceManagementDeviceCompliancePolicyIdDeviceStatusId{
				DeviceCompliancePolicyId:       "deviceCompliancePolicyId",
				DeviceComplianceDeviceStatusId: "deviceComplianceDeviceStatusId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/deviceStatuses/deviceComplianceDeviceStatusId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD/dEvIcEsTaTuSeS/dEvIcEcOmPlIaNcEdEvIcEsTaTuSiD",
			Expected: &DeviceManagementDeviceCompliancePolicyIdDeviceStatusId{
				DeviceCompliancePolicyId:       "dEvIcEcOmPlIaNcEpOlIcYiD",
				DeviceComplianceDeviceStatusId: "dEvIcEcOmPlIaNcEdEvIcEsTaTuSiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD/dEvIcEsTaTuSeS/dEvIcEcOmPlIaNcEdEvIcEsTaTuSiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicyIdDeviceStatusIDInsensitively(v.Input)
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

		if actual.DeviceComplianceDeviceStatusId != v.Expected.DeviceComplianceDeviceStatusId {
			t.Fatalf("Expected %q but got %q for DeviceComplianceDeviceStatusId", v.Expected.DeviceComplianceDeviceStatusId, actual.DeviceComplianceDeviceStatusId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceCompliancePolicyIdDeviceStatusId(t *testing.T) {
	segments := DeviceManagementDeviceCompliancePolicyIdDeviceStatusId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceCompliancePolicyIdDeviceStatusId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
