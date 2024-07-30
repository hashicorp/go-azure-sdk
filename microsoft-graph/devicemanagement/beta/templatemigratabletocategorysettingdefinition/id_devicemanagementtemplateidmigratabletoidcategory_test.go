package templatemigratabletocategorysettingdefinition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdMigratableToIdCategoryId{}

func TestNewDeviceManagementTemplateIdMigratableToIdCategoryID(t *testing.T) {
	id := NewDeviceManagementTemplateIdMigratableToIdCategoryID("deviceManagementTemplateIdValue", "deviceManagementTemplateId1Value", "deviceManagementTemplateSettingCategoryIdValue")

	if id.DeviceManagementTemplateId != "deviceManagementTemplateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateId'", id.DeviceManagementTemplateId, "deviceManagementTemplateIdValue")
	}

	if id.DeviceManagementTemplateId1 != "deviceManagementTemplateId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateId1'", id.DeviceManagementTemplateId1, "deviceManagementTemplateId1Value")
	}

	if id.DeviceManagementTemplateSettingCategoryId != "deviceManagementTemplateSettingCategoryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateSettingCategoryId'", id.DeviceManagementTemplateSettingCategoryId, "deviceManagementTemplateSettingCategoryIdValue")
	}
}

func TestFormatDeviceManagementTemplateIdMigratableToIdCategoryID(t *testing.T) {
	actual := NewDeviceManagementTemplateIdMigratableToIdCategoryID("deviceManagementTemplateIdValue", "deviceManagementTemplateId1Value", "deviceManagementTemplateSettingCategoryIdValue").ID()
	expected := "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value/categories/deviceManagementTemplateSettingCategoryIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementTemplateIdMigratableToIdCategoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateIdMigratableToIdCategoryId
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
			Input: "/deviceManagement/templates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value/categories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value/categories/deviceManagementTemplateSettingCategoryIdValue",
			Expected: &DeviceManagementTemplateIdMigratableToIdCategoryId{
				DeviceManagementTemplateId:                "deviceManagementTemplateIdValue",
				DeviceManagementTemplateId1:               "deviceManagementTemplateId1Value",
				DeviceManagementTemplateSettingCategoryId: "deviceManagementTemplateSettingCategoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value/categories/deviceManagementTemplateSettingCategoryIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateIdMigratableToIdCategoryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementTemplateId != v.Expected.DeviceManagementTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementTemplateId", v.Expected.DeviceManagementTemplateId, actual.DeviceManagementTemplateId)
		}

		if actual.DeviceManagementTemplateId1 != v.Expected.DeviceManagementTemplateId1 {
			t.Fatalf("Expected %q but got %q for DeviceManagementTemplateId1", v.Expected.DeviceManagementTemplateId1, actual.DeviceManagementTemplateId1)
		}

		if actual.DeviceManagementTemplateSettingCategoryId != v.Expected.DeviceManagementTemplateSettingCategoryId {
			t.Fatalf("Expected %q but got %q for DeviceManagementTemplateSettingCategoryId", v.Expected.DeviceManagementTemplateSettingCategoryId, actual.DeviceManagementTemplateSettingCategoryId)
		}

	}
}

func TestParseDeviceManagementTemplateIdMigratableToIdCategoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateIdMigratableToIdCategoryId
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
			Input: "/deviceManagement/templates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe/mIgRaTaBlEtO",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1VaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value/categories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1VaLuE/cAtEgOrIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value/categories/deviceManagementTemplateSettingCategoryIdValue",
			Expected: &DeviceManagementTemplateIdMigratableToIdCategoryId{
				DeviceManagementTemplateId:                "deviceManagementTemplateIdValue",
				DeviceManagementTemplateId1:               "deviceManagementTemplateId1Value",
				DeviceManagementTemplateSettingCategoryId: "deviceManagementTemplateSettingCategoryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value/categories/deviceManagementTemplateSettingCategoryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1VaLuE/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyIdVaLuE",
			Expected: &DeviceManagementTemplateIdMigratableToIdCategoryId{
				DeviceManagementTemplateId:                "dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe",
				DeviceManagementTemplateId1:               "dEvIcEmAnAgEmEnTtEmPlAtEiD1VaLuE",
				DeviceManagementTemplateSettingCategoryId: "dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1VaLuE/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateIdMigratableToIdCategoryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementTemplateId != v.Expected.DeviceManagementTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementTemplateId", v.Expected.DeviceManagementTemplateId, actual.DeviceManagementTemplateId)
		}

		if actual.DeviceManagementTemplateId1 != v.Expected.DeviceManagementTemplateId1 {
			t.Fatalf("Expected %q but got %q for DeviceManagementTemplateId1", v.Expected.DeviceManagementTemplateId1, actual.DeviceManagementTemplateId1)
		}

		if actual.DeviceManagementTemplateSettingCategoryId != v.Expected.DeviceManagementTemplateSettingCategoryId {
			t.Fatalf("Expected %q but got %q for DeviceManagementTemplateSettingCategoryId", v.Expected.DeviceManagementTemplateSettingCategoryId, actual.DeviceManagementTemplateSettingCategoryId)
		}

	}
}

func TestSegmentsForDeviceManagementTemplateIdMigratableToIdCategoryId(t *testing.T) {
	segments := DeviceManagementTemplateIdMigratableToIdCategoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementTemplateIdMigratableToIdCategoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
