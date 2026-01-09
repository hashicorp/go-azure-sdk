package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AdministrativeUnitIdMemberId{}

func TestNewAdministrativeUnitIdMemberID(t *testing.T) {
	id := NewAdministrativeUnitIdMemberID("administrativeUnitId", "directoryObjectId")

	if id.AdministrativeUnitId != "administrativeUnitId" {
		t.Fatalf("Expected %q but got %q for Segment 'AdministrativeUnitId'", id.AdministrativeUnitId, "administrativeUnitId")
	}

	if id.DirectoryObjectId != "directoryObjectId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectId")
	}
}

func TestFormatAdministrativeUnitIdMemberID(t *testing.T) {
	actual := NewAdministrativeUnitIdMemberID("administrativeUnitId", "directoryObjectId").ID()
	expected := "/administrativeUnits/administrativeUnitId/members/directoryObjectId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAdministrativeUnitIdMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AdministrativeUnitIdMemberId
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
			Input: "/administrativeUnits/administrativeUnitId/members",
			Error: true,
		},
		{
			// Valid URI
			Input: "/administrativeUnits/administrativeUnitId/members/directoryObjectId",
			Expected: &AdministrativeUnitIdMemberId{
				AdministrativeUnitId: "administrativeUnitId",
				DirectoryObjectId:    "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/administrativeUnits/administrativeUnitId/members/directoryObjectId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAdministrativeUnitIdMemberID(v.Input)
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

func TestParseAdministrativeUnitIdMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AdministrativeUnitIdMemberId
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
			Input: "/administrativeUnits/administrativeUnitId/members",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiD/mEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/administrativeUnits/administrativeUnitId/members/directoryObjectId",
			Expected: &AdministrativeUnitIdMemberId{
				AdministrativeUnitId: "administrativeUnitId",
				DirectoryObjectId:    "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/administrativeUnits/administrativeUnitId/members/directoryObjectId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiD/mEmBeRs/dIrEcToRyObJeCtId",
			Expected: &AdministrativeUnitIdMemberId{
				AdministrativeUnitId: "aDmInIsTrAtIvEuNiTiD",
				DirectoryObjectId:    "dIrEcToRyObJeCtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiD/mEmBeRs/dIrEcToRyObJeCtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAdministrativeUnitIdMemberIDInsensitively(v.Input)
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

func TestSegmentsForAdministrativeUnitIdMemberId(t *testing.T) {
	segments := AdministrativeUnitIdMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AdministrativeUnitIdMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
