package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdScheduleTimeOffRequestId{}

func TestNewMeJoinedTeamIdScheduleTimeOffRequestID(t *testing.T) {
	id := NewMeJoinedTeamIdScheduleTimeOffRequestID("teamId", "timeOffRequestId")

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.TimeOffRequestId != "timeOffRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'TimeOffRequestId'", id.TimeOffRequestId, "timeOffRequestId")
	}
}

func TestFormatMeJoinedTeamIdScheduleTimeOffRequestID(t *testing.T) {
	actual := NewMeJoinedTeamIdScheduleTimeOffRequestID("teamId", "timeOffRequestId").ID()
	expected := "/me/joinedTeams/teamId/schedule/timeOffRequests/timeOffRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamIdScheduleTimeOffRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdScheduleTimeOffRequestId
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
			Input: "/me/joinedTeams/teamId/schedule/timeOffRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/schedule/timeOffRequests/timeOffRequestId",
			Expected: &MeJoinedTeamIdScheduleTimeOffRequestId{
				TeamId:           "teamId",
				TimeOffRequestId: "timeOffRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/schedule/timeOffRequests/timeOffRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdScheduleTimeOffRequestID(v.Input)
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

		if actual.TimeOffRequestId != v.Expected.TimeOffRequestId {
			t.Fatalf("Expected %q but got %q for TimeOffRequestId", v.Expected.TimeOffRequestId, actual.TimeOffRequestId)
		}

	}
}

func TestParseMeJoinedTeamIdScheduleTimeOffRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdScheduleTimeOffRequestId
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
			Input: "/me/joinedTeams/teamId/schedule/timeOffRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE/tImEoFfReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/schedule/timeOffRequests/timeOffRequestId",
			Expected: &MeJoinedTeamIdScheduleTimeOffRequestId{
				TeamId:           "teamId",
				TimeOffRequestId: "timeOffRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/schedule/timeOffRequests/timeOffRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE/tImEoFfReQuEsTs/tImEoFfReQuEsTiD",
			Expected: &MeJoinedTeamIdScheduleTimeOffRequestId{
				TeamId:           "tEaMiD",
				TimeOffRequestId: "tImEoFfReQuEsTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE/tImEoFfReQuEsTs/tImEoFfReQuEsTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdScheduleTimeOffRequestIDInsensitively(v.Input)
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

		if actual.TimeOffRequestId != v.Expected.TimeOffRequestId {
			t.Fatalf("Expected %q but got %q for TimeOffRequestId", v.Expected.TimeOffRequestId, actual.TimeOffRequestId)
		}

	}
}

func TestSegmentsForMeJoinedTeamIdScheduleTimeOffRequestId(t *testing.T) {
	segments := MeJoinedTeamIdScheduleTimeOffRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamIdScheduleTimeOffRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
