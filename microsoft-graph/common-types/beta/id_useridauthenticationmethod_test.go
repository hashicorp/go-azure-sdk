package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationMethodId{}

func TestNewUserIdAuthenticationMethodID(t *testing.T) {
	id := NewUserIdAuthenticationMethodID("userId", "authenticationMethodId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.AuthenticationMethodId != "authenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationMethodId'", id.AuthenticationMethodId, "authenticationMethodId")
	}
}

func TestFormatUserIdAuthenticationMethodID(t *testing.T) {
	actual := NewUserIdAuthenticationMethodID("userId", "authenticationMethodId").ID()
	expected := "/users/userId/authentication/methods/authenticationMethodId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAuthenticationMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationMethodId
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
			Input: "/users/userId/authentication/methods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/methods/authenticationMethodId",
			Expected: &UserIdAuthenticationMethodId{
				UserId:                 "userId",
				AuthenticationMethodId: "authenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/methods/authenticationMethodId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationMethodID(v.Input)
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

		if actual.AuthenticationMethodId != v.Expected.AuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for AuthenticationMethodId", v.Expected.AuthenticationMethodId, actual.AuthenticationMethodId)
		}

	}
}

func TestParseUserIdAuthenticationMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationMethodId
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
			Input: "/users/userId/authentication/methods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/mEtHoDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/methods/authenticationMethodId",
			Expected: &UserIdAuthenticationMethodId{
				UserId:                 "userId",
				AuthenticationMethodId: "authenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/methods/authenticationMethodId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/mEtHoDs/aUtHeNtIcAtIoNmEtHoDiD",
			Expected: &UserIdAuthenticationMethodId{
				UserId:                 "uSeRiD",
				AuthenticationMethodId: "aUtHeNtIcAtIoNmEtHoDiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/mEtHoDs/aUtHeNtIcAtIoNmEtHoDiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationMethodIDInsensitively(v.Input)
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

		if actual.AuthenticationMethodId != v.Expected.AuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for AuthenticationMethodId", v.Expected.AuthenticationMethodId, actual.AuthenticationMethodId)
		}

	}
}

func TestSegmentsForUserIdAuthenticationMethodId(t *testing.T) {
	segments := UserIdAuthenticationMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAuthenticationMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
