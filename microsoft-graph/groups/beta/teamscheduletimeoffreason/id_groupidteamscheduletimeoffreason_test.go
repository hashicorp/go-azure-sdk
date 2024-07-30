package teamscheduletimeoffreason

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleTimeOffReasonId{}

func TestNewGroupIdTeamScheduleTimeOffReasonID(t *testing.T) {
	id := NewGroupIdTeamScheduleTimeOffReasonID("groupIdValue", "timeOffReasonIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.TimeOffReasonId != "timeOffReasonIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TimeOffReasonId'", id.TimeOffReasonId, "timeOffReasonIdValue")
	}
}

func TestFormatGroupIdTeamScheduleTimeOffReasonID(t *testing.T) {
	actual := NewGroupIdTeamScheduleTimeOffReasonID("groupIdValue", "timeOffReasonIdValue").ID()
	expected := "/groups/groupIdValue/team/schedule/timeOffReasons/timeOffReasonIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamScheduleTimeOffReasonID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleTimeOffReasonId
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
			Input: "/groups/groupIdValue/team/schedule/timeOffReasons",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/schedule/timeOffReasons/timeOffReasonIdValue",
			Expected: &GroupIdTeamScheduleTimeOffReasonId{
				GroupId:         "groupIdValue",
				TimeOffReasonId: "timeOffReasonIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/schedule/timeOffReasons/timeOffReasonIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleTimeOffReasonID(v.Input)
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

		if actual.TimeOffReasonId != v.Expected.TimeOffReasonId {
			t.Fatalf("Expected %q but got %q for TimeOffReasonId", v.Expected.TimeOffReasonId, actual.TimeOffReasonId)
		}

	}
}

func TestParseGroupIdTeamScheduleTimeOffReasonIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleTimeOffReasonId
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
			Input: "/groups/groupIdValue/team/schedule/timeOffReasons",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/tImEoFfReAsOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/schedule/timeOffReasons/timeOffReasonIdValue",
			Expected: &GroupIdTeamScheduleTimeOffReasonId{
				GroupId:         "groupIdValue",
				TimeOffReasonId: "timeOffReasonIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/schedule/timeOffReasons/timeOffReasonIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/tImEoFfReAsOnS/tImEoFfReAsOnIdVaLuE",
			Expected: &GroupIdTeamScheduleTimeOffReasonId{
				GroupId:         "gRoUpIdVaLuE",
				TimeOffReasonId: "tImEoFfReAsOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/tImEoFfReAsOnS/tImEoFfReAsOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleTimeOffReasonIDInsensitively(v.Input)
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

		if actual.TimeOffReasonId != v.Expected.TimeOffReasonId {
			t.Fatalf("Expected %q but got %q for TimeOffReasonId", v.Expected.TimeOffReasonId, actual.TimeOffReasonId)
		}

	}
}

func TestSegmentsForGroupIdTeamScheduleTimeOffReasonId(t *testing.T) {
	segments := GroupIdTeamScheduleTimeOffReasonId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamScheduleTimeOffReasonId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
