package userplannerplanbucket

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPlannerPlanId{}

// UserPlannerPlanId is a struct representing the Resource ID for a User Planner Plan
type UserPlannerPlanId struct {
	UserId        string
	PlannerPlanId string
}

// NewUserPlannerPlanID returns a new UserPlannerPlanId struct
func NewUserPlannerPlanID(userId string, plannerPlanId string) UserPlannerPlanId {
	return UserPlannerPlanId{
		UserId:        userId,
		PlannerPlanId: plannerPlanId,
	}
}

// ParseUserPlannerPlanID parses 'input' into a UserPlannerPlanId
func ParseUserPlannerPlanID(input string) (*UserPlannerPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerPlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerPlanId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	return &id, nil
}

// ParseUserPlannerPlanIDInsensitively parses 'input' case-insensitively into a UserPlannerPlanId
// note: this method should only be used for API response data and not user input
func ParseUserPlannerPlanIDInsensitively(input string) (*UserPlannerPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerPlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerPlanId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	return &id, nil
}

// ValidateUserPlannerPlanID checks that 'input' can be parsed as a User Planner Plan ID
func ValidateUserPlannerPlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPlannerPlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Planner Plan ID
func (id UserPlannerPlanId) ID() string {
	fmtString := "/users/%s/planner/plans/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Planner Plan ID
func (id UserPlannerPlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPlanner", "planner", "planner"),
		resourceids.StaticSegment("staticPlans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanIdValue"),
	}
}

// String returns a human-readable description of this User Planner Plan ID
func (id UserPlannerPlanId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
	}
	return fmt.Sprintf("User Planner Plan (%s)", strings.Join(components, "\n"))
}
