package groupteaminstalledappteamsapp

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamInstalledAppId{}

func TestNewGroupTeamInstalledAppID(t *testing.T) {
	id := NewGroupTeamInstalledAppID("groupIdValue", "teamsAppInstallationIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.TeamsAppInstallationId != "teamsAppInstallationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamsAppInstallationId'", id.TeamsAppInstallationId, "teamsAppInstallationIdValue")
	}
}

func TestFormatGroupTeamInstalledAppID(t *testing.T) {
	actual := NewGroupTeamInstalledAppID("groupIdValue", "teamsAppInstallationIdValue").ID()
	expected := "/groups/groupIdValue/team/installedApps/teamsAppInstallationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupTeamInstalledAppID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamInstalledAppId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/installedApps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/installedApps/teamsAppInstallationIdValue",
			Expected: &GroupTeamInstalledAppId{
				GroupId:                "groupIdValue",
				TeamsAppInstallationId: "teamsAppInstallationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/installedApps/teamsAppInstallationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamInstalledAppID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.TeamsAppInstallationId != v.Expected.TeamsAppInstallationId {
			t.Fatalf("Expected %q but got %q for TeamsAppInstallationId", v.Expected.TeamsAppInstallationId, actual.TeamsAppInstallationId)
		}

	}
}

func TestParseGroupTeamInstalledAppIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamInstalledAppId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/installedApps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/iNsTaLlEdApPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/installedApps/teamsAppInstallationIdValue",
			Expected: &GroupTeamInstalledAppId{
				GroupId:                "groupIdValue",
				TeamsAppInstallationId: "teamsAppInstallationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/installedApps/teamsAppInstallationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/iNsTaLlEdApPs/tEaMsApPiNsTaLlAtIoNiDvAlUe",
			Expected: &GroupTeamInstalledAppId{
				GroupId:                "gRoUpIdVaLuE",
				TeamsAppInstallationId: "tEaMsApPiNsTaLlAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/iNsTaLlEdApPs/tEaMsApPiNsTaLlAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamInstalledAppIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.TeamsAppInstallationId != v.Expected.TeamsAppInstallationId {
			t.Fatalf("Expected %q but got %q for TeamsAppInstallationId", v.Expected.TeamsAppInstallationId, actual.TeamsAppInstallationId)
		}

	}
}

func TestSegmentsForGroupTeamInstalledAppId(t *testing.T) {
	segments := GroupTeamInstalledAppId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupTeamInstalledAppId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
