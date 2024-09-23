package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdPrimaryChannelSharedWithTeamId{}

func TestNewMeJoinedTeamIdPrimaryChannelSharedWithTeamID(t *testing.T) {
	id := NewMeJoinedTeamIdPrimaryChannelSharedWithTeamID("teamId", "sharedWithChannelTeamInfoId")

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.SharedWithChannelTeamInfoId != "sharedWithChannelTeamInfoId" {
		t.Fatalf("Expected %q but got %q for Segment 'SharedWithChannelTeamInfoId'", id.SharedWithChannelTeamInfoId, "sharedWithChannelTeamInfoId")
	}
}

func TestFormatMeJoinedTeamIdPrimaryChannelSharedWithTeamID(t *testing.T) {
	actual := NewMeJoinedTeamIdPrimaryChannelSharedWithTeamID("teamId", "sharedWithChannelTeamInfoId").ID()
	expected := "/me/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamIdPrimaryChannelSharedWithTeamID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdPrimaryChannelSharedWithTeamId
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
			Input: "/me/joinedTeams/teamId/primaryChannel",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/primaryChannel/sharedWithTeams",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId",
			Expected: &MeJoinedTeamIdPrimaryChannelSharedWithTeamId{
				TeamId:                      "teamId",
				SharedWithChannelTeamInfoId: "sharedWithChannelTeamInfoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdPrimaryChannelSharedWithTeamID(v.Input)
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

		if actual.SharedWithChannelTeamInfoId != v.Expected.SharedWithChannelTeamInfoId {
			t.Fatalf("Expected %q but got %q for SharedWithChannelTeamInfoId", v.Expected.SharedWithChannelTeamInfoId, actual.SharedWithChannelTeamInfoId)
		}

	}
}

func TestParseMeJoinedTeamIdPrimaryChannelSharedWithTeamIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdPrimaryChannelSharedWithTeamId
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
			Input: "/me/joinedTeams/teamId/primaryChannel",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/primaryChannel/sharedWithTeams",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/sHaReDwItHtEaMs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId",
			Expected: &MeJoinedTeamIdPrimaryChannelSharedWithTeamId{
				TeamId:                      "teamId",
				SharedWithChannelTeamInfoId: "sharedWithChannelTeamInfoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoId",
			Expected: &MeJoinedTeamIdPrimaryChannelSharedWithTeamId{
				TeamId:                      "tEaMiD",
				SharedWithChannelTeamInfoId: "sHaReDwItHcHaNnElTeAmInFoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/pRiMaRyChAnNeL/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdPrimaryChannelSharedWithTeamIDInsensitively(v.Input)
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

		if actual.SharedWithChannelTeamInfoId != v.Expected.SharedWithChannelTeamInfoId {
			t.Fatalf("Expected %q but got %q for SharedWithChannelTeamInfoId", v.Expected.SharedWithChannelTeamInfoId, actual.SharedWithChannelTeamInfoId)
		}

	}
}

func TestSegmentsForMeJoinedTeamIdPrimaryChannelSharedWithTeamId(t *testing.T) {
	segments := MeJoinedTeamIdPrimaryChannelSharedWithTeamId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamIdPrimaryChannelSharedWithTeamId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
