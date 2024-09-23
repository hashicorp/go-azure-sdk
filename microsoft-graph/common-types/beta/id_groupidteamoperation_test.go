package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamOperationId{}

func TestNewGroupIdTeamOperationID(t *testing.T) {
	id := NewGroupIdTeamOperationID("groupId", "teamsAsyncOperationId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.TeamsAsyncOperationId != "teamsAsyncOperationId" {
		t.Fatalf("Expected %q but got %q for Segment 'TeamsAsyncOperationId'", id.TeamsAsyncOperationId, "teamsAsyncOperationId")
	}
}

func TestFormatGroupIdTeamOperationID(t *testing.T) {
	actual := NewGroupIdTeamOperationID("groupId", "teamsAsyncOperationId").ID()
	expected := "/groups/groupId/team/operations/teamsAsyncOperationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamOperationId
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
			Input: "/groups/groupId/team/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/operations/teamsAsyncOperationId",
			Expected: &GroupIdTeamOperationId{
				GroupId:               "groupId",
				TeamsAsyncOperationId: "teamsAsyncOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/operations/teamsAsyncOperationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamOperationID(v.Input)
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

		if actual.TeamsAsyncOperationId != v.Expected.TeamsAsyncOperationId {
			t.Fatalf("Expected %q but got %q for TeamsAsyncOperationId", v.Expected.TeamsAsyncOperationId, actual.TeamsAsyncOperationId)
		}

	}
}

func TestParseGroupIdTeamOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamOperationId
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
			Input: "/groups/groupId/team/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/operations/teamsAsyncOperationId",
			Expected: &GroupIdTeamOperationId{
				GroupId:               "groupId",
				TeamsAsyncOperationId: "teamsAsyncOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/operations/teamsAsyncOperationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/oPeRaTiOnS/tEaMsAsYnCoPeRaTiOnId",
			Expected: &GroupIdTeamOperationId{
				GroupId:               "gRoUpId",
				TeamsAsyncOperationId: "tEaMsAsYnCoPeRaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/oPeRaTiOnS/tEaMsAsYnCoPeRaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamOperationIDInsensitively(v.Input)
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

		if actual.TeamsAsyncOperationId != v.Expected.TeamsAsyncOperationId {
			t.Fatalf("Expected %q but got %q for TeamsAsyncOperationId", v.Expected.TeamsAsyncOperationId, actual.TeamsAsyncOperationId)
		}

	}
}

func TestSegmentsForGroupIdTeamOperationId(t *testing.T) {
	segments := GroupIdTeamOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
