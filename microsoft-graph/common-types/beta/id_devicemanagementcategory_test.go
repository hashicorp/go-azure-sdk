package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCategoryId{}

func TestNewDeviceManagementCategoryID(t *testing.T) {
	id := NewDeviceManagementCategoryID("deviceManagementSettingCategoryId")

	if id.DeviceManagementSettingCategoryId != "deviceManagementSettingCategoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementSettingCategoryId'", id.DeviceManagementSettingCategoryId, "deviceManagementSettingCategoryId")
	}
}

func TestFormatDeviceManagementCategoryID(t *testing.T) {
	actual := NewDeviceManagementCategoryID("deviceManagementSettingCategoryId").ID()
	expected := "/deviceManagement/categories/deviceManagementSettingCategoryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementCategoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCategoryId
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
			Input: "/deviceManagement/categories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/categories/deviceManagementSettingCategoryId",
			Expected: &DeviceManagementCategoryId{
				DeviceManagementSettingCategoryId: "deviceManagementSettingCategoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/categories/deviceManagementSettingCategoryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCategoryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementSettingCategoryId != v.Expected.DeviceManagementSettingCategoryId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingCategoryId", v.Expected.DeviceManagementSettingCategoryId, actual.DeviceManagementSettingCategoryId)
		}

	}
}

func TestParseDeviceManagementCategoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCategoryId
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
			Input: "/deviceManagement/categories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cAtEgOrIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/categories/deviceManagementSettingCategoryId",
			Expected: &DeviceManagementCategoryId{
				DeviceManagementSettingCategoryId: "deviceManagementSettingCategoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/categories/deviceManagementSettingCategoryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cAtEgOrIeS/dEvIcEmAnAgEmEnTsEtTiNgCaTeGoRyId",
			Expected: &DeviceManagementCategoryId{
				DeviceManagementSettingCategoryId: "dEvIcEmAnAgEmEnTsEtTiNgCaTeGoRyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cAtEgOrIeS/dEvIcEmAnAgEmEnTsEtTiNgCaTeGoRyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCategoryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementSettingCategoryId != v.Expected.DeviceManagementSettingCategoryId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingCategoryId", v.Expected.DeviceManagementSettingCategoryId, actual.DeviceManagementSettingCategoryId)
		}

	}
}

func TestSegmentsForDeviceManagementCategoryId(t *testing.T) {
	segments := DeviceManagementCategoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementCategoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
