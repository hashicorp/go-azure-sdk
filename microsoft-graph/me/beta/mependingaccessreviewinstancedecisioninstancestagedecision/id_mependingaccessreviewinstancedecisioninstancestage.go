package mependingaccessreviewinstancedecisioninstancestagedecision

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePendingAccessReviewInstanceDecisionInstanceStageId{}

// MePendingAccessReviewInstanceDecisionInstanceStageId is a struct representing the Resource ID for a Me Pending Access Review Instance Decision Instance Stage
type MePendingAccessReviewInstanceDecisionInstanceStageId struct {
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
	AccessReviewStageId                string
}

// NewMePendingAccessReviewInstanceDecisionInstanceStageID returns a new MePendingAccessReviewInstanceDecisionInstanceStageId struct
func NewMePendingAccessReviewInstanceDecisionInstanceStageID(accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, accessReviewStageId string) MePendingAccessReviewInstanceDecisionInstanceStageId {
	return MePendingAccessReviewInstanceDecisionInstanceStageId{
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		AccessReviewStageId:                accessReviewStageId,
	}
}

// ParseMePendingAccessReviewInstanceDecisionInstanceStageID parses 'input' into a MePendingAccessReviewInstanceDecisionInstanceStageId
func ParseMePendingAccessReviewInstanceDecisionInstanceStageID(input string) (*MePendingAccessReviewInstanceDecisionInstanceStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceDecisionInstanceStageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceDecisionInstanceStageId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceDecisionInstanceStageIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceDecisionInstanceStageId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceDecisionInstanceStageIDInsensitively(input string) (*MePendingAccessReviewInstanceDecisionInstanceStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceDecisionInstanceStageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceDecisionInstanceStageId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	return &id, nil
}

// ValidateMePendingAccessReviewInstanceDecisionInstanceStageID checks that 'input' can be parsed as a Me Pending Access Review Instance Decision Instance Stage ID
func ValidateMePendingAccessReviewInstanceDecisionInstanceStageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceDecisionInstanceStageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Decision Instance Stage ID
func (id MePendingAccessReviewInstanceDecisionInstanceStageId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/decisions/%s/instance/stages/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewStageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Decision Instance Stage ID
func (id MePendingAccessReviewInstanceDecisionInstanceStageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemIdValue"),
		resourceids.StaticSegment("staticInstance", "instance", "instance"),
		resourceids.StaticSegment("staticStages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageIdValue"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Decision Instance Stage ID
func (id MePendingAccessReviewInstanceDecisionInstanceStageId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Decision Instance Stage (%s)", strings.Join(components, "\n"))
}
