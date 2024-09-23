package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryDeviceLocalCredentialId{}

func TestNewDirectoryDeviceLocalCredentialID(t *testing.T) {
	id := NewDirectoryDeviceLocalCredentialID("deviceLocalCredentialInfoId")

	if id.DeviceLocalCredentialInfoId != "deviceLocalCredentialInfoId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceLocalCredentialInfoId'", id.DeviceLocalCredentialInfoId, "deviceLocalCredentialInfoId")
	}
}

func TestFormatDirectoryDeviceLocalCredentialID(t *testing.T) {
	actual := NewDirectoryDeviceLocalCredentialID("deviceLocalCredentialInfoId").ID()
	expected := "/directory/deviceLocalCredentials/deviceLocalCredentialInfoId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryDeviceLocalCredentialID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryDeviceLocalCredentialId
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
			Input: "/directory/deviceLocalCredentials",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/deviceLocalCredentials/deviceLocalCredentialInfoId",
			Expected: &DirectoryDeviceLocalCredentialId{
				DeviceLocalCredentialInfoId: "deviceLocalCredentialInfoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/deviceLocalCredentials/deviceLocalCredentialInfoId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryDeviceLocalCredentialID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceLocalCredentialInfoId != v.Expected.DeviceLocalCredentialInfoId {
			t.Fatalf("Expected %q but got %q for DeviceLocalCredentialInfoId", v.Expected.DeviceLocalCredentialInfoId, actual.DeviceLocalCredentialInfoId)
		}

	}
}

func TestParseDirectoryDeviceLocalCredentialIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryDeviceLocalCredentialId
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
			Input: "/directory/deviceLocalCredentials",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/dEvIcElOcAlCrEdEnTiAlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/deviceLocalCredentials/deviceLocalCredentialInfoId",
			Expected: &DirectoryDeviceLocalCredentialId{
				DeviceLocalCredentialInfoId: "deviceLocalCredentialInfoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/deviceLocalCredentials/deviceLocalCredentialInfoId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/dEvIcElOcAlCrEdEnTiAlS/dEvIcElOcAlCrEdEnTiAlInFoId",
			Expected: &DirectoryDeviceLocalCredentialId{
				DeviceLocalCredentialInfoId: "dEvIcElOcAlCrEdEnTiAlInFoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/dEvIcElOcAlCrEdEnTiAlS/dEvIcElOcAlCrEdEnTiAlInFoId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryDeviceLocalCredentialIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceLocalCredentialInfoId != v.Expected.DeviceLocalCredentialInfoId {
			t.Fatalf("Expected %q but got %q for DeviceLocalCredentialInfoId", v.Expected.DeviceLocalCredentialInfoId, actual.DeviceLocalCredentialInfoId)
		}

	}
}

func TestSegmentsForDirectoryDeviceLocalCredentialId(t *testing.T) {
	segments := DirectoryDeviceLocalCredentialId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryDeviceLocalCredentialId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
