package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdCategoryIdSettingId{}

func TestNewDeviceManagementIntentIdCategoryIdSettingID(t *testing.T) {
	id := NewDeviceManagementIntentIdCategoryIdSettingID("deviceManagementIntentIdValue", "deviceManagementIntentSettingCategoryIdValue", "deviceManagementSettingInstanceIdValue")

	if id.DeviceManagementIntentId != "deviceManagementIntentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentId'", id.DeviceManagementIntentId, "deviceManagementIntentIdValue")
	}

	if id.DeviceManagementIntentSettingCategoryId != "deviceManagementIntentSettingCategoryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentSettingCategoryId'", id.DeviceManagementIntentSettingCategoryId, "deviceManagementIntentSettingCategoryIdValue")
	}

	if id.DeviceManagementSettingInstanceId != "deviceManagementSettingInstanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementSettingInstanceId'", id.DeviceManagementSettingInstanceId, "deviceManagementSettingInstanceIdValue")
	}
}

func TestFormatDeviceManagementIntentIdCategoryIdSettingID(t *testing.T) {
	actual := NewDeviceManagementIntentIdCategoryIdSettingID("deviceManagementIntentIdValue", "deviceManagementIntentSettingCategoryIdValue", "deviceManagementSettingInstanceIdValue").ID()
	expected := "/deviceManagement/intents/deviceManagementIntentIdValue/categories/deviceManagementIntentSettingCategoryIdValue/settings/deviceManagementSettingInstanceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementIntentIdCategoryIdSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntentIdCategoryIdSettingId
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
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/categories/deviceManagementIntentSettingCategoryIdValue/settings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/categories/deviceManagementIntentSettingCategoryIdValue/settings/deviceManagementSettingInstanceIdValue",
			Expected: &DeviceManagementIntentIdCategoryIdSettingId{
				DeviceManagementIntentId:                "deviceManagementIntentIdValue",
				DeviceManagementIntentSettingCategoryId: "deviceManagementIntentSettingCategoryIdValue",
				DeviceManagementSettingInstanceId:       "deviceManagementSettingInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/categories/deviceManagementIntentSettingCategoryIdValue/settings/deviceManagementSettingInstanceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntentIdCategoryIdSettingID(v.Input)
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

		if actual.DeviceManagementSettingInstanceId != v.Expected.DeviceManagementSettingInstanceId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingInstanceId", v.Expected.DeviceManagementSettingInstanceId, actual.DeviceManagementSettingInstanceId)
		}

	}
}

func TestParseDeviceManagementIntentIdCategoryIdSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntentIdCategoryIdSettingId
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
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/categories/deviceManagementIntentSettingCategoryIdValue/settings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/cAtEgOrIeS/dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyIdVaLuE/sEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/categories/deviceManagementIntentSettingCategoryIdValue/settings/deviceManagementSettingInstanceIdValue",
			Expected: &DeviceManagementIntentIdCategoryIdSettingId{
				DeviceManagementIntentId:                "deviceManagementIntentIdValue",
				DeviceManagementIntentSettingCategoryId: "deviceManagementIntentSettingCategoryIdValue",
				DeviceManagementSettingInstanceId:       "deviceManagementSettingInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentIdValue/categories/deviceManagementIntentSettingCategoryIdValue/settings/deviceManagementSettingInstanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/cAtEgOrIeS/dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyIdVaLuE/sEtTiNgS/dEvIcEmAnAgEmEnTsEtTiNgInStAnCeIdVaLuE",
			Expected: &DeviceManagementIntentIdCategoryIdSettingId{
				DeviceManagementIntentId:                "dEvIcEmAnAgEmEnTiNtEnTiDvAlUe",
				DeviceManagementIntentSettingCategoryId: "dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyIdVaLuE",
				DeviceManagementSettingInstanceId:       "dEvIcEmAnAgEmEnTsEtTiNgInStAnCeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiDvAlUe/cAtEgOrIeS/dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyIdVaLuE/sEtTiNgS/dEvIcEmAnAgEmEnTsEtTiNgInStAnCeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntentIdCategoryIdSettingIDInsensitively(v.Input)
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

		if actual.DeviceManagementSettingInstanceId != v.Expected.DeviceManagementSettingInstanceId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingInstanceId", v.Expected.DeviceManagementSettingInstanceId, actual.DeviceManagementSettingInstanceId)
		}

	}
}

func TestSegmentsForDeviceManagementIntentIdCategoryIdSettingId(t *testing.T) {
	segments := DeviceManagementIntentIdCategoryIdSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementIntentIdCategoryIdSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
