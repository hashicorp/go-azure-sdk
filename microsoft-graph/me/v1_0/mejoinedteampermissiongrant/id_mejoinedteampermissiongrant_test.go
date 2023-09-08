package mejoinedteampermissiongrant

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamPermissionGrantId{}

func TestNewMeJoinedTeamPermissionGrantID(t *testing.T) {
	id := NewMeJoinedTeamPermissionGrantID("teamIdValue", "resourceSpecificPermissionGrantIdValue")

	if id.TeamId != "teamIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamId'", id.TeamId, "teamIdValue")
	}

	if id.ResourceSpecificPermissionGrantId != "resourceSpecificPermissionGrantIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceSpecificPermissionGrantId'", id.ResourceSpecificPermissionGrantId, "resourceSpecificPermissionGrantIdValue")
	}
}

func TestFormatMeJoinedTeamPermissionGrantID(t *testing.T) {
	actual := NewMeJoinedTeamPermissionGrantID("teamIdValue", "resourceSpecificPermissionGrantIdValue").ID()
	expected := "/me/joinedTeams/teamIdValue/permissionGrants/resourceSpecificPermissionGrantIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeJoinedTeamPermissionGrantID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamPermissionGrantId
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
			Input: "/me/joinedTeams/teamIdValue/permissionGrants",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/permissionGrants/resourceSpecificPermissionGrantIdValue",
			Expected: &MeJoinedTeamPermissionGrantId{
				TeamId:                            "teamIdValue",
				ResourceSpecificPermissionGrantId: "resourceSpecificPermissionGrantIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/permissionGrants/resourceSpecificPermissionGrantIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamPermissionGrantID(v.Input)
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

		if actual.ResourceSpecificPermissionGrantId != v.Expected.ResourceSpecificPermissionGrantId {
			t.Fatalf("Expected %q but got %q for ResourceSpecificPermissionGrantId", v.Expected.ResourceSpecificPermissionGrantId, actual.ResourceSpecificPermissionGrantId)
		}

	}
}

func TestParseMeJoinedTeamPermissionGrantIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeJoinedTeamPermissionGrantId
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
			Input: "/me/joinedTeams/teamIdValue/permissionGrants",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/pErMiSsIoNgRaNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/joinedTeams/teamIdValue/permissionGrants/resourceSpecificPermissionGrantIdValue",
			Expected: &MeJoinedTeamPermissionGrantId{
				TeamId:                            "teamIdValue",
				ResourceSpecificPermissionGrantId: "resourceSpecificPermissionGrantIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/joinedTeams/teamIdValue/permissionGrants/resourceSpecificPermissionGrantIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/pErMiSsIoNgRaNtS/rEsOuRcEsPeCiFiCpErMiSsIoNgRaNtIdVaLuE",
			Expected: &MeJoinedTeamPermissionGrantId{
				TeamId:                            "tEaMiDvAlUe",
				ResourceSpecificPermissionGrantId: "rEsOuRcEsPeCiFiCpErMiSsIoNgRaNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/jOiNeDtEaMs/tEaMiDvAlUe/pErMiSsIoNgRaNtS/rEsOuRcEsPeCiFiCpErMiSsIoNgRaNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeJoinedTeamPermissionGrantIDInsensitively(v.Input)
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

		if actual.ResourceSpecificPermissionGrantId != v.Expected.ResourceSpecificPermissionGrantId {
			t.Fatalf("Expected %q but got %q for ResourceSpecificPermissionGrantId", v.Expected.ResourceSpecificPermissionGrantId, actual.ResourceSpecificPermissionGrantId)
		}

	}
}

func TestSegmentsForMeJoinedTeamPermissionGrantId(t *testing.T) {
	segments := MeJoinedTeamPermissionGrantId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeJoinedTeamPermissionGrantId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
