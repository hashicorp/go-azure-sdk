package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileWebsiteId{}

func TestNewMeProfileWebsiteID(t *testing.T) {
	id := NewMeProfileWebsiteID("personWebsiteId")

	if id.PersonWebsiteId != "personWebsiteId" {
		t.Fatalf("Expected %q but got %q for Segment 'PersonWebsiteId'", id.PersonWebsiteId, "personWebsiteId")
	}
}

func TestFormatMeProfileWebsiteID(t *testing.T) {
	actual := NewMeProfileWebsiteID("personWebsiteId").ID()
	expected := "/me/profile/websites/personWebsiteId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfileWebsiteID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileWebsiteId
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
			Input: "/me/profile/websites",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/websites/personWebsiteId",
			Expected: &MeProfileWebsiteId{
				PersonWebsiteId: "personWebsiteId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/websites/personWebsiteId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileWebsiteID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PersonWebsiteId != v.Expected.PersonWebsiteId {
			t.Fatalf("Expected %q but got %q for PersonWebsiteId", v.Expected.PersonWebsiteId, actual.PersonWebsiteId)
		}

	}
}

func TestParseMeProfileWebsiteIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileWebsiteId
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
			Input: "/me/profile/websites",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/wEbSiTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/websites/personWebsiteId",
			Expected: &MeProfileWebsiteId{
				PersonWebsiteId: "personWebsiteId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/websites/personWebsiteId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/wEbSiTeS/pErSoNwEbSiTeId",
			Expected: &MeProfileWebsiteId{
				PersonWebsiteId: "pErSoNwEbSiTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/wEbSiTeS/pErSoNwEbSiTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileWebsiteIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PersonWebsiteId != v.Expected.PersonWebsiteId {
			t.Fatalf("Expected %q but got %q for PersonWebsiteId", v.Expected.PersonWebsiteId, actual.PersonWebsiteId)
		}

	}
}

func TestSegmentsForMeProfileWebsiteId(t *testing.T) {
	segments := MeProfileWebsiteId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfileWebsiteId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
