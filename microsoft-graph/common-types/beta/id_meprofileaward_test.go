package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileAwardId{}

func TestNewMeProfileAwardID(t *testing.T) {
	id := NewMeProfileAwardID("personAwardIdValue")

	if id.PersonAwardId != "personAwardIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonAwardId'", id.PersonAwardId, "personAwardIdValue")
	}
}

func TestFormatMeProfileAwardID(t *testing.T) {
	actual := NewMeProfileAwardID("personAwardIdValue").ID()
	expected := "/me/profile/awards/personAwardIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfileAwardID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileAwardId
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
			Input: "/me/profile/awards",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/awards/personAwardIdValue",
			Expected: &MeProfileAwardId{
				PersonAwardId: "personAwardIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/awards/personAwardIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileAwardID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PersonAwardId != v.Expected.PersonAwardId {
			t.Fatalf("Expected %q but got %q for PersonAwardId", v.Expected.PersonAwardId, actual.PersonAwardId)
		}

	}
}

func TestParseMeProfileAwardIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileAwardId
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
			Input: "/me/profile/awards",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/aWaRdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/awards/personAwardIdValue",
			Expected: &MeProfileAwardId{
				PersonAwardId: "personAwardIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/awards/personAwardIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/aWaRdS/pErSoNaWaRdIdVaLuE",
			Expected: &MeProfileAwardId{
				PersonAwardId: "pErSoNaWaRdIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/aWaRdS/pErSoNaWaRdIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileAwardIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PersonAwardId != v.Expected.PersonAwardId {
			t.Fatalf("Expected %q but got %q for PersonAwardId", v.Expected.PersonAwardId, actual.PersonAwardId)
		}

	}
}

func TestSegmentsForMeProfileAwardId(t *testing.T) {
	segments := MeProfileAwardId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfileAwardId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
