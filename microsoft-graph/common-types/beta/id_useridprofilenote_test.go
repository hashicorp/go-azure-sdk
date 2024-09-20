package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileNoteId{}

func TestNewUserIdProfileNoteID(t *testing.T) {
	id := NewUserIdProfileNoteID("userId", "personAnnotationId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.PersonAnnotationId != "personAnnotationId" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonAnnotationId'", id.PersonAnnotationId, "personAnnotationId")
	}
}

func TestFormatUserIdProfileNoteID(t *testing.T) {
	actual := NewUserIdProfileNoteID("userId", "personAnnotationId").ID()
	expected := "/users/userId/profile/notes/personAnnotationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfileNoteID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileNoteId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/profile",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/profile/notes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/notes/personAnnotationId",
			Expected: &UserIdProfileNoteId{
				UserId:             "userId",
				PersonAnnotationId: "personAnnotationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/notes/personAnnotationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileNoteID(v.Input)
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

func TestParseUserIdProfileNoteIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileNoteId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/profile",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/profile/notes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/nOtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/notes/personAnnotationId",
			Expected: &UserIdProfileNoteId{
				UserId:             "userId",
				PersonAnnotationId: "personAnnotationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/notes/personAnnotationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/nOtEs/pErSoNaNnOtAtIoNiD",
			Expected: &UserIdProfileNoteId{
				UserId:             "uSeRiD",
				PersonAnnotationId: "pErSoNaNnOtAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/nOtEs/pErSoNaNnOtAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileNoteIDInsensitively(v.Input)
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

func TestSegmentsForUserIdProfileNoteId(t *testing.T) {
	segments := UserIdProfileNoteId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfileNoteId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
