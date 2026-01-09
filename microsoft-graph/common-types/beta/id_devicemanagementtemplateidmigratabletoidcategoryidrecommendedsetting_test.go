package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId{}

func TestNewDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID(t *testing.T) {
	id := NewDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID("deviceManagementTemplateId", "deviceManagementTemplateId1", "deviceManagementTemplateSettingCategoryId", "deviceManagementSettingInstanceId")

	if id.DeviceManagementTemplateId != "deviceManagementTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateId'", id.DeviceManagementTemplateId, "deviceManagementTemplateId")
	}

	if id.DeviceManagementTemplateId1 != "deviceManagementTemplateId1" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateId1'", id.DeviceManagementTemplateId1, "deviceManagementTemplateId1")
	}

	if id.DeviceManagementTemplateSettingCategoryId != "deviceManagementTemplateSettingCategoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateSettingCategoryId'", id.DeviceManagementTemplateSettingCategoryId, "deviceManagementTemplateSettingCategoryId")
	}

	if id.DeviceManagementSettingInstanceId != "deviceManagementSettingInstanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementSettingInstanceId'", id.DeviceManagementSettingInstanceId, "deviceManagementSettingInstanceId")
	}
}

func TestFormatDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID(t *testing.T) {
	actual := NewDeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingID("deviceManagementTemplateId", "deviceManagementTemplateId1", "deviceManagementTemplateSettingCategoryId", "deviceManagementSettingInstanceId").ID()
	expected := "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId/recommendedSettings/deviceManagementSettingInstanceId"
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
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId/recommendedSettings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId/recommendedSettings/deviceManagementSettingInstanceId",
			Expected: &DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId{
				DeviceManagementTemplateId:                "deviceManagementTemplateId",
				DeviceManagementTemplateId1:               "deviceManagementTemplateId1",
				DeviceManagementTemplateSettingCategoryId: "deviceManagementTemplateSettingCategoryId",
				DeviceManagementSettingInstanceId:         "deviceManagementSettingInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId/recommendedSettings/deviceManagementSettingInstanceId/extra",
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
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId/recommendedSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyId/rEcOmMeNdEdSeTtInGs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId/recommendedSettings/deviceManagementSettingInstanceId",
			Expected: &DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId{
				DeviceManagementTemplateId:                "deviceManagementTemplateId",
				DeviceManagementTemplateId1:               "deviceManagementTemplateId1",
				DeviceManagementTemplateSettingCategoryId: "deviceManagementTemplateSettingCategoryId",
				DeviceManagementSettingInstanceId:         "deviceManagementSettingInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId/recommendedSettings/deviceManagementSettingInstanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyId/rEcOmMeNdEdSeTtInGs/dEvIcEmAnAgEmEnTsEtTiNgInStAnCeId",
			Expected: &DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId{
				DeviceManagementTemplateId:                "dEvIcEmAnAgEmEnTtEmPlAtEiD",
				DeviceManagementTemplateId1:               "dEvIcEmAnAgEmEnTtEmPlAtEiD1",
				DeviceManagementTemplateSettingCategoryId: "dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyId",
				DeviceManagementSettingInstanceId:         "dEvIcEmAnAgEmEnTsEtTiNgInStAnCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyId/rEcOmMeNdEdSeTtInGs/dEvIcEmAnAgEmEnTsEtTiNgInStAnCeId/extra",
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
