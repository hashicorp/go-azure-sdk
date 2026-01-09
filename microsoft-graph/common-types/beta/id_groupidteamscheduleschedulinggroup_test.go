package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleSchedulingGroupId{}

func TestNewGroupIdTeamScheduleSchedulingGroupID(t *testing.T) {
	id := NewGroupIdTeamScheduleSchedulingGroupID("groupId", "schedulingGroupId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.SchedulingGroupId != "schedulingGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'SchedulingGroupId'", id.SchedulingGroupId, "schedulingGroupId")
	}
}

func TestFormatGroupIdTeamScheduleSchedulingGroupID(t *testing.T) {
	actual := NewGroupIdTeamScheduleSchedulingGroupID("groupId", "schedulingGroupId").ID()
	expected := "/groups/groupId/team/schedule/schedulingGroups/schedulingGroupId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamScheduleSchedulingGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleSchedulingGroupId
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
			Input: "/groups/groupId/team/schedule/schedulingGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/schedule/schedulingGroups/schedulingGroupId",
			Expected: &GroupIdTeamScheduleSchedulingGroupId{
				GroupId:           "groupId",
				SchedulingGroupId: "schedulingGroupId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/schedule/schedulingGroups/schedulingGroupId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleSchedulingGroupID(v.Input)
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

		if actual.SchedulingGroupId != v.Expected.SchedulingGroupId {
			t.Fatalf("Expected %q but got %q for SchedulingGroupId", v.Expected.SchedulingGroupId, actual.SchedulingGroupId)
		}

	}
}

func TestParseGroupIdTeamScheduleSchedulingGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleSchedulingGroupId
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
			Input: "/groups/groupId/team/schedule/schedulingGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/sChEdUlInGgRoUpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/schedule/schedulingGroups/schedulingGroupId",
			Expected: &GroupIdTeamScheduleSchedulingGroupId{
				GroupId:           "groupId",
				SchedulingGroupId: "schedulingGroupId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/schedule/schedulingGroups/schedulingGroupId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/sChEdUlInGgRoUpS/sChEdUlInGgRoUpId",
			Expected: &GroupIdTeamScheduleSchedulingGroupId{
				GroupId:           "gRoUpId",
				SchedulingGroupId: "sChEdUlInGgRoUpId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/sChEdUlInGgRoUpS/sChEdUlInGgRoUpId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleSchedulingGroupIDInsensitively(v.Input)
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

		if actual.SchedulingGroupId != v.Expected.SchedulingGroupId {
			t.Fatalf("Expected %q but got %q for SchedulingGroupId", v.Expected.SchedulingGroupId, actual.SchedulingGroupId)
		}

	}
}

func TestSegmentsForGroupIdTeamScheduleSchedulingGroupId(t *testing.T) {
	segments := GroupIdTeamScheduleSchedulingGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamScheduleSchedulingGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
