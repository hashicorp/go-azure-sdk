package comanageddevicemanageddevicemobileappconfigurationstate

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId{}

func TestNewDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateID(t *testing.T) {
	id := NewDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateID("managedDeviceIdValue", "managedDeviceMobileAppConfigurationStateIdValue")

	if id.ManagedDeviceId != "managedDeviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceIdValue")
	}

	if id.ManagedDeviceMobileAppConfigurationStateId != "managedDeviceMobileAppConfigurationStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceMobileAppConfigurationStateId'", id.ManagedDeviceMobileAppConfigurationStateId, "managedDeviceMobileAppConfigurationStateIdValue")
	}
}

func TestFormatDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateID(t *testing.T) {
	actual := NewDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateID("managedDeviceIdValue", "managedDeviceMobileAppConfigurationStateIdValue").ID()
	expected := "/deviceManagement/comanagedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates/managedDeviceMobileAppConfigurationStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId
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
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates/managedDeviceMobileAppConfigurationStateIdValue",
			Expected: &DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId{
				ManagedDeviceId: "managedDeviceIdValue",
				ManagedDeviceMobileAppConfigurationStateId: "managedDeviceMobileAppConfigurationStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates/managedDeviceMobileAppConfigurationStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateID(v.Input)
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

func TestParseDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId
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
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates/managedDeviceMobileAppConfigurationStateIdValue",
			Expected: &DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId{
				ManagedDeviceId: "managedDeviceIdValue",
				ManagedDeviceMobileAppConfigurationStateId: "managedDeviceMobileAppConfigurationStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/comanagedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates/managedDeviceMobileAppConfigurationStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEs/mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEiDvAlUe",
			Expected: &DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId{
				ManagedDeviceId: "mAnAgEdDeViCeIdVaLuE",
				ManagedDeviceMobileAppConfigurationStateId: "mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEs/mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateIDInsensitively(v.Input)
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

func TestSegmentsForDeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId(t *testing.T) {
	segments := DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementComanagedDeviceIdManagedDeviceMobileAppConfigurationStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
