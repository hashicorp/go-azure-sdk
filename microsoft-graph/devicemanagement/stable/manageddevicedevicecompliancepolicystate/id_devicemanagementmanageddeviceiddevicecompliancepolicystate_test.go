package manageddevicedevicecompliancepolicystate

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId{}

func TestNewDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateID(t *testing.T) {
	id := NewDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateID("managedDeviceIdValue", "deviceCompliancePolicyStateIdValue")

	if id.ManagedDeviceId != "managedDeviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceIdValue")
	}

	if id.DeviceCompliancePolicyStateId != "deviceCompliancePolicyStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCompliancePolicyStateId'", id.DeviceCompliancePolicyStateId, "deviceCompliancePolicyStateIdValue")
	}
}

func TestFormatDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateID(t *testing.T) {
	actual := NewDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateID("managedDeviceIdValue", "deviceCompliancePolicyStateIdValue").ID()
	expected := "/deviceManagement/managedDevices/managedDeviceIdValue/deviceCompliancePolicyStates/deviceCompliancePolicyStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId
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
			Input: "/deviceManagement/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/deviceCompliancePolicyStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/deviceCompliancePolicyStates/deviceCompliancePolicyStateIdValue",
			Expected: &DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId{
				ManagedDeviceId:               "managedDeviceIdValue",
				DeviceCompliancePolicyStateId: "deviceCompliancePolicyStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/deviceCompliancePolicyStates/deviceCompliancePolicyStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.DeviceCompliancePolicyStateId != v.Expected.DeviceCompliancePolicyStateId {
			t.Fatalf("Expected %q but got %q for DeviceCompliancePolicyStateId", v.Expected.DeviceCompliancePolicyStateId, actual.DeviceCompliancePolicyStateId)
		}

	}
}

func TestParseDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId
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
			Input: "/deviceManagement/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/deviceCompliancePolicyStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/dEvIcEcOmPlIaNcEpOlIcYsTaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/deviceCompliancePolicyStates/deviceCompliancePolicyStateIdValue",
			Expected: &DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId{
				ManagedDeviceId:               "managedDeviceIdValue",
				DeviceCompliancePolicyStateId: "deviceCompliancePolicyStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/deviceCompliancePolicyStates/deviceCompliancePolicyStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/dEvIcEcOmPlIaNcEpOlIcYsTaTeS/dEvIcEcOmPlIaNcEpOlIcYsTaTeIdVaLuE",
			Expected: &DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId{
				ManagedDeviceId:               "mAnAgEdDeViCeIdVaLuE",
				DeviceCompliancePolicyStateId: "dEvIcEcOmPlIaNcEpOlIcYsTaTeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/dEvIcEcOmPlIaNcEpOlIcYsTaTeS/dEvIcEcOmPlIaNcEpOlIcYsTaTeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.DeviceCompliancePolicyStateId != v.Expected.DeviceCompliancePolicyStateId {
			t.Fatalf("Expected %q but got %q for DeviceCompliancePolicyStateId", v.Expected.DeviceCompliancePolicyStateId, actual.DeviceCompliancePolicyStateId)
		}

	}
}

func TestSegmentsForDeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId(t *testing.T) {
	segments := DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementManagedDeviceIdDeviceCompliancePolicyStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
