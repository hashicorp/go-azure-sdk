package mependingaccessreviewinstancestagedecisioninstancedecision

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePendingAccessReviewInstanceStageDecisionInstanceDecisionId{}

// MePendingAccessReviewInstanceStageDecisionInstanceDecisionId is a struct representing the Resource ID for a Me Pending Access Review Instance Stage Decision Instance Decision
type MePendingAccessReviewInstanceStageDecisionInstanceDecisionId struct {
	AccessReviewInstanceId              string
	AccessReviewStageId                 string
	AccessReviewInstanceDecisionItemId  string
	AccessReviewInstanceDecisionItemId1 string
}

// NewMePendingAccessReviewInstanceStageDecisionInstanceDecisionID returns a new MePendingAccessReviewInstanceStageDecisionInstanceDecisionId struct
func NewMePendingAccessReviewInstanceStageDecisionInstanceDecisionID(accessReviewInstanceId string, accessReviewStageId string, accessReviewInstanceDecisionItemId string, accessReviewInstanceDecisionItemId1 string) MePendingAccessReviewInstanceStageDecisionInstanceDecisionId {
	return MePendingAccessReviewInstanceStageDecisionInstanceDecisionId{
		AccessReviewInstanceId:              accessReviewInstanceId,
		AccessReviewStageId:                 accessReviewStageId,
		AccessReviewInstanceDecisionItemId:  accessReviewInstanceDecisionItemId,
		AccessReviewInstanceDecisionItemId1: accessReviewInstanceDecisionItemId1,
	}
}

// ParseMePendingAccessReviewInstanceStageDecisionInstanceDecisionID parses 'input' into a MePendingAccessReviewInstanceStageDecisionInstanceDecisionId
func ParseMePendingAccessReviewInstanceStageDecisionInstanceDecisionID(input string) (*MePendingAccessReviewInstanceStageDecisionInstanceDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceStageDecisionInstanceDecisionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceStageDecisionInstanceDecisionId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId1, ok = parsed.Parsed["accessReviewInstanceDecisionItemId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId1", *parsed)
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceStageDecisionInstanceDecisionIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceStageDecisionInstanceDecisionId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceStageDecisionInstanceDecisionIDInsensitively(input string) (*MePendingAccessReviewInstanceStageDecisionInstanceDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MePendingAccessReviewInstanceStageDecisionInstanceDecisionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MePendingAccessReviewInstanceStageDecisionInstanceDecisionId{}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = parsed.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", *parsed)
	}

	if id.AccessReviewInstanceDecisionItemId1, ok = parsed.Parsed["accessReviewInstanceDecisionItemId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId1", *parsed)
	}

	return &id, nil
}

// ValidateMePendingAccessReviewInstanceStageDecisionInstanceDecisionID checks that 'input' can be parsed as a Me Pending Access Review Instance Stage Decision Instance Decision ID
func ValidateMePendingAccessReviewInstanceStageDecisionInstanceDecisionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceStageDecisionInstanceDecisionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Stage Decision Instance Decision ID
func (id MePendingAccessReviewInstanceStageDecisionInstanceDecisionId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/stages/%s/decisions/%s/instance/decisions/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewInstanceDecisionItemId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Stage Decision Instance Decision ID
func (id MePendingAccessReviewInstanceStageDecisionInstanceDecisionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticStages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageIdValue"),
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemIdValue"),
		resourceids.StaticSegment("staticInstance", "instance", "instance"),
		resourceids.StaticSegment("staticDecisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId1", "accessReviewInstanceDecisionItemId1Value"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Stage Decision Instance Decision ID
func (id MePendingAccessReviewInstanceStageDecisionInstanceDecisionId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Instance Decision Item Id 1: %q", id.AccessReviewInstanceDecisionItemId1),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Stage Decision Instance Decision (%s)", strings.Join(components, "\n"))
}
