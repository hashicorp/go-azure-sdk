package meplannerplanbuckettaskdetail

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePlannerPlanBucketTaskId{}

// MePlannerPlanBucketTaskId is a struct representing the Resource ID for a Me Planner Plan Bucket Task
type MePlannerPlanBucketTaskId struct {
	PlannerPlanId   string
	PlannerBucketId string
	PlannerTaskId   string
}

// NewMePlannerPlanBucketTaskID returns a new MePlannerPlanBucketTaskId struct
func NewMePlannerPlanBucketTaskID(plannerPlanId string, plannerBucketId string, plannerTaskId string) MePlannerPlanBucketTaskId {
	return MePlannerPlanBucketTaskId{
		PlannerPlanId:   plannerPlanId,
		PlannerBucketId: plannerBucketId,
		PlannerTaskId:   plannerTaskId,
	}
}

// ParseMePlannerPlanBucketTaskID parses 'input' into a MePlannerPlanBucketTaskId
func ParseMePlannerPlanBucketTaskID(input string) (*MePlannerPlanBucketTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePlannerPlanBucketTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePlannerPlanBucketTaskId{}

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

// ParseMePlannerPlanBucketTaskIDInsensitively parses 'input' case-insensitively into a MePlannerPlanBucketTaskId
// note: this method should only be used for API response data and not user input
func ParseMePlannerPlanBucketTaskIDInsensitively(input string) (*MePlannerPlanBucketTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePlannerPlanBucketTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePlannerPlanBucketTaskId{}

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

// ValidateMePlannerPlanBucketTaskID checks that 'input' can be parsed as a Me Planner Plan Bucket Task ID
func ValidateMePlannerPlanBucketTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePlannerPlanBucketTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Planner Plan Bucket Task ID
func (id MePlannerPlanBucketTaskId) ID() string {
	fmtString := "/me/planner/plans/%s/buckets/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.PlannerPlanId, id.PlannerBucketId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Planner Plan Bucket Task ID
func (id MePlannerPlanBucketTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticPlanner", "planner", "planner"),
		resourceids.StaticSegment("staticPlans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanIdValue"),
		resourceids.StaticSegment("staticBuckets", "buckets", "buckets"),
		resourceids.UserSpecifiedSegment("plannerBucketId", "plannerBucketIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskIdValue"),
	}
}

// String returns a human-readable description of this Me Planner Plan Bucket Task ID
func (id MePlannerPlanBucketTaskId) String() string {
	components := []string{
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Bucket: %q", id.PlannerBucketId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("Me Planner Plan Bucket Task (%s)", strings.Join(components, "\n"))
}
