package authenticationpasswordlessmicrosoftauthenticatormethod

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{}

func TestNewUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(t *testing.T) {
	id := NewUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodID("userIdValue", "passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId != "passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PasswordlessMicrosoftAuthenticatorAuthenticationMethodId'", id.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId, "passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue")
	}
}

func TestFormatUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(t *testing.T) {
	actual := NewUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodID("userIdValue", "passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue").ID()
	expected := "/users/userIdValue/authentication/passwordlessMicrosoftAuthenticatorMethods/passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId
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
			Input: "/users/userIdValue/authentication/passwordlessMicrosoftAuthenticatorMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/passwordlessMicrosoftAuthenticatorMethods/passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue",
			Expected: &UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{
				UserId: "userIdValue",
				PasswordlessMicrosoftAuthenticatorAuthenticationMethodId: "passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/passwordlessMicrosoftAuthenticatorMethods/passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(v.Input)
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

		if actual.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId != v.Expected.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for PasswordlessMicrosoftAuthenticatorAuthenticationMethodId", v.Expected.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId, actual.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId)
		}

	}
}

func TestParseUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId
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
			Input: "/users/userIdValue/authentication/passwordlessMicrosoftAuthenticatorMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/pAsSwOrDlEsSmIcRoSoFtAuThEnTiCaToRmEtHoDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/passwordlessMicrosoftAuthenticatorMethods/passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue",
			Expected: &UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{
				UserId: "userIdValue",
				PasswordlessMicrosoftAuthenticatorAuthenticationMethodId: "passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/passwordlessMicrosoftAuthenticatorMethods/passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/pAsSwOrDlEsSmIcRoSoFtAuThEnTiCaToRmEtHoDs/pAsSwOrDlEsSmIcRoSoFtAuThEnTiCaToRaUtHeNtIcAtIoNmEtHoDiDvAlUe",
			Expected: &UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{
				UserId: "uSeRiDvAlUe",
				PasswordlessMicrosoftAuthenticatorAuthenticationMethodId: "pAsSwOrDlEsSmIcRoSoFtAuThEnTiCaToRaUtHeNtIcAtIoNmEtHoDiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/pAsSwOrDlEsSmIcRoSoFtAuThEnTiCaToRmEtHoDs/pAsSwOrDlEsSmIcRoSoFtAuThEnTiCaToRaUtHeNtIcAtIoNmEtHoDiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodIDInsensitively(v.Input)
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

		if actual.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId != v.Expected.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for PasswordlessMicrosoftAuthenticatorAuthenticationMethodId", v.Expected.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId, actual.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId)
		}

	}
}

func TestSegmentsForUserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId(t *testing.T) {
	segments := UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAuthenticationPasswordlessMicrosoftAuthenticatorMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
