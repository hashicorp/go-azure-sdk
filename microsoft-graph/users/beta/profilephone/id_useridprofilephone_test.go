package profilephone

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfilePhoneId{}

func TestNewUserIdProfilePhoneID(t *testing.T) {
	id := NewUserIdProfilePhoneID("userIdValue", "itemPhoneIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ItemPhoneId != "itemPhoneIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemPhoneId'", id.ItemPhoneId, "itemPhoneIdValue")
	}
}

func TestFormatUserIdProfilePhoneID(t *testing.T) {
	actual := NewUserIdProfilePhoneID("userIdValue", "itemPhoneIdValue").ID()
	expected := "/users/userIdValue/profile/phones/itemPhoneIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfilePhoneID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfilePhoneId
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
			Input: "/users/userIdValue/profile/phones",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/phones/itemPhoneIdValue",
			Expected: &UserIdProfilePhoneId{
				UserId:      "userIdValue",
				ItemPhoneId: "itemPhoneIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/phones/itemPhoneIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfilePhoneID(v.Input)
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

		if actual.ItemPhoneId != v.Expected.ItemPhoneId {
			t.Fatalf("Expected %q but got %q for ItemPhoneId", v.Expected.ItemPhoneId, actual.ItemPhoneId)
		}

	}
}

func TestParseUserIdProfilePhoneIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfilePhoneId
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
			Input: "/users/userIdValue/profile/phones",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/pHoNeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/phones/itemPhoneIdValue",
			Expected: &UserIdProfilePhoneId{
				UserId:      "userIdValue",
				ItemPhoneId: "itemPhoneIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/phones/itemPhoneIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/pHoNeS/iTeMpHoNeIdVaLuE",
			Expected: &UserIdProfilePhoneId{
				UserId:      "uSeRiDvAlUe",
				ItemPhoneId: "iTeMpHoNeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/pHoNeS/iTeMpHoNeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfilePhoneIDInsensitively(v.Input)
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

		if actual.ItemPhoneId != v.Expected.ItemPhoneId {
			t.Fatalf("Expected %q but got %q for ItemPhoneId", v.Expected.ItemPhoneId, actual.ItemPhoneId)
		}

	}
}

func TestSegmentsForUserIdProfilePhoneId(t *testing.T) {
	segments := UserIdProfilePhoneId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfilePhoneId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
