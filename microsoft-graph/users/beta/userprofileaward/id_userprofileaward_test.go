package userprofileaward

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileAwardId{}

func TestNewUserProfileAwardID(t *testing.T) {
	id := NewUserProfileAwardID("userIdValue", "personAwardIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.PersonAwardId != "personAwardIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonAwardId'", id.PersonAwardId, "personAwardIdValue")
	}
}

func TestFormatUserProfileAwardID(t *testing.T) {
	actual := NewUserProfileAwardID("userIdValue", "personAwardIdValue").ID()
	expected := "/users/userIdValue/profile/awards/personAwardIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserProfileAwardID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileAwardId
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
			Input: "/users/userIdValue/profile/awards",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/awards/personAwardIdValue",
			Expected: &UserProfileAwardId{
				UserId:        "userIdValue",
				PersonAwardId: "personAwardIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/awards/personAwardIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileAwardID(v.Input)
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

		if actual.PersonAwardId != v.Expected.PersonAwardId {
			t.Fatalf("Expected %q but got %q for PersonAwardId", v.Expected.PersonAwardId, actual.PersonAwardId)
		}

	}
}

func TestParseUserProfileAwardIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileAwardId
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
			Input: "/users/userIdValue/profile/awards",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/aWaRdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/awards/personAwardIdValue",
			Expected: &UserProfileAwardId{
				UserId:        "userIdValue",
				PersonAwardId: "personAwardIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/awards/personAwardIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/aWaRdS/pErSoNaWaRdIdVaLuE",
			Expected: &UserProfileAwardId{
				UserId:        "uSeRiDvAlUe",
				PersonAwardId: "pErSoNaWaRdIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/aWaRdS/pErSoNaWaRdIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileAwardIDInsensitively(v.Input)
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

		if actual.PersonAwardId != v.Expected.PersonAwardId {
			t.Fatalf("Expected %q but got %q for PersonAwardId", v.Expected.PersonAwardId, actual.PersonAwardId)
		}

	}
}

func TestSegmentsForUserProfileAwardId(t *testing.T) {
	segments := UserProfileAwardId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserProfileAwardId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
