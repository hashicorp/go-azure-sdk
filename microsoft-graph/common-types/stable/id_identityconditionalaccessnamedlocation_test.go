package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccessNamedLocationId{}

func TestNewIdentityConditionalAccessNamedLocationID(t *testing.T) {
	id := NewIdentityConditionalAccessNamedLocationID("namedLocationId")

	if id.NamedLocationId != "namedLocationId" {
		t.Fatalf("Expected %q but got %q for Segment 'NamedLocationId'", id.NamedLocationId, "namedLocationId")
	}
}

func TestFormatIdentityConditionalAccessNamedLocationID(t *testing.T) {
	actual := NewIdentityConditionalAccessNamedLocationID("namedLocationId").ID()
	expected := "/identity/conditionalAccess/namedLocations/namedLocationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityConditionalAccessNamedLocationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccessNamedLocationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess/namedLocations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/conditionalAccess/namedLocations/namedLocationId",
			Expected: &IdentityConditionalAccessNamedLocationId{
				NamedLocationId: "namedLocationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/namedLocations/namedLocationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccessNamedLocationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.NamedLocationId != v.Expected.NamedLocationId {
			t.Fatalf("Expected %q but got %q for NamedLocationId", v.Expected.NamedLocationId, actual.NamedLocationId)
		}

	}
}

func TestParseIdentityConditionalAccessNamedLocationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccessNamedLocationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess/namedLocations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/nAmEdLoCaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/conditionalAccess/namedLocations/namedLocationId",
			Expected: &IdentityConditionalAccessNamedLocationId{
				NamedLocationId: "namedLocationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/namedLocations/namedLocationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/nAmEdLoCaTiOnS/nAmEdLoCaTiOnId",
			Expected: &IdentityConditionalAccessNamedLocationId{
				NamedLocationId: "nAmEdLoCaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/nAmEdLoCaTiOnS/nAmEdLoCaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccessNamedLocationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.NamedLocationId != v.Expected.NamedLocationId {
			t.Fatalf("Expected %q but got %q for NamedLocationId", v.Expected.NamedLocationId, actual.NamedLocationId)
		}

	}
}

func TestSegmentsForIdentityConditionalAccessNamedLocationId(t *testing.T) {
	segments := IdentityConditionalAccessNamedLocationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityConditionalAccessNamedLocationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
