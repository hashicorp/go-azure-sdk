package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleShiftsRoleDefinitionId{}

func TestNewGroupIdTeamScheduleShiftsRoleDefinitionID(t *testing.T) {
	id := NewGroupIdTeamScheduleShiftsRoleDefinitionID("groupId", "shiftsRoleDefinitionId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.ShiftsRoleDefinitionId != "shiftsRoleDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ShiftsRoleDefinitionId'", id.ShiftsRoleDefinitionId, "shiftsRoleDefinitionId")
	}
}

func TestFormatGroupIdTeamScheduleShiftsRoleDefinitionID(t *testing.T) {
	actual := NewGroupIdTeamScheduleShiftsRoleDefinitionID("groupId", "shiftsRoleDefinitionId").ID()
	expected := "/groups/groupId/team/schedule/shiftsRoleDefinitions/shiftsRoleDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamScheduleShiftsRoleDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleShiftsRoleDefinitionId
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
			Input: "/groups/groupId/team/schedule",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/schedule/shiftsRoleDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/schedule/shiftsRoleDefinitions/shiftsRoleDefinitionId",
			Expected: &GroupIdTeamScheduleShiftsRoleDefinitionId{
				GroupId:                "groupId",
				ShiftsRoleDefinitionId: "shiftsRoleDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/schedule/shiftsRoleDefinitions/shiftsRoleDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleShiftsRoleDefinitionID(v.Input)
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

		if actual.ShiftsRoleDefinitionId != v.Expected.ShiftsRoleDefinitionId {
			t.Fatalf("Expected %q but got %q for ShiftsRoleDefinitionId", v.Expected.ShiftsRoleDefinitionId, actual.ShiftsRoleDefinitionId)
		}

	}
}

func TestParseGroupIdTeamScheduleShiftsRoleDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleShiftsRoleDefinitionId
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
			Input: "/groups/groupId/team/schedule",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/schedule/shiftsRoleDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/sHiFtSrOlEdEfInItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/schedule/shiftsRoleDefinitions/shiftsRoleDefinitionId",
			Expected: &GroupIdTeamScheduleShiftsRoleDefinitionId{
				GroupId:                "groupId",
				ShiftsRoleDefinitionId: "shiftsRoleDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/schedule/shiftsRoleDefinitions/shiftsRoleDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/sHiFtSrOlEdEfInItIoNs/sHiFtSrOlEdEfInItIoNiD",
			Expected: &GroupIdTeamScheduleShiftsRoleDefinitionId{
				GroupId:                "gRoUpId",
				ShiftsRoleDefinitionId: "sHiFtSrOlEdEfInItIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/sHiFtSrOlEdEfInItIoNs/sHiFtSrOlEdEfInItIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleShiftsRoleDefinitionIDInsensitively(v.Input)
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

		if actual.ShiftsRoleDefinitionId != v.Expected.ShiftsRoleDefinitionId {
			t.Fatalf("Expected %q but got %q for ShiftsRoleDefinitionId", v.Expected.ShiftsRoleDefinitionId, actual.ShiftsRoleDefinitionId)
		}

	}
}

func TestSegmentsForGroupIdTeamScheduleShiftsRoleDefinitionId(t *testing.T) {
	segments := GroupIdTeamScheduleShiftsRoleDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamScheduleShiftsRoleDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
