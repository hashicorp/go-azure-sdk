package directoryadministrativeunitextension

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DirectoryAdministrativeUnitExtensionId{}

func TestNewDirectoryAdministrativeUnitExtensionID(t *testing.T) {
	id := NewDirectoryAdministrativeUnitExtensionID("administrativeUnitIdValue", "extensionIdValue")

	if id.AdministrativeUnitId != "administrativeUnitIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AdministrativeUnitId'", id.AdministrativeUnitId, "administrativeUnitIdValue")
	}

	if id.ExtensionId != "extensionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionIdValue")
	}
}

func TestFormatDirectoryAdministrativeUnitExtensionID(t *testing.T) {
	actual := NewDirectoryAdministrativeUnitExtensionID("administrativeUnitIdValue", "extensionIdValue").ID()
	expected := "/directory/administrativeUnits/administrativeUnitIdValue/extensions/extensionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryAdministrativeUnitExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryAdministrativeUnitExtensionId
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
			Input: "/directory/administrativeUnits",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/administrativeUnits/administrativeUnitIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/administrativeUnits/administrativeUnitIdValue/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/administrativeUnits/administrativeUnitIdValue/extensions/extensionIdValue",
			Expected: &DirectoryAdministrativeUnitExtensionId{
				AdministrativeUnitId: "administrativeUnitIdValue",
				ExtensionId:          "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/administrativeUnits/administrativeUnitIdValue/extensions/extensionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryAdministrativeUnitExtensionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AdministrativeUnitId != v.Expected.AdministrativeUnitId {
			t.Fatalf("Expected %q but got %q for AdministrativeUnitId", v.Expected.AdministrativeUnitId, actual.AdministrativeUnitId)
		}

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestParseDirectoryAdministrativeUnitExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryAdministrativeUnitExtensionId
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
			Input: "/directory/administrativeUnits",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aDmInIsTrAtIvEuNiTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/administrativeUnits/administrativeUnitIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/administrativeUnits/administrativeUnitIdValue/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiDvAlUe/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/administrativeUnits/administrativeUnitIdValue/extensions/extensionIdValue",
			Expected: &DirectoryAdministrativeUnitExtensionId{
				AdministrativeUnitId: "administrativeUnitIdValue",
				ExtensionId:          "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/administrativeUnits/administrativeUnitIdValue/extensions/extensionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiDvAlUe/eXtEnSiOnS/eXtEnSiOnIdVaLuE",
			Expected: &DirectoryAdministrativeUnitExtensionId{
				AdministrativeUnitId: "aDmInIsTrAtIvEuNiTiDvAlUe",
				ExtensionId:          "eXtEnSiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiDvAlUe/eXtEnSiOnS/eXtEnSiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryAdministrativeUnitExtensionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AdministrativeUnitId != v.Expected.AdministrativeUnitId {
			t.Fatalf("Expected %q but got %q for AdministrativeUnitId", v.Expected.AdministrativeUnitId, actual.AdministrativeUnitId)
		}

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestSegmentsForDirectoryAdministrativeUnitExtensionId(t *testing.T) {
	segments := DirectoryAdministrativeUnitExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryAdministrativeUnitExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
