package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamChannelIdSharedWithTeamId{}

func TestNewGroupIdTeamChannelIdSharedWithTeamID(t *testing.T) {
	id := NewGroupIdTeamChannelIdSharedWithTeamID("groupId", "channelId", "sharedWithChannelTeamInfoId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.ChannelId != "channelId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChannelId'", id.ChannelId, "channelId")
	}

	if id.SharedWithChannelTeamInfoId != "sharedWithChannelTeamInfoId" {
		t.Fatalf("Expected %q but got %q for Segment 'SharedWithChannelTeamInfoId'", id.SharedWithChannelTeamInfoId, "sharedWithChannelTeamInfoId")
	}
}

func TestFormatGroupIdTeamChannelIdSharedWithTeamID(t *testing.T) {
	actual := NewGroupIdTeamChannelIdSharedWithTeamID("groupId", "channelId", "sharedWithChannelTeamInfoId").ID()
	expected := "/groups/groupId/team/channels/channelId/sharedWithTeams/sharedWithChannelTeamInfoId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamChannelIdSharedWithTeamID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamChannelIdSharedWithTeamId
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
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/sharedWithTeams",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/channels/channelId/sharedWithTeams/sharedWithChannelTeamInfoId",
			Expected: &GroupIdTeamChannelIdSharedWithTeamId{
				GroupId:                     "groupId",
				ChannelId:                   "channelId",
				SharedWithChannelTeamInfoId: "sharedWithChannelTeamInfoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/channels/channelId/sharedWithTeams/sharedWithChannelTeamInfoId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamChannelIdSharedWithTeamID(v.Input)
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

func TestParseGroupIdTeamChannelIdSharedWithTeamIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamChannelIdSharedWithTeamId
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
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/sharedWithTeams",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/sHaReDwItHtEaMs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/channels/channelId/sharedWithTeams/sharedWithChannelTeamInfoId",
			Expected: &GroupIdTeamChannelIdSharedWithTeamId{
				GroupId:                     "groupId",
				ChannelId:                   "channelId",
				SharedWithChannelTeamInfoId: "sharedWithChannelTeamInfoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/channels/channelId/sharedWithTeams/sharedWithChannelTeamInfoId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoId",
			Expected: &GroupIdTeamChannelIdSharedWithTeamId{
				GroupId:                     "gRoUpId",
				ChannelId:                   "cHaNnElId",
				SharedWithChannelTeamInfoId: "sHaReDwItHcHaNnElTeAmInFoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/sHaReDwItHtEaMs/sHaReDwItHcHaNnElTeAmInFoId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamChannelIdSharedWithTeamIDInsensitively(v.Input)
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

func TestSegmentsForGroupIdTeamChannelIdSharedWithTeamId(t *testing.T) {
	segments := GroupIdTeamChannelIdSharedWithTeamId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamChannelIdSharedWithTeamId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
