package groupplannerplantask

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupPlannerPlanTaskId{}

// GroupPlannerPlanTaskId is a struct representing the Resource ID for a Group Planner Plan Task
type GroupPlannerPlanTaskId struct {
	GroupId       string
	PlannerPlanId string
	PlannerTaskId string
}

// NewGroupPlannerPlanTaskID returns a new GroupPlannerPlanTaskId struct
func NewGroupPlannerPlanTaskID(groupId string, plannerPlanId string, plannerTaskId string) GroupPlannerPlanTaskId {
	return GroupPlannerPlanTaskId{
		GroupId:       groupId,
		PlannerPlanId: plannerPlanId,
		PlannerTaskId: plannerTaskId,
	}
}

// ParseGroupPlannerPlanTaskID parses 'input' into a GroupPlannerPlanTaskId
func ParseGroupPlannerPlanTaskID(input string) (*GroupPlannerPlanTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupPlannerPlanTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupPlannerPlanTaskId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	if id.PlannerTaskId, ok = parsed.Parsed["plannerTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", *parsed)
	}

	return &id, nil
}

// ParseGroupPlannerPlanTaskIDInsensitively parses 'input' case-insensitively into a GroupPlannerPlanTaskId
// note: this method should only be used for API response data and not user input
func ParseGroupPlannerPlanTaskIDInsensitively(input string) (*GroupPlannerPlanTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupPlannerPlanTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupPlannerPlanTaskId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	if id.PlannerTaskId, ok = parsed.Parsed["plannerTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", *parsed)
	}

	return &id, nil
}

// ValidateGroupPlannerPlanTaskID checks that 'input' can be parsed as a Group Planner Plan Task ID
func ValidateGroupPlannerPlanTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupPlannerPlanTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Planner Plan Task ID
func (id GroupPlannerPlanTaskId) ID() string {
	fmtString := "/groups/%s/planner/plans/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.PlannerPlanId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Planner Plan Task ID
func (id GroupPlannerPlanTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticPlanner", "planner", "planner"),
		resourceids.StaticSegment("staticPlans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskIdValue"),
	}
}

// String returns a human-readable description of this Group Planner Plan Task ID
func (id GroupPlannerPlanTaskId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("Group Planner Plan Task (%s)", strings.Join(components, "\n"))
}
