package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPersonId{}

func TestNewUserIdPersonID(t *testing.T) {
	id := NewUserIdPersonID("userId", "personId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.PersonId != "personId" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonId'", id.PersonId, "personId")
	}
}

func TestFormatUserIdPersonID(t *testing.T) {
	actual := NewUserIdPersonID("userId", "personId").ID()
	expected := "/users/userId/people/personId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdPersonID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPersonId
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
			Input: "/users/userId/people",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/people/personId",
			Expected: &UserIdPersonId{
				UserId:   "userId",
				PersonId: "personId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/people/personId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPersonID(v.Input)
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

func TestParseUserIdPersonIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPersonId
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
			Input: "/users/userId/people",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pEoPlE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/people/personId",
			Expected: &UserIdPersonId{
				UserId:   "userId",
				PersonId: "personId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/people/personId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pEoPlE/pErSoNiD",
			Expected: &UserIdPersonId{
				UserId:   "uSeRiD",
				PersonId: "pErSoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pEoPlE/pErSoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPersonIDInsensitively(v.Input)
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

func TestSegmentsForUserIdPersonId(t *testing.T) {
	segments := UserIdPersonId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdPersonId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
