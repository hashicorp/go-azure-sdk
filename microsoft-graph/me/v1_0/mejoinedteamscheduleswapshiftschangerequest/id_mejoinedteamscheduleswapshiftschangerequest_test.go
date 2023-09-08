package mejoinedteamscheduleswapshiftschangerequest

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamScheduleSwapShiftsChangeRequestId{}

func TestNewMeJoinedTeamScheduleSwapShiftsChangeRequestID(t *testing.T) {
	id := NewMeJoinedTeamScheduleSwapShiftsChangeRequestID("teamIdValue", "swapShiftsChangeRequestIdValue")

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.SwapShiftsChangeRequestId != "swapShiftsChangeRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SwapShiftsChangeRequestId'", id.SwapShiftsChangeRequestId, "swapShiftsChangeRequestIdValue")
	}
}

func TestFormatMeJoinedTeamScheduleSwapShiftsChangeRequestID(t *testing.T) {
	actual := NewMeJoinedTeamScheduleSwapShiftsChangeRequestID("teamIdValue", "swapShiftsChangeRequestIdValue").ID()
	expected := "/me/joinedTeams/teamIdValue/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamScheduleSwapShiftsChangeRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamScheduleSwapShiftsChangeRequestId
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
			Input: "/me/joinedTeams/teamIdValue/schedule/swapShiftsChangeRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestIdValue",
			Expected: &MeJoinedTeamScheduleSwapShiftsChangeRequestId{
				TeamId:                    "teamIdValue",
				SwapShiftsChangeRequestId: "swapShiftsChangeRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamScheduleSwapShiftsChangeRequestID(v.Input)
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

func TestParseMeJoinedTeamScheduleSwapShiftsChangeRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamScheduleSwapShiftsChangeRequestId
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
			Input: "/me/joinedTeams/teamIdValue/schedule/swapShiftsChangeRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/sWaPsHiFtScHaNgErEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestIdValue",
			Expected: &MeJoinedTeamScheduleSwapShiftsChangeRequestId{
				TeamId:                    "teamIdValue",
				SwapShiftsChangeRequestId: "swapShiftsChangeRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/sWaPsHiFtScHaNgErEqUeStS/sWaPsHiFtScHaNgErEqUeStIdVaLuE",
			Expected: &MeJoinedTeamScheduleSwapShiftsChangeRequestId{
				TeamId:                    "tEaMiDvAlUe",
				SwapShiftsChangeRequestId: "sWaPsHiFtScHaNgErEqUeStIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/sWaPsHiFtScHaNgErEqUeStS/sWaPsHiFtScHaNgErEqUeStIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamScheduleSwapShiftsChangeRequestIDInsensitively(v.Input)
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

func TestSegmentsForMeJoinedTeamScheduleSwapShiftsChangeRequestId(t *testing.T) {
	segments := MeJoinedTeamScheduleSwapShiftsChangeRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamScheduleSwapShiftsChangeRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
