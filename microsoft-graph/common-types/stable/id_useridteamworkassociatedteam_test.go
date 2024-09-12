package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTeamworkAssociatedTeamId{}

func TestNewUserIdTeamworkAssociatedTeamID(t *testing.T) {
	id := NewUserIdTeamworkAssociatedTeamID("userIdValue", "associatedTeamInfoIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.AssociatedTeamInfoId != "associatedTeamInfoIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AssociatedTeamInfoId'", id.AssociatedTeamInfoId, "associatedTeamInfoIdValue")
	}
}

func TestFormatUserIdTeamworkAssociatedTeamID(t *testing.T) {
	actual := NewUserIdTeamworkAssociatedTeamID("userIdValue", "associatedTeamInfoIdValue").ID()
	expected := "/users/userIdValue/teamwork/associatedTeams/associatedTeamInfoIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdTeamworkAssociatedTeamID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdTeamworkAssociatedTeamId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/teamwork",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/teamwork/associatedTeams",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/teamwork/associatedTeams/associatedTeamInfoIdValue",
			Expected: &UserIdTeamworkAssociatedTeamId{
				UserId:               "userIdValue",
				AssociatedTeamInfoId: "associatedTeamInfoIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/teamwork/associatedTeams/associatedTeamInfoIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdTeamworkAssociatedTeamID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.AssociatedTeamInfoId != v.Expected.AssociatedTeamInfoId {
			t.Fatalf("Expected %q but got %q for AssociatedTeamInfoId", v.Expected.AssociatedTeamInfoId, actual.AssociatedTeamInfoId)
		}

	}
}

func TestParseUserIdTeamworkAssociatedTeamIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdTeamworkAssociatedTeamId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/teamwork",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tEaMwOrK",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/teamwork/associatedTeams",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tEaMwOrK/aSsOcIaTeDtEaMs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/teamwork/associatedTeams/associatedTeamInfoIdValue",
			Expected: &UserIdTeamworkAssociatedTeamId{
				UserId:               "userIdValue",
				AssociatedTeamInfoId: "associatedTeamInfoIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/teamwork/associatedTeams/associatedTeamInfoIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tEaMwOrK/aSsOcIaTeDtEaMs/aSsOcIaTeDtEaMiNfOiDvAlUe",
			Expected: &UserIdTeamworkAssociatedTeamId{
				UserId:               "uSeRiDvAlUe",
				AssociatedTeamInfoId: "aSsOcIaTeDtEaMiNfOiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tEaMwOrK/aSsOcIaTeDtEaMs/aSsOcIaTeDtEaMiNfOiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdTeamworkAssociatedTeamIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.AssociatedTeamInfoId != v.Expected.AssociatedTeamInfoId {
			t.Fatalf("Expected %q but got %q for AssociatedTeamInfoId", v.Expected.AssociatedTeamInfoId, actual.AssociatedTeamInfoId)
		}

	}
}

func TestSegmentsForUserIdTeamworkAssociatedTeamId(t *testing.T) {
	segments := UserIdTeamworkAssociatedTeamId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdTeamworkAssociatedTeamId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
