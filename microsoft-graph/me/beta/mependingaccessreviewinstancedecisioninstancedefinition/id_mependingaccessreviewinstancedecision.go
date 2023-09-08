package mependingaccessreviewinstancedecisioninstancedefinition

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePendingAccessReviewInstanceDecisionId{}

// MePendingAccessReviewInstanceDecisionId is a struct representing the Resource ID for a Me Pending Access Review Instance Decision
type MePendingAccessReviewInstanceDecisionId struct {
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
}

// NewMePendingAccessReviewInstanceDecisionID returns a new MePendingAccessReviewInstanceDecisionId struct
func NewMePendingAccessReviewInstanceDecisionID(accessReviewInstanceId string, accessReviewInstanceDecisionItemId string) MePendingAccessReviewInstanceDecisionId {
	return MePendingAccessReviewInstanceDecisionId{
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
	}
}

// ParseMePendingAccessReviewInstanceDecisionID parses 'input' into a MePendingAccessReviewInstanceDecisionId
func ParseMePendingAccessReviewInstanceDecisionID(input string) (*MePendingAccessReviewInstanceDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceDecisionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceDecisionId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceDecisionIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceDecisionId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceDecisionIDInsensitively(input string) (*MePendingAccessReviewInstanceDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceDecisionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceDecisionId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	return &id, nil
}

// ValidateMePendingAccessReviewInstanceDecisionID checks that 'input' can be parsed as a Me Pending Access Review Instance Decision ID
func ValidateMePendingAccessReviewInstanceDecisionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceDecisionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Decision ID
func (id MePendingAccessReviewInstanceDecisionId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/decisions/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Decision ID
func (id MePendingAccessReviewInstanceDecisionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemIdValue"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Decision ID
func (id MePendingAccessReviewInstanceDecisionId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Decision (%s)", strings.Join(components, "\n"))
}
