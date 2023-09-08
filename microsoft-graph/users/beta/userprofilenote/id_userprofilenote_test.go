package userprofilenote

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileNoteId{}

func TestNewUserProfileNoteID(t *testing.T) {
	id := NewUserProfileNoteID("userIdValue", "personAnnotationIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.PersonAnnotationId != "personAnnotationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonAnnotationId'", id.PersonAnnotationId, "personAnnotationIdValue")
	}
}

func TestFormatUserProfileNoteID(t *testing.T) {
	actual := NewUserProfileNoteID("userIdValue", "personAnnotationIdValue").ID()
	expected := "/users/userIdValue/profile/notes/personAnnotationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserProfileNoteID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileNoteId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/profile",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/profile/notes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/notes/personAnnotationIdValue",
			Expected: &UserProfileNoteId{
				UserId:             "userIdValue",
				PersonAnnotationId: "personAnnotationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/notes/personAnnotationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileNoteID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.PersonAnnotationId != v.Expected.PersonAnnotationId {
			t.Fatalf("Expected %q but got %q for PersonAnnotationId", v.Expected.PersonAnnotationId, actual.PersonAnnotationId)
		}

	}
}

func TestParseUserProfileNoteIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileNoteId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/profile",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/profile/notes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/nOtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/notes/personAnnotationIdValue",
			Expected: &UserProfileNoteId{
				UserId:             "userIdValue",
				PersonAnnotationId: "personAnnotationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/notes/personAnnotationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/nOtEs/pErSoNaNnOtAtIoNiDvAlUe",
			Expected: &UserProfileNoteId{
				UserId:             "uSeRiDvAlUe",
				PersonAnnotationId: "pErSoNaNnOtAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/nOtEs/pErSoNaNnOtAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileNoteIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.PersonAnnotationId != v.Expected.PersonAnnotationId {
			t.Fatalf("Expected %q but got %q for PersonAnnotationId", v.Expected.PersonAnnotationId, actual.PersonAnnotationId)
		}

	}
}

func TestSegmentsForUserProfileNoteId(t *testing.T) {
	segments := UserProfileNoteId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserProfileNoteId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
