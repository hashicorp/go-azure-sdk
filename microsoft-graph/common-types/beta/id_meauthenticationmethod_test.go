package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationMethodId{}

func TestNewMeAuthenticationMethodID(t *testing.T) {
	id := NewMeAuthenticationMethodID("authenticationMethodId")

	if id.AuthenticationMethodId != "authenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationMethodId'", id.AuthenticationMethodId, "authenticationMethodId")
	}
}

func TestFormatMeAuthenticationMethodID(t *testing.T) {
	actual := NewMeAuthenticationMethodID("authenticationMethodId").ID()
	expected := "/me/authentication/methods/authenticationMethodId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAuthenticationMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationMethodId
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
			Input: "/me/authentication/methods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/methods/authenticationMethodId",
			Expected: &MeAuthenticationMethodId{
				AuthenticationMethodId: "authenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/methods/authenticationMethodId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationMethodID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationMethodId != v.Expected.AuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for AuthenticationMethodId", v.Expected.AuthenticationMethodId, actual.AuthenticationMethodId)
		}

	}
}

func TestParseMeAuthenticationMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationMethodId
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
			Input: "/me/authentication/methods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/mEtHoDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/methods/authenticationMethodId",
			Expected: &MeAuthenticationMethodId{
				AuthenticationMethodId: "authenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/methods/authenticationMethodId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/mEtHoDs/aUtHeNtIcAtIoNmEtHoDiD",
			Expected: &MeAuthenticationMethodId{
				AuthenticationMethodId: "aUtHeNtIcAtIoNmEtHoDiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/mEtHoDs/aUtHeNtIcAtIoNmEtHoDiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationMethodIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationMethodId != v.Expected.AuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for AuthenticationMethodId", v.Expected.AuthenticationMethodId, actual.AuthenticationMethodId)
		}

	}
}

func TestSegmentsForMeAuthenticationMethodId(t *testing.T) {
	segments := MeAuthenticationMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAuthenticationMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
