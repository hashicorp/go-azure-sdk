package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationFido2MethodId{}

func TestNewMeAuthenticationFido2MethodID(t *testing.T) {
	id := NewMeAuthenticationFido2MethodID("fido2AuthenticationMethodId")

	if id.Fido2AuthenticationMethodId != "fido2AuthenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'Fido2AuthenticationMethodId'", id.Fido2AuthenticationMethodId, "fido2AuthenticationMethodId")
	}
}

func TestFormatMeAuthenticationFido2MethodID(t *testing.T) {
	actual := NewMeAuthenticationFido2MethodID("fido2AuthenticationMethodId").ID()
	expected := "/me/authentication/fido2Methods/fido2AuthenticationMethodId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAuthenticationFido2MethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationFido2MethodId
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
			Input: "/me/authentication/fido2Methods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/fido2Methods/fido2AuthenticationMethodId",
			Expected: &MeAuthenticationFido2MethodId{
				Fido2AuthenticationMethodId: "fido2AuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/fido2Methods/fido2AuthenticationMethodId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationFido2MethodID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Fido2AuthenticationMethodId != v.Expected.Fido2AuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for Fido2AuthenticationMethodId", v.Expected.Fido2AuthenticationMethodId, actual.Fido2AuthenticationMethodId)
		}

	}
}

func TestParseMeAuthenticationFido2MethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationFido2MethodId
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
			Input: "/me/authentication/fido2Methods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/fIdO2MeThOdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/fido2Methods/fido2AuthenticationMethodId",
			Expected: &MeAuthenticationFido2MethodId{
				Fido2AuthenticationMethodId: "fido2AuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/fido2Methods/fido2AuthenticationMethodId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/fIdO2MeThOdS/fIdO2AuThEnTiCaTiOnMeThOdId",
			Expected: &MeAuthenticationFido2MethodId{
				Fido2AuthenticationMethodId: "fIdO2AuThEnTiCaTiOnMeThOdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/fIdO2MeThOdS/fIdO2AuThEnTiCaTiOnMeThOdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationFido2MethodIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Fido2AuthenticationMethodId != v.Expected.Fido2AuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for Fido2AuthenticationMethodId", v.Expected.Fido2AuthenticationMethodId, actual.Fido2AuthenticationMethodId)
		}

	}
}

func TestSegmentsForMeAuthenticationFido2MethodId(t *testing.T) {
	segments := MeAuthenticationFido2MethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAuthenticationFido2MethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
