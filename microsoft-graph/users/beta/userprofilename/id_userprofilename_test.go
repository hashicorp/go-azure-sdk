package userprofilename

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileNameId{}

func TestNewUserProfileNameID(t *testing.T) {
	id := NewUserProfileNameID("userIdValue", "personNameIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.PersonNameId != "personNameIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonNameId'", id.PersonNameId, "personNameIdValue")
	}
}

func TestFormatUserProfileNameID(t *testing.T) {
	actual := NewUserProfileNameID("userIdValue", "personNameIdValue").ID()
	expected := "/users/userIdValue/profile/names/personNameIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserProfileNameID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileNameId
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
			Input: "/users/userIdValue/profile/names",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/names/personNameIdValue",
			Expected: &UserProfileNameId{
				UserId:       "userIdValue",
				PersonNameId: "personNameIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/names/personNameIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileNameID(v.Input)
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

		if actual.PersonNameId != v.Expected.PersonNameId {
			t.Fatalf("Expected %q but got %q for PersonNameId", v.Expected.PersonNameId, actual.PersonNameId)
		}

	}
}

func TestParseUserProfileNameIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileNameId
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
			Input: "/users/userIdValue/profile/names",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/nAmEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/names/personNameIdValue",
			Expected: &UserProfileNameId{
				UserId:       "userIdValue",
				PersonNameId: "personNameIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/names/personNameIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/nAmEs/pErSoNnAmEiDvAlUe",
			Expected: &UserProfileNameId{
				UserId:       "uSeRiDvAlUe",
				PersonNameId: "pErSoNnAmEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/nAmEs/pErSoNnAmEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileNameIDInsensitively(v.Input)
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

		if actual.PersonNameId != v.Expected.PersonNameId {
			t.Fatalf("Expected %q but got %q for PersonNameId", v.Expected.PersonNameId, actual.PersonNameId)
		}

	}
}

func TestSegmentsForUserProfileNameId(t *testing.T) {
	segments := UserProfileNameId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserProfileNameId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
