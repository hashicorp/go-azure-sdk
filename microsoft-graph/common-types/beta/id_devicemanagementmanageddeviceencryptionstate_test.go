package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceEncryptionStateId{}

func TestNewDeviceManagementManagedDeviceEncryptionStateID(t *testing.T) {
	id := NewDeviceManagementManagedDeviceEncryptionStateID("managedDeviceEncryptionStateId")

	if id.ManagedDeviceEncryptionStateId != "managedDeviceEncryptionStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceEncryptionStateId'", id.ManagedDeviceEncryptionStateId, "managedDeviceEncryptionStateId")
	}
}

func TestFormatDeviceManagementManagedDeviceEncryptionStateID(t *testing.T) {
	actual := NewDeviceManagementManagedDeviceEncryptionStateID("managedDeviceEncryptionStateId").ID()
	expected := "/deviceManagement/managedDeviceEncryptionStates/managedDeviceEncryptionStateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementManagedDeviceEncryptionStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceEncryptionStateId
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
			Input: "/deviceManagement/managedDeviceEncryptionStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDeviceEncryptionStates/managedDeviceEncryptionStateId",
			Expected: &DeviceManagementManagedDeviceEncryptionStateId{
				ManagedDeviceEncryptionStateId: "managedDeviceEncryptionStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDeviceEncryptionStates/managedDeviceEncryptionStateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceEncryptionStateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceEncryptionStateId != v.Expected.ManagedDeviceEncryptionStateId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceEncryptionStateId", v.Expected.ManagedDeviceEncryptionStateId, actual.ManagedDeviceEncryptionStateId)
		}

	}
}

func TestParseDeviceManagementManagedDeviceEncryptionStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceEncryptionStateId
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
			Input: "/deviceManagement/managedDeviceEncryptionStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeEnCrYpTiOnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDeviceEncryptionStates/managedDeviceEncryptionStateId",
			Expected: &DeviceManagementManagedDeviceEncryptionStateId{
				ManagedDeviceEncryptionStateId: "managedDeviceEncryptionStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDeviceEncryptionStates/managedDeviceEncryptionStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeEnCrYpTiOnStAtEs/mAnAgEdDeViCeEnCrYpTiOnStAtEiD",
			Expected: &DeviceManagementManagedDeviceEncryptionStateId{
				ManagedDeviceEncryptionStateId: "mAnAgEdDeViCeEnCrYpTiOnStAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeEnCrYpTiOnStAtEs/mAnAgEdDeViCeEnCrYpTiOnStAtEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceEncryptionStateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceEncryptionStateId != v.Expected.ManagedDeviceEncryptionStateId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceEncryptionStateId", v.Expected.ManagedDeviceEncryptionStateId, actual.ManagedDeviceEncryptionStateId)
		}

	}
}

func TestSegmentsForDeviceManagementManagedDeviceEncryptionStateId(t *testing.T) {
	segments := DeviceManagementManagedDeviceEncryptionStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementManagedDeviceEncryptionStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
