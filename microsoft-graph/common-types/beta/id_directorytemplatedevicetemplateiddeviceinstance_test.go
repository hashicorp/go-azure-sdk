package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryTemplateDeviceTemplateIdDeviceInstanceId{}

func TestNewDirectoryTemplateDeviceTemplateIdDeviceInstanceID(t *testing.T) {
	id := NewDirectoryTemplateDeviceTemplateIdDeviceInstanceID("deviceTemplateId", "deviceId")

	if id.DeviceTemplateId != "deviceTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceTemplateId'", id.DeviceTemplateId, "deviceTemplateId")
	}

	if id.DeviceId != "deviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceId'", id.DeviceId, "deviceId")
	}
}

func TestFormatDirectoryTemplateDeviceTemplateIdDeviceInstanceID(t *testing.T) {
	actual := NewDirectoryTemplateDeviceTemplateIdDeviceInstanceID("deviceTemplateId", "deviceId").ID()
	expected := "/directory/templates/deviceTemplates/deviceTemplateId/deviceInstances/deviceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryTemplateDeviceTemplateIdDeviceInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryTemplateDeviceTemplateIdDeviceInstanceId
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
			// Incomplete URI
			Input: "/directory/templates/deviceTemplates/deviceTemplateId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/templates/deviceTemplates/deviceTemplateId/deviceInstances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/templates/deviceTemplates/deviceTemplateId/deviceInstances/deviceId",
			Expected: &DirectoryTemplateDeviceTemplateIdDeviceInstanceId{
				DeviceTemplateId: "deviceTemplateId",
				DeviceId:         "deviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/templates/deviceTemplates/deviceTemplateId/deviceInstances/deviceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryTemplateDeviceTemplateIdDeviceInstanceID(v.Input)
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

		if actual.DeviceId != v.Expected.DeviceId {
			t.Fatalf("Expected %q but got %q for DeviceId", v.Expected.DeviceId, actual.DeviceId)
		}

	}
}

func TestParseDirectoryTemplateDeviceTemplateIdDeviceInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryTemplateDeviceTemplateIdDeviceInstanceId
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
			// Incomplete URI
			Input: "/directory/templates/deviceTemplates/deviceTemplateId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/tEmPlAtEs/dEvIcEtEmPlAtEs/dEvIcEtEmPlAtEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/templates/deviceTemplates/deviceTemplateId/deviceInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/tEmPlAtEs/dEvIcEtEmPlAtEs/dEvIcEtEmPlAtEiD/dEvIcEiNsTaNcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/templates/deviceTemplates/deviceTemplateId/deviceInstances/deviceId",
			Expected: &DirectoryTemplateDeviceTemplateIdDeviceInstanceId{
				DeviceTemplateId: "deviceTemplateId",
				DeviceId:         "deviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/templates/deviceTemplates/deviceTemplateId/deviceInstances/deviceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/tEmPlAtEs/dEvIcEtEmPlAtEs/dEvIcEtEmPlAtEiD/dEvIcEiNsTaNcEs/dEvIcEiD",
			Expected: &DirectoryTemplateDeviceTemplateIdDeviceInstanceId{
				DeviceTemplateId: "dEvIcEtEmPlAtEiD",
				DeviceId:         "dEvIcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/tEmPlAtEs/dEvIcEtEmPlAtEs/dEvIcEtEmPlAtEiD/dEvIcEiNsTaNcEs/dEvIcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryTemplateDeviceTemplateIdDeviceInstanceIDInsensitively(v.Input)
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

		if actual.DeviceId != v.Expected.DeviceId {
			t.Fatalf("Expected %q but got %q for DeviceId", v.Expected.DeviceId, actual.DeviceId)
		}

	}
}

func TestSegmentsForDirectoryTemplateDeviceTemplateIdDeviceInstanceId(t *testing.T) {
	segments := DirectoryTemplateDeviceTemplateIdDeviceInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryTemplateDeviceTemplateIdDeviceInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
