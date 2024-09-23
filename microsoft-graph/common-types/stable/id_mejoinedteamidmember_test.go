package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdMemberId{}

func TestNewMeJoinedTeamIdMemberID(t *testing.T) {
	id := NewMeJoinedTeamIdMemberID("teamId", "conversationMemberId")

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.ConversationMemberId != "conversationMemberId" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationMemberId'", id.ConversationMemberId, "conversationMemberId")
	}
}

func TestFormatMeJoinedTeamIdMemberID(t *testing.T) {
	actual := NewMeJoinedTeamIdMemberID("teamId", "conversationMemberId").ID()
	expected := "/me/joinedTeams/teamId/members/conversationMemberId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamIdMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdMemberId
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
			Input: "/me/joinedTeams/teamId/members",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/members/conversationMemberId",
			Expected: &MeJoinedTeamIdMemberId{
				TeamId:               "teamId",
				ConversationMemberId: "conversationMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/members/conversationMemberId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdMemberID(v.Input)
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

		if actual.ConversationMemberId != v.Expected.ConversationMemberId {
			t.Fatalf("Expected %q but got %q for ConversationMemberId", v.Expected.ConversationMemberId, actual.ConversationMemberId)
		}

	}
}

func TestParseMeJoinedTeamIdMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdMemberId
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
			Input: "/me/joinedTeams/teamId/members",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/mEmBeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/members/conversationMemberId",
			Expected: &MeJoinedTeamIdMemberId{
				TeamId:               "teamId",
				ConversationMemberId: "conversationMemberId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/members/conversationMemberId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/mEmBeRs/cOnVeRsAtIoNmEmBeRiD",
			Expected: &MeJoinedTeamIdMemberId{
				TeamId:               "tEaMiD",
				ConversationMemberId: "cOnVeRsAtIoNmEmBeRiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/mEmBeRs/cOnVeRsAtIoNmEmBeRiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdMemberIDInsensitively(v.Input)
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

		if actual.ConversationMemberId != v.Expected.ConversationMemberId {
			t.Fatalf("Expected %q but got %q for ConversationMemberId", v.Expected.ConversationMemberId, actual.ConversationMemberId)
		}

	}
}

func TestSegmentsForMeJoinedTeamIdMemberId(t *testing.T) {
	segments := MeJoinedTeamIdMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamIdMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
