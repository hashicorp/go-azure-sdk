package profilepatent

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfilePatentId{}

func TestNewMeProfilePatentID(t *testing.T) {
	id := NewMeProfilePatentID("itemPatentIdValue")

	if id.ItemPatentId != "itemPatentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemPatentId'", id.ItemPatentId, "itemPatentIdValue")
	}
}

func TestFormatMeProfilePatentID(t *testing.T) {
	actual := NewMeProfilePatentID("itemPatentIdValue").ID()
	expected := "/me/profile/patents/itemPatentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfilePatentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfilePatentId
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
			Input: "/me/profile/patents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/patents/itemPatentIdValue",
			Expected: &MeProfilePatentId{
				ItemPatentId: "itemPatentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/patents/itemPatentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfilePatentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ItemPatentId != v.Expected.ItemPatentId {
			t.Fatalf("Expected %q but got %q for ItemPatentId", v.Expected.ItemPatentId, actual.ItemPatentId)
		}

	}
}

func TestParseMeProfilePatentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfilePatentId
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
			Input: "/me/profile/patents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/pAtEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/patents/itemPatentIdValue",
			Expected: &MeProfilePatentId{
				ItemPatentId: "itemPatentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/patents/itemPatentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/pAtEnTs/iTeMpAtEnTiDvAlUe",
			Expected: &MeProfilePatentId{
				ItemPatentId: "iTeMpAtEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/pAtEnTs/iTeMpAtEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfilePatentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ItemPatentId != v.Expected.ItemPatentId {
			t.Fatalf("Expected %q but got %q for ItemPatentId", v.Expected.ItemPatentId, actual.ItemPatentId)
		}

	}
}

func TestSegmentsForMeProfilePatentId(t *testing.T) {
	segments := MeProfilePatentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfilePatentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
