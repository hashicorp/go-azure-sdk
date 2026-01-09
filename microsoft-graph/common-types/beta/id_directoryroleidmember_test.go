package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryRoleIdMemberId{}

func TestNewDirectoryRoleIdMemberID(t *testing.T) {
	id := NewDirectoryRoleIdMemberID("directoryRoleId", "directoryObjectId")

	if id.DirectoryRoleId != "directoryRoleId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryRoleId'", id.DirectoryRoleId, "directoryRoleId")
	}

	if id.DirectoryObjectId != "directoryObjectId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectId")
	}
}

func TestFormatDirectoryRoleIdMemberID(t *testing.T) {
	actual := NewDirectoryRoleIdMemberID("directoryRoleId", "directoryObjectId").ID()
	expected := "/directoryRoles/directoryRoleId/members/directoryObjectId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryRoleIdMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryRoleIdMemberId
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
			Input: "/directoryRoles/directoryRoleId/members",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directoryRoles/directoryRoleId/members/directoryObjectId",
			Expected: &DirectoryRoleIdMemberId{
				DirectoryRoleId:   "directoryRoleId",
				DirectoryObjectId: "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directoryRoles/directoryRoleId/members/directoryObjectId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryRoleIdMemberID(v.Input)
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

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParseDirectoryRoleIdMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryRoleIdMemberId
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
			Input: "/directoryRoles/directoryRoleId/members",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeS/dIrEcToRyRoLeId/mEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directoryRoles/directoryRoleId/members/directoryObjectId",
			Expected: &DirectoryRoleIdMemberId{
				DirectoryRoleId:   "directoryRoleId",
				DirectoryObjectId: "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directoryRoles/directoryRoleId/members/directoryObjectId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeS/dIrEcToRyRoLeId/mEmBeRs/dIrEcToRyObJeCtId",
			Expected: &DirectoryRoleIdMemberId{
				DirectoryRoleId:   "dIrEcToRyRoLeId",
				DirectoryObjectId: "dIrEcToRyObJeCtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeS/dIrEcToRyRoLeId/mEmBeRs/dIrEcToRyObJeCtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryRoleIdMemberIDInsensitively(v.Input)
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

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForDirectoryRoleIdMemberId(t *testing.T) {
	segments := DirectoryRoleIdMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryRoleIdMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
