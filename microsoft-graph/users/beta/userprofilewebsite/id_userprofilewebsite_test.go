package userprofilewebsite

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileWebsiteId{}

func TestNewUserProfileWebsiteID(t *testing.T) {
	id := NewUserProfileWebsiteID("userIdValue", "personWebsiteIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.PersonWebsiteId != "personWebsiteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonWebsiteId'", id.PersonWebsiteId, "personWebsiteIdValue")
	}
}

func TestFormatUserProfileWebsiteID(t *testing.T) {
	actual := NewUserProfileWebsiteID("userIdValue", "personWebsiteIdValue").ID()
	expected := "/users/userIdValue/profile/websites/personWebsiteIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserProfileWebsiteID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileWebsiteId
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
			Input: "/users/userIdValue/profile/websites",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/websites/personWebsiteIdValue",
			Expected: &UserProfileWebsiteId{
				UserId:          "userIdValue",
				PersonWebsiteId: "personWebsiteIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/websites/personWebsiteIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileWebsiteID(v.Input)
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

func TestParseUserProfileWebsiteIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileWebsiteId
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
			Input: "/users/userIdValue/profile/websites",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/wEbSiTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/websites/personWebsiteIdValue",
			Expected: &UserProfileWebsiteId{
				UserId:          "userIdValue",
				PersonWebsiteId: "personWebsiteIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/websites/personWebsiteIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/wEbSiTeS/pErSoNwEbSiTeIdVaLuE",
			Expected: &UserProfileWebsiteId{
				UserId:          "uSeRiDvAlUe",
				PersonWebsiteId: "pErSoNwEbSiTeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/wEbSiTeS/pErSoNwEbSiTeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileWebsiteIDInsensitively(v.Input)
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

func TestSegmentsForUserProfileWebsiteId(t *testing.T) {
	segments := UserProfileWebsiteId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserProfileWebsiteId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
