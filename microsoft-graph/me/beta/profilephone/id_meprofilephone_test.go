package profilephone

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfilePhoneId{}

func TestNewMeProfilePhoneID(t *testing.T) {
	id := NewMeProfilePhoneID("itemPhoneIdValue")

	if id.ItemPhoneId != "itemPhoneIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemPhoneId'", id.ItemPhoneId, "itemPhoneIdValue")
	}
}

func TestFormatMeProfilePhoneID(t *testing.T) {
	actual := NewMeProfilePhoneID("itemPhoneIdValue").ID()
	expected := "/me/profile/phones/itemPhoneIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfilePhoneID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfilePhoneId
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
			Input: "/me/profile/phones",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/phones/itemPhoneIdValue",
			Expected: &MeProfilePhoneId{
				ItemPhoneId: "itemPhoneIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/phones/itemPhoneIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfilePhoneID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ItemPhoneId != v.Expected.ItemPhoneId {
			t.Fatalf("Expected %q but got %q for ItemPhoneId", v.Expected.ItemPhoneId, actual.ItemPhoneId)
		}

	}
}

func TestParseMeProfilePhoneIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfilePhoneId
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
			Input: "/me/profile/phones",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/pHoNeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/phones/itemPhoneIdValue",
			Expected: &MeProfilePhoneId{
				ItemPhoneId: "itemPhoneIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/phones/itemPhoneIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/pHoNeS/iTeMpHoNeIdVaLuE",
			Expected: &MeProfilePhoneId{
				ItemPhoneId: "iTeMpHoNeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/pHoNeS/iTeMpHoNeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfilePhoneIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ItemPhoneId != v.Expected.ItemPhoneId {
			t.Fatalf("Expected %q but got %q for ItemPhoneId", v.Expected.ItemPhoneId, actual.ItemPhoneId)
		}

	}
}

func TestSegmentsForMeProfilePhoneId(t *testing.T) {
	segments := MeProfilePhoneId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfilePhoneId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
