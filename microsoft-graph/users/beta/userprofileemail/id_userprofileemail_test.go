package userprofileemail

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileEmailId{}

func TestNewUserProfileEmailID(t *testing.T) {
	id := NewUserProfileEmailID("userIdValue", "itemEmailIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ItemEmailId != "itemEmailIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemEmailId'", id.ItemEmailId, "itemEmailIdValue")
	}
}

func TestFormatUserProfileEmailID(t *testing.T) {
	actual := NewUserProfileEmailID("userIdValue", "itemEmailIdValue").ID()
	expected := "/users/userIdValue/profile/emails/itemEmailIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserProfileEmailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileEmailId
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
			Input: "/users/userIdValue/profile/emails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/emails/itemEmailIdValue",
			Expected: &UserProfileEmailId{
				UserId:      "userIdValue",
				ItemEmailId: "itemEmailIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/emails/itemEmailIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileEmailID(v.Input)
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

		if actual.ItemEmailId != v.Expected.ItemEmailId {
			t.Fatalf("Expected %q but got %q for ItemEmailId", v.Expected.ItemEmailId, actual.ItemEmailId)
		}

	}
}

func TestParseUserProfileEmailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileEmailId
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
			Input: "/users/userIdValue/profile/emails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/eMaIlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/emails/itemEmailIdValue",
			Expected: &UserProfileEmailId{
				UserId:      "userIdValue",
				ItemEmailId: "itemEmailIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/emails/itemEmailIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/eMaIlS/iTeMeMaIlIdVaLuE",
			Expected: &UserProfileEmailId{
				UserId:      "uSeRiDvAlUe",
				ItemEmailId: "iTeMeMaIlIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/eMaIlS/iTeMeMaIlIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileEmailIDInsensitively(v.Input)
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

		if actual.ItemEmailId != v.Expected.ItemEmailId {
			t.Fatalf("Expected %q but got %q for ItemEmailId", v.Expected.ItemEmailId, actual.ItemEmailId)
		}

	}
}

func TestSegmentsForUserProfileEmailId(t *testing.T) {
	segments := UserProfileEmailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserProfileEmailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
