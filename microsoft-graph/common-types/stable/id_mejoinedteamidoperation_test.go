package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdOperationId{}

func TestNewMeJoinedTeamIdOperationID(t *testing.T) {
	id := NewMeJoinedTeamIdOperationID("teamIdValue", "teamsAsyncOperationIdValue")

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.TeamsAsyncOperationId != "teamsAsyncOperationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamsAsyncOperationId'", id.TeamsAsyncOperationId, "teamsAsyncOperationIdValue")
	}
}

func TestFormatMeJoinedTeamIdOperationID(t *testing.T) {
	actual := NewMeJoinedTeamIdOperationID("teamIdValue", "teamsAsyncOperationIdValue").ID()
	expected := "/me/joinedTeams/teamIdValue/operations/teamsAsyncOperationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamIdOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdOperationId
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
			Input: "/me/joinedTeams/teamIdValue/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/operations/teamsAsyncOperationIdValue",
			Expected: &MeJoinedTeamIdOperationId{
				TeamId:                "teamIdValue",
				TeamsAsyncOperationId: "teamsAsyncOperationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/operations/teamsAsyncOperationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdOperationID(v.Input)
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

		if actual.TeamsAsyncOperationId != v.Expected.TeamsAsyncOperationId {
			t.Fatalf("Expected %q but got %q for TeamsAsyncOperationId", v.Expected.TeamsAsyncOperationId, actual.TeamsAsyncOperationId)
		}

	}
}

func TestParseMeJoinedTeamIdOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamIdOperationId
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
			Input: "/me/joinedTeams/teamIdValue/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/operations/teamsAsyncOperationIdValue",
			Expected: &MeJoinedTeamIdOperationId{
				TeamId:                "teamIdValue",
				TeamsAsyncOperationId: "teamsAsyncOperationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/operations/teamsAsyncOperationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/oPeRaTiOnS/tEaMsAsYnCoPeRaTiOnIdVaLuE",
			Expected: &MeJoinedTeamIdOperationId{
				TeamId:                "tEaMiDvAlUe",
				TeamsAsyncOperationId: "tEaMsAsYnCoPeRaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/oPeRaTiOnS/tEaMsAsYnCoPeRaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamIdOperationIDInsensitively(v.Input)
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

		if actual.TeamsAsyncOperationId != v.Expected.TeamsAsyncOperationId {
			t.Fatalf("Expected %q but got %q for TeamsAsyncOperationId", v.Expected.TeamsAsyncOperationId, actual.TeamsAsyncOperationId)
		}

	}
}

func TestSegmentsForMeJoinedTeamIdOperationId(t *testing.T) {
	segments := MeJoinedTeamIdOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamIdOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
