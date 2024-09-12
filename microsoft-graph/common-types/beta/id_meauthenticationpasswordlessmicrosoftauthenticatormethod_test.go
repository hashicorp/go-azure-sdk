package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{}

func TestNewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(t *testing.T) {
	id := NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodID("passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue")

	if id.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId != "passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PasswordlessMicrosoftAuthenticatorAuthenticationMethodId'", id.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId, "passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue")
	}
}

func TestFormatMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(t *testing.T) {
	actual := NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodID("passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue").ID()
	expected := "/me/authentication/passwordlessMicrosoftAuthenticatorMethods/passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication/passwordlessMicrosoftAuthenticatorMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/passwordlessMicrosoftAuthenticatorMethods/passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue",
			Expected: &MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{
				PasswordlessMicrosoftAuthenticatorAuthenticationMethodId: "passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/passwordlessMicrosoftAuthenticatorMethods/passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId != v.Expected.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for PasswordlessMicrosoftAuthenticatorAuthenticationMethodId", v.Expected.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId, actual.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId)
		}

	}
}

func TestParseMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication/passwordlessMicrosoftAuthenticatorMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/pAsSwOrDlEsSmIcRoSoFtAuThEnTiCaToRmEtHoDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/passwordlessMicrosoftAuthenticatorMethods/passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue",
			Expected: &MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{
				PasswordlessMicrosoftAuthenticatorAuthenticationMethodId: "passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/passwordlessMicrosoftAuthenticatorMethods/passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/pAsSwOrDlEsSmIcRoSoFtAuThEnTiCaToRmEtHoDs/pAsSwOrDlEsSmIcRoSoFtAuThEnTiCaToRaUtHeNtIcAtIoNmEtHoDiDvAlUe",
			Expected: &MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{
				PasswordlessMicrosoftAuthenticatorAuthenticationMethodId: "pAsSwOrDlEsSmIcRoSoFtAuThEnTiCaToRaUtHeNtIcAtIoNmEtHoDiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/pAsSwOrDlEsSmIcRoSoFtAuThEnTiCaToRmEtHoDs/pAsSwOrDlEsSmIcRoSoFtAuThEnTiCaToRaUtHeNtIcAtIoNmEtHoDiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId != v.Expected.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for PasswordlessMicrosoftAuthenticatorAuthenticationMethodId", v.Expected.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId, actual.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId)
		}

	}
}

func TestSegmentsForMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId(t *testing.T) {
	segments := MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
