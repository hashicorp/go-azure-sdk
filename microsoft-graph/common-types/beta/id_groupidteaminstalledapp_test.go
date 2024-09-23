package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamInstalledAppId{}

func TestNewGroupIdTeamInstalledAppID(t *testing.T) {
	id := NewGroupIdTeamInstalledAppID("groupId", "teamsAppInstallationId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.TeamsAppInstallationId != "teamsAppInstallationId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamsAppInstallationId'", id.TeamsAppInstallationId, "teamsAppInstallationId")
	}
}

func TestFormatGroupIdTeamInstalledAppID(t *testing.T) {
	actual := NewGroupIdTeamInstalledAppID("groupId", "teamsAppInstallationId").ID()
	expected := "/groups/groupId/team/installedApps/teamsAppInstallationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamInstalledAppID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamInstalledAppId
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
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/installedApps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/installedApps/teamsAppInstallationId",
			Expected: &GroupIdTeamInstalledAppId{
				GroupId:                "groupId",
				TeamsAppInstallationId: "teamsAppInstallationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/installedApps/teamsAppInstallationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamInstalledAppID(v.Input)
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

func TestParseGroupIdTeamInstalledAppIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamInstalledAppId
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
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/installedApps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/iNsTaLlEdApPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/installedApps/teamsAppInstallationId",
			Expected: &GroupIdTeamInstalledAppId{
				GroupId:                "groupId",
				TeamsAppInstallationId: "teamsAppInstallationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/installedApps/teamsAppInstallationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/iNsTaLlEdApPs/tEaMsApPiNsTaLlAtIoNiD",
			Expected: &GroupIdTeamInstalledAppId{
				GroupId:                "gRoUpId",
				TeamsAppInstallationId: "tEaMsApPiNsTaLlAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/iNsTaLlEdApPs/tEaMsApPiNsTaLlAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamInstalledAppIDInsensitively(v.Input)
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

func TestSegmentsForGroupIdTeamInstalledAppId(t *testing.T) {
	segments := GroupIdTeamInstalledAppId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamInstalledAppId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
