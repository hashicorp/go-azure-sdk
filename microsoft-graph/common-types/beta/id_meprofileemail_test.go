package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileEmailId{}

func TestNewMeProfileEmailID(t *testing.T) {
	id := NewMeProfileEmailID("itemEmailIdValue")

	if id.ItemEmailId != "itemEmailIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemEmailId'", id.ItemEmailId, "itemEmailIdValue")
	}
}

func TestFormatMeProfileEmailID(t *testing.T) {
	actual := NewMeProfileEmailID("itemEmailIdValue").ID()
	expected := "/me/profile/emails/itemEmailIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfileEmailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileEmailId
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
			Input: "/me/profile/emails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/emails/itemEmailIdValue",
			Expected: &MeProfileEmailId{
				ItemEmailId: "itemEmailIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/emails/itemEmailIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileEmailID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ItemEmailId != v.Expected.ItemEmailId {
			t.Fatalf("Expected %q but got %q for ItemEmailId", v.Expected.ItemEmailId, actual.ItemEmailId)
		}

	}
}

func TestParseMeProfileEmailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileEmailId
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
			Input: "/me/profile/emails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/eMaIlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/emails/itemEmailIdValue",
			Expected: &MeProfileEmailId{
				ItemEmailId: "itemEmailIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/emails/itemEmailIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/eMaIlS/iTeMeMaIlIdVaLuE",
			Expected: &MeProfileEmailId{
				ItemEmailId: "iTeMeMaIlIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/eMaIlS/iTeMeMaIlIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileEmailIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ItemEmailId != v.Expected.ItemEmailId {
			t.Fatalf("Expected %q but got %q for ItemEmailId", v.Expected.ItemEmailId, actual.ItemEmailId)
		}

	}
}

func TestSegmentsForMeProfileEmailId(t *testing.T) {
	segments := MeProfileEmailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfileEmailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
