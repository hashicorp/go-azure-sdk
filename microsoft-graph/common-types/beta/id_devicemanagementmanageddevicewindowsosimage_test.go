package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceWindowsOSImageId{}

func TestNewDeviceManagementManagedDeviceWindowsOSImageID(t *testing.T) {
	id := NewDeviceManagementManagedDeviceWindowsOSImageID("managedDeviceWindowsOperatingSystemImageIdValue")

	if id.ManagedDeviceWindowsOperatingSystemImageId != "managedDeviceWindowsOperatingSystemImageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceWindowsOperatingSystemImageId'", id.ManagedDeviceWindowsOperatingSystemImageId, "managedDeviceWindowsOperatingSystemImageIdValue")
	}
}

func TestFormatDeviceManagementManagedDeviceWindowsOSImageID(t *testing.T) {
	actual := NewDeviceManagementManagedDeviceWindowsOSImageID("managedDeviceWindowsOperatingSystemImageIdValue").ID()
	expected := "/deviceManagement/managedDeviceWindowsOSImages/managedDeviceWindowsOperatingSystemImageIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementManagedDeviceWindowsOSImageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceWindowsOSImageId
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
			Input: "/deviceManagement/managedDeviceWindowsOSImages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDeviceWindowsOSImages/managedDeviceWindowsOperatingSystemImageIdValue",
			Expected: &DeviceManagementManagedDeviceWindowsOSImageId{
				ManagedDeviceWindowsOperatingSystemImageId: "managedDeviceWindowsOperatingSystemImageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDeviceWindowsOSImages/managedDeviceWindowsOperatingSystemImageIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceWindowsOSImageID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceWindowsOperatingSystemImageId != v.Expected.ManagedDeviceWindowsOperatingSystemImageId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceWindowsOperatingSystemImageId", v.Expected.ManagedDeviceWindowsOperatingSystemImageId, actual.ManagedDeviceWindowsOperatingSystemImageId)
		}

	}
}

func TestParseDeviceManagementManagedDeviceWindowsOSImageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceWindowsOSImageId
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
			Input: "/deviceManagement/managedDeviceWindowsOSImages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeWiNdOwSoSiMaGeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDeviceWindowsOSImages/managedDeviceWindowsOperatingSystemImageIdValue",
			Expected: &DeviceManagementManagedDeviceWindowsOSImageId{
				ManagedDeviceWindowsOperatingSystemImageId: "managedDeviceWindowsOperatingSystemImageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDeviceWindowsOSImages/managedDeviceWindowsOperatingSystemImageIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeWiNdOwSoSiMaGeS/mAnAgEdDeViCeWiNdOwSoPeRaTiNgSyStEmImAgEiDvAlUe",
			Expected: &DeviceManagementManagedDeviceWindowsOSImageId{
				ManagedDeviceWindowsOperatingSystemImageId: "mAnAgEdDeViCeWiNdOwSoPeRaTiNgSyStEmImAgEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeWiNdOwSoSiMaGeS/mAnAgEdDeViCeWiNdOwSoPeRaTiNgSyStEmImAgEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceWindowsOSImageIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceWindowsOperatingSystemImageId != v.Expected.ManagedDeviceWindowsOperatingSystemImageId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceWindowsOperatingSystemImageId", v.Expected.ManagedDeviceWindowsOperatingSystemImageId, actual.ManagedDeviceWindowsOperatingSystemImageId)
		}

	}
}

func TestSegmentsForDeviceManagementManagedDeviceWindowsOSImageId(t *testing.T) {
	segments := DeviceManagementManagedDeviceWindowsOSImageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementManagedDeviceWindowsOSImageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
