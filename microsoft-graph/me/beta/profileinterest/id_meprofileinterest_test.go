package profileinterest

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileInterestId{}

func TestNewMeProfileInterestID(t *testing.T) {
	id := NewMeProfileInterestID("personInterestIdValue")

	if id.PersonInterestId != "personInterestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonInterestId'", id.PersonInterestId, "personInterestIdValue")
	}
}

func TestFormatMeProfileInterestID(t *testing.T) {
	actual := NewMeProfileInterestID("personInterestIdValue").ID()
	expected := "/me/profile/interests/personInterestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfileInterestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileInterestId
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
			Input: "/me/profile/interests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/interests/personInterestIdValue",
			Expected: &MeProfileInterestId{
				PersonInterestId: "personInterestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/interests/personInterestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileInterestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PersonInterestId != v.Expected.PersonInterestId {
			t.Fatalf("Expected %q but got %q for PersonInterestId", v.Expected.PersonInterestId, actual.PersonInterestId)
		}

	}
}

func TestParseMeProfileInterestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileInterestId
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
			Input: "/me/profile/interests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/iNtErEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/interests/personInterestIdValue",
			Expected: &MeProfileInterestId{
				PersonInterestId: "personInterestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/interests/personInterestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/iNtErEsTs/pErSoNiNtErEsTiDvAlUe",
			Expected: &MeProfileInterestId{
				PersonInterestId: "pErSoNiNtErEsTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/iNtErEsTs/pErSoNiNtErEsTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileInterestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PersonInterestId != v.Expected.PersonInterestId {
			t.Fatalf("Expected %q but got %q for PersonInterestId", v.Expected.PersonInterestId, actual.PersonInterestId)
		}

	}
}

func TestSegmentsForMeProfileInterestId(t *testing.T) {
	segments := MeProfileInterestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfileInterestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
