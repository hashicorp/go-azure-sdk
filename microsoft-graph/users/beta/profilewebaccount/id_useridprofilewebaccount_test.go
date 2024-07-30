package profilewebaccount

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileWebAccountId{}

func TestNewUserIdProfileWebAccountID(t *testing.T) {
	id := NewUserIdProfileWebAccountID("userIdValue", "webAccountIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.WebAccountId != "webAccountIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WebAccountId'", id.WebAccountId, "webAccountIdValue")
	}
}

func TestFormatUserIdProfileWebAccountID(t *testing.T) {
	actual := NewUserIdProfileWebAccountID("userIdValue", "webAccountIdValue").ID()
	expected := "/users/userIdValue/profile/webAccounts/webAccountIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfileWebAccountID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileWebAccountId
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
			Input: "/users/userIdValue/profile/webAccounts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/webAccounts/webAccountIdValue",
			Expected: &UserIdProfileWebAccountId{
				UserId:       "userIdValue",
				WebAccountId: "webAccountIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/webAccounts/webAccountIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileWebAccountID(v.Input)
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

		if actual.WebAccountId != v.Expected.WebAccountId {
			t.Fatalf("Expected %q but got %q for WebAccountId", v.Expected.WebAccountId, actual.WebAccountId)
		}

	}
}

func TestParseUserIdProfileWebAccountIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileWebAccountId
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
			Input: "/users/userIdValue/profile/webAccounts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/wEbAcCoUnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/webAccounts/webAccountIdValue",
			Expected: &UserIdProfileWebAccountId{
				UserId:       "userIdValue",
				WebAccountId: "webAccountIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/webAccounts/webAccountIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/wEbAcCoUnTs/wEbAcCoUnTiDvAlUe",
			Expected: &UserIdProfileWebAccountId{
				UserId:       "uSeRiDvAlUe",
				WebAccountId: "wEbAcCoUnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/wEbAcCoUnTs/wEbAcCoUnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileWebAccountIDInsensitively(v.Input)
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

		if actual.WebAccountId != v.Expected.WebAccountId {
			t.Fatalf("Expected %q but got %q for WebAccountId", v.Expected.WebAccountId, actual.WebAccountId)
		}

	}
}

func TestSegmentsForUserIdProfileWebAccountId(t *testing.T) {
	segments := UserIdProfileWebAccountId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfileWebAccountId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
