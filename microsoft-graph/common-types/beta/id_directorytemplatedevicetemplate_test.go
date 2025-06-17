package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryTemplateDeviceTemplateId{}

func TestNewDirectoryTemplateDeviceTemplateID(t *testing.T) {
	id := NewDirectoryTemplateDeviceTemplateID("deviceTemplateId")

	if id.DeviceTemplateId != "deviceTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceTemplateId'", id.DeviceTemplateId, "deviceTemplateId")
	}
}

func TestFormatDirectoryTemplateDeviceTemplateID(t *testing.T) {
	actual := NewDirectoryTemplateDeviceTemplateID("deviceTemplateId").ID()
	expected := "/directory/templates/deviceTemplates/deviceTemplateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryTemplateDeviceTemplateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryTemplateDeviceTemplateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/templates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/templates/deviceTemplates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/templates/deviceTemplates/deviceTemplateId",
			Expected: &DirectoryTemplateDeviceTemplateId{
				DeviceTemplateId: "deviceTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/templates/deviceTemplates/deviceTemplateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryTemplateDeviceTemplateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceTemplateId != v.Expected.DeviceTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceTemplateId", v.Expected.DeviceTemplateId, actual.DeviceTemplateId)
		}

	}
}

func TestParseDirectoryTemplateDeviceTemplateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryTemplateDeviceTemplateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/templates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/tEmPlAtEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/templates/deviceTemplates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/tEmPlAtEs/dEvIcEtEmPlAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/templates/deviceTemplates/deviceTemplateId",
			Expected: &DirectoryTemplateDeviceTemplateId{
				DeviceTemplateId: "deviceTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/templates/deviceTemplates/deviceTemplateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/tEmPlAtEs/dEvIcEtEmPlAtEs/dEvIcEtEmPlAtEiD",
			Expected: &DirectoryTemplateDeviceTemplateId{
				DeviceTemplateId: "dEvIcEtEmPlAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/tEmPlAtEs/dEvIcEtEmPlAtEs/dEvIcEtEmPlAtEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryTemplateDeviceTemplateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceTemplateId != v.Expected.DeviceTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceTemplateId", v.Expected.DeviceTemplateId, actual.DeviceTemplateId)
		}

	}
}

func TestSegmentsForDirectoryTemplateDeviceTemplateId(t *testing.T) {
	segments := DirectoryTemplateDeviceTemplateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryTemplateDeviceTemplateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
