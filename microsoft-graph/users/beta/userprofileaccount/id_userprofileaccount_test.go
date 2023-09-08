package userprofileaccount

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileAccountId{}

func TestNewUserProfileAccountID(t *testing.T) {
	id := NewUserProfileAccountID("userIdValue", "userAccountInformationIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.UserAccountInformationId != "userAccountInformationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserAccountInformationId'", id.UserAccountInformationId, "userAccountInformationIdValue")
	}
}

func TestFormatUserProfileAccountID(t *testing.T) {
	actual := NewUserProfileAccountID("userIdValue", "userAccountInformationIdValue").ID()
	expected := "/users/userIdValue/profile/account/userAccountInformationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserProfileAccountID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileAccountId
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
			Input: "/users/userIdValue/profile/account",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/account/userAccountInformationIdValue",
			Expected: &UserProfileAccountId{
				UserId:                   "userIdValue",
				UserAccountInformationId: "userAccountInformationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/account/userAccountInformationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileAccountID(v.Input)
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

		if actual.UserAccountInformationId != v.Expected.UserAccountInformationId {
			t.Fatalf("Expected %q but got %q for UserAccountInformationId", v.Expected.UserAccountInformationId, actual.UserAccountInformationId)
		}

	}
}

func TestParseUserProfileAccountIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileAccountId
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
			Input: "/users/userIdValue/profile/account",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/aCcOuNt",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/account/userAccountInformationIdValue",
			Expected: &UserProfileAccountId{
				UserId:                   "userIdValue",
				UserAccountInformationId: "userAccountInformationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/account/userAccountInformationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/aCcOuNt/uSeRaCcOuNtInFoRmAtIoNiDvAlUe",
			Expected: &UserProfileAccountId{
				UserId:                   "uSeRiDvAlUe",
				UserAccountInformationId: "uSeRaCcOuNtInFoRmAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/aCcOuNt/uSeRaCcOuNtInFoRmAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileAccountIDInsensitively(v.Input)
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

		if actual.UserAccountInformationId != v.Expected.UserAccountInformationId {
			t.Fatalf("Expected %q but got %q for UserAccountInformationId", v.Expected.UserAccountInformationId, actual.UserAccountInformationId)
		}

	}
}

func TestSegmentsForUserProfileAccountId(t *testing.T) {
	segments := UserProfileAccountId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserProfileAccountId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
