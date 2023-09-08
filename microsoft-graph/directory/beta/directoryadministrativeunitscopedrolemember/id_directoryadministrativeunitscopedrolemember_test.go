package directoryadministrativeunitscopedrolemember

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DirectoryAdministrativeUnitScopedRoleMemberId{}

func TestNewDirectoryAdministrativeUnitScopedRoleMemberID(t *testing.T) {
	id := NewDirectoryAdministrativeUnitScopedRoleMemberID("administrativeUnitIdValue", "scopedRoleMembershipIdValue")

	if id.AdministrativeUnitId != "administrativeUnitIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AdministrativeUnitId'", id.AdministrativeUnitId, "administrativeUnitIdValue")
	}

	if id.ScopedRoleMembershipId != "scopedRoleMembershipIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ScopedRoleMembershipId'", id.ScopedRoleMembershipId, "scopedRoleMembershipIdValue")
	}
}

func TestFormatDirectoryAdministrativeUnitScopedRoleMemberID(t *testing.T) {
	actual := NewDirectoryAdministrativeUnitScopedRoleMemberID("administrativeUnitIdValue", "scopedRoleMembershipIdValue").ID()
	expected := "/directory/administrativeUnits/administrativeUnitIdValue/scopedRoleMembers/scopedRoleMembershipIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryAdministrativeUnitScopedRoleMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryAdministrativeUnitScopedRoleMemberId
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
			Input: "/directory/administrativeUnits/administrativeUnitIdValue/scopedRoleMembers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/administrativeUnits/administrativeUnitIdValue/scopedRoleMembers/scopedRoleMembershipIdValue",
			Expected: &DirectoryAdministrativeUnitScopedRoleMemberId{
				AdministrativeUnitId:   "administrativeUnitIdValue",
				ScopedRoleMembershipId: "scopedRoleMembershipIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/administrativeUnits/administrativeUnitIdValue/scopedRoleMembers/scopedRoleMembershipIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryAdministrativeUnitScopedRoleMemberID(v.Input)
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

		if actual.ScopedRoleMembershipId != v.Expected.ScopedRoleMembershipId {
			t.Fatalf("Expected %q but got %q for ScopedRoleMembershipId", v.Expected.ScopedRoleMembershipId, actual.ScopedRoleMembershipId)
		}

	}
}

func TestParseDirectoryAdministrativeUnitScopedRoleMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryAdministrativeUnitScopedRoleMemberId
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
			Input: "/directory/administrativeUnits/administrativeUnitIdValue/scopedRoleMembers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiDvAlUe/sCoPeDrOlEmEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/administrativeUnits/administrativeUnitIdValue/scopedRoleMembers/scopedRoleMembershipIdValue",
			Expected: &DirectoryAdministrativeUnitScopedRoleMemberId{
				AdministrativeUnitId:   "administrativeUnitIdValue",
				ScopedRoleMembershipId: "scopedRoleMembershipIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/administrativeUnits/administrativeUnitIdValue/scopedRoleMembers/scopedRoleMembershipIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiDvAlUe/sCoPeDrOlEmEmBeRs/sCoPeDrOlEmEmBeRsHiPiDvAlUe",
			Expected: &DirectoryAdministrativeUnitScopedRoleMemberId{
				AdministrativeUnitId:   "aDmInIsTrAtIvEuNiTiDvAlUe",
				ScopedRoleMembershipId: "sCoPeDrOlEmEmBeRsHiPiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiDvAlUe/sCoPeDrOlEmEmBeRs/sCoPeDrOlEmEmBeRsHiPiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryAdministrativeUnitScopedRoleMemberIDInsensitively(v.Input)
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

		if actual.ScopedRoleMembershipId != v.Expected.ScopedRoleMembershipId {
			t.Fatalf("Expected %q but got %q for ScopedRoleMembershipId", v.Expected.ScopedRoleMembershipId, actual.ScopedRoleMembershipId)
		}

	}
}

func TestSegmentsForDirectoryAdministrativeUnitScopedRoleMemberId(t *testing.T) {
	segments := DirectoryAdministrativeUnitScopedRoleMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryAdministrativeUnitScopedRoleMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
