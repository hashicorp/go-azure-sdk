package userplannerfavoriteplan

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPlannerFavoritePlanId{}

// UserPlannerFavoritePlanId is a struct representing the Resource ID for a User Planner Favorite Plan
type UserPlannerFavoritePlanId struct {
	UserId        string
	PlannerPlanId string
}

// NewUserPlannerFavoritePlanID returns a new UserPlannerFavoritePlanId struct
func NewUserPlannerFavoritePlanID(userId string, plannerPlanId string) UserPlannerFavoritePlanId {
	return UserPlannerFavoritePlanId{
		UserId:        userId,
		PlannerPlanId: plannerPlanId,
	}
}

// ParseUserPlannerFavoritePlanID parses 'input' into a UserPlannerFavoritePlanId
func ParseUserPlannerFavoritePlanID(input string) (*UserPlannerFavoritePlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerFavoritePlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerFavoritePlanId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	return &id, nil
}

// ParseUserPlannerFavoritePlanIDInsensitively parses 'input' case-insensitively into a UserPlannerFavoritePlanId
// note: this method should only be used for API response data and not user input
func ParseUserPlannerFavoritePlanIDInsensitively(input string) (*UserPlannerFavoritePlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerFavoritePlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerFavoritePlanId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	return &id, nil
}

// ValidateUserPlannerFavoritePlanID checks that 'input' can be parsed as a User Planner Favorite Plan ID
func ValidateUserPlannerFavoritePlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPlannerFavoritePlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Planner Favorite Plan ID
func (id UserPlannerFavoritePlanId) ID() string {
	fmtString := "/users/%s/planner/favoritePlans/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Planner Favorite Plan ID
func (id UserPlannerFavoritePlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPlanner", "planner", "planner"),
		resourceids.StaticSegment("staticFavoritePlans", "favoritePlans", "favoritePlans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanIdValue"),
	}
}

// String returns a human-readable description of this User Planner Favorite Plan ID
func (id UserPlannerFavoritePlanId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
	}
	return fmt.Sprintf("User Planner Favorite Plan (%s)", strings.Join(components, "\n"))
}
