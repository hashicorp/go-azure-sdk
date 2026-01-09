package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryAdministrativeUnitIdScopedRoleMemberId{}

func TestNewDirectoryAdministrativeUnitIdScopedRoleMemberID(t *testing.T) {
	id := NewDirectoryAdministrativeUnitIdScopedRoleMemberID("administrativeUnitId", "scopedRoleMembershipId")

	if id.AdministrativeUnitId != "administrativeUnitId" {
		t.Fatalf("Expected %q but got %q for Segment 'AdministrativeUnitId'", id.AdministrativeUnitId, "administrativeUnitId")
	}

	if id.ScopedRoleMembershipId != "scopedRoleMembershipId" {
		t.Fatalf("Expected %q but got %q for Segment 'ScopedRoleMembershipId'", id.ScopedRoleMembershipId, "scopedRoleMembershipId")
	}
}

func TestFormatDirectoryAdministrativeUnitIdScopedRoleMemberID(t *testing.T) {
	actual := NewDirectoryAdministrativeUnitIdScopedRoleMemberID("administrativeUnitId", "scopedRoleMembershipId").ID()
	expected := "/directory/administrativeUnits/administrativeUnitId/scopedRoleMembers/scopedRoleMembershipId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryAdministrativeUnitIdScopedRoleMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryAdministrativeUnitIdScopedRoleMemberId
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
			Input: "/directory/administrativeUnits/administrativeUnitId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/administrativeUnits/administrativeUnitId/scopedRoleMembers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/administrativeUnits/administrativeUnitId/scopedRoleMembers/scopedRoleMembershipId",
			Expected: &DirectoryAdministrativeUnitIdScopedRoleMemberId{
				AdministrativeUnitId:   "administrativeUnitId",
				ScopedRoleMembershipId: "scopedRoleMembershipId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/administrativeUnits/administrativeUnitId/scopedRoleMembers/scopedRoleMembershipId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryAdministrativeUnitIdScopedRoleMemberID(v.Input)
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

func TestParseDirectoryAdministrativeUnitIdScopedRoleMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryAdministrativeUnitIdScopedRoleMemberId
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
			Input: "/directory/administrativeUnits/administrativeUnitId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/administrativeUnits/administrativeUnitId/scopedRoleMembers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiD/sCoPeDrOlEmEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/administrativeUnits/administrativeUnitId/scopedRoleMembers/scopedRoleMembershipId",
			Expected: &DirectoryAdministrativeUnitIdScopedRoleMemberId{
				AdministrativeUnitId:   "administrativeUnitId",
				ScopedRoleMembershipId: "scopedRoleMembershipId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/administrativeUnits/administrativeUnitId/scopedRoleMembers/scopedRoleMembershipId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiD/sCoPeDrOlEmEmBeRs/sCoPeDrOlEmEmBeRsHiPiD",
			Expected: &DirectoryAdministrativeUnitIdScopedRoleMemberId{
				AdministrativeUnitId:   "aDmInIsTrAtIvEuNiTiD",
				ScopedRoleMembershipId: "sCoPeDrOlEmEmBeRsHiPiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiD/sCoPeDrOlEmEmBeRs/sCoPeDrOlEmEmBeRsHiPiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryAdministrativeUnitIdScopedRoleMemberIDInsensitively(v.Input)
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

func TestSegmentsForDirectoryAdministrativeUnitIdScopedRoleMemberId(t *testing.T) {
	segments := DirectoryAdministrativeUnitIdScopedRoleMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryAdministrativeUnitIdScopedRoleMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
