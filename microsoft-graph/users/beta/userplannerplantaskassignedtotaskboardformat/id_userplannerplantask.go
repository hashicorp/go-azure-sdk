package userplannerplantaskassignedtotaskboardformat

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPlannerPlanTaskId{}

// UserPlannerPlanTaskId is a struct representing the Resource ID for a User Planner Plan Task
type UserPlannerPlanTaskId struct {
	UserId        string
	PlannerPlanId string
	PlannerTaskId string
}

// NewUserPlannerPlanTaskID returns a new UserPlannerPlanTaskId struct
func NewUserPlannerPlanTaskID(userId string, plannerPlanId string, plannerTaskId string) UserPlannerPlanTaskId {
	return UserPlannerPlanTaskId{
		UserId:        userId,
		PlannerPlanId: plannerPlanId,
		PlannerTaskId: plannerTaskId,
	}
}

// ParseUserPlannerPlanTaskID parses 'input' into a UserPlannerPlanTaskId
func ParseUserPlannerPlanTaskID(input string) (*UserPlannerPlanTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerPlanTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerPlanTaskId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	if id.PlannerTaskId, ok = parsed.Parsed["plannerTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", *parsed)
	}

	return &id, nil
}

// ParseUserPlannerPlanTaskIDInsensitively parses 'input' case-insensitively into a UserPlannerPlanTaskId
// note: this method should only be used for API response data and not user input
func ParseUserPlannerPlanTaskIDInsensitively(input string) (*UserPlannerPlanTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerPlanTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerPlanTaskId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	if id.PlannerTaskId, ok = parsed.Parsed["plannerTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", *parsed)
	}

	return &id, nil
}

// ValidateUserPlannerPlanTaskID checks that 'input' can be parsed as a User Planner Plan Task ID
func ValidateUserPlannerPlanTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPlannerPlanTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Planner Plan Task ID
func (id UserPlannerPlanTaskId) ID() string {
	fmtString := "/users/%s/planner/plans/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerPlanId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Planner Plan Task ID
func (id UserPlannerPlanTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPlanner", "planner", "planner"),
		resourceids.StaticSegment("staticPlans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskIdValue"),
	}
}

// String returns a human-readable description of this User Planner Plan Task ID
func (id UserPlannerPlanTaskId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("User Planner Plan Task (%s)", strings.Join(components, "\n"))
}
