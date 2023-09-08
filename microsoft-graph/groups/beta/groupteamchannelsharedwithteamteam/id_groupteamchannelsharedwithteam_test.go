package groupteamchannelsharedwithteamteam

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamChannelSharedWithTeamId{}

func TestNewGroupTeamChannelSharedWithTeamID(t *testing.T) {
	id := NewGroupTeamChannelSharedWithTeamID("groupIdValue", "channelIdValue", "sharedWithChannelTeamInfoIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.ChannelId != "channelIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChannelId'", id.ChannelId, "channelIdValue")
	}

	if id.SharedWithChannelTeamInfoId != "sharedWithChannelTeamInfoIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SharedWithChannelTeamInfoId'", id.SharedWithChannelTeamInfoId, "sharedWithChannelTeamInfoIdValue")
	}
}

func TestFormatGroupTeamChannelSharedWithTeamID(t *testing.T) {
	actual := NewGroupTeamChannelSharedWithTeamID("groupIdValue", "channelIdValue", "sharedWithChannelTeamInfoIdValue").ID()
	expected := "/groups/groupIdValue/team/channels/channelIdValue/sharedWithTeams/sharedWithChannelTeamInfoIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupTeamChannelSharedWithTeamID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamChannelSharedWithTeamId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/channels",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/channels/channelIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/channels/channelIdValue/sharedWithTeams",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/channels/channelIdValue/sharedWithTeams/sharedWithChannelTeamInfoIdValue",
			Expected: &GroupTeamChannelSharedWithTeamId{
				GroupId:                     "groupIdValue",
				ChannelId:                   "channelIdValue",
				SharedWithChannelTeamInfoId: "sharedWithChannelTeamInfoIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/channels/channelIdValue/sharedWithTeams/sharedWithChannelTeamInfoIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamChannelSharedWithTeamID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.ChannelId != v.Expected.ChannelId {
			t.Fatalf("Expected %q but got %q for ChannelId", v.Expected.ChannelId, actual.ChannelId)
		}

		if actual.SharedWithChannelTeamInfoId != v.Expected.SharedWithChannelTeamInfoId {
			t.Fatalf("Expected %q but got %q for SharedWithChannelTeamInfoId", v.Expected.SharedWithChannelTeamInfoId, actual.SharedWithChannelTeamInfoId)
		}

	}
}

func TestParseGroupTeamChannelSharedWithTeamIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamChannelSharedWithTeamId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/channels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/cHaNnElS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/channels/channelIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/cHaNnElS/cHaNnElIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/channels/channelIdValue/sharedWithTeams",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/cHaNnElS/cHaNnElIdVaLuE/sHaReDwItHtEaMs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/channels/channelIdValue/sharedWithTeams/sharedWithChannelTeamInfoIdValue",
			Expected: &GroupTeamChannelSharedWithTeamId{
				GroupId:                     "groupIdValue",
				ChannelId:                   "channelIdValue",
				SharedWithChannelTeamInfoId: "sharedWithChannelTeamInfoIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/channels/channelIdValue/sharedWithTeams/sharedWithChannelTeamInfoIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/cHaNnElS/cHaNnElIdVaLuE/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoIdVaLuE",
			Expected: &GroupTeamChannelSharedWithTeamId{
				GroupId:                     "gRoUpIdVaLuE",
				ChannelId:                   "cHaNnElIdVaLuE",
				SharedWithChannelTeamInfoId: "sHaReDwItHcHaNnElTeAmInFoIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/cHaNnElS/cHaNnElIdVaLuE/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamChannelSharedWithTeamIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.ChannelId != v.Expected.ChannelId {
			t.Fatalf("Expected %q but got %q for ChannelId", v.Expected.ChannelId, actual.ChannelId)
		}

		if actual.SharedWithChannelTeamInfoId != v.Expected.SharedWithChannelTeamInfoId {
			t.Fatalf("Expected %q but got %q for SharedWithChannelTeamInfoId", v.Expected.SharedWithChannelTeamInfoId, actual.SharedWithChannelTeamInfoId)
		}

	}
}

func TestSegmentsForGroupTeamChannelSharedWithTeamId(t *testing.T) {
	segments := GroupTeamChannelSharedWithTeamId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupTeamChannelSharedWithTeamId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
