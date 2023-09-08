package groupteamscheduletimesoff

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamScheduleTimesOffId{}

func TestNewGroupTeamScheduleTimesOffID(t *testing.T) {
	id := NewGroupTeamScheduleTimesOffID("groupIdValue", "timeOffIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.TimeOffId != "timeOffIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TimeOffId'", id.TimeOffId, "timeOffIdValue")
	}
}

func TestFormatGroupTeamScheduleTimesOffID(t *testing.T) {
	actual := NewGroupTeamScheduleTimesOffID("groupIdValue", "timeOffIdValue").ID()
	expected := "/groups/groupIdValue/team/schedule/timesOff/timeOffIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupTeamScheduleTimesOffID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamScheduleTimesOffId
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
			Input: "/groups/groupIdValue/team/schedule/timesOff",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/schedule/timesOff/timeOffIdValue",
			Expected: &GroupTeamScheduleTimesOffId{
				GroupId:   "groupIdValue",
				TimeOffId: "timeOffIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/schedule/timesOff/timeOffIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamScheduleTimesOffID(v.Input)
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

func TestParseGroupTeamScheduleTimesOffIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamScheduleTimesOffId
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
			Input: "/groups/groupIdValue/team/schedule/timesOff",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/tImEsOfF",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/schedule/timesOff/timeOffIdValue",
			Expected: &GroupTeamScheduleTimesOffId{
				GroupId:   "groupIdValue",
				TimeOffId: "timeOffIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/schedule/timesOff/timeOffIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/tImEsOfF/tImEoFfIdVaLuE",
			Expected: &GroupTeamScheduleTimesOffId{
				GroupId:   "gRoUpIdVaLuE",
				TimeOffId: "tImEoFfIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/tImEsOfF/tImEoFfIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamScheduleTimesOffIDInsensitively(v.Input)
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

func TestSegmentsForGroupTeamScheduleTimesOffId(t *testing.T) {
	segments := GroupTeamScheduleTimesOffId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupTeamScheduleTimesOffId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
