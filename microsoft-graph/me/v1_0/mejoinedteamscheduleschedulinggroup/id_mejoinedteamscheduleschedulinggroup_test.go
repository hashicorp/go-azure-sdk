package mejoinedteamscheduleschedulinggroup

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamScheduleSchedulingGroupId{}

func TestNewMeJoinedTeamScheduleSchedulingGroupID(t *testing.T) {
	id := NewMeJoinedTeamScheduleSchedulingGroupID("teamIdValue", "schedulingGroupIdValue")

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.SchedulingGroupId != "schedulingGroupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SchedulingGroupId'", id.SchedulingGroupId, "schedulingGroupIdValue")
	}
}

func TestFormatMeJoinedTeamScheduleSchedulingGroupID(t *testing.T) {
	actual := NewMeJoinedTeamScheduleSchedulingGroupID("teamIdValue", "schedulingGroupIdValue").ID()
	expected := "/me/joinedTeams/teamIdValue/schedule/schedulingGroups/schedulingGroupIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamScheduleSchedulingGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamScheduleSchedulingGroupId
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
			Input: "/me/joinedTeams/teamIdValue/schedule",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/schedule/schedulingGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/schedule/schedulingGroups/schedulingGroupIdValue",
			Expected: &MeJoinedTeamScheduleSchedulingGroupId{
				TeamId:            "teamIdValue",
				SchedulingGroupId: "schedulingGroupIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/schedule/schedulingGroups/schedulingGroupIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamScheduleSchedulingGroupID(v.Input)
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

func TestParseMeJoinedTeamScheduleSchedulingGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamScheduleSchedulingGroupId
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
			Input: "/me/joinedTeams/teamIdValue/schedule",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/joinedTeams/teamIdValue/schedule/schedulingGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/sChEdUlInGgRoUpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/schedule/schedulingGroups/schedulingGroupIdValue",
			Expected: &MeJoinedTeamScheduleSchedulingGroupId{
				TeamId:            "teamIdValue",
				SchedulingGroupId: "schedulingGroupIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/schedule/schedulingGroups/schedulingGroupIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/sChEdUlInGgRoUpS/sChEdUlInGgRoUpIdVaLuE",
			Expected: &MeJoinedTeamScheduleSchedulingGroupId{
				TeamId:            "tEaMiDvAlUe",
				SchedulingGroupId: "sChEdUlInGgRoUpIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/sChEdUlInGgRoUpS/sChEdUlInGgRoUpIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamScheduleSchedulingGroupIDInsensitively(v.Input)
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

func TestSegmentsForMeJoinedTeamScheduleSchedulingGroupId(t *testing.T) {
	segments := MeJoinedTeamScheduleSchedulingGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamScheduleSchedulingGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
