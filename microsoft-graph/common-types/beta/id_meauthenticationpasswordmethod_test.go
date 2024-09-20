package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationPasswordMethodId{}

func TestNewMeAuthenticationPasswordMethodID(t *testing.T) {
	id := NewMeAuthenticationPasswordMethodID("passwordAuthenticationMethodId")

	if id.PasswordAuthenticationMethodId != "passwordAuthenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'PasswordAuthenticationMethodId'", id.PasswordAuthenticationMethodId, "passwordAuthenticationMethodId")
	}
}

func TestFormatMeAuthenticationPasswordMethodID(t *testing.T) {
	actual := NewMeAuthenticationPasswordMethodID("passwordAuthenticationMethodId").ID()
	expected := "/me/authentication/passwordMethods/passwordAuthenticationMethodId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAuthenticationPasswordMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationPasswordMethodId
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
			Input: "/me/authentication/passwordMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/passwordMethods/passwordAuthenticationMethodId",
			Expected: &MeAuthenticationPasswordMethodId{
				PasswordAuthenticationMethodId: "passwordAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/passwordMethods/passwordAuthenticationMethodId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationPasswordMethodID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PasswordAuthenticationMethodId != v.Expected.PasswordAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for PasswordAuthenticationMethodId", v.Expected.PasswordAuthenticationMethodId, actual.PasswordAuthenticationMethodId)
		}

	}
}

func TestParseMeAuthenticationPasswordMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationPasswordMethodId
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
			Input: "/me/authentication/passwordMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/pAsSwOrDmEtHoDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/passwordMethods/passwordAuthenticationMethodId",
			Expected: &MeAuthenticationPasswordMethodId{
				PasswordAuthenticationMethodId: "passwordAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/passwordMethods/passwordAuthenticationMethodId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/pAsSwOrDmEtHoDs/pAsSwOrDaUtHeNtIcAtIoNmEtHoDiD",
			Expected: &MeAuthenticationPasswordMethodId{
				PasswordAuthenticationMethodId: "pAsSwOrDaUtHeNtIcAtIoNmEtHoDiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/pAsSwOrDmEtHoDs/pAsSwOrDaUtHeNtIcAtIoNmEtHoDiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationPasswordMethodIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PasswordAuthenticationMethodId != v.Expected.PasswordAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for PasswordAuthenticationMethodId", v.Expected.PasswordAuthenticationMethodId, actual.PasswordAuthenticationMethodId)
		}

	}
}

func TestSegmentsForMeAuthenticationPasswordMethodId(t *testing.T) {
	segments := MeAuthenticationPasswordMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAuthenticationPasswordMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
