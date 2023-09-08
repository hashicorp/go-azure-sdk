package groupteamscheduleschedulinggroup

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamScheduleSchedulingGroupId{}

func TestNewGroupTeamScheduleSchedulingGroupID(t *testing.T) {
	id := NewGroupTeamScheduleSchedulingGroupID("groupIdValue", "schedulingGroupIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SchedulingGroupId != "schedulingGroupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SchedulingGroupId'", id.SchedulingGroupId, "schedulingGroupIdValue")
	}
}

func TestFormatGroupTeamScheduleSchedulingGroupID(t *testing.T) {
	actual := NewGroupTeamScheduleSchedulingGroupID("groupIdValue", "schedulingGroupIdValue").ID()
	expected := "/groups/groupIdValue/team/schedule/schedulingGroups/schedulingGroupIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupTeamScheduleSchedulingGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamScheduleSchedulingGroupId
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
			Input: "/groups/groupIdValue/team/schedule",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/schedule/schedulingGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/schedule/schedulingGroups/schedulingGroupIdValue",
			Expected: &GroupTeamScheduleSchedulingGroupId{
				GroupId:           "groupIdValue",
				SchedulingGroupId: "schedulingGroupIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/schedule/schedulingGroups/schedulingGroupIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamScheduleSchedulingGroupID(v.Input)
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

func TestParseGroupTeamScheduleSchedulingGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamScheduleSchedulingGroupId
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
			Input: "/groups/groupIdValue/team/schedule",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/schedule/schedulingGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/sChEdUlInGgRoUpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/schedule/schedulingGroups/schedulingGroupIdValue",
			Expected: &GroupTeamScheduleSchedulingGroupId{
				GroupId:           "groupIdValue",
				SchedulingGroupId: "schedulingGroupIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/schedule/schedulingGroups/schedulingGroupIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/sChEdUlInGgRoUpS/sChEdUlInGgRoUpIdVaLuE",
			Expected: &GroupTeamScheduleSchedulingGroupId{
				GroupId:           "gRoUpIdVaLuE",
				SchedulingGroupId: "sChEdUlInGgRoUpIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/sChEdUlInGgRoUpS/sChEdUlInGgRoUpIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamScheduleSchedulingGroupIDInsensitively(v.Input)
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

func TestSegmentsForGroupTeamScheduleSchedulingGroupId(t *testing.T) {
	segments := GroupTeamScheduleSchedulingGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupTeamScheduleSchedulingGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
