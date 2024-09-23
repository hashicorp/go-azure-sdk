package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComplianceCategoryId{}

func TestNewDeviceManagementComplianceCategoryID(t *testing.T) {
	id := NewDeviceManagementComplianceCategoryID("deviceManagementConfigurationCategoryId")

	if id.DeviceManagementConfigurationCategoryId != "deviceManagementConfigurationCategoryId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementConfigurationCategoryId'", id.DeviceManagementConfigurationCategoryId, "deviceManagementConfigurationCategoryId")
	}
}

func TestFormatDeviceManagementComplianceCategoryID(t *testing.T) {
	actual := NewDeviceManagementComplianceCategoryID("deviceManagementConfigurationCategoryId").ID()
	expected := "/deviceManagement/complianceCategories/deviceManagementConfigurationCategoryId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementComplianceCategoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementComplianceCategoryId
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
			Input: "/deviceManagement/complianceCategories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/complianceCategories/deviceManagementConfigurationCategoryId",
			Expected: &DeviceManagementComplianceCategoryId{
				DeviceManagementConfigurationCategoryId: "deviceManagementConfigurationCategoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/complianceCategories/deviceManagementConfigurationCategoryId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementComplianceCategoryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationCategoryId != v.Expected.DeviceManagementConfigurationCategoryId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationCategoryId", v.Expected.DeviceManagementConfigurationCategoryId, actual.DeviceManagementConfigurationCategoryId)
		}

	}
}

func TestParseDeviceManagementComplianceCategoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementComplianceCategoryId
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
			Input: "/deviceManagement/complianceCategories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEcAtEgOrIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/complianceCategories/deviceManagementConfigurationCategoryId",
			Expected: &DeviceManagementComplianceCategoryId{
				DeviceManagementConfigurationCategoryId: "deviceManagementConfigurationCategoryId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/complianceCategories/deviceManagementConfigurationCategoryId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEcAtEgOrIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnCaTeGoRyId",
			Expected: &DeviceManagementComplianceCategoryId{
				DeviceManagementConfigurationCategoryId: "dEvIcEmAnAgEmEnTcOnFiGuRaTiOnCaTeGoRyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEcAtEgOrIeS/dEvIcEmAnAgEmEnTcOnFiGuRaTiOnCaTeGoRyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementComplianceCategoryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementConfigurationCategoryId != v.Expected.DeviceManagementConfigurationCategoryId {
			t.Fatalf("Expected %q but got %q for DeviceManagementConfigurationCategoryId", v.Expected.DeviceManagementConfigurationCategoryId, actual.DeviceManagementConfigurationCategoryId)
		}

	}
}

func TestSegmentsForDeviceManagementComplianceCategoryId(t *testing.T) {
	segments := DeviceManagementComplianceCategoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementComplianceCategoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
