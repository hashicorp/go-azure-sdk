package userplannerrosterplan

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPlannerRosterPlanId{}

// UserPlannerRosterPlanId is a struct representing the Resource ID for a User Planner Roster Plan
type UserPlannerRosterPlanId struct {
	UserId        string
	PlannerPlanId string
}

// NewUserPlannerRosterPlanID returns a new UserPlannerRosterPlanId struct
func NewUserPlannerRosterPlanID(userId string, plannerPlanId string) UserPlannerRosterPlanId {
	return UserPlannerRosterPlanId{
		UserId:        userId,
		PlannerPlanId: plannerPlanId,
	}
}

// ParseUserPlannerRosterPlanID parses 'input' into a UserPlannerRosterPlanId
func ParseUserPlannerRosterPlanID(input string) (*UserPlannerRosterPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerRosterPlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerRosterPlanId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	return &id, nil
}

// ParseUserPlannerRosterPlanIDInsensitively parses 'input' case-insensitively into a UserPlannerRosterPlanId
// note: this method should only be used for API response data and not user input
func ParseUserPlannerRosterPlanIDInsensitively(input string) (*UserPlannerRosterPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerRosterPlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerRosterPlanId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	return &id, nil
}

// ValidateUserPlannerRosterPlanID checks that 'input' can be parsed as a User Planner Roster Plan ID
func ValidateUserPlannerRosterPlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPlannerRosterPlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Planner Roster Plan ID
func (id UserPlannerRosterPlanId) ID() string {
	fmtString := "/users/%s/planner/rosterPlans/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Planner Roster Plan ID
func (id UserPlannerRosterPlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPlanner", "planner", "planner"),
		resourceids.StaticSegment("staticRosterPlans", "rosterPlans", "rosterPlans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanIdValue"),
	}
}

// String returns a human-readable description of this User Planner Roster Plan ID
func (id UserPlannerRosterPlanId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
	}
	return fmt.Sprintf("User Planner Roster Plan (%s)", strings.Join(components, "\n"))
}
