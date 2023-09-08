package userplannerrecentplan

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPlannerRecentPlanId{}

// UserPlannerRecentPlanId is a struct representing the Resource ID for a User Planner Recent Plan
type UserPlannerRecentPlanId struct {
	UserId        string
	PlannerPlanId string
}

// NewUserPlannerRecentPlanID returns a new UserPlannerRecentPlanId struct
func NewUserPlannerRecentPlanID(userId string, plannerPlanId string) UserPlannerRecentPlanId {
	return UserPlannerRecentPlanId{
		UserId:        userId,
		PlannerPlanId: plannerPlanId,
	}
}

// ParseUserPlannerRecentPlanID parses 'input' into a UserPlannerRecentPlanId
func ParseUserPlannerRecentPlanID(input string) (*UserPlannerRecentPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerRecentPlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerRecentPlanId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	return &id, nil
}

// ParseUserPlannerRecentPlanIDInsensitively parses 'input' case-insensitively into a UserPlannerRecentPlanId
// note: this method should only be used for API response data and not user input
func ParseUserPlannerRecentPlanIDInsensitively(input string) (*UserPlannerRecentPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerRecentPlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerRecentPlanId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	return &id, nil
}

// ValidateUserPlannerRecentPlanID checks that 'input' can be parsed as a User Planner Recent Plan ID
func ValidateUserPlannerRecentPlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPlannerRecentPlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Planner Recent Plan ID
func (id UserPlannerRecentPlanId) ID() string {
	fmtString := "/users/%s/planner/recentPlans/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Planner Recent Plan ID
func (id UserPlannerRecentPlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPlanner", "planner", "planner"),
		resourceids.StaticSegment("staticRecentPlans", "recentPlans", "recentPlans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanIdValue"),
	}
}

// String returns a human-readable description of this User Planner Recent Plan ID
func (id UserPlannerRecentPlanId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
	}
	return fmt.Sprintf("User Planner Recent Plan (%s)", strings.Join(components, "\n"))
}
