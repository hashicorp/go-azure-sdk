package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AdministrativeUnitIdExtensionId{}

func TestNewAdministrativeUnitIdExtensionID(t *testing.T) {
	id := NewAdministrativeUnitIdExtensionID("administrativeUnitId", "extensionId")

	if id.AdministrativeUnitId != "administrativeUnitId" {
		t.Fatalf("Expected %q but got %q for Segment 'AdministrativeUnitId'", id.AdministrativeUnitId, "administrativeUnitId")
	}

	if id.ExtensionId != "extensionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionId")
	}
}

func TestFormatAdministrativeUnitIdExtensionID(t *testing.T) {
	actual := NewAdministrativeUnitIdExtensionID("administrativeUnitId", "extensionId").ID()
	expected := "/administrativeUnits/administrativeUnitId/extensions/extensionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAdministrativeUnitIdExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AdministrativeUnitIdExtensionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/administrativeUnits",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/administrativeUnits/administrativeUnitId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/administrativeUnits/administrativeUnitId/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/administrativeUnits/administrativeUnitId/extensions/extensionId",
			Expected: &AdministrativeUnitIdExtensionId{
				AdministrativeUnitId: "administrativeUnitId",
				ExtensionId:          "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/administrativeUnits/administrativeUnitId/extensions/extensionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAdministrativeUnitIdExtensionID(v.Input)
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

func TestParseAdministrativeUnitIdExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AdministrativeUnitIdExtensionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/administrativeUnits",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aDmInIsTrAtIvEuNiTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/administrativeUnits/administrativeUnitId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/administrativeUnits/administrativeUnitId/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiD/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/administrativeUnits/administrativeUnitId/extensions/extensionId",
			Expected: &AdministrativeUnitIdExtensionId{
				AdministrativeUnitId: "administrativeUnitId",
				ExtensionId:          "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/administrativeUnits/administrativeUnitId/extensions/extensionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiD/eXtEnSiOnS/eXtEnSiOnId",
			Expected: &AdministrativeUnitIdExtensionId{
				AdministrativeUnitId: "aDmInIsTrAtIvEuNiTiD",
				ExtensionId:          "eXtEnSiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiD/eXtEnSiOnS/eXtEnSiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAdministrativeUnitIdExtensionIDInsensitively(v.Input)
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

func TestSegmentsForAdministrativeUnitIdExtensionId(t *testing.T) {
	segments := AdministrativeUnitIdExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AdministrativeUnitIdExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
