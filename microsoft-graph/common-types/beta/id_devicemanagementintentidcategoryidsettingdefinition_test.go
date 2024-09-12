package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdCategoryIdSettingDefinitionId{}

func TestNewDeviceManagementIntentIdCategoryIdSettingDefinitionID(t *testing.T) {
	id := NewDeviceManagementIntentIdCategoryIdSettingDefinitionID("deviceManagementIntentIdValue", "deviceManagementIntentSettingCategoryIdValue", "deviceManagementSettingDefinitionIdValue")

	if id.DeviceManagementIntentId != "deviceManagementIntentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentId'", id.DeviceManagementIntentId, "deviceManagementIntentIdValue")
	}

	if id.DeviceManagementIntentSettingCategoryId != "deviceManagementIntentSettingCategoryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentSettingCategoryId'", id.DeviceManagementIntentSettingCategoryId, "deviceManagementIntentSettingCategoryIdValue")
	}

	if id.DeviceManagementSettingDefinitionId != "deviceManagementSettingDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementSettingDefinitionId'", id.DeviceManagementSettingDefinitionId, "deviceManagementSettingDefinitionIdValue")
	}
}

func TestFormatDeviceManagementIntentIdCategoryIdSettingDefinitionID(t *testing.T) {
	actual := NewDeviceManagementIntentIdCategoryIdSettingDefinitionID("deviceManagementIntentIdValue", "deviceManagementIntentSettingCategoryIdValue", "deviceManagementSettingDefinitionIdValue").ID()
	expected := "/deviceManagement/intents/deviceManagementIntentIdValue/categories/deviceManagementIntentSettingCategoryIdValue/settingDefinitions/deviceManagementSettingDefinitionIdValue"
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
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/categories",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/categories/deviceManagementIntentSettingCategoryIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/categories/deviceManagementIntentSettingCategoryIdValue/settingDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/categories/deviceManagementIntentSettingCategoryIdValue/settingDefinitions/deviceManagementSettingDefinitionIdValue",
			Expected: &DeviceManagementIntentIdCategoryIdSettingDefinitionId{
				DeviceManagementIntentId:                "deviceManagementIntentIdValue",
				DeviceManagementIntentSettingCategoryId: "deviceManagementIntentSettingCategoryIdValue",
				DeviceManagementSettingDefinitionId:     "deviceManagementSettingDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/categories/deviceManagementIntentSettingCategoryIdValue/settingDefinitions/deviceManagementSettingDefinitionIdValue/extra",
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
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/categories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/cAtEgOrIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/categories/deviceManagementIntentSettingCategoryIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/cAtEgOrIeS/dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/categories/deviceManagementIntentSettingCategoryIdValue/settingDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/cAtEgOrIeS/dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyIdVaLuE/sEtTiNgDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/categories/deviceManagementIntentSettingCategoryIdValue/settingDefinitions/deviceManagementSettingDefinitionIdValue",
			Expected: &DeviceManagementIntentIdCategoryIdSettingDefinitionId{
				DeviceManagementIntentId:                "deviceManagementIntentIdValue",
				DeviceManagementIntentSettingCategoryId: "deviceManagementIntentSettingCategoryIdValue",
				DeviceManagementSettingDefinitionId:     "deviceManagementSettingDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/categories/deviceManagementIntentSettingCategoryIdValue/settingDefinitions/deviceManagementSettingDefinitionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/cAtEgOrIeS/dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyIdVaLuE/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTsEtTiNgDeFiNiTiOnIdVaLuE",
			Expected: &DeviceManagementIntentIdCategoryIdSettingDefinitionId{
				DeviceManagementIntentId:                "dEvIcEmAnAgEmEnTiNtEnTiDvAlUe",
				DeviceManagementIntentSettingCategoryId: "dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyIdVaLuE",
				DeviceManagementSettingDefinitionId:     "dEvIcEmAnAgEmEnTsEtTiNgDeFiNiTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/cAtEgOrIeS/dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyIdVaLuE/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTsEtTiNgDeFiNiTiOnIdVaLuE/extra",
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
