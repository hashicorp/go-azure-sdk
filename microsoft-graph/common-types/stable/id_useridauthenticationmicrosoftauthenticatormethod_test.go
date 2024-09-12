package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationMicrosoftAuthenticatorMethodId{}

func TestNewUserIdAuthenticationMicrosoftAuthenticatorMethodID(t *testing.T) {
	id := NewUserIdAuthenticationMicrosoftAuthenticatorMethodID("userIdValue", "microsoftAuthenticatorAuthenticationMethodIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.MicrosoftAuthenticatorAuthenticationMethodId != "microsoftAuthenticatorAuthenticationMethodIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MicrosoftAuthenticatorAuthenticationMethodId'", id.MicrosoftAuthenticatorAuthenticationMethodId, "microsoftAuthenticatorAuthenticationMethodIdValue")
	}
}

func TestFormatUserIdAuthenticationMicrosoftAuthenticatorMethodID(t *testing.T) {
	actual := NewUserIdAuthenticationMicrosoftAuthenticatorMethodID("userIdValue", "microsoftAuthenticatorAuthenticationMethodIdValue").ID()
	expected := "/users/userIdValue/authentication/microsoftAuthenticatorMethods/microsoftAuthenticatorAuthenticationMethodIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAuthenticationMicrosoftAuthenticatorMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationMicrosoftAuthenticatorMethodId
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
			Input: "/users/userIdValue/authentication/microsoftAuthenticatorMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/microsoftAuthenticatorMethods/microsoftAuthenticatorAuthenticationMethodIdValue",
			Expected: &UserIdAuthenticationMicrosoftAuthenticatorMethodId{
				UserId: "userIdValue",
				MicrosoftAuthenticatorAuthenticationMethodId: "microsoftAuthenticatorAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/microsoftAuthenticatorMethods/microsoftAuthenticatorAuthenticationMethodIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationMicrosoftAuthenticatorMethodID(v.Input)
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

		if actual.MicrosoftAuthenticatorAuthenticationMethodId != v.Expected.MicrosoftAuthenticatorAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for MicrosoftAuthenticatorAuthenticationMethodId", v.Expected.MicrosoftAuthenticatorAuthenticationMethodId, actual.MicrosoftAuthenticatorAuthenticationMethodId)
		}

	}
}

func TestParseUserIdAuthenticationMicrosoftAuthenticatorMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationMicrosoftAuthenticatorMethodId
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
			Input: "/users/userIdValue/authentication/microsoftAuthenticatorMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/mIcRoSoFtAuThEnTiCaToRmEtHoDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/microsoftAuthenticatorMethods/microsoftAuthenticatorAuthenticationMethodIdValue",
			Expected: &UserIdAuthenticationMicrosoftAuthenticatorMethodId{
				UserId: "userIdValue",
				MicrosoftAuthenticatorAuthenticationMethodId: "microsoftAuthenticatorAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/microsoftAuthenticatorMethods/microsoftAuthenticatorAuthenticationMethodIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/mIcRoSoFtAuThEnTiCaToRmEtHoDs/mIcRoSoFtAuThEnTiCaToRaUtHeNtIcAtIoNmEtHoDiDvAlUe",
			Expected: &UserIdAuthenticationMicrosoftAuthenticatorMethodId{
				UserId: "uSeRiDvAlUe",
				MicrosoftAuthenticatorAuthenticationMethodId: "mIcRoSoFtAuThEnTiCaToRaUtHeNtIcAtIoNmEtHoDiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/mIcRoSoFtAuThEnTiCaToRmEtHoDs/mIcRoSoFtAuThEnTiCaToRaUtHeNtIcAtIoNmEtHoDiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationMicrosoftAuthenticatorMethodIDInsensitively(v.Input)
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

		if actual.MicrosoftAuthenticatorAuthenticationMethodId != v.Expected.MicrosoftAuthenticatorAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for MicrosoftAuthenticatorAuthenticationMethodId", v.Expected.MicrosoftAuthenticatorAuthenticationMethodId, actual.MicrosoftAuthenticatorAuthenticationMethodId)
		}

	}
}

func TestSegmentsForUserIdAuthenticationMicrosoftAuthenticatorMethodId(t *testing.T) {
	segments := UserIdAuthenticationMicrosoftAuthenticatorMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAuthenticationMicrosoftAuthenticatorMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
