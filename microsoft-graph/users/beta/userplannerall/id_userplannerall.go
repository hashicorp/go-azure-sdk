package userplannerall

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPlannerAllId{}

// UserPlannerAllId is a struct representing the Resource ID for a User Planner All
type UserPlannerAllId struct {
	UserId         string
	PlannerDeltaId string
}

// NewUserPlannerAllID returns a new UserPlannerAllId struct
func NewUserPlannerAllID(userId string, plannerDeltaId string) UserPlannerAllId {
	return UserPlannerAllId{
		UserId:         userId,
		PlannerDeltaId: plannerDeltaId,
	}
}

// ParseUserPlannerAllID parses 'input' into a UserPlannerAllId
func ParseUserPlannerAllID(input string) (*UserPlannerAllId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerAllId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerAllId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PlannerDeltaId, ok = parsed.Parsed["plannerDeltaId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerDeltaId", *parsed)
	}

	return &id, nil
}

// ParseUserPlannerAllIDInsensitively parses 'input' case-insensitively into a UserPlannerAllId
// note: this method should only be used for API response data and not user input
func ParseUserPlannerAllIDInsensitively(input string) (*UserPlannerAllId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerAllId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerAllId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PlannerDeltaId, ok = parsed.Parsed["plannerDeltaId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerDeltaId", *parsed)
	}

	return &id, nil
}

// ValidateUserPlannerAllID checks that 'input' can be parsed as a User Planner All ID
func ValidateUserPlannerAllID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPlannerAllID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Planner All ID
func (id UserPlannerAllId) ID() string {
	fmtString := "/users/%s/planner/all/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerDeltaId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Planner All ID
func (id UserPlannerAllId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPlanner", "planner", "planner"),
		resourceids.StaticSegment("staticAll", "all", "all"),
		resourceids.UserSpecifiedSegment("plannerDeltaId", "plannerDeltaIdValue"),
	}
}

// String returns a human-readable description of this User Planner All ID
func (id UserPlannerAllId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Delta: %q", id.PlannerDeltaId),
	}
	return fmt.Sprintf("User Planner All (%s)", strings.Join(components, "\n"))
}
