package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileAccountId{}

func TestNewUserIdProfileAccountID(t *testing.T) {
	id := NewUserIdProfileAccountID("userId", "userAccountInformationId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.UserAccountInformationId != "userAccountInformationId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserAccountInformationId'", id.UserAccountInformationId, "userAccountInformationId")
	}
}

func TestFormatUserIdProfileAccountID(t *testing.T) {
	actual := NewUserIdProfileAccountID("userId", "userAccountInformationId").ID()
	expected := "/users/userId/profile/account/userAccountInformationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfileAccountID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileAccountId
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
			Input: "/users/userId/profile/account",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/account/userAccountInformationId",
			Expected: &UserIdProfileAccountId{
				UserId:                   "userId",
				UserAccountInformationId: "userAccountInformationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/account/userAccountInformationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileAccountID(v.Input)
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

func TestParseUserIdProfileAccountIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileAccountId
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
			Input: "/users/userId/profile/account",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/aCcOuNt",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/account/userAccountInformationId",
			Expected: &UserIdProfileAccountId{
				UserId:                   "userId",
				UserAccountInformationId: "userAccountInformationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/account/userAccountInformationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/aCcOuNt/uSeRaCcOuNtInFoRmAtIoNiD",
			Expected: &UserIdProfileAccountId{
				UserId:                   "uSeRiD",
				UserAccountInformationId: "uSeRaCcOuNtInFoRmAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/aCcOuNt/uSeRaCcOuNtInFoRmAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileAccountIDInsensitively(v.Input)
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

func TestSegmentsForUserIdProfileAccountId(t *testing.T) {
	segments := UserIdProfileAccountId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfileAccountId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
