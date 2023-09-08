package userplannerplanbucket

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPlannerPlanBucketId{}

// UserPlannerPlanBucketId is a struct representing the Resource ID for a User Planner Plan Bucket
type UserPlannerPlanBucketId struct {
	UserId          string
	PlannerPlanId   string
	PlannerBucketId string
}

// NewUserPlannerPlanBucketID returns a new UserPlannerPlanBucketId struct
func NewUserPlannerPlanBucketID(userId string, plannerPlanId string, plannerBucketId string) UserPlannerPlanBucketId {
	return UserPlannerPlanBucketId{
		UserId:          userId,
		PlannerPlanId:   plannerPlanId,
		PlannerBucketId: plannerBucketId,
	}
}

// ParseUserPlannerPlanBucketID parses 'input' into a UserPlannerPlanBucketId
func ParseUserPlannerPlanBucketID(input string) (*UserPlannerPlanBucketId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerPlanBucketId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerPlanBucketId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	if id.PlannerBucketId, ok = parsed.Parsed["plannerBucketId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerBucketId", *parsed)
	}

	return &id, nil
}

// ParseUserPlannerPlanBucketIDInsensitively parses 'input' case-insensitively into a UserPlannerPlanBucketId
// note: this method should only be used for API response data and not user input
func ParseUserPlannerPlanBucketIDInsensitively(input string) (*UserPlannerPlanBucketId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPlannerPlanBucketId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPlannerPlanBucketId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	if id.PlannerBucketId, ok = parsed.Parsed["plannerBucketId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerBucketId", *parsed)
	}

	return &id, nil
}

// ValidateUserPlannerPlanBucketID checks that 'input' can be parsed as a User Planner Plan Bucket ID
func ValidateUserPlannerPlanBucketID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPlannerPlanBucketID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Planner Plan Bucket ID
func (id UserPlannerPlanBucketId) ID() string {
	fmtString := "/users/%s/planner/plans/%s/buckets/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PlannerPlanId, id.PlannerBucketId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Planner Plan Bucket ID
func (id UserPlannerPlanBucketId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPlanner", "planner", "planner"),
		resourceids.StaticSegment("staticPlans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanIdValue"),
		resourceids.StaticSegment("staticBuckets", "buckets", "buckets"),
		resourceids.UserSpecifiedSegment("plannerBucketId", "plannerBucketIdValue"),
	}
}

// String returns a human-readable description of this User Planner Plan Bucket ID
func (id UserPlannerPlanBucketId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Bucket: %q", id.PlannerBucketId),
	}
	return fmt.Sprintf("User Planner Plan Bucket (%s)", strings.Join(components, "\n"))
}
