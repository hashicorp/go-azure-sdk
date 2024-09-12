package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileAddressId{}

func TestNewUserIdProfileAddressID(t *testing.T) {
	id := NewUserIdProfileAddressID("userIdValue", "itemAddressIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ItemAddressId != "itemAddressIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemAddressId'", id.ItemAddressId, "itemAddressIdValue")
	}
}

func TestFormatUserIdProfileAddressID(t *testing.T) {
	actual := NewUserIdProfileAddressID("userIdValue", "itemAddressIdValue").ID()
	expected := "/users/userIdValue/profile/addresses/itemAddressIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfileAddressID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileAddressId
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
			Expected: &UserIdProfileAddressId{
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

		actual, err := ParseUserIdProfileAddressID(v.Input)
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

func TestParseUserIdProfileAddressIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileAddressId
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
			Expected: &UserIdProfileAddressId{
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
			Expected: &UserIdProfileAddressId{
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

		actual, err := ParseUserIdProfileAddressIDInsensitively(v.Input)
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

func TestSegmentsForUserIdProfileAddressId(t *testing.T) {
	segments := UserIdProfileAddressId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfileAddressId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
