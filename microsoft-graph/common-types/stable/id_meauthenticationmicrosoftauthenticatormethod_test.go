package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationMicrosoftAuthenticatorMethodId{}

func TestNewMeAuthenticationMicrosoftAuthenticatorMethodID(t *testing.T) {
	id := NewMeAuthenticationMicrosoftAuthenticatorMethodID("microsoftAuthenticatorAuthenticationMethodId")

	if id.MicrosoftAuthenticatorAuthenticationMethodId != "microsoftAuthenticatorAuthenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'MicrosoftAuthenticatorAuthenticationMethodId'", id.MicrosoftAuthenticatorAuthenticationMethodId, "microsoftAuthenticatorAuthenticationMethodId")
	}
}

func TestFormatMeAuthenticationMicrosoftAuthenticatorMethodID(t *testing.T) {
	actual := NewMeAuthenticationMicrosoftAuthenticatorMethodID("microsoftAuthenticatorAuthenticationMethodId").ID()
	expected := "/me/authentication/microsoftAuthenticatorMethods/microsoftAuthenticatorAuthenticationMethodId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAuthenticationMicrosoftAuthenticatorMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationMicrosoftAuthenticatorMethodId
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
			Input: "/me/authentication/microsoftAuthenticatorMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/microsoftAuthenticatorMethods/microsoftAuthenticatorAuthenticationMethodId",
			Expected: &MeAuthenticationMicrosoftAuthenticatorMethodId{
				MicrosoftAuthenticatorAuthenticationMethodId: "microsoftAuthenticatorAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/microsoftAuthenticatorMethods/microsoftAuthenticatorAuthenticationMethodId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationMicrosoftAuthenticatorMethodID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MicrosoftAuthenticatorAuthenticationMethodId != v.Expected.MicrosoftAuthenticatorAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for MicrosoftAuthenticatorAuthenticationMethodId", v.Expected.MicrosoftAuthenticatorAuthenticationMethodId, actual.MicrosoftAuthenticatorAuthenticationMethodId)
		}

	}
}

func TestParseMeAuthenticationMicrosoftAuthenticatorMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationMicrosoftAuthenticatorMethodId
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
			Input: "/me/authentication/microsoftAuthenticatorMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/mIcRoSoFtAuThEnTiCaToRmEtHoDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/microsoftAuthenticatorMethods/microsoftAuthenticatorAuthenticationMethodId",
			Expected: &MeAuthenticationMicrosoftAuthenticatorMethodId{
				MicrosoftAuthenticatorAuthenticationMethodId: "microsoftAuthenticatorAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/microsoftAuthenticatorMethods/microsoftAuthenticatorAuthenticationMethodId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/mIcRoSoFtAuThEnTiCaToRmEtHoDs/mIcRoSoFtAuThEnTiCaToRaUtHeNtIcAtIoNmEtHoDiD",
			Expected: &MeAuthenticationMicrosoftAuthenticatorMethodId{
				MicrosoftAuthenticatorAuthenticationMethodId: "mIcRoSoFtAuThEnTiCaToRaUtHeNtIcAtIoNmEtHoDiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/mIcRoSoFtAuThEnTiCaToRmEtHoDs/mIcRoSoFtAuThEnTiCaToRaUtHeNtIcAtIoNmEtHoDiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationMicrosoftAuthenticatorMethodIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MicrosoftAuthenticatorAuthenticationMethodId != v.Expected.MicrosoftAuthenticatorAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for MicrosoftAuthenticatorAuthenticationMethodId", v.Expected.MicrosoftAuthenticatorAuthenticationMethodId, actual.MicrosoftAuthenticatorAuthenticationMethodId)
		}

	}
}

func TestSegmentsForMeAuthenticationMicrosoftAuthenticatorMethodId(t *testing.T) {
	segments := MeAuthenticationMicrosoftAuthenticatorMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAuthenticationMicrosoftAuthenticatorMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
