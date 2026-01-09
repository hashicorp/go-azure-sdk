package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileNoteId{}

func TestNewMeProfileNoteID(t *testing.T) {
	id := NewMeProfileNoteID("personAnnotationId")

	if id.PersonAnnotationId != "personAnnotationId" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonAnnotationId'", id.PersonAnnotationId, "personAnnotationId")
	}
}

func TestFormatMeProfileNoteID(t *testing.T) {
	actual := NewMeProfileNoteID("personAnnotationId").ID()
	expected := "/me/profile/notes/personAnnotationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfileNoteID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileNoteId
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
			Input: "/me/profile/notes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/notes/personAnnotationId",
			Expected: &MeProfileNoteId{
				PersonAnnotationId: "personAnnotationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/notes/personAnnotationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileNoteID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PersonAnnotationId != v.Expected.PersonAnnotationId {
			t.Fatalf("Expected %q but got %q for PersonAnnotationId", v.Expected.PersonAnnotationId, actual.PersonAnnotationId)
		}

	}
}

func TestParseMeProfileNoteIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileNoteId
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
			Input: "/me/profile/notes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/nOtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/notes/personAnnotationId",
			Expected: &MeProfileNoteId{
				PersonAnnotationId: "personAnnotationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/notes/personAnnotationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/nOtEs/pErSoNaNnOtAtIoNiD",
			Expected: &MeProfileNoteId{
				PersonAnnotationId: "pErSoNaNnOtAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/nOtEs/pErSoNaNnOtAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileNoteIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PersonAnnotationId != v.Expected.PersonAnnotationId {
			t.Fatalf("Expected %q but got %q for PersonAnnotationId", v.Expected.PersonAnnotationId, actual.PersonAnnotationId)
		}

	}
}

func TestSegmentsForMeProfileNoteId(t *testing.T) {
	segments := MeProfileNoteId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfileNoteId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
