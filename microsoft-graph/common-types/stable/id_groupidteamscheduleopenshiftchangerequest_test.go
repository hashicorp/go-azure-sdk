package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleOpenShiftChangeRequestId{}

func TestNewGroupIdTeamScheduleOpenShiftChangeRequestID(t *testing.T) {
	id := NewGroupIdTeamScheduleOpenShiftChangeRequestID("groupId", "openShiftChangeRequestId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.OpenShiftChangeRequestId != "openShiftChangeRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'OpenShiftChangeRequestId'", id.OpenShiftChangeRequestId, "openShiftChangeRequestId")
	}
}

func TestFormatGroupIdTeamScheduleOpenShiftChangeRequestID(t *testing.T) {
	actual := NewGroupIdTeamScheduleOpenShiftChangeRequestID("groupId", "openShiftChangeRequestId").ID()
	expected := "/groups/groupId/team/schedule/openShiftChangeRequests/openShiftChangeRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamScheduleOpenShiftChangeRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleOpenShiftChangeRequestId
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
			Input: "/groups/groupId/team/schedule/openShiftChangeRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/schedule/openShiftChangeRequests/openShiftChangeRequestId",
			Expected: &GroupIdTeamScheduleOpenShiftChangeRequestId{
				GroupId:                  "groupId",
				OpenShiftChangeRequestId: "openShiftChangeRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/schedule/openShiftChangeRequests/openShiftChangeRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleOpenShiftChangeRequestID(v.Input)
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

		if actual.OpenShiftChangeRequestId != v.Expected.OpenShiftChangeRequestId {
			t.Fatalf("Expected %q but got %q for OpenShiftChangeRequestId", v.Expected.OpenShiftChangeRequestId, actual.OpenShiftChangeRequestId)
		}

	}
}

func TestParseGroupIdTeamScheduleOpenShiftChangeRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleOpenShiftChangeRequestId
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
			Input: "/groups/groupId/team/schedule/openShiftChangeRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/oPeNsHiFtChAnGeReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/schedule/openShiftChangeRequests/openShiftChangeRequestId",
			Expected: &GroupIdTeamScheduleOpenShiftChangeRequestId{
				GroupId:                  "groupId",
				OpenShiftChangeRequestId: "openShiftChangeRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/schedule/openShiftChangeRequests/openShiftChangeRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/oPeNsHiFtChAnGeReQuEsTs/oPeNsHiFtChAnGeReQuEsTiD",
			Expected: &GroupIdTeamScheduleOpenShiftChangeRequestId{
				GroupId:                  "gRoUpId",
				OpenShiftChangeRequestId: "oPeNsHiFtChAnGeReQuEsTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/sChEdUlE/oPeNsHiFtChAnGeReQuEsTs/oPeNsHiFtChAnGeReQuEsTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleOpenShiftChangeRequestIDInsensitively(v.Input)
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

		if actual.OpenShiftChangeRequestId != v.Expected.OpenShiftChangeRequestId {
			t.Fatalf("Expected %q but got %q for OpenShiftChangeRequestId", v.Expected.OpenShiftChangeRequestId, actual.OpenShiftChangeRequestId)
		}

	}
}

func TestSegmentsForGroupIdTeamScheduleOpenShiftChangeRequestId(t *testing.T) {
	segments := GroupIdTeamScheduleOpenShiftChangeRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamScheduleOpenShiftChangeRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
