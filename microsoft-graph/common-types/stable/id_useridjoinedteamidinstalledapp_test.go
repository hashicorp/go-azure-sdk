package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdJoinedTeamIdInstalledAppId{}

func TestNewUserIdJoinedTeamIdInstalledAppID(t *testing.T) {
	id := NewUserIdJoinedTeamIdInstalledAppID("userId", "teamId", "teamsAppInstallationId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.TeamId != "teamId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamId")
	}

	if id.TeamsAppInstallationId != "teamsAppInstallationId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamsAppInstallationId'", id.TeamsAppInstallationId, "teamsAppInstallationId")
	}
}

func TestFormatUserIdJoinedTeamIdInstalledAppID(t *testing.T) {
	actual := NewUserIdJoinedTeamIdInstalledAppID("userId", "teamId", "teamsAppInstallationId").ID()
	expected := "/users/userId/joinedTeams/teamId/installedApps/teamsAppInstallationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdJoinedTeamIdInstalledAppID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdInstalledAppId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/installedApps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/joinedTeams/teamId/installedApps/teamsAppInstallationId",
			Expected: &UserIdJoinedTeamIdInstalledAppId{
				UserId:                 "userId",
				TeamId:                 "teamId",
				TeamsAppInstallationId: "teamsAppInstallationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/joinedTeams/teamId/installedApps/teamsAppInstallationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdInstalledAppID(v.Input)
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

		if actual.TeamsAppInstallationId != v.Expected.TeamsAppInstallationId {
			t.Fatalf("Expected %q but got %q for TeamsAppInstallationId", v.Expected.TeamsAppInstallationId, actual.TeamsAppInstallationId)
		}

	}
}

func TestParseUserIdJoinedTeamIdInstalledAppIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdJoinedTeamIdInstalledAppId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/joinedTeams/teamId/installedApps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/iNsTaLlEdApPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/joinedTeams/teamId/installedApps/teamsAppInstallationId",
			Expected: &UserIdJoinedTeamIdInstalledAppId{
				UserId:                 "userId",
				TeamId:                 "teamId",
				TeamsAppInstallationId: "teamsAppInstallationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/joinedTeams/teamId/installedApps/teamsAppInstallationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/iNsTaLlEdApPs/tEaMsApPiNsTaLlAtIoNiD",
			Expected: &UserIdJoinedTeamIdInstalledAppId{
				UserId:                 "uSeRiD",
				TeamId:                 "tEaMiD",
				TeamsAppInstallationId: "tEaMsApPiNsTaLlAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/jOiNeDtEaMs/tEaMiD/iNsTaLlEdApPs/tEaMsApPiNsTaLlAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdJoinedTeamIdInstalledAppIDInsensitively(v.Input)
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

		if actual.TeamsAppInstallationId != v.Expected.TeamsAppInstallationId {
			t.Fatalf("Expected %q but got %q for TeamsAppInstallationId", v.Expected.TeamsAppInstallationId, actual.TeamsAppInstallationId)
		}

	}
}

func TestSegmentsForUserIdJoinedTeamIdInstalledAppId(t *testing.T) {
	segments := UserIdJoinedTeamIdInstalledAppId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdJoinedTeamIdInstalledAppId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
