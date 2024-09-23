package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDeviceCategoryId{}

func TestNewDeviceManagementDeviceCategoryID(t *testing.T) {
	id := NewDeviceManagementDeviceCategoryID("deviceCategoryId")

	if id.DeviceCategoryId != "deviceCategoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCategoryId'", id.DeviceCategoryId, "deviceCategoryId")
	}
}

func TestFormatDeviceManagementDeviceCategoryID(t *testing.T) {
	actual := NewDeviceManagementDeviceCategoryID("deviceCategoryId").ID()
	expected := "/deviceManagement/deviceCategories/deviceCategoryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDeviceCategoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCategoryId
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
			Input: "/deviceManagement/deviceCategories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCategories/deviceCategoryId",
			Expected: &DeviceManagementDeviceCategoryId{
				DeviceCategoryId: "deviceCategoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCategories/deviceCategoryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCategoryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceCategoryId != v.Expected.DeviceCategoryId {
			t.Fatalf("Expected %q but got %q for DeviceCategoryId", v.Expected.DeviceCategoryId, actual.DeviceCategoryId)
		}

	}
}

func TestParseDeviceManagementDeviceCategoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDeviceCategoryId
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
			Input: "/deviceManagement/deviceCategories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcAtEgOrIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/deviceCategories/deviceCategoryId",
			Expected: &DeviceManagementDeviceCategoryId{
				DeviceCategoryId: "deviceCategoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/deviceCategories/deviceCategoryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcAtEgOrIeS/dEvIcEcAtEgOrYiD",
			Expected: &DeviceManagementDeviceCategoryId{
				DeviceCategoryId: "dEvIcEcAtEgOrYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEvIcEcAtEgOrIeS/dEvIcEcAtEgOrYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDeviceCategoryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceCategoryId != v.Expected.DeviceCategoryId {
			t.Fatalf("Expected %q but got %q for DeviceCategoryId", v.Expected.DeviceCategoryId, actual.DeviceCategoryId)
		}

	}
}

func TestSegmentsForDeviceManagementDeviceCategoryId(t *testing.T) {
	segments := DeviceManagementDeviceCategoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDeviceCategoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
