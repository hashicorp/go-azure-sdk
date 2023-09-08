package administrativeunitmember

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AdministrativeUnitMemberId{}

func TestNewAdministrativeUnitMemberID(t *testing.T) {
	id := NewAdministrativeUnitMemberID("administrativeUnitIdValue", "directoryObjectIdValue")

	if id.AdministrativeUnitId != "administrativeUnitIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AdministrativeUnitId'", id.AdministrativeUnitId, "administrativeUnitIdValue")
	}

	if id.DirectoryObjectId != "directoryObjectIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectIdValue")
	}
}

func TestFormatAdministrativeUnitMemberID(t *testing.T) {
	actual := NewAdministrativeUnitMemberID("administrativeUnitIdValue", "directoryObjectIdValue").ID()
	expected := "/administrativeUnits/administrativeUnitIdValue/members/directoryObjectIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAdministrativeUnitMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AdministrativeUnitMemberId
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
			Input: "/administrativeUnits/administrativeUnitIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/administrativeUnits/administrativeUnitIdValue/members",
			Error: true,
		},
		{
			// Valid URI
			Input: "/administrativeUnits/administrativeUnitIdValue/members/directoryObjectIdValue",
			Expected: &AdministrativeUnitMemberId{
				AdministrativeUnitId: "administrativeUnitIdValue",
				DirectoryObjectId:    "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/administrativeUnits/administrativeUnitIdValue/members/directoryObjectIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAdministrativeUnitMemberID(v.Input)
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

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParseAdministrativeUnitMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AdministrativeUnitMemberId
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
			Input: "/administrativeUnits/administrativeUnitIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/administrativeUnits/administrativeUnitIdValue/members",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiDvAlUe/mEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/administrativeUnits/administrativeUnitIdValue/members/directoryObjectIdValue",
			Expected: &AdministrativeUnitMemberId{
				AdministrativeUnitId: "administrativeUnitIdValue",
				DirectoryObjectId:    "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/administrativeUnits/administrativeUnitIdValue/members/directoryObjectIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiDvAlUe/mEmBeRs/dIrEcToRyObJeCtIdVaLuE",
			Expected: &AdministrativeUnitMemberId{
				AdministrativeUnitId: "aDmInIsTrAtIvEuNiTiDvAlUe",
				DirectoryObjectId:    "dIrEcToRyObJeCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiDvAlUe/mEmBeRs/dIrEcToRyObJeCtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAdministrativeUnitMemberIDInsensitively(v.Input)
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

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForAdministrativeUnitMemberId(t *testing.T) {
	segments := AdministrativeUnitMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AdministrativeUnitMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
