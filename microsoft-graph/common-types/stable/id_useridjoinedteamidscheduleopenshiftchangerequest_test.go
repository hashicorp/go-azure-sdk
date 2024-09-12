package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId{}

func TestNewUserIdJoinedTeamIdScheduleOpenShiftChangeRequestID(t *testing.T) {
	id := NewUserIdJoinedTeamIdScheduleOpenShiftChangeRequestID("userIdValue", "teamIdValue", "openShiftChangeRequestIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.OpenShiftChangeRequestId != "openShiftChangeRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OpenShiftChangeRequestId'", id.OpenShiftChangeRequestId, "openShiftChangeRequestIdValue")
	}
}

func TestFormatUserIdJoinedTeamIdScheduleOpenShiftChangeRequestID(t *testing.T) {
	actual := NewUserIdJoinedTeamIdScheduleOpenShiftChangeRequestID("userIdValue", "teamIdValue", "openShiftChangeRequestIdValue").ID()
	expected := "/users/userIdValue/joinedTeams/teamIdValue/schedule/openShiftChangeRequests/openShiftChangeRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdJoinedTeamIdScheduleOpenShiftChangeRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId
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
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/openShiftChangeRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/openShiftChangeRequests/openShiftChangeRequestIdValue",
			Expected: &UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId{
				UserId:                   "userIdValue",
				TeamId:                   "teamIdValue",
				OpenShiftChangeRequestId: "openShiftChangeRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/openShiftChangeRequests/openShiftChangeRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdScheduleOpenShiftChangeRequestID(v.Input)
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

		if actual.OpenShiftChangeRequestId != v.Expected.OpenShiftChangeRequestId {
			t.Fatalf("Expected %q but got %q for OpenShiftChangeRequestId", v.Expected.OpenShiftChangeRequestId, actual.OpenShiftChangeRequestId)
		}

	}
}

func TestParseUserIdJoinedTeamIdScheduleOpenShiftChangeRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId
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
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/openShiftChangeRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/oPeNsHiFtChAnGeReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/openShiftChangeRequests/openShiftChangeRequestIdValue",
			Expected: &UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId{
				UserId:                   "userIdValue",
				TeamId:                   "teamIdValue",
				OpenShiftChangeRequestId: "openShiftChangeRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/openShiftChangeRequests/openShiftChangeRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/oPeNsHiFtChAnGeReQuEsTs/oPeNsHiFtChAnGeReQuEsTiDvAlUe",
			Expected: &UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId{
				UserId:                   "uSeRiDvAlUe",
				TeamId:                   "tEaMiDvAlUe",
				OpenShiftChangeRequestId: "oPeNsHiFtChAnGeReQuEsTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/oPeNsHiFtChAnGeReQuEsTs/oPeNsHiFtChAnGeReQuEsTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdScheduleOpenShiftChangeRequestIDInsensitively(v.Input)
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

		if actual.OpenShiftChangeRequestId != v.Expected.OpenShiftChangeRequestId {
			t.Fatalf("Expected %q but got %q for OpenShiftChangeRequestId", v.Expected.OpenShiftChangeRequestId, actual.OpenShiftChangeRequestId)
		}

	}
}

func TestSegmentsForUserIdJoinedTeamIdScheduleOpenShiftChangeRequestId(t *testing.T) {
	segments := UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdJoinedTeamIdScheduleOpenShiftChangeRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
