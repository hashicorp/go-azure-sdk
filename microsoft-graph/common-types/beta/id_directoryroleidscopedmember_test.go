package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryRoleIdScopedMemberId{}

func TestNewDirectoryRoleIdScopedMemberID(t *testing.T) {
	id := NewDirectoryRoleIdScopedMemberID("directoryRoleId", "scopedRoleMembershipId")

	if id.DirectoryRoleId != "directoryRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryRoleId'", id.DirectoryRoleId, "directoryRoleId")
	}

	if id.ScopedRoleMembershipId != "scopedRoleMembershipId" {
		t.Fatalf("Expected %q but got %q for Segment 'ScopedRoleMembershipId'", id.ScopedRoleMembershipId, "scopedRoleMembershipId")
	}
}

func TestFormatDirectoryRoleIdScopedMemberID(t *testing.T) {
	actual := NewDirectoryRoleIdScopedMemberID("directoryRoleId", "scopedRoleMembershipId").ID()
	expected := "/directoryRoles/directoryRoleId/scopedMembers/scopedRoleMembershipId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryRoleIdScopedMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryRoleIdScopedMemberId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directoryRoles",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directoryRoles/directoryRoleId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directoryRoles/directoryRoleId/scopedMembers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directoryRoles/directoryRoleId/scopedMembers/scopedRoleMembershipId",
			Expected: &DirectoryRoleIdScopedMemberId{
				DirectoryRoleId:        "directoryRoleId",
				ScopedRoleMembershipId: "scopedRoleMembershipId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directoryRoles/directoryRoleId/scopedMembers/scopedRoleMembershipId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryRoleIdScopedMemberID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DirectoryRoleId != v.Expected.DirectoryRoleId {
			t.Fatalf("Expected %q but got %q for DirectoryRoleId", v.Expected.DirectoryRoleId, actual.DirectoryRoleId)
		}

		if actual.ScopedRoleMembershipId != v.Expected.ScopedRoleMembershipId {
			t.Fatalf("Expected %q but got %q for ScopedRoleMembershipId", v.Expected.ScopedRoleMembershipId, actual.ScopedRoleMembershipId)
		}

	}
}

func TestParseDirectoryRoleIdScopedMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryRoleIdScopedMemberId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directoryRoles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directoryRoles/directoryRoleId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeS/dIrEcToRyRoLeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directoryRoles/directoryRoleId/scopedMembers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeS/dIrEcToRyRoLeId/sCoPeDmEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directoryRoles/directoryRoleId/scopedMembers/scopedRoleMembershipId",
			Expected: &DirectoryRoleIdScopedMemberId{
				DirectoryRoleId:        "directoryRoleId",
				ScopedRoleMembershipId: "scopedRoleMembershipId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directoryRoles/directoryRoleId/scopedMembers/scopedRoleMembershipId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeS/dIrEcToRyRoLeId/sCoPeDmEmBeRs/sCoPeDrOlEmEmBeRsHiPiD",
			Expected: &DirectoryRoleIdScopedMemberId{
				DirectoryRoleId:        "dIrEcToRyRoLeId",
				ScopedRoleMembershipId: "sCoPeDrOlEmEmBeRsHiPiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeS/dIrEcToRyRoLeId/sCoPeDmEmBeRs/sCoPeDrOlEmEmBeRsHiPiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryRoleIdScopedMemberIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DirectoryRoleId != v.Expected.DirectoryRoleId {
			t.Fatalf("Expected %q but got %q for DirectoryRoleId", v.Expected.DirectoryRoleId, actual.DirectoryRoleId)
		}

		if actual.ScopedRoleMembershipId != v.Expected.ScopedRoleMembershipId {
			t.Fatalf("Expected %q but got %q for ScopedRoleMembershipId", v.Expected.ScopedRoleMembershipId, actual.ScopedRoleMembershipId)
		}

	}
}

func TestSegmentsForDirectoryRoleIdScopedMemberId(t *testing.T) {
	segments := DirectoryRoleIdScopedMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryRoleIdScopedMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
