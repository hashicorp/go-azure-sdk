package userjoinedteamprimarychannelsharedwithteamallowedmember

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{}

func TestNewUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID(t *testing.T) {
	id := NewUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID("userIdValue", "teamIdValue", "sharedWithChannelTeamInfoIdValue", "conversationMemberIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.SharedWithChannelTeamInfoId != "sharedWithChannelTeamInfoIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SharedWithChannelTeamInfoId'", id.SharedWithChannelTeamInfoId, "sharedWithChannelTeamInfoIdValue")
	}

	if id.ConversationMemberId != "conversationMemberIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConversationMemberId'", id.ConversationMemberId, "conversationMemberIdValue")
	}
}

func TestFormatUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID(t *testing.T) {
	actual := NewUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID("userIdValue", "teamIdValue", "sharedWithChannelTeamInfoIdValue", "conversationMemberIdValue").ID()
	expected := "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue/allowedMembers/conversationMemberIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId
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
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/sharedWithTeams",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue/allowedMembers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue/allowedMembers/conversationMemberIdValue",
			Expected: &UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{
				UserId:                      "userIdValue",
				TeamId:                      "teamIdValue",
				SharedWithChannelTeamInfoId: "sharedWithChannelTeamInfoIdValue",
				ConversationMemberId:        "conversationMemberIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue/allowedMembers/conversationMemberIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID(v.Input)
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

		if actual.SharedWithChannelTeamInfoId != v.Expected.SharedWithChannelTeamInfoId {
			t.Fatalf("Expected %q but got %q for SharedWithChannelTeamInfoId", v.Expected.SharedWithChannelTeamInfoId, actual.SharedWithChannelTeamInfoId)
		}

		if actual.ConversationMemberId != v.Expected.ConversationMemberId {
			t.Fatalf("Expected %q but got %q for ConversationMemberId", v.Expected.ConversationMemberId, actual.ConversationMemberId)
		}

	}
}

func TestParseUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId
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
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/sharedWithTeams",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/sHaReDwItHtEaMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue/allowedMembers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoIdVaLuE/aLlOwEdMeMbErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue/allowedMembers/conversationMemberIdValue",
			Expected: &UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{
				UserId:                      "userIdValue",
				TeamId:                      "teamIdValue",
				SharedWithChannelTeamInfoId: "sharedWithChannelTeamInfoIdValue",
				ConversationMemberId:        "conversationMemberIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/joinedTeams/teamIdValue/primaryChannel/sharedWithTeams/sharedWithChannelTeamInfoIdValue/allowedMembers/conversationMemberIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoIdVaLuE/aLlOwEdMeMbErS/cOnVeRsAtIoNmEmBeRiDvAlUe",
			Expected: &UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{
				UserId:                      "uSeRiDvAlUe",
				TeamId:                      "tEaMiDvAlUe",
				SharedWithChannelTeamInfoId: "sHaReDwItHcHaNnElTeAmInFoIdVaLuE",
				ConversationMemberId:        "cOnVeRsAtIoNmEmBeRiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/pRiMaRyChAnNeL/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoIdVaLuE/aLlOwEdMeMbErS/cOnVeRsAtIoNmEmBeRiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberIDInsensitively(v.Input)
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

		if actual.SharedWithChannelTeamInfoId != v.Expected.SharedWithChannelTeamInfoId {
			t.Fatalf("Expected %q but got %q for SharedWithChannelTeamInfoId", v.Expected.SharedWithChannelTeamInfoId, actual.SharedWithChannelTeamInfoId)
		}

		if actual.ConversationMemberId != v.Expected.ConversationMemberId {
			t.Fatalf("Expected %q but got %q for ConversationMemberId", v.Expected.ConversationMemberId, actual.ConversationMemberId)
		}

	}
}

func TestSegmentsForUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId(t *testing.T) {
	segments := UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
