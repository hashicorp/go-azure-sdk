package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdMigratableToIdCategoryId{}

func TestNewDeviceManagementTemplateIdMigratableToIdCategoryID(t *testing.T) {
	id := NewDeviceManagementTemplateIdMigratableToIdCategoryID("deviceManagementTemplateId", "deviceManagementTemplateId1", "deviceManagementTemplateSettingCategoryId")

	if id.DeviceManagementTemplateId != "deviceManagementTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateId'", id.DeviceManagementTemplateId, "deviceManagementTemplateId")
	}

	if id.DeviceManagementTemplateId1 != "deviceManagementTemplateId1" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateId1'", id.DeviceManagementTemplateId1, "deviceManagementTemplateId1")
	}

	if id.DeviceManagementTemplateSettingCategoryId != "deviceManagementTemplateSettingCategoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateSettingCategoryId'", id.DeviceManagementTemplateSettingCategoryId, "deviceManagementTemplateSettingCategoryId")
	}
}

func TestFormatDeviceManagementTemplateIdMigratableToIdCategoryID(t *testing.T) {
	actual := NewDeviceManagementTemplateIdMigratableToIdCategoryID("deviceManagementTemplateId", "deviceManagementTemplateId1", "deviceManagementTemplateSettingCategoryId").ID()
	expected := "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId"
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
			// Valid URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId",
			Expected: &DeviceManagementTemplateIdMigratableToIdCategoryId{
				DeviceManagementTemplateId:                "deviceManagementTemplateId",
				DeviceManagementTemplateId1:               "deviceManagementTemplateId1",
				DeviceManagementTemplateSettingCategoryId: "deviceManagementTemplateSettingCategoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId/extra",
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
			// Valid URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId",
			Expected: &DeviceManagementTemplateIdMigratableToIdCategoryId{
				DeviceManagementTemplateId:                "deviceManagementTemplateId",
				DeviceManagementTemplateId1:               "deviceManagementTemplateId1",
				DeviceManagementTemplateSettingCategoryId: "deviceManagementTemplateSettingCategoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templates/deviceManagementTemplateId/migratableTo/deviceManagementTemplateId1/categories/deviceManagementTemplateSettingCategoryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyId",
			Expected: &DeviceManagementTemplateIdMigratableToIdCategoryId{
				DeviceManagementTemplateId:                "dEvIcEmAnAgEmEnTtEmPlAtEiD",
				DeviceManagementTemplateId1:               "dEvIcEmAnAgEmEnTtEmPlAtEiD1",
				DeviceManagementTemplateSettingCategoryId: "dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/mIgRaTaBlEtO/dEvIcEmAnAgEmEnTtEmPlAtEiD1/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyId/extra",
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
