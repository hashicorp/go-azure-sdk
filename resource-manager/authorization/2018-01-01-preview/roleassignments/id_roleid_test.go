package roleassignments

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleIdId{}

func TestNewRoleIdID(t *testing.T) {
	id := NewRoleIdID("roleId")

	if id.RoleId != "roleId" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleId'", id.RoleId, "roleId")
	}
}

func TestFormatRoleIdID(t *testing.T) {
	actual := NewRoleIdID("roleId").ID()
	expected := "/roleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleIdID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleIdId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleId",
			Expected: &RoleIdId{
				RoleId: "roleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleIdID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RoleId != v.Expected.RoleId {
			t.Fatalf("Expected %q but got %q for RoleId", v.Expected.RoleId, actual.RoleId)
		}

	}
}

func TestParseRoleIdIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleIdId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleId",
			Expected: &RoleIdId{
				RoleId: "roleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEiD",
			Expected: &RoleIdId{
				RoleId: "rOlEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleIdIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RoleId != v.Expected.RoleId {
			t.Fatalf("Expected %q but got %q for RoleId", v.Expected.RoleId, actual.RoleId)
		}

	}
}

func TestSegmentsForRoleIdId(t *testing.T) {
	segments := RoleIdId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleIdId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
