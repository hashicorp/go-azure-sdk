package userplannertaskassignedtotaskboardformat

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPlannerTaskId{}

// UserPlannerTaskId is a struct representing the Resource ID for a User Planner Task
type UserPlannerTaskId struct {
	UserId        string
	PlannerTaskId string
}

// NewUserPlannerTaskID returns a new UserPlannerTaskId struct
func NewUserPlannerTaskID(userId string, plannerTaskId string) UserPlannerTaskId {
	return UserPlannerTaskId{
		UserId:        userId,
		PlannerTaskId: plannerTaskId,
	}
}

// ParseUserPlannerTaskID parses 'input' into a UserPlannerTaskId
func ParseUserPlannerTaskID(input string) (*UserPlannerTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerTaskId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PlannerTaskId, ok = parsed.Parsed["plannerTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", *parsed)
	}

	return &id, nil
}

// ParseUserPlannerTaskIDInsensitively parses 'input' case-insensitively into a UserPlannerTaskId
// note: this method should only be used for API response data and not user input
func ParseUserPlannerTaskIDInsensitively(input string) (*UserPlannerTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerTaskId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PlannerTaskId, ok = parsed.Parsed["plannerTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerTaskId", *parsed)
	}

	return &id, nil
}

// ValidateUserPlannerTaskID checks that 'input' can be parsed as a User Planner Task ID
func ValidateUserPlannerTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPlannerTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Planner Task ID
func (id UserPlannerTaskId) ID() string {
	fmtString := "/users/%s/planner/tasks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Planner Task ID
func (id UserPlannerTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPlanner", "planner", "planner"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskIdValue"),
	}
}

// String returns a human-readable description of this User Planner Task ID
func (id UserPlannerTaskId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("User Planner Task (%s)", strings.Join(components, "\n"))
}
