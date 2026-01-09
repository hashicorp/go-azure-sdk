package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleSwapShiftsChangeRequestId{}

func TestNewGroupIdTeamScheduleSwapShiftsChangeRequestID(t *testing.T) {
	id := NewGroupIdTeamScheduleSwapShiftsChangeRequestID("groupId", "swapShiftsChangeRequestId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.SwapShiftsChangeRequestId != "swapShiftsChangeRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'SwapShiftsChangeRequestId'", id.SwapShiftsChangeRequestId, "swapShiftsChangeRequestId")
	}
}

func TestFormatGroupIdTeamScheduleSwapShiftsChangeRequestID(t *testing.T) {
	actual := NewGroupIdTeamScheduleSwapShiftsChangeRequestID("groupId", "swapShiftsChangeRequestId").ID()
	expected := "/groups/groupId/team/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamScheduleSwapShiftsChangeRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleSwapShiftsChangeRequestId
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
			Input: "/groups/groupId/team/schedule/swapShiftsChangeRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestId",
			Expected: &GroupIdTeamScheduleSwapShiftsChangeRequestId{
				GroupId:                   "groupId",
				SwapShiftsChangeRequestId: "swapShiftsChangeRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleSwapShiftsChangeRequestID(v.Input)
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

		if actual.SwapShiftsChangeRequestId != v.Expected.SwapShiftsChangeRequestId {
			t.Fatalf("Expected %q but got %q for SwapShiftsChangeRequestId", v.Expected.SwapShiftsChangeRequestId, actual.SwapShiftsChangeRequestId)
		}

	}
}

func TestParseGroupIdTeamScheduleSwapShiftsChangeRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleSwapShiftsChangeRequestId
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
			Input: "/groups/groupId/team/schedule/swapShiftsChangeRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/sWaPsHiFtScHaNgErEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestId",
			Expected: &GroupIdTeamScheduleSwapShiftsChangeRequestId{
				GroupId:                   "groupId",
				SwapShiftsChangeRequestId: "swapShiftsChangeRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/schedule/swapShiftsChangeRequests/swapShiftsChangeRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/sWaPsHiFtScHaNgErEqUeStS/sWaPsHiFtScHaNgErEqUeStId",
			Expected: &GroupIdTeamScheduleSwapShiftsChangeRequestId{
				GroupId:                   "gRoUpId",
				SwapShiftsChangeRequestId: "sWaPsHiFtScHaNgErEqUeStId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/sWaPsHiFtScHaNgErEqUeStS/sWaPsHiFtScHaNgErEqUeStId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleSwapShiftsChangeRequestIDInsensitively(v.Input)
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

		if actual.SwapShiftsChangeRequestId != v.Expected.SwapShiftsChangeRequestId {
			t.Fatalf("Expected %q but got %q for SwapShiftsChangeRequestId", v.Expected.SwapShiftsChangeRequestId, actual.SwapShiftsChangeRequestId)
		}

	}
}

func TestSegmentsForGroupIdTeamScheduleSwapShiftsChangeRequestId(t *testing.T) {
	segments := GroupIdTeamScheduleSwapShiftsChangeRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamScheduleSwapShiftsChangeRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
