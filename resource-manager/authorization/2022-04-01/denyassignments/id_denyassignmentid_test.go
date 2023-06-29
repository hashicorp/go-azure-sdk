package denyassignments

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DenyAssignmentIdId{}

func TestNewDenyAssignmentIdID(t *testing.T) {
	id := NewDenyAssignmentIdID("denyAssignmentIdValue")

	if id.DenyAssignmentId != "denyAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DenyAssignmentId'", id.DenyAssignmentId, "denyAssignmentIdValue")
	}
}

func TestFormatDenyAssignmentIdID(t *testing.T) {
	actual := NewDenyAssignmentIdID("denyAssignmentIdValue").ID()
	expected := "/denyAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDenyAssignmentIdID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DenyAssignmentIdId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Valid URI
			Input: "/denyAssignmentIdValue",
			Expected: &DenyAssignmentIdId{
				DenyAssignmentId: "denyAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/denyAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDenyAssignmentIdID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DenyAssignmentId != v.Expected.DenyAssignmentId {
			t.Fatalf("Expected %q but got %q for DenyAssignmentId", v.Expected.DenyAssignmentId, actual.DenyAssignmentId)
		}

	}
}

func TestParseDenyAssignmentIdIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DenyAssignmentIdId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Valid URI
			Input: "/denyAssignmentIdValue",
			Expected: &DenyAssignmentIdId{
				DenyAssignmentId: "denyAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/denyAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEnYaSsIgNmEnTiDvAlUe",
			Expected: &DenyAssignmentIdId{
				DenyAssignmentId: "dEnYaSsIgNmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEnYaSsIgNmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDenyAssignmentIdIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DenyAssignmentId != v.Expected.DenyAssignmentId {
			t.Fatalf("Expected %q but got %q for DenyAssignmentId", v.Expected.DenyAssignmentId, actual.DenyAssignmentId)
		}

	}
}

func TestSegmentsForDenyAssignmentIdId(t *testing.T) {
	segments := DenyAssignmentIdId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DenyAssignmentIdId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
