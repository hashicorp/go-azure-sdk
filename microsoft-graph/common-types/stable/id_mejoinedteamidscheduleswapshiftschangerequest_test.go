package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdScheduleSwapShiftsChangeRequestId{}

func TestNewMeJoinedTeamIdScheduleSwapShiftsChangeRequestID(t *testing.T) {
	id := NewMeJoinedTeamIdScheduleSwapShiftsChangeRequestID("teamId", "swapShiftsChangeRequestId")

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.SwapShiftsChangeRequestId != "swapShiftsChangeRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'SwapShiftsChangeRequestId'", id.SwapShiftsChangeRequestId, "swapShiftsChangeRequestId")
	}
}

func TestFormatMeJoinedTeamIdScheduleSwapShiftsChangeRequestID(t *testing.T) {
	actual := NewMeJoinedTeamIdScheduleSwapShiftsChangeRequestID("teamId", "swapShiftsChangeRequestId").ID()
	expected := "/me/joinedTeams/teamId/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamIdScheduleSwapShiftsChangeRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdScheduleSwapShiftsChangeRequestId
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
			Input: "/me/joinedTeams/teamId/schedule/swapShiftsChangeRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestId",
			Expected: &MeJoinedTeamIdScheduleSwapShiftsChangeRequestId{
				TeamId:                    "teamId",
				SwapShiftsChangeRequestId: "swapShiftsChangeRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdScheduleSwapShiftsChangeRequestID(v.Input)
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

		if actual.SwapShiftsChangeRequestId != v.Expected.SwapShiftsChangeRequestId {
			t.Fatalf("Expected %q but got %q for SwapShiftsChangeRequestId", v.Expected.SwapShiftsChangeRequestId, actual.SwapShiftsChangeRequestId)
		}

	}
}

func TestParseMeJoinedTeamIdScheduleSwapShiftsChangeRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdScheduleSwapShiftsChangeRequestId
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
			Input: "/me/joinedTeams/teamId/schedule/swapShiftsChangeRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE/sWaPsHiFtScHaNgErEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamId/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestId",
			Expected: &MeJoinedTeamIdScheduleSwapShiftsChangeRequestId{
				TeamId:                    "teamId",
				SwapShiftsChangeRequestId: "swapShiftsChangeRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamId/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE/sWaPsHiFtScHaNgErEqUeStS/sWaPsHiFtScHaNgErEqUeStId",
			Expected: &MeJoinedTeamIdScheduleSwapShiftsChangeRequestId{
				TeamId:                    "tEaMiD",
				SwapShiftsChangeRequestId: "sWaPsHiFtScHaNgErEqUeStId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiD/sChEdUlE/sWaPsHiFtScHaNgErEqUeStS/sWaPsHiFtScHaNgErEqUeStId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdScheduleSwapShiftsChangeRequestIDInsensitively(v.Input)
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

		if actual.SwapShiftsChangeRequestId != v.Expected.SwapShiftsChangeRequestId {
			t.Fatalf("Expected %q but got %q for SwapShiftsChangeRequestId", v.Expected.SwapShiftsChangeRequestId, actual.SwapShiftsChangeRequestId)
		}

	}
}

func TestSegmentsForMeJoinedTeamIdScheduleSwapShiftsChangeRequestId(t *testing.T) {
	segments := MeJoinedTeamIdScheduleSwapShiftsChangeRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamIdScheduleSwapShiftsChangeRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
