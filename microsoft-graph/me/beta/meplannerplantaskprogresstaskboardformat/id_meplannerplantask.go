package meplannerplantaskprogresstaskboardformat

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePlannerPlanTaskId{}

// MePlannerPlanTaskId is a struct representing the Resource ID for a Me Planner Plan Task
type MePlannerPlanTaskId struct {
	PlannerPlanId string
	PlannerTaskId string
}

// NewMePlannerPlanTaskID returns a new MePlannerPlanTaskId struct
func NewMePlannerPlanTaskID(plannerPlanId string, plannerTaskId string) MePlannerPlanTaskId {
	return MePlannerPlanTaskId{
		PlannerPlanId: plannerPlanId,
		PlannerTaskId: plannerTaskId,
	}
}

// ParseMePlannerPlanTaskID parses 'input' into a MePlannerPlanTaskId
func ParseMePlannerPlanTaskID(input string) (*MePlannerPlanTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePlannerPlanTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePlannerPlanTaskId{}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	if id.PlannerTaskId, ok = parsed.Parsed["plannerTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", *parsed)
	}

	return &id, nil
}

// ParseMePlannerPlanTaskIDInsensitively parses 'input' case-insensitively into a MePlannerPlanTaskId
// note: this method should only be used for API response data and not user input
func ParseMePlannerPlanTaskIDInsensitively(input string) (*MePlannerPlanTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePlannerPlanTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePlannerPlanTaskId{}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	if id.PlannerTaskId, ok = parsed.Parsed["plannerTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", *parsed)
	}

	return &id, nil
}

// ValidateMePlannerPlanTaskID checks that 'input' can be parsed as a Me Planner Plan Task ID
func ValidateMePlannerPlanTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePlannerPlanTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Planner Plan Task ID
func (id MePlannerPlanTaskId) ID() string {
	fmtString := "/me/planner/plans/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.PlannerPlanId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Planner Plan Task ID
func (id MePlannerPlanTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticPlanner", "planner", "planner"),
		resourceids.StaticSegment("staticPlans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskIdValue"),
	}
}

// String returns a human-readable description of this Me Planner Plan Task ID
func (id MePlannerPlanTaskId) String() string {
	components := []string{
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("Me Planner Plan Task (%s)", strings.Join(components, "\n"))
}
