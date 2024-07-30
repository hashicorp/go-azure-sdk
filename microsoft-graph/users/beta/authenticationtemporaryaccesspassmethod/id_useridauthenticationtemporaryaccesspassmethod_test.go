package authenticationtemporaryaccesspassmethod

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationTemporaryAccessPassMethodId{}

func TestNewUserIdAuthenticationTemporaryAccessPassMethodID(t *testing.T) {
	id := NewUserIdAuthenticationTemporaryAccessPassMethodID("userIdValue", "temporaryAccessPassAuthenticationMethodIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.TemporaryAccessPassAuthenticationMethodId != "temporaryAccessPassAuthenticationMethodIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TemporaryAccessPassAuthenticationMethodId'", id.TemporaryAccessPassAuthenticationMethodId, "temporaryAccessPassAuthenticationMethodIdValue")
	}
}

func TestFormatUserIdAuthenticationTemporaryAccessPassMethodID(t *testing.T) {
	actual := NewUserIdAuthenticationTemporaryAccessPassMethodID("userIdValue", "temporaryAccessPassAuthenticationMethodIdValue").ID()
	expected := "/users/userIdValue/authentication/temporaryAccessPassMethods/temporaryAccessPassAuthenticationMethodIdValue"
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/authentication",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/authentication/temporaryAccessPassMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/temporaryAccessPassMethods/temporaryAccessPassAuthenticationMethodIdValue",
			Expected: &UserIdAuthenticationTemporaryAccessPassMethodId{
				UserId: "userIdValue",
				TemporaryAccessPassAuthenticationMethodId: "temporaryAccessPassAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/temporaryAccessPassMethods/temporaryAccessPassAuthenticationMethodIdValue/extra",
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
			Input: "/users/userIdValue/authentication",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/authentication/temporaryAccessPassMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/tEmPoRaRyAcCeSsPaSsMeThOdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/temporaryAccessPassMethods/temporaryAccessPassAuthenticationMethodIdValue",
			Expected: &UserIdAuthenticationTemporaryAccessPassMethodId{
				UserId: "userIdValue",
				TemporaryAccessPassAuthenticationMethodId: "temporaryAccessPassAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/temporaryAccessPassMethods/temporaryAccessPassAuthenticationMethodIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/tEmPoRaRyAcCeSsPaSsMeThOdS/tEmPoRaRyAcCeSsPaSsAuThEnTiCaTiOnMeThOdIdVaLuE",
			Expected: &UserIdAuthenticationTemporaryAccessPassMethodId{
				UserId: "uSeRiDvAlUe",
				TemporaryAccessPassAuthenticationMethodId: "tEmPoRaRyAcCeSsPaSsAuThEnTiCaTiOnMeThOdIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/tEmPoRaRyAcCeSsPaSsMeThOdS/tEmPoRaRyAcCeSsPaSsAuThEnTiCaTiOnMeThOdIdVaLuE/extra",
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
