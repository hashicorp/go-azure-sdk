package userprofilepublication

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfilePublicationId{}

func TestNewUserProfilePublicationID(t *testing.T) {
	id := NewUserProfilePublicationID("userIdValue", "itemPublicationIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ItemPublicationId != "itemPublicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemPublicationId'", id.ItemPublicationId, "itemPublicationIdValue")
	}
}

func TestFormatUserProfilePublicationID(t *testing.T) {
	actual := NewUserProfilePublicationID("userIdValue", "itemPublicationIdValue").ID()
	expected := "/users/userIdValue/profile/publications/itemPublicationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserProfilePublicationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfilePublicationId
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
			Input: "/users/userIdValue/profile/publications",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/publications/itemPublicationIdValue",
			Expected: &UserProfilePublicationId{
				UserId:            "userIdValue",
				ItemPublicationId: "itemPublicationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/publications/itemPublicationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfilePublicationID(v.Input)
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

		if actual.ItemPublicationId != v.Expected.ItemPublicationId {
			t.Fatalf("Expected %q but got %q for ItemPublicationId", v.Expected.ItemPublicationId, actual.ItemPublicationId)
		}

	}
}

func TestParseUserProfilePublicationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfilePublicationId
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
			Input: "/users/userIdValue/profile/publications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/pUbLiCaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/publications/itemPublicationIdValue",
			Expected: &UserProfilePublicationId{
				UserId:            "userIdValue",
				ItemPublicationId: "itemPublicationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/publications/itemPublicationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/pUbLiCaTiOnS/iTeMpUbLiCaTiOnIdVaLuE",
			Expected: &UserProfilePublicationId{
				UserId:            "uSeRiDvAlUe",
				ItemPublicationId: "iTeMpUbLiCaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/pUbLiCaTiOnS/iTeMpUbLiCaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfilePublicationIDInsensitively(v.Input)
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

		if actual.ItemPublicationId != v.Expected.ItemPublicationId {
			t.Fatalf("Expected %q but got %q for ItemPublicationId", v.Expected.ItemPublicationId, actual.ItemPublicationId)
		}

	}
}

func TestSegmentsForUserProfilePublicationId(t *testing.T) {
	segments := UserProfilePublicationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserProfilePublicationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
