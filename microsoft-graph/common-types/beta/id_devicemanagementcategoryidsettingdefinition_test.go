package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCategoryIdSettingDefinitionId{}

func TestNewDeviceManagementCategoryIdSettingDefinitionID(t *testing.T) {
	id := NewDeviceManagementCategoryIdSettingDefinitionID("deviceManagementSettingCategoryId", "deviceManagementSettingDefinitionId")

	if id.DeviceManagementSettingCategoryId != "deviceManagementSettingCategoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementSettingCategoryId'", id.DeviceManagementSettingCategoryId, "deviceManagementSettingCategoryId")
	}

	if id.DeviceManagementSettingDefinitionId != "deviceManagementSettingDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementSettingDefinitionId'", id.DeviceManagementSettingDefinitionId, "deviceManagementSettingDefinitionId")
	}
}

func TestFormatDeviceManagementCategoryIdSettingDefinitionID(t *testing.T) {
	actual := NewDeviceManagementCategoryIdSettingDefinitionID("deviceManagementSettingCategoryId", "deviceManagementSettingDefinitionId").ID()
	expected := "/deviceManagement/categories/deviceManagementSettingCategoryId/settingDefinitions/deviceManagementSettingDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementCategoryIdSettingDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCategoryIdSettingDefinitionId
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
			// Incomplete URI
			Input: "/deviceManagement/categories/deviceManagementSettingCategoryId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/categories/deviceManagementSettingCategoryId/settingDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/categories/deviceManagementSettingCategoryId/settingDefinitions/deviceManagementSettingDefinitionId",
			Expected: &DeviceManagementCategoryIdSettingDefinitionId{
				DeviceManagementSettingCategoryId:   "deviceManagementSettingCategoryId",
				DeviceManagementSettingDefinitionId: "deviceManagementSettingDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/categories/deviceManagementSettingCategoryId/settingDefinitions/deviceManagementSettingDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCategoryIdSettingDefinitionID(v.Input)
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

		if actual.DeviceManagementSettingDefinitionId != v.Expected.DeviceManagementSettingDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingDefinitionId", v.Expected.DeviceManagementSettingDefinitionId, actual.DeviceManagementSettingDefinitionId)
		}

	}
}

func TestParseDeviceManagementCategoryIdSettingDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCategoryIdSettingDefinitionId
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
			// Incomplete URI
			Input: "/deviceManagement/categories/deviceManagementSettingCategoryId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cAtEgOrIeS/dEvIcEmAnAgEmEnTsEtTiNgCaTeGoRyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/categories/deviceManagementSettingCategoryId/settingDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cAtEgOrIeS/dEvIcEmAnAgEmEnTsEtTiNgCaTeGoRyId/sEtTiNgDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/categories/deviceManagementSettingCategoryId/settingDefinitions/deviceManagementSettingDefinitionId",
			Expected: &DeviceManagementCategoryIdSettingDefinitionId{
				DeviceManagementSettingCategoryId:   "deviceManagementSettingCategoryId",
				DeviceManagementSettingDefinitionId: "deviceManagementSettingDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/categories/deviceManagementSettingCategoryId/settingDefinitions/deviceManagementSettingDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cAtEgOrIeS/dEvIcEmAnAgEmEnTsEtTiNgCaTeGoRyId/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTsEtTiNgDeFiNiTiOnId",
			Expected: &DeviceManagementCategoryIdSettingDefinitionId{
				DeviceManagementSettingCategoryId:   "dEvIcEmAnAgEmEnTsEtTiNgCaTeGoRyId",
				DeviceManagementSettingDefinitionId: "dEvIcEmAnAgEmEnTsEtTiNgDeFiNiTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cAtEgOrIeS/dEvIcEmAnAgEmEnTsEtTiNgCaTeGoRyId/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTsEtTiNgDeFiNiTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCategoryIdSettingDefinitionIDInsensitively(v.Input)
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

		if actual.DeviceManagementSettingDefinitionId != v.Expected.DeviceManagementSettingDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingDefinitionId", v.Expected.DeviceManagementSettingDefinitionId, actual.DeviceManagementSettingDefinitionId)
		}

	}
}

func TestSegmentsForDeviceManagementCategoryIdSettingDefinitionId(t *testing.T) {
	segments := DeviceManagementCategoryIdSettingDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementCategoryIdSettingDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
