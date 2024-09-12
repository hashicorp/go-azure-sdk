package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryRoleId{}

func TestNewDirectoryRoleID(t *testing.T) {
	id := NewDirectoryRoleID("directoryRoleIdValue")

	if id.DirectoryRoleId != "directoryRoleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryRoleId'", id.DirectoryRoleId, "directoryRoleIdValue")
	}
}

func TestFormatDirectoryRoleID(t *testing.T) {
	actual := NewDirectoryRoleID("directoryRoleIdValue").ID()
	expected := "/directoryRoles/directoryRoleIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryRoleId
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
			// Valid URI
			Input: "/directoryRoles/directoryRoleIdValue",
			Expected: &DirectoryRoleId{
				DirectoryRoleId: "directoryRoleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directoryRoles/directoryRoleIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryRoleID(v.Input)
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

	}
}

func TestParseDirectoryRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryRoleId
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
			// Valid URI
			Input: "/directoryRoles/directoryRoleIdValue",
			Expected: &DirectoryRoleId{
				DirectoryRoleId: "directoryRoleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directoryRoles/directoryRoleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeS/dIrEcToRyRoLeIdVaLuE",
			Expected: &DirectoryRoleId{
				DirectoryRoleId: "dIrEcToRyRoLeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeS/dIrEcToRyRoLeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryRoleIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDirectoryRoleId(t *testing.T) {
	segments := DirectoryRoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryRoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
