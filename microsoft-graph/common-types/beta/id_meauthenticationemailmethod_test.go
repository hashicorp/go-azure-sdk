package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationEmailMethodId{}

func TestNewMeAuthenticationEmailMethodID(t *testing.T) {
	id := NewMeAuthenticationEmailMethodID("emailAuthenticationMethodId")

	if id.EmailAuthenticationMethodId != "emailAuthenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'EmailAuthenticationMethodId'", id.EmailAuthenticationMethodId, "emailAuthenticationMethodId")
	}
}

func TestFormatMeAuthenticationEmailMethodID(t *testing.T) {
	actual := NewMeAuthenticationEmailMethodID("emailAuthenticationMethodId").ID()
	expected := "/me/authentication/emailMethods/emailAuthenticationMethodId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAuthenticationEmailMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationEmailMethodId
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
			Input: "/me/authentication/emailMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/emailMethods/emailAuthenticationMethodId",
			Expected: &MeAuthenticationEmailMethodId{
				EmailAuthenticationMethodId: "emailAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/emailMethods/emailAuthenticationMethodId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationEmailMethodID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.EmailAuthenticationMethodId != v.Expected.EmailAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for EmailAuthenticationMethodId", v.Expected.EmailAuthenticationMethodId, actual.EmailAuthenticationMethodId)
		}

	}
}

func TestParseMeAuthenticationEmailMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationEmailMethodId
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
			Input: "/me/authentication/emailMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/eMaIlMeThOdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/emailMethods/emailAuthenticationMethodId",
			Expected: &MeAuthenticationEmailMethodId{
				EmailAuthenticationMethodId: "emailAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/emailMethods/emailAuthenticationMethodId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/eMaIlMeThOdS/eMaIlAuThEnTiCaTiOnMeThOdId",
			Expected: &MeAuthenticationEmailMethodId{
				EmailAuthenticationMethodId: "eMaIlAuThEnTiCaTiOnMeThOdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/eMaIlMeThOdS/eMaIlAuThEnTiCaTiOnMeThOdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationEmailMethodIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.EmailAuthenticationMethodId != v.Expected.EmailAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for EmailAuthenticationMethodId", v.Expected.EmailAuthenticationMethodId, actual.EmailAuthenticationMethodId)
		}

	}
}

func TestSegmentsForMeAuthenticationEmailMethodId(t *testing.T) {
	segments := MeAuthenticationEmailMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAuthenticationEmailMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
