package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdScheduleSchedulingGroupId{}

func TestNewMeJoinedTeamIdScheduleSchedulingGroupID(t *testing.T) {
	id := NewMeJoinedTeamIdScheduleSchedulingGroupID("teamId", "schedulingGroupId")

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.SchedulingGroupId != "schedulingGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'SchedulingGroupId'", id.SchedulingGroupId, "schedulingGroupId")
	}
}

func TestFormatMeJoinedTeamIdScheduleSchedulingGroupID(t *testing.T) {
	actual := NewMeJoinedTeamIdScheduleSchedulingGroupID("teamId", "schedulingGroupId").ID()
	expected := "/me/joinedTeams/teamId/schedule/schedulingGroups/schedulingGroupId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamIdScheduleSchedulingGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdScheduleSchedulingGroupId
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
			Input: "/me/joinedTeams/teamId/schedule",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/schedule/schedulingGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/schedule/schedulingGroups/schedulingGroupId",
			Expected: &MeJoinedTeamIdScheduleSchedulingGroupId{
				TeamId:            "teamId",
				SchedulingGroupId: "schedulingGroupId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/schedule/schedulingGroups/schedulingGroupId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdScheduleSchedulingGroupID(v.Input)
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

		if actual.SchedulingGroupId != v.Expected.SchedulingGroupId {
			t.Fatalf("Expected %q but got %q for SchedulingGroupId", v.Expected.SchedulingGroupId, actual.SchedulingGroupId)
		}

	}
}

func TestParseMeJoinedTeamIdScheduleSchedulingGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdScheduleSchedulingGroupId
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
			Input: "/me/joinedTeams/teamId/schedule",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamId/schedule/schedulingGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE/sChEdUlInGgRoUpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/schedule/schedulingGroups/schedulingGroupId",
			Expected: &MeJoinedTeamIdScheduleSchedulingGroupId{
				TeamId:            "teamId",
				SchedulingGroupId: "schedulingGroupId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/schedule/schedulingGroups/schedulingGroupId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE/sChEdUlInGgRoUpS/sChEdUlInGgRoUpId",
			Expected: &MeJoinedTeamIdScheduleSchedulingGroupId{
				TeamId:            "tEaMiD",
				SchedulingGroupId: "sChEdUlInGgRoUpId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE/sChEdUlInGgRoUpS/sChEdUlInGgRoUpId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdScheduleSchedulingGroupIDInsensitively(v.Input)
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

		if actual.SchedulingGroupId != v.Expected.SchedulingGroupId {
			t.Fatalf("Expected %q but got %q for SchedulingGroupId", v.Expected.SchedulingGroupId, actual.SchedulingGroupId)
		}

	}
}

func TestSegmentsForMeJoinedTeamIdScheduleSchedulingGroupId(t *testing.T) {
	segments := MeJoinedTeamIdScheduleSchedulingGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamIdScheduleSchedulingGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
