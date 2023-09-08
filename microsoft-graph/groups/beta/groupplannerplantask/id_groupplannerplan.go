package groupplannerplantask

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupPlannerPlanId{}

// GroupPlannerPlanId is a struct representing the Resource ID for a Group Planner Plan
type GroupPlannerPlanId struct {
	GroupId       string
	PlannerPlanId string
}

// NewGroupPlannerPlanID returns a new GroupPlannerPlanId struct
func NewGroupPlannerPlanID(groupId string, plannerPlanId string) GroupPlannerPlanId {
	return GroupPlannerPlanId{
		GroupId:       groupId,
		PlannerPlanId: plannerPlanId,
	}
}

// ParseGroupPlannerPlanID parses 'input' into a GroupPlannerPlanId
func ParseGroupPlannerPlanID(input string) (*GroupPlannerPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupPlannerPlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupPlannerPlanId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	return &id, nil
}

// ParseGroupPlannerPlanIDInsensitively parses 'input' case-insensitively into a GroupPlannerPlanId
// note: this method should only be used for API response data and not user input
func ParseGroupPlannerPlanIDInsensitively(input string) (*GroupPlannerPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupPlannerPlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupPlannerPlanId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	return &id, nil
}

// ValidateGroupPlannerPlanID checks that 'input' can be parsed as a Group Planner Plan ID
func ValidateGroupPlannerPlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupPlannerPlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Planner Plan ID
func (id GroupPlannerPlanId) ID() string {
	fmtString := "/groups/%s/planner/plans/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.PlannerPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Planner Plan ID
func (id GroupPlannerPlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticPlanner", "planner", "planner"),
		resourceids.StaticSegment("staticPlans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanIdValue"),
	}
}

// String returns a human-readable description of this Group Planner Plan ID
func (id GroupPlannerPlanId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
	}
	return fmt.Sprintf("Group Planner Plan (%s)", strings.Join(components, "\n"))
}
