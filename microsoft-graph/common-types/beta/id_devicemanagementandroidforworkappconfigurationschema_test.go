package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAndroidForWorkAppConfigurationSchemaId{}

func TestNewDeviceManagementAndroidForWorkAppConfigurationSchemaID(t *testing.T) {
	id := NewDeviceManagementAndroidForWorkAppConfigurationSchemaID("androidForWorkAppConfigurationSchemaIdValue")

	if id.AndroidForWorkAppConfigurationSchemaId != "androidForWorkAppConfigurationSchemaIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AndroidForWorkAppConfigurationSchemaId'", id.AndroidForWorkAppConfigurationSchemaId, "androidForWorkAppConfigurationSchemaIdValue")
	}
}

func TestFormatDeviceManagementAndroidForWorkAppConfigurationSchemaID(t *testing.T) {
	actual := NewDeviceManagementAndroidForWorkAppConfigurationSchemaID("androidForWorkAppConfigurationSchemaIdValue").ID()
	expected := "/deviceManagement/androidForWorkAppConfigurationSchemas/androidForWorkAppConfigurationSchemaIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementAndroidForWorkAppConfigurationSchemaID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAndroidForWorkAppConfigurationSchemaId
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
			Input: "/deviceManagement/androidForWorkAppConfigurationSchemas",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/androidForWorkAppConfigurationSchemas/androidForWorkAppConfigurationSchemaIdValue",
			Expected: &DeviceManagementAndroidForWorkAppConfigurationSchemaId{
				AndroidForWorkAppConfigurationSchemaId: "androidForWorkAppConfigurationSchemaIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/androidForWorkAppConfigurationSchemas/androidForWorkAppConfigurationSchemaIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAndroidForWorkAppConfigurationSchemaID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AndroidForWorkAppConfigurationSchemaId != v.Expected.AndroidForWorkAppConfigurationSchemaId {
			t.Fatalf("Expected %q but got %q for AndroidForWorkAppConfigurationSchemaId", v.Expected.AndroidForWorkAppConfigurationSchemaId, actual.AndroidForWorkAppConfigurationSchemaId)
		}

	}
}

func TestParseDeviceManagementAndroidForWorkAppConfigurationSchemaIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAndroidForWorkAppConfigurationSchemaId
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
			Input: "/deviceManagement/androidForWorkAppConfigurationSchemas",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aNdRoIdFoRwOrKaPpCoNfIgUrAtIoNsChEmAs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/androidForWorkAppConfigurationSchemas/androidForWorkAppConfigurationSchemaIdValue",
			Expected: &DeviceManagementAndroidForWorkAppConfigurationSchemaId{
				AndroidForWorkAppConfigurationSchemaId: "androidForWorkAppConfigurationSchemaIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/androidForWorkAppConfigurationSchemas/androidForWorkAppConfigurationSchemaIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aNdRoIdFoRwOrKaPpCoNfIgUrAtIoNsChEmAs/aNdRoIdFoRwOrKaPpCoNfIgUrAtIoNsChEmAiDvAlUe",
			Expected: &DeviceManagementAndroidForWorkAppConfigurationSchemaId{
				AndroidForWorkAppConfigurationSchemaId: "aNdRoIdFoRwOrKaPpCoNfIgUrAtIoNsChEmAiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aNdRoIdFoRwOrKaPpCoNfIgUrAtIoNsChEmAs/aNdRoIdFoRwOrKaPpCoNfIgUrAtIoNsChEmAiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAndroidForWorkAppConfigurationSchemaIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AndroidForWorkAppConfigurationSchemaId != v.Expected.AndroidForWorkAppConfigurationSchemaId {
			t.Fatalf("Expected %q but got %q for AndroidForWorkAppConfigurationSchemaId", v.Expected.AndroidForWorkAppConfigurationSchemaId, actual.AndroidForWorkAppConfigurationSchemaId)
		}

	}
}

func TestSegmentsForDeviceManagementAndroidForWorkAppConfigurationSchemaId(t *testing.T) {
	segments := DeviceManagementAndroidForWorkAppConfigurationSchemaId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementAndroidForWorkAppConfigurationSchemaId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
