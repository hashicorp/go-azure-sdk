package groupteamscheduleshift

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamScheduleShiftId{}

func TestNewGroupTeamScheduleShiftID(t *testing.T) {
	id := NewGroupTeamScheduleShiftID("groupIdValue", "shiftIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.ShiftId != "shiftIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ShiftId'", id.ShiftId, "shiftIdValue")
	}
}

func TestFormatGroupTeamScheduleShiftID(t *testing.T) {
	actual := NewGroupTeamScheduleShiftID("groupIdValue", "shiftIdValue").ID()
	expected := "/groups/groupIdValue/team/schedule/shifts/shiftIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupTeamScheduleShiftID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamScheduleShiftId
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
			Input: "/groups/groupIdValue/team/schedule/shifts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/schedule/shifts/shiftIdValue",
			Expected: &GroupTeamScheduleShiftId{
				GroupId: "groupIdValue",
				ShiftId: "shiftIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/schedule/shifts/shiftIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamScheduleShiftID(v.Input)
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

		if actual.ShiftId != v.Expected.ShiftId {
			t.Fatalf("Expected %q but got %q for ShiftId", v.Expected.ShiftId, actual.ShiftId)
		}

	}
}

func TestParseGroupTeamScheduleShiftIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamScheduleShiftId
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
			Input: "/groups/groupIdValue/team/schedule/shifts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/sHiFtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/schedule/shifts/shiftIdValue",
			Expected: &GroupTeamScheduleShiftId{
				GroupId: "groupIdValue",
				ShiftId: "shiftIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/schedule/shifts/shiftIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/sHiFtS/sHiFtIdVaLuE",
			Expected: &GroupTeamScheduleShiftId{
				GroupId: "gRoUpIdVaLuE",
				ShiftId: "sHiFtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/sHiFtS/sHiFtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamScheduleShiftIDInsensitively(v.Input)
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

		if actual.ShiftId != v.Expected.ShiftId {
			t.Fatalf("Expected %q but got %q for ShiftId", v.Expected.ShiftId, actual.ShiftId)
		}

	}
}

func TestSegmentsForGroupTeamScheduleShiftId(t *testing.T) {
	segments := GroupTeamScheduleShiftId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupTeamScheduleShiftId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
