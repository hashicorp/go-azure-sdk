package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileAwardId{}

func TestNewUserIdProfileAwardID(t *testing.T) {
	id := NewUserIdProfileAwardID("userId", "personAwardId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.PersonAwardId != "personAwardId" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonAwardId'", id.PersonAwardId, "personAwardId")
	}
}

func TestFormatUserIdProfileAwardID(t *testing.T) {
	actual := NewUserIdProfileAwardID("userId", "personAwardId").ID()
	expected := "/users/userId/profile/awards/personAwardId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfileAwardID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileAwardId
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
			Input: "/users/userId/profile/awards",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/awards/personAwardId",
			Expected: &UserIdProfileAwardId{
				UserId:        "userId",
				PersonAwardId: "personAwardId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/awards/personAwardId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileAwardID(v.Input)
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

func TestParseUserIdProfileAwardIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileAwardId
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
			Input: "/users/userId/profile/awards",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/aWaRdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/awards/personAwardId",
			Expected: &UserIdProfileAwardId{
				UserId:        "userId",
				PersonAwardId: "personAwardId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/awards/personAwardId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/aWaRdS/pErSoNaWaRdId",
			Expected: &UserIdProfileAwardId{
				UserId:        "uSeRiD",
				PersonAwardId: "pErSoNaWaRdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/aWaRdS/pErSoNaWaRdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileAwardIDInsensitively(v.Input)
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

func TestSegmentsForUserIdProfileAwardId(t *testing.T) {
	segments := UserIdProfileAwardId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfileAwardId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
