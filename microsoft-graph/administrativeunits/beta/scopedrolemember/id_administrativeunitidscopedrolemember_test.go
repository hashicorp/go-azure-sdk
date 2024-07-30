package scopedrolemember

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AdministrativeUnitIdScopedRoleMemberId{}

func TestNewAdministrativeUnitIdScopedRoleMemberID(t *testing.T) {
	id := NewAdministrativeUnitIdScopedRoleMemberID("administrativeUnitIdValue", "scopedRoleMembershipIdValue")

	if id.AdministrativeUnitId != "administrativeUnitIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AdministrativeUnitId'", id.AdministrativeUnitId, "administrativeUnitIdValue")
	}

	if id.ScopedRoleMembershipId != "scopedRoleMembershipIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ScopedRoleMembershipId'", id.ScopedRoleMembershipId, "scopedRoleMembershipIdValue")
	}
}

func TestFormatAdministrativeUnitIdScopedRoleMemberID(t *testing.T) {
	actual := NewAdministrativeUnitIdScopedRoleMemberID("administrativeUnitIdValue", "scopedRoleMembershipIdValue").ID()
	expected := "/administrativeUnits/administrativeUnitIdValue/scopedRoleMembers/scopedRoleMembershipIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAdministrativeUnitIdScopedRoleMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AdministrativeUnitIdScopedRoleMemberId
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
			Input: "/administrativeUnits/administrativeUnitIdValue/scopedRoleMembers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/administrativeUnits/administrativeUnitIdValue/scopedRoleMembers/scopedRoleMembershipIdValue",
			Expected: &AdministrativeUnitIdScopedRoleMemberId{
				AdministrativeUnitId:   "administrativeUnitIdValue",
				ScopedRoleMembershipId: "scopedRoleMembershipIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/administrativeUnits/administrativeUnitIdValue/scopedRoleMembers/scopedRoleMembershipIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAdministrativeUnitIdScopedRoleMemberID(v.Input)
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

func TestParseAdministrativeUnitIdScopedRoleMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AdministrativeUnitIdScopedRoleMemberId
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
			Input: "/administrativeUnits/administrativeUnitIdValue/scopedRoleMembers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiDvAlUe/sCoPeDrOlEmEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/administrativeUnits/administrativeUnitIdValue/scopedRoleMembers/scopedRoleMembershipIdValue",
			Expected: &AdministrativeUnitIdScopedRoleMemberId{
				AdministrativeUnitId:   "administrativeUnitIdValue",
				ScopedRoleMembershipId: "scopedRoleMembershipIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/administrativeUnits/administrativeUnitIdValue/scopedRoleMembers/scopedRoleMembershipIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiDvAlUe/sCoPeDrOlEmEmBeRs/sCoPeDrOlEmEmBeRsHiPiDvAlUe",
			Expected: &AdministrativeUnitIdScopedRoleMemberId{
				AdministrativeUnitId:   "aDmInIsTrAtIvEuNiTiDvAlUe",
				ScopedRoleMembershipId: "sCoPeDrOlEmEmBeRsHiPiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aDmInIsTrAtIvEuNiTs/aDmInIsTrAtIvEuNiTiDvAlUe/sCoPeDrOlEmEmBeRs/sCoPeDrOlEmEmBeRsHiPiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAdministrativeUnitIdScopedRoleMemberIDInsensitively(v.Input)
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

func TestSegmentsForAdministrativeUnitIdScopedRoleMemberId(t *testing.T) {
	segments := AdministrativeUnitIdScopedRoleMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AdministrativeUnitIdScopedRoleMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
