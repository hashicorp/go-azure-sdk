package joinedteamscheduleswapshiftschangerequest

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId{}

func TestNewUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestID(t *testing.T) {
	id := NewUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestID("userIdValue", "teamIdValue", "swapShiftsChangeRequestIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.SwapShiftsChangeRequestId != "swapShiftsChangeRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SwapShiftsChangeRequestId'", id.SwapShiftsChangeRequestId, "swapShiftsChangeRequestIdValue")
	}
}

func TestFormatUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestID(t *testing.T) {
	actual := NewUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestID("userIdValue", "teamIdValue", "swapShiftsChangeRequestIdValue").ID()
	expected := "/users/userIdValue/joinedTeams/teamIdValue/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId
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
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/swapShiftsChangeRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestIdValue",
			Expected: &UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId{
				UserId:                    "userIdValue",
				TeamId:                    "teamIdValue",
				SwapShiftsChangeRequestId: "swapShiftsChangeRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestID(v.Input)
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

		if actual.SwapShiftsChangeRequestId != v.Expected.SwapShiftsChangeRequestId {
			t.Fatalf("Expected %q but got %q for SwapShiftsChangeRequestId", v.Expected.SwapShiftsChangeRequestId, actual.SwapShiftsChangeRequestId)
		}

	}
}

func TestParseUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId
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
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/swapShiftsChangeRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/sWaPsHiFtScHaNgErEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestIdValue",
			Expected: &UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId{
				UserId:                    "userIdValue",
				TeamId:                    "teamIdValue",
				SwapShiftsChangeRequestId: "swapShiftsChangeRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/joinedTeams/teamIdValue/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/sWaPsHiFtScHaNgErEqUeStS/sWaPsHiFtScHaNgErEqUeStIdVaLuE",
			Expected: &UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId{
				UserId:                    "uSeRiDvAlUe",
				TeamId:                    "tEaMiDvAlUe",
				SwapShiftsChangeRequestId: "sWaPsHiFtScHaNgErEqUeStIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/jOiNeDtEaMs/tEaMiDvAlUe/sChEdUlE/sWaPsHiFtScHaNgErEqUeStS/sWaPsHiFtScHaNgErEqUeStIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestIDInsensitively(v.Input)
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

		if actual.SwapShiftsChangeRequestId != v.Expected.SwapShiftsChangeRequestId {
			t.Fatalf("Expected %q but got %q for SwapShiftsChangeRequestId", v.Expected.SwapShiftsChangeRequestId, actual.SwapShiftsChangeRequestId)
		}

	}
}

func TestSegmentsForUserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId(t *testing.T) {
	segments := UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdJoinedTeamIdScheduleSwapShiftsChangeRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
