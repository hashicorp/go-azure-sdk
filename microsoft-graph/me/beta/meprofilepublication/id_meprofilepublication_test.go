package meprofilepublication

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeProfilePublicationId{}

func TestNewMeProfilePublicationID(t *testing.T) {
	id := NewMeProfilePublicationID("itemPublicationIdValue")

	if id.ItemPublicationId != "itemPublicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemPublicationId'", id.ItemPublicationId, "itemPublicationIdValue")
	}
}

func TestFormatMeProfilePublicationID(t *testing.T) {
	actual := NewMeProfilePublicationID("itemPublicationIdValue").ID()
	expected := "/me/profile/publications/itemPublicationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfilePublicationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfilePublicationId
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
			Input: "/me/profile/publications",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/publications/itemPublicationIdValue",
			Expected: &MeProfilePublicationId{
				ItemPublicationId: "itemPublicationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/publications/itemPublicationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfilePublicationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ItemPublicationId != v.Expected.ItemPublicationId {
			t.Fatalf("Expected %q but got %q for ItemPublicationId", v.Expected.ItemPublicationId, actual.ItemPublicationId)
		}

	}
}

func TestParseMeProfilePublicationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfilePublicationId
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
			Input: "/me/profile/publications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/pUbLiCaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/publications/itemPublicationIdValue",
			Expected: &MeProfilePublicationId{
				ItemPublicationId: "itemPublicationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/publications/itemPublicationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/pUbLiCaTiOnS/iTeMpUbLiCaTiOnIdVaLuE",
			Expected: &MeProfilePublicationId{
				ItemPublicationId: "iTeMpUbLiCaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/pUbLiCaTiOnS/iTeMpUbLiCaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfilePublicationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ItemPublicationId != v.Expected.ItemPublicationId {
			t.Fatalf("Expected %q but got %q for ItemPublicationId", v.Expected.ItemPublicationId, actual.ItemPublicationId)
		}

	}
}

func TestSegmentsForMeProfilePublicationId(t *testing.T) {
	segments := MeProfilePublicationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfilePublicationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
