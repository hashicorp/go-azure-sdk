package manageddevicemanageddevicemobileappconfigurationstate

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{}

func TestNewDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateID(t *testing.T) {
	id := NewDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateID("managedDeviceIdValue", "managedDeviceMobileAppConfigurationStateIdValue")

	if id.ManagedDeviceId != "managedDeviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceIdValue")
	}

	if id.ManagedDeviceMobileAppConfigurationStateId != "managedDeviceMobileAppConfigurationStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceMobileAppConfigurationStateId'", id.ManagedDeviceMobileAppConfigurationStateId, "managedDeviceMobileAppConfigurationStateIdValue")
	}
}

func TestFormatDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateID(t *testing.T) {
	actual := NewDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateID("managedDeviceIdValue", "managedDeviceMobileAppConfigurationStateIdValue").ID()
	expected := "/deviceManagement/managedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates/managedDeviceMobileAppConfigurationStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId
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
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates/managedDeviceMobileAppConfigurationStateIdValue",
			Expected: &DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{
				ManagedDeviceId: "managedDeviceIdValue",
				ManagedDeviceMobileAppConfigurationStateId: "managedDeviceMobileAppConfigurationStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates/managedDeviceMobileAppConfigurationStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateID(v.Input)
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

		if actual.ManagedDeviceMobileAppConfigurationStateId != v.Expected.ManagedDeviceMobileAppConfigurationStateId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceMobileAppConfigurationStateId", v.Expected.ManagedDeviceMobileAppConfigurationStateId, actual.ManagedDeviceMobileAppConfigurationStateId)
		}

	}
}

func TestParseDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId
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
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates/managedDeviceMobileAppConfigurationStateIdValue",
			Expected: &DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{
				ManagedDeviceId: "managedDeviceIdValue",
				ManagedDeviceMobileAppConfigurationStateId: "managedDeviceMobileAppConfigurationStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates/managedDeviceMobileAppConfigurationStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEs/mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEiDvAlUe",
			Expected: &DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{
				ManagedDeviceId: "mAnAgEdDeViCeIdVaLuE",
				ManagedDeviceMobileAppConfigurationStateId: "mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEs/mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateIDInsensitively(v.Input)
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

		if actual.ManagedDeviceMobileAppConfigurationStateId != v.Expected.ManagedDeviceMobileAppConfigurationStateId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceMobileAppConfigurationStateId", v.Expected.ManagedDeviceMobileAppConfigurationStateId, actual.ManagedDeviceMobileAppConfigurationStateId)
		}

	}
}

func TestSegmentsForDeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId(t *testing.T) {
	segments := DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementManagedDeviceIdManagedDeviceMobileAppConfigurationStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
