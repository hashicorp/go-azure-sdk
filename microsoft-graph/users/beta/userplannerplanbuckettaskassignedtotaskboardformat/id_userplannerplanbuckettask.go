package userplannerplanbuckettaskassignedtotaskboardformat

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPlannerPlanBucketTaskId{}

// UserPlannerPlanBucketTaskId is a struct representing the Resource ID for a User Planner Plan Bucket Task
type UserPlannerPlanBucketTaskId struct {
	UserId          string
	PlannerPlanId   string
	PlannerBucketId string
	PlannerTaskId   string
}

// NewUserPlannerPlanBucketTaskID returns a new UserPlannerPlanBucketTaskId struct
func NewUserPlannerPlanBucketTaskID(userId string, plannerPlanId string, plannerBucketId string, plannerTaskId string) UserPlannerPlanBucketTaskId {
	return UserPlannerPlanBucketTaskId{
		UserId:          userId,
		PlannerPlanId:   plannerPlanId,
		PlannerBucketId: plannerBucketId,
		PlannerTaskId:   plannerTaskId,
	}
}

// ParseUserPlannerPlanBucketTaskID parses 'input' into a UserPlannerPlanBucketTaskId
func ParseUserPlannerPlanBucketTaskID(input string) (*UserPlannerPlanBucketTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerPlanBucketTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerPlanBucketTaskId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
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

// ParseUserPlannerPlanBucketTaskIDInsensitively parses 'input' case-insensitively into a UserPlannerPlanBucketTaskId
// note: this method should only be used for API response data and not user input
func ParseUserPlannerPlanBucketTaskIDInsensitively(input string) (*UserPlannerPlanBucketTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerPlanBucketTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerPlanBucketTaskId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
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

// ValidateUserPlannerPlanBucketTaskID checks that 'input' can be parsed as a User Planner Plan Bucket Task ID
func ValidateUserPlannerPlanBucketTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPlannerPlanBucketTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Planner Plan Bucket Task ID
func (id UserPlannerPlanBucketTaskId) ID() string {
	fmtString := "/users/%s/planner/plans/%s/buckets/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerPlanId, id.PlannerBucketId, id.PlannerTaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Planner Plan Bucket Task ID
func (id UserPlannerPlanBucketTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPlanner", "planner", "planner"),
		resourceids.StaticSegment("staticPlans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanIdValue"),
		resourceids.StaticSegment("staticBuckets", "buckets", "buckets"),
		resourceids.UserSpecifiedSegment("plannerBucketId", "plannerBucketIdValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("plannerTaskId", "plannerTaskIdValue"),
	}
}

// String returns a human-readable description of this User Planner Plan Bucket Task ID
func (id UserPlannerPlanBucketTaskId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Bucket: %q", id.PlannerBucketId),
		fmt.Sprintf("Planner Task: %q", id.PlannerTaskId),
	}
	return fmt.Sprintf("User Planner Plan Bucket Task (%s)", strings.Join(components, "\n"))
}
