package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleTimeOffRequestId{}

func TestNewGroupIdTeamScheduleTimeOffRequestID(t *testing.T) {
	id := NewGroupIdTeamScheduleTimeOffRequestID("groupId", "timeOffRequestId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.TimeOffRequestId != "timeOffRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'TimeOffRequestId'", id.TimeOffRequestId, "timeOffRequestId")
	}
}

func TestFormatGroupIdTeamScheduleTimeOffRequestID(t *testing.T) {
	actual := NewGroupIdTeamScheduleTimeOffRequestID("groupId", "timeOffRequestId").ID()
	expected := "/groups/groupId/team/schedule/timeOffRequests/timeOffRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamScheduleTimeOffRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleTimeOffRequestId
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
			Input: "/groups/groupId/team/schedule/timeOffRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/schedule/timeOffRequests/timeOffRequestId",
			Expected: &GroupIdTeamScheduleTimeOffRequestId{
				GroupId:          "groupId",
				TimeOffRequestId: "timeOffRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/schedule/timeOffRequests/timeOffRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleTimeOffRequestID(v.Input)
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

		if actual.TimeOffRequestId != v.Expected.TimeOffRequestId {
			t.Fatalf("Expected %q but got %q for TimeOffRequestId", v.Expected.TimeOffRequestId, actual.TimeOffRequestId)
		}

	}
}

func TestParseGroupIdTeamScheduleTimeOffRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleTimeOffRequestId
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
			Input: "/groups/groupId/team/schedule/timeOffRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/tImEoFfReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/schedule/timeOffRequests/timeOffRequestId",
			Expected: &GroupIdTeamScheduleTimeOffRequestId{
				GroupId:          "groupId",
				TimeOffRequestId: "timeOffRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/schedule/timeOffRequests/timeOffRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/tImEoFfReQuEsTs/tImEoFfReQuEsTiD",
			Expected: &GroupIdTeamScheduleTimeOffRequestId{
				GroupId:          "gRoUpId",
				TimeOffRequestId: "tImEoFfReQuEsTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/tImEoFfReQuEsTs/tImEoFfReQuEsTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleTimeOffRequestIDInsensitively(v.Input)
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

		if actual.TimeOffRequestId != v.Expected.TimeOffRequestId {
			t.Fatalf("Expected %q but got %q for TimeOffRequestId", v.Expected.TimeOffRequestId, actual.TimeOffRequestId)
		}

	}
}

func TestSegmentsForGroupIdTeamScheduleTimeOffRequestId(t *testing.T) {
	segments := GroupIdTeamScheduleTimeOffRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamScheduleTimeOffRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
