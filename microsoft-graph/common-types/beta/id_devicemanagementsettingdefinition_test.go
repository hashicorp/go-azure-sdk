package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementSettingDefinitionId{}

func TestNewDeviceManagementSettingDefinitionID(t *testing.T) {
	id := NewDeviceManagementSettingDefinitionID("deviceManagementSettingDefinitionId")

	if id.DeviceManagementSettingDefinitionId != "deviceManagementSettingDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementSettingDefinitionId'", id.DeviceManagementSettingDefinitionId, "deviceManagementSettingDefinitionId")
	}
}

func TestFormatDeviceManagementSettingDefinitionID(t *testing.T) {
	actual := NewDeviceManagementSettingDefinitionID("deviceManagementSettingDefinitionId").ID()
	expected := "/deviceManagement/settingDefinitions/deviceManagementSettingDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementSettingDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementSettingDefinitionId
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
			Input: "/deviceManagement/settingDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/settingDefinitions/deviceManagementSettingDefinitionId",
			Expected: &DeviceManagementSettingDefinitionId{
				DeviceManagementSettingDefinitionId: "deviceManagementSettingDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/settingDefinitions/deviceManagementSettingDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementSettingDefinitionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementSettingDefinitionId != v.Expected.DeviceManagementSettingDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingDefinitionId", v.Expected.DeviceManagementSettingDefinitionId, actual.DeviceManagementSettingDefinitionId)
		}

	}
}

func TestParseDeviceManagementSettingDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementSettingDefinitionId
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
			Input: "/deviceManagement/settingDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/sEtTiNgDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/settingDefinitions/deviceManagementSettingDefinitionId",
			Expected: &DeviceManagementSettingDefinitionId{
				DeviceManagementSettingDefinitionId: "deviceManagementSettingDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/settingDefinitions/deviceManagementSettingDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTsEtTiNgDeFiNiTiOnId",
			Expected: &DeviceManagementSettingDefinitionId{
				DeviceManagementSettingDefinitionId: "dEvIcEmAnAgEmEnTsEtTiNgDeFiNiTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/sEtTiNgDeFiNiTiOnS/dEvIcEmAnAgEmEnTsEtTiNgDeFiNiTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementSettingDefinitionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementSettingDefinitionId != v.Expected.DeviceManagementSettingDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingDefinitionId", v.Expected.DeviceManagementSettingDefinitionId, actual.DeviceManagementSettingDefinitionId)
		}

	}
}

func TestSegmentsForDeviceManagementSettingDefinitionId(t *testing.T) {
	segments := DeviceManagementSettingDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementSettingDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
