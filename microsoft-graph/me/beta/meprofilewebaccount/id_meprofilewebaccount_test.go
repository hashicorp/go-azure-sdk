package meprofilewebaccount

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeProfileWebAccountId{}

func TestNewMeProfileWebAccountID(t *testing.T) {
	id := NewMeProfileWebAccountID("webAccountIdValue")

	if id.WebAccountId != "webAccountIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WebAccountId'", id.WebAccountId, "webAccountIdValue")
	}
}

func TestFormatMeProfileWebAccountID(t *testing.T) {
	actual := NewMeProfileWebAccountID("webAccountIdValue").ID()
	expected := "/me/profile/webAccounts/webAccountIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfileWebAccountID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileWebAccountId
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
			Input: "/me/profile",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile/webAccounts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/webAccounts/webAccountIdValue",
			Expected: &MeProfileWebAccountId{
				WebAccountId: "webAccountIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/webAccounts/webAccountIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileWebAccountID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WebAccountId != v.Expected.WebAccountId {
			t.Fatalf("Expected %q but got %q for WebAccountId", v.Expected.WebAccountId, actual.WebAccountId)
		}

	}
}

func TestParseMeProfileWebAccountIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileWebAccountId
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
			Input: "/me/profile",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile/webAccounts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/wEbAcCoUnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/webAccounts/webAccountIdValue",
			Expected: &MeProfileWebAccountId{
				WebAccountId: "webAccountIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/webAccounts/webAccountIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/wEbAcCoUnTs/wEbAcCoUnTiDvAlUe",
			Expected: &MeProfileWebAccountId{
				WebAccountId: "wEbAcCoUnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/wEbAcCoUnTs/wEbAcCoUnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileWebAccountIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WebAccountId != v.Expected.WebAccountId {
			t.Fatalf("Expected %q but got %q for WebAccountId", v.Expected.WebAccountId, actual.WebAccountId)
		}

	}
}

func TestSegmentsForMeProfileWebAccountId(t *testing.T) {
	segments := MeProfileWebAccountId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfileWebAccountId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
