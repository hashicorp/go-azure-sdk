package meplannerplanbucket

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePlannerPlanBucketId{}

// MePlannerPlanBucketId is a struct representing the Resource ID for a Me Planner Plan Bucket
type MePlannerPlanBucketId struct {
	PlannerPlanId   string
	PlannerBucketId string
}

// NewMePlannerPlanBucketID returns a new MePlannerPlanBucketId struct
func NewMePlannerPlanBucketID(plannerPlanId string, plannerBucketId string) MePlannerPlanBucketId {
	return MePlannerPlanBucketId{
		PlannerPlanId:   plannerPlanId,
		PlannerBucketId: plannerBucketId,
	}
}

// ParseMePlannerPlanBucketID parses 'input' into a MePlannerPlanBucketId
func ParseMePlannerPlanBucketID(input string) (*MePlannerPlanBucketId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePlannerPlanBucketId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePlannerPlanBucketId{}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	if id.PlannerBucketId, ok = parsed.Parsed["plannerBucketId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerBucketId", *parsed)
	}

	return &id, nil
}

// ParseMePlannerPlanBucketIDInsensitively parses 'input' case-insensitively into a MePlannerPlanBucketId
// note: this method should only be used for API response data and not user input
func ParseMePlannerPlanBucketIDInsensitively(input string) (*MePlannerPlanBucketId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePlannerPlanBucketId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePlannerPlanBucketId{}

	if id.PlannerPlanId, ok = parsed.Parsed["plannerPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerPlanId", *parsed)
	}

	if id.PlannerBucketId, ok = parsed.Parsed["plannerBucketId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "plannerBucketId", *parsed)
	}

	return &id, nil
}

// ValidateMePlannerPlanBucketID checks that 'input' can be parsed as a Me Planner Plan Bucket ID
func ValidateMePlannerPlanBucketID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePlannerPlanBucketID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Planner Plan Bucket ID
func (id MePlannerPlanBucketId) ID() string {
	fmtString := "/me/planner/plans/%s/buckets/%s"
	return fmt.Sprintf(fmtString, id.PlannerPlanId, id.PlannerBucketId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Planner Plan Bucket ID
func (id MePlannerPlanBucketId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticPlanner", "planner", "planner"),
		resourceids.StaticSegment("staticPlans", "plans", "plans"),
		resourceids.UserSpecifiedSegment("plannerPlanId", "plannerPlanIdValue"),
		resourceids.StaticSegment("staticBuckets", "buckets", "buckets"),
		resourceids.UserSpecifiedSegment("plannerBucketId", "plannerBucketIdValue"),
	}
}

// String returns a human-readable description of this Me Planner Plan Bucket ID
func (id MePlannerPlanBucketId) String() string {
	components := []string{
		fmt.Sprintf("Planner Plan: %q", id.PlannerPlanId),
		fmt.Sprintf("Planner Bucket: %q", id.PlannerBucketId),
	}
	return fmt.Sprintf("Me Planner Plan Bucket (%s)", strings.Join(components, "\n"))
}
