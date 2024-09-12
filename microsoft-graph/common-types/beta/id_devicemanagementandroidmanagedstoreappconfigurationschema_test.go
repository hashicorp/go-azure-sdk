package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementAndroidManagedStoreAppConfigurationSchemaId{}

func TestNewDeviceManagementAndroidManagedStoreAppConfigurationSchemaID(t *testing.T) {
	id := NewDeviceManagementAndroidManagedStoreAppConfigurationSchemaID("androidManagedStoreAppConfigurationSchemaIdValue")

	if id.AndroidManagedStoreAppConfigurationSchemaId != "androidManagedStoreAppConfigurationSchemaIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AndroidManagedStoreAppConfigurationSchemaId'", id.AndroidManagedStoreAppConfigurationSchemaId, "androidManagedStoreAppConfigurationSchemaIdValue")
	}
}

func TestFormatDeviceManagementAndroidManagedStoreAppConfigurationSchemaID(t *testing.T) {
	actual := NewDeviceManagementAndroidManagedStoreAppConfigurationSchemaID("androidManagedStoreAppConfigurationSchemaIdValue").ID()
	expected := "/deviceManagement/androidManagedStoreAppConfigurationSchemas/androidManagedStoreAppConfigurationSchemaIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementAndroidManagedStoreAppConfigurationSchemaID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAndroidManagedStoreAppConfigurationSchemaId
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
			Input: "/deviceManagement/androidManagedStoreAppConfigurationSchemas",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/androidManagedStoreAppConfigurationSchemas/androidManagedStoreAppConfigurationSchemaIdValue",
			Expected: &DeviceManagementAndroidManagedStoreAppConfigurationSchemaId{
				AndroidManagedStoreAppConfigurationSchemaId: "androidManagedStoreAppConfigurationSchemaIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/androidManagedStoreAppConfigurationSchemas/androidManagedStoreAppConfigurationSchemaIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAndroidManagedStoreAppConfigurationSchemaID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AndroidManagedStoreAppConfigurationSchemaId != v.Expected.AndroidManagedStoreAppConfigurationSchemaId {
			t.Fatalf("Expected %q but got %q for AndroidManagedStoreAppConfigurationSchemaId", v.Expected.AndroidManagedStoreAppConfigurationSchemaId, actual.AndroidManagedStoreAppConfigurationSchemaId)
		}

	}
}

func TestParseDeviceManagementAndroidManagedStoreAppConfigurationSchemaIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementAndroidManagedStoreAppConfigurationSchemaId
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
			Input: "/deviceManagement/androidManagedStoreAppConfigurationSchemas",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aNdRoIdMaNaGeDsToReApPcOnFiGuRaTiOnScHeMaS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/androidManagedStoreAppConfigurationSchemas/androidManagedStoreAppConfigurationSchemaIdValue",
			Expected: &DeviceManagementAndroidManagedStoreAppConfigurationSchemaId{
				AndroidManagedStoreAppConfigurationSchemaId: "androidManagedStoreAppConfigurationSchemaIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/androidManagedStoreAppConfigurationSchemas/androidManagedStoreAppConfigurationSchemaIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aNdRoIdMaNaGeDsToReApPcOnFiGuRaTiOnScHeMaS/aNdRoIdMaNaGeDsToReApPcOnFiGuRaTiOnScHeMaIdVaLuE",
			Expected: &DeviceManagementAndroidManagedStoreAppConfigurationSchemaId{
				AndroidManagedStoreAppConfigurationSchemaId: "aNdRoIdMaNaGeDsToReApPcOnFiGuRaTiOnScHeMaIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/aNdRoIdMaNaGeDsToReApPcOnFiGuRaTiOnScHeMaS/aNdRoIdMaNaGeDsToReApPcOnFiGuRaTiOnScHeMaIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementAndroidManagedStoreAppConfigurationSchemaIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AndroidManagedStoreAppConfigurationSchemaId != v.Expected.AndroidManagedStoreAppConfigurationSchemaId {
			t.Fatalf("Expected %q but got %q for AndroidManagedStoreAppConfigurationSchemaId", v.Expected.AndroidManagedStoreAppConfigurationSchemaId, actual.AndroidManagedStoreAppConfigurationSchemaId)
		}

	}
}

func TestSegmentsForDeviceManagementAndroidManagedStoreAppConfigurationSchemaId(t *testing.T) {
	segments := DeviceManagementAndroidManagedStoreAppConfigurationSchemaId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementAndroidManagedStoreAppConfigurationSchemaId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
