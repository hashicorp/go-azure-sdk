package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCompliancePolicyId{}

func TestNewDeviceManagementCompliancePolicyID(t *testing.T) {
	id := NewDeviceManagementCompliancePolicyID("deviceManagementCompliancePolicyId")

	if id.DeviceManagementCompliancePolicyId != "deviceManagementCompliancePolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementCompliancePolicyId'", id.DeviceManagementCompliancePolicyId, "deviceManagementCompliancePolicyId")
	}
}

func TestFormatDeviceManagementCompliancePolicyID(t *testing.T) {
	actual := NewDeviceManagementCompliancePolicyID("deviceManagementCompliancePolicyId").ID()
	expected := "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementCompliancePolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCompliancePolicyId
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
			Input: "/deviceManagement/compliancePolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId",
			Expected: &DeviceManagementCompliancePolicyId{
				DeviceManagementCompliancePolicyId: "deviceManagementCompliancePolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCompliancePolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementCompliancePolicyId != v.Expected.DeviceManagementCompliancePolicyId {
			t.Fatalf("Expected %q but got %q for DeviceManagementCompliancePolicyId", v.Expected.DeviceManagementCompliancePolicyId, actual.DeviceManagementCompliancePolicyId)
		}

	}
}

func TestParseDeviceManagementCompliancePolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCompliancePolicyId
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
			Input: "/deviceManagement/compliancePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId",
			Expected: &DeviceManagementCompliancePolicyId{
				DeviceManagementCompliancePolicyId: "deviceManagementCompliancePolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/compliancePolicies/deviceManagementCompliancePolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiD",
			Expected: &DeviceManagementCompliancePolicyId{
				DeviceManagementCompliancePolicyId: "dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEpOlIcIeS/dEvIcEmAnAgEmEnTcOmPlIaNcEpOlIcYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCompliancePolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementCompliancePolicyId != v.Expected.DeviceManagementCompliancePolicyId {
			t.Fatalf("Expected %q but got %q for DeviceManagementCompliancePolicyId", v.Expected.DeviceManagementCompliancePolicyId, actual.DeviceManagementCompliancePolicyId)
		}

	}
}

func TestSegmentsForDeviceManagementCompliancePolicyId(t *testing.T) {
	segments := DeviceManagementCompliancePolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementCompliancePolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
