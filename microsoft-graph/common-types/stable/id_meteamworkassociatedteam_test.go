package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeTeamworkAssociatedTeamId{}

func TestNewMeTeamworkAssociatedTeamID(t *testing.T) {
	id := NewMeTeamworkAssociatedTeamID("associatedTeamInfoIdValue")

	if id.AssociatedTeamInfoId != "associatedTeamInfoIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AssociatedTeamInfoId'", id.AssociatedTeamInfoId, "associatedTeamInfoIdValue")
	}
}

func TestFormatMeTeamworkAssociatedTeamID(t *testing.T) {
	actual := NewMeTeamworkAssociatedTeamID("associatedTeamInfoIdValue").ID()
	expected := "/me/teamwork/associatedTeams/associatedTeamInfoIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeTeamworkAssociatedTeamID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeTeamworkAssociatedTeamId
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
			Input: "/me/teamwork",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/teamwork/associatedTeams",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/teamwork/associatedTeams/associatedTeamInfoIdValue",
			Expected: &MeTeamworkAssociatedTeamId{
				AssociatedTeamInfoId: "associatedTeamInfoIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/teamwork/associatedTeams/associatedTeamInfoIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeTeamworkAssociatedTeamID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AssociatedTeamInfoId != v.Expected.AssociatedTeamInfoId {
			t.Fatalf("Expected %q but got %q for AssociatedTeamInfoId", v.Expected.AssociatedTeamInfoId, actual.AssociatedTeamInfoId)
		}

	}
}

func TestParseMeTeamworkAssociatedTeamIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeTeamworkAssociatedTeamId
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
			Input: "/me/teamwork",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tEaMwOrK",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/teamwork/associatedTeams",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tEaMwOrK/aSsOcIaTeDtEaMs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/teamwork/associatedTeams/associatedTeamInfoIdValue",
			Expected: &MeTeamworkAssociatedTeamId{
				AssociatedTeamInfoId: "associatedTeamInfoIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/teamwork/associatedTeams/associatedTeamInfoIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tEaMwOrK/aSsOcIaTeDtEaMs/aSsOcIaTeDtEaMiNfOiDvAlUe",
			Expected: &MeTeamworkAssociatedTeamId{
				AssociatedTeamInfoId: "aSsOcIaTeDtEaMiNfOiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/tEaMwOrK/aSsOcIaTeDtEaMs/aSsOcIaTeDtEaMiNfOiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeTeamworkAssociatedTeamIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AssociatedTeamInfoId != v.Expected.AssociatedTeamInfoId {
			t.Fatalf("Expected %q but got %q for AssociatedTeamInfoId", v.Expected.AssociatedTeamInfoId, actual.AssociatedTeamInfoId)
		}

	}
}

func TestSegmentsForMeTeamworkAssociatedTeamId(t *testing.T) {
	segments := MeTeamworkAssociatedTeamId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeTeamworkAssociatedTeamId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
