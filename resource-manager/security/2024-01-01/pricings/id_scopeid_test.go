package pricings

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopeIdId{}

func TestNewScopeIdID(t *testing.T) {
	id := NewScopeIdID("scopeId")

	if id.ScopeId != "scopeId" {
		t.Fatalf("Expected %q but got %q for Segment 'ScopeId'", id.ScopeId, "scopeId")
	}
}

func TestFormatScopeIdID(t *testing.T) {
	actual := NewScopeIdID("scopeId").ID()
	expected := "/scopeId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopeIdID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopeIdId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Valid URI
			Input: "/scopeId",
			Expected: &ScopeIdId{
				ScopeId: "scopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/scopeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopeIdID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ScopeId != v.Expected.ScopeId {
			t.Fatalf("Expected %q but got %q for ScopeId", v.Expected.ScopeId, actual.ScopeId)
		}

	}
}

func TestParseScopeIdIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopeIdId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Valid URI
			Input: "/scopeId",
			Expected: &ScopeIdId{
				ScopeId: "scopeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/scopeId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sCoPeId",
			Expected: &ScopeIdId{
				ScopeId: "sCoPeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sCoPeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopeIdIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ScopeId != v.Expected.ScopeId {
			t.Fatalf("Expected %q but got %q for ScopeId", v.Expected.ScopeId, actual.ScopeId)
		}

	}
}

func TestSegmentsForScopeIdId(t *testing.T) {
	segments := ScopeIdId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopeIdId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
