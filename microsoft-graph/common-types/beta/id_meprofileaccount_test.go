package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileAccountId{}

func TestNewMeProfileAccountID(t *testing.T) {
	id := NewMeProfileAccountID("userAccountInformationId")

	if id.UserAccountInformationId != "userAccountInformationId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserAccountInformationId'", id.UserAccountInformationId, "userAccountInformationId")
	}
}

func TestFormatMeProfileAccountID(t *testing.T) {
	actual := NewMeProfileAccountID("userAccountInformationId").ID()
	expected := "/me/profile/account/userAccountInformationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfileAccountID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileAccountId
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
			Input: "/me/profile/account",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/account/userAccountInformationId",
			Expected: &MeProfileAccountId{
				UserAccountInformationId: "userAccountInformationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/account/userAccountInformationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileAccountID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserAccountInformationId != v.Expected.UserAccountInformationId {
			t.Fatalf("Expected %q but got %q for UserAccountInformationId", v.Expected.UserAccountInformationId, actual.UserAccountInformationId)
		}

	}
}

func TestParseMeProfileAccountIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileAccountId
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
			Input: "/me/profile/account",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/aCcOuNt",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/account/userAccountInformationId",
			Expected: &MeProfileAccountId{
				UserAccountInformationId: "userAccountInformationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/account/userAccountInformationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/aCcOuNt/uSeRaCcOuNtInFoRmAtIoNiD",
			Expected: &MeProfileAccountId{
				UserAccountInformationId: "uSeRaCcOuNtInFoRmAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/aCcOuNt/uSeRaCcOuNtInFoRmAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileAccountIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserAccountInformationId != v.Expected.UserAccountInformationId {
			t.Fatalf("Expected %q but got %q for UserAccountInformationId", v.Expected.UserAccountInformationId, actual.UserAccountInformationId)
		}

	}
}

func TestSegmentsForMeProfileAccountId(t *testing.T) {
	segments := MeProfileAccountId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfileAccountId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
