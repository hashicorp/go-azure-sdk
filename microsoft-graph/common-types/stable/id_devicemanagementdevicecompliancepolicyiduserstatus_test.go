package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicyIdUserStatusId{}

func TestNewDeviceManagementDeviceCompliancePolicyIdUserStatusID(t *testing.T) {
	id := NewDeviceManagementDeviceCompliancePolicyIdUserStatusID("deviceCompliancePolicyId", "deviceComplianceUserStatusId")

	if id.DeviceCompliancePolicyId != "deviceCompliancePolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCompliancePolicyId'", id.DeviceCompliancePolicyId, "deviceCompliancePolicyId")
	}

	if id.DeviceComplianceUserStatusId != "deviceComplianceUserStatusId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceComplianceUserStatusId'", id.DeviceComplianceUserStatusId, "deviceComplianceUserStatusId")
	}
}

func TestFormatDeviceManagementDeviceCompliancePolicyIdUserStatusID(t *testing.T) {
	actual := NewDeviceManagementDeviceCompliancePolicyIdUserStatusID("deviceCompliancePolicyId", "deviceComplianceUserStatusId").ID()
	expected := "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/userStatuses/deviceComplianceUserStatusId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceCompliancePolicyIdUserStatusID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicyIdUserStatusId
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
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/userStatuses",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/userStatuses/deviceComplianceUserStatusId",
			Expected: &DeviceManagementDeviceCompliancePolicyIdUserStatusId{
				DeviceCompliancePolicyId:     "deviceCompliancePolicyId",
				DeviceComplianceUserStatusId: "deviceComplianceUserStatusId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/userStatuses/deviceComplianceUserStatusId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicyIdUserStatusID(v.Input)
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

		if actual.DeviceComplianceUserStatusId != v.Expected.DeviceComplianceUserStatusId {
			t.Fatalf("Expected %q but got %q for DeviceComplianceUserStatusId", v.Expected.DeviceComplianceUserStatusId, actual.DeviceComplianceUserStatusId)
		}

	}
}

func TestParseDeviceManagementDeviceCompliancePolicyIdUserStatusIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCompliancePolicyIdUserStatusId
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
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/userStatuses",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD/uSeRsTaTuSeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/userStatuses/deviceComplianceUserStatusId",
			Expected: &DeviceManagementDeviceCompliancePolicyIdUserStatusId{
				DeviceCompliancePolicyId:     "deviceCompliancePolicyId",
				DeviceComplianceUserStatusId: "deviceComplianceUserStatusId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyId/userStatuses/deviceComplianceUserStatusId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD/uSeRsTaTuSeS/dEvIcEcOmPlIaNcEuSeRsTaTuSiD",
			Expected: &DeviceManagementDeviceCompliancePolicyIdUserStatusId{
				DeviceCompliancePolicyId:     "dEvIcEcOmPlIaNcEpOlIcYiD",
				DeviceComplianceUserStatusId: "dEvIcEcOmPlIaNcEuSeRsTaTuSiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiD/uSeRsTaTuSeS/dEvIcEcOmPlIaNcEuSeRsTaTuSiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCompliancePolicyIdUserStatusIDInsensitively(v.Input)
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

		if actual.DeviceComplianceUserStatusId != v.Expected.DeviceComplianceUserStatusId {
			t.Fatalf("Expected %q but got %q for DeviceComplianceUserStatusId", v.Expected.DeviceComplianceUserStatusId, actual.DeviceComplianceUserStatusId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceCompliancePolicyIdUserStatusId(t *testing.T) {
	segments := DeviceManagementDeviceCompliancePolicyIdUserStatusId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceCompliancePolicyIdUserStatusId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
