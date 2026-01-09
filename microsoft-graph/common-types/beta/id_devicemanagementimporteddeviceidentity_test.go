package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementImportedDeviceIdentityId{}

func TestNewDeviceManagementImportedDeviceIdentityID(t *testing.T) {
	id := NewDeviceManagementImportedDeviceIdentityID("importedDeviceIdentityId")

	if id.ImportedDeviceIdentityId != "importedDeviceIdentityId" {
		t.Fatalf("Expected %q but got %q for Segment 'ImportedDeviceIdentityId'", id.ImportedDeviceIdentityId, "importedDeviceIdentityId")
	}
}

func TestFormatDeviceManagementImportedDeviceIdentityID(t *testing.T) {
	actual := NewDeviceManagementImportedDeviceIdentityID("importedDeviceIdentityId").ID()
	expected := "/deviceManagement/importedDeviceIdentities/importedDeviceIdentityId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementImportedDeviceIdentityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementImportedDeviceIdentityId
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
			Input: "/deviceManagement/importedDeviceIdentities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/importedDeviceIdentities/importedDeviceIdentityId",
			Expected: &DeviceManagementImportedDeviceIdentityId{
				ImportedDeviceIdentityId: "importedDeviceIdentityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/importedDeviceIdentities/importedDeviceIdentityId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementImportedDeviceIdentityID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ImportedDeviceIdentityId != v.Expected.ImportedDeviceIdentityId {
			t.Fatalf("Expected %q but got %q for ImportedDeviceIdentityId", v.Expected.ImportedDeviceIdentityId, actual.ImportedDeviceIdentityId)
		}

	}
}

func TestParseDeviceManagementImportedDeviceIdentityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementImportedDeviceIdentityId
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
			Input: "/deviceManagement/importedDeviceIdentities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iMpOrTeDdEvIcEiDeNtItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/importedDeviceIdentities/importedDeviceIdentityId",
			Expected: &DeviceManagementImportedDeviceIdentityId{
				ImportedDeviceIdentityId: "importedDeviceIdentityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/importedDeviceIdentities/importedDeviceIdentityId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iMpOrTeDdEvIcEiDeNtItIeS/iMpOrTeDdEvIcEiDeNtItYiD",
			Expected: &DeviceManagementImportedDeviceIdentityId{
				ImportedDeviceIdentityId: "iMpOrTeDdEvIcEiDeNtItYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/iMpOrTeDdEvIcEiDeNtItIeS/iMpOrTeDdEvIcEiDeNtItYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementImportedDeviceIdentityIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ImportedDeviceIdentityId != v.Expected.ImportedDeviceIdentityId {
			t.Fatalf("Expected %q but got %q for ImportedDeviceIdentityId", v.Expected.ImportedDeviceIdentityId, actual.ImportedDeviceIdentityId)
		}

	}
}

func TestSegmentsForDeviceManagementImportedDeviceIdentityId(t *testing.T) {
	segments := DeviceManagementImportedDeviceIdentityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementImportedDeviceIdentityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
