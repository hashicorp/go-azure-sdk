package userprofileaddress

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileAddressId{}

func TestNewUserProfileAddressID(t *testing.T) {
	id := NewUserProfileAddressID("userIdValue", "itemAddressIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ItemAddressId != "itemAddressIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemAddressId'", id.ItemAddressId, "itemAddressIdValue")
	}
}

func TestFormatUserProfileAddressID(t *testing.T) {
	actual := NewUserProfileAddressID("userIdValue", "itemAddressIdValue").ID()
	expected := "/users/userIdValue/profile/addresses/itemAddressIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserProfileAddressID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileAddressId
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
			Input: "/users/userIdValue/profile/addresses",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/addresses/itemAddressIdValue",
			Expected: &UserProfileAddressId{
				UserId:        "userIdValue",
				ItemAddressId: "itemAddressIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/addresses/itemAddressIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileAddressID(v.Input)
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

		if actual.ItemAddressId != v.Expected.ItemAddressId {
			t.Fatalf("Expected %q but got %q for ItemAddressId", v.Expected.ItemAddressId, actual.ItemAddressId)
		}

	}
}

func TestParseUserProfileAddressIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileAddressId
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
			Input: "/users/userIdValue/profile/addresses",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/aDdReSsEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/addresses/itemAddressIdValue",
			Expected: &UserProfileAddressId{
				UserId:        "userIdValue",
				ItemAddressId: "itemAddressIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/addresses/itemAddressIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/aDdReSsEs/iTeMaDdReSsIdVaLuE",
			Expected: &UserProfileAddressId{
				UserId:        "uSeRiDvAlUe",
				ItemAddressId: "iTeMaDdReSsIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/aDdReSsEs/iTeMaDdReSsIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileAddressIDInsensitively(v.Input)
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

		if actual.ItemAddressId != v.Expected.ItemAddressId {
			t.Fatalf("Expected %q but got %q for ItemAddressId", v.Expected.ItemAddressId, actual.ItemAddressId)
		}

	}
}

func TestSegmentsForUserProfileAddressId(t *testing.T) {
	segments := UserProfileAddressId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserProfileAddressId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
