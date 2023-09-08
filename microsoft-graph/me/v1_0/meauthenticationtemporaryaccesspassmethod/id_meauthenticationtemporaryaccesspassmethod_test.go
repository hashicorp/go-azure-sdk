package meauthenticationtemporaryaccesspassmethod

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeAuthenticationTemporaryAccessPassMethodId{}

func TestNewMeAuthenticationTemporaryAccessPassMethodID(t *testing.T) {
	id := NewMeAuthenticationTemporaryAccessPassMethodID("temporaryAccessPassAuthenticationMethodIdValue")

	if id.TemporaryAccessPassAuthenticationMethodId != "temporaryAccessPassAuthenticationMethodIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TemporaryAccessPassAuthenticationMethodId'", id.TemporaryAccessPassAuthenticationMethodId, "temporaryAccessPassAuthenticationMethodIdValue")
	}
}

func TestFormatMeAuthenticationTemporaryAccessPassMethodID(t *testing.T) {
	actual := NewMeAuthenticationTemporaryAccessPassMethodID("temporaryAccessPassAuthenticationMethodIdValue").ID()
	expected := "/me/authentication/temporaryAccessPassMethods/temporaryAccessPassAuthenticationMethodIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAuthenticationTemporaryAccessPassMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationTemporaryAccessPassMethodId
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
			Input: "/me/authentication/temporaryAccessPassMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/temporaryAccessPassMethods/temporaryAccessPassAuthenticationMethodIdValue",
			Expected: &MeAuthenticationTemporaryAccessPassMethodId{
				TemporaryAccessPassAuthenticationMethodId: "temporaryAccessPassAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/temporaryAccessPassMethods/temporaryAccessPassAuthenticationMethodIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationTemporaryAccessPassMethodID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TemporaryAccessPassAuthenticationMethodId != v.Expected.TemporaryAccessPassAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for TemporaryAccessPassAuthenticationMethodId", v.Expected.TemporaryAccessPassAuthenticationMethodId, actual.TemporaryAccessPassAuthenticationMethodId)
		}

	}
}

func TestParseMeAuthenticationTemporaryAccessPassMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationTemporaryAccessPassMethodId
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
			Input: "/me/authentication/temporaryAccessPassMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/tEmPoRaRyAcCeSsPaSsMeThOdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/temporaryAccessPassMethods/temporaryAccessPassAuthenticationMethodIdValue",
			Expected: &MeAuthenticationTemporaryAccessPassMethodId{
				TemporaryAccessPassAuthenticationMethodId: "temporaryAccessPassAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/temporaryAccessPassMethods/temporaryAccessPassAuthenticationMethodIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/tEmPoRaRyAcCeSsPaSsMeThOdS/tEmPoRaRyAcCeSsPaSsAuThEnTiCaTiOnMeThOdIdVaLuE",
			Expected: &MeAuthenticationTemporaryAccessPassMethodId{
				TemporaryAccessPassAuthenticationMethodId: "tEmPoRaRyAcCeSsPaSsAuThEnTiCaTiOnMeThOdIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/tEmPoRaRyAcCeSsPaSsMeThOdS/tEmPoRaRyAcCeSsPaSsAuThEnTiCaTiOnMeThOdIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationTemporaryAccessPassMethodIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TemporaryAccessPassAuthenticationMethodId != v.Expected.TemporaryAccessPassAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for TemporaryAccessPassAuthenticationMethodId", v.Expected.TemporaryAccessPassAuthenticationMethodId, actual.TemporaryAccessPassAuthenticationMethodId)
		}

	}
}

func TestSegmentsForMeAuthenticationTemporaryAccessPassMethodId(t *testing.T) {
	segments := MeAuthenticationTemporaryAccessPassMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAuthenticationTemporaryAccessPassMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
