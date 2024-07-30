package manageddevicedeviceconfigurationstate

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceIdDeviceConfigurationStateId{}

func TestNewDeviceManagementManagedDeviceIdDeviceConfigurationStateID(t *testing.T) {
	id := NewDeviceManagementManagedDeviceIdDeviceConfigurationStateID("managedDeviceIdValue", "deviceConfigurationStateIdValue")

	if id.ManagedDeviceId != "managedDeviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceIdValue")
	}

	if id.DeviceConfigurationStateId != "deviceConfigurationStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationStateId'", id.DeviceConfigurationStateId, "deviceConfigurationStateIdValue")
	}
}

func TestFormatDeviceManagementManagedDeviceIdDeviceConfigurationStateID(t *testing.T) {
	actual := NewDeviceManagementManagedDeviceIdDeviceConfigurationStateID("managedDeviceIdValue", "deviceConfigurationStateIdValue").ID()
	expected := "/deviceManagement/managedDevices/managedDeviceIdValue/deviceConfigurationStates/deviceConfigurationStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementManagedDeviceIdDeviceConfigurationStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceIdDeviceConfigurationStateId
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
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/deviceConfigurationStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/deviceConfigurationStates/deviceConfigurationStateIdValue",
			Expected: &DeviceManagementManagedDeviceIdDeviceConfigurationStateId{
				ManagedDeviceId:            "managedDeviceIdValue",
				DeviceConfigurationStateId: "deviceConfigurationStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/deviceConfigurationStates/deviceConfigurationStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceIdDeviceConfigurationStateID(v.Input)
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

		if actual.DeviceConfigurationStateId != v.Expected.DeviceConfigurationStateId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationStateId", v.Expected.DeviceConfigurationStateId, actual.DeviceConfigurationStateId)
		}

	}
}

func TestParseDeviceManagementManagedDeviceIdDeviceConfigurationStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceIdDeviceConfigurationStateId
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
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/deviceConfigurationStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/dEvIcEcOnFiGuRaTiOnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/deviceConfigurationStates/deviceConfigurationStateIdValue",
			Expected: &DeviceManagementManagedDeviceIdDeviceConfigurationStateId{
				ManagedDeviceId:            "managedDeviceIdValue",
				DeviceConfigurationStateId: "deviceConfigurationStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/deviceConfigurationStates/deviceConfigurationStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/dEvIcEcOnFiGuRaTiOnStAtEs/dEvIcEcOnFiGuRaTiOnStAtEiDvAlUe",
			Expected: &DeviceManagementManagedDeviceIdDeviceConfigurationStateId{
				ManagedDeviceId:            "mAnAgEdDeViCeIdVaLuE",
				DeviceConfigurationStateId: "dEvIcEcOnFiGuRaTiOnStAtEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/dEvIcEcOnFiGuRaTiOnStAtEs/dEvIcEcOnFiGuRaTiOnStAtEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceIdDeviceConfigurationStateIDInsensitively(v.Input)
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

		if actual.DeviceConfigurationStateId != v.Expected.DeviceConfigurationStateId {
			t.Fatalf("Expected %q but got %q for DeviceConfigurationStateId", v.Expected.DeviceConfigurationStateId, actual.DeviceConfigurationStateId)
		}

	}
}

func TestSegmentsForDeviceManagementManagedDeviceIdDeviceConfigurationStateId(t *testing.T) {
	segments := DeviceManagementManagedDeviceIdDeviceConfigurationStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementManagedDeviceIdDeviceConfigurationStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
