package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId{}

func TestNewDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionID(t *testing.T) {
	id := NewDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionID("deviceManagementTemplateId", "deviceManagementTemplateId1", "deviceManagementTemplateSettingCategoryId", "deviceManagementSettingDefinitionId")

	if id.DeviceManagementTemplateId != "deviceManagementTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateId'", id.DeviceManagementTemplateId, "deviceManagementTemplateId")
	}

	if id.DeviceManagementTemplateId1 != "deviceManagementTemplateId1" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateId1'", id.DeviceManagementTemplateId1, "deviceManagementTemplateId1")
	}

	if id.DeviceManagementTemplateSettingCategoryId != "deviceManagementTemplateSettingCategoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateSettingCategoryId'", id.DeviceManagementTemplateSettingCategoryId, "deviceManagementTemplateSettingCategoryId")
	}

	if id.DeviceManagementSettingDefinitionId != "deviceManagementSettingDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementSettingDefinitionId'", id.DeviceManagementSettingDefinitionId, "deviceManagementSettingDefinitionId")
	}
}

func TestFormatDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionID(t *testing.T) {
	actual := NewDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionID("deviceManagementTemplateId", "deviceManagementTemplateId1", "deviceManagementTemplateSettingCategoryId", "deviceManagementSettingDefinitionId").ID()
	expected := "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId/settingDefinitions/deviceManagementSettingDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId
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
			Input: "/deviceManagement/templates/deviceManagementTemplateId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId/settingDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId/settingDefinitions/deviceManagementSettingDefinitionId",
			Expected: &DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId{
				DeviceManagementTemplateId:                "deviceManagementTemplateId",
				DeviceManagementTemplateId1:               "deviceManagementTemplateId1",
				DeviceManagementTemplateSettingCategoryId: "deviceManagementTemplateSettingCategoryId",
				DeviceManagementSettingDefinitionId:       "deviceManagementSettingDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId/settingDefinitions/deviceManagementSettingDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionID(v.Input)
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

		if actual.DeviceManagementSettingDefinitionId != v.Expected.DeviceManagementSettingDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingDefinitionId", v.Expected.DeviceManagementSettingDefinitionId, actual.DeviceManagementSettingDefinitionId)
		}

	}
}

func TestParseDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId
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
			Input: "/deviceManagement/templates/deviceManagementTemplateId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/mIgRaTaBlEtO",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1/cAtEgOrIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId/settingDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyId/sEtTiNgDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId/settingDefinitions/deviceManagementSettingDefinitionId",
			Expected: &DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId{
				DeviceManagementTemplateId:                "deviceManagementTemplateId",
				DeviceManagementTemplateId1:               "deviceManagementTemplateId1",
				DeviceManagementTemplateSettingCategoryId: "deviceManagementTemplateSettingCategoryId",
				DeviceManagementSettingDefinitionId:       "deviceManagementSettingDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId/settingDefinitions/deviceManagementSettingDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyId/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTsEtTiNgDeFiNiTiOnId",
			Expected: &DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId{
				DeviceManagementTemplateId:                "dEvIcEmAnAgEmEnTtEmPlAtEiD",
				DeviceManagementTemplateId1:               "dEvIcEmAnAgEmEnTtEmPlAtEiD1",
				DeviceManagementTemplateSettingCategoryId: "dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyId",
				DeviceManagementSettingDefinitionId:       "dEvIcEmAnAgEmEnTsEtTiNgDeFiNiTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyId/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTsEtTiNgDeFiNiTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionIDInsensitively(v.Input)
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

		if actual.DeviceManagementSettingDefinitionId != v.Expected.DeviceManagementSettingDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingDefinitionId", v.Expected.DeviceManagementSettingDefinitionId, actual.DeviceManagementSettingDefinitionId)
		}

	}
}

func TestSegmentsForDeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId(t *testing.T) {
	segments := DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
