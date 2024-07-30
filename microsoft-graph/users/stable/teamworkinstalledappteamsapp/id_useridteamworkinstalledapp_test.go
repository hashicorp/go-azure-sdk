package teamworkinstalledappteamsapp

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdTeamworkInstalledAppId{}

func TestNewUserIdTeamworkInstalledAppID(t *testing.T) {
	id := NewUserIdTeamworkInstalledAppID("userIdValue", "userScopeTeamsAppInstallationIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.UserScopeTeamsAppInstallationId != "userScopeTeamsAppInstallationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserScopeTeamsAppInstallationId'", id.UserScopeTeamsAppInstallationId, "userScopeTeamsAppInstallationIdValue")
	}
}

func TestFormatUserIdTeamworkInstalledAppID(t *testing.T) {
	actual := NewUserIdTeamworkInstalledAppID("userIdValue", "userScopeTeamsAppInstallationIdValue").ID()
	expected := "/users/userIdValue/teamwork/installedApps/userScopeTeamsAppInstallationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdTeamworkInstalledAppID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdTeamworkInstalledAppId
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
			Input: "/users/userIdValue/teamwork",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/teamwork/installedApps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/teamwork/installedApps/userScopeTeamsAppInstallationIdValue",
			Expected: &UserIdTeamworkInstalledAppId{
				UserId:                          "userIdValue",
				UserScopeTeamsAppInstallationId: "userScopeTeamsAppInstallationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/teamwork/installedApps/userScopeTeamsAppInstallationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdTeamworkInstalledAppID(v.Input)
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

		if actual.UserScopeTeamsAppInstallationId != v.Expected.UserScopeTeamsAppInstallationId {
			t.Fatalf("Expected %q but got %q for UserScopeTeamsAppInstallationId", v.Expected.UserScopeTeamsAppInstallationId, actual.UserScopeTeamsAppInstallationId)
		}

	}
}

func TestParseUserIdTeamworkInstalledAppIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdTeamworkInstalledAppId
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
			Input: "/users/userIdValue/teamwork",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tEaMwOrK",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/teamwork/installedApps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tEaMwOrK/iNsTaLlEdApPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/teamwork/installedApps/userScopeTeamsAppInstallationIdValue",
			Expected: &UserIdTeamworkInstalledAppId{
				UserId:                          "userIdValue",
				UserScopeTeamsAppInstallationId: "userScopeTeamsAppInstallationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/teamwork/installedApps/userScopeTeamsAppInstallationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tEaMwOrK/iNsTaLlEdApPs/uSeRsCoPeTeAmSaPpInStAlLaTiOnIdVaLuE",
			Expected: &UserIdTeamworkInstalledAppId{
				UserId:                          "uSeRiDvAlUe",
				UserScopeTeamsAppInstallationId: "uSeRsCoPeTeAmSaPpInStAlLaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/tEaMwOrK/iNsTaLlEdApPs/uSeRsCoPeTeAmSaPpInStAlLaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdTeamworkInstalledAppIDInsensitively(v.Input)
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

		if actual.UserScopeTeamsAppInstallationId != v.Expected.UserScopeTeamsAppInstallationId {
			t.Fatalf("Expected %q but got %q for UserScopeTeamsAppInstallationId", v.Expected.UserScopeTeamsAppInstallationId, actual.UserScopeTeamsAppInstallationId)
		}

	}
}

func TestSegmentsForUserIdTeamworkInstalledAppId(t *testing.T) {
	segments := UserIdTeamworkInstalledAppId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdTeamworkInstalledAppId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
