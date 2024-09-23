package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdCategoryId{}

func TestNewDeviceManagementTemplateIdCategoryID(t *testing.T) {
	id := NewDeviceManagementTemplateIdCategoryID("deviceManagementTemplateId", "deviceManagementTemplateSettingCategoryId")

	if id.DeviceManagementTemplateId != "deviceManagementTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateId'", id.DeviceManagementTemplateId, "deviceManagementTemplateId")
	}

	if id.DeviceManagementTemplateSettingCategoryId != "deviceManagementTemplateSettingCategoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateSettingCategoryId'", id.DeviceManagementTemplateSettingCategoryId, "deviceManagementTemplateSettingCategoryId")
	}
}

func TestFormatDeviceManagementTemplateIdCategoryID(t *testing.T) {
	actual := NewDeviceManagementTemplateIdCategoryID("deviceManagementTemplateId", "deviceManagementTemplateSettingCategoryId").ID()
	expected := "/deviceManagement/templates/deviceManagementTemplateId/categories/deviceManagementTemplateSettingCategoryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementTemplateIdCategoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateIdCategoryId
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
			Input: "/deviceManagement/templates/deviceManagementTemplateId/categories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/categories/deviceManagementTemplateSettingCategoryId",
			Expected: &DeviceManagementTemplateIdCategoryId{
				DeviceManagementTemplateId:                "deviceManagementTemplateId",
				DeviceManagementTemplateSettingCategoryId: "deviceManagementTemplateSettingCategoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templates/deviceManagementTemplateId/categories/deviceManagementTemplateSettingCategoryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateIdCategoryID(v.Input)
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

		if actual.DeviceManagementTemplateSettingCategoryId != v.Expected.DeviceManagementTemplateSettingCategoryId {
			t.Fatalf("Expected %q but got %q for DeviceManagementTemplateSettingCategoryId", v.Expected.DeviceManagementTemplateSettingCategoryId, actual.DeviceManagementTemplateSettingCategoryId)
		}

	}
}

func TestParseDeviceManagementTemplateIdCategoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateIdCategoryId
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
			Input: "/deviceManagement/templates/deviceManagementTemplateId/categories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/cAtEgOrIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templates/deviceManagementTemplateId/categories/deviceManagementTemplateSettingCategoryId",
			Expected: &DeviceManagementTemplateIdCategoryId{
				DeviceManagementTemplateId:                "deviceManagementTemplateId",
				DeviceManagementTemplateSettingCategoryId: "deviceManagementTemplateSettingCategoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templates/deviceManagementTemplateId/categories/deviceManagementTemplateSettingCategoryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyId",
			Expected: &DeviceManagementTemplateIdCategoryId{
				DeviceManagementTemplateId:                "dEvIcEmAnAgEmEnTtEmPlAtEiD",
				DeviceManagementTemplateSettingCategoryId: "dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiD/cAtEgOrIeS/dEvIcEmAnAgEmEnTtEmPlAtEsEtTiNgCaTeGoRyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateIdCategoryIDInsensitively(v.Input)
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

		if actual.DeviceManagementTemplateSettingCategoryId != v.Expected.DeviceManagementTemplateSettingCategoryId {
			t.Fatalf("Expected %q but got %q for DeviceManagementTemplateSettingCategoryId", v.Expected.DeviceManagementTemplateSettingCategoryId, actual.DeviceManagementTemplateSettingCategoryId)
		}

	}
}

func TestSegmentsForDeviceManagementTemplateIdCategoryId(t *testing.T) {
	segments := DeviceManagementTemplateIdCategoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementTemplateIdCategoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
