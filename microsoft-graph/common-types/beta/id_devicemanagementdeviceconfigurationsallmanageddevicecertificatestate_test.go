package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId{}

func TestNewDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateID(t *testing.T) {
	id := NewDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateID("managedAllDeviceCertificateStateId")

	if id.ManagedAllDeviceCertificateStateId != "managedAllDeviceCertificateStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedAllDeviceCertificateStateId'", id.ManagedAllDeviceCertificateStateId, "managedAllDeviceCertificateStateId")
	}
}

func TestFormatDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateID(t *testing.T) {
	actual := NewDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateID("managedAllDeviceCertificateStateId").ID()
	expected := "/deviceManagement/deviceConfigurationsAllManagedDeviceCertificateStates/managedAllDeviceCertificateStateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId
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
			Input: "/deviceManagement/deviceConfigurationsAllManagedDeviceCertificateStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurationsAllManagedDeviceCertificateStates/managedAllDeviceCertificateStateId",
			Expected: &DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId{
				ManagedAllDeviceCertificateStateId: "managedAllDeviceCertificateStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurationsAllManagedDeviceCertificateStates/managedAllDeviceCertificateStateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedAllDeviceCertificateStateId != v.Expected.ManagedAllDeviceCertificateStateId {
			t.Fatalf("Expected %q but got %q for ManagedAllDeviceCertificateStateId", v.Expected.ManagedAllDeviceCertificateStateId, actual.ManagedAllDeviceCertificateStateId)
		}

	}
}

func TestParseDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId
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
			Input: "/deviceManagement/deviceConfigurationsAllManagedDeviceCertificateStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnSaLlMaNaGeDdEvIcEcErTiFiCaTeStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceConfigurationsAllManagedDeviceCertificateStates/managedAllDeviceCertificateStateId",
			Expected: &DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId{
				ManagedAllDeviceCertificateStateId: "managedAllDeviceCertificateStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceConfigurationsAllManagedDeviceCertificateStates/managedAllDeviceCertificateStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnSaLlMaNaGeDdEvIcEcErTiFiCaTeStAtEs/mAnAgEdAlLdEvIcEcErTiFiCaTeStAtEiD",
			Expected: &DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId{
				ManagedAllDeviceCertificateStateId: "mAnAgEdAlLdEvIcEcErTiFiCaTeStAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcOnFiGuRaTiOnSaLlMaNaGeDdEvIcEcErTiFiCaTeStAtEs/mAnAgEdAlLdEvIcEcErTiFiCaTeStAtEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedAllDeviceCertificateStateId != v.Expected.ManagedAllDeviceCertificateStateId {
			t.Fatalf("Expected %q but got %q for ManagedAllDeviceCertificateStateId", v.Expected.ManagedAllDeviceCertificateStateId, actual.ManagedAllDeviceCertificateStateId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId(t *testing.T) {
	segments := DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
