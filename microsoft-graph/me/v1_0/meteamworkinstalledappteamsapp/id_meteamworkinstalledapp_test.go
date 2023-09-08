package meteamworkinstalledappteamsapp

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeTeamworkInstalledAppId{}

func TestNewMeTeamworkInstalledAppID(t *testing.T) {
	id := NewMeTeamworkInstalledAppID("userScopeTeamsAppInstallationIdValue")

	if id.UserScopeTeamsAppInstallationId != "userScopeTeamsAppInstallationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserScopeTeamsAppInstallationId'", id.UserScopeTeamsAppInstallationId, "userScopeTeamsAppInstallationIdValue")
	}
}

func TestFormatMeTeamworkInstalledAppID(t *testing.T) {
	actual := NewMeTeamworkInstalledAppID("userScopeTeamsAppInstallationIdValue").ID()
	expected := "/me/teamwork/installedApps/userScopeTeamsAppInstallationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeTeamworkInstalledAppID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeTeamworkInstalledAppId
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
			Input: "/me/teamwork",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/teamwork/installedApps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/teamwork/installedApps/userScopeTeamsAppInstallationIdValue",
			Expected: &MeTeamworkInstalledAppId{
				UserScopeTeamsAppInstallationId: "userScopeTeamsAppInstallationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/teamwork/installedApps/userScopeTeamsAppInstallationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeTeamworkInstalledAppID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserScopeTeamsAppInstallationId != v.Expected.UserScopeTeamsAppInstallationId {
			t.Fatalf("Expected %q but got %q for UserScopeTeamsAppInstallationId", v.Expected.UserScopeTeamsAppInstallationId, actual.UserScopeTeamsAppInstallationId)
		}

	}
}

func TestParseMeTeamworkInstalledAppIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeTeamworkInstalledAppId
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
			Input: "/me/teamwork",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tEaMwOrK",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/teamwork/installedApps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tEaMwOrK/iNsTaLlEdApPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/teamwork/installedApps/userScopeTeamsAppInstallationIdValue",
			Expected: &MeTeamworkInstalledAppId{
				UserScopeTeamsAppInstallationId: "userScopeTeamsAppInstallationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/teamwork/installedApps/userScopeTeamsAppInstallationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/tEaMwOrK/iNsTaLlEdApPs/uSeRsCoPeTeAmSaPpInStAlLaTiOnIdVaLuE",
			Expected: &MeTeamworkInstalledAppId{
				UserScopeTeamsAppInstallationId: "uSeRsCoPeTeAmSaPpInStAlLaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/tEaMwOrK/iNsTaLlEdApPs/uSeRsCoPeTeAmSaPpInStAlLaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeTeamworkInstalledAppIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserScopeTeamsAppInstallationId != v.Expected.UserScopeTeamsAppInstallationId {
			t.Fatalf("Expected %q but got %q for UserScopeTeamsAppInstallationId", v.Expected.UserScopeTeamsAppInstallationId, actual.UserScopeTeamsAppInstallationId)
		}

	}
}

func TestSegmentsForMeTeamworkInstalledAppId(t *testing.T) {
	segments := MeTeamworkInstalledAppId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeTeamworkInstalledAppId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
