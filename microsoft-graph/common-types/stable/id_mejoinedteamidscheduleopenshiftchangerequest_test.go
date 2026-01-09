package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdScheduleOpenShiftChangeRequestId{}

func TestNewMeJoinedTeamIdScheduleOpenShiftChangeRequestID(t *testing.T) {
	id := NewMeJoinedTeamIdScheduleOpenShiftChangeRequestID("teamId", "openShiftChangeRequestId")

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.OpenShiftChangeRequestId != "openShiftChangeRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'OpenShiftChangeRequestId'", id.OpenShiftChangeRequestId, "openShiftChangeRequestId")
	}
}

func TestFormatMeJoinedTeamIdScheduleOpenShiftChangeRequestID(t *testing.T) {
	actual := NewMeJoinedTeamIdScheduleOpenShiftChangeRequestID("teamId", "openShiftChangeRequestId").ID()
	expected := "/me/joinedTeams/teamId/schedule/openShiftChangeRequests/openShiftChangeRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamIdScheduleOpenShiftChangeRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdScheduleOpenShiftChangeRequestId
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
			Input: "/me/joinedTeams/teamId/schedule/openShiftChangeRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/schedule/openShiftChangeRequests/openShiftChangeRequestId",
			Expected: &MeJoinedTeamIdScheduleOpenShiftChangeRequestId{
				TeamId:                   "teamId",
				OpenShiftChangeRequestId: "openShiftChangeRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/schedule/openShiftChangeRequests/openShiftChangeRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdScheduleOpenShiftChangeRequestID(v.Input)
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

		if actual.OpenShiftChangeRequestId != v.Expected.OpenShiftChangeRequestId {
			t.Fatalf("Expected %q but got %q for OpenShiftChangeRequestId", v.Expected.OpenShiftChangeRequestId, actual.OpenShiftChangeRequestId)
		}

	}
}

func TestParseMeJoinedTeamIdScheduleOpenShiftChangeRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdScheduleOpenShiftChangeRequestId
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
			Input: "/me/joinedTeams/teamId/schedule/openShiftChangeRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE/oPeNsHiFtChAnGeReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/schedule/openShiftChangeRequests/openShiftChangeRequestId",
			Expected: &MeJoinedTeamIdScheduleOpenShiftChangeRequestId{
				TeamId:                   "teamId",
				OpenShiftChangeRequestId: "openShiftChangeRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/schedule/openShiftChangeRequests/openShiftChangeRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE/oPeNsHiFtChAnGeReQuEsTs/oPeNsHiFtChAnGeReQuEsTiD",
			Expected: &MeJoinedTeamIdScheduleOpenShiftChangeRequestId{
				TeamId:                   "tEaMiD",
				OpenShiftChangeRequestId: "oPeNsHiFtChAnGeReQuEsTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE/oPeNsHiFtChAnGeReQuEsTs/oPeNsHiFtChAnGeReQuEsTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdScheduleOpenShiftChangeRequestIDInsensitively(v.Input)
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

		if actual.OpenShiftChangeRequestId != v.Expected.OpenShiftChangeRequestId {
			t.Fatalf("Expected %q but got %q for OpenShiftChangeRequestId", v.Expected.OpenShiftChangeRequestId, actual.OpenShiftChangeRequestId)
		}

	}
}

func TestSegmentsForMeJoinedTeamIdScheduleOpenShiftChangeRequestId(t *testing.T) {
	segments := MeJoinedTeamIdScheduleOpenShiftChangeRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamIdScheduleOpenShiftChangeRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
