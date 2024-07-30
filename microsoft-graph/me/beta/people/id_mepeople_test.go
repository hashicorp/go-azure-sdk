package people

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePeopleId{}

func TestNewMePeopleID(t *testing.T) {
	id := NewMePeopleID("personIdValue")

	if id.PersonId != "personIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonId'", id.PersonId, "personIdValue")
	}
}

func TestFormatMePeopleID(t *testing.T) {
	actual := NewMePeopleID("personIdValue").ID()
	expected := "/me/people/personIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMePeopleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePeopleId
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
			Input: "/me/people/personIdValue",
			Expected: &MePeopleId{
				PersonId: "personIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/people/personIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePeopleID(v.Input)
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

func TestParseMePeopleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePeopleId
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
			Input: "/me/people/personIdValue",
			Expected: &MePeopleId{
				PersonId: "personIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/people/personIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEoPlE/pErSoNiDvAlUe",
			Expected: &MePeopleId{
				PersonId: "pErSoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pEoPlE/pErSoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePeopleIDInsensitively(v.Input)
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

func TestSegmentsForMePeopleId(t *testing.T) {
	segments := MePeopleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MePeopleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
