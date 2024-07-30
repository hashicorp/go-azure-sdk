package devicecompliancepolicy

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicyId{}

func TestNewDeviceManagementDeviceCompliancePolicyID(t *testing.T) {
	id := NewDeviceManagementDeviceCompliancePolicyID("deviceCompliancePolicyIdValue")

	if id.DeviceCompliancePolicyId != "deviceCompliancePolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCompliancePolicyId'", id.DeviceCompliancePolicyId, "deviceCompliancePolicyIdValue")
	}
}

func TestFormatDeviceManagementDeviceCompliancePolicyID(t *testing.T) {
	actual := NewDeviceManagementDeviceCompliancePolicyID("deviceCompliancePolicyIdValue").ID()
	expected := "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceCompliancePolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicyId
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
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue",
			Expected: &DeviceManagementDeviceCompliancePolicyId{
				DeviceCompliancePolicyId: "deviceCompliancePolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicyID(v.Input)
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

	}
}

func TestParseDeviceManagementDeviceCompliancePolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicyId
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
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue",
			Expected: &DeviceManagementDeviceCompliancePolicyId{
				DeviceCompliancePolicyId: "deviceCompliancePolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiDvAlUe",
			Expected: &DeviceManagementDeviceCompliancePolicyId{
				DeviceCompliancePolicyId: "dEvIcEcOmPlIaNcEpOlIcYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicyIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementDeviceCompliancePolicyId(t *testing.T) {
	segments := DeviceManagementDeviceCompliancePolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceCompliancePolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
