package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleOpenShiftId{}

func TestNewGroupIdTeamScheduleOpenShiftID(t *testing.T) {
	id := NewGroupIdTeamScheduleOpenShiftID("groupId", "openShiftId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.OpenShiftId != "openShiftId" {
		t.Fatalf("Expected %q but got %q for Segment 'OpenShiftId'", id.OpenShiftId, "openShiftId")
	}
}

func TestFormatGroupIdTeamScheduleOpenShiftID(t *testing.T) {
	actual := NewGroupIdTeamScheduleOpenShiftID("groupId", "openShiftId").ID()
	expected := "/groups/groupId/team/schedule/openShifts/openShiftId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamScheduleOpenShiftID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleOpenShiftId
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
			Input: "/groups/groupId/team/schedule/openShifts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/schedule/openShifts/openShiftId",
			Expected: &GroupIdTeamScheduleOpenShiftId{
				GroupId:     "groupId",
				OpenShiftId: "openShiftId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/schedule/openShifts/openShiftId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleOpenShiftID(v.Input)
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

		if actual.OpenShiftId != v.Expected.OpenShiftId {
			t.Fatalf("Expected %q but got %q for OpenShiftId", v.Expected.OpenShiftId, actual.OpenShiftId)
		}

	}
}

func TestParseGroupIdTeamScheduleOpenShiftIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleOpenShiftId
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
			Input: "/groups/groupId/team/schedule/openShifts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/oPeNsHiFtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/schedule/openShifts/openShiftId",
			Expected: &GroupIdTeamScheduleOpenShiftId{
				GroupId:     "groupId",
				OpenShiftId: "openShiftId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/schedule/openShifts/openShiftId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/oPeNsHiFtS/oPeNsHiFtId",
			Expected: &GroupIdTeamScheduleOpenShiftId{
				GroupId:     "gRoUpId",
				OpenShiftId: "oPeNsHiFtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/oPeNsHiFtS/oPeNsHiFtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleOpenShiftIDInsensitively(v.Input)
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

		if actual.OpenShiftId != v.Expected.OpenShiftId {
			t.Fatalf("Expected %q but got %q for OpenShiftId", v.Expected.OpenShiftId, actual.OpenShiftId)
		}

	}
}

func TestSegmentsForGroupIdTeamScheduleOpenShiftId(t *testing.T) {
	segments := GroupIdTeamScheduleOpenShiftId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamScheduleOpenShiftId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
