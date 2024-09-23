package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdIncomingChannelId{}

func TestNewMeJoinedTeamIdIncomingChannelID(t *testing.T) {
	id := NewMeJoinedTeamIdIncomingChannelID("teamId", "channelId")

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.ChannelId != "channelId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChannelId'", id.ChannelId, "channelId")
	}
}

func TestFormatMeJoinedTeamIdIncomingChannelID(t *testing.T) {
	actual := NewMeJoinedTeamIdIncomingChannelID("teamId", "channelId").ID()
	expected := "/me/joinedTeams/teamId/incomingChannels/channelId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamIdIncomingChannelID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdIncomingChannelId
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
			Input: "/me/joinedTeams/teamId/incomingChannels",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/incomingChannels/channelId",
			Expected: &MeJoinedTeamIdIncomingChannelId{
				TeamId:    "teamId",
				ChannelId: "channelId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/incomingChannels/channelId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdIncomingChannelID(v.Input)
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

	}
}

func TestParseMeJoinedTeamIdIncomingChannelIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdIncomingChannelId
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
			Input: "/me/joinedTeams/teamId/incomingChannels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/iNcOmInGcHaNnElS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/incomingChannels/channelId",
			Expected: &MeJoinedTeamIdIncomingChannelId{
				TeamId:    "teamId",
				ChannelId: "channelId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/incomingChannels/channelId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/iNcOmInGcHaNnElS/cHaNnElId",
			Expected: &MeJoinedTeamIdIncomingChannelId{
				TeamId:    "tEaMiD",
				ChannelId: "cHaNnElId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/iNcOmInGcHaNnElS/cHaNnElId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdIncomingChannelIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeJoinedTeamIdIncomingChannelId(t *testing.T) {
	segments := MeJoinedTeamIdIncomingChannelId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamIdIncomingChannelId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
