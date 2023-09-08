package groupplannerplanbuckettask

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupPlannerPlanBucketId{}

// GroupPlannerPlanBucketId is a struct representing the Resource ID for a Group Planner Plan Bucket
type GroupPlannerPlanBucketId struct {
	GroupId         string
	PlannerPlanId   string
	PlannerBucketId string
}

// NewGroupPlannerPlanBucketID returns a new GroupPlannerPlanBucketId struct
func NewGroupPlannerPlanBucketID(groupId string, plannerPlanId string, plannerBucketId string) GroupPlannerPlanBucketId {
	return GroupPlannerPlanBucketId{
		GroupId:         groupId,
		PlannerPlanId:   plannerPlanId,
		PlannerBucketId: plannerBucketId,
	}
}

// ParseGroupPlannerPlanBucketID parses 'input' into a GroupPlannerPlanBucketId
func ParseGroupPlannerPlanBucketID(input string) (*GroupPlannerPlanBucketId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupPlannerPlanBucketId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupPlannerPlanBucketId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	if id.PlannerBucketId, ok = parsed.Parsed["plannerBucketId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerBucketId", *parsed)
	}

	return &id, nil
}

// ParseGroupPlannerPlanBucketIDInsensitively parses 'input' case-insensitively into a GroupPlannerPlanBucketId
// note: this method should only be used for API response data and not user input
func ParseGroupPlannerPlanBucketIDInsensitively(input string) (*GroupPlannerPlanBucketId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupPlannerPlanBucketId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupPlannerPlanBucketId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	if id.PlannerBucketId, ok = parsed.Parsed["plannerBucketId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerBucketId", *parsed)
	}

	return &id, nil
}

// ValidateGroupPlannerPlanBucketID checks that 'input' can be parsed as a Group Planner Plan Bucket ID
func ValidateGroupPlannerPlanBucketID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupPlannerPlanBucketID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Planner Plan Bucket ID
func (id GroupPlannerPlanBucketId) ID() string {
	fmtString := "/groups/%s/planner/plans/%s/buckets/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.PlannerPlanId, id.PlannerBucketId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Planner Plan Bucket ID
func (id GroupPlannerPlanBucketId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticPlanner", "planner", "planner"),
		resourceids.StaticSegment("staticPlans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanIdValue"),
		resourceids.StaticSegment("staticBuckets", "buckets", "buckets"),
		resourceids.UserSpecifiedSegment("plannerBucketId", "plannerBucketIdValue"),
	}
}

// String returns a human-readable description of this Group Planner Plan Bucket ID
func (id GroupPlannerPlanBucketId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Bucket: %q", id.PlannerBucketId),
	}
	return fmt.Sprintf("Group Planner Plan Bucket (%s)", strings.Join(components, "\n"))
}
