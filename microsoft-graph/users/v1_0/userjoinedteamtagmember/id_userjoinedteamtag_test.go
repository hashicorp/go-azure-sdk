package userjoinedteamtagmember

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamTagId{}

func TestNewUserJoinedTeamTagID(t *testing.T) {
	id := NewUserJoinedTeamTagID("userIdValue", "teamIdValue", "teamworkTagIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.TeamworkTagId != "teamworkTagIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamworkTagId'", id.TeamworkTagId, "teamworkTagIdValue")
	}
}

func TestFormatUserJoinedTeamTagID(t *testing.T) {
	actual := NewUserJoinedTeamTagID("userIdValue", "teamIdValue", "teamworkTagIdValue").ID()
	expected := "/users/userIdValue/joinedTeams/teamIdValue/tags/teamworkTagIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserJoinedTeamTagID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserJoinedTeamTagId
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
			Input: "/users/userIdValue/joinedTeams",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/tags",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/tags/teamworkTagIdValue",
			Expected: &UserJoinedTeamTagId{
				UserId:        "userIdValue",
				TeamId:        "teamIdValue",
				TeamworkTagId: "teamworkTagIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/joinedTeams/teamIdValue/tags/teamworkTagIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserJoinedTeamTagID(v.Input)
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

		if actual.TeamId != v.Expected.TeamId {
			t.Fatalf("Expected %q but got %q for TeamId", v.Expected.TeamId, actual.TeamId)
		}

		if actual.TeamworkTagId != v.Expected.TeamworkTagId {
			t.Fatalf("Expected %q but got %q for TeamworkTagId", v.Expected.TeamworkTagId, actual.TeamworkTagId)
		}

	}
}

func TestParseUserJoinedTeamTagIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserJoinedTeamTagId
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
			Input: "/users/userIdValue/joinedTeams",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/tags",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/tAgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/tags/teamworkTagIdValue",
			Expected: &UserJoinedTeamTagId{
				UserId:        "userIdValue",
				TeamId:        "teamIdValue",
				TeamworkTagId: "teamworkTagIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/joinedTeams/teamIdValue/tags/teamworkTagIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/tAgS/tEaMwOrKtAgIdVaLuE",
			Expected: &UserJoinedTeamTagId{
				UserId:        "uSeRiDvAlUe",
				TeamId:        "tEaMiDvAlUe",
				TeamworkTagId: "tEaMwOrKtAgIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/tAgS/tEaMwOrKtAgIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserJoinedTeamTagIDInsensitively(v.Input)
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

		if actual.TeamId != v.Expected.TeamId {
			t.Fatalf("Expected %q but got %q for TeamId", v.Expected.TeamId, actual.TeamId)
		}

		if actual.TeamworkTagId != v.Expected.TeamworkTagId {
			t.Fatalf("Expected %q but got %q for TeamworkTagId", v.Expected.TeamworkTagId, actual.TeamworkTagId)
		}

	}
}

func TestSegmentsForUserJoinedTeamTagId(t *testing.T) {
	segments := UserJoinedTeamTagId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserJoinedTeamTagId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
