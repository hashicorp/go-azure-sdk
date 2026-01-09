package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdTagIdMemberId{}

func TestNewMeJoinedTeamIdTagIdMemberID(t *testing.T) {
	id := NewMeJoinedTeamIdTagIdMemberID("teamId", "teamworkTagId", "teamworkTagMemberId")

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.TeamworkTagId != "teamworkTagId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamworkTagId'", id.TeamworkTagId, "teamworkTagId")
	}

	if id.TeamworkTagMemberId != "teamworkTagMemberId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamworkTagMemberId'", id.TeamworkTagMemberId, "teamworkTagMemberId")
	}
}

func TestFormatMeJoinedTeamIdTagIdMemberID(t *testing.T) {
	actual := NewMeJoinedTeamIdTagIdMemberID("teamId", "teamworkTagId", "teamworkTagMemberId").ID()
	expected := "/me/joinedTeams/teamId/tags/teamworkTagId/members/teamworkTagMemberId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamIdTagIdMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdTagIdMemberId
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
			Input: "/me/joinedTeams",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/tags",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/tags/teamworkTagId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/tags/teamworkTagId/members",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/tags/teamworkTagId/members/teamworkTagMemberId",
			Expected: &MeJoinedTeamIdTagIdMemberId{
				TeamId:              "teamId",
				TeamworkTagId:       "teamworkTagId",
				TeamworkTagMemberId: "teamworkTagMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/tags/teamworkTagId/members/teamworkTagMemberId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdTagIdMemberID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TeamId != v.Expected.TeamId {
			t.Fatalf("Expected %q but got %q for TeamId", v.Expected.TeamId, actual.TeamId)
		}

		if actual.TeamworkTagId != v.Expected.TeamworkTagId {
			t.Fatalf("Expected %q but got %q for TeamworkTagId", v.Expected.TeamworkTagId, actual.TeamworkTagId)
		}

		if actual.TeamworkTagMemberId != v.Expected.TeamworkTagMemberId {
			t.Fatalf("Expected %q but got %q for TeamworkTagMemberId", v.Expected.TeamworkTagMemberId, actual.TeamworkTagMemberId)
		}

	}
}

func TestParseMeJoinedTeamIdTagIdMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdTagIdMemberId
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
			Input: "/me/joinedTeams",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/tags",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/tAgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/tags/teamworkTagId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/tAgS/tEaMwOrKtAgId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/tags/teamworkTagId/members",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/tAgS/tEaMwOrKtAgId/mEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/tags/teamworkTagId/members/teamworkTagMemberId",
			Expected: &MeJoinedTeamIdTagIdMemberId{
				TeamId:              "teamId",
				TeamworkTagId:       "teamworkTagId",
				TeamworkTagMemberId: "teamworkTagMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/tags/teamworkTagId/members/teamworkTagMemberId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/tAgS/tEaMwOrKtAgId/mEmBeRs/tEaMwOrKtAgMeMbErId",
			Expected: &MeJoinedTeamIdTagIdMemberId{
				TeamId:              "tEaMiD",
				TeamworkTagId:       "tEaMwOrKtAgId",
				TeamworkTagMemberId: "tEaMwOrKtAgMeMbErId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/tAgS/tEaMwOrKtAgId/mEmBeRs/tEaMwOrKtAgMeMbErId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdTagIdMemberIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TeamId != v.Expected.TeamId {
			t.Fatalf("Expected %q but got %q for TeamId", v.Expected.TeamId, actual.TeamId)
		}

		if actual.TeamworkTagId != v.Expected.TeamworkTagId {
			t.Fatalf("Expected %q but got %q for TeamworkTagId", v.Expected.TeamworkTagId, actual.TeamworkTagId)
		}

		if actual.TeamworkTagMemberId != v.Expected.TeamworkTagMemberId {
			t.Fatalf("Expected %q but got %q for TeamworkTagMemberId", v.Expected.TeamworkTagMemberId, actual.TeamworkTagMemberId)
		}

	}
}

func TestSegmentsForMeJoinedTeamIdTagIdMemberId(t *testing.T) {
	segments := MeJoinedTeamIdTagIdMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamIdTagIdMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
