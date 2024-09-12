package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIosUpdateStatusId{}

func TestNewDeviceManagementIosUpdateStatusID(t *testing.T) {
	id := NewDeviceManagementIosUpdateStatusID("iosUpdateDeviceStatusIdValue")

	if id.IosUpdateDeviceStatusId != "iosUpdateDeviceStatusIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'IosUpdateDeviceStatusId'", id.IosUpdateDeviceStatusId, "iosUpdateDeviceStatusIdValue")
	}
}

func TestFormatDeviceManagementIosUpdateStatusID(t *testing.T) {
	actual := NewDeviceManagementIosUpdateStatusID("iosUpdateDeviceStatusIdValue").ID()
	expected := "/deviceManagement/iosUpdateStatuses/iosUpdateDeviceStatusIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementIosUpdateStatusID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIosUpdateStatusId
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
			Input: "/deviceManagement/iosUpdateStatuses",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/iosUpdateStatuses/iosUpdateDeviceStatusIdValue",
			Expected: &DeviceManagementIosUpdateStatusId{
				IosUpdateDeviceStatusId: "iosUpdateDeviceStatusIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/iosUpdateStatuses/iosUpdateDeviceStatusIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIosUpdateStatusID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.IosUpdateDeviceStatusId != v.Expected.IosUpdateDeviceStatusId {
			t.Fatalf("Expected %q but got %q for IosUpdateDeviceStatusId", v.Expected.IosUpdateDeviceStatusId, actual.IosUpdateDeviceStatusId)
		}

	}
}

func TestParseDeviceManagementIosUpdateStatusIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIosUpdateStatusId
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
			Input: "/deviceManagement/iosUpdateStatuses",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iOsUpDaTeStAtUsEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/iosUpdateStatuses/iosUpdateDeviceStatusIdValue",
			Expected: &DeviceManagementIosUpdateStatusId{
				IosUpdateDeviceStatusId: "iosUpdateDeviceStatusIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/iosUpdateStatuses/iosUpdateDeviceStatusIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iOsUpDaTeStAtUsEs/iOsUpDaTeDeViCeStAtUsIdVaLuE",
			Expected: &DeviceManagementIosUpdateStatusId{
				IosUpdateDeviceStatusId: "iOsUpDaTeDeViCeStAtUsIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iOsUpDaTeStAtUsEs/iOsUpDaTeDeViCeStAtUsIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIosUpdateStatusIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.IosUpdateDeviceStatusId != v.Expected.IosUpdateDeviceStatusId {
			t.Fatalf("Expected %q but got %q for IosUpdateDeviceStatusId", v.Expected.IosUpdateDeviceStatusId, actual.IosUpdateDeviceStatusId)
		}

	}
}

func TestSegmentsForDeviceManagementIosUpdateStatusId(t *testing.T) {
	segments := DeviceManagementIosUpdateStatusId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementIosUpdateStatusId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
