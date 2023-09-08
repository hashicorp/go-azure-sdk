package mependingaccessreviewinstancestagedecisioninstance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePendingAccessReviewInstanceStageDecisionId{}

// MePendingAccessReviewInstanceStageDecisionId is a struct representing the Resource ID for a Me Pending Access Review Instance Stage Decision
type MePendingAccessReviewInstanceStageDecisionId struct {
	AccessReviewInstanceId             string
	AccessReviewStageId                string
	AccessReviewInstanceDecisionItemId string
}

// NewMePendingAccessReviewInstanceStageDecisionID returns a new MePendingAccessReviewInstanceStageDecisionId struct
func NewMePendingAccessReviewInstanceStageDecisionID(accessReviewInstanceId string, accessReviewStageId string, accessReviewInstanceDecisionItemId string) MePendingAccessReviewInstanceStageDecisionId {
	return MePendingAccessReviewInstanceStageDecisionId{
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewStageId:                accessReviewStageId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
	}
}

// ParseMePendingAccessReviewInstanceStageDecisionID parses 'input' into a MePendingAccessReviewInstanceStageDecisionId
func ParseMePendingAccessReviewInstanceStageDecisionID(input string) (*MePendingAccessReviewInstanceStageDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceStageDecisionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceStageDecisionId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceStageDecisionIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceStageDecisionId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceStageDecisionIDInsensitively(input string) (*MePendingAccessReviewInstanceStageDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceStageDecisionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceStageDecisionId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	return &id, nil
}

// ValidateMePendingAccessReviewInstanceStageDecisionID checks that 'input' can be parsed as a Me Pending Access Review Instance Stage Decision ID
func ValidateMePendingAccessReviewInstanceStageDecisionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceStageDecisionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Stage Decision ID
func (id MePendingAccessReviewInstanceStageDecisionId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/stages/%s/decisions/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Stage Decision ID
func (id MePendingAccessReviewInstanceStageDecisionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticStages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageIdValue"),
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemIdValue"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Stage Decision ID
func (id MePendingAccessReviewInstanceStageDecisionId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Stage Decision (%s)", strings.Join(components, "\n"))
}
