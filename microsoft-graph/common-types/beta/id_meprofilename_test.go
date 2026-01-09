package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileNameId{}

func TestNewMeProfileNameID(t *testing.T) {
	id := NewMeProfileNameID("personNameId")

	if id.PersonNameId != "personNameId" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonNameId'", id.PersonNameId, "personNameId")
	}
}

func TestFormatMeProfileNameID(t *testing.T) {
	actual := NewMeProfileNameID("personNameId").ID()
	expected := "/me/profile/names/personNameId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfileNameID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileNameId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile/names",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/names/personNameId",
			Expected: &MeProfileNameId{
				PersonNameId: "personNameId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/names/personNameId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileNameID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PersonNameId != v.Expected.PersonNameId {
			t.Fatalf("Expected %q but got %q for PersonNameId", v.Expected.PersonNameId, actual.PersonNameId)
		}

	}
}

func TestParseMeProfileNameIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileNameId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile/names",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/nAmEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/names/personNameId",
			Expected: &MeProfileNameId{
				PersonNameId: "personNameId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/names/personNameId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/nAmEs/pErSoNnAmEiD",
			Expected: &MeProfileNameId{
				PersonNameId: "pErSoNnAmEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/nAmEs/pErSoNnAmEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileNameIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PersonNameId != v.Expected.PersonNameId {
			t.Fatalf("Expected %q but got %q for PersonNameId", v.Expected.PersonNameId, actual.PersonNameId)
		}

	}
}

func TestSegmentsForMeProfileNameId(t *testing.T) {
	segments := MeProfileNameId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfileNameId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
