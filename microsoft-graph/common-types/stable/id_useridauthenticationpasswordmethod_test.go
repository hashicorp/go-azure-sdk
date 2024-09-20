package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationPasswordMethodId{}

func TestNewUserIdAuthenticationPasswordMethodID(t *testing.T) {
	id := NewUserIdAuthenticationPasswordMethodID("userId", "passwordAuthenticationMethodId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.PasswordAuthenticationMethodId != "passwordAuthenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'PasswordAuthenticationMethodId'", id.PasswordAuthenticationMethodId, "passwordAuthenticationMethodId")
	}
}

func TestFormatUserIdAuthenticationPasswordMethodID(t *testing.T) {
	actual := NewUserIdAuthenticationPasswordMethodID("userId", "passwordAuthenticationMethodId").ID()
	expected := "/users/userId/authentication/passwordMethods/passwordAuthenticationMethodId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAuthenticationPasswordMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationPasswordMethodId
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
			Input: "/users/userId/authentication/passwordMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/passwordMethods/passwordAuthenticationMethodId",
			Expected: &UserIdAuthenticationPasswordMethodId{
				UserId:                         "userId",
				PasswordAuthenticationMethodId: "passwordAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/passwordMethods/passwordAuthenticationMethodId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationPasswordMethodID(v.Input)
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

func TestParseUserIdAuthenticationPasswordMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationPasswordMethodId
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
			Input: "/users/userId/authentication/passwordMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/pAsSwOrDmEtHoDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/passwordMethods/passwordAuthenticationMethodId",
			Expected: &UserIdAuthenticationPasswordMethodId{
				UserId:                         "userId",
				PasswordAuthenticationMethodId: "passwordAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/passwordMethods/passwordAuthenticationMethodId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/pAsSwOrDmEtHoDs/pAsSwOrDaUtHeNtIcAtIoNmEtHoDiD",
			Expected: &UserIdAuthenticationPasswordMethodId{
				UserId:                         "uSeRiD",
				PasswordAuthenticationMethodId: "pAsSwOrDaUtHeNtIcAtIoNmEtHoDiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/pAsSwOrDmEtHoDs/pAsSwOrDaUtHeNtIcAtIoNmEtHoDiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationPasswordMethodIDInsensitively(v.Input)
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

func TestSegmentsForUserIdAuthenticationPasswordMethodId(t *testing.T) {
	segments := UserIdAuthenticationPasswordMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAuthenticationPasswordMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
