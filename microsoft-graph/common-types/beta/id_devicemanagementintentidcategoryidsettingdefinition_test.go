package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdCategoryIdSettingDefinitionId{}

func TestNewDeviceManagementIntentIdCategoryIdSettingDefinitionID(t *testing.T) {
	id := NewDeviceManagementIntentIdCategoryIdSettingDefinitionID("deviceManagementIntentId", "deviceManagementIntentSettingCategoryId", "deviceManagementSettingDefinitionId")

	if id.DeviceManagementIntentId != "deviceManagementIntentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentId'", id.DeviceManagementIntentId, "deviceManagementIntentId")
	}

	if id.DeviceManagementIntentSettingCategoryId != "deviceManagementIntentSettingCategoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentSettingCategoryId'", id.DeviceManagementIntentSettingCategoryId, "deviceManagementIntentSettingCategoryId")
	}

	if id.DeviceManagementSettingDefinitionId != "deviceManagementSettingDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementSettingDefinitionId'", id.DeviceManagementSettingDefinitionId, "deviceManagementSettingDefinitionId")
	}
}

func TestFormatDeviceManagementIntentIdCategoryIdSettingDefinitionID(t *testing.T) {
	actual := NewDeviceManagementIntentIdCategoryIdSettingDefinitionID("deviceManagementIntentId", "deviceManagementIntentSettingCategoryId", "deviceManagementSettingDefinitionId").ID()
	expected := "/deviceManagement/intents/deviceManagementIntentId/categories/deviceManagementIntentSettingCategoryId/settingDefinitions/deviceManagementSettingDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementIntentIdCategoryIdSettingDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntentIdCategoryIdSettingDefinitionId
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
			Input: "/deviceManagement/intents",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/categories",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/categories/deviceManagementIntentSettingCategoryId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/categories/deviceManagementIntentSettingCategoryId/settingDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/categories/deviceManagementIntentSettingCategoryId/settingDefinitions/deviceManagementSettingDefinitionId",
			Expected: &DeviceManagementIntentIdCategoryIdSettingDefinitionId{
				DeviceManagementIntentId:                "deviceManagementIntentId",
				DeviceManagementIntentSettingCategoryId: "deviceManagementIntentSettingCategoryId",
				DeviceManagementSettingDefinitionId:     "deviceManagementSettingDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentId/categories/deviceManagementIntentSettingCategoryId/settingDefinitions/deviceManagementSettingDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntentIdCategoryIdSettingDefinitionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementIntentId != v.Expected.DeviceManagementIntentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentId", v.Expected.DeviceManagementIntentId, actual.DeviceManagementIntentId)
		}

		if actual.DeviceManagementIntentSettingCategoryId != v.Expected.DeviceManagementIntentSettingCategoryId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentSettingCategoryId", v.Expected.DeviceManagementIntentSettingCategoryId, actual.DeviceManagementIntentSettingCategoryId)
		}

		if actual.DeviceManagementSettingDefinitionId != v.Expected.DeviceManagementSettingDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingDefinitionId", v.Expected.DeviceManagementSettingDefinitionId, actual.DeviceManagementSettingDefinitionId)
		}

	}
}

func TestParseDeviceManagementIntentIdCategoryIdSettingDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntentIdCategoryIdSettingDefinitionId
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
			Input: "/deviceManagement/intents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/categories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiD/cAtEgOrIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/categories/deviceManagementIntentSettingCategoryId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiD/cAtEgOrIeS/dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/categories/deviceManagementIntentSettingCategoryId/settingDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiD/cAtEgOrIeS/dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyId/sEtTiNgDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/categories/deviceManagementIntentSettingCategoryId/settingDefinitions/deviceManagementSettingDefinitionId",
			Expected: &DeviceManagementIntentIdCategoryIdSettingDefinitionId{
				DeviceManagementIntentId:                "deviceManagementIntentId",
				DeviceManagementIntentSettingCategoryId: "deviceManagementIntentSettingCategoryId",
				DeviceManagementSettingDefinitionId:     "deviceManagementSettingDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentId/categories/deviceManagementIntentSettingCategoryId/settingDefinitions/deviceManagementSettingDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiD/cAtEgOrIeS/dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyId/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTsEtTiNgDeFiNiTiOnId",
			Expected: &DeviceManagementIntentIdCategoryIdSettingDefinitionId{
				DeviceManagementIntentId:                "dEvIcEmAnAgEmEnTiNtEnTiD",
				DeviceManagementIntentSettingCategoryId: "dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyId",
				DeviceManagementSettingDefinitionId:     "dEvIcEmAnAgEmEnTsEtTiNgDeFiNiTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiD/cAtEgOrIeS/dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyId/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTsEtTiNgDeFiNiTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntentIdCategoryIdSettingDefinitionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementIntentId != v.Expected.DeviceManagementIntentId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentId", v.Expected.DeviceManagementIntentId, actual.DeviceManagementIntentId)
		}

		if actual.DeviceManagementIntentSettingCategoryId != v.Expected.DeviceManagementIntentSettingCategoryId {
			t.Fatalf("Expected %q but got %q for DeviceManagementIntentSettingCategoryId", v.Expected.DeviceManagementIntentSettingCategoryId, actual.DeviceManagementIntentSettingCategoryId)
		}

		if actual.DeviceManagementSettingDefinitionId != v.Expected.DeviceManagementSettingDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingDefinitionId", v.Expected.DeviceManagementSettingDefinitionId, actual.DeviceManagementSettingDefinitionId)
		}

	}
}

func TestSegmentsForDeviceManagementIntentIdCategoryIdSettingDefinitionId(t *testing.T) {
	segments := DeviceManagementIntentIdCategoryIdSettingDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementIntentIdCategoryIdSettingDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
