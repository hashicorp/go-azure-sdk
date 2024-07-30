package templatemigratabletocategoryrecommendedsetting

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId{}

func TestNewDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID(t *testing.T) {
	id := NewDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID("deviceManagementTemplateIdValue", "deviceManagementTemplateId1Value", "deviceManagementTemplateSettingCategoryIdValue", "deviceManagementSettingInstanceIdValue")

	if id.DeviceManagementTemplateId != "deviceManagementTemplateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateId'", id.DeviceManagementTemplateId, "deviceManagementTemplateIdValue")
	}

	if id.DeviceManagementTemplateId1 != "deviceManagementTemplateId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateId1'", id.DeviceManagementTemplateId1, "deviceManagementTemplateId1Value")
	}

	if id.DeviceManagementTemplateSettingCategoryId != "deviceManagementTemplateSettingCategoryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateSettingCategoryId'", id.DeviceManagementTemplateSettingCategoryId, "deviceManagementTemplateSettingCategoryIdValue")
	}

	if id.DeviceManagementSettingInstanceId != "deviceManagementSettingInstanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementSettingInstanceId'", id.DeviceManagementSettingInstanceId, "deviceManagementSettingInstanceIdValue")
	}
}

func TestFormatDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID(t *testing.T) {
	actual := NewDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID("deviceManagementTemplateIdValue", "deviceManagementTemplateId1Value", "deviceManagementTemplateSettingCategoryIdValue", "deviceManagementSettingInstanceIdValue").ID()
	expected := "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value/categories/deviceManagementTemplateSettingCategoryIdValue/recommendedSettings/deviceManagementSettingInstanceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId
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
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value/categories/deviceManagementTemplateSettingCategoryIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value/categories/deviceManagementTemplateSettingCategoryIdValue/recommendedSettings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value/categories/deviceManagementTemplateSettingCategoryIdValue/recommendedSettings/deviceManagementSettingInstanceIdValue",
			Expected: &DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId{
				DeviceManagementTemplateId:                "deviceManagementTemplateIdValue",
				DeviceManagementTemplateId1:               "deviceManagementTemplateId1Value",
				DeviceManagementTemplateSettingCategoryId: "deviceManagementTemplateSettingCategoryIdValue",
				DeviceManagementSettingInstanceId:         "deviceManagementSettingInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value/categories/deviceManagementTemplateSettingCategoryIdValue/recommendedSettings/deviceManagementSettingInstanceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID(v.Input)
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

		if actual.DeviceManagementSettingInstanceId != v.Expected.DeviceManagementSettingInstanceId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingInstanceId", v.Expected.DeviceManagementSettingInstanceId, actual.DeviceManagementSettingInstanceId)
		}

	}
}

func TestParseDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId
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
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value/categories/deviceManagementTemplateSettingCategoryIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1VaLuE/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value/categories/deviceManagementTemplateSettingCategoryIdValue/recommendedSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1VaLuE/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyIdVaLuE/rEcOmMeNdEdSeTtInGs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value/categories/deviceManagementTemplateSettingCategoryIdValue/recommendedSettings/deviceManagementSettingInstanceIdValue",
			Expected: &DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId{
				DeviceManagementTemplateId:                "deviceManagementTemplateIdValue",
				DeviceManagementTemplateId1:               "deviceManagementTemplateId1Value",
				DeviceManagementTemplateSettingCategoryId: "deviceManagementTemplateSettingCategoryIdValue",
				DeviceManagementSettingInstanceId:         "deviceManagementSettingInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/migratableTo/deviceManagementTemplateId1Value/categories/deviceManagementTemplateSettingCategoryIdValue/recommendedSettings/deviceManagementSettingInstanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1VaLuE/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyIdVaLuE/rEcOmMeNdEdSeTtInGs/dEvIcEmAnAgEmEnTsEtTiNgInStAnCeIdVaLuE",
			Expected: &DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId{
				DeviceManagementTemplateId:                "dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe",
				DeviceManagementTemplateId1:               "dEvIcEmAnAgEmEnTtEmPlAtEiD1VaLuE",
				DeviceManagementTemplateSettingCategoryId: "dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyIdVaLuE",
				DeviceManagementSettingInstanceId:         "dEvIcEmAnAgEmEnTsEtTiNgInStAnCeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1VaLuE/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyIdVaLuE/rEcOmMeNdEdSeTtInGs/dEvIcEmAnAgEmEnTsEtTiNgInStAnCeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingIDInsensitively(v.Input)
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

		if actual.DeviceManagementSettingInstanceId != v.Expected.DeviceManagementSettingInstanceId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingInstanceId", v.Expected.DeviceManagementSettingInstanceId, actual.DeviceManagementSettingInstanceId)
		}

	}
}

func TestSegmentsForDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId(t *testing.T) {
	segments := DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
