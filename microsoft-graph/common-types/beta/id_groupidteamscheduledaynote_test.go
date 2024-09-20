package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleDayNoteId{}

func TestNewGroupIdTeamScheduleDayNoteID(t *testing.T) {
	id := NewGroupIdTeamScheduleDayNoteID("groupId", "dayNoteId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.DayNoteId != "dayNoteId" {
		t.Fatalf("Expected %q but got %q for Segment 'DayNoteId'", id.DayNoteId, "dayNoteId")
	}
}

func TestFormatGroupIdTeamScheduleDayNoteID(t *testing.T) {
	actual := NewGroupIdTeamScheduleDayNoteID("groupId", "dayNoteId").ID()
	expected := "/groups/groupId/team/schedule/dayNotes/dayNoteId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamScheduleDayNoteID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleDayNoteId
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
			Input: "/groups/groupId/team/schedule/dayNotes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/schedule/dayNotes/dayNoteId",
			Expected: &GroupIdTeamScheduleDayNoteId{
				GroupId:   "groupId",
				DayNoteId: "dayNoteId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/schedule/dayNotes/dayNoteId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleDayNoteID(v.Input)
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

		if actual.DayNoteId != v.Expected.DayNoteId {
			t.Fatalf("Expected %q but got %q for DayNoteId", v.Expected.DayNoteId, actual.DayNoteId)
		}

	}
}

func TestParseGroupIdTeamScheduleDayNoteIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleDayNoteId
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
			Input: "/groups/groupId/team/schedule/dayNotes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/dAyNoTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/schedule/dayNotes/dayNoteId",
			Expected: &GroupIdTeamScheduleDayNoteId{
				GroupId:   "groupId",
				DayNoteId: "dayNoteId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/schedule/dayNotes/dayNoteId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/dAyNoTeS/dAyNoTeId",
			Expected: &GroupIdTeamScheduleDayNoteId{
				GroupId:   "gRoUpId",
				DayNoteId: "dAyNoTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/dAyNoTeS/dAyNoTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleDayNoteIDInsensitively(v.Input)
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

		if actual.DayNoteId != v.Expected.DayNoteId {
			t.Fatalf("Expected %q but got %q for DayNoteId", v.Expected.DayNoteId, actual.DayNoteId)
		}

	}
}

func TestSegmentsForGroupIdTeamScheduleDayNoteId(t *testing.T) {
	segments := GroupIdTeamScheduleDayNoteId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamScheduleDayNoteId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
