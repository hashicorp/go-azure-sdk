package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePersonId{}

func TestNewMePersonID(t *testing.T) {
	id := NewMePersonID("personId")

	if id.PersonId != "personId" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonId'", id.PersonId, "personId")
	}
}

func TestFormatMePersonID(t *testing.T) {
	actual := NewMePersonID("personId").ID()
	expected := "/me/people/personId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMePersonID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePersonId
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
			Input: "/me/people",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/people/personId",
			Expected: &MePersonId{
				PersonId: "personId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/people/personId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePersonID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PersonId != v.Expected.PersonId {
			t.Fatalf("Expected %q but got %q for PersonId", v.Expected.PersonId, actual.PersonId)
		}

	}
}

func TestParseMePersonIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePersonId
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
			Input: "/me/people",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEoPlE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/people/personId",
			Expected: &MePersonId{
				PersonId: "personId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/people/personId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEoPlE/pErSoNiD",
			Expected: &MePersonId{
				PersonId: "pErSoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pEoPlE/pErSoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePersonIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PersonId != v.Expected.PersonId {
			t.Fatalf("Expected %q but got %q for PersonId", v.Expected.PersonId, actual.PersonId)
		}

	}
}

func TestSegmentsForMePersonId(t *testing.T) {
	segments := MePersonId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MePersonId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
