package groupteamscheduletimecard

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamScheduleTimeCardId{}

func TestNewGroupTeamScheduleTimeCardID(t *testing.T) {
	id := NewGroupTeamScheduleTimeCardID("groupIdValue", "timeCardIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.TimeCardId != "timeCardIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TimeCardId'", id.TimeCardId, "timeCardIdValue")
	}
}

func TestFormatGroupTeamScheduleTimeCardID(t *testing.T) {
	actual := NewGroupTeamScheduleTimeCardID("groupIdValue", "timeCardIdValue").ID()
	expected := "/groups/groupIdValue/team/schedule/timeCards/timeCardIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupTeamScheduleTimeCardID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamScheduleTimeCardId
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
			Input: "/groups/groupIdValue/team/schedule/timeCards",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/schedule/timeCards/timeCardIdValue",
			Expected: &GroupTeamScheduleTimeCardId{
				GroupId:    "groupIdValue",
				TimeCardId: "timeCardIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/schedule/timeCards/timeCardIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamScheduleTimeCardID(v.Input)
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

		if actual.TimeCardId != v.Expected.TimeCardId {
			t.Fatalf("Expected %q but got %q for TimeCardId", v.Expected.TimeCardId, actual.TimeCardId)
		}

	}
}

func TestParseGroupTeamScheduleTimeCardIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamScheduleTimeCardId
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
			Input: "/groups/groupIdValue/team/schedule/timeCards",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/tImEcArDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/schedule/timeCards/timeCardIdValue",
			Expected: &GroupTeamScheduleTimeCardId{
				GroupId:    "groupIdValue",
				TimeCardId: "timeCardIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/schedule/timeCards/timeCardIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/tImEcArDs/tImEcArDiDvAlUe",
			Expected: &GroupTeamScheduleTimeCardId{
				GroupId:    "gRoUpIdVaLuE",
				TimeCardId: "tImEcArDiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/tImEcArDs/tImEcArDiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamScheduleTimeCardIDInsensitively(v.Input)
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

		if actual.TimeCardId != v.Expected.TimeCardId {
			t.Fatalf("Expected %q but got %q for TimeCardId", v.Expected.TimeCardId, actual.TimeCardId)
		}

	}
}

func TestSegmentsForGroupTeamScheduleTimeCardId(t *testing.T) {
	segments := GroupTeamScheduleTimeCardId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupTeamScheduleTimeCardId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
