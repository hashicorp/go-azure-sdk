package identityconditionalaccesnamedlocation

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityConditionalAccesNamedLocationId{}

func TestNewIdentityConditionalAccesNamedLocationID(t *testing.T) {
	id := NewIdentityConditionalAccesNamedLocationID("namedLocationIdValue")

	if id.NamedLocationId != "namedLocationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'NamedLocationId'", id.NamedLocationId, "namedLocationIdValue")
	}
}

func TestFormatIdentityConditionalAccesNamedLocationID(t *testing.T) {
	actual := NewIdentityConditionalAccesNamedLocationID("namedLocationIdValue").ID()
	expected := "/identity/conditionalAccess/namedLocations/namedLocationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityConditionalAccesNamedLocationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccesNamedLocationId
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
			Input: "/identity/conditionalAccess/namedLocations/namedLocationIdValue",
			Expected: &IdentityConditionalAccesNamedLocationId{
				NamedLocationId: "namedLocationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/namedLocations/namedLocationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccesNamedLocationID(v.Input)
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

func TestParseIdentityConditionalAccesNamedLocationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccesNamedLocationId
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
			Input: "/identity/conditionalAccess/namedLocations/namedLocationIdValue",
			Expected: &IdentityConditionalAccesNamedLocationId{
				NamedLocationId: "namedLocationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/namedLocations/namedLocationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/nAmEdLoCaTiOnS/nAmEdLoCaTiOnIdVaLuE",
			Expected: &IdentityConditionalAccesNamedLocationId{
				NamedLocationId: "nAmEdLoCaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/nAmEdLoCaTiOnS/nAmEdLoCaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccesNamedLocationIDInsensitively(v.Input)
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

func TestSegmentsForIdentityConditionalAccesNamedLocationId(t *testing.T) {
	segments := IdentityConditionalAccesNamedLocationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityConditionalAccesNamedLocationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
