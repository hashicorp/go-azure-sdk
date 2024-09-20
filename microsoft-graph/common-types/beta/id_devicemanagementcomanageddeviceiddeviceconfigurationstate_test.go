package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagedDeviceIdDeviceConfigurationStateId{}

func TestNewDeviceManagementComanagedDeviceIdDeviceConfigurationStateID(t *testing.T) {
	id := NewDeviceManagementComanagedDeviceIdDeviceConfigurationStateID("managedDeviceId", "deviceConfigurationStateId")

	if id.ManagedDeviceId != "managedDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceId")
	}

	if id.DeviceConfigurationStateId != "deviceConfigurationStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationStateId'", id.DeviceConfigurationStateId, "deviceConfigurationStateId")
	}
}

func TestFormatDeviceManagementComanagedDeviceIdDeviceConfigurationStateID(t *testing.T) {
	actual := NewDeviceManagementComanagedDeviceIdDeviceConfigurationStateID("managedDeviceId", "deviceConfigurationStateId").ID()
	expected := "/deviceManagement/comanagedDevices/managedDeviceId/deviceConfigurationStates/deviceConfigurationStateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementComanagedDeviceIdDeviceConfigurationStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementComanagedDeviceIdDeviceConfigurationStateId
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
			Input: "/deviceManagement/comanagedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/deviceConfigurationStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/deviceConfigurationStates/deviceConfigurationStateId",
			Expected: &DeviceManagementComanagedDeviceIdDeviceConfigurationStateId{
				ManagedDeviceId:            "managedDeviceId",
				DeviceConfigurationStateId: "deviceConfigurationStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/deviceConfigurationStates/deviceConfigurationStateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementComanagedDeviceIdDeviceConfigurationStateID(v.Input)
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

func TestParseDeviceManagementComanagedDeviceIdDeviceConfigurationStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementComanagedDeviceIdDeviceConfigurationStateId
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
			Input: "/deviceManagement/comanagedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/deviceConfigurationStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId/dEvIcEcOnFiGuRaTiOnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/deviceConfigurationStates/deviceConfigurationStateId",
			Expected: &DeviceManagementComanagedDeviceIdDeviceConfigurationStateId{
				ManagedDeviceId:            "managedDeviceId",
				DeviceConfigurationStateId: "deviceConfigurationStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/deviceConfigurationStates/deviceConfigurationStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId/dEvIcEcOnFiGuRaTiOnStAtEs/dEvIcEcOnFiGuRaTiOnStAtEiD",
			Expected: &DeviceManagementComanagedDeviceIdDeviceConfigurationStateId{
				ManagedDeviceId:            "mAnAgEdDeViCeId",
				DeviceConfigurationStateId: "dEvIcEcOnFiGuRaTiOnStAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId/dEvIcEcOnFiGuRaTiOnStAtEs/dEvIcEcOnFiGuRaTiOnStAtEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementComanagedDeviceIdDeviceConfigurationStateIDInsensitively(v.Input)
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

func TestSegmentsForDeviceManagementComanagedDeviceIdDeviceConfigurationStateId(t *testing.T) {
	segments := DeviceManagementComanagedDeviceIdDeviceConfigurationStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementComanagedDeviceIdDeviceConfigurationStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
