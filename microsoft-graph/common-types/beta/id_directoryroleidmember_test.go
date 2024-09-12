package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryRoleIdMemberId{}

func TestNewDirectoryRoleIdMemberID(t *testing.T) {
	id := NewDirectoryRoleIdMemberID("directoryRoleIdValue", "directoryObjectIdValue")

	if id.DirectoryRoleId != "directoryRoleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryRoleId'", id.DirectoryRoleId, "directoryRoleIdValue")
	}

	if id.DirectoryObjectId != "directoryObjectIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectIdValue")
	}
}

func TestFormatDirectoryRoleIdMemberID(t *testing.T) {
	actual := NewDirectoryRoleIdMemberID("directoryRoleIdValue", "directoryObjectIdValue").ID()
	expected := "/directoryRoles/directoryRoleIdValue/members/directoryObjectIdValue"
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
			Input: "/directoryRoles/directoryRoleIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directoryRoles/directoryRoleIdValue/members",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directoryRoles/directoryRoleIdValue/members/directoryObjectIdValue",
			Expected: &DirectoryRoleIdMemberId{
				DirectoryRoleId:   "directoryRoleIdValue",
				DirectoryObjectId: "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directoryRoles/directoryRoleIdValue/members/directoryObjectIdValue/extra",
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
			Input: "/directoryRoles/directoryRoleIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeS/dIrEcToRyRoLeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directoryRoles/directoryRoleIdValue/members",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeS/dIrEcToRyRoLeIdVaLuE/mEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directoryRoles/directoryRoleIdValue/members/directoryObjectIdValue",
			Expected: &DirectoryRoleIdMemberId{
				DirectoryRoleId:   "directoryRoleIdValue",
				DirectoryObjectId: "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directoryRoles/directoryRoleIdValue/members/directoryObjectIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeS/dIrEcToRyRoLeIdVaLuE/mEmBeRs/dIrEcToRyObJeCtIdVaLuE",
			Expected: &DirectoryRoleIdMemberId{
				DirectoryRoleId:   "dIrEcToRyRoLeIdVaLuE",
				DirectoryObjectId: "dIrEcToRyObJeCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeS/dIrEcToRyRoLeIdVaLuE/mEmBeRs/dIrEcToRyObJeCtIdVaLuE/extra",
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
