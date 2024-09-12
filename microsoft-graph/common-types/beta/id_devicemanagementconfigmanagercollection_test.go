package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigManagerCollectionId{}

func TestNewDeviceManagementConfigManagerCollectionID(t *testing.T) {
	id := NewDeviceManagementConfigManagerCollectionID("configManagerCollectionIdValue")

	if id.ConfigManagerCollectionId != "configManagerCollectionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConfigManagerCollectionId'", id.ConfigManagerCollectionId, "configManagerCollectionIdValue")
	}
}

func TestFormatDeviceManagementConfigManagerCollectionID(t *testing.T) {
	actual := NewDeviceManagementConfigManagerCollectionID("configManagerCollectionIdValue").ID()
	expected := "/deviceManagement/configManagerCollections/configManagerCollectionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementConfigManagerCollectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigManagerCollectionId
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
			Input: "/deviceManagement/configManagerCollections",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configManagerCollections/configManagerCollectionIdValue",
			Expected: &DeviceManagementConfigManagerCollectionId{
				ConfigManagerCollectionId: "configManagerCollectionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configManagerCollections/configManagerCollectionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigManagerCollectionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ConfigManagerCollectionId != v.Expected.ConfigManagerCollectionId {
			t.Fatalf("Expected %q but got %q for ConfigManagerCollectionId", v.Expected.ConfigManagerCollectionId, actual.ConfigManagerCollectionId)
		}

	}
}

func TestParseDeviceManagementConfigManagerCollectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementConfigManagerCollectionId
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
			Input: "/deviceManagement/configManagerCollections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGmAnAgErCoLlEcTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/configManagerCollections/configManagerCollectionIdValue",
			Expected: &DeviceManagementConfigManagerCollectionId{
				ConfigManagerCollectionId: "configManagerCollectionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/configManagerCollections/configManagerCollectionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGmAnAgErCoLlEcTiOnS/cOnFiGmAnAgErCoLlEcTiOnIdVaLuE",
			Expected: &DeviceManagementConfigManagerCollectionId{
				ConfigManagerCollectionId: "cOnFiGmAnAgErCoLlEcTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOnFiGmAnAgErCoLlEcTiOnS/cOnFiGmAnAgErCoLlEcTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementConfigManagerCollectionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ConfigManagerCollectionId != v.Expected.ConfigManagerCollectionId {
			t.Fatalf("Expected %q but got %q for ConfigManagerCollectionId", v.Expected.ConfigManagerCollectionId, actual.ConfigManagerCollectionId)
		}

	}
}

func TestSegmentsForDeviceManagementConfigManagerCollectionId(t *testing.T) {
	segments := DeviceManagementConfigManagerCollectionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementConfigManagerCollectionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
