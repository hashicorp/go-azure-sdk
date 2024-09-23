package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementIntentIdCategoryId{}

func TestNewDeviceManagementIntentIdCategoryID(t *testing.T) {
	id := NewDeviceManagementIntentIdCategoryID("deviceManagementIntentId", "deviceManagementIntentSettingCategoryId")

	if id.DeviceManagementIntentId != "deviceManagementIntentId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentId'", id.DeviceManagementIntentId, "deviceManagementIntentId")
	}

	if id.DeviceManagementIntentSettingCategoryId != "deviceManagementIntentSettingCategoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementIntentSettingCategoryId'", id.DeviceManagementIntentSettingCategoryId, "deviceManagementIntentSettingCategoryId")
	}
}

func TestFormatDeviceManagementIntentIdCategoryID(t *testing.T) {
	actual := NewDeviceManagementIntentIdCategoryID("deviceManagementIntentId", "deviceManagementIntentSettingCategoryId").ID()
	expected := "/deviceManagement/intents/deviceManagementIntentId/categories/deviceManagementIntentSettingCategoryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementIntentIdCategoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntentIdCategoryId
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
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/categories/deviceManagementIntentSettingCategoryId",
			Expected: &DeviceManagementIntentIdCategoryId{
				DeviceManagementIntentId:                "deviceManagementIntentId",
				DeviceManagementIntentSettingCategoryId: "deviceManagementIntentSettingCategoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentId/categories/deviceManagementIntentSettingCategoryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntentIdCategoryID(v.Input)
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

	}
}

func TestParseDeviceManagementIntentIdCategoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementIntentIdCategoryId
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
			// Valid URI
			Input: "/deviceManagement/intents/deviceManagementIntentId/categories/deviceManagementIntentSettingCategoryId",
			Expected: &DeviceManagementIntentIdCategoryId{
				DeviceManagementIntentId:                "deviceManagementIntentId",
				DeviceManagementIntentSettingCategoryId: "deviceManagementIntentSettingCategoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/intents/deviceManagementIntentId/categories/deviceManagementIntentSettingCategoryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiD/cAtEgOrIeS/dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyId",
			Expected: &DeviceManagementIntentIdCategoryId{
				DeviceManagementIntentId:                "dEvIcEmAnAgEmEnTiNtEnTiD",
				DeviceManagementIntentSettingCategoryId: "dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iNtEnTs/dEvIcEmAnAgEmEnTiNtEnTiD/cAtEgOrIeS/dEvIcEmAnAgEmEnTiNtEnTsEtTiNgCaTeGoRyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementIntentIdCategoryIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementIntentIdCategoryId(t *testing.T) {
	segments := DeviceManagementIntentIdCategoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementIntentIdCategoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
