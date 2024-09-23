package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileWebAccountId{}

func TestNewUserIdProfileWebAccountID(t *testing.T) {
	id := NewUserIdProfileWebAccountID("userId", "webAccountId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.WebAccountId != "webAccountId" {
		t.Fatalf("Expected %q but got %q for Segment 'WebAccountId'", id.WebAccountId, "webAccountId")
	}
}

func TestFormatUserIdProfileWebAccountID(t *testing.T) {
	actual := NewUserIdProfileWebAccountID("userId", "webAccountId").ID()
	expected := "/users/userId/profile/webAccounts/webAccountId"
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
			Input: "/users/userId/profile/webAccounts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/webAccounts/webAccountId",
			Expected: &UserIdProfileWebAccountId{
				UserId:       "userId",
				WebAccountId: "webAccountId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/webAccounts/webAccountId/extra",
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
			Input: "/users/userId/profile/webAccounts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/wEbAcCoUnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/webAccounts/webAccountId",
			Expected: &UserIdProfileWebAccountId{
				UserId:       "userId",
				WebAccountId: "webAccountId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/webAccounts/webAccountId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/wEbAcCoUnTs/wEbAcCoUnTiD",
			Expected: &UserIdProfileWebAccountId{
				UserId:       "uSeRiD",
				WebAccountId: "wEbAcCoUnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/wEbAcCoUnTs/wEbAcCoUnTiD/extra",
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
