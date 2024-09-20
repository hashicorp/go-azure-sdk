package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileWebsiteId{}

func TestNewUserIdProfileWebsiteID(t *testing.T) {
	id := NewUserIdProfileWebsiteID("userId", "personWebsiteId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.PersonWebsiteId != "personWebsiteId" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonWebsiteId'", id.PersonWebsiteId, "personWebsiteId")
	}
}

func TestFormatUserIdProfileWebsiteID(t *testing.T) {
	actual := NewUserIdProfileWebsiteID("userId", "personWebsiteId").ID()
	expected := "/users/userId/profile/websites/personWebsiteId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfileWebsiteID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileWebsiteId
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
			Input: "/users/userId/profile/websites",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/websites/personWebsiteId",
			Expected: &UserIdProfileWebsiteId{
				UserId:          "userId",
				PersonWebsiteId: "personWebsiteId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/websites/personWebsiteId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileWebsiteID(v.Input)
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

		if actual.PersonWebsiteId != v.Expected.PersonWebsiteId {
			t.Fatalf("Expected %q but got %q for PersonWebsiteId", v.Expected.PersonWebsiteId, actual.PersonWebsiteId)
		}

	}
}

func TestParseUserIdProfileWebsiteIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileWebsiteId
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
			Input: "/users/userId/profile/websites",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/wEbSiTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/websites/personWebsiteId",
			Expected: &UserIdProfileWebsiteId{
				UserId:          "userId",
				PersonWebsiteId: "personWebsiteId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/websites/personWebsiteId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/wEbSiTeS/pErSoNwEbSiTeId",
			Expected: &UserIdProfileWebsiteId{
				UserId:          "uSeRiD",
				PersonWebsiteId: "pErSoNwEbSiTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/wEbSiTeS/pErSoNwEbSiTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileWebsiteIDInsensitively(v.Input)
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

		if actual.PersonWebsiteId != v.Expected.PersonWebsiteId {
			t.Fatalf("Expected %q but got %q for PersonWebsiteId", v.Expected.PersonWebsiteId, actual.PersonWebsiteId)
		}

	}
}

func TestSegmentsForUserIdProfileWebsiteId(t *testing.T) {
	segments := UserIdProfileWebsiteId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfileWebsiteId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
