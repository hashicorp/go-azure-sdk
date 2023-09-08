package groupplannerplanbuckettaskassignedtotaskboardformat

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupPlannerPlanBucketTaskId{}

// GroupPlannerPlanBucketTaskId is a struct representing the Resource ID for a Group Planner Plan Bucket Task
type GroupPlannerPlanBucketTaskId struct {
	GroupId         string
	PlannerPlanId   string
	PlannerBucketId string
	PlannerTaskId   string
}

// NewGroupPlannerPlanBucketTaskID returns a new GroupPlannerPlanBucketTaskId struct
func NewGroupPlannerPlanBucketTaskID(groupId string, plannerPlanId string, plannerBucketId string, plannerTaskId string) GroupPlannerPlanBucketTaskId {
	return GroupPlannerPlanBucketTaskId{
		GroupId:         groupId,
		PlannerPlanId:   plannerPlanId,
		PlannerBucketId: plannerBucketId,
		PlannerTaskId:   plannerTaskId,
	}
}

// ParseGroupPlannerPlanBucketTaskID parses 'input' into a GroupPlannerPlanBucketTaskId
func ParseGroupPlannerPlanBucketTaskID(input string) (*GroupPlannerPlanBucketTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupPlannerPlanBucketTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupPlannerPlanBucketTaskId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	if id.PlannerBucketId, ok = parsed.Parsed["plannerBucketId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerBucketId", *parsed)
	}

	if id.PlannerTaskId, ok = parsed.Parsed["plannerTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", *parsed)
	}

	return &id, nil
}

// ParseGroupPlannerPlanBucketTaskIDInsensitively parses 'input' case-insensitively into a GroupPlannerPlanBucketTaskId
// note: this method should only be used for API response data and not user input
func ParseGroupPlannerPlanBucketTaskIDInsensitively(input string) (*GroupPlannerPlanBucketTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupPlannerPlanBucketTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupPlannerPlanBucketTaskId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	if id.PlannerBucketId, ok = parsed.Parsed["plannerBucketId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerBucketId", *parsed)
	}

	if id.PlannerTaskId, ok = parsed.Parsed["plannerTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", *parsed)
	}

	return &id, nil
}

// ValidateGroupPlannerPlanBucketTaskID checks that 'input' can be parsed as a Group Planner Plan Bucket Task ID
func ValidateGroupPlannerPlanBucketTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupPlannerPlanBucketTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Planner Plan Bucket Task ID
func (id GroupPlannerPlanBucketTaskId) ID() string {
	fmtString := "/groups/%s/planner/plans/%s/buckets/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.PlannerPlanId, id.PlannerBucketId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Planner Plan Bucket Task ID
func (id GroupPlannerPlanBucketTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticPlanner", "planner", "planner"),
		resourceids.StaticSegment("staticPlans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanIdValue"),
		resourceids.StaticSegment("staticBuckets", "buckets", "buckets"),
		resourceids.UserSpecifiedSegment("plannerBucketId", "plannerBucketIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskIdValue"),
	}
}

// String returns a human-readable description of this Group Planner Plan Bucket Task ID
func (id GroupPlannerPlanBucketTaskId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Bucket: %q", id.PlannerBucketId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("Group Planner Plan Bucket Task (%s)", strings.Join(components, "\n"))
}
