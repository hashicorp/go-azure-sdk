package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementResourceAccessProfileId{}

func TestNewDeviceManagementResourceAccessProfileID(t *testing.T) {
	id := NewDeviceManagementResourceAccessProfileID("deviceManagementResourceAccessProfileBaseIdValue")

	if id.DeviceManagementResourceAccessProfileBaseId != "deviceManagementResourceAccessProfileBaseIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementResourceAccessProfileBaseId'", id.DeviceManagementResourceAccessProfileBaseId, "deviceManagementResourceAccessProfileBaseIdValue")
	}
}

func TestFormatDeviceManagementResourceAccessProfileID(t *testing.T) {
	actual := NewDeviceManagementResourceAccessProfileID("deviceManagementResourceAccessProfileBaseIdValue").ID()
	expected := "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementResourceAccessProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementResourceAccessProfileId
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
			Input: "/deviceManagement/resourceAccessProfiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseIdValue",
			Expected: &DeviceManagementResourceAccessProfileId{
				DeviceManagementResourceAccessProfileBaseId: "deviceManagementResourceAccessProfileBaseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementResourceAccessProfileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementResourceAccessProfileBaseId != v.Expected.DeviceManagementResourceAccessProfileBaseId {
			t.Fatalf("Expected %q but got %q for DeviceManagementResourceAccessProfileBaseId", v.Expected.DeviceManagementResourceAccessProfileBaseId, actual.DeviceManagementResourceAccessProfileBaseId)
		}

	}
}

func TestParseDeviceManagementResourceAccessProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementResourceAccessProfileId
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
			Input: "/deviceManagement/resourceAccessProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEsOuRcEaCcEsSpRoFiLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseIdValue",
			Expected: &DeviceManagementResourceAccessProfileId{
				DeviceManagementResourceAccessProfileBaseId: "deviceManagementResourceAccessProfileBaseIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/resourceAccessProfiles/deviceManagementResourceAccessProfileBaseIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEsOuRcEaCcEsSpRoFiLeS/dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeBaSeIdVaLuE",
			Expected: &DeviceManagementResourceAccessProfileId{
				DeviceManagementResourceAccessProfileBaseId: "dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeBaSeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/rEsOuRcEaCcEsSpRoFiLeS/dEvIcEmAnAgEmEnTrEsOuRcEaCcEsSpRoFiLeBaSeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementResourceAccessProfileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementResourceAccessProfileBaseId != v.Expected.DeviceManagementResourceAccessProfileBaseId {
			t.Fatalf("Expected %q but got %q for DeviceManagementResourceAccessProfileBaseId", v.Expected.DeviceManagementResourceAccessProfileBaseId, actual.DeviceManagementResourceAccessProfileBaseId)
		}

	}
}

func TestSegmentsForDeviceManagementResourceAccessProfileId(t *testing.T) {
	segments := DeviceManagementResourceAccessProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementResourceAccessProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
