package groupteamscheduletimeoffrequest

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamScheduleTimeOffRequestId{}

func TestNewGroupTeamScheduleTimeOffRequestID(t *testing.T) {
	id := NewGroupTeamScheduleTimeOffRequestID("groupIdValue", "timeOffRequestIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.TimeOffRequestId != "timeOffRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TimeOffRequestId'", id.TimeOffRequestId, "timeOffRequestIdValue")
	}
}

func TestFormatGroupTeamScheduleTimeOffRequestID(t *testing.T) {
	actual := NewGroupTeamScheduleTimeOffRequestID("groupIdValue", "timeOffRequestIdValue").ID()
	expected := "/groups/groupIdValue/team/schedule/timeOffRequests/timeOffRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupTeamScheduleTimeOffRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamScheduleTimeOffRequestId
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
			Input: "/groups/groupIdValue/team/schedule/timeOffRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/schedule/timeOffRequests/timeOffRequestIdValue",
			Expected: &GroupTeamScheduleTimeOffRequestId{
				GroupId:          "groupIdValue",
				TimeOffRequestId: "timeOffRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/schedule/timeOffRequests/timeOffRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamScheduleTimeOffRequestID(v.Input)
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

func TestParseGroupTeamScheduleTimeOffRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupTeamScheduleTimeOffRequestId
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
			Input: "/groups/groupIdValue/team/schedule/timeOffRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/tImEoFfReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/schedule/timeOffRequests/timeOffRequestIdValue",
			Expected: &GroupTeamScheduleTimeOffRequestId{
				GroupId:          "groupIdValue",
				TimeOffRequestId: "timeOffRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/schedule/timeOffRequests/timeOffRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/tImEoFfReQuEsTs/tImEoFfReQuEsTiDvAlUe",
			Expected: &GroupTeamScheduleTimeOffRequestId{
				GroupId:          "gRoUpIdVaLuE",
				TimeOffRequestId: "tImEoFfReQuEsTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/tImEoFfReQuEsTs/tImEoFfReQuEsTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupTeamScheduleTimeOffRequestIDInsensitively(v.Input)
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

func TestSegmentsForGroupTeamScheduleTimeOffRequestId(t *testing.T) {
	segments := GroupTeamScheduleTimeOffRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupTeamScheduleTimeOffRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
