package mejoinedteamscheduleopenshiftchangerequest

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamScheduleOpenShiftChangeRequestId{}

func TestNewMeJoinedTeamScheduleOpenShiftChangeRequestID(t *testing.T) {
	id := NewMeJoinedTeamScheduleOpenShiftChangeRequestID("teamIdValue", "openShiftChangeRequestIdValue")

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.OpenShiftChangeRequestId != "openShiftChangeRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OpenShiftChangeRequestId'", id.OpenShiftChangeRequestId, "openShiftChangeRequestIdValue")
	}
}

func TestFormatMeJoinedTeamScheduleOpenShiftChangeRequestID(t *testing.T) {
	actual := NewMeJoinedTeamScheduleOpenShiftChangeRequestID("teamIdValue", "openShiftChangeRequestIdValue").ID()
	expected := "/me/joinedTeams/teamIdValue/schedule/openShiftChangeRequests/openShiftChangeRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamScheduleOpenShiftChangeRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamScheduleOpenShiftChangeRequestId
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
			Input: "/me/joinedTeams/teamIdValue/schedule/openShiftChangeRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/schedule/openShiftChangeRequests/openShiftChangeRequestIdValue",
			Expected: &MeJoinedTeamScheduleOpenShiftChangeRequestId{
				TeamId:                   "teamIdValue",
				OpenShiftChangeRequestId: "openShiftChangeRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/schedule/openShiftChangeRequests/openShiftChangeRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamScheduleOpenShiftChangeRequestID(v.Input)
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

func TestParseMeJoinedTeamScheduleOpenShiftChangeRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamScheduleOpenShiftChangeRequestId
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
			Input: "/me/joinedTeams/teamIdValue/schedule/openShiftChangeRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/oPeNsHiFtChAnGeReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/schedule/openShiftChangeRequests/openShiftChangeRequestIdValue",
			Expected: &MeJoinedTeamScheduleOpenShiftChangeRequestId{
				TeamId:                   "teamIdValue",
				OpenShiftChangeRequestId: "openShiftChangeRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/schedule/openShiftChangeRequests/openShiftChangeRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/oPeNsHiFtChAnGeReQuEsTs/oPeNsHiFtChAnGeReQuEsTiDvAlUe",
			Expected: &MeJoinedTeamScheduleOpenShiftChangeRequestId{
				TeamId:                   "tEaMiDvAlUe",
				OpenShiftChangeRequestId: "oPeNsHiFtChAnGeReQuEsTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/oPeNsHiFtChAnGeReQuEsTs/oPeNsHiFtChAnGeReQuEsTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamScheduleOpenShiftChangeRequestIDInsensitively(v.Input)
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

func TestSegmentsForMeJoinedTeamScheduleOpenShiftChangeRequestId(t *testing.T) {
	segments := MeJoinedTeamScheduleOpenShiftChangeRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamScheduleOpenShiftChangeRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
