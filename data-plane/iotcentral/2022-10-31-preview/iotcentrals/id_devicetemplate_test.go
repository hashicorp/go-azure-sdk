package iotcentrals

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceTemplateId{}

func TestNewDeviceTemplateID(t *testing.T) {
	id := NewDeviceTemplateID("https://endpoint-url.example.com", "deviceTemplateId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.DeviceTemplateId != "deviceTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceTemplateId'", id.DeviceTemplateId, "deviceTemplateId")
	}
}

func TestFormatDeviceTemplateID(t *testing.T) {
	actual := NewDeviceTemplateID("https://endpoint-url.example.com", "deviceTemplateId").ID()
	expected := "https://endpoint-url.example.com/deviceTemplates/deviceTemplateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceTemplateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceTemplateId
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
			Input: "https://endpoint-url.example.com/deviceTemplates",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/deviceTemplates/deviceTemplateId",
			Expected: &DeviceTemplateId{
				BaseURI:          "https://endpoint-url.example.com",
				DeviceTemplateId: "deviceTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/deviceTemplates/deviceTemplateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceTemplateID(v.Input)
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

		if actual.DeviceTemplateId != v.Expected.DeviceTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceTemplateId", v.Expected.DeviceTemplateId, actual.DeviceTemplateId)
		}

	}
}

func TestParseDeviceTemplateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceTemplateId
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
			Input: "https://endpoint-url.example.com/deviceTemplates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEtEmPlAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/deviceTemplates/deviceTemplateId",
			Expected: &DeviceTemplateId{
				BaseURI:          "https://endpoint-url.example.com",
				DeviceTemplateId: "deviceTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/deviceTemplates/deviceTemplateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEtEmPlAtEs/dEvIcEtEmPlAtEiD",
			Expected: &DeviceTemplateId{
				BaseURI:          "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				DeviceTemplateId: "dEvIcEtEmPlAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEtEmPlAtEs/dEvIcEtEmPlAtEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceTemplateIDInsensitively(v.Input)
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

		if actual.DeviceTemplateId != v.Expected.DeviceTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceTemplateId", v.Expected.DeviceTemplateId, actual.DeviceTemplateId)
		}

	}
}

func TestSegmentsForDeviceTemplateId(t *testing.T) {
	segments := DeviceTemplateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceTemplateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
