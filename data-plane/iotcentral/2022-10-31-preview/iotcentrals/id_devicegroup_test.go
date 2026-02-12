package iotcentrals

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceGroupId{}

func TestNewDeviceGroupID(t *testing.T) {
	id := NewDeviceGroupID("https://endpoint-url.example.com", "deviceGroupId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.DeviceGroupId != "deviceGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceGroupId'", id.DeviceGroupId, "deviceGroupId")
	}
}

func TestFormatDeviceGroupID(t *testing.T) {
	actual := NewDeviceGroupID("https://endpoint-url.example.com", "deviceGroupId").ID()
	expected := "https://endpoint-url.example.com/deviceGroups/deviceGroupId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceGroupId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/deviceGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/deviceGroups/deviceGroupId",
			Expected: &DeviceGroupId{
				BaseURI:       "https://endpoint-url.example.com",
				DeviceGroupId: "deviceGroupId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/deviceGroups/deviceGroupId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceGroupID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.DeviceGroupId != v.Expected.DeviceGroupId {
			t.Fatalf("Expected %q but got %q for DeviceGroupId", v.Expected.DeviceGroupId, actual.DeviceGroupId)
		}

	}
}

func TestParseDeviceGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceGroupId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/deviceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEgRoUpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/deviceGroups/deviceGroupId",
			Expected: &DeviceGroupId{
				BaseURI:       "https://endpoint-url.example.com",
				DeviceGroupId: "deviceGroupId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/deviceGroups/deviceGroupId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEgRoUpS/dEvIcEgRoUpId",
			Expected: &DeviceGroupId{
				BaseURI:       "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				DeviceGroupId: "dEvIcEgRoUpId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEgRoUpS/dEvIcEgRoUpId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceGroupIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.DeviceGroupId != v.Expected.DeviceGroupId {
			t.Fatalf("Expected %q but got %q for DeviceGroupId", v.Expected.DeviceGroupId, actual.DeviceGroupId)
		}

	}
}

func TestSegmentsForDeviceGroupId(t *testing.T) {
	segments := DeviceGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
