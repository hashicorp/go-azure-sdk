package people

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPeopleId{}

func TestNewUserIdPeopleID(t *testing.T) {
	id := NewUserIdPeopleID("userIdValue", "personIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.PersonId != "personIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonId'", id.PersonId, "personIdValue")
	}
}

func TestFormatUserIdPeopleID(t *testing.T) {
	actual := NewUserIdPeopleID("userIdValue", "personIdValue").ID()
	expected := "/users/userIdValue/people/personIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdPeopleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPeopleId
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
			Input: "/users/userIdValue/people",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/people/personIdValue",
			Expected: &UserIdPeopleId{
				UserId:   "userIdValue",
				PersonId: "personIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/people/personIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPeopleID(v.Input)
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

		if actual.PersonId != v.Expected.PersonId {
			t.Fatalf("Expected %q but got %q for PersonId", v.Expected.PersonId, actual.PersonId)
		}

	}
}

func TestParseUserIdPeopleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPeopleId
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
			Input: "/users/userIdValue/people",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEoPlE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/people/personIdValue",
			Expected: &UserIdPeopleId{
				UserId:   "userIdValue",
				PersonId: "personIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/people/personIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEoPlE/pErSoNiDvAlUe",
			Expected: &UserIdPeopleId{
				UserId:   "uSeRiDvAlUe",
				PersonId: "pErSoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEoPlE/pErSoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPeopleIDInsensitively(v.Input)
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

		if actual.PersonId != v.Expected.PersonId {
			t.Fatalf("Expected %q but got %q for PersonId", v.Expected.PersonId, actual.PersonId)
		}

	}
}

func TestSegmentsForUserIdPeopleId(t *testing.T) {
	segments := UserIdPeopleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdPeopleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
