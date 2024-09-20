package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationTemporaryAccessPassMethodId{}

func TestNewUserIdAuthenticationTemporaryAccessPassMethodID(t *testing.T) {
	id := NewUserIdAuthenticationTemporaryAccessPassMethodID("userId", "temporaryAccessPassAuthenticationMethodId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.TemporaryAccessPassAuthenticationMethodId != "temporaryAccessPassAuthenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'TemporaryAccessPassAuthenticationMethodId'", id.TemporaryAccessPassAuthenticationMethodId, "temporaryAccessPassAuthenticationMethodId")
	}
}

func TestFormatUserIdAuthenticationTemporaryAccessPassMethodID(t *testing.T) {
	actual := NewUserIdAuthenticationTemporaryAccessPassMethodID("userId", "temporaryAccessPassAuthenticationMethodId").ID()
	expected := "/users/userId/authentication/temporaryAccessPassMethods/temporaryAccessPassAuthenticationMethodId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAuthenticationTemporaryAccessPassMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationTemporaryAccessPassMethodId
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
			Input: "/users/userId/authentication",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/authentication/temporaryAccessPassMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/temporaryAccessPassMethods/temporaryAccessPassAuthenticationMethodId",
			Expected: &UserIdAuthenticationTemporaryAccessPassMethodId{
				UserId: "userId",
				TemporaryAccessPassAuthenticationMethodId: "temporaryAccessPassAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/temporaryAccessPassMethods/temporaryAccessPassAuthenticationMethodId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationTemporaryAccessPassMethodID(v.Input)
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

		if actual.TemporaryAccessPassAuthenticationMethodId != v.Expected.TemporaryAccessPassAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for TemporaryAccessPassAuthenticationMethodId", v.Expected.TemporaryAccessPassAuthenticationMethodId, actual.TemporaryAccessPassAuthenticationMethodId)
		}

	}
}

func TestParseUserIdAuthenticationTemporaryAccessPassMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationTemporaryAccessPassMethodId
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
			Input: "/users/userId/authentication",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/authentication/temporaryAccessPassMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/tEmPoRaRyAcCeSsPaSsMeThOdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/temporaryAccessPassMethods/temporaryAccessPassAuthenticationMethodId",
			Expected: &UserIdAuthenticationTemporaryAccessPassMethodId{
				UserId: "userId",
				TemporaryAccessPassAuthenticationMethodId: "temporaryAccessPassAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/temporaryAccessPassMethods/temporaryAccessPassAuthenticationMethodId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/tEmPoRaRyAcCeSsPaSsMeThOdS/tEmPoRaRyAcCeSsPaSsAuThEnTiCaTiOnMeThOdId",
			Expected: &UserIdAuthenticationTemporaryAccessPassMethodId{
				UserId: "uSeRiD",
				TemporaryAccessPassAuthenticationMethodId: "tEmPoRaRyAcCeSsPaSsAuThEnTiCaTiOnMeThOdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/tEmPoRaRyAcCeSsPaSsMeThOdS/tEmPoRaRyAcCeSsPaSsAuThEnTiCaTiOnMeThOdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationTemporaryAccessPassMethodIDInsensitively(v.Input)
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

		if actual.TemporaryAccessPassAuthenticationMethodId != v.Expected.TemporaryAccessPassAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for TemporaryAccessPassAuthenticationMethodId", v.Expected.TemporaryAccessPassAuthenticationMethodId, actual.TemporaryAccessPassAuthenticationMethodId)
		}

	}
}

func TestSegmentsForUserIdAuthenticationTemporaryAccessPassMethodId(t *testing.T) {
	segments := UserIdAuthenticationTemporaryAccessPassMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAuthenticationTemporaryAccessPassMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
