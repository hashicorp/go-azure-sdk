package userauthenticationmethod

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAuthenticationMethodId{}

func TestNewUserAuthenticationMethodID(t *testing.T) {
	id := NewUserAuthenticationMethodID("userIdValue", "authenticationMethodIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.AuthenticationMethodId != "authenticationMethodIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationMethodId'", id.AuthenticationMethodId, "authenticationMethodIdValue")
	}
}

func TestFormatUserAuthenticationMethodID(t *testing.T) {
	actual := NewUserAuthenticationMethodID("userIdValue", "authenticationMethodIdValue").ID()
	expected := "/users/userIdValue/authentication/methods/authenticationMethodIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserAuthenticationMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserAuthenticationMethodId
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
			Input: "/users/userIdValue/authentication/methods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/methods/authenticationMethodIdValue",
			Expected: &UserAuthenticationMethodId{
				UserId:                 "userIdValue",
				AuthenticationMethodId: "authenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/methods/authenticationMethodIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserAuthenticationMethodID(v.Input)
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

func TestParseUserAuthenticationMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserAuthenticationMethodId
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
			Input: "/users/userIdValue/authentication/methods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/mEtHoDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/methods/authenticationMethodIdValue",
			Expected: &UserAuthenticationMethodId{
				UserId:                 "userIdValue",
				AuthenticationMethodId: "authenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/methods/authenticationMethodIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/mEtHoDs/aUtHeNtIcAtIoNmEtHoDiDvAlUe",
			Expected: &UserAuthenticationMethodId{
				UserId:                 "uSeRiDvAlUe",
				AuthenticationMethodId: "aUtHeNtIcAtIoNmEtHoDiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/mEtHoDs/aUtHeNtIcAtIoNmEtHoDiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserAuthenticationMethodIDInsensitively(v.Input)
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

func TestSegmentsForUserAuthenticationMethodId(t *testing.T) {
	segments := UserAuthenticationMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserAuthenticationMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
