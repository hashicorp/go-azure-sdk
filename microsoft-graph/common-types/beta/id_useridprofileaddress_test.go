package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileAddressId{}

func TestNewUserIdProfileAddressID(t *testing.T) {
	id := NewUserIdProfileAddressID("userId", "itemAddressId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ItemAddressId != "itemAddressId" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemAddressId'", id.ItemAddressId, "itemAddressId")
	}
}

func TestFormatUserIdProfileAddressID(t *testing.T) {
	actual := NewUserIdProfileAddressID("userId", "itemAddressId").ID()
	expected := "/users/userId/profile/addresses/itemAddressId"
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
			Input: "/users/userId/profile/addresses",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/addresses/itemAddressId",
			Expected: &UserIdProfileAddressId{
				UserId:        "userId",
				ItemAddressId: "itemAddressId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/addresses/itemAddressId/extra",
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
			Input: "/users/userId/profile/addresses",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/aDdReSsEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/addresses/itemAddressId",
			Expected: &UserIdProfileAddressId{
				UserId:        "userId",
				ItemAddressId: "itemAddressId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/addresses/itemAddressId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/aDdReSsEs/iTeMaDdReSsId",
			Expected: &UserIdProfileAddressId{
				UserId:        "uSeRiD",
				ItemAddressId: "iTeMaDdReSsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/aDdReSsEs/iTeMaDdReSsId/extra",
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
