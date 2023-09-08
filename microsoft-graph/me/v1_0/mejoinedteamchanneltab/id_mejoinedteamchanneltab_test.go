package mejoinedteamchanneltab

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamChannelTabId{}

func TestNewMeJoinedTeamChannelTabID(t *testing.T) {
	id := NewMeJoinedTeamChannelTabID("teamIdValue", "channelIdValue", "teamsTabIdValue")

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.ChannelId != "channelIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChannelId'", id.ChannelId, "channelIdValue")
	}

	if id.TeamsTabId != "teamsTabIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamsTabId'", id.TeamsTabId, "teamsTabIdValue")
	}
}

func TestFormatMeJoinedTeamChannelTabID(t *testing.T) {
	actual := NewMeJoinedTeamChannelTabID("teamIdValue", "channelIdValue", "teamsTabIdValue").ID()
	expected := "/me/joinedTeams/teamIdValue/channels/channelIdValue/tabs/teamsTabIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamChannelTabID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamChannelTabId
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
			Input: "/me/joinedTeams/teamIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/tabs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/tabs/teamsTabIdValue",
			Expected: &MeJoinedTeamChannelTabId{
				TeamId:     "teamIdValue",
				ChannelId:  "channelIdValue",
				TeamsTabId: "teamsTabIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/tabs/teamsTabIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamChannelTabID(v.Input)
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

		if actual.ChannelId != v.Expected.ChannelId {
			t.Fatalf("Expected %q but got %q for ChannelId", v.Expected.ChannelId, actual.ChannelId)
		}

		if actual.TeamsTabId != v.Expected.TeamsTabId {
			t.Fatalf("Expected %q but got %q for TeamsTabId", v.Expected.TeamsTabId, actual.TeamsTabId)
		}

	}
}

func TestParseMeJoinedTeamChannelTabIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamChannelTabId
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
			Input: "/me/joinedTeams/teamIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/cHaNnElS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/cHaNnElS/cHaNnElIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/tabs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/cHaNnElS/cHaNnElIdVaLuE/tAbS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/tabs/teamsTabIdValue",
			Expected: &MeJoinedTeamChannelTabId{
				TeamId:     "teamIdValue",
				ChannelId:  "channelIdValue",
				TeamsTabId: "teamsTabIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/channels/channelIdValue/tabs/teamsTabIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/cHaNnElS/cHaNnElIdVaLuE/tAbS/tEaMsTaBiDvAlUe",
			Expected: &MeJoinedTeamChannelTabId{
				TeamId:     "tEaMiDvAlUe",
				ChannelId:  "cHaNnElIdVaLuE",
				TeamsTabId: "tEaMsTaBiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/cHaNnElS/cHaNnElIdVaLuE/tAbS/tEaMsTaBiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamChannelTabIDInsensitively(v.Input)
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

		if actual.ChannelId != v.Expected.ChannelId {
			t.Fatalf("Expected %q but got %q for ChannelId", v.Expected.ChannelId, actual.ChannelId)
		}

		if actual.TeamsTabId != v.Expected.TeamsTabId {
			t.Fatalf("Expected %q but got %q for TeamsTabId", v.Expected.TeamsTabId, actual.TeamsTabId)
		}

	}
}

func TestSegmentsForMeJoinedTeamChannelTabId(t *testing.T) {
	segments := MeJoinedTeamChannelTabId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamChannelTabId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
