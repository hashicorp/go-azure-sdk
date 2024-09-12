package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationCategoryId{}

func TestNewDeviceManagementConfigurationCategoryID(t *testing.T) {
	id := NewDeviceManagementConfigurationCategoryID("deviceManagementConfigurationCategoryIdValue")

	if id.DeviceManagementConfigurationCategoryId != "deviceManagementConfigurationCategoryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationCategoryId'", id.DeviceManagementConfigurationCategoryId, "deviceManagementConfigurationCategoryIdValue")
	}
}

func TestFormatDeviceManagementConfigurationCategoryID(t *testing.T) {
	actual := NewDeviceManagementConfigurationCategoryID("deviceManagementConfigurationCategoryIdValue").ID()
	expected := "/deviceManagement/configurationCategories/deviceManagementConfigurationCategoryIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementConfigurationCategoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigurationCategoryId
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
			Input: "/deviceManagement/configurationCategories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationCategories/deviceManagementConfigurationCategoryIdValue",
			Expected: &DeviceManagementConfigurationCategoryId{
				DeviceManagementConfigurationCategoryId: "deviceManagementConfigurationCategoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationCategories/deviceManagementConfigurationCategoryIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigurationCategoryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationCategoryId != v.Expected.DeviceManagementConfigurationCategoryId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationCategoryId", v.Expected.DeviceManagementConfigurationCategoryId, actual.DeviceManagementConfigurationCategoryId)
		}

	}
}

func TestParseDeviceManagementConfigurationCategoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigurationCategoryId
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
			Input: "/deviceManagement/configurationCategories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnCaTeGoRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configurationCategories/deviceManagementConfigurationCategoryIdValue",
			Expected: &DeviceManagementConfigurationCategoryId{
				DeviceManagementConfigurationCategoryId: "deviceManagementConfigurationCategoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configurationCategories/deviceManagementConfigurationCategoryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnCaTeGoRiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnCaTeGoRyIdVaLuE",
			Expected: &DeviceManagementConfigurationCategoryId{
				DeviceManagementConfigurationCategoryId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnCaTeGoRyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGuRaTiOnCaTeGoRiEs/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnCaTeGoRyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigurationCategoryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationCategoryId != v.Expected.DeviceManagementConfigurationCategoryId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationCategoryId", v.Expected.DeviceManagementConfigurationCategoryId, actual.DeviceManagementConfigurationCategoryId)
		}

	}
}

func TestSegmentsForDeviceManagementConfigurationCategoryId(t *testing.T) {
	segments := DeviceManagementConfigurationCategoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementConfigurationCategoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
