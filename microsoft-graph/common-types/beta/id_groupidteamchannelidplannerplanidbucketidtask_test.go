package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamChannelIdPlannerPlanIdBucketIdTaskId{}

func TestNewGroupIdTeamChannelIdPlannerPlanIdBucketIdTaskID(t *testing.T) {
	id := NewGroupIdTeamChannelIdPlannerPlanIdBucketIdTaskID("groupId", "channelId", "plannerPlanId", "plannerBucketId", "plannerTaskId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.ChannelId != "channelId" {
		t.Fatalf("Expected %q but got %q for Segment 'ChannelId'", id.ChannelId, "channelId")
	}

	if id.PlannerPlanId != "plannerPlanId" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerPlanId'", id.PlannerPlanId, "plannerPlanId")
	}

	if id.PlannerBucketId != "plannerBucketId" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerBucketId'", id.PlannerBucketId, "plannerBucketId")
	}

	if id.PlannerTaskId != "plannerTaskId" {
		t.Fatalf("Expected %q but got %q for Segment 'PlannerTaskId'", id.PlannerTaskId, "plannerTaskId")
	}
}

func TestFormatGroupIdTeamChannelIdPlannerPlanIdBucketIdTaskID(t *testing.T) {
	actual := NewGroupIdTeamChannelIdPlannerPlanIdBucketIdTaskID("groupId", "channelId", "plannerPlanId", "plannerBucketId", "plannerTaskId").ID()
	expected := "/groups/groupId/team/channels/channelId/planner/plans/plannerPlanId/buckets/plannerBucketId/tasks/plannerTaskId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamChannelIdPlannerPlanIdBucketIdTaskID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamChannelIdPlannerPlanIdBucketIdTaskId
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
			Input: "/groups/groupId/team/channels",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/planner",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/planner/plans",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/planner/plans/plannerPlanId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/planner/plans/plannerPlanId/buckets",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/planner/plans/plannerPlanId/buckets/plannerBucketId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/planner/plans/plannerPlanId/buckets/plannerBucketId/tasks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/channels/channelId/planner/plans/plannerPlanId/buckets/plannerBucketId/tasks/plannerTaskId",
			Expected: &GroupIdTeamChannelIdPlannerPlanIdBucketIdTaskId{
				GroupId:         "groupId",
				ChannelId:       "channelId",
				PlannerPlanId:   "plannerPlanId",
				PlannerBucketId: "plannerBucketId",
				PlannerTaskId:   "plannerTaskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/channels/channelId/planner/plans/plannerPlanId/buckets/plannerBucketId/tasks/plannerTaskId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamChannelIdPlannerPlanIdBucketIdTaskID(v.Input)
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

		if actual.ChannelId != v.Expected.ChannelId {
			t.Fatalf("Expected %q but got %q for ChannelId", v.Expected.ChannelId, actual.ChannelId)
		}

		if actual.PlannerPlanId != v.Expected.PlannerPlanId {
			t.Fatalf("Expected %q but got %q for PlannerPlanId", v.Expected.PlannerPlanId, actual.PlannerPlanId)
		}

		if actual.PlannerBucketId != v.Expected.PlannerBucketId {
			t.Fatalf("Expected %q but got %q for PlannerBucketId", v.Expected.PlannerBucketId, actual.PlannerBucketId)
		}

		if actual.PlannerTaskId != v.Expected.PlannerTaskId {
			t.Fatalf("Expected %q but got %q for PlannerTaskId", v.Expected.PlannerTaskId, actual.PlannerTaskId)
		}

	}
}

func TestParseGroupIdTeamChannelIdPlannerPlanIdBucketIdTaskIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamChannelIdPlannerPlanIdBucketIdTaskId
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
			Input: "/groups/groupId/team/channels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/planner",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/pLaNnEr",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/planner/plans",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/pLaNnEr/pLaNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/planner/plans/plannerPlanId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/pLaNnEr/pLaNs/pLaNnErPlAnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/planner/plans/plannerPlanId/buckets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/pLaNnEr/pLaNs/pLaNnErPlAnId/bUcKeTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/planner/plans/plannerPlanId/buckets/plannerBucketId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/pLaNnEr/pLaNs/pLaNnErPlAnId/bUcKeTs/pLaNnErBuCkEtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/team/channels/channelId/planner/plans/plannerPlanId/buckets/plannerBucketId/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/pLaNnEr/pLaNs/pLaNnErPlAnId/bUcKeTs/pLaNnErBuCkEtId/tAsKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/team/channels/channelId/planner/plans/plannerPlanId/buckets/plannerBucketId/tasks/plannerTaskId",
			Expected: &GroupIdTeamChannelIdPlannerPlanIdBucketIdTaskId{
				GroupId:         "groupId",
				ChannelId:       "channelId",
				PlannerPlanId:   "plannerPlanId",
				PlannerBucketId: "plannerBucketId",
				PlannerTaskId:   "plannerTaskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/team/channels/channelId/planner/plans/plannerPlanId/buckets/plannerBucketId/tasks/plannerTaskId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/pLaNnEr/pLaNs/pLaNnErPlAnId/bUcKeTs/pLaNnErBuCkEtId/tAsKs/pLaNnErTaSkId",
			Expected: &GroupIdTeamChannelIdPlannerPlanIdBucketIdTaskId{
				GroupId:         "gRoUpId",
				ChannelId:       "cHaNnElId",
				PlannerPlanId:   "pLaNnErPlAnId",
				PlannerBucketId: "pLaNnErBuCkEtId",
				PlannerTaskId:   "pLaNnErTaSkId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/tEaM/cHaNnElS/cHaNnElId/pLaNnEr/pLaNs/pLaNnErPlAnId/bUcKeTs/pLaNnErBuCkEtId/tAsKs/pLaNnErTaSkId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamChannelIdPlannerPlanIdBucketIdTaskIDInsensitively(v.Input)
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

		if actual.ChannelId != v.Expected.ChannelId {
			t.Fatalf("Expected %q but got %q for ChannelId", v.Expected.ChannelId, actual.ChannelId)
		}

		if actual.PlannerPlanId != v.Expected.PlannerPlanId {
			t.Fatalf("Expected %q but got %q for PlannerPlanId", v.Expected.PlannerPlanId, actual.PlannerPlanId)
		}

		if actual.PlannerBucketId != v.Expected.PlannerBucketId {
			t.Fatalf("Expected %q but got %q for PlannerBucketId", v.Expected.PlannerBucketId, actual.PlannerBucketId)
		}

		if actual.PlannerTaskId != v.Expected.PlannerTaskId {
			t.Fatalf("Expected %q but got %q for PlannerTaskId", v.Expected.PlannerTaskId, actual.PlannerTaskId)
		}

	}
}

func TestSegmentsForGroupIdTeamChannelIdPlannerPlanIdBucketIdTaskId(t *testing.T) {
	segments := GroupIdTeamChannelIdPlannerPlanIdBucketIdTaskId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamChannelIdPlannerPlanIdBucketIdTaskId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
