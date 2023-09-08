package userauthenticationpasswordmethod

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAuthenticationPasswordMethodId{}

func TestNewUserAuthenticationPasswordMethodID(t *testing.T) {
	id := NewUserAuthenticationPasswordMethodID("userIdValue", "passwordAuthenticationMethodIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.PasswordAuthenticationMethodId != "passwordAuthenticationMethodIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PasswordAuthenticationMethodId'", id.PasswordAuthenticationMethodId, "passwordAuthenticationMethodIdValue")
	}
}

func TestFormatUserAuthenticationPasswordMethodID(t *testing.T) {
	actual := NewUserAuthenticationPasswordMethodID("userIdValue", "passwordAuthenticationMethodIdValue").ID()
	expected := "/users/userIdValue/authentication/passwordMethods/passwordAuthenticationMethodIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserAuthenticationPasswordMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserAuthenticationPasswordMethodId
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
			Input: "/users/userIdValue/authentication/passwordMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/passwordMethods/passwordAuthenticationMethodIdValue",
			Expected: &UserAuthenticationPasswordMethodId{
				UserId:                         "userIdValue",
				PasswordAuthenticationMethodId: "passwordAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/passwordMethods/passwordAuthenticationMethodIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserAuthenticationPasswordMethodID(v.Input)
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

		if actual.PasswordAuthenticationMethodId != v.Expected.PasswordAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for PasswordAuthenticationMethodId", v.Expected.PasswordAuthenticationMethodId, actual.PasswordAuthenticationMethodId)
		}

	}
}

func TestParseUserAuthenticationPasswordMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserAuthenticationPasswordMethodId
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
			Input: "/users/userIdValue/authentication/passwordMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/pAsSwOrDmEtHoDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/passwordMethods/passwordAuthenticationMethodIdValue",
			Expected: &UserAuthenticationPasswordMethodId{
				UserId:                         "userIdValue",
				PasswordAuthenticationMethodId: "passwordAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/passwordMethods/passwordAuthenticationMethodIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/pAsSwOrDmEtHoDs/pAsSwOrDaUtHeNtIcAtIoNmEtHoDiDvAlUe",
			Expected: &UserAuthenticationPasswordMethodId{
				UserId:                         "uSeRiDvAlUe",
				PasswordAuthenticationMethodId: "pAsSwOrDaUtHeNtIcAtIoNmEtHoDiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/pAsSwOrDmEtHoDs/pAsSwOrDaUtHeNtIcAtIoNmEtHoDiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserAuthenticationPasswordMethodIDInsensitively(v.Input)
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

		if actual.PasswordAuthenticationMethodId != v.Expected.PasswordAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for PasswordAuthenticationMethodId", v.Expected.PasswordAuthenticationMethodId, actual.PasswordAuthenticationMethodId)
		}

	}
}

func TestSegmentsForUserAuthenticationPasswordMethodId(t *testing.T) {
	segments := UserAuthenticationPasswordMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserAuthenticationPasswordMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
