package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTeamworkAssociatedTeamId{}

func TestNewUserIdTeamworkAssociatedTeamID(t *testing.T) {
	id := NewUserIdTeamworkAssociatedTeamID("userId", "associatedTeamInfoId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.AssociatedTeamInfoId != "associatedTeamInfoId" {
		t.Fatalf("Expected %q but got %q for Segment 'AssociatedTeamInfoId'", id.AssociatedTeamInfoId, "associatedTeamInfoId")
	}
}

func TestFormatUserIdTeamworkAssociatedTeamID(t *testing.T) {
	actual := NewUserIdTeamworkAssociatedTeamID("userId", "associatedTeamInfoId").ID()
	expected := "/users/userId/teamwork/associatedTeams/associatedTeamInfoId"
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/teamwork",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/teamwork/associatedTeams",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/teamwork/associatedTeams/associatedTeamInfoId",
			Expected: &UserIdTeamworkAssociatedTeamId{
				UserId:               "userId",
				AssociatedTeamInfoId: "associatedTeamInfoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/teamwork/associatedTeams/associatedTeamInfoId/extra",
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/teamwork",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/tEaMwOrK",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/teamwork/associatedTeams",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/tEaMwOrK/aSsOcIaTeDtEaMs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/teamwork/associatedTeams/associatedTeamInfoId",
			Expected: &UserIdTeamworkAssociatedTeamId{
				UserId:               "userId",
				AssociatedTeamInfoId: "associatedTeamInfoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/teamwork/associatedTeams/associatedTeamInfoId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/tEaMwOrK/aSsOcIaTeDtEaMs/aSsOcIaTeDtEaMiNfOiD",
			Expected: &UserIdTeamworkAssociatedTeamId{
				UserId:               "uSeRiD",
				AssociatedTeamInfoId: "aSsOcIaTeDtEaMiNfOiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/tEaMwOrK/aSsOcIaTeDtEaMs/aSsOcIaTeDtEaMiNfOiD/extra",
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
