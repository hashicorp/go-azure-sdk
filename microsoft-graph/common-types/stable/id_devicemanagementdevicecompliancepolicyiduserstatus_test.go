package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCompliancePolicyIdUserStatusId{}

func TestNewDeviceManagementDeviceCompliancePolicyIdUserStatusID(t *testing.T) {
	id := NewDeviceManagementDeviceCompliancePolicyIdUserStatusID("deviceCompliancePolicyIdValue", "deviceComplianceUserStatusIdValue")

	if id.DeviceCompliancePolicyId != "deviceCompliancePolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCompliancePolicyId'", id.DeviceCompliancePolicyId, "deviceCompliancePolicyIdValue")
	}

	if id.DeviceComplianceUserStatusId != "deviceComplianceUserStatusIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceComplianceUserStatusId'", id.DeviceComplianceUserStatusId, "deviceComplianceUserStatusIdValue")
	}
}

func TestFormatDeviceManagementDeviceCompliancePolicyIdUserStatusID(t *testing.T) {
	actual := NewDeviceManagementDeviceCompliancePolicyIdUserStatusID("deviceCompliancePolicyIdValue", "deviceComplianceUserStatusIdValue").ID()
	expected := "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/userStatuses/deviceComplianceUserStatusIdValue"
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
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/userStatuses",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/userStatuses/deviceComplianceUserStatusIdValue",
			Expected: &DeviceManagementDeviceCompliancePolicyIdUserStatusId{
				DeviceCompliancePolicyId:     "deviceCompliancePolicyIdValue",
				DeviceComplianceUserStatusId: "deviceComplianceUserStatusIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/userStatuses/deviceComplianceUserStatusIdValue/extra",
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
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/userStatuses",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiDvAlUe/uSeRsTaTuSeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/userStatuses/deviceComplianceUserStatusIdValue",
			Expected: &DeviceManagementDeviceCompliancePolicyIdUserStatusId{
				DeviceCompliancePolicyId:     "deviceCompliancePolicyIdValue",
				DeviceComplianceUserStatusId: "deviceComplianceUserStatusIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCompliancePolicies/deviceCompliancePolicyIdValue/userStatuses/deviceComplianceUserStatusIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiDvAlUe/uSeRsTaTuSeS/dEvIcEcOmPlIaNcEuSeRsTaTuSiDvAlUe",
			Expected: &DeviceManagementDeviceCompliancePolicyIdUserStatusId{
				DeviceCompliancePolicyId:     "dEvIcEcOmPlIaNcEpOlIcYiDvAlUe",
				DeviceComplianceUserStatusId: "dEvIcEcOmPlIaNcEuSeRsTaTuSiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOmPlIaNcEpOlIcIeS/dEvIcEcOmPlIaNcEpOlIcYiDvAlUe/uSeRsTaTuSeS/dEvIcEcOmPlIaNcEuSeRsTaTuSiDvAlUe/extra",
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
