package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleTimesOffId{}

func TestNewGroupIdTeamScheduleTimesOffID(t *testing.T) {
	id := NewGroupIdTeamScheduleTimesOffID("groupId", "timeOffId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.TimeOffId != "timeOffId" {
		t.Fatalf("Expected %q but got %q for Segment 'TimeOffId'", id.TimeOffId, "timeOffId")
	}
}

func TestFormatGroupIdTeamScheduleTimesOffID(t *testing.T) {
	actual := NewGroupIdTeamScheduleTimesOffID("groupId", "timeOffId").ID()
	expected := "/groups/groupId/team/schedule/timesOff/timeOffId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamScheduleTimesOffID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleTimesOffId
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
			Input: "/groups/groupId/team/schedule/timesOff",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/schedule/timesOff/timeOffId",
			Expected: &GroupIdTeamScheduleTimesOffId{
				GroupId:   "groupId",
				TimeOffId: "timeOffId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/schedule/timesOff/timeOffId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleTimesOffID(v.Input)
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

		if actual.TimeOffId != v.Expected.TimeOffId {
			t.Fatalf("Expected %q but got %q for TimeOffId", v.Expected.TimeOffId, actual.TimeOffId)
		}

	}
}

func TestParseGroupIdTeamScheduleTimesOffIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleTimesOffId
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
			Input: "/groups/groupId/team/schedule/timesOff",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/tImEsOfF",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/schedule/timesOff/timeOffId",
			Expected: &GroupIdTeamScheduleTimesOffId{
				GroupId:   "groupId",
				TimeOffId: "timeOffId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/schedule/timesOff/timeOffId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/tImEsOfF/tImEoFfId",
			Expected: &GroupIdTeamScheduleTimesOffId{
				GroupId:   "gRoUpId",
				TimeOffId: "tImEoFfId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/tImEsOfF/tImEoFfId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleTimesOffIDInsensitively(v.Input)
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

		if actual.TimeOffId != v.Expected.TimeOffId {
			t.Fatalf("Expected %q but got %q for TimeOffId", v.Expected.TimeOffId, actual.TimeOffId)
		}

	}
}

func TestSegmentsForGroupIdTeamScheduleTimesOffId(t *testing.T) {
	segments := GroupIdTeamScheduleTimesOffId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamScheduleTimesOffId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
