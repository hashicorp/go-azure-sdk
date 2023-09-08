package meprofileaddress

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeProfileAddressId{}

func TestNewMeProfileAddressID(t *testing.T) {
	id := NewMeProfileAddressID("itemAddressIdValue")

	if id.ItemAddressId != "itemAddressIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemAddressId'", id.ItemAddressId, "itemAddressIdValue")
	}
}

func TestFormatMeProfileAddressID(t *testing.T) {
	actual := NewMeProfileAddressID("itemAddressIdValue").ID()
	expected := "/me/profile/addresses/itemAddressIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfileAddressID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileAddressId
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
			Input: "/me/profile/addresses",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/addresses/itemAddressIdValue",
			Expected: &MeProfileAddressId{
				ItemAddressId: "itemAddressIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/addresses/itemAddressIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileAddressID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ItemAddressId != v.Expected.ItemAddressId {
			t.Fatalf("Expected %q but got %q for ItemAddressId", v.Expected.ItemAddressId, actual.ItemAddressId)
		}

	}
}

func TestParseMeProfileAddressIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileAddressId
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
			Input: "/me/profile/addresses",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/aDdReSsEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/addresses/itemAddressIdValue",
			Expected: &MeProfileAddressId{
				ItemAddressId: "itemAddressIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/addresses/itemAddressIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/aDdReSsEs/iTeMaDdReSsIdVaLuE",
			Expected: &MeProfileAddressId{
				ItemAddressId: "iTeMaDdReSsIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/aDdReSsEs/iTeMaDdReSsIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileAddressIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ItemAddressId != v.Expected.ItemAddressId {
			t.Fatalf("Expected %q but got %q for ItemAddressId", v.Expected.ItemAddressId, actual.ItemAddressId)
		}

	}
}

func TestSegmentsForMeProfileAddressId(t *testing.T) {
	segments := MeProfileAddressId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfileAddressId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
