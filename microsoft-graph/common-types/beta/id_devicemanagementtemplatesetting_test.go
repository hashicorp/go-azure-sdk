package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateSettingId{}

func TestNewDeviceManagementTemplateSettingID(t *testing.T) {
	id := NewDeviceManagementTemplateSettingID("deviceManagementConfigurationSettingTemplateId")

	if id.DeviceManagementConfigurationSettingTemplateId != "deviceManagementConfigurationSettingTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationSettingTemplateId'", id.DeviceManagementConfigurationSettingTemplateId, "deviceManagementConfigurationSettingTemplateId")
	}
}

func TestFormatDeviceManagementTemplateSettingID(t *testing.T) {
	actual := NewDeviceManagementTemplateSettingID("deviceManagementConfigurationSettingTemplateId").ID()
	expected := "/deviceManagement/templateSettings/deviceManagementConfigurationSettingTemplateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementTemplateSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateSettingId
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
			Input: "/deviceManagement/templateSettings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templateSettings/deviceManagementConfigurationSettingTemplateId",
			Expected: &DeviceManagementTemplateSettingId{
				DeviceManagementConfigurationSettingTemplateId: "deviceManagementConfigurationSettingTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templateSettings/deviceManagementConfigurationSettingTemplateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateSettingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationSettingTemplateId != v.Expected.DeviceManagementConfigurationSettingTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingTemplateId", v.Expected.DeviceManagementConfigurationSettingTemplateId, actual.DeviceManagementConfigurationSettingTemplateId)
		}

	}
}

func TestParseDeviceManagementTemplateSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateSettingId
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
			Input: "/deviceManagement/templateSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEsEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templateSettings/deviceManagementConfigurationSettingTemplateId",
			Expected: &DeviceManagementTemplateSettingId{
				DeviceManagementConfigurationSettingTemplateId: "deviceManagementConfigurationSettingTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templateSettings/deviceManagementConfigurationSettingTemplateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEsEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiD",
			Expected: &DeviceManagementTemplateSettingId{
				DeviceManagementConfigurationSettingTemplateId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEsEtTiNgS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnSeTtInGtEmPlAtEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateSettingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationSettingTemplateId != v.Expected.DeviceManagementConfigurationSettingTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationSettingTemplateId", v.Expected.DeviceManagementConfigurationSettingTemplateId, actual.DeviceManagementConfigurationSettingTemplateId)
		}

	}
}

func TestSegmentsForDeviceManagementTemplateSettingId(t *testing.T) {
	segments := DeviceManagementTemplateSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementTemplateSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
