package profileinterest

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileInterestId{}

func TestNewUserIdProfileInterestID(t *testing.T) {
	id := NewUserIdProfileInterestID("userIdValue", "personInterestIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.PersonInterestId != "personInterestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonInterestId'", id.PersonInterestId, "personInterestIdValue")
	}
}

func TestFormatUserIdProfileInterestID(t *testing.T) {
	actual := NewUserIdProfileInterestID("userIdValue", "personInterestIdValue").ID()
	expected := "/users/userIdValue/profile/interests/personInterestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfileInterestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileInterestId
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
			Input: "/users/userIdValue/profile/interests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/interests/personInterestIdValue",
			Expected: &UserIdProfileInterestId{
				UserId:           "userIdValue",
				PersonInterestId: "personInterestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/interests/personInterestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileInterestID(v.Input)
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

		if actual.PersonInterestId != v.Expected.PersonInterestId {
			t.Fatalf("Expected %q but got %q for PersonInterestId", v.Expected.PersonInterestId, actual.PersonInterestId)
		}

	}
}

func TestParseUserIdProfileInterestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileInterestId
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
			Input: "/users/userIdValue/profile/interests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/iNtErEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/interests/personInterestIdValue",
			Expected: &UserIdProfileInterestId{
				UserId:           "userIdValue",
				PersonInterestId: "personInterestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/interests/personInterestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/iNtErEsTs/pErSoNiNtErEsTiDvAlUe",
			Expected: &UserIdProfileInterestId{
				UserId:           "uSeRiDvAlUe",
				PersonInterestId: "pErSoNiNtErEsTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/iNtErEsTs/pErSoNiNtErEsTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileInterestIDInsensitively(v.Input)
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

		if actual.PersonInterestId != v.Expected.PersonInterestId {
			t.Fatalf("Expected %q but got %q for PersonInterestId", v.Expected.PersonInterestId, actual.PersonInterestId)
		}

	}
}

func TestSegmentsForUserIdProfileInterestId(t *testing.T) {
	segments := UserIdProfileInterestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfileInterestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
