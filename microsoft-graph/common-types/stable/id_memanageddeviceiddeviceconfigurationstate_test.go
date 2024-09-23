package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeManagedDeviceIdDeviceConfigurationStateId{}

func TestNewMeManagedDeviceIdDeviceConfigurationStateID(t *testing.T) {
	id := NewMeManagedDeviceIdDeviceConfigurationStateID("managedDeviceId", "deviceConfigurationStateId")

	if id.ManagedDeviceId != "managedDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceId")
	}

	if id.DeviceConfigurationStateId != "deviceConfigurationStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceConfigurationStateId'", id.DeviceConfigurationStateId, "deviceConfigurationStateId")
	}
}

func TestFormatMeManagedDeviceIdDeviceConfigurationStateID(t *testing.T) {
	actual := NewMeManagedDeviceIdDeviceConfigurationStateID("managedDeviceId", "deviceConfigurationStateId").ID()
	expected := "/me/managedDevices/managedDeviceId/deviceConfigurationStates/deviceConfigurationStateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeManagedDeviceIdDeviceConfigurationStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeManagedDeviceIdDeviceConfigurationStateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceId/deviceConfigurationStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/managedDevices/managedDeviceId/deviceConfigurationStates/deviceConfigurationStateId",
			Expected: &MeManagedDeviceIdDeviceConfigurationStateId{
				ManagedDeviceId:            "managedDeviceId",
				DeviceConfigurationStateId: "deviceConfigurationStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/managedDevices/managedDeviceId/deviceConfigurationStates/deviceConfigurationStateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeManagedDeviceIdDeviceConfigurationStateID(v.Input)
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

func TestParseMeManagedDeviceIdDeviceConfigurationStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeManagedDeviceIdDeviceConfigurationStateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceId/deviceConfigurationStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeId/dEvIcEcOnFiGuRaTiOnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/managedDevices/managedDeviceId/deviceConfigurationStates/deviceConfigurationStateId",
			Expected: &MeManagedDeviceIdDeviceConfigurationStateId{
				ManagedDeviceId:            "managedDeviceId",
				DeviceConfigurationStateId: "deviceConfigurationStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/managedDevices/managedDeviceId/deviceConfigurationStates/deviceConfigurationStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeId/dEvIcEcOnFiGuRaTiOnStAtEs/dEvIcEcOnFiGuRaTiOnStAtEiD",
			Expected: &MeManagedDeviceIdDeviceConfigurationStateId{
				ManagedDeviceId:            "mAnAgEdDeViCeId",
				DeviceConfigurationStateId: "dEvIcEcOnFiGuRaTiOnStAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeId/dEvIcEcOnFiGuRaTiOnStAtEs/dEvIcEcOnFiGuRaTiOnStAtEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeManagedDeviceIdDeviceConfigurationStateIDInsensitively(v.Input)
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

func TestSegmentsForMeManagedDeviceIdDeviceConfigurationStateId(t *testing.T) {
	segments := MeManagedDeviceIdDeviceConfigurationStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeManagedDeviceIdDeviceConfigurationStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
